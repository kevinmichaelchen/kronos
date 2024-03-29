// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	UserID               string            `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	SessionID            string            `protobuf:"bytes,2,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	TimeMs               int64             `protobuf:"varint,3,opt,name=timeMs,proto3" json:"timeMs,omitempty"`
	Properties           map[string]string `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (m *Event) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Event) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *Event) GetTimeMs() int64 {
	if m != nil {
		return m.TimeMs
	}
	return 0
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

type UserQuery struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserQuery) Reset()         { *m = UserQuery{} }
func (m *UserQuery) String() string { return proto.CompactTextString(m) }
func (*UserQuery) ProtoMessage()    {}
func (*UserQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}

func (m *UserQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserQuery.Unmarshal(m, b)
}
func (m *UserQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserQuery.Marshal(b, m, deterministic)
}
func (m *UserQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserQuery.Merge(m, src)
}
func (m *UserQuery) XXX_Size() int {
	return xxx_messageInfo_UserQuery.Size(m)
}
func (m *UserQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_UserQuery.DiscardUnknown(m)
}

var xxx_messageInfo_UserQuery proto.InternalMessageInfo

func (m *UserQuery) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserQuery) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *UserQuery) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type DurationResponse struct {
	DurationMs           int64    `protobuf:"varint,1,opt,name=durationMs,proto3" json:"durationMs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DurationResponse) Reset()         { *m = DurationResponse{} }
func (m *DurationResponse) String() string { return proto.CompactTextString(m) }
func (*DurationResponse) ProtoMessage()    {}
func (*DurationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{3}
}

func (m *DurationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DurationResponse.Unmarshal(m, b)
}
func (m *DurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DurationResponse.Marshal(b, m, deterministic)
}
func (m *DurationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DurationResponse.Merge(m, src)
}
func (m *DurationResponse) XXX_Size() int {
	return xxx_messageInfo_DurationResponse.Size(m)
}
func (m *DurationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DurationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DurationResponse proto.InternalMessageInfo

func (m *DurationResponse) GetDurationMs() int64 {
	if m != nil {
		return m.DurationMs
	}
	return 0
}

type CountResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResponse) Reset()         { *m = CountResponse{} }
func (m *CountResponse) String() string { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()    {}
func (*CountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{4}
}

func (m *CountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountResponse.Unmarshal(m, b)
}
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountResponse.Marshal(b, m, deterministic)
}
func (m *CountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResponse.Merge(m, src)
}
func (m *CountResponse) XXX_Size() int {
	return xxx_messageInfo_CountResponse.Size(m)
}
func (m *CountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountResponse proto.InternalMessageInfo

func (m *CountResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "proto.Event")
	proto.RegisterMapType((map[string]string)(nil), "proto.Event.PropertiesEntry")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*UserQuery)(nil), "proto.UserQuery")
	proto.RegisterType((*DurationResponse)(nil), "proto.DurationResponse")
	proto.RegisterType((*CountResponse)(nil), "proto.CountResponse")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdb, 0x6a, 0xe2, 0x40,
	0x18, 0x36, 0x66, 0xe3, 0xe2, 0xaf, 0xbb, 0xeb, 0x0e, 0xc1, 0x0d, 0x22, 0x8b, 0x04, 0x16, 0x64,
	0x2f, 0x72, 0x91, 0xbd, 0x59, 0x4a, 0xbd, 0xaa, 0xc1, 0x4a, 0x6b, 0x0f, 0x91, 0x3e, 0x40, 0x34,
	0x7f, 0x4b, 0x68, 0x9d, 0x09, 0x33, 0x13, 0x21, 0xef, 0xd3, 0x27, 0xea, 0x13, 0x95, 0x99, 0x4c,
	0x53, 0xb5, 0x2d, 0xf4, 0xca, 0xf9, 0x0e, 0xff, 0xe9, 0x33, 0xd0, 0xc1, 0x2d, 0x52, 0x19, 0xe4,
	0x9c, 0x49, 0x46, 0x1c, 0xfd, 0xe3, 0x3f, 0x59, 0xe0, 0x44, 0x8a, 0x26, 0x7d, 0x68, 0x15, 0x02,
	0xf9, 0x7c, 0xea, 0x59, 0x23, 0x6b, 0xdc, 0x8e, 0x0d, 0x22, 0x43, 0x68, 0x0b, 0x14, 0x22, 0x63,
	0x74, 0x3e, 0xf5, 0x9a, 0x5a, 0x7a, 0x25, 0x54, 0x95, 0xcc, 0x36, 0xb8, 0x10, 0x9e, 0x3d, 0xb2,
	0xc6, 0x76, 0x6c, 0x10, 0x39, 0x06, 0xc8, 0x39, 0xcb, 0x91, 0xcb, 0x0c, 0x85, 0xf7, 0x65, 0x64,
	0x8f, 0x3b, 0xe1, 0xb0, 0x1a, 0x1d, 0xe8, 0x79, 0xc1, 0x55, 0x2d, 0x47, 0x54, 0xf2, 0x32, 0xde,
	0xf1, 0x0f, 0x26, 0xf0, 0xe3, 0x40, 0x26, 0x3d, 0xb0, 0xef, 0xb1, 0x34, 0xbb, 0xa9, 0x27, 0x71,
	0xc1, 0xd9, 0x26, 0x0f, 0x05, 0x9a, 0xa5, 0x2a, 0x70, 0xd4, 0xfc, 0x6f, 0xf9, 0x5f, 0xc1, 0x89,
	0x36, 0xb9, 0x2c, 0xfd, 0x33, 0x68, 0xdf, 0x08, 0xe4, 0xd7, 0x05, 0xf2, 0xf2, 0xc3, 0x03, 0x5d,
	0x70, 0x84, 0x4c, 0xb8, 0xd4, 0x7d, 0xec, 0xb8, 0x02, 0x6a, 0x1e, 0xd2, 0xd4, 0x5c, 0xa5, 0x9e,
	0x7e, 0x08, 0xbd, 0x69, 0xc1, 0x13, 0x99, 0x31, 0x1a, 0xa3, 0xc8, 0x19, 0x15, 0x48, 0x7e, 0x03,
	0xa4, 0x86, 0x5b, 0x08, 0xdd, 0xd7, 0x8e, 0x77, 0x18, 0xff, 0x0f, 0x7c, 0x3b, 0x61, 0x05, 0x95,
	0x75, 0x81, 0x0b, 0xce, 0x5a, 0x11, 0xc6, 0x5b, 0x81, 0xf0, 0xb1, 0x09, 0x5d, 0x9d, 0xca, 0x12,
	0xf9, 0x36, 0x5b, 0x23, 0x09, 0xe0, 0xfb, 0x12, 0x69, 0x7a, 0xce, 0xee, 0x32, 0x5a, 0xfd, 0x3d,
	0xdd, 0xdd, 0xf0, 0x06, 0x35, 0xd2, 0x67, 0x36, 0x48, 0x08, 0x44, 0xf9, 0x4f, 0x31, 0xe1, 0x72,
	0x85, 0x89, 0xfc, 0x4c, 0xcd, 0x5f, 0x80, 0x18, 0x93, 0x54, 0x8b, 0x82, 0xec, 0xa9, 0x6f, 0xbc,
	0x13, 0xf8, 0x39, 0x43, 0x79, 0x51, 0x6c, 0x56, 0xc8, 0x2f, 0x6f, 0xf5, 0x5a, 0x82, 0xf4, 0x8c,
	0xa9, 0x8e, 0x78, 0xe0, 0x1a, 0x66, 0xef, 0x66, 0xbf, 0x41, 0x22, 0xe8, 0xcf, 0x50, 0x2a, 0xdf,
	0xb2, 0xfa, 0x72, 0x5e, 0x82, 0x7c, 0xa7, 0xc7, 0x2f, 0xc3, 0x1c, 0x66, 0xed, 0x37, 0x56, 0x2d,
	0xad, 0xfc, 0x7b, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x6f, 0x52, 0x49, 0xc9, 0x02, 0x00, 0x00,
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
	SendLoginEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error)
	SendHeartbeatEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error)
	// Read all events
	ReadEvents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetNumberOfLogins(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*CountResponse, error)
	// Find out how long Alice was in VR between X and Y.
	GetUserSessionDuration(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*DurationResponse, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) SendLoginEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.EventService/SendLoginEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) SendHeartbeatEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.EventService/SendHeartbeatEvent", in, out, opts...)
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

