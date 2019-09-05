// Code generated by protoc-gen-go. DO NOT EDIT.
// source: npuzzle.proto

package npuzzle

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4181303827119de, []int{0}
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

type Handshake struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Handshake) Reset()         { *m = Handshake{} }
func (m *Handshake) String() string { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()    {}
func (*Handshake) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4181303827119de, []int{1}
}

func (m *Handshake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Handshake.Unmarshal(m, b)
}
func (m *Handshake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Handshake.Marshal(b, m, deterministic)
}
func (m *Handshake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handshake.Merge(m, src)
}
func (m *Handshake) XXX_Size() int {
	return xxx_messageInfo_Handshake.Size(m)
}
func (m *Handshake) XXX_DiscardUnknown() {
	xxx_messageInfo_Handshake.DiscardUnknown(m)
}

var xxx_messageInfo_Handshake proto.InternalMessageInfo

func (m *Handshake) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Problem struct {
	Heuristic            string   `protobuf:"bytes,1,opt,name=heuristic,proto3" json:"heuristic,omitempty"`
	Search               string   `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
	Matrix               string   `protobuf:"bytes,3,opt,name=matrix,proto3" json:"matrix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Problem) Reset()         { *m = Problem{} }
func (m *Problem) String() string { return proto.CompactTextString(m) }
func (*Problem) ProtoMessage()    {}
func (*Problem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4181303827119de, []int{2}
}

func (m *Problem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Problem.Unmarshal(m, b)
}
func (m *Problem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Problem.Marshal(b, m, deterministic)
}
func (m *Problem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Problem.Merge(m, src)
}
func (m *Problem) XXX_Size() int {
	return xxx_messageInfo_Problem.Size(m)
}
func (m *Problem) XXX_DiscardUnknown() {
	xxx_messageInfo_Problem.DiscardUnknown(m)
}

var xxx_messageInfo_Problem proto.InternalMessageInfo

func (m *Problem) GetHeuristic() string {
	if m != nil {
		return m.Heuristic
	}
	return ""
}

func (m *Problem) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *Problem) GetMatrix() string {
	if m != nil {
		return m.Matrix
	}
	return ""
}

