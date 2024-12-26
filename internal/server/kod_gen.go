// Code generated by "kod generate". DO NOT EDIT.
//go:build !ignoreKodGen

package server

import (
	"context"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod/interceptor"
	"github.com/jhump/protoreflect/v2/grpcdynamic"
	"github.com/nautilus/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"net/http"
	"reflect"
)

// Full method names for components.
const (
	// GraphqlCaller_Call_FullMethodName is the full name of the method [graphqlCaller.Call].
	GraphqlCaller_Call_FullMethodName = "github.com/go-kod/grpc-gateway/internal/server/GraphqlCaller.Call"
	// GraphqlReflection_ListPackages_FullMethodName is the full name of the method [graphqlReflection.ListPackages].
	GraphqlReflection_ListPackages_FullMethodName = "github.com/go-kod/grpc-gateway/internal/server/GraphqlReflection.ListPackages"
	// GraphqlQueryer_Query_FullMethodName is the full name of the method [graphqlQueryer.Query].
	GraphqlQueryer_Query_FullMethodName = "github.com/go-kod/grpc-gateway/internal/server/GraphqlQueryer.Query"
	// HttpUpstreamInvoker_Invoke_FullMethodName is the full name of the method [httpUpstreamInvoker.Invoke].
	HttpUpstreamInvoker_Invoke_FullMethodName = "github.com/go-kod/grpc-gateway/internal/server/HttpUpstreamInvoker.Invoke"
	// HttpUpstream_Register_FullMethodName is the full name of the method [httpUpstream.Register].
	HttpUpstream_Register_FullMethodName = "github.com/go-kod/grpc-gateway/internal/server/HttpUpstream.Register"
)

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/grpc-gateway/internal/server/Gateway",
		Interface: reflect.TypeOf((*Gateway)(nil)).Elem(),
		Impl:      reflect.TypeOf(server{}),
		Refs: `⟦08088651:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/Gateway→github.com/go-kod/grpc-gateway/internal/config/Config⟧,
⟦1be0b044:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/Gateway→github.com/go-kod/grpc-gateway/internal/server/GraphqlCaller⟧,
⟦c3524ed4:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/Gateway→github.com/go-kod/grpc-gateway/internal/server/GraphqlQueryer⟧,
⟦d0ac0283:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/Gateway→github.com/go-kod/grpc-gateway/internal/server/GraphqlCallerRegistry⟧,
⟦b40d358d:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/Gateway→github.com/go-kod/grpc-gateway/internal/server/HttpUpstream⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return gateway_local_stub{
				impl:        info.Impl.(Gateway),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/grpc-gateway/internal/server/GraphqlCaller",
		Interface: reflect.TypeOf((*GraphqlCaller)(nil)).Elem(),
		Impl:      reflect.TypeOf(graphqlCaller{}),
		Refs: `⟦4908c6e6:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/GraphqlCaller→github.com/go-kod/grpc-gateway/internal/config/Config⟧,
⟦18d2bb35:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/GraphqlCaller→github.com/go-kod/grpc-gateway/internal/server/GraphqlCallerRegistry⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return graphqlCaller_local_stub{
				impl:        info.Impl.(GraphqlCaller),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/grpc-gateway/internal/server/GraphqlCallerRegistry",
		Interface: reflect.TypeOf((*GraphqlCallerRegistry)(nil)).Elem(),
		Impl:      reflect.TypeOf(graphqlCallerRegistry{}),
		Refs: `⟦da87aa60:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/GraphqlCallerRegistry→github.com/go-kod/grpc-gateway/internal/config/Config⟧,
⟦514cb4fe:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/GraphqlCallerRegistry→github.com/go-kod/grpc-gateway/internal/server/GraphqlReflection⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return graphqlCallerRegistry_local_stub{
				impl:        info.Impl.(GraphqlCallerRegistry),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/grpc-gateway/internal/server/GraphqlReflection",
		Interface: reflect.TypeOf((*GraphqlReflection)(nil)).Elem(),
		Impl:      reflect.TypeOf(graphqlReflection{}),
		Refs:      ``,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return graphqlReflection_local_stub{
				impl:        info.Impl.(GraphqlReflection),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/grpc-gateway/internal/server/GraphqlQueryer",
		Interface: reflect.TypeOf((*GraphqlQueryer)(nil)).Elem(),
		Impl:      reflect.TypeOf(graphqlQueryer{}),
		Refs: `⟦6b17ad14:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/GraphqlQueryer→github.com/go-kod/grpc-gateway/internal/config/Config⟧,
⟦0bd5a48c:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/GraphqlQueryer→github.com/go-kod/grpc-gateway/internal/server/GraphqlCaller⟧,
⟦e28805f0:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/GraphqlQueryer→github.com/go-kod/grpc-gateway/internal/server/GraphqlCallerRegistry⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return graphqlQueryer_local_stub{
				impl:        info.Impl.(GraphqlQueryer),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/grpc-gateway/internal/server/HttpUpstreamInvoker",
		Interface: reflect.TypeOf((*HttpUpstreamInvoker)(nil)).Elem(),
		Impl:      reflect.TypeOf(httpUpstreamInvoker{}),
		Refs:      `⟦064c132d:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/HttpUpstreamInvoker→github.com/go-kod/grpc-gateway/internal/config/Config⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return httpUpstreamInvoker_local_stub{
				impl:        info.Impl.(HttpUpstreamInvoker),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/grpc-gateway/internal/server/HttpUpstream",
		Interface: reflect.TypeOf((*HttpUpstream)(nil)).Elem(),
		Impl:      reflect.TypeOf(httpUpstream{}),
		Refs: `⟦8e0a637a:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/HttpUpstream→github.com/go-kod/grpc-gateway/internal/server/HttpUpstreamInvoker⟧,
⟦d1d5181e:KoDeDgE:github.com/go-kod/grpc-gateway/internal/server/HttpUpstream→github.com/go-kod/grpc-gateway/internal/config/Config⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return httpUpstream_local_stub{
				impl:        info.Impl.(HttpUpstream),
				interceptor: info.Interceptor,
			}
		},
	})
}

