package server

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/go-kod/grpc-gateway/internal/config"
	"github.com/go-kod/kod"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type server struct {
	kod.Implements[Gateway]

	config       kod.Ref[config.Config]
	_            kod.Ref[GraphqlCaller]
	queryer      kod.Ref[GraphqlQueryer]
	registry     kod.Ref[GraphqlCallerRegistry]
	httpUpstream kod.Ref[HttpUpstream]
}

func (ins *server) Init(ctx context.Context) error {
	cfg := ins.config.Get().Config()
	if cfg.Engine.Pyroscope.Enable {
		err := ins.config.Get().Config().Engine.Pyroscope.Init(ctx)
		if err != nil {
			return err
		}
	}

	err := cfg.Engine.Otel.Init(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (ins *server) Shutdown(ctx context.Context) error {
	return nil
}

func (s *server) BuildServer() (http.Handler, error) {
	queryFactory := gateway.QueryerFactory(func(ctx *gateway.PlanningContext, url string) graphql.Queryer {
		return s.queryer.Get()
	})

	sources := []*graphql.RemoteSchema{{URL: "url1"}}
	sources[0].Schema = s.registry.Get().GraphQLSchema()

	// formatter.NewFormatter(os.Stdout).FormatSchema(sources[0].Schema)

	opts := []gateway.Option{
		gateway.WithLogger(&noopLogger{}),
		gateway.WithQueryerFactory(&queryFactory),
	}
	if s.config.Get().Config().Server.GraphQL.QueryCache {
		opts = append(opts, gateway.WithQueryPlanCache(NewQueryPlanCacher()))
	}

	g, err := gateway.New(sources, opts...)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	cfg := s.config.Get().Config()

	if !cfg.Server.GraphQL.Disable {
		mux.HandleFunc("/query", g.GraphQLHandler)
		if cfg.Server.GraphQL.Playground {
			mux.HandleFunc("/playground", g.PlaygroundHandler)
		}
	}

	var handler http.Handler = addHeader(mux)
	handler = otelhttp.NewMiddleware("graphql-gateway")(handler)

	if cfg.Server.GraphQL.Jwt.Enable {
		handler = s.jwtAuthHandler(handler)
	}

	return handler, nil
}

func (s *server) BuildHTTPServer() (http.Handler, error) {
	mux := http.NewServeMux()
	cfg := s.config.Get().Config()

	if !cfg.Server.HTTP.Disable {
		s.httpUpstream.Get().Register(context.Background(), mux)
	}

	var handler http.Handler = addHeader(mux)
	handler = otelhttp.NewMiddleware("graphql-gateway")(handler)

	if cfg.Server.HTTP.Jwt.Enable {
		handler = s.jwtAuthHandler(handler)
	}

	return handler, nil
}

type noopLogger struct {
	gateway.Logger
}

func (noopLogger) Debug(args ...interface{})                               {}
func (noopLogger) Info(args ...interface{})                                {}
func (noopLogger) Warn(args ...interface{})                                {}
func (l noopLogger) WithFields(fields gateway.LoggerFields) gateway.Logger { return l }
func (noopLogger) QueryPlanStep(step *gateway.QueryPlanStep)               {}

type queryPlanCacher struct {
	cache *expirable.LRU[string, gateway.QueryPlanList]
}

func NewQueryPlanCacher() *queryPlanCacher {
	cache := expirable.NewLRU[string, gateway.QueryPlanList](1024, nil, time.Hour)
	return &queryPlanCacher{cache: cache}
}

func (c *queryPlanCacher) Retrieve(ctx *gateway.PlanningContext, hash *string, planner gateway.QueryPlanner) (gateway.QueryPlanList, error) {
	// if there is no hash
	if *hash == "" {
		hashString := sha256.Sum256([]byte(ctx.Query))
		// generate a hash that will identify the query for later use
		*hash = hex.EncodeToString(hashString[:])
	}

	if plan, ok := c.cache.Get(*hash); ok {
		return plan, nil
	}

	// compute the plan
	plan, err := planner.Plan(ctx)
	if err != nil {
		return nil, err
	}

	c.cache.Add(*hash, plan)

	return plan, nil
}
