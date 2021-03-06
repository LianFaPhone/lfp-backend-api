// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Error int32

const (
	Error_OK                  Error = 0
	Error_INVALID_REQUEST     Error = 1
	Error_UNIMPLEMENTED       Error = 2
	Error_INVALID_ARGUMENT    Error = 3
	Error_UNSURPPORT_TOPIC    Error = 4
	Error_UNSURPPORT_CHANNEL  Error = 5
	Error_ALREADY_EXISTS      Error = 6
	Error_PERMISSION_DENIED   Error = 7
	Error_UNAUTHENTICATED     Error = 8
	Error_BAD_MEMORY_ALLOCATE Error = 9
)

var Error_name = map[int32]string{
	0: "OK",
	1: "INVALID_REQUEST",
	2: "UNIMPLEMENTED",
	3: "INVALID_ARGUMENT",
	4: "UNSURPPORT_TOPIC",
	5: "UNSURPPORT_CHANNEL",
	6: "ALREADY_EXISTS",
	7: "PERMISSION_DENIED",
	8: "UNAUTHENTICATED",
	9: "BAD_MEMORY_ALLOCATE",
}
var Error_value = map[string]int32{
	"OK":                  0,
	"INVALID_REQUEST":     1,
	"UNIMPLEMENTED":       2,
	"INVALID_ARGUMENT":    3,
	"UNSURPPORT_TOPIC":    4,
	"UNSURPPORT_CHANNEL":  5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"UNAUTHENTICATED":     8,
	"BAD_MEMORY_ALLOCATE": 9,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}
func (Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterEnum("rpc.Error", Error_name, Error_value)
}

func init() { proto.RegisterFile("error.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x4b, 0x4e, 0x03, 0x31,
	0x0c, 0xc6, 0x71, 0xda, 0xd2, 0x01, 0x8c, 0x00, 0xd7, 0xe5, 0x71, 0x07, 0x16, 0x6c, 0x38, 0x41,
	0x98, 0x58, 0x34, 0x22, 0x71, 0x42, 0x1e, 0x88, 0xae, 0x22, 0x81, 0x58, 0xcf, 0x28, 0xe2, 0xa8,
	0x1c, 0x08, 0x05, 0x09, 0xa9, 0xdb, 0x9f, 0xbd, 0xf8, 0x7f, 0x70, 0xfe, 0xd5, 0xda, 0xd4, 0x1e,
	0xe6, 0x36, 0x7d, 0x4f, 0xb4, 0x6a, 0xf3, 0xe7, 0xfd, 0xcf, 0x02, 0xd6, 0xdc, 0x91, 0x06, 0x58,
	0xfa, 0x17, 0x3c, 0xa2, 0x2d, 0x5c, 0x19, 0x79, 0x53, 0xd6, 0xe8, 0x1a, 0xf9, 0xb5, 0x70, 0xca,
	0xb8, 0xa0, 0x0d, 0x5c, 0x14, 0x31, 0x2e, 0x58, 0x76, 0x2c, 0x99, 0x35, 0x2e, 0xe9, 0x1a, 0xf0,
	0xff, 0x4f, 0xc5, 0xe7, 0xd2, 0x1d, 0x57, 0x5d, 0x8b, 0xa4, 0x12, 0x43, 0xf0, 0x31, 0xd7, 0xec,
	0x83, 0x19, 0xf1, 0x98, 0x6e, 0x81, 0x0e, 0x74, 0xdc, 0x29, 0x11, 0xb6, 0xb8, 0x26, 0x82, 0x4b,
	0x65, 0x23, 0x2b, 0xbd, 0xaf, 0xfc, 0x6e, 0x52, 0x4e, 0x38, 0xd0, 0x0d, 0x6c, 0x02, 0x47, 0x67,
	0x52, 0x32, 0x5e, 0xaa, 0x66, 0x31, 0xac, 0xf1, 0xa4, 0x67, 0x15, 0x51, 0x25, 0xef, 0x58, 0xb2,
	0x19, 0x55, 0x6f, 0x38, 0xa5, 0x3b, 0xd8, 0x3e, 0x29, 0x5d, 0x1d, 0x3b, 0x1f, 0xf7, 0x55, 0x59,
	0xeb, 0xfb, 0x05, 0xcf, 0x3e, 0x86, 0xbf, 0x89, 0x8f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x68,
	0xa2, 0xd8, 0xda, 0xf1, 0x00, 0x00, 0x00,
}
