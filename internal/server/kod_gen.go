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
	// Caller_Call_FullMethodName is the full name of the method [caller.Call].
	Caller_Call_FullMethodName = "github.com/sysulq/graphql-grpc-gateway/internal/server/Caller.Call"
	// CallerRegistry_FindMethodByName_FullMethodName is the full name of the method [callerRegistry.FindMethodByName].
	CallerRegistry_FindMethodByName_FullMethodName = "github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry.FindMethodByName"
	// CallerRegistry_GetCallerStub_FullMethodName is the full name of the method [callerRegistry.GetCallerStub].
	CallerRegistry_GetCallerStub_FullMethodName = "github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry.GetCallerStub"
	// CallerRegistry_GraphQLSchema_FullMethodName is the full name of the method [callerRegistry.GraphQLSchema].
	CallerRegistry_GraphQLSchema_FullMethodName = "github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry.GraphQLSchema"
	// CallerRegistry_Marshal_FullMethodName is the full name of the method [callerRegistry.Marshal].
	CallerRegistry_Marshal_FullMethodName = "github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry.Marshal"
	// CallerRegistry_Unmarshal_FullMethodName is the full name of the method [callerRegistry.Unmarshal].
	CallerRegistry_Unmarshal_FullMethodName = "github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry.Unmarshal"
	// Reflection_ListPackages_FullMethodName is the full name of the method [reflection.ListPackages].
	Reflection_ListPackages_FullMethodName = "github.com/sysulq/graphql-grpc-gateway/internal/server/Reflection.ListPackages"
	// Gateway_BuildServer_FullMethodName is the full name of the method [server.BuildServer].
	Gateway_BuildServer_FullMethodName = "github.com/sysulq/graphql-grpc-gateway/internal/server/Gateway.BuildServer"
	// Queryer_Query_FullMethodName is the full name of the method [queryer.Query].
	Queryer_Query_FullMethodName = "github.com/sysulq/graphql-grpc-gateway/internal/server/Queryer.Query"
)

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/sysulq/graphql-grpc-gateway/internal/server/Caller",
		Interface: reflect.TypeOf((*Caller)(nil)).Elem(),
		Impl:      reflect.TypeOf(caller{}),
		Refs: `⟦6b1d0901:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/Caller→github.com/sysulq/graphql-grpc-gateway/internal/config/Config⟧,
⟦09e993b0:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/Caller→github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return caller_local_stub{
				impl:        info.Impl.(Caller),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry",
		Interface: reflect.TypeOf((*CallerRegistry)(nil)).Elem(),
		Impl:      reflect.TypeOf(callerRegistry{}),
		Refs: `⟦86a35e79:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry→github.com/sysulq/graphql-grpc-gateway/internal/config/Config⟧,
⟦a92f40b6:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry→github.com/sysulq/graphql-grpc-gateway/internal/server/Reflection⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return callerRegistry_local_stub{
				impl:        info.Impl.(CallerRegistry),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/sysulq/graphql-grpc-gateway/internal/server/Reflection",
		Interface: reflect.TypeOf((*Reflection)(nil)).Elem(),
		Impl:      reflect.TypeOf(reflection{}),
		Refs:      ``,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return reflection_local_stub{
				impl:        info.Impl.(Reflection),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/sysulq/graphql-grpc-gateway/internal/server/Gateway",
		Interface: reflect.TypeOf((*Gateway)(nil)).Elem(),
		Impl:      reflect.TypeOf(server{}),
		Refs: `⟦88a4dee9:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/Gateway→github.com/sysulq/graphql-grpc-gateway/internal/config/Config⟧,
⟦f59f8a3c:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/Gateway→github.com/sysulq/graphql-grpc-gateway/internal/server/Caller⟧,
⟦b39287d6:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/Gateway→github.com/sysulq/graphql-grpc-gateway/internal/server/Queryer⟧,
⟦2bafdbff:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/Gateway→github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return gateway_local_stub{
				impl:        info.Impl.(Gateway),
				interceptor: info.Interceptor,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/sysulq/graphql-grpc-gateway/internal/server/Queryer",
		Interface: reflect.TypeOf((*Queryer)(nil)).Elem(),
		Impl:      reflect.TypeOf(queryer{}),
		Refs: `⟦20a41d92:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/Queryer→github.com/sysulq/graphql-grpc-gateway/internal/config/Config⟧,
⟦8d32f9dd:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/Queryer→github.com/sysulq/graphql-grpc-gateway/internal/server/Caller⟧,
⟦bcaec4e2:KoDeDgE:github.com/sysulq/graphql-grpc-gateway/internal/server/Queryer→github.com/sysulq/graphql-grpc-gateway/internal/server/CallerRegistry⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			return queryer_local_stub{
				impl:        info.Impl.(Queryer),
				interceptor: info.Interceptor,
			}
		},
	})
}

