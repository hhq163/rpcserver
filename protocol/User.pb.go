// Code generated by protoc-gen-go. DO NOT EDIT.
// source: User.proto

package protocol

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

type CashType int32

const (
	CashType_UNKNOWN      CashType = 0
	CashType_RECHARGE     CashType = 1
	CashType_DEDUCT_MONEY CashType = 2
	CashType_ADD_GOLD     CashType = 3
	CashType_PAYOUT       CashType = 4
)

var CashType_name = map[int32]string{
	0: "UNKNOWN",
	1: "RECHARGE",
	2: "DEDUCT_MONEY",
	3: "ADD_GOLD",
	4: "PAYOUT",
}

var CashType_value = map[string]int32{
	"UNKNOWN":      0,
	"RECHARGE":     1,
	"DEDUCT_MONEY": 2,
	"ADD_GOLD":     3,
	"PAYOUT":       4,
}

func (x CashType) String() string {
	return proto.EnumName(CashType_name, int32(x))
}

func (CashType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{0}
}

//现金操作请求
type CashOperRequest struct {
	HallId               int32    `protobuf:"varint,1,opt,name=hall_id,json=hallId,proto3" json:"hall_id,omitempty"`
	AgentId              int32    `protobuf:"varint,2,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	HallName             string   `protobuf:"bytes,4,opt,name=hall_name,json=hallName,proto3" json:"hall_name,omitempty"`
	UserName             string   `protobuf:"bytes,5,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Amount               float64  `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Type                 CashType `protobuf:"varint,7,opt,name=type,proto3,enum=protocol.CashType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CashOperRequest) Reset()         { *m = CashOperRequest{} }
func (m *CashOperRequest) String() string { return proto.CompactTextString(m) }
func (*CashOperRequest) ProtoMessage()    {}
func (*CashOperRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{0}
}

func (m *CashOperRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CashOperRequest.Unmarshal(m, b)
}
func (m *CashOperRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CashOperRequest.Marshal(b, m, deterministic)
}
func (m *CashOperRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CashOperRequest.Merge(m, src)
}
func (m *CashOperRequest) XXX_Size() int {
	return xxx_messageInfo_CashOperRequest.Size(m)
}
func (m *CashOperRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CashOperRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CashOperRequest proto.InternalMessageInfo

func (m *CashOperRequest) GetHallId() int32 {
	if m != nil {
		return m.HallId
	}
	return 0
}

func (m *CashOperRequest) GetAgentId() int32 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *CashOperRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CashOperRequest) GetHallName() string {
	if m != nil {
		return m.HallName
	}
	return ""
}

func (m *CashOperRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CashOperRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CashOperRequest) GetType() CashType {
	if m != nil {
		return m.Type
	}
	return CashType_UNKNOWN
}

