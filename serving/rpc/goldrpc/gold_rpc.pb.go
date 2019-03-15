// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gold_rpc.proto

package goldrpc

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SyncRequest struct {
	Data                 *SyncData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_860aae571b0964f2, []int{0}
}

func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRequest.Unmarshal(m, b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
}
func (m *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(m, src)
}
func (m *SyncRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRequest.Size(m)
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

func (m *SyncRequest) GetData() *SyncData {
	if m != nil {
		return m.Data
	}
	return nil
}

type SyncResponse struct {
	Data                 *SyncData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SyncResponse) Reset()         { *m = SyncResponse{} }
func (m *SyncResponse) String() string { return proto.CompactTextString(m) }
func (*SyncResponse) ProtoMessage()    {}
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_860aae571b0964f2, []int{1}
}

func (m *SyncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncResponse.Unmarshal(m, b)
}
func (m *SyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncResponse.Marshal(b, m, deterministic)
}
func (m *SyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncResponse.Merge(m, src)
}
func (m *SyncResponse) XXX_Size() int {
	return xxx_messageInfo_SyncResponse.Size(m)
}
func (m *SyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncResponse proto.InternalMessageInfo

func (m *SyncResponse) GetData() *SyncData {
	if m != nil {
		return m.Data
	}
	return nil
}

type SyncData struct {
	// the sender of data (service name).
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// data source endpoint.
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// the timestamp of data generation.
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// real transfer data in json string format.
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncData) Reset()         { *m = SyncData{} }
func (m *SyncData) String() string { return proto.CompactTextString(m) }
func (*SyncData) ProtoMessage()    {}
func (*SyncData) Descriptor() ([]byte, []int) {
	return fileDescriptor_860aae571b0964f2, []int{2}
}

func (m *SyncData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncData.Unmarshal(m, b)
}
func (m *SyncData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncData.Marshal(b, m, deterministic)
}
func (m *SyncData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncData.Merge(m, src)
}
func (m *SyncData) XXX_Size() int {
	return xxx_messageInfo_SyncData.Size(m)
}
func (m *SyncData) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncData.DiscardUnknown(m)
}

var xxx_messageInfo_SyncData proto.InternalMessageInfo

func (m *SyncData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SyncData) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *SyncData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SyncData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*SyncRequest)(nil), "goldrpc.SyncRequest")
	proto.RegisterType((*SyncResponse)(nil), "goldrpc.SyncResponse")
	proto.RegisterType((*SyncData)(nil), "goldrpc.SyncData")
}

func init() { proto.RegisterFile("gold_rpc.proto", fileDescriptor_860aae571b0964f2) }

var fileDescriptor_860aae571b0964f2 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0x8d, 0x5d, 0xfa, 0x31, 0x2d, 0x8a, 0x83, 0xca, 0x52, 0x3c, 0xd4, 0x05, 0xa1, 0xa7,
	0x1c, 0xaa, 0xbd, 0x4b, 0x15, 0xbc, 0x96, 0xf8, 0x00, 0x12, 0x93, 0x50, 0x16, 0xd2, 0x64, 0x4c,
	0xe2, 0xc1, 0xb7, 0x97, 0xa4, 0x51, 0xd1, 0x93, 0xb7, 0xfc, 0x3f, 0x7e, 0x61, 0x66, 0xe0, 0x64,
	0xe7, 0xad, 0x7e, 0x09, 0xa4, 0x38, 0x05, 0x9f, 0x3c, 0x8e, 0xb2, 0x0e, 0xa4, 0xba, 0x3b, 0x98,
	0x3e, 0x7f, 0x38, 0x25, 0xcc, 0xdb, 0xbb, 0x89, 0x09, 0x6f, 0xa0, 0xd1, 0x32, 0xc9, 0x96, 0x2d,
	0xd8, 0x72, 0xba, 0x3a, 0xe3, 0xb5, 0xc6, 0x73, 0xe7, 0x51, 0x26, 0x29, 0x4a, 0xdc, 0xad, 0x61,
	0x76, 0xa0, 0x22, 0x79, 0x17, 0xcd, 0x7f, 0x31, 0x82, 0xf1, 0x97, 0x83, 0x97, 0x30, 0x8c, 0xc6,
	0x69, 0x13, 0x0a, 0x34, 0x11, 0x55, 0xe1, 0x1c, 0xc6, 0xc6, 0x69, 0xf2, 0xbd, 0x4b, 0xed, 0x71,
	0x49, 0xbe, 0x35, 0x5e, 0xc1, 0x24, 0xf5, 0x7b, 0x13, 0x93, 0xdc, 0x53, 0x3b, 0x58, 0xb0, 0xe5,
	0x40, 0xfc, 0x18, 0x88, 0x75, 0x88, 0xa6, 0x50, 0xe5, 0xbd, 0xba, 0x87, 0xd1, 0x93, 0xb7, 0x5a,
	0x90, 0xc2, 0x35, 0x34, 0x0f, 0xd2, 0x5a, 0x3c, 0xff, 0x35, 0x5d, 0x5d, 0x7c, 0x7e, 0xf1, 0xc7,
	0x3d, 0x2c, 0xd6, 0x1d, 0x6d, 0xae, 0xe1, 0xb4, 0xf7, 0x7c, 0x97, 0x93, 0xda, 0xd8, 0xcc, 0xea,
	0x97, 0xdb, 0x7c, 0xca, 0x2d, 0x7b, 0x1d, 0x96, 0x9b, 0xde, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x4f, 0x9d, 0x29, 0xf1, 0x65, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GoldRpcClient is the client API for GoldRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoldRpcClient interface {
	Call(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
}

type goldRpcClient struct {
	cc *grpc.ClientConn
}

func NewGoldRpcClient(cc *grpc.ClientConn) GoldRpcClient {
	return &goldRpcClient{cc}
}

func (c *goldRpcClient) Call(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, "/goldrpc.GoldRpc/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoldRpcServer is the server API for GoldRpc service.
type GoldRpcServer interface {
	Call(context.Context, *SyncRequest) (*SyncResponse, error)
}

func RegisterGoldRpcServer(s *grpc.Server, srv GoldRpcServer) {
	s.RegisterService(&_GoldRpc_serviceDesc, srv)
}

func _GoldRpc_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoldRpcServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goldrpc.GoldRpc/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoldRpcServer).Call(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoldRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goldrpc.GoldRpc",
	HandlerType: (*GoldRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _GoldRpc_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gold_rpc.proto",
}