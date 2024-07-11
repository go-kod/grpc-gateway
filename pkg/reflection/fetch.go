package reflection

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/v2/grpcreflect"
	"github.com/jhump/protoreflect/v2/protowrap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

func isReflectionServiceName(name string) bool {
	return name == grpc_reflection_v1.ServerReflection_ServiceDesc.ServiceName ||
		name == grpc_reflection_v1alpha.ServerReflection_ServiceDesc.ServiceName
}

var ErrTLSHandshakeFailed = errors.New("TLS handshake failed")

// Client defines gRPC reflection client.
type Client interface {
	// ListPackages lists file descriptors from the gRPC reflection server.
	// ListPackages returns these errors:
	//   - ErrTLSHandshakeFailed: TLS misconfig.
	ListPackages() ([]protoreflect.FileDescriptor, error)
}

type client struct {
	client *grpcreflect.Client
}

// NewClient returns an instance of gRPC reflection client for gRPC protocol.
func NewClient(conn grpc.ClientConnInterface) Client {
	return &client{
		client: grpcreflect.NewClientAuto(context.Background(), conn),
	}
}

func (c *client) ListPackages() ([]protoreflect.FileDescriptor, error) {
	// c.client.FileContainingExtension()
	ssvcs, err := c.client.ListServices()
	if err != nil {
		msg := status.Convert(err).Message()
		// Check whether the error message contains TLS related error.
		// If the server didn't enable TLS, the error message contains the first string.
		// If Evans didn't enable TLS against to the TLS enabled server, the error message contains
		// the second string.
		if strings.Contains(msg, "tls: first record does not look like a TLS handshake") ||
			strings.Contains(msg, "latest connection error: <nil>") {
			return nil, ErrTLSHandshakeFailed
		}
		return nil, fmt.Errorf("failed to list services from reflecton enabled gRPC server: %w", err)
	}

	var fds []protoreflect.FileDescriptor
	for _, s := range ssvcs {
		if isReflectionServiceName(string(s)) {
			continue
		}
		svc, err := c.client.FileContainingSymbol(s)
		if err != nil {
			return nil, err
		}

		fds = append(fds, svc)
	}
	return fds, nil
}

func (c *client) Reset() {
	c.client.Reset()
}

func NewClientWithImportsResolver(conn grpc.ClientConnInterface) Client {
	return &clientV2{client: grpc_reflection_v1.NewServerReflectionClient(conn), fileMap: map[*descriptorpb.FileDescriptorProto]struct{}{}, mu: &sync.RWMutex{}}
}

type clientV2 struct {
	client  grpc_reflection_v1.ServerReflectionClient
	fileMap map[*descriptorpb.FileDescriptorProto]struct{}
	mu      *sync.RWMutex
}

func (c *clientV2) ListPackages() (descriptors []protoreflect.FileDescriptor, err error) {
	ctx := context.Background()
	stream, err := c.client.ServerReflectionInfo(ctx)
	if err != nil {
		return nil, err
	}
	err = stream.Send(&grpc_reflection_v1.ServerReflectionRequest{
		MessageRequest: &grpc_reflection_v1.ServerReflectionRequest_ListServices{
			ListServices: "*",
		},
	})
	if err != nil {
		return nil, err
	}

	request, response, filebus := NewFileBus()
	defer close(request)
	defer close(response)

	errCh := make(chan error)
	go func() {
		for req := range request {
			if err := stream.Send(req); err != nil {
				errCh <- err
				return
			}
		}
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					errCh <- errors.New("connection with reflect server closed before retrieving all data")
					return
				}
				errCh <- err
				return
			}

			if res.GetListServicesResponse() != nil {
				// TODO move this logic to the bus
				go func() {
					// TODO stop signal when this for ends
					for _, svc := range res.GetListServicesResponse().Service {
						if isReflectionServiceName(svc.GetName()) {
							continue
						}
						files, err := filebus.GetFilesForSymbol(svc.GetName())
						if err != nil {
							errCh <- err
							return
						}
						for _, file := range files {
							if err := c.processFile(file, filebus); err != nil {
								errCh <- err
								return
							}
						}
					}

					errCh <- nil
				}()
				continue
			}

			select {
			case response <- res:
			case _, ok := <-errCh:
				if !ok {
					return
				}
			}
		}
	}()
	err, ok := <-errCh
	if ok && err == nil {
		var descriptors []protoreflect.FileDescriptor
		c.mu.RLock()
		for f := range c.fileMap {
			desc, err := protowrap.FromFileDescriptorProto(f, protoregistry.GlobalFiles)
			if err != nil {
				return nil, err
			}
			descriptors = append(descriptors, desc)
		}
		c.mu.RUnlock()

		return descriptors, nil
	}
	return nil, err
}

