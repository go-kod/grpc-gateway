// Code generated by kod struct2interface; DO NOT EDIT.

package server

import (
	"context"
	"net/http"

	"github.com/jhump/protoreflect/desc"
	"github.com/nautilus/graphql"
	"github.com/sysulq/graphql-grpc-gateway/pkg/generator"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/protoadapt"
)

// reflection is a component that implements Reflection.
type Reflection interface {
	ListPackages(ctx context.Context, cc grpc.ClientConnInterface) ([]*desc.FileDescriptor, error)
}

// server is a component that implements Gateway.
type Gateway interface {
	BuildServer() (http.Handler, error)
}

// caller is a component that implements Caller.
type Caller interface {
	GetDescs() []*desc.FileDescriptor
	Call(ctx context.Context, rpc *desc.MethodDescriptor, message protoadapt.MessageV1) (protoadapt.MessageV1, error)
}

// queryer is a component that implements Queryer.
type Queryer interface {
	Query(ctx context.Context, input *graphql.QueryInput, result interface{}) error
}

// repository is a component that implements Registry.
type Registry interface {
	SchemaDescriptorList() generator.SchemaDescriptorList
	FindMethodByName(op ast.Operation, name string) *desc.MethodDescriptor
	FindObjectByName(name string) *desc.MessageDescriptor
	FindObjectByFullyQualifiedName(fqn string) (*desc.MessageDescriptor, *ast.Definition)
	FindFieldByName(msg desc.Descriptor, name string) *desc.FieldDescriptor
	FindUnionFieldByMessageFQNAndName(fqn, name string) *desc.FieldDescriptor
	FindGraphqlFieldByProtoField(msg *ast.Definition, name string) *ast.FieldDefinition
}
