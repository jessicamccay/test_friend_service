// Code generated by protoc-gen-go.
// source: friends.uacf.io/apps/friendsapi/rpc/friendsapi.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	friends.uacf.io/apps/friendsapi/rpc/friendsapi.proto

It has these top-level messages:
	Friendship
	CreateRequest
	CreateResponse
	GetRequest
	GetResponse
	ListRequest
	ListResponse
	MethodRequest
	MethodResponse
	PingRequest
	PongResponse
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "third_party/google/api"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Messages for resource operations (see service notes below)
type Friendship struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	FromUserId   string `protobuf:"bytes,2,opt,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"`
	ToUserId     string `protobuf:"bytes,3,opt,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
	FriendsSince string `protobuf:"bytes,4,opt,name=friends_since,json=friendsSince" json:"friends_since,omitempty"`
	Status       string `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *Friendship) Reset()                    { *m = Friendship{} }
func (m *Friendship) String() string            { return proto.CompactTextString(m) }
func (*Friendship) ProtoMessage()               {}
func (*Friendship) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateRequest struct {
	Friendship *Friendship `protobuf:"bytes,1,opt,name=friendship" json:"friendship,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetFriendship() *Friendship {
	if m != nil {
		return m.Friendship
	}
	return nil
}

type CreateResponse struct {
	Friendship *Friendship `protobuf:"bytes,1,opt,name=friendship" json:"friendship,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateResponse) GetFriendship() *Friendship {
	if m != nil {
		return m.Friendship
	}
	return nil
}

type GetRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetResponse struct {
	Friendship *Friendship `protobuf:"bytes,1,opt,name=friendship" json:"friendship,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetResponse) GetFriendship() *Friendship {
	if m != nil {
		return m.Friendship
	}
	return nil
}

type ListRequest struct {
	Status     string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	FromUserId string `protobuf:"bytes,2,opt,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"`
	ToUserId   string `protobuf:"bytes,3,opt,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListResponse struct {
	Friendships []*Friendship `protobuf:"bytes,1,rep,name=friendships" json:"friendships,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListResponse) GetFriendships() []*Friendship {
	if m != nil {
		return m.Friendships
	}
	return nil
}

// Messages for imperative operations (see service notes below)
type MethodRequest struct {
	Parameter1 string `protobuf:"bytes,1,opt,name=parameter1" json:"parameter1,omitempty"`
	Parameter2 string `protobuf:"bytes,2,opt,name=parameter2" json:"parameter2,omitempty"`
}

func (m *MethodRequest) Reset()                    { *m = MethodRequest{} }
func (m *MethodRequest) String() string            { return proto.CompactTextString(m) }
func (*MethodRequest) ProtoMessage()               {}
func (*MethodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type MethodResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *MethodResponse) Reset()                    { *m = MethodResponse{} }
func (m *MethodResponse) String() string            { return proto.CompactTextString(m) }
func (*MethodResponse) ProtoMessage()               {}
func (*MethodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// Ping request/response
type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type PongResponse struct {
	Pong bool `protobuf:"varint,1,opt,name=pong" json:"pong,omitempty"`
}

func (m *PongResponse) Reset()                    { *m = PongResponse{} }
func (m *PongResponse) String() string            { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()               {}
func (*PongResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*Friendship)(nil), "io.uacf.friends.Friendship")
	proto.RegisterType((*CreateRequest)(nil), "io.uacf.friends.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "io.uacf.friends.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "io.uacf.friends.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "io.uacf.friends.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "io.uacf.friends.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "io.uacf.friends.ListResponse")
	proto.RegisterType((*MethodRequest)(nil), "io.uacf.friends.MethodRequest")
	proto.RegisterType((*MethodResponse)(nil), "io.uacf.friends.MethodResponse")
	proto.RegisterType((*PingRequest)(nil), "io.uacf.friends.PingRequest")
	proto.RegisterType((*PongResponse)(nil), "io.uacf.friends.PongResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for FriendsApiService service

type FriendsApiServiceClient interface {
	// -------------------------
	// Ping operation for determining service health
	// -------------------------
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type friendsApiServiceClient struct {
	cc *grpc.ClientConn
}

func NewFriendsApiServiceClient(cc *grpc.ClientConn) FriendsApiServiceClient {
	return &friendsApiServiceClient{cc}
}

func (c *friendsApiServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := grpc.Invoke(ctx, "/io.uacf.friends.FriendsApiService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsApiServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/io.uacf.friends.FriendsApiService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsApiServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/io.uacf.friends.FriendsApiService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FriendsApiService service

type FriendsApiServiceServer interface {
	// -------------------------
	// Ping operation for determining service health
	// -------------------------
	Ping(context.Context, *PingRequest) (*PongResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterFriendsApiServiceServer(s *grpc.Server, srv FriendsApiServiceServer) {
	s.RegisterService(&_FriendsApiService_serviceDesc, srv)
}

func _FriendsApiService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsApiServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.uacf.friends.FriendsApiService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsApiServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsApiService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsApiServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.uacf.friends.FriendsApiService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsApiServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsApiService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsApiServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.uacf.friends.FriendsApiService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsApiServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FriendsApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.uacf.friends.FriendsApiService",
	HandlerType: (*FriendsApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _FriendsApiService_Ping_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _FriendsApiService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FriendsApiService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("friends.uacf.io/apps/friendsapi/rpc/friendsapi.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0x7f, 0x59, 0xa6, 0x4d, 0x11, 0x46, 0xac, 0xa2, 0x6e, 0x41, 0x95, 0xb9, 0xf4, 0x94,
	0x68, 0x0b, 0x37, 0xc4, 0x01, 0x90, 0x58, 0x81, 0xb6, 0x62, 0x95, 0x15, 0x17, 0x2e, 0xc5, 0x9b,
	0x4c, 0x5b, 0x4b, 0xbb, 0xb1, 0xb1, 0x9d, 0x95, 0x10, 0xe2, 0xc2, 0x2b, 0x70, 0xe2, 0x21, 0x78,
	0x1a, 0x5e, 0x81, 0x07, 0x41, 0x71, 0x9c, 0xc6, 0xd0, 0x4a, 0x1c, 0x96, 0x5b, 0x66, 0x3e, 0xcf,
	0xf7, 0x7d, 0xf3, 0x39, 0x32, 0x3c, 0x59, 0x29, 0x8e, 0x79, 0xa6, 0xa3, 0x82, 0xa5, 0xab, 0x88,
	0x8b, 0x98, 0x49, 0xa9, 0x63, 0xd7, 0x64, 0x92, 0xc7, 0x4a, 0xa6, 0x5e, 0x19, 0x49, 0x25, 0x8c,
	0x20, 0x77, 0xb8, 0xa8, 0x06, 0x1c, 0x32, 0x9e, 0x99, 0x0d, 0x57, 0xd9, 0x52, 0x32, 0x65, 0x3e,
	0xc5, 0x6b, 0x21, 0xd6, 0x97, 0x18, 0x97, 0xd3, 0x2c, 0xcf, 0x85, 0x61, 0x86, 0x8b, 0x5c, 0x57,
	0xa3, 0xf4, 0x7b, 0x0b, 0xe0, 0x55, 0x35, 0xb5, 0xe1, 0x92, 0x8c, 0xa0, 0xcd, 0xb3, 0xb0, 0x35,
	0x6d, 0xcd, 0x6e, 0x27, 0x6d, 0x9e, 0x91, 0x29, 0x0c, 0x57, 0x4a, 0x5c, 0x2d, 0x0b, 0x8d, 0x6a,
	0xc9, 0xb3, 0xb0, 0x6d, 0x11, 0x28, 0x7b, 0xef, 0x34, 0xaa, 0xd7, 0x19, 0x99, 0x00, 0x18, 0xb1,
	0xc5, 0x3b, 0x16, 0x3f, 0x30, 0xc2, 0xa1, 0x8f, 0x20, 0x70, 0x9e, 0x96, 0x9a, 0xe7, 0x29, 0x86,
	0x5d, 0x7b, 0x60, 0xe8, 0x9a, 0xe7, 0x65, 0x8f, 0x1c, 0x42, 0x5f, 0x1b, 0x66, 0x0a, 0x1d, 0xf6,
	0x2c, 0xea, 0x2a, 0x7a, 0x0a, 0xc1, 0x4b, 0x85, 0xcc, 0x60, 0x82, 0x1f, 0x0b, 0xd4, 0x86, 0x3c,
	0x05, 0x58, 0x6d, 0xbd, 0x5a, 0x97, 0x83, 0xf9, 0x51, 0xf4, 0xd7, 0xf2, 0x51, 0xb3, 0x4e, 0xe2,
	0x1d, 0xa7, 0x0b, 0x18, 0xd5, 0x6c, 0x5a, 0x8a, 0x5c, 0xe3, 0xcd, 0xe8, 0x26, 0x00, 0x27, 0x68,
	0x6a, 0x67, 0x4d, 0x6e, 0x9d, 0x32, 0x37, 0xfa, 0x06, 0x06, 0x16, 0xfd, 0x1f, 0x4a, 0x08, 0x83,
	0x53, 0xae, 0xb7, 0x52, 0x4d, 0x5a, 0x2d, 0x3f, 0xad, 0x9b, 0x5e, 0x15, 0x5d, 0xc0, 0xb0, 0x92,
	0x71, 0x9e, 0x9f, 0xc1, 0xa0, 0x31, 0x51, 0x8a, 0x75, 0xfe, 0x65, 0xda, 0x3f, 0x4f, 0xdf, 0x42,
	0xb0, 0x40, 0xb3, 0x11, 0x59, 0xed, 0xfb, 0x21, 0x80, 0x64, 0x8a, 0x5d, 0xa1, 0x41, 0x75, 0xec,
	0xbc, 0x7b, 0x9d, 0x3f, 0xf0, 0x79, 0xed, 0xbe, 0xe9, 0xd0, 0x19, 0x8c, 0x6a, 0x42, 0xe7, 0xf0,
	0x10, 0xfa, 0x0a, 0x75, 0x71, 0x69, 0xea, 0x24, 0xaa, 0x8a, 0x06, 0x30, 0x38, 0xe3, 0xf9, 0xda,
	0x09, 0x53, 0x0a, 0xc3, 0x33, 0x51, 0x96, 0x6e, 0x8c, 0x40, 0x57, 0x8a, 0x7c, 0x6d, 0x87, 0x0e,
	0x12, 0xfb, 0x3d, 0xff, 0xd1, 0x86, 0xbb, 0x6e, 0x93, 0xe7, 0x92, 0x9f, 0xa3, 0xba, 0xe6, 0x29,
	0x92, 0x04, 0xba, 0x25, 0x11, 0x99, 0xec, 0x6c, 0xed, 0xf1, 0x8f, 0x1f, 0xec, 0xa2, 0x9e, 0x1c,
	0x0d, 0xbe, 0xfe, 0xfc, 0xf5, 0xad, 0x7d, 0x8b, 0xf4, 0x62, 0x59, 0x72, 0x5d, 0x40, 0xe7, 0x04,
	0x0d, 0xd9, 0x0d, 0xb2, 0xf9, 0x9b, 0xc6, 0x93, 0xfd, 0xa0, 0x23, 0x9c, 0x5a, 0xc2, 0x31, 0x09,
	0xe3, 0xeb, 0xe3, 0x58, 0xa1, 0x36, 0xfe, 0xfb, 0xf0, 0x99, 0x67, 0x5f, 0xc8, 0x07, 0xe8, 0x96,
	0x57, 0xb9, 0xc7, 0xb7, 0xf7, 0x23, 0xed, 0xf1, 0xed, 0xdf, 0x3f, 0x3d, 0xb2, 0x32, 0xf7, 0xc9,
	0xbd, 0x3d, 0x32, 0x2f, 0x7a, 0xef, 0x3b, 0x4a, 0xa6, 0x17, 0x7d, 0xfb, 0x88, 0x3c, 0xfe, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x6f, 0x0c, 0x21, 0xbb, 0xb7, 0x04, 0x00, 0x00,
}
