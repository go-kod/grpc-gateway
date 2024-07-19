package server

import (
	"context"
	"slices"
	"strconv"
	"strings"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod/ext/registry"
	"github.com/go-kod/kod/interceptor"
	"github.com/go-kod/kod/interceptor/kcircuitbreaker"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/samber/lo"
	"github.com/sysulq/graphql-grpc-gateway/internal/config"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"
)

type caller struct {
	kod.Implements[Caller]

	reflection kod.Ref[Reflection]
	config     kod.Ref[config.Config]

	serviceStub  map[string]grpcdynamic.Stub
	singleflight singleflight.Group
	descs        []*desc.FileDescriptor
}

func (c *caller) Init(ctx context.Context) (err error) {
	config := c.config.Get().Config().Grpc

	serviceStub := map[string]grpcdynamic.Stub{}
	descs := make([]*desc.FileDescriptor, 0)
	descsconn := map[string]*grpc.ClientConn{}
	var etcd registry.Registry

	if len(c.config.Get().Config().Grpc.Etcd.Endpoints) > 0 {
		etcd = lo.Must(c.config.Get().Config().Grpc.Etcd.Build(ctx))
	}

	for _, service := range config.Services {

		if etcd != nil {
			service = service.WithRegistry(etcd)
		}

		conn := service.Build()

		newDescs, err := c.reflection.Get().ListPackages(ctx, conn)
		if err != nil {
			return err
		}

		for _, d := range newDescs {
			descsconn[d.GetFullyQualifiedName()] = conn
		}
		descs = append(descs, newDescs...)
	}

	for _, d := range descs {
		for _, svc := range d.GetServices() {
			serviceStub[svc.GetFullyQualifiedName()] = grpcdynamic.NewStub(descsconn[d.GetFullyQualifiedName()])
		}
	}

	descs = lo.UniqBy(descs, func(item *desc.FileDescriptor) string {
		return item.GetFile().GetFullyQualifiedName()
	})

	c.descs = descs
	c.serviceStub = serviceStub

	return nil
}

func (c *caller) GetDescs() []*desc.FileDescriptor {
	return c.descs
}

func (c *caller) Call(ctx context.Context, rpc *desc.MethodDescriptor, message protoadapt.MessageV1) (protoadapt.MessageV1, error) {
	if c.config.Get().Config().GraphQL.SingleFlight {
		if enable, ok := ctx.Value(allowSingleFlightKey).(bool); ok && enable {
			hash := Hash64.Get()
			defer Hash64.Put(hash)

			md, ok := metadata.FromOutgoingContext(ctx)
			if ok {
				hd := make([]string, 0, len(md))
				for k, v := range md {
					// skip grpc gateway prefixed metadata
					if strings.Contains(k, MetadataPrefix) {
						continue
					}
					hd = append(hd, k+strings.Join(v, ","))
				}
				slices.Sort(hd)
				for _, v := range hd {
					_, err := hash.Write([]byte(v))
					if err != nil {
						return nil, err
					}
				}
			}

			msg, err := proto.Marshal(protoadapt.MessageV2Of(message))
			if err != nil {
				return nil, err
			}

			// generate hash based on rpc pointer
			_, err = hash.Write([]byte(rpc.GetFullyQualifiedName()))
			if err != nil {
				return nil, err
			}
			_, err = hash.Write(msg)
			if err != nil {
				return nil, err
			}
			sum := hash.Sum64()
			key := strconv.FormatUint(sum, 10)

			res, err, _ := c.singleflight.Do(key, func() (interface{}, error) {
				return c.serviceStub[rpc.GetService().GetFullyQualifiedName()].InvokeRpc(ctx, rpc, message)
			})
			if err != nil {
				return nil, err
			}

			return res.(protoadapt.MessageV1), nil
		}
	}

	res, err := c.serviceStub[rpc.GetService().GetFullyQualifiedName()].InvokeRpc(ctx, rpc, message)
	return res, err
}

func (c *caller) Interceptors() []interceptor.Interceptor {
	if c.config.Get().Config().Engine.CircuitBreaker {
		return []interceptor.Interceptor{
			kcircuitbreaker.Interceptor(),
		}
	}
	return nil
}

var allowSingleFlightKey struct{}