// CodeGen version check.
var _ kod.CodeGenLatestVersion = kod.CodeGenVersion[[0][1]struct{}](`
ERROR: You generated this file with 'kod generate' (codegen
version v0.1.0). The generated code is incompatible with the version of the
github.com/go-kod/kod module that you're using. The kod module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/go-kod/kod

We recommend updating the kod module and the 'kod generate' command by
running the following.

    go get github.com/go-kod/kod@latest
    go install github.com/go-kod/kod/cmd/kod@latest

Then, re-run 'kod generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/go-kod/kod/issues.
`)

// kod.InstanceOf checks.
var _ kod.InstanceOf[Gateway] = (*server)(nil)
var _ kod.InstanceOf[GraphqlCaller] = (*graphqlCaller)(nil)
var _ kod.InstanceOf[GraphqlCallerRegistry] = (*graphqlCallerRegistry)(nil)
var _ kod.InstanceOf[GraphqlReflection] = (*graphqlReflection)(nil)
var _ kod.InstanceOf[GraphqlQueryer] = (*graphqlQueryer)(nil)
var _ kod.InstanceOf[HttpUpstreamInvoker] = (*httpUpstreamInvoker)(nil)
var _ kod.InstanceOf[HttpUpstream] = (*httpUpstream)(nil)

// Local stub implementations.
// gateway_local_stub is a local stub implementation of [Gateway].
type gateway_local_stub struct {
	impl        Gateway
	interceptor interceptor.Interceptor
}

// Check that [gateway_local_stub] implements the [Gateway] interface.
var _ Gateway = (*gateway_local_stub)(nil)

// BuildHTTPServer wraps the method [server.BuildHTTPServer].
func (s gateway_local_stub) BuildHTTPServer() (r0 http.Handler, err error) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0, err = s.impl.BuildHTTPServer()
	return
}

// BuildServer wraps the method [server.BuildServer].
func (s gateway_local_stub) BuildServer() (r0 http.Handler, err error) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0, err = s.impl.BuildServer()
	return
}

// graphqlCaller_local_stub is a local stub implementation of [GraphqlCaller].
type graphqlCaller_local_stub struct {
	impl        GraphqlCaller
	interceptor interceptor.Interceptor
}

// Check that [graphqlCaller_local_stub] implements the [GraphqlCaller] interface.
var _ GraphqlCaller = (*graphqlCaller_local_stub)(nil)

// Call wraps the method [graphqlCaller.Call].
func (s graphqlCaller_local_stub) Call(ctx context.Context, a1 protoreflect.MethodDescriptor, a2 protoreflect.ProtoMessage) (r0 protoreflect.ProtoMessage, err error) {

	if s.interceptor == nil {
		r0, err = s.impl.Call(ctx, a1, a2)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		r0, err = s.impl.Call(ctx, a1, a2)
		res[0] = r0
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		FullMethod: GraphqlCaller_Call_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{a1, a2}, []any{r0}, call)
	return
}

// graphqlCallerRegistry_local_stub is a local stub implementation of [GraphqlCallerRegistry].
type graphqlCallerRegistry_local_stub struct {
	impl        GraphqlCallerRegistry
	interceptor interceptor.Interceptor
}

// Check that [graphqlCallerRegistry_local_stub] implements the [GraphqlCallerRegistry] interface.
var _ GraphqlCallerRegistry = (*graphqlCallerRegistry_local_stub)(nil)

// FindMethodByName wraps the method [graphqlCallerRegistry.FindMethodByName].
func (s graphqlCallerRegistry_local_stub) FindMethodByName(a0 ast.Operation, a1 string) (r0 protoreflect.MethodDescriptor) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0 = s.impl.FindMethodByName(a0, a1)
	return
}

