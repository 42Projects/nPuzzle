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

type Matrix struct {
	Success              bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Rows                 []*Matrix_Row `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Matrix) Reset()         { *m = Matrix{} }
func (m *Matrix) String() string { return proto.CompactTextString(m) }
func (*Matrix) ProtoMessage()    {}
func (*Matrix) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4181303827119de, []int{1}
}

func (m *Matrix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Matrix.Unmarshal(m, b)
}
func (m *Matrix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Matrix.Marshal(b, m, deterministic)
}
func (m *Matrix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Matrix.Merge(m, src)
}
func (m *Matrix) XXX_Size() int {
	return xxx_messageInfo_Matrix.Size(m)
}
func (m *Matrix) XXX_DiscardUnknown() {
	xxx_messageInfo_Matrix.DiscardUnknown(m)
}

var xxx_messageInfo_Matrix proto.InternalMessageInfo

func (m *Matrix) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Matrix) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Matrix) GetRows() []*Matrix_Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type Matrix_Row struct {
	Num                  []uint32 `protobuf:"varint,1,rep,packed,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Matrix_Row) Reset()         { *m = Matrix_Row{} }
func (m *Matrix_Row) String() string { return proto.CompactTextString(m) }
func (*Matrix_Row) ProtoMessage()    {}
func (*Matrix_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4181303827119de, []int{1, 0}
}

func (m *Matrix_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Matrix_Row.Unmarshal(m, b)
}
func (m *Matrix_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Matrix_Row.Marshal(b, m, deterministic)
}
func (m *Matrix_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Matrix_Row.Merge(m, src)
}
func (m *Matrix_Row) XXX_Size() int {
	return xxx_messageInfo_Matrix_Row.Size(m)
}
func (m *Matrix_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Matrix_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Matrix_Row proto.InternalMessageInfo

func (m *Matrix_Row) GetNum() []uint32 {
	if m != nil {
		return m.Num
	}
	return nil
}

type Message struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4181303827119de, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Problem struct {
	Heuristic            string         `protobuf:"bytes,1,opt,name=heuristic,proto3" json:"heuristic,omitempty"`
	Search               string         `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
	Rows                 []*Problem_Row `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Problem) Reset()         { *m = Problem{} }
func (m *Problem) String() string { return proto.CompactTextString(m) }
func (*Problem) ProtoMessage()    {}
func (*Problem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4181303827119de, []int{3}
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

func (m *Problem) GetRows() []*Problem_Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type Problem_Row struct {
	Num                  []uint32 `protobuf:"varint,1,rep,packed,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Problem_Row) Reset()         { *m = Problem_Row{} }
func (m *Problem_Row) String() string { return proto.CompactTextString(m) }
func (*Problem_Row) ProtoMessage()    {}
func (*Problem_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4181303827119de, []int{3, 0}
}

func (m *Problem_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Problem_Row.Unmarshal(m, b)
}
func (m *Problem_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Problem_Row.Marshal(b, m, deterministic)
}
func (m *Problem_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Problem_Row.Merge(m, src)
}
func (m *Problem_Row) XXX_Size() int {
	return xxx_messageInfo_Problem_Row.Size(m)
}
func (m *Problem_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Problem_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Problem_Row proto.InternalMessageInfo

func (m *Problem_Row) GetNum() []uint32 {
	if m != nil {
		return m.Num
	}
	return nil
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
	return fileDescriptor_a4181303827119de, []int{4}
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
	proto.RegisterType((*Matrix)(nil), "npuzzle.Matrix")
	proto.RegisterType((*Matrix_Row)(nil), "npuzzle.Matrix.Row")
	proto.RegisterType((*Message)(nil), "npuzzle.Message")
	proto.RegisterType((*Problem)(nil), "npuzzle.Problem")
	proto.RegisterType((*Problem_Row)(nil), "npuzzle.Problem.Row")
	proto.RegisterType((*Result)(nil), "npuzzle.Result")
}

func init() { proto.RegisterFile("npuzzle.proto", fileDescriptor_a4181303827119de) }

var fileDescriptor_a4181303827119de = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6e, 0xf2, 0x30,
	0x10, 0xc5, 0x5f, 0x48, 0xf2, 0x31, 0x08, 0x15, 0xb9, 0xa8, 0xb5, 0x50, 0x17, 0x51, 0xba, 0x68,
	0x16, 0x55, 0x16, 0xf4, 0x0c, 0x55, 0x57, 0x54, 0xc8, 0x9c, 0x20, 0x44, 0xa3, 0x12, 0x29, 0xc6,
	0xc8, 0x76, 0x80, 0xb2, 0xec, 0x1d, 0x7a, 0x81, 0x9e, 0xb4, 0x8a, 0x6d, 0x0a, 0xfd, 0x59, 0xb4,
	0xbb, 0x79, 0x6f, 0x9e, 0x3c, 0x6f, 0xde, 0x18, 0x06, 0xab, 0x75, 0xb3, 0xdf, 0xd7, 0x98, 0xaf,
	0x95, 0x34, 0x92, 0xc6, 0x1e, 0xa6, 0x31, 0x84, 0xf7, 0x62, 0x6d, 0x9e, 0xd3, 0x3d, 0x44, 0xd3,
	0xc2, 0xa8, 0x6a, 0x47, 0x19, 0xc4, 0xba, 0x29, 0x4b, 0xd4, 0x9a, 0x91, 0x84, 0x64, 0xff, 0xf9,
	0x01, 0xd2, 0x11, 0x84, 0xa8, 0x94, 0x54, 0xec, 0x5f, 0x42, 0xb2, 0x1e, 0x77, 0x80, 0xde, 0x40,
	0x57, 0xc9, 0xad, 0x66, 0x41, 0x12, 0x64, 0xfd, 0xc9, 0x79, 0x7e, 0x98, 0xe4, 0x9e, 0xcb, 0xb9,
	0xdc, 0x72, 0x2b, 0x18, 0x5f, 0x42, 0xc0, 0xe5, 0x96, 0x0e, 0x21, 0x58, 0x35, 0x82, 0x91, 0x24,
	0xc8, 0x06, 0xbc, 0x2d, 0xd3, 0x6b, 0x88, 0xa7, 0xa8, 0x75, 0xf1, 0x84, 0xed, 0x70, 0xe1, 0x4a,
	0x3b, 0xbc, 0xc7, 0x0f, 0x30, 0x7d, 0x21, 0x10, 0xcf, 0x94, 0x5c, 0xd4, 0x28, 0xe8, 0x15, 0xf4,
	0x96, 0xd8, 0xa8, 0x4a, 0x9b, 0xaa, 0xf4, 0xba, 0x23, 0x41, 0x2f, 0x20, 0xd2, 0x58, 0xa8, 0x72,
	0xe9, 0x7d, 0x7a, 0x44, 0xb3, 0x4f, 0x46, 0x47, 0x1f, 0x46, 0xfd, 0xab, 0xbf, 0x71, 0xfa, 0x46,
	0x20, 0xe2, 0xa8, 0x9b, 0xda, 0xfc, 0x39, 0x26, 0x0a, 0x5d, 0x53, 0x09, 0x64, 0x81, 0x25, 0x6d,
	0xdd, 0x2a, 0x85, 0xdc, 0xa0, 0x66, 0xdd, 0x84, 0x64, 0x21, 0x77, 0x80, 0x26, 0xd0, 0x37, 0xd2,
	0x14, 0xf5, 0xdc, 0x14, 0x06, 0x35, 0x0b, 0x6d, 0xef, 0x94, 0x6a, 0xf7, 0x17, 0xc5, 0xce, 0xf7,
	0x23, 0xdb, 0x3f, 0x12, 0x93, 0x57, 0x02, 0xf1, 0xa3, 0xdb, 0x8d, 0xe6, 0x10, 0x3d, 0x28, 0x44,
	0xa3, 0xe9, 0xf0, 0x78, 0x18, 0x97, 0xe8, 0xf8, 0x1b, 0x93, 0x76, 0xe8, 0x2d, 0x84, 0xb3, 0x42,
	0x69, 0xfc, 0x41, 0x7e, 0xf6, 0xe5, 0xb2, 0x4e, 0x3d, 0x97, 0xf5, 0xe6, 0x54, 0xed, 0xc3, 0x3c,
	0x51, 0xbb, 0xbc, 0xd2, 0xce, 0x22, 0xb2, 0x7f, 0xef, 0xee, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x23,
	0xed, 0x52, 0xee, 0x8c, 0x02, 0x00, 0x00,
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
	Greets(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Parse(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Matrix, error)
	Solve(ctx context.Context, in *Problem, opts ...grpc.CallOption) (*Result, error)
}

type npuzzleClient struct {
	cc *grpc.ClientConn
}

func NewNpuzzleClient(cc *grpc.ClientConn) NpuzzleClient {
	return &npuzzleClient{cc}
}

func (c *npuzzleClient) Greets(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/npuzzle.Npuzzle/Greets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *npuzzleClient) Parse(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Matrix, error) {
	out := new(Matrix)
	err := c.cc.Invoke(ctx, "/npuzzle.Npuzzle/Parse", in, out, opts...)
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
	Greets(context.Context, *Message) (*Message, error)
	Parse(context.Context, *Message) (*Matrix, error)
	Solve(context.Context, *Problem) (*Result, error)
}

// UnimplementedNpuzzleServer can be embedded to have forward compatible implementations.
type UnimplementedNpuzzleServer struct {
}

func (*UnimplementedNpuzzleServer) Greets(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greets not implemented")
}
func (*UnimplementedNpuzzleServer) Parse(ctx context.Context, req *Message) (*Matrix, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (*UnimplementedNpuzzleServer) Solve(ctx context.Context, req *Problem) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solve not implemented")
}

func RegisterNpuzzleServer(s *grpc.Server, srv NpuzzleServer) {
	s.RegisterService(&_Npuzzle_serviceDesc, srv)
}

func _Npuzzle_Greets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
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
		return srv.(NpuzzleServer).Greets(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Npuzzle_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NpuzzleServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/npuzzle.Npuzzle/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NpuzzleServer).Parse(ctx, req.(*Message))
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
			MethodName: "Parse",
			Handler:    _Npuzzle_Parse_Handler,
		},
		{
			MethodName: "Solve",
			Handler:    _Npuzzle_Solve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npuzzle.proto",
}
