// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: MaxEvenNumber.proto

package proto

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

type EvenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EvenNumber int32 `protobuf:"varint,1,opt,name=evenNumber,proto3" json:"evenNumber,omitempty"`
}

func (x *EvenRequest) Reset() {
	*x = EvenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MaxEvenNumber_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvenRequest) ProtoMessage() {}

func (x *EvenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MaxEvenNumber_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvenRequest.ProtoReflect.Descriptor instead.
func (*EvenRequest) Descriptor() ([]byte, []int) {
	return file_MaxEvenNumber_proto_rawDescGZIP(), []int{0}
}

func (x *EvenRequest) GetEvenNumber() int32 {
	if x != nil {
		return x.EvenNumber
	}
	return 0
}

type EvenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *EvenResponse) Reset() {
	*x = EvenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MaxEvenNumber_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvenResponse) ProtoMessage() {}

func (x *EvenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MaxEvenNumber_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvenResponse.ProtoReflect.Descriptor instead.
func (*EvenResponse) Descriptor() ([]byte, []int) {
	return file_MaxEvenNumber_proto_rawDescGZIP(), []int{1}
}

func (x *EvenResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_MaxEvenNumber_proto protoreflect.FileDescriptor

var file_MaxEvenNumber_proto_rawDesc = []byte{
	0x0a, 0x13, 0x4d, 0x61, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x22, 0x2d, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x26, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6c, 0x6c, 0x61, 0x72, 0x61, 0x73, 0x75,
	0x53, 0x4a, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MaxEvenNumber_proto_rawDescOnce sync.Once
	file_MaxEvenNumber_proto_rawDescData = file_MaxEvenNumber_proto_rawDesc
)

func file_MaxEvenNumber_proto_rawDescGZIP() []byte {
	file_MaxEvenNumber_proto_rawDescOnce.Do(func() {
		file_MaxEvenNumber_proto_rawDescData = protoimpl.X.CompressGZIP(file_MaxEvenNumber_proto_rawDescData)
	})
	return file_MaxEvenNumber_proto_rawDescData
}

var file_MaxEvenNumber_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_MaxEvenNumber_proto_goTypes = []interface{}{
	(*EvenRequest)(nil),  // 0: calculator.EvenRequest
	(*EvenResponse)(nil), // 1: calculator.EvenResponse
}
var file_MaxEvenNumber_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MaxEvenNumber_proto_init() }
func file_MaxEvenNumber_proto_init() {
	if File_MaxEvenNumber_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MaxEvenNumber_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvenRequest); i {
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
		file_MaxEvenNumber_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvenResponse); i {
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
			RawDescriptor: file_MaxEvenNumber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MaxEvenNumber_proto_goTypes,
		DependencyIndexes: file_MaxEvenNumber_proto_depIdxs,
		MessageInfos:      file_MaxEvenNumber_proto_msgTypes,
	}.Build()
	File_MaxEvenNumber_proto = out.File
	file_MaxEvenNumber_proto_rawDesc = nil
	file_MaxEvenNumber_proto_goTypes = nil
	file_MaxEvenNumber_proto_depIdxs = nil
}
