// Code generated by protoc-gen-go.
// source: ExchangeMessage.proto
// DO NOT EDIT!

/*
Package exchange_message_def is a generated protocol buffer package.

It is generated from these files:
	ExchangeMessage.proto

It has these top-level messages:
	Person
	Manager
*/
package exchange_message_def

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person struct {
	UserName        string                     `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	FavouriteNumber int64                      `protobuf:"varint,2,opt,name=favourite_number,json=favouriteNumber" json:"favourite_number,omitempty"`
	Company         string                     `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	Date            *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=date" json:"date,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Person) GetFavouriteNumber() int64 {
	if m != nil {
		return m.FavouriteNumber
	}
	return 0
}

func (m *Person) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Person) GetDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type Manager struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Company  string `protobuf:"bytes,2,opt,name=company" json:"company,omitempty"`
	Stars    string `protobuf:"bytes,3,opt,name=stars" json:"stars,omitempty"`
}

func (m *Manager) Reset()                    { *m = Manager{} }
func (m *Manager) String() string            { return proto.CompactTextString(m) }
func (*Manager) ProtoMessage()               {}
func (*Manager) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Manager) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Manager) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Manager) GetStars() string {
	if m != nil {
		return m.Stars
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "exchange_message_def.Person")
	proto.RegisterType((*Manager)(nil), "exchange_message_def.Manager")
}

func init() { proto.RegisterFile("ExchangeMessage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xc9, 0x36, 0x37, 0xf7, 0x89, 0x4c, 0xc2, 0x84, 0x30, 0x0f, 0x96, 0x9e, 0xea, 0x25,
	0x03, 0xfd, 0x07, 0x03, 0xc1, 0xcb, 0x86, 0x14, 0xf1, 0x5a, 0xbe, 0x6e, 0x5f, 0xe3, 0x98, 0x49,
	0x4a, 0x92, 0x8a, 0xfe, 0x15, 0x7f, 0xad, 0x98, 0xb6, 0xa2, 0x1e, 0x76, 0xfc, 0x1e, 0xde, 0xe4,
	0x7d, 0x78, 0xe1, 0xf2, 0xfe, 0x7d, 0xfb, 0x82, 0x46, 0xd1, 0x9a, 0xbc, 0x47, 0x45, 0xb2, 0x76,
	0x36, 0x58, 0x3e, 0xa7, 0x0e, 0x17, 0xba, 0xe5, 0xc5, 0x8e, 0xaa, 0xc5, 0xb5, 0xb2, 0x56, 0xbd,
	0xd2, 0x32, 0x66, 0xca, 0xa6, 0x5a, 0x86, 0xbd, 0x26, 0x1f, 0x50, 0xd7, 0xed, 0xb3, 0xf4, 0x93,
	0xc1, 0xf8, 0x91, 0x9c, 0xb7, 0x86, 0x5f, 0xc1, 0xb4, 0xf1, 0xe4, 0x0a, 0x83, 0x9a, 0x04, 0x4b,
	0x58, 0x36, 0xcd, 0x4f, 0xbf, 0xc1, 0x06, 0x35, 0xf1, 0x1b, 0xb8, 0xa8, 0xf0, 0xcd, 0x36, 0x6e,
	0x1f, 0xa8, 0x30, 0x8d, 0x2e, 0xc9, 0x89, 0x41, 0xc2, 0xb2, 0x61, 0x3e, 0xfb, 0xe1, 0x9b, 0x88,
	0xb9, 0x80, 0xc9, 0xd6, 0xea, 0x1a, 0xcd, 0x87, 0x18, 0xc6, 0x5f, 0xfa, 0x93, 0x4b, 0x18, 0xed,
	0x30, 0x90, 0x18, 0x25, 0x2c, 0x3b, 0xbb, 0x5d, 0xc8, 0x56, 0x4e, 0xf6, 0x72, 0xf2, 0xa9, 0x97,
	0xcb, 0x63, 0x2e, 0x7d, 0x86, 0xc9, 0x1a, 0x0d, 0x2a, 0x72, 0xc7, 0xe5, 0x7e, 0x35, 0x0e, 0xfe,
	0x36, 0xce, 0xe1, 0xc4, 0x07, 0x74, 0xbe, 0x33, 0x69, 0x8f, 0x55, 0x0a, 0xe7, 0x07, 0xac, 0x0e,
	0x28, 0xbb, 0xa9, 0x56, 0xb3, 0x7f, 0x9b, 0x3e, 0xb0, 0x72, 0x1c, 0xad, 0xee, 0xbe, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0x0d, 0x14, 0x87, 0x6f, 0x01, 0x00, 0x00,
}
