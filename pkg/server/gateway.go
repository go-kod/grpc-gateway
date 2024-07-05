package server

import (
	"context"
	"net/http"

	"github.com/go-kod/kod"
	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/sysulq/graphql-gateway/pkg/generator"
)

type server struct {
	kod.Implements[ServerComponent]

	config  kod.Ref[ConfigComponent]
	caller  kod.Ref[Caller]
	queryer kod.Ref[Queryer]
}

func (s *server) BuildServer(ctx context.Context) (http.Handler, error) {
	cfg := s.config.Get().Config()

	descs := s.caller.Get().GetDescs()

	gqlDesc, err := generator.NewSchemas(descs, true, true, nil)
	if err != nil {
		return nil, err
	}

	repo := generator.NewRegistry(gqlDesc)

	queryFactory := gateway.QueryerFactory(func(ctx *gateway.PlanningContext, url string) graphql.Queryer {
		q := s.queryer.Get()
		q.SetPM(repo)
		return QueryerLogger{q}
	})
	sources := []*graphql.RemoteSchema{{URL: "url1"}}
	sources[0].Schema = gqlDesc.AsGraphql()[0]

	g, err := gateway.New(sources, gateway.WithQueryerFactory(&queryFactory))
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/query", g.GraphQLHandler)
	if cfg.Playground != nil && *cfg.Playground {
		mux.HandleFunc("/playground", g.PlaygroundHandler)
	}

	var handler http.Handler = mux
	handler = otelhttp.NewMiddleware("graphql-gateway")(handler)

	if cfg.Cors == nil {
		return cors.Default().Handler(handler), nil
	}

	return cors.New(*cfg.Cors).Handler(handler), nil
}