// GetCallerStub wraps the method [graphqlCallerRegistry.GetCallerStub].
func (s graphqlCallerRegistry_local_stub) GetCallerStub(a0 string) (r0 *grpcdynamic.Stub) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0 = s.impl.GetCallerStub(a0)
	return
}

// GraphQLSchema wraps the method [graphqlCallerRegistry.GraphQLSchema].
func (s graphqlCallerRegistry_local_stub) GraphQLSchema() (r0 *ast.Schema) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0 = s.impl.GraphQLSchema()
	return
}

// Marshal wraps the method [graphqlCallerRegistry.Marshal].
func (s graphqlCallerRegistry_local_stub) Marshal(a0 protoreflect.ProtoMessage, a1 *ast.Field) (r0 interface{}, err error) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0, err = s.impl.Marshal(a0, a1)
	return
}

// Unmarshal wraps the method [graphqlCallerRegistry.Unmarshal].
func (s graphqlCallerRegistry_local_stub) Unmarshal(a0 protoreflect.MessageDescriptor, a1 *ast.Field, a2 map[string]interface{}) (r0 protoreflect.ProtoMessage, err error) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0, err = s.impl.Unmarshal(a0, a1, a2)
	return
}

// graphqlReflection_local_stub is a local stub implementation of [GraphqlReflection].
type graphqlReflection_local_stub struct {
	impl        GraphqlReflection
	interceptor interceptor.Interceptor
}

// Check that [graphqlReflection_local_stub] implements the [GraphqlReflection] interface.
var _ GraphqlReflection = (*graphqlReflection_local_stub)(nil)

// ListPackages wraps the method [graphqlReflection.ListPackages].
func (s graphqlReflection_local_stub) ListPackages(ctx context.Context, a1 grpc.ClientConnInterface) (r0 []protoreflect.FileDescriptor, err error) {

	if s.interceptor == nil {
		r0, err = s.impl.ListPackages(ctx, a1)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		r0, err = s.impl.ListPackages(ctx, a1)
		res[0] = r0
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		FullMethod: GraphqlReflection_ListPackages_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{r0}, call)
	return
}

// graphqlQueryer_local_stub is a local stub implementation of [GraphqlQueryer].
type graphqlQueryer_local_stub struct {
	impl        GraphqlQueryer
	interceptor interceptor.Interceptor
}

// Check that [graphqlQueryer_local_stub] implements the [GraphqlQueryer] interface.
var _ GraphqlQueryer = (*graphqlQueryer_local_stub)(nil)

// Query wraps the method [graphqlQueryer.Query].
func (s graphqlQueryer_local_stub) Query(ctx context.Context, a1 *graphql.QueryInput, a2 interface{}) (err error) {

	if s.interceptor == nil {
		err = s.impl.Query(ctx, a1, a2)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		err = s.impl.Query(ctx, a1, a2)
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		FullMethod: GraphqlQueryer_Query_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{a1, a2}, []any{}, call)
	return
}

// httpUpstreamInvoker_local_stub is a local stub implementation of [HttpUpstreamInvoker].
type httpUpstreamInvoker_local_stub struct {
	impl        HttpUpstreamInvoker
	interceptor interceptor.Interceptor
}

// Check that [httpUpstreamInvoker_local_stub] implements the [HttpUpstreamInvoker] interface.
var _ HttpUpstreamInvoker = (*httpUpstreamInvoker_local_stub)(nil)

// Invoke wraps the method [httpUpstreamInvoker.Invoke].
func (s httpUpstreamInvoker_local_stub) Invoke(ctx context.Context, a1 http.ResponseWriter, a2 *http.Request, a3 upstreamInfo, a4 string, a5 []string) {

	if s.interceptor == nil {
		s.impl.Invoke(ctx, a1, a2, a3, a4, a5)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		s.impl.Invoke(ctx, a1, a2, a3, a4, a5)
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		FullMethod: HttpUpstreamInvoker_Invoke_FullMethodName,
	}

	_ = s.interceptor(ctx, info, []any{a1, a2, a3, a4, a5}, []any{}, call)
}

// httpUpstream_local_stub is a local stub implementation of [HttpUpstream].
type httpUpstream_local_stub struct {
	impl        HttpUpstream
	interceptor interceptor.Interceptor
}

// Check that [httpUpstream_local_stub] implements the [HttpUpstream] interface.
var _ HttpUpstream = (*httpUpstream_local_stub)(nil)

// Register wraps the method [httpUpstream.Register].
func (s httpUpstream_local_stub) Register(ctx context.Context, a1 *http.ServeMux) {

	if s.interceptor == nil {
		s.impl.Register(ctx, a1)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		s.impl.Register(ctx, a1)
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		FullMethod: HttpUpstream_Register_FullMethodName,
	}

	_ = s.interceptor(ctx, info, []any{a1}, []any{}, call)
}
