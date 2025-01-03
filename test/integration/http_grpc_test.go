package integration

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kod/grpc-gateway/internal/config"
	"github.com/go-kod/grpc-gateway/internal/server"
	"github.com/go-kod/grpc-gateway/test"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod-ext/client/kgrpc"
	"github.com/go-kod/kod-ext/registry/etcdv3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHTTP2Grpc(t *testing.T) {
	infos := test.SetupDeps(t)

	mockConfig := config.NewMockConfig(gomock.NewController(t))
	mockConfig.EXPECT().Config().Return(&config.ConfigInfo{
		Engine: config.EngineConfig{
			RateLimit:      true,
			CircuitBreaker: true,
		},
		Grpc: config.Grpc{
			Etcd: etcdv3.Config{
				Endpoints: []string{"localhost:2379"},
			},
			Services: []kgrpc.Config{
				{
					Target: infos.ConstructsServerAddr.Addr().String(),
				},
				{
					Target: infos.OptionsServerAddr.Addr().String(),
				},
				{
					Target: infos.HelloworldServerAddr.Addr().String(),
				},
			},
		},
		Server: config.ServerConfig{
			GraphQL: config.GraphQLConfig{
				Playground:             true,
				GenerateUnboundMethods: true,
				SingleFlight:           true,
				QueryCache:             true,
			},
		},
	}).AnyTimes()

	t.Run("http to grpc", func(t *testing.T) {
		kod.RunTest(t, func(ctx context.Context, up server.HttpUpstream) {
			router := http.NewServeMux()
			up.Register(ctx, router)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/say/bob", nil)
			req.Header.Set("Content-Type", "application/json; charset=utf-8")

			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "application/json; charset=utf-8", rec.Header().Get("Content-Type"))
			assert.Equal(t, "{\"message\":\"Hello bob\"}", rec.Body.String())
		}, kod.WithFakes(kod.Fake[config.Config](mockConfig)))
	})

	t.Run("not found", func(t *testing.T) {
		kod.RunTest(t, func(ctx context.Context, up server.HttpUpstream) {
			router := http.NewServeMux()
			up.Register(ctx, router)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/say-notfound", nil)
			req.Header.Set("Content-Type", "application/json; charset=utf-8")

			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "text/plain; charset=utf-8", rec.Header().Get("Content-Type"))
			assert.Equal(t, "404 page not found\n", rec.Body.String())
		}, kod.WithFakes(kod.Fake[config.Config](mockConfig)))
	})

	t.Run("error", func(t *testing.T) {
		kod.RunTest(t, func(ctx context.Context, up server.HttpUpstream) {
			router := http.NewServeMux()
			up.Register(ctx, router)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/say/error", nil)
			req.Header.Set("Content-Type", "application/json; charset=utf-8")

			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "application/json; charset=utf-8", rec.Header().Get("Content-Type"))
			assert.Equal(t, `{"code":2,"message":"error","details":[]}`, rec.Body.String())
		}, kod.WithFakes(kod.Fake[config.Config](mockConfig)))
	})

	t.Run("http body", func(t *testing.T) {
		kod.RunTest(t, func(ctx context.Context, up server.HttpUpstream) {
			router := http.NewServeMux()
			up.Register(ctx, router)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/say/sam", bytes.NewBufferString("{\"name\":\"bob\"}"))
			req.Header.Set("Content-Type", "application/json; charset=utf-8")

			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "application/json; charset=utf-8", rec.Header().Get("Content-Type"))
			assert.Equal(t, `{"message":"Hello sam"}`, rec.Body.String())
		}, kod.WithFakes(kod.Fake[config.Config](mockConfig)))
	})

	t.Run("invalid http body", func(t *testing.T) {
		kod.RunTest(t, func(ctx context.Context, up server.HttpUpstream) {
			router := http.NewServeMux()
			up.Register(ctx, router)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/say/sam", bytes.NewBufferString("{invalid data}"))
			req.Header.Set("Content-Type", "application/json; charset=utf-8")

			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "text/plain; charset=utf-8", rec.Header().Get("Content-Type"))
			assert.Equal(t, "rpc error: code = InvalidArgument desc = invalid character 'i' looking for beginning of object key string\n", rec.Body.String())
		}, kod.WithFakes(kod.Fake[config.Config](mockConfig)))
	})
}

func TestHTTP2Grpc_Singleflight(t *testing.T) {
	infos := test.SetupDeps(t)

	t.Run("singleflight", func(t *testing.T) {
		mockConfig := config.NewMockConfig(gomock.NewController(t))
		mockConfig.EXPECT().Config().Return(&config.ConfigInfo{
			Grpc: config.Grpc{
				Etcd: etcdv3.Config{
					Endpoints: []string{"localhost:2379"},
				},
				Services: []kgrpc.Config{
					{
						Target: infos.ConstructsServerAddr.Addr().String(),
					},
					{
						Target: infos.OptionsServerAddr.Addr().String(),
					},
					{
						Target: infos.HelloworldServerAddr.Addr().String(),
					},
				},
			},
			Server: config.ServerConfig{
				HTTP: config.HTTPConfig{
					SingleFlight: true,
				},
			},
		}).AnyTimes()

		kod.RunTest(t, func(ctx context.Context, up server.HttpUpstream) {
			router := http.NewServeMux()
			up.Register(ctx, router)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/say/bob", nil)
			req.Header.Set("Content-Type", "application/json; charset=utf-8")

			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "application/json; charset=utf-8", rec.Header().Get("Content-Type"))
			assert.Equal(t, "{\"message\":\"Hello bob\"}", rec.Body.String())
		}, kod.WithFakes(kod.Fake[config.Config](mockConfig)))
	})
}