//现金操作响应
type CashOperResponse struct {
	ResultCode           int32    `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Restult              *Result  `protobuf:"bytes,3,opt,name=restult,proto3" json:"restult,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CashOperResponse) Reset()         { *m = CashOperResponse{} }
func (m *CashOperResponse) String() string { return proto.CompactTextString(m) }
func (*CashOperResponse) ProtoMessage()    {}
func (*CashOperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{1}
}

func (m *CashOperResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CashOperResponse.Unmarshal(m, b)
}
func (m *CashOperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CashOperResponse.Marshal(b, m, deterministic)
}
func (m *CashOperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CashOperResponse.Merge(m, src)
}
func (m *CashOperResponse) XXX_Size() int {
	return xxx_messageInfo_CashOperResponse.Size(m)
}
func (m *CashOperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CashOperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CashOperResponse proto.InternalMessageInfo

func (m *CashOperResponse) GetResultCode() int32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

func (m *CashOperResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *CashOperResponse) GetRestult() *Result {
	if m != nil {
		return m.Restult
	}
	return nil
}

type Result struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	OrderSn              string   `protobuf:"bytes,2,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_979821478719c248, []int{2}
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

func (m *Result) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Result) GetOrderSn() string {
	if m != nil {
		return m.OrderSn
	}
	return ""
}

func init() {
	proto.RegisterEnum("protocol.CashType", CashType_name, CashType_value)
	proto.RegisterType((*CashOperRequest)(nil), "protocol.CashOperRequest")
	proto.RegisterType((*CashOperResponse)(nil), "protocol.CashOperResponse")
	proto.RegisterType((*Result)(nil), "protocol.Result")
}

func init() { proto.RegisterFile("User.proto", fileDescriptor_979821478719c248) }

var fileDescriptor_979821478719c248 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0xbb, 0x8d, 0x6b, 0x3b, 0x93, 0x0a, 0x56, 0x73, 0x80, 0xa4, 0x1c, 0xb0, 0x72, 0x40,
	0x56, 0x0f, 0x39, 0x84, 0x23, 0xa7, 0x28, 0x8e, 0x4a, 0x44, 0xb1, 0x61, 0x49, 0x84, 0x7a, 0xb2,
	0x4c, 0x76, 0x44, 0x91, 0x1c, 0xaf, 0xd9, 0x5d, 0x1f, 0xfa, 0x4f, 0xf9, 0x39, 0x68, 0xd7, 0xb1,
	0xda, 0x4a, 0x3d, 0xd9, 0xef, 0x7d, 0x1e, 0xcf, 0xbc, 0x19, 0x80, 0xbd, 0x21, 0xbd, 0x68, 0xb5,
	0xb2, 0x0a, 0x63, 0xff, 0x38, 0xa8, 0x7a, 0xfe, 0x8f, 0xc1, 0xeb, 0x75, 0x65, 0xee, 0x8b, 0x96,
	0xb4, 0xa0, 0xbf, 0x1d, 0x19, 0x8b, 0x6f, 0x21, 0xba, 0xaf, 0xea, 0xba, 0xfc, 0x23, 0xa7, 0x2c,
	0x61, 0xe9, 0x85, 0x08, 0x9d, 0xdc, 0x4a, 0x9c, 0x41, 0x5c, 0xfd, 0xa6, 0xc6, 0x3a, 0x72, 0xee,
	0x49, 0xe4, 0xf5, 0x56, 0xba, 0x9a, 0xce, 0x90, 0x76, 0x64, 0x94, 0xb0, 0x74, 0x24, 0x42, 0x27,
	0xb7, 0x12, 0xdf, 0xc1, 0xd8, 0xff, 0xac, 0xa9, 0x8e, 0x34, 0x0d, 0x12, 0x96, 0x8e, 0x45, 0xec,
	0x8c, 0xbc, 0x3a, 0x92, 0x83, 0xbe, 0xca, 0xc3, 0x8b, 0x1e, 0x3a, 0xc3, 0xc3, 0x37, 0x10, 0x56,
	0x47, 0xd5, 0x35, 0x76, 0x1a, 0x26, 0x2c, 0x65, 0xe2, 0xa4, 0xf0, 0x03, 0x04, 0xf6, 0xa1, 0xa5,
	0x69, 0x94, 0xb0, 0xf4, 0xd5, 0x12, 0x17, 0x43, 0x96, 0x85, 0xcb, 0xb1, 0x7b, 0x68, 0x49, 0x78,
	0x3e, 0x37, 0xc0, 0x1f, 0x93, 0x99, 0x56, 0x35, 0x86, 0xf0, 0x3d, 0x4c, 0x34, 0x99, 0xae, 0xb6,
	0xe5, 0x41, 0x49, 0x3a, 0xc5, 0x83, 0xde, 0x5a, 0x2b, 0x49, 0x88, 0x10, 0x48, 0x32, 0x07, 0x1f,
	0x6f, 0x2c, 0xfc, 0x3b, 0x5e, 0x43, 0xa4, 0xc9, 0xd8, 0xae, 0xb6, 0x3e, 0xdb, 0x64, 0xc9, 0x1f,
	0x7b, 0x0a, 0x5f, 0x2a, 0x86, 0x0f, 0xe6, 0x9f, 0x20, 0xec, 0xad, 0x27, 0xe3, 0xb3, 0x67, 0xe3,
	0xcf, 0x20, 0x56, 0x5a, 0x92, 0x2e, 0x4d, 0x73, 0xea, 0x12, 0x79, 0xfd, 0xa3, 0xb9, 0xfe, 0x0e,
	0xf1, 0x90, 0x01, 0x27, 0x10, 0xed, 0xf3, 0x2f, 0x79, 0xf1, 0x33, 0xe7, 0x67, 0x78, 0x09, 0xb1,
	0xd8, 0xac, 0x3f, 0xaf, 0xc4, 0xcd, 0x86, 0x33, 0xe4, 0x70, 0x99, 0x6d, 0xb2, 0xfd, 0x7a, 0x57,
	0x7e, 0x2d, 0xf2, 0xcd, 0x1d, 0x3f, 0x77, 0x7c, 0x95, 0x65, 0xe5, 0x4d, 0x71, 0x9b, 0xf1, 0x11,
	0x02, 0x84, 0xdf, 0x56, 0x77, 0xc5, 0x7e, 0xc7, 0x83, 0xe5, 0x2d, 0x04, 0xee, 0xee, 0x98, 0xc1,
	0x78, 0x58, 0x46, 0x85, 0xb3, 0xe7, 0x3b, 0x7b, 0x72, 0xfb, 0xab, 0xab, 0x97, 0x50, 0xbf, 0xbc,
	0xf9, 0xd9, 0xaf, 0xd0, 0xc3, 0x8f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x1d, 0xdf, 0xa3,
	0x4c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	CashOpera(ctx context.Context, in *CashOperRequest, opts ...grpc.CallOption) (*CashOperResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) CashOpera(ctx context.Context, in *CashOperRequest, opts ...grpc.CallOption) (*CashOperResponse, error) {
	out := new(CashOperResponse)
	err := c.cc.Invoke(ctx, "/protocol.User/CashOpera", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	CashOpera(context.Context, *CashOperRequest) (*CashOperResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) CashOpera(ctx context.Context, req *CashOperRequest) (*CashOperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CashOpera not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_CashOpera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashOperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CashOpera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.User/CashOpera",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CashOpera(ctx, req.(*CashOperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CashOpera",
			Handler:    _User_CashOpera_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "User.proto",
}
