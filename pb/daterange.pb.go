// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: agokit/pb/daterange.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Date From
	From *timestamp.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	//Date To
	To *timestamp.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	//Default field is created_at
	Field string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agokit_pb_daterange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_agokit_pb_daterange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRange.ProtoReflect.Descriptor instead.
func (*DateRange) Descriptor() ([]byte, []int) {
	return file_agokit_pb_daterange_proto_rawDescGZIP(), []int{0}
}

func (x *DateRange) GetFrom() *timestamp.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *DateRange) GetTo() *timestamp.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *DateRange) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

var File_agokit_pb_daterange_proto protoreflect.FileDescriptor

var file_agokit_pb_daterange_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x67, 0x6f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7d, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x0b, 0x5a, 0x09, 0x61, 0x67, 0x6f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agokit_pb_daterange_proto_rawDescOnce sync.Once
	file_agokit_pb_daterange_proto_rawDescData = file_agokit_pb_daterange_proto_rawDesc
)

func file_agokit_pb_daterange_proto_rawDescGZIP() []byte {
	file_agokit_pb_daterange_proto_rawDescOnce.Do(func() {
		file_agokit_pb_daterange_proto_rawDescData = protoimpl.X.CompressGZIP(file_agokit_pb_daterange_proto_rawDescData)
	})
	return file_agokit_pb_daterange_proto_rawDescData
}

var file_agokit_pb_daterange_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_agokit_pb_daterange_proto_goTypes = []interface{}{
	(*DateRange)(nil),           // 0: pb.DateRange
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_agokit_pb_daterange_proto_depIdxs = []int32{
	1, // 0: pb.DateRange.from:type_name -> google.protobuf.Timestamp
	1, // 1: pb.DateRange.to:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_agokit_pb_daterange_proto_init() }
func file_agokit_pb_daterange_proto_init() {
	if File_agokit_pb_daterange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agokit_pb_daterange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateRange); i {
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
			RawDescriptor: file_agokit_pb_daterange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_agokit_pb_daterange_proto_goTypes,
		DependencyIndexes: file_agokit_pb_daterange_proto_depIdxs,
		MessageInfos:      file_agokit_pb_daterange_proto_msgTypes,
	}.Build()
	File_agokit_pb_daterange_proto = out.File
	file_agokit_pb_daterange_proto_rawDesc = nil
	file_agokit_pb_daterange_proto_goTypes = nil
	file_agokit_pb_daterange_proto_depIdxs = nil
}
