package config

import (
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-ext/client/kgrpc"
	"github.com/go-kod/kod-ext/core/otel"
	"github.com/go-kod/kod-ext/core/pyroscope"
	"github.com/go-kod/kod-ext/registry/etcdv3"
)

type config struct {
	kod.Implements[Config]
	kod.WithGlobalConfig[ConfigInfo]
}

type ConfigInfo struct {
	Server ServerConfig
	Engine EngineConfig
	Grpc   Grpc
}

type Pyroscope struct {
	pyroscope.Config `koanf:",squash"`
	Enable           bool
}

type Jwt struct {
	Enable               bool
	LocalJwks            string
	ForwardPayloadHeader string
}

type JwtClaimToHeader struct {
	HeaderName string
	ClaimName  string
}

type EngineConfig struct {
	Otel           otel.Config
	Pyroscope      Pyroscope
	RateLimit      bool
	CircuitBreaker bool
}

type ServerConfig struct {
	GraphQL GraphQLConfig
	HTTP    HTTPConfig
}

type HTTPConfig struct {
	Address      string
	Disable      bool
	Jwt          Jwt
	SingleFlight bool
}

type GraphQLConfig struct {
	Address                string
	Disable                bool
	Playground             bool
	Jwt                    Jwt
	GenerateUnboundMethods bool
	QueryCache             bool
	SingleFlight           bool
}

type Grpc struct {
	Etcd     etcdv3.Config
	Services []kgrpc.Config
}

func (ins *config) Config() *ConfigInfo {
	return ins.WithGlobalConfig.Config()
}