// kod.InstanceOf checks.
var _ kod.InstanceOf[Caller] = (*caller)(nil)
var _ kod.InstanceOf[CallerRegistry] = (*callerRegistry)(nil)
var _ kod.InstanceOf[Reflection] = (*reflection)(nil)
var _ kod.InstanceOf[Gateway] = (*server)(nil)
var _ kod.InstanceOf[Queryer] = (*queryer)(nil)

// Local stub implementations.
// caller_local_stub is a local stub implementation of [Caller].
type caller_local_stub struct {
	impl        Caller
	interceptor interceptor.Interceptor
}

// Check that [caller_local_stub] implements the [Caller] interface.
var _ Caller = (*caller_local_stub)(nil)

// Call wraps the method [caller.Call].
func (s caller_local_stub) Call(ctx context.Context, a1 protoreflect.MethodDescriptor, a2 protoreflect.ProtoMessage) (r0 protoreflect.ProtoMessage, err error) {

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
		FullMethod: Caller_Call_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{a1, a2}, []any{r0}, call)
	return
}

// callerRegistry_local_stub is a local stub implementation of [CallerRegistry].
type callerRegistry_local_stub struct {
	impl        CallerRegistry
	interceptor interceptor.Interceptor
}

// Check that [callerRegistry_local_stub] implements the [CallerRegistry] interface.
var _ CallerRegistry = (*callerRegistry_local_stub)(nil)

// FindMethodByName wraps the method [callerRegistry.FindMethodByName].
func (s callerRegistry_local_stub) FindMethodByName(a0 ast.Operation, a1 string) (r0 protoreflect.MethodDescriptor) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0 = s.impl.FindMethodByName(a0, a1)
	return
}

// GetCallerStub wraps the method [callerRegistry.GetCallerStub].
func (s callerRegistry_local_stub) GetCallerStub(a0 string) (r0 *grpcdynamic.Stub) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0 = s.impl.GetCallerStub(a0)
	return
}

// GraphQLSchema wraps the method [callerRegistry.GraphQLSchema].
func (s callerRegistry_local_stub) GraphQLSchema() (r0 *ast.Schema) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0 = s.impl.GraphQLSchema()
	return
}

// Marshal wraps the method [callerRegistry.Marshal].
func (s callerRegistry_local_stub) Marshal(a0 protoreflect.ProtoMessage, a1 *ast.Field) (r0 interface{}, err error) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0, err = s.impl.Marshal(a0, a1)
	return
}

// Unmarshal wraps the method [callerRegistry.Unmarshal].
func (s callerRegistry_local_stub) Unmarshal(a0 protoreflect.MessageDescriptor, a1 *ast.Field, a2 map[string]interface{}) (r0 protoreflect.ProtoMessage, err error) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0, err = s.impl.Unmarshal(a0, a1, a2)
	return
}

// reflection_local_stub is a local stub implementation of [Reflection].
type reflection_local_stub struct {
	impl        Reflection
	interceptor interceptor.Interceptor
}

// Check that [reflection_local_stub] implements the [Reflection] interface.
var _ Reflection = (*reflection_local_stub)(nil)

// ListPackages wraps the method [reflection.ListPackages].
func (s reflection_local_stub) ListPackages(ctx context.Context, a1 grpc.ClientConnInterface) (r0 []protoreflect.FileDescriptor, err error) {

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
		FullMethod: Reflection_ListPackages_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{r0}, call)
	return
}

// gateway_local_stub is a local stub implementation of [Gateway].
type gateway_local_stub struct {
	impl        Gateway
	interceptor interceptor.Interceptor
}

// Check that [gateway_local_stub] implements the [Gateway] interface.
var _ Gateway = (*gateway_local_stub)(nil)

// BuildServer wraps the method [server.BuildServer].
func (s gateway_local_stub) BuildServer() (r0 http.Handler, err error) {
	// Because the first argument is not context.Context, so interceptors are not supported.
	r0, err = s.impl.BuildServer()
	return
}

// queryer_local_stub is a local stub implementation of [Queryer].
type queryer_local_stub struct {
	impl        Queryer
	interceptor interceptor.Interceptor
}

// Check that [queryer_local_stub] implements the [Queryer] interface.
var _ Queryer = (*queryer_local_stub)(nil)

// Query wraps the method [queryer.Query].
func (s queryer_local_stub) Query(ctx context.Context, a1 *graphql.QueryInput, a2 interface{}) (err error) {

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
		FullMethod: Queryer_Query_FullMethodName,
	}

	err = s.interceptor(ctx, info, []any{a1, a2}, []any{}, call)
	return
}
