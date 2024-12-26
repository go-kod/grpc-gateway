package config

import (
	"context"
	"testing"
	"time"

	"github.com/go-kod/kod"
	"github.com/go-kod/kod-ext/client/kgrpc"
	"github.com/go-kod/kod-ext/core/pyroscope"
	"github.com/go-kod/kod-ext/registry/etcdv3"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	kod.RunTest(t, func(ctx context.Context, c Config) {
		require.Equal(t, &ConfigInfo{
			Engine: EngineConfig{
				RateLimit:      true,
				CircuitBreaker: true,

				Pyroscope: Pyroscope{
					Enable: true,
					Config: pyroscope.Config{ServerAddress: ""},
				},
			},
			Server: ServerConfig{
				GraphQL: GraphQLConfig{
					Address:                ":8080",
					Disable:                false,
					Playground:             true,
					GenerateUnboundMethods: true,
					QueryCache:             true,
					SingleFlight:           true,
				},
				HTTP: HTTPConfig{
					Address:      ":9090",
					SingleFlight: true,
				},
			},
			Grpc: Grpc{
				Etcd: etcdv3.Config{
					Endpoints: []string{"localhost:2379"},
					Timeout:   3 * time.Second,
					TTL:       60,
				},
				Services: []kgrpc.Config{
					{
						Target:  "etcd:///local/optionsserver/grpc",
						Timeout: time.Second,
					},
					{
						Target:  "etcd:///local/constructsserver/grpc",
						Timeout: time.Second,
					},
					{
						Target:  "etcd:///local/helloworld/grpc",
						Timeout: time.Second,
					},
				},
			},
		}, c.Config())
	}, kod.WithConfigFile("../../example/gateway/config.yaml"))
}