type Result struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Moves                int32    `protobuf:"varint,4,opt,name=moves,proto3" json:"moves,omitempty"`
	TotalStates          int32    `protobuf:"varint,5,opt,name=totalStates,proto3" json:"totalStates,omitempty"`
	MaxStates            int32    `protobuf:"varint,6,opt,name=maxStates,proto3" json:"maxStates,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4181303827119de, []int{3}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Result) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Result) GetMoves() int32 {
	if m != nil {
		return m.Moves
	}
	return 0
}

func (m *Result) GetTotalStates() int32 {
	if m != nil {
		return m.TotalStates
	}
	return 0
}

func (m *Result) GetMaxStates() int32 {
	if m != nil {
		return m.MaxStates
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "npuzzle.Empty")
	proto.RegisterType((*Handshake)(nil), "npuzzle.Handshake")
	proto.RegisterType((*Problem)(nil), "npuzzle.Problem")
	proto.RegisterType((*Result)(nil), "npuzzle.Result")
}

func init() { proto.RegisterFile("npuzzle.proto", fileDescriptor_a4181303827119de) }

var fileDescriptor_a4181303827119de = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x57, 0x5d, 0x5b, 0xf7, 0x44, 0x94, 0x87, 0x48, 0x18, 0x1e, 0x46, 0x40, 0xd8, 0x41,
	0x76, 0x98, 0x9f, 0x41, 0xf4, 0x24, 0xd2, 0x1d, 0x3c, 0x67, 0xf5, 0xe1, 0xca, 0x9a, 0xa5, 0xe4,
	0xa5, 0x63, 0xee, 0xe3, 0xf8, 0x49, 0xa5, 0x49, 0xda, 0x79, 0xf0, 0x96, 0xdf, 0xef, 0xff, 0x78,
	0x21, 0xff, 0xc0, 0xd5, 0xae, 0x69, 0x8f, 0xc7, 0x9a, 0x16, 0x8d, 0x35, 0xce, 0x60, 0x1e, 0x51,
	0xe6, 0x90, 0x3e, 0xeb, 0xc6, 0x7d, 0xcb, 0x07, 0x98, 0xbc, 0xaa, 0xdd, 0x27, 0x6f, 0xd4, 0x96,
	0x50, 0x40, 0xae, 0x89, 0x59, 0x7d, 0x91, 0x48, 0x66, 0xc9, 0x7c, 0x52, 0xf4, 0x28, 0x3f, 0x20,
	0x7f, 0xb7, 0x66, 0x5d, 0x93, 0xc6, 0x7b, 0x98, 0x6c, 0xa8, 0xb5, 0x15, 0xbb, 0xaa, 0x8c, 0x63,
	0x27, 0x81, 0x77, 0x90, 0x31, 0x29, 0x5b, 0x6e, 0xc4, 0x99, 0x8f, 0x22, 0x75, 0x5e, 0x2b, 0x67,
	0xab, 0x83, 0x38, 0x0f, 0x3e, 0x90, 0xfc, 0x49, 0x20, 0x2b, 0x88, 0xdb, 0xda, 0x75, 0xb7, 0x73,
	0x5b, 0x96, 0xc4, 0xec, 0xd7, 0x5e, 0x14, 0x3d, 0xe2, 0x2d, 0xa4, 0x64, 0xad, 0xb1, 0x71, 0x67,
	0x00, 0x44, 0x18, 0xbb, 0x4a, 0x53, 0x5c, 0xe8, 0xcf, 0xdd, 0xa4, 0x36, 0x7b, 0x62, 0x31, 0x9e,
	0x25, 0xf3, 0xb4, 0x08, 0x80, 0x33, 0xb8, 0x74, 0xc6, 0xa9, 0x7a, 0xe5, 0x94, 0x23, 0x16, 0xa9,
	0xcf, 0xfe, 0xaa, 0xee, 0x51, 0x5a, 0x1d, 0x62, 0x9e, 0xf9, 0xfc, 0x24, 0x96, 0x5b, 0xc8, 0xdf,
	0x42, 0x71, 0xb8, 0x84, 0xec, 0xc5, 0x12, 0x39, 0x46, 0x5c, 0xf4, 0xdd, 0x0e, 0x05, 0x4e, 0xff,
	0x71, 0x72, 0x84, 0x8f, 0x90, 0xae, 0x4c, 0xbd, 0x27, 0xbc, 0x19, 0xe2, 0x58, 0xe6, 0xf4, 0x7a,
	0x30, 0xa1, 0x04, 0x39, 0x5a, 0x67, 0xfe, 0xab, 0x9e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0xbe, 0xa9, 0x4e, 0xbb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NpuzzleClient is the client API for Npuzzle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NpuzzleClient interface {
	Greets(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (*Handshake, error)
	Solve(ctx context.Context, in *Problem, opts ...grpc.CallOption) (*Result, error)
}

type npuzzleClient struct {
	cc *grpc.ClientConn
}

func NewNpuzzleClient(cc *grpc.ClientConn) NpuzzleClient {
	return &npuzzleClient{cc}
}

func (c *npuzzleClient) Greets(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (*Handshake, error) {
	out := new(Handshake)
	err := c.cc.Invoke(ctx, "/npuzzle.Npuzzle/Greets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *npuzzleClient) Solve(ctx context.Context, in *Problem, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/npuzzle.Npuzzle/Solve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NpuzzleServer is the server API for Npuzzle service.
type NpuzzleServer interface {
	Greets(context.Context, *Handshake) (*Handshake, error)
	Solve(context.Context, *Problem) (*Result, error)
}

// UnimplementedNpuzzleServer can be embedded to have forward compatible implementations.
type UnimplementedNpuzzleServer struct {
}

func (*UnimplementedNpuzzleServer) Greets(ctx context.Context, req *Handshake) (*Handshake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greets not implemented")
}
func (*UnimplementedNpuzzleServer) Solve(ctx context.Context, req *Problem) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solve not implemented")
}

func RegisterNpuzzleServer(s *grpc.Server, srv NpuzzleServer) {
	s.RegisterService(&_Npuzzle_serviceDesc, srv)
}

func _Npuzzle_Greets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Handshake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NpuzzleServer).Greets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/npuzzle.Npuzzle/Greets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NpuzzleServer).Greets(ctx, req.(*Handshake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Npuzzle_Solve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Problem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NpuzzleServer).Solve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/npuzzle.Npuzzle/Solve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NpuzzleServer).Solve(ctx, req.(*Problem))
	}
	return interceptor(ctx, in, info, handler)
}

var _Npuzzle_serviceDesc = grpc.ServiceDesc{
	ServiceName: "npuzzle.Npuzzle",
	HandlerType: (*NpuzzleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greets",
			Handler:    _Npuzzle_Greets_Handler,
		},
		{
			MethodName: "Solve",
			Handler:    _Npuzzle_Solve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npuzzle.proto",
}
