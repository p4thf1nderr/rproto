// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.13.0
// source: api/pdfcompose.proto

package grpc

import (
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

type SendIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Order int64  `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SendIn) Reset() {
	*x = SendIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pdfcompose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendIn) ProtoMessage() {}

func (x *SendIn) ProtoReflect() protoreflect.Message {
	mi := &file_api_pdfcompose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendIn.ProtoReflect.Descriptor instead.
func (*SendIn) Descriptor() ([]byte, []int) {
	return file_api_pdfcompose_proto_rawDescGZIP(), []int{0}
}

func (x *SendIn) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *SendIn) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *SendIn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SendOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *SendOut) Reset() {
	*x = SendOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pdfcompose_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOut) ProtoMessage() {}

func (x *SendOut) ProtoReflect() protoreflect.Message {
	mi := &file_api_pdfcompose_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOut.ProtoReflect.Descriptor instead.
func (*SendOut) Descriptor() ([]byte, []int) {
	return file_api_pdfcompose_proto_rawDescGZIP(), []int{1}
}

func (x *SendOut) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_api_pdfcompose_proto protoreflect.FileDescriptor

var file_api_pdfcompose_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x65, 0x22, 0x48, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x07,
	0x53, 0x65, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0x4a, 0x0a,
	0x11, 0x50, 0x44, 0x46, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x64, 0x66,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x1a, 0x13,
	0x2e, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x4f, 0x75, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x64, 0x66,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_pdfcompose_proto_rawDescOnce sync.Once
	file_api_pdfcompose_proto_rawDescData = file_api_pdfcompose_proto_rawDesc
)

func file_api_pdfcompose_proto_rawDescGZIP() []byte {
	file_api_pdfcompose_proto_rawDescOnce.Do(func() {
		file_api_pdfcompose_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_pdfcompose_proto_rawDescData)
	})
	return file_api_pdfcompose_proto_rawDescData
}

var file_api_pdfcompose_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_pdfcompose_proto_goTypes = []interface{}{
	(*SendIn)(nil),  // 0: pdfcompose.SendIn
	(*SendOut)(nil), // 1: pdfcompose.SendOut
}
var file_api_pdfcompose_proto_depIdxs = []int32{
	0, // 0: pdfcompose.PDFComposeService.Send:input_type -> pdfcompose.SendIn
	1, // 1: pdfcompose.PDFComposeService.Send:output_type -> pdfcompose.SendOut
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_pdfcompose_proto_init() }
func file_api_pdfcompose_proto_init() {
	if File_api_pdfcompose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_pdfcompose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendIn); i {
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
		file_api_pdfcompose_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOut); i {
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
			RawDescriptor: file_api_pdfcompose_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_pdfcompose_proto_goTypes,
		DependencyIndexes: file_api_pdfcompose_proto_depIdxs,
		MessageInfos:      file_api_pdfcompose_proto_msgTypes,
	}.Build()
	File_api_pdfcompose_proto = out.File
	file_api_pdfcompose_proto_rawDesc = nil
	file_api_pdfcompose_proto_goTypes = nil
	file_api_pdfcompose_proto_depIdxs = nil
}
