// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/user.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
	StatusResponse
	CreateUserRequest
	User
	Empty
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto1.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto1.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type StatusResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto1.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StatusResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateUserRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto1.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto1.RegisterType((*HelloRequest)(nil), "proto.HelloRequest")
	proto1.RegisterType((*HelloResponse)(nil), "proto.HelloResponse")
	proto1.RegisterType((*StatusResponse)(nil), "proto.StatusResponse")
	proto1.RegisterType((*CreateUserRequest)(nil), "proto.CreateUserRequest")
	proto1.RegisterType((*User)(nil), "proto.User")
	proto1.RegisterType((*Empty)(nil), "proto.Empty")
}

func init() { proto1.RegisterFile("proto/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x6d, 0x6a, 0x62, 0xed, 0x54, 0x8b, 0x8e, 0x15, 0x42, 0x45, 0x94, 0x3d, 0x09, 0x42, 0x85,
	0xda, 0xb3, 0x1e, 0x44, 0xf0, 0xe4, 0xa1, 0xe2, 0x59, 0xd6, 0x66, 0x2c, 0x81, 0xa4, 0x1b, 0x77,
	0x36, 0x8a, 0xf7, 0x7e, 0xb8, 0x64, 0x76, 0x53, 0x2d, 0x9e, 0x32, 0x6f, 0xde, 0xcb, 0xec, 0x7b,
	0x0f, 0x0e, 0x2b, 0x6b, 0x9c, 0xb9, 0xae, 0x99, 0xec, 0x44, 0x46, 0x4c, 0xe4, 0xa3, 0x14, 0xec,
	0x3f, 0x52, 0x51, 0x98, 0x39, 0x7d, 0xd4, 0xc4, 0x0e, 0x11, 0xe2, 0x95, 0x2e, 0x29, 0x8d, 0x2e,
	0xa2, 0xcb, 0xfe, 0x5c, 0x66, 0x75, 0x05, 0x07, 0x41, 0xc3, 0x95, 0x59, 0x31, 0xe1, 0x18, 0xf6,
	0x96, 0x96, 0xc8, 0xe5, 0xab, 0x65, 0xda, 0x15, 0xe1, 0x06, 0xab, 0x5b, 0x18, 0x3e, 0x3b, 0xed,
	0x6a, 0xde, 0xa8, 0x11, 0xe2, 0x85, 0xc9, 0xfc, 0xc9, 0x64, 0x2e, 0x33, 0xa6, 0xd0, 0x2b, 0x89,
	0x59, 0x2f, 0x29, 0x1c, 0x68, 0xa1, 0x9a, 0xc1, 0xd1, 0xbd, 0x25, 0xed, 0xe8, 0x85, 0xc9, 0xb6,
	0xae, 0xce, 0x21, 0x6e, 0xac, 0xcb, 0x89, 0xc1, 0x74, 0xe0, 0x23, 0x4c, 0x44, 0x21, 0x84, 0x5a,
	0x47, 0x10, 0x37, 0x10, 0x87, 0xd0, 0xcd, 0xb3, 0xe0, 0xbe, 0x9b, 0x67, 0x78, 0x06, 0xf0, 0x9e,
	0x5b, 0x76, 0xaf, 0x92, 0xca, 0xbf, 0xd5, 0x97, 0xcd, 0x93, 0x2e, 0x09, 0x4f, 0xa1, 0x5f, 0xe8,
	0x96, 0xdd, 0xf1, 0x51, 0x9a, 0x85, 0x90, 0x23, 0x48, 0xa8, 0xd4, 0x79, 0x91, 0xc6, 0x42, 0x78,
	0xd0, 0x84, 0xaf, 0x34, 0xf3, 0x97, 0xb1, 0x59, 0x9a, 0xf8, 0x3f, 0x5a, 0xac, 0x7a, 0x90, 0x3c,
	0x94, 0x95, 0xfb, 0x9e, 0xae, 0x23, 0x18, 0x34, 0x7e, 0x98, 0xec, 0x67, 0xbe, 0x20, 0x9c, 0x41,
	0x22, 0x15, 0xe2, 0x71, 0xf0, 0xfe, 0xb7, 0xf4, 0xf1, 0x68, 0x7b, 0xe9, 0x7b, 0x53, 0x1d, 0xbc,
	0x03, 0xf8, 0xed, 0x02, 0xd3, 0xa0, 0xfa, 0x57, 0xcf, 0xf8, 0x24, 0x30, 0xdb, 0xc5, 0xab, 0xce,
	0xdb, 0xae, 0xec, 0x6f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x63, 0xc0, 0xda, 0x7d, 0xff, 0x01,
	0x00, 0x00,
}