func (c *eventServiceClient) GetNumberOfLogins(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/proto.EventService/GetNumberOfLogins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetUserSessionDuration(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*DurationResponse, error) {
	out := new(DurationResponse)
	err := c.cc.Invoke(ctx, "/proto.EventService/GetUserSessionDuration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	SendLoginEvent(context.Context, *Event) (*Empty, error)
	SendHeartbeatEvent(context.Context, *Event) (*Empty, error)
	// Read all events
	ReadEvents(context.Context, *Empty) (*Empty, error)
	GetNumberOfLogins(context.Context, *UserQuery) (*CountResponse, error)
	// Find out how long Alice was in VR between X and Y.
	GetUserSessionDuration(context.Context, *UserQuery) (*DurationResponse, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) SendLoginEvent(ctx context.Context, req *Event) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLoginEvent not implemented")
}
func (*UnimplementedEventServiceServer) SendHeartbeatEvent(ctx context.Context, req *Event) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartbeatEvent not implemented")
}
func (*UnimplementedEventServiceServer) ReadEvents(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEvents not implemented")
}
func (*UnimplementedEventServiceServer) GetNumberOfLogins(ctx context.Context, req *UserQuery) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberOfLogins not implemented")
}
func (*UnimplementedEventServiceServer) GetUserSessionDuration(ctx context.Context, req *UserQuery) (*DurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSessionDuration not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_SendLoginEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SendLoginEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EventService/SendLoginEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SendLoginEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_SendHeartbeatEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SendHeartbeatEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EventService/SendHeartbeatEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SendHeartbeatEvent(ctx, req.(*Event))
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

func _EventService_GetNumberOfLogins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetNumberOfLogins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EventService/GetNumberOfLogins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetNumberOfLogins(ctx, req.(*UserQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetUserSessionDuration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetUserSessionDuration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EventService/GetUserSessionDuration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetUserSessionDuration(ctx, req.(*UserQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLoginEvent",
			Handler:    _EventService_SendLoginEvent_Handler,
		},
		{
			MethodName: "SendHeartbeatEvent",
			Handler:    _EventService_SendHeartbeatEvent_Handler,
		},
		{
			MethodName: "ReadEvents",
			Handler:    _EventService_ReadEvents_Handler,
		},
		{
			MethodName: "GetNumberOfLogins",
			Handler:    _EventService_GetNumberOfLogins_Handler,
		},
		{
			MethodName: "GetUserSessionDuration",
			Handler:    _EventService_GetUserSessionDuration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event.proto",
}