func (c *clientV2) processFile(file *descriptorpb.FileDescriptorProto, filebus FileBus) error {
	c.mu.RLock()
	_, ok := c.fileMap[file]
	c.mu.RUnlock()
	if ok {
		return nil
	}
	c.mu.Lock()
	c.fileMap[file] = struct{}{}
	c.mu.Unlock()
	file.Name = proto.String(file.GetPackage() + "/" + file.GetName())
	file.Dependency = nil
	deps := map[*descriptorpb.FileDescriptorProto]struct{}{}

	resolveDependencies := func(ff []*descriptorpb.FileDescriptorProto, err error) error {
		if err != nil {
			return err
		}
		for _, f := range ff {
			if err := c.processFile(f, filebus); err != nil {
				return err
			}
			if _, ok := deps[f]; ok || f == file {
				continue
			}
			deps[f] = struct{}{}
			file.Dependency = append(file.Dependency, f.GetName())
		}
		return nil
	}
	getFilesForOptions := func(op interface {
		proto.Message
		ProtoReflect() protoreflect.Message
	}, dd string,
	) error {
		exx, err := proto.ExtensionDescs(op)
		if err != nil {
			return nil
		}
		for _, e := range exx {
			if err := resolveDependencies(filebus.GetFilesForExtension(
				string(op.ProtoReflect().Descriptor().FullName()),
				int32(e.TypeDescriptor().Number()))); err != nil {
				return err
			}
		}
		return nil
	}
	getFilesForSymbol := func(symbol, dd string) error {
		return resolveDependencies(filebus.GetFilesForSymbol(symbol))
	}
	if err := getFilesForOptions(file.GetOptions(), file.GetName()); err != nil {
		return err
	}
	for _, svc := range file.GetService() {
		if err := getFilesForOptions(svc.GetOptions(), svc.GetName()); err != nil {
			return err
		}
		for _, rpc := range svc.GetMethod() {
			if err := getFilesForOptions(rpc.GetOptions(), rpc.GetName()); err != nil {
				return err
			}
			if err := getFilesForSymbol(rpc.GetInputType(), rpc.GetName()); err != nil {
				return err
			}
			if err := getFilesForSymbol(rpc.GetOutputType(), rpc.GetName()); err != nil {
				return err
			}
		}
	}
	for _, ex := range file.GetExtension() {
		if err := getFilesForSymbol(ex.GetExtendee(), ex.GetName()); err != nil {
			return err
		}
	}
	for _, m := range file.GetMessageType() {
		if err := getFilesForOptions(m.GetOptions(), m.GetName()); err != nil {
			return err
		}
		for _, f := range m.GetField() {
			if err := getFilesForOptions(f.GetOptions(), f.GetName()); err != nil {
				return err
			}
			switch f.GetType() {
			case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, descriptorpb.FieldDescriptorProto_TYPE_ENUM:
				if err := getFilesForSymbol(f.GetTypeName(), f.GetName()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FileBus interface {
	GetFilesForSymbol(symbol string) (files []*descriptorpb.FileDescriptorProto, err error)
	GetFilesForExtension(symbol string, extensionNr int32) (files []*descriptorpb.FileDescriptorProto, err error)
}

func NewFileBus() (chan *grpc_reflection_v1.ServerReflectionRequest, chan *grpc_reflection_v1.ServerReflectionResponse, FileBus) {
	request := make(chan *grpc_reflection_v1.ServerReflectionRequest)
	response := make(chan *grpc_reflection_v1.ServerReflectionResponse)
	fb := &fileBus{
		symbols:       map[string]*descriptorpb.FileDescriptorProto{},
		extensions:    map[extensionKey]*descriptorpb.FileDescriptorProto{},
		subscriptions: map[*grpc_reflection_v1.ServerReflectionRequest]chan *grpc_reflection_v1.ServerReflectionResponse{},
		request:       request,
	}
	go func() {
		for {
			select {
			case res, ok := <-response:
				if !ok {
					for _, sub := range fb.subscriptions {
						close(sub)
					}
					return
				}
				for _, sub := range fb.subscriptions {
					sub <- res
				}
			}
		}
	}()

	return request, response, fb
}

type fileBus struct {
	symbols       map[string]*descriptorpb.FileDescriptorProto
	extensions    map[extensionKey]*descriptorpb.FileDescriptorProto
	subscriptions map[*grpc_reflection_v1.ServerReflectionRequest]chan *grpc_reflection_v1.ServerReflectionResponse
	request       chan *grpc_reflection_v1.ServerReflectionRequest
}

func (f *fileBus) getFiles(request *grpc_reflection_v1.ServerReflectionRequest) (files []*descriptorpb.FileDescriptorProto, err error) {
	subscription := make(chan *grpc_reflection_v1.ServerReflectionResponse)
	f.subscriptions[request] = subscription
	f.request <- request
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	for {
		select {
		case res, ok := <-subscription:
			if !ok {
				return
			}
			if proto.Equal(res.GetOriginalRequest(), request) {
				delete(f.subscriptions, request)
				close(subscription)
				if res.GetErrorResponse() != nil {
					log.Println("WARNING", fmt.Errorf("error status %d; msg: %s",
						res.GetErrorResponse().GetErrorCode(), res.GetErrorResponse().GetErrorMessage()))
				}
				for _, fileBytes := range res.GetFileDescriptorResponse().GetFileDescriptorProto() {
					file := &descriptorpb.FileDescriptorProto{}
					if err := proto.Unmarshal(fileBytes, file); err != nil {
						return nil, err
					}
					files = append(files, file)
				}
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("reflection server takes too long to respond %w", ctx.Err())
		}
	}
}

func (f *fileBus) GetFilesForSymbol(symbol string) (files []*descriptorpb.FileDescriptorProto, err error) {
	symbol = strings.TrimPrefix(symbol, ".")
	file, ok := f.symbols[symbol]
	if ok {
		return []*descriptorpb.FileDescriptorProto{file}, nil
	}
	defer func() { f.addToCache(files) }()
	request := &grpc_reflection_v1.ServerReflectionRequest{MessageRequest: &grpc_reflection_v1.ServerReflectionRequest_FileContainingSymbol{
		FileContainingSymbol: symbol,
	}}
	return f.getFiles(request)
}

func (f *fileBus) GetFilesForExtension(symbol string, extensionNr int32) (files []*descriptorpb.FileDescriptorProto, err error) {
	symbol = strings.TrimPrefix(symbol, ".")
	file, ok := f.extensions[extensionKey{symbol, extensionNr}]
	if ok {
		return []*descriptorpb.FileDescriptorProto{file}, nil
	}
	defer func() { f.addToCache(files) }()
	request := &grpc_reflection_v1.ServerReflectionRequest{MessageRequest: &grpc_reflection_v1.ServerReflectionRequest_FileContainingExtension{
		FileContainingExtension: &grpc_reflection_v1.ExtensionRequest{
			ContainingType:  symbol,
			ExtensionNumber: extensionNr,
		},
	}}

	return f.getFiles(request)
}

type extensionKey struct {
	symbol      string
	extensionNr int32
}

func (f *fileBus) addMessagesAndEnums(file *descriptorpb.FileDescriptorProto, msg *descriptorpb.DescriptorProto, parents ...string) {
	symbol := file.GetPackage() + "." + strings.Join(parents, ".")
	f.symbols[symbol] = file
	for _, nt := range msg.GetNestedType() {
		f.addMessagesAndEnums(file, nt, append(parents, nt.GetName())...)
	}
	for _, nt := range msg.GetEnumType() {
		symbol := file.GetPackage() + "." + strings.Join(parents, ".") + "." + nt.GetName()
		f.symbols[symbol] = file
	}
	return
}

func (f *fileBus) addToCache(files []*descriptorpb.FileDescriptorProto) {
	for _, fileDesc := range files {
		for _, e := range fileDesc.GetExtension() {
			key := extensionKey{strings.TrimPrefix(e.GetExtendee(), "."), e.GetNumber()}
			f.extensions[key] = fileDesc
		}
		for _, m := range fileDesc.GetMessageType() {
			f.addMessagesAndEnums(fileDesc, m, m.GetName())
		}
		for _, e := range fileDesc.GetEnumType() {
			symbol := fileDesc.GetPackage() + "." + e.GetName()
			f.symbols[symbol] = fileDesc
		}
		for _, s := range fileDesc.GetService() {
			symbol := fileDesc.GetPackage() + "." + s.GetName()
			f.symbols[symbol] = fileDesc
		}
	}
}
