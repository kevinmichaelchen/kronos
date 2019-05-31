// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Event struct {
	Event                string            `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	UserID               string            `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Properties           map[string]string `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Event) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Event) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Event)(nil), "proto.Event")
	proto.RegisterMapType((map[string]string)(nil), "proto.Event.PropertiesEntry")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x9b, 0x18, 0xb9, 0x58,
	0x5d, 0x41, 0xc2, 0x42, 0x22, 0x5c, 0xac, 0x60, 0x79, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x08, 0x47, 0x48, 0x8c, 0x8b, 0xad, 0xb4, 0x38, 0xb5, 0xc8, 0xd3, 0x45, 0x82, 0x09, 0x2c, 0x0c,
	0xe5, 0x09, 0xd9, 0x70, 0x71, 0x15, 0x14, 0xe5, 0x17, 0xa4, 0x16, 0x95, 0x64, 0xa6, 0x16, 0x4b,
	0x30, 0x2b, 0x30, 0x6b, 0x70, 0x1b, 0xc9, 0x40, 0x8c, 0xd6, 0x03, 0x9b, 0xa7, 0x17, 0x00, 0x97,
	0x76, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x42, 0x52, 0x2f, 0x65, 0xcb, 0xc5, 0x8f, 0x26, 0x2d, 0x24,
	0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0xb5, 0x1c, 0xc4, 0x04, 0x39, 0xa8, 0x2c, 0x31, 0xa7, 0x34,
	0x15, 0x6a, 0x33, 0x84, 0x63, 0xc5, 0x64, 0xc1, 0xa8, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50,
	0x52, 0x69, 0x94, 0xca, 0xc5, 0x03, 0xb6, 0x2c, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48,
	0x93, 0x8b, 0x33, 0x38, 0x35, 0x2f, 0x05, 0xe2, 0x21, 0x1e, 0x64, 0xe7, 0x48, 0xc1, 0x79, 0x20,
	0x8d, 0x4a, 0x0c, 0x42, 0x5a, 0x5c, 0x5c, 0x41, 0xa9, 0x89, 0x10, 0xa5, 0xc5, 0x42, 0x28, 0xb2,
	0xe8, 0x6a, 0x93, 0xd8, 0xc0, 0x5c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x88, 0xab,
	0xdf, 0x41, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	SendEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error)
	ReadEvents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) SendEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.EventService/SendEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) ReadEvents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.EventService/ReadEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	SendEvent(context.Context, *Event) (*Empty, error)
	ReadEvents(context.Context, *Empty) (*Empty, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) SendEvent(ctx context.Context, req *Event) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEvent not implemented")
}
func (*UnimplementedEventServiceServer) ReadEvents(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEvents not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_SendEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SendEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EventService/SendEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SendEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_ReadEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).ReadEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EventService/ReadEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).ReadEvents(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEvent",
			Handler:    _EventService_SendEvent_Handler,
		},
		{
			MethodName: "ReadEvents",
			Handler:    _EventService_ReadEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event.proto",
}
