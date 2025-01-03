package protojson

import (
	"io"

	// nolint
	"github.com/golang/protobuf/jsonpb"
	// nolint
	"github.com/golang/protobuf/proto"
	// nolint
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type EventHandler struct {
	Status    *status.Status
	Message   proto.Message
	Marshaler jsonpb.Marshaler
}

func NewEventHandler(writer io.Writer, resolver jsonpb.AnyResolver) *EventHandler {
	return &EventHandler{
		Marshaler: jsonpb.Marshaler{
			OrigName:     false,
			EmitDefaults: true,
			EnumsAsInts:  true,
			Indent:       "",
			AnyResolver:  resolver,
		},
	}
}

func (h *EventHandler) OnReceiveResponse(message proto.Message) {
	h.Message = message
}

func (h *EventHandler) OnReceiveTrailers(status *status.Status, _ metadata.MD) {
	h.Status = status
}

func (h *EventHandler) OnResolveMethod(_ *desc.MethodDescriptor) {
}

func (h *EventHandler) OnSendHeaders(_ metadata.MD) {
}

func (h *EventHandler) OnReceiveHeaders(_ metadata.MD) {
}
