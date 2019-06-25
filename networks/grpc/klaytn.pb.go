// Code generated by protoc-gen-go. DO NOT EDIT.
// source: klaytn.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d8429895d2d55b, []int{0}
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

type RPCRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Params               []byte   `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCRequest) Reset()         { *m = RPCRequest{} }
func (m *RPCRequest) String() string { return proto.CompactTextString(m) }
func (*RPCRequest) ProtoMessage()    {}
func (*RPCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d8429895d2d55b, []int{1}
}

func (m *RPCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCRequest.Unmarshal(m, b)
}
func (m *RPCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCRequest.Marshal(b, m, deterministic)
}
func (m *RPCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCRequest.Merge(m, src)
}
func (m *RPCRequest) XXX_Size() int {
	return xxx_messageInfo_RPCRequest.Size(m)
}
func (m *RPCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPCRequest proto.InternalMessageInfo

func (m *RPCRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *RPCRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *RPCRequest) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

type RPCResponse struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCResponse) Reset()         { *m = RPCResponse{} }
func (m *RPCResponse) String() string { return proto.CompactTextString(m) }
func (*RPCResponse) ProtoMessage()    {}
func (*RPCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d8429895d2d55b, []int{2}
}

func (m *RPCResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCResponse.Unmarshal(m, b)
}
func (m *RPCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCResponse.Marshal(b, m, deterministic)
}
func (m *RPCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCResponse.Merge(m, src)
}
func (m *RPCResponse) XXX_Size() int {
	return xxx_messageInfo_RPCResponse.Size(m)
}
func (m *RPCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RPCResponse proto.InternalMessageInfo

func (m *RPCResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "grpc.Empty")
	proto.RegisterType((*RPCRequest)(nil), "grpc.RPCRequest")
	proto.RegisterType((*RPCResponse)(nil), "grpc.RPCResponse")
}

func init() { proto.RegisterFile("klaytn.proto", fileDescriptor_c6d8429895d2d55b) }

var fileDescriptor_c6d8429895d2d55b = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xfd, 0x66, 0xed, 0xd8, 0x67, 0x61, 0x9a, 0x83, 0x14, 0x4f, 0xa3, 0x17, 0x0b, 0x62,
	0x19, 0xce, 0x5f, 0xd0, 0xe1, 0x49, 0x90, 0x12, 0xc1, 0x7b, 0xda, 0x7e, 0x68, 0xb1, 0x59, 0x62,
	0x92, 0x09, 0xfd, 0x3b, 0x1e, 0xfc, 0x9d, 0x92, 0xd8, 0x51, 0x8f, 0x7a, 0x7c, 0xdf, 0x24, 0x4f,
	0x1e, 0x5e, 0x4c, 0xde, 0x7a, 0x31, 0xb8, 0x5d, 0xa1, 0x8d, 0x72, 0x8a, 0x45, 0x2f, 0x46, 0x37,
	0xd9, 0x1c, 0x4f, 0xee, 0xa5, 0x76, 0x43, 0xf6, 0x8c, 0xc8, 0xab, 0x2d, 0xa7, 0xf7, 0x3d, 0x59,
	0xc7, 0x52, 0x9c, 0x5b, 0x32, 0x1f, 0x5d, 0x43, 0x29, 0xac, 0x20, 0x5f, 0xf0, 0x43, 0x64, 0x17,
	0x18, 0x4b, 0x72, 0xaf, 0xaa, 0x4d, 0x67, 0xe1, 0x60, 0x4c, 0xbe, 0xd7, 0xc2, 0x08, 0x69, 0xd3,
	0xe3, 0x15, 0xe4, 0x09, 0x1f, 0x53, 0x76, 0x85, 0xa7, 0x81, 0x6b, 0xb5, 0xda, 0x59, 0xf2, 0x60,
	0x2d, 0x86, 0x5e, 0x89, 0x36, 0x80, 0x13, 0x7e, 0x88, 0xb7, 0x5f, 0x80, 0xf8, 0x10, 0x04, 0x1f,
	0x55, 0x4b, 0xec, 0x06, 0xa3, 0xad, 0xe8, 0x7b, 0x76, 0x56, 0x78, 0xcf, 0x62, 0x72, 0xbb, 0x3c,
	0xff, 0xd5, 0xfc, 0x50, 0xb3, 0x23, 0x76, 0x87, 0x8b, 0xa7, 0x7d, 0x6d, 0x1b, 0xd3, 0xd5, 0xf4,
	0xc7, 0x37, 0x6b, 0x60, 0x1b, 0x8c, 0xcb, 0xee, 0x1f, 0xdf, 0xe4, 0xb0, 0x86, 0xf2, 0x1a, 0x97,
	0x8d, 0x92, 0xc5, 0x38, 0xa6, 0xbf, 0x54, 0x2e, 0x27, 0xf1, 0xca, 0x8f, 0x5b, 0xc1, 0xe7, 0x2c,
	0xf2, 0x5d, 0x1d, 0x87, 0xb1, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75, 0x07, 0x2f, 0xb1,
	0x7c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KlaytnNodeClient is the client API for KlaytnNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KlaytnNodeClient interface {
	Call(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*RPCResponse, error)
	Subscribe(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (KlaytnNode_SubscribeClient, error)
	BiCall(ctx context.Context, opts ...grpc.CallOption) (KlaytnNode_BiCallClient, error)
}

type klaytnNodeClient struct {
	cc *grpc.ClientConn
}

func NewKlaytnNodeClient(cc *grpc.ClientConn) KlaytnNodeClient {
	return &klaytnNodeClient{cc}
}

func (c *klaytnNodeClient) Call(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*RPCResponse, error) {
	out := new(RPCResponse)
	err := c.cc.Invoke(ctx, "/grpc.KlaytnNode/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klaytnNodeClient) Subscribe(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (KlaytnNode_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KlaytnNode_serviceDesc.Streams[0], "/grpc.KlaytnNode/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &klaytnNodeSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KlaytnNode_SubscribeClient interface {
	Recv() (*RPCResponse, error)
	grpc.ClientStream
}

type klaytnNodeSubscribeClient struct {
	grpc.ClientStream
}

func (x *klaytnNodeSubscribeClient) Recv() (*RPCResponse, error) {
	m := new(RPCResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *klaytnNodeClient) BiCall(ctx context.Context, opts ...grpc.CallOption) (KlaytnNode_BiCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KlaytnNode_serviceDesc.Streams[1], "/grpc.KlaytnNode/BiCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &klaytnNodeBiCallClient{stream}
	return x, nil
}

type KlaytnNode_BiCallClient interface {
	Send(*RPCRequest) error
	Recv() (*RPCResponse, error)
	grpc.ClientStream
}

type klaytnNodeBiCallClient struct {
	grpc.ClientStream
}

func (x *klaytnNodeBiCallClient) Send(m *RPCRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *klaytnNodeBiCallClient) Recv() (*RPCResponse, error) {
	m := new(RPCResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KlaytnNodeServer is the server API for KlaytnNode service.
type KlaytnNodeServer interface {
	Call(context.Context, *RPCRequest) (*RPCResponse, error)
	Subscribe(*RPCRequest, KlaytnNode_SubscribeServer) error
	BiCall(KlaytnNode_BiCallServer) error
}

func RegisterKlaytnNodeServer(s *grpc.Server, srv KlaytnNodeServer) {
	s.RegisterService(&_KlaytnNode_serviceDesc, srv)
}

func _KlaytnNode_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlaytnNodeServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KlaytnNode/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlaytnNodeServer).Call(ctx, req.(*RPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlaytnNode_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RPCRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KlaytnNodeServer).Subscribe(m, &klaytnNodeSubscribeServer{stream})
}

type KlaytnNode_SubscribeServer interface {
	Send(*RPCResponse) error
	grpc.ServerStream
}

type klaytnNodeSubscribeServer struct {
	grpc.ServerStream
}

func (x *klaytnNodeSubscribeServer) Send(m *RPCResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KlaytnNode_BiCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KlaytnNodeServer).BiCall(&klaytnNodeBiCallServer{stream})
}

type KlaytnNode_BiCallServer interface {
	Send(*RPCResponse) error
	Recv() (*RPCRequest, error)
	grpc.ServerStream
}

type klaytnNodeBiCallServer struct {
	grpc.ServerStream
}

func (x *klaytnNodeBiCallServer) Send(m *RPCResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *klaytnNodeBiCallServer) Recv() (*RPCRequest, error) {
	m := new(RPCRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _KlaytnNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.KlaytnNode",
	HandlerType: (*KlaytnNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _KlaytnNode_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _KlaytnNode_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BiCall",
			Handler:       _KlaytnNode_BiCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "klaytn.proto",
}
