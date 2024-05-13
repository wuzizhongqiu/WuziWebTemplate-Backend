// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: http/post/post_http.proto

package post

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Tags   string `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	UserId int64  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_post_post_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_post_post_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_http_post_post_http_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreatePostRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *CreatePostRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreatePostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePostReply) Reset() {
	*x = CreatePostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_post_post_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostReply) ProtoMessage() {}

func (x *CreatePostReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_post_post_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostReply.ProtoReflect.Descriptor instead.
func (*CreatePostReply) Descriptor() ([]byte, []int) {
	return file_http_post_post_http_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_http_post_post_http_proto protoreflect.FileDescriptor

var file_http_post_post_http_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x1b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x6e, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x66, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1c,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x42, 0x1e, 0x5a, 0x1c, 0x77, 0x75, 0x7a, 0x69,
	0x67, 0x6f, 0x77, 0x65, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_post_post_http_proto_rawDescOnce sync.Once
	file_http_post_post_http_proto_rawDescData = file_http_post_post_http_proto_rawDesc
)

func file_http_post_post_http_proto_rawDescGZIP() []byte {
	file_http_post_post_http_proto_rawDescOnce.Do(func() {
		file_http_post_post_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_post_post_http_proto_rawDescData)
	})
	return file_http_post_post_http_proto_rawDescData
}

var file_http_post_post_http_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_http_post_post_http_proto_goTypes = []interface{}{
	(*CreatePostRequest)(nil), // 0: http.post.CreatePostRequest
	(*CreatePostReply)(nil),   // 1: http.post.CreatePostReply
}
var file_http_post_post_http_proto_depIdxs = []int32{
	0, // 0: http.post.Post.CreatePost:input_type -> http.post.CreatePostRequest
	1, // 1: http.post.Post.CreatePost:output_type -> http.post.CreatePostReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_http_post_post_http_proto_init() }
func file_http_post_post_http_proto_init() {
	if File_http_post_post_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_post_post_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_http_post_post_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostReply); i {
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
			RawDescriptor: file_http_post_post_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_http_post_post_http_proto_goTypes,
		DependencyIndexes: file_http_post_post_http_proto_depIdxs,
		MessageInfos:      file_http_post_post_http_proto_msgTypes,
	}.Build()
	File_http_post_post_http_proto = out.File
	file_http_post_post_http_proto_rawDesc = nil
	file_http_post_post_http_proto_goTypes = nil
	file_http_post_post_http_proto_depIdxs = nil
}