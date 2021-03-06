// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: agokit/pb/query.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Number of Items per page
	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	//Current page, by default = 1
	Page uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	//Keyword to search  by string contains in name, description
	Keyword string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	//Order by string
	Order string `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
	//Query multiple date fields
	Dates []*DateRange `protobuf:"bytes,5,rep,name=dates,proto3" json:"dates,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agokit_pb_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_agokit_pb_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_agokit_pb_query_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Query) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Query) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *Query) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *Query) GetDates() []*DateRange {
	if x != nil {
		return x.Dates
	}
	return nil
}

var File_agokit_pb_query_proto protoreflect.FileDescriptor

var file_agokit_pb_query_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x67, 0x6f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x61, 0x67, 0x6f,
	0x6b, 0x69, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x42,
	0x0b, 0x5a, 0x09, 0x61, 0x67, 0x6f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agokit_pb_query_proto_rawDescOnce sync.Once
	file_agokit_pb_query_proto_rawDescData = file_agokit_pb_query_proto_rawDesc
)

func file_agokit_pb_query_proto_rawDescGZIP() []byte {
	file_agokit_pb_query_proto_rawDescOnce.Do(func() {
		file_agokit_pb_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_agokit_pb_query_proto_rawDescData)
	})
	return file_agokit_pb_query_proto_rawDescData
}

var file_agokit_pb_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_agokit_pb_query_proto_goTypes = []interface{}{
	(*Query)(nil),     // 0: pb.Query
	(*DateRange)(nil), // 1: pb.DateRange
}
var file_agokit_pb_query_proto_depIdxs = []int32{
	1, // 0: pb.Query.dates:type_name -> pb.DateRange
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_agokit_pb_query_proto_init() }
func file_agokit_pb_query_proto_init() {
	if File_agokit_pb_query_proto != nil {
		return
	}
	file_agokit_pb_daterange_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_agokit_pb_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
			RawDescriptor: file_agokit_pb_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_agokit_pb_query_proto_goTypes,
		DependencyIndexes: file_agokit_pb_query_proto_depIdxs,
		MessageInfos:      file_agokit_pb_query_proto_msgTypes,
	}.Build()
	File_agokit_pb_query_proto = out.File
	file_agokit_pb_query_proto_rawDesc = nil
	file_agokit_pb_query_proto_goTypes = nil
	file_agokit_pb_query_proto_depIdxs = nil
}
