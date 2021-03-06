// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache-service.proto

package cacheService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Item is what is stored in the cache
type Item struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expiration           string   `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Item) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Item) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

type GetKey struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKey) Reset()         { *m = GetKey{} }
func (m *GetKey) String() string { return proto.CompactTextString(m) }
func (*GetKey) ProtoMessage()    {}
func (*GetKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{1}
}

func (m *GetKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKey.Unmarshal(m, b)
}
func (m *GetKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKey.Marshal(b, m, deterministic)
}
func (m *GetKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKey.Merge(m, src)
}
func (m *GetKey) XXX_Size() int {
	return xxx_messageInfo_GetKey.Size(m)
}
func (m *GetKey) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKey.DiscardUnknown(m)
}

var xxx_messageInfo_GetKey proto.InternalMessageInfo

func (m *GetKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type AllItems struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllItems) Reset()         { *m = AllItems{} }
func (m *AllItems) String() string { return proto.CompactTextString(m) }
func (*AllItems) ProtoMessage()    {}
func (*AllItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{2}
}

func (m *AllItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllItems.Unmarshal(m, b)
}
func (m *AllItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllItems.Marshal(b, m, deterministic)
}
func (m *AllItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllItems.Merge(m, src)
}
func (m *AllItems) XXX_Size() int {
	return xxx_messageInfo_AllItems.Size(m)
}
func (m *AllItems) XXX_DiscardUnknown() {
	xxx_messageInfo_AllItems.DiscardUnknown(m)
}

var xxx_messageInfo_AllItems proto.InternalMessageInfo

func (m *AllItems) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type Success struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Success) Reset()         { *m = Success{} }
func (m *Success) String() string { return proto.CompactTextString(m) }
func (*Success) ProtoMessage()    {}
func (*Success) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{3}
}

func (m *Success) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Success.Unmarshal(m, b)
}
func (m *Success) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Success.Marshal(b, m, deterministic)
}
func (m *Success) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Success.Merge(m, src)
}
func (m *Success) XXX_Size() int {
	return xxx_messageInfo_Success.Size(m)
}
func (m *Success) XXX_DiscardUnknown() {
	xxx_messageInfo_Success.DiscardUnknown(m)
}

var xxx_messageInfo_Success proto.InternalMessageInfo

func (m *Success) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Item)(nil), "cacheService.Item")
	proto.RegisterType((*GetKey)(nil), "cacheService.GetKey")
	proto.RegisterType((*AllItems)(nil), "cacheService.AllItems")
	proto.RegisterType((*Success)(nil), "cacheService.Success")
}

func init() { proto.RegisterFile("cache-service.proto", fileDescriptor_9b0ee6ef8b54c4e4) }

var fileDescriptor_9b0ee6ef8b54c4e4 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x6a, 0x83, 0x40,
	0x10, 0xc7, 0x31, 0x36, 0x5f, 0x93, 0x1c, 0xca, 0x36, 0x0d, 0x62, 0xa1, 0x04, 0x7b, 0xf1, 0x92,
	0x15, 0xd2, 0x1e, 0x4a, 0x4b, 0x0f, 0xf6, 0x03, 0x29, 0x85, 0x52, 0xcc, 0x13, 0x18, 0x33, 0x49,
	0xa5, 0x9b, 0x2a, 0xba, 0x86, 0xf8, 0xb6, 0x7d, 0x94, 0xb2, 0xbb, 0x18, 0x94, 0xe8, 0x6d, 0xe6,
	0x37, 0xf3, 0xff, 0xcf, 0xee, 0x0c, 0x5c, 0x84, 0x41, 0xf8, 0x8d, 0xf3, 0x0c, 0xd3, 0x7d, 0x14,
	0x22, 0x4d, 0xd2, 0x98, 0xc7, 0x64, 0x2c, 0xe1, 0x52, 0x31, 0xf3, 0x6a, 0x1b, 0xc7, 0x5b, 0x86,
	0x8e, 0xac, 0xad, 0xf2, 0x8d, 0x83, 0xbb, 0x84, 0x17, 0xaa, 0xd5, 0xfa, 0x84, 0xb3, 0x77, 0x8e,
	0x3b, 0x72, 0x0e, 0xfa, 0x0f, 0x16, 0x86, 0x36, 0xd3, 0xec, 0xa1, 0x2f, 0x42, 0x32, 0x81, 0xee,
	0x3e, 0x60, 0x39, 0x1a, 0x1d, 0xc9, 0x54, 0x42, 0xae, 0x01, 0xf0, 0x90, 0x44, 0x69, 0xc0, 0xa3,
	0xf8, 0xd7, 0xd0, 0x65, 0xa9, 0x42, 0x2c, 0x13, 0x7a, 0x1e, 0xf2, 0x0f, 0x2c, 0x4e, 0x1d, 0xad,
	0x3b, 0x18, 0xb8, 0x8c, 0x89, 0x71, 0x19, 0xb1, 0xa1, 0x1b, 0x89, 0xc0, 0xd0, 0x66, 0xba, 0x3d,
	0x5a, 0x10, 0x5a, 0x7d, 0x32, 0x15, 0x3d, 0xbe, 0x6a, 0xb0, 0x6e, 0xa0, 0xbf, 0xcc, 0xc3, 0x10,
	0xb3, 0x8c, 0x18, 0xd0, 0xcf, 0x54, 0x28, 0x6d, 0x07, 0x7e, 0x99, 0x2e, 0xfe, 0x3a, 0x30, 0x7e,
	0xa9, 0x38, 0x90, 0x39, 0xe8, 0xee, 0x7a, 0x4d, 0x1a, 0x7c, 0xcd, 0x06, 0x46, 0x1c, 0xd0, 0x3d,
	0xe4, 0x64, 0x52, 0x2f, 0xa9, 0x9f, 0x34, 0x0a, 0x1e, 0x61, 0xe4, 0x21, 0x7f, 0x2e, 0xbe, 0x52,
	0xdc, 0x44, 0x87, 0x16, 0xe1, 0xb4, 0x4e, 0x8f, 0x9f, 0x7f, 0x92, 0xe2, 0x63, 0x3a, 0xa5, 0xea,
	0x42, 0xb4, 0xbc, 0x10, 0x7d, 0x13, 0x17, 0x6a, 0x95, 0xdf, 0xc3, 0xf0, 0x15, 0x19, 0x72, 0x14,
	0x6b, 0x6e, 0x9e, 0x7c, 0x59, 0xa7, 0xe5, 0x02, 0x1f, 0x4a, 0xa5, 0xcb, 0x58, 0xeb, 0xd8, 0x66,
	0xed, 0xaa, 0x27, 0xdb, 0x6e, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x18, 0xca, 0xb9, 0x72,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheServiceClient interface {
	Add(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	Get(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*Item, error)
	GetByPrefix(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*AllItems, error)
	GetAllItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllItems, error)
	DeleteKey(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*Success, error)
	DeleteAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Success, error)
}

type cacheServiceClient struct {
	cc *grpc.ClientConn
}

func NewCacheServiceClient(cc *grpc.ClientConn) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Add(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Get(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetByPrefix(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*AllItems, error) {
	out := new(AllItems)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/GetByPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetAllItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllItems, error) {
	out := new(AllItems)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/GetAllItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) DeleteKey(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/DeleteKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) DeleteAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/DeleteAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
type CacheServiceServer interface {
	Add(context.Context, *Item) (*Item, error)
	Get(context.Context, *GetKey) (*Item, error)
	GetByPrefix(context.Context, *GetKey) (*AllItems, error)
	GetAllItems(context.Context, *empty.Empty) (*AllItems, error)
	DeleteKey(context.Context, *GetKey) (*Success, error)
	DeleteAll(context.Context, *empty.Empty) (*Success, error)
}

// UnimplementedCacheServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (*UnimplementedCacheServiceServer) Add(ctx context.Context, req *Item) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCacheServiceServer) Get(ctx context.Context, req *GetKey) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCacheServiceServer) GetByPrefix(ctx context.Context, req *GetKey) (*AllItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPrefix not implemented")
}
func (*UnimplementedCacheServiceServer) GetAllItems(ctx context.Context, req *empty.Empty) (*AllItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllItems not implemented")
}
func (*UnimplementedCacheServiceServer) DeleteKey(ctx context.Context, req *GetKey) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
func (*UnimplementedCacheServiceServer) DeleteAll(ctx context.Context, req *empty.Empty) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAll not implemented")
}

func RegisterCacheServiceServer(s *grpc.Server, srv CacheServiceServer) {
	s.RegisterService(&_CacheService_serviceDesc, srv)
}

func _CacheService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Add(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Get(ctx, req.(*GetKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetByPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetByPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/GetByPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetByPrefix(ctx, req.(*GetKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetAllItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetAllItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/GetAllItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetAllItems(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/DeleteKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).DeleteKey(ctx, req.(*GetKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_DeleteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).DeleteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/DeleteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).DeleteAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cacheService.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CacheService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CacheService_Get_Handler,
		},
		{
			MethodName: "GetByPrefix",
			Handler:    _CacheService_GetByPrefix_Handler,
		},
		{
			MethodName: "GetAllItems",
			Handler:    _CacheService_GetAllItems_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _CacheService_DeleteKey_Handler,
		},
		{
			MethodName: "DeleteAll",
			Handler:    _CacheService_DeleteAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache-service.proto",
}
