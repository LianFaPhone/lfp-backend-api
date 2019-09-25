// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Side int32

const (
	Side_NO_SIDE Side = 0
	Side_BUY     Side = 1
	Side_SELL    Side = -1
)

var Side_name = map[int32]string{
	0:  "NO_SIDE",
	1:  "BUY",
	-1: "SELL",
}
var Side_value = map[string]int32{
	"NO_SIDE": 0,
	"BUY":     1,
	"SELL":    -1,
}

func (x Side) String() string {
	return proto.EnumName(Side_name, int32(x))
}
func (Side) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type State int32

const (
	State_STATE_NONE  State = 0
	State_NEW         State = 1
	State_REVISE      State = 2
	State_CANCEL      State = 3
	State_PARTIALFILL State = 4
	State_TRADE       State = 5
	State_REMOVED     State = 6
)

var State_name = map[int32]string{
	0: "STATE_NONE",
	1: "NEW",
	2: "REVISE",
	3: "CANCEL",
	4: "PARTIALFILL",
	5: "TRADE",
	6: "REMOVED",
}
var State_value = map[string]int32{
	"STATE_NONE":  0,
	"NEW":         1,
	"REVISE":      2,
	"CANCEL":      3,
	"PARTIALFILL": 4,
	"TRADE":       5,
	"REMOVED":     6,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}
func (State) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type Type int32

const (
	Type_TYPE_NONE    Type = 0
	Type_LIMIT        Type = 1
	Type_MARKET       Type = 2
	Type_CANCEL_ORDER Type = 15
)

var Type_name = map[int32]string{
	0:  "TYPE_NONE",
	1:  "LIMIT",
	2:  "MARKET",
	15: "CANCEL_ORDER",
}
var Type_value = map[string]int32{
	"TYPE_NONE":    0,
	"LIMIT":        1,
	"MARKET":       2,
	"CANCEL_ORDER": 15,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type Role int32

const (
	Role_ROLE_NONE Role = 0
	Role_MAKER     Role = 1
	Role_TAKER     Role = 2
)

var Role_name = map[int32]string{
	0: "ROLE_NONE",
	1: "MAKER",
	2: "TAKER",
}
var Role_value = map[string]int32{
	"ROLE_NONE": 0,
	"MAKER":     1,
	"TAKER":     2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}
func (Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterEnum("rpc.Side", Side_name, Side_value)
	proto.RegisterEnum("rpc.State", State_name, State_value)
	proto.RegisterEnum("rpc.Type", Type_name, Type_value)
	proto.RegisterEnum("rpc.Role", Role_name, Role_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x71, 0xf3, 0xb7, 0x64, 0x5a, 0xed, 0x38, 0x8f, 0x11, 0x41, 0x0f, 0xde, 0x85, 0xb5,
	0x19, 0x61, 0xe9, 0x26, 0x5b, 0x66, 0xd7, 0x4a, 0x4f, 0xa1, 0xd6, 0x1c, 0x04, 0xeb, 0x86, 0x92,
	0x8b, 0x4f, 0xaf, 0x34, 0x55, 0xe9, 0x9c, 0x7e, 0x97, 0x2f, 0x1f, 0x18, 0x98, 0xed, 0xc2, 0x7e,
	0x1f, 0x3e, 0x6f, 0xfb, 0x43, 0x18, 0x02, 0x25, 0x87, 0x7e, 0x57, 0xde, 0x41, 0xea, 0xde, 0xdf,
	0x3a, 0x9a, 0xc2, 0xa4, 0xb1, 0xad, 0xd3, 0x15, 0xe3, 0x05, 0x4d, 0x20, 0x79, 0x7c, 0xde, 0x60,
	0x44, 0xd7, 0x90, 0x3a, 0x36, 0x06, 0xbf, 0xff, 0x2e, 0x2a, 0xb7, 0x90, 0xb9, 0x61, 0x3b, 0x74,
	0x74, 0x05, 0xe0, 0xbc, 0xf2, 0xdc, 0x36, 0xb6, 0xf9, 0x8d, 0x1a, 0x7e, 0xc1, 0x88, 0x00, 0x72,
	0xe1, 0xb5, 0x76, 0x8c, 0xf1, 0x71, 0x2f, 0x54, 0xb3, 0x60, 0x83, 0x09, 0xcd, 0x61, 0xba, 0x52,
	0xe2, 0xb5, 0x32, 0x4f, 0xda, 0x18, 0x4c, 0xa9, 0x80, 0xcc, 0x8b, 0xaa, 0x18, 0xb3, 0x23, 0x2f,
	0x5c, 0xdb, 0x35, 0x57, 0x98, 0x97, 0x0f, 0x90, 0xfa, 0xaf, 0xbe, 0xa3, 0x4b, 0x28, 0xfc, 0x66,
	0xf5, 0x0f, 0x14, 0x90, 0x19, 0x5d, 0x6b, 0x7f, 0x22, 0x6a, 0x25, 0x4b, 0xf6, 0x18, 0x13, 0xc2,
	0xec, 0x44, 0xb4, 0x56, 0x2a, 0x16, 0x9c, 0x97, 0x37, 0x90, 0x4a, 0xf8, 0x18, 0x7b, 0xb1, 0xe6,
	0xbc, 0xaf, 0xd5, 0x92, 0x05, 0xa3, 0x51, 0x1e, 0x67, 0xfc, 0x9a, 0x8f, 0xcf, 0xb8, 0xff, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x2b, 0xbf, 0xba, 0x26, 0x1c, 0x01, 0x00, 0x00,
}
