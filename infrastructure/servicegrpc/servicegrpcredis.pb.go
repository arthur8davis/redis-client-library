// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: servicegrpcredis.proto

package servicegrpc

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

type RequestGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RequestGet) Reset() {
	*x = RequestGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicegrpcredis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGet) ProtoMessage() {}

func (x *RequestGet) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpcredis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGet.ProtoReflect.Descriptor instead.
func (*RequestGet) Descriptor() ([]byte, []int) {
	return file_servicegrpcredis_proto_rawDescGZIP(), []int{0}
}

func (x *RequestGet) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ResponseGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value              string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Message            string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	IsCacheKeyNotFound bool   `protobuf:"varint,3,opt,name=isCacheKeyNotFound,proto3" json:"isCacheKeyNotFound,omitempty"`
}

func (x *ResponseGet) Reset() {
	*x = ResponseGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicegrpcredis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGet) ProtoMessage() {}

func (x *ResponseGet) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpcredis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGet.ProtoReflect.Descriptor instead.
func (*ResponseGet) Descriptor() ([]byte, []int) {
	return file_servicegrpcredis_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseGet) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ResponseGet) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseGet) GetIsCacheKeyNotFound() bool {
	if x != nil {
		return x.IsCacheKeyNotFound
	}
	return false
}

type RequestSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key                 string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value               string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ExpirationInSeconds int64  `protobuf:"varint,3,opt,name=expirationInSeconds,proto3" json:"expirationInSeconds,omitempty"`
}

func (x *RequestSet) Reset() {
	*x = RequestSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicegrpcredis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSet) ProtoMessage() {}

func (x *RequestSet) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpcredis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSet.ProtoReflect.Descriptor instead.
func (*RequestSet) Descriptor() ([]byte, []int) {
	return file_servicegrpcredis_proto_rawDescGZIP(), []int{2}
}

func (x *RequestSet) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RequestSet) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RequestSet) GetExpirationInSeconds() int64 {
	if x != nil {
		return x.ExpirationInSeconds
	}
	return 0
}

type ResponseSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseSet) Reset() {
	*x = ResponseSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicegrpcredis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseSet) ProtoMessage() {}

func (x *ResponseSet) ProtoReflect() protoreflect.Message {
	mi := &file_servicegrpcredis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseSet.ProtoReflect.Descriptor instead.
func (*ResponseSet) Descriptor() ([]byte, []int) {
	return file_servicegrpcredis_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseSet) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_servicegrpcredis_proto protoreflect.FileDescriptor

var file_servicegrpcredis_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x67, 0x72, 0x70, 0x63, 0x22, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x47, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x6d, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x47, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b,
	0x65, 0x79, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x69, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x4e, 0x6f, 0x74, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x22, 0x66, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x27, 0x0a, 0x0b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x86, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x64, 0x69, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x65, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x74, 0x22, 0x00, 0x42, 0x27,
	0x5a, 0x25, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_servicegrpcredis_proto_rawDescOnce sync.Once
	file_servicegrpcredis_proto_rawDescData = file_servicegrpcredis_proto_rawDesc
)

func file_servicegrpcredis_proto_rawDescGZIP() []byte {
	file_servicegrpcredis_proto_rawDescOnce.Do(func() {
		file_servicegrpcredis_proto_rawDescData = protoimpl.X.CompressGZIP(file_servicegrpcredis_proto_rawDescData)
	})
	return file_servicegrpcredis_proto_rawDescData
}

var file_servicegrpcredis_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_servicegrpcredis_proto_goTypes = []interface{}{
	(*RequestGet)(nil),  // 0: servicegrpc.RequestGet
	(*ResponseGet)(nil), // 1: servicegrpc.ResponseGet
	(*RequestSet)(nil),  // 2: servicegrpc.RequestSet
	(*ResponseSet)(nil), // 3: servicegrpc.ResponseSet
}
var file_servicegrpcredis_proto_depIdxs = []int32{
	0, // 0: servicegrpc.RedisService.Get:input_type -> servicegrpc.RequestGet
	2, // 1: servicegrpc.RedisService.Set:input_type -> servicegrpc.RequestSet
	1, // 2: servicegrpc.RedisService.Get:output_type -> servicegrpc.ResponseGet
	3, // 3: servicegrpc.RedisService.Set:output_type -> servicegrpc.ResponseSet
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_servicegrpcredis_proto_init() }
func file_servicegrpcredis_proto_init() {
	if File_servicegrpcredis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_servicegrpcredis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGet); i {
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
		file_servicegrpcredis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGet); i {
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
		file_servicegrpcredis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestSet); i {
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
		file_servicegrpcredis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseSet); i {
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
			RawDescriptor: file_servicegrpcredis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_servicegrpcredis_proto_goTypes,
		DependencyIndexes: file_servicegrpcredis_proto_depIdxs,
		MessageInfos:      file_servicegrpcredis_proto_msgTypes,
	}.Build()
	File_servicegrpcredis_proto = out.File
	file_servicegrpcredis_proto_rawDesc = nil
	file_servicegrpcredis_proto_goTypes = nil
	file_servicegrpcredis_proto_depIdxs = nil
}
