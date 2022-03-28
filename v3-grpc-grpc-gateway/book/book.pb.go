// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: book.proto

package v3

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求书详情的参数结构  book_id 32位整形
type BookInfoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId int32 `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *BookInfoParams) Reset() {
	*x = BookInfoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookInfoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookInfoParams) ProtoMessage() {}

func (x *BookInfoParams) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookInfoParams.ProtoReflect.Descriptor instead.
func (*BookInfoParams) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *BookInfoParams) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

// 书详情信息的结构   book_name字符串类型
type BookInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId   int32  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BookName string `protobuf:"bytes,2,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
}

func (x *BookInfo) Reset() {
	*x = BookInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookInfo) ProtoMessage() {}

func (x *BookInfo) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookInfo.ProtoReflect.Descriptor instead.
func (*BookInfo) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
}

func (x *BookInfo) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *BookInfo) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

// 请求书列表的参数结构  page、limit   32位整形
type BookListParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *BookListParams) Reset() {
	*x = BookListParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookListParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookListParams) ProtoMessage() {}

func (x *BookListParams) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookListParams.ProtoReflect.Descriptor instead.
func (*BookListParams) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
}

func (x *BookListParams) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *BookListParams) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// 书列表的结构    BookInfo结构数组
type BookList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookList []*BookInfo `protobuf:"bytes,1,rep,name=book_list,json=bookList,proto3" json:"book_list,omitempty"`
}

func (x *BookList) Reset() {
	*x = BookList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookList) ProtoMessage() {}

func (x *BookList) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookList.ProtoReflect.Descriptor instead.
func (*BookList) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{3}
}

func (x *BookList) GetBookList() []*BookInfo {
	if x != nil {
		return x.BookList
	}
	return nil
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x33,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29,
	0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x08, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x0e, 0x42,
	0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x35, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xf7,
	0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e,
	0x76, 0x33, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x76, 0x33,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x4c, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x76, 0x33,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x0c, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x76, 0x33, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0c, 0x2e,
	0x76, 0x33, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x3b, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_proto_rawDescOnce sync.Once
	file_book_proto_rawDescData = file_book_proto_rawDesc
)

func file_book_proto_rawDescGZIP() []byte {
	file_book_proto_rawDescOnce.Do(func() {
		file_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_proto_rawDescData)
	})
	return file_book_proto_rawDescData
}

var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_book_proto_goTypes = []interface{}{
	(*BookInfoParams)(nil), // 0: v3.BookInfoParams
	(*BookInfo)(nil),       // 1: v3.BookInfo
	(*BookListParams)(nil), // 2: v3.BookListParams
	(*BookList)(nil),       // 3: v3.BookList
}
var file_book_proto_depIdxs = []int32{
	1, // 0: v3.BookList.book_list:type_name -> v3.BookInfo
	1, // 1: v3.BookService.AddBookInfo:input_type -> v3.BookInfo
	0, // 2: v3.BookService.GetBookInfo:input_type -> v3.BookInfoParams
	2, // 3: v3.BookService.GetBookList:input_type -> v3.BookListParams
	0, // 4: v3.BookService.AddBookInfo:output_type -> v3.BookInfoParams
	1, // 5: v3.BookService.GetBookInfo:output_type -> v3.BookInfo
	3, // 6: v3.BookService.GetBookList:output_type -> v3.BookList
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_book_proto_init() }
func file_book_proto_init() {
	if File_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookInfoParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookListParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_proto_goTypes,
		DependencyIndexes: file_book_proto_depIdxs,
		MessageInfos:      file_book_proto_msgTypes,
	}.Build()
	File_book_proto = out.File
	file_book_proto_rawDesc = nil
	file_book_proto_goTypes = nil
	file_book_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServiceClient interface {
	// 添加书籍信息
	AddBookInfo(ctx context.Context, in *BookInfo, opts ...grpc.CallOption) (*BookInfoParams, error)
	// 获取书籍信息
	GetBookInfo(ctx context.Context, in *BookInfoParams, opts ...grpc.CallOption) (*BookInfo, error)
	// 获取图书列表
	GetBookList(ctx context.Context, in *BookListParams, opts ...grpc.CallOption) (*BookList, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) AddBookInfo(ctx context.Context, in *BookInfo, opts ...grpc.CallOption) (*BookInfoParams, error) {
	out := new(BookInfoParams)
	err := c.cc.Invoke(ctx, "/v3.BookService/AddBookInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBookInfo(ctx context.Context, in *BookInfoParams, opts ...grpc.CallOption) (*BookInfo, error) {
	out := new(BookInfo)
	err := c.cc.Invoke(ctx, "/v3.BookService/GetBookInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBookList(ctx context.Context, in *BookListParams, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/v3.BookService/GetBookList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	// 添加书籍信息
	AddBookInfo(context.Context, *BookInfo) (*BookInfoParams, error)
	// 获取书籍信息
	GetBookInfo(context.Context, *BookInfoParams) (*BookInfo, error)
	// 获取图书列表
	GetBookList(context.Context, *BookListParams) (*BookList, error)
}

// UnimplementedBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (*UnimplementedBookServiceServer) AddBookInfo(context.Context, *BookInfo) (*BookInfoParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBookInfo not implemented")
}
func (*UnimplementedBookServiceServer) GetBookInfo(context.Context, *BookInfoParams) (*BookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfo not implemented")
}
func (*UnimplementedBookServiceServer) GetBookList(context.Context, *BookListParams) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookList not implemented")
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_AddBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).AddBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.BookService/AddBookInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).AddBookInfo(ctx, req.(*BookInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.BookService/GetBookInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookInfo(ctx, req.(*BookInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.BookService/GetBookList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookList(ctx, req.(*BookListParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v3.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBookInfo",
			Handler:    _BookService_AddBookInfo_Handler,
		},
		{
			MethodName: "GetBookInfo",
			Handler:    _BookService_GetBookInfo_Handler,
		},
		{
			MethodName: "GetBookList",
			Handler:    _BookService_GetBookList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}
