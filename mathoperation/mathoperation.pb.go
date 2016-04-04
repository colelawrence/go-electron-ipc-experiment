// Code generated by protoc-gen-go.
// source: mathoperation/mathoperation.proto
// DO NOT EDIT!

/*
Package mathoperation is a generated protocol buffer package.

It is generated from these files:
	mathoperation/mathoperation.proto

It has these top-level messages:
	MathOperation
	MathResult
*/
package mathoperation

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MathOperation_OperationType int32

const (
	MathOperation_MULTIPLY MathOperation_OperationType = 0
	MathOperation_DIVIDE   MathOperation_OperationType = 1
	MathOperation_ADD      MathOperation_OperationType = 2
	MathOperation_SUBTRACT MathOperation_OperationType = 3
)

var MathOperation_OperationType_name = map[int32]string{
	0: "MULTIPLY",
	1: "DIVIDE",
	2: "ADD",
	3: "SUBTRACT",
}
var MathOperation_OperationType_value = map[string]int32{
	"MULTIPLY": 0,
	"DIVIDE":   1,
	"ADD":      2,
	"SUBTRACT": 3,
}

func (x MathOperation_OperationType) String() string {
	return proto.EnumName(MathOperation_OperationType_name, int32(x))
}
func (MathOperation_OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type MathOperation struct {
	Id            string                      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	LeftHandSide  float32                     `protobuf:"fixed32,2,opt,name=left_hand_side,json=leftHandSide" json:"left_hand_side,omitempty"`
	RightHandSide float32                     `protobuf:"fixed32,3,opt,name=right_hand_side,json=rightHandSide" json:"right_hand_side,omitempty"`
	Operation     MathOperation_OperationType `protobuf:"varint,4,opt,name=operation,enum=mathoperation.MathOperation_OperationType" json:"operation,omitempty"`
}

func (m *MathOperation) Reset()                    { *m = MathOperation{} }
func (m *MathOperation) String() string            { return proto.CompactTextString(m) }
func (*MathOperation) ProtoMessage()               {}
func (*MathOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MathResult struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Result float32 `protobuf:"fixed32,2,opt,name=result" json:"result,omitempty"`
	Error  string  `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *MathResult) Reset()                    { *m = MathResult{} }
func (m *MathResult) String() string            { return proto.CompactTextString(m) }
func (*MathResult) ProtoMessage()               {}
func (*MathResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*MathOperation)(nil), "mathoperation.MathOperation")
	proto.RegisterType((*MathResult)(nil), "mathoperation.MathResult")
	proto.RegisterEnum("mathoperation.MathOperation_OperationType", MathOperation_OperationType_name, MathOperation_OperationType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Calculator service

type CalculatorClient interface {
	// Computes the value
	Compute(ctx context.Context, in *MathOperation, opts ...grpc.CallOption) (*MathResult, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Compute(ctx context.Context, in *MathOperation, opts ...grpc.CallOption) (*MathResult, error) {
	out := new(MathResult)
	err := grpc.Invoke(ctx, "/mathoperation.Calculator/Compute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calculator service

type CalculatorServer interface {
	// Computes the value
	Compute(context.Context, *MathOperation) (*MathResult, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MathOperation)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CalculatorServer).Compute(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mathoperation.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _Calculator_Compute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x9b, 0x54, 0x53, 0x33, 0x34, 0x31, 0x0c, 0x22, 0x51, 0x3c, 0x68, 0x10, 0x11, 0x0f,
	0x11, 0xea, 0x1f, 0xb0, 0x26, 0x42, 0x23, 0x2d, 0xca, 0x36, 0x15, 0x3c, 0x95, 0x68, 0x56, 0x13,
	0x48, 0xbb, 0x61, 0xbb, 0x39, 0xf8, 0xd3, 0xbd, 0xb9, 0x49, 0x6b, 0x74, 0x51, 0xbc, 0xcd, 0x7c,
	0xef, 0x2d, 0xf3, 0x1e, 0x0b, 0x27, 0x8b, 0x44, 0x64, 0xac, 0xa4, 0x3c, 0x11, 0x39, 0x5b, 0x5e,
	0x2a, 0x9b, 0x5f, 0x72, 0x26, 0x18, 0x5a, 0x0a, 0xf4, 0x3e, 0x34, 0xb0, 0x26, 0x92, 0xdc, 0x7f,
	0x11, 0xb4, 0x41, 0xcf, 0x53, 0x57, 0x3b, 0xd6, 0xce, 0x4d, 0x22, 0x27, 0x3c, 0x05, 0xbb, 0xa0,
	0xaf, 0x62, 0x9e, 0x25, 0xcb, 0x74, 0xbe, 0xca, 0x53, 0xea, 0xea, 0x52, 0xd3, 0x49, 0xbf, 0xa6,
	0x23, 0x09, 0xa7, 0x92, 0xe1, 0x19, 0xec, 0xf2, 0xfc, 0x2d, 0xfb, 0x69, 0xeb, 0x36, 0x36, 0xab,
	0xc1, 0xad, 0x6f, 0x04, 0x66, 0x7b, 0xdc, 0xdd, 0x92, 0x0e, 0x7b, 0x70, 0xe1, 0xab, 0x39, 0x95,
	0x38, 0x7e, 0x3b, 0xc5, 0xef, 0x25, 0x25, 0xdf, 0x8f, 0xbd, 0x6b, 0xb0, 0x14, 0x0d, 0xfb, 0xb0,
	0x33, 0x99, 0x8d, 0xe3, 0xe8, 0x61, 0xfc, 0xe4, 0x74, 0x10, 0xc0, 0x08, 0xa3, 0xc7, 0x28, 0xbc,
	0x75, 0x34, 0xec, 0x41, 0x77, 0x18, 0x86, 0x8e, 0x5e, 0x5b, 0xa6, 0xb3, 0x9b, 0x98, 0x0c, 0x83,
	0xd8, 0xe9, 0x7a, 0x77, 0x00, 0xf5, 0x2d, 0x42, 0x57, 0x55, 0x21, 0x7e, 0xf5, 0xde, 0x07, 0x83,
	0x37, 0xca, 0xa6, 0xef, 0x66, 0xc3, 0x3d, 0xd8, 0xa6, 0x9c, 0x33, 0xde, 0xf4, 0x33, 0xc9, 0x7a,
	0x19, 0x10, 0x80, 0x20, 0x29, 0x5e, 0xaa, 0x22, 0x11, 0x8c, 0x63, 0x08, 0xbd, 0x80, 0x2d, 0xca,
	0x4a, 0x50, 0x3c, 0xfa, 0xaf, 0xdd, 0xe1, 0xc1, 0x1f, 0xea, 0x3a, 0x8f, 0xd7, 0x79, 0x36, 0x9a,
	0x1f, 0xbb, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x79, 0x4f, 0x0e, 0x28, 0xd6, 0x01, 0x00, 0x00,
}