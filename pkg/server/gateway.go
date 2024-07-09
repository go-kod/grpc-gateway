package server

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/go-kod/kod"
	"github.com/grafana/pyroscope-go"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type server struct {
	kod.Implements[ServerComponent]

	profiler *pyroscope.Profiler

	config   kod.Ref[ConfigComponent]
	queryer  kod.Ref[Queryer]
	registry kod.Ref[Registry]
}

func (ins *server) Init(ctx context.Context) error {
	profiler, err := ins.config.Get().Config().Pyroscope.Build(ctx)
	if err != nil {
		return err
	}

	ins.profiler = profiler

	return nil
}

func (ins *server) Shutdown(ctx context.Context) error {
	if ins.profiler != nil {
		ins.profiler.Stop()
	}

	return nil
}

func (s *server) BuildServer() (http.Handler, error) {
	queryFactory := gateway.QueryerFactory(func(ctx *gateway.PlanningContext, url string) graphql.Queryer {
		return s.queryer.Get()
	})

	sources := []*graphql.RemoteSchema{{URL: "url1"}}
	sources[0].Schema = s.registry.Get().SchemaDescriptorList().AsGraphql()[0]

	g, err := gateway.New(sources,
		gateway.WithLogger(&noopLogger{}),
		gateway.WithQueryPlanCache(NewQueryPlanCacher()),
		gateway.WithQueryerFactory(&queryFactory))
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/query", g.GraphQLHandler)

	cfg := s.config.Get().Config()
	if cfg.Playground {
		mux.HandleFunc("/playground", g.PlaygroundHandler)
	}

	var handler http.Handler = addHeader(mux)
	handler = otelhttp.NewMiddleware("graphql-gateway")(handler)

	return cors.New(cfg.Cors).Handler(handler), nil
}

func addHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		md := httpHeadersToGRPCMetadata(r.Header)
		ctx := metadata.NewOutgoingContext(r.Context(), md)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
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

// httpHeadersToGRPCMetadata converts HTTP headers to gRPC metadata.
func httpHeadersToGRPCMetadata(headers http.Header) metadata.MD {
	grpcMetadata := metadata.MD{}
	for key, values := range headers {
		grpcKey, ok := DefaultHeaderMatcher(key)
		if ok {
			for _, value := range values {
				grpcMetadata.Append(grpcKey, value)
			}
		}
	}
	return grpcMetadata
}

const (
	MetadataHeaderPrefix = "Grpc-Metadata-"
	MetadataPrefix       = "grpcgateway-"
)

func DefaultHeaderMatcher(key string) (string, bool) {
	switch key = textproto.CanonicalMIMEHeaderKey(key); {
	case isPermanentHTTPHeader(key):
		return MetadataPrefix + key, true
	case strings.HasPrefix(key, MetadataHeaderPrefix):
		return key[len(MetadataHeaderPrefix):], true
	}
	return "", false
}

// isPermanentHTTPHeader checks whether hdr belongs to the list of
// permanent request headers maintained by IANA.
// http://www.iana.org/assignments/message-headers/message-headers.xml
func isPermanentHTTPHeader(hdr string) bool {
	switch hdr {
	case
		"Accept",
		"Accept-Charset",
		"Accept-Language",
		"Accept-Ranges",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Cookie",
		"Date",
		"Expect",
		"From",
		"Host",
		"If-Match",
		"If-Modified-Since",
		"If-None-Match",
		"If-Schedule-Tag-Match",
		"If-Unmodified-Since",
		"Max-Forwards",
		"Origin",
		"Pragma",
		"Referer",
		"User-Agent",
		"Via",
		"Warning":
		return true
	}
	return false
}
