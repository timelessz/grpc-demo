// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: book.proto

package v8

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

// 请求书籍 详情   根据id等查询
type BookQueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author string  `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Price  float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Page   int32   `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32   `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *BookQueryParams) Reset() {
	*x = BookQueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookQueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookQueryParams) ProtoMessage() {}

func (x *BookQueryParams) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BookQueryParams.ProtoReflect.Descriptor instead.
func (*BookQueryParams) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *BookQueryParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BookQueryParams) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *BookQueryParams) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BookQueryParams) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *BookQueryParams) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type OneBookQueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OneBookQueryParams) Reset() {
	*x = OneBookQueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneBookQueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneBookQueryParams) ProtoMessage() {}

func (x *OneBookQueryParams) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OneBookQueryParams.ProtoReflect.Descriptor instead.
func (*OneBookQueryParams) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
}

func (x *OneBookQueryParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author      string  `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Price       float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Description string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
}

func (x *Book) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Book) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type BookList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	Total int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
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

func (x *BookList) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

func (x *BookList) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0f, 0x42, 0x6f,
	0x6f, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x4f, 0x6e, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x7a, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x08, 0x42,
	0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xce, 0x01, 0x0a, 0x0b, 0x42,
	0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x0d,
	0x2f, 0x76, 0x33, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x13,
	0x2e, 0x4f, 0x6e, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0c, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x6f, 0x6e, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x36, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x3b, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*BookQueryParams)(nil),    // 0: BookQueryParams
	(*OneBookQueryParams)(nil), // 1: OneBookQueryParams
	(*Book)(nil),               // 2: Book
	(*BookList)(nil),           // 3: BookList
}
var file_book_proto_depIdxs = []int32{
	2, // 0: BookList.books:type_name -> Book
	0, // 1: BookService.GetBookList:input_type -> BookQueryParams
	1, // 2: BookService.GetOneBook:input_type -> OneBookQueryParams
	2, // 3: BookService.CreateBook:input_type -> Book
	3, // 4: BookService.GetBookList:output_type -> BookList
	2, // 5: BookService.GetOneBook:output_type -> Book
	2, // 6: BookService.CreateBook:output_type -> Book
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
			switch v := v.(*BookQueryParams); i {
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
			switch v := v.(*OneBookQueryParams); i {
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
			switch v := v.(*Book); i {
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
	GetBookList(ctx context.Context, in *BookQueryParams, opts ...grpc.CallOption) (*BookList, error)
	GetOneBook(ctx context.Context, in *OneBookQueryParams, opts ...grpc.CallOption) (*Book, error)
	CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBookList(ctx context.Context, in *BookQueryParams, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/BookService/GetBookList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetOneBook(ctx context.Context, in *OneBookQueryParams, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/BookService/GetOneBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/BookService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	GetBookList(context.Context, *BookQueryParams) (*BookList, error)
	GetOneBook(context.Context, *OneBookQueryParams) (*Book, error)
	CreateBook(context.Context, *Book) (*Book, error)
}

// UnimplementedBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (*UnimplementedBookServiceServer) GetBookList(context.Context, *BookQueryParams) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookList not implemented")
}
func (*UnimplementedBookServiceServer) GetOneBook(context.Context, *OneBookQueryParams) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneBook not implemented")
}
func (*UnimplementedBookServiceServer) CreateBook(context.Context, *Book) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/GetBookList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookList(ctx, req.(*BookQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetOneBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OneBookQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetOneBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/GetOneBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetOneBook(ctx, req.(*OneBookQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookList",
			Handler:    _BookService_GetBookList_Handler,
		},
		{
			MethodName: "GetOneBook",
			Handler:    _BookService_GetOneBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}