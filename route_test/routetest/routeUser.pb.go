// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: routeUser/routeUser.proto

package __

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

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"fixed32,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Mobile   string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeUser_routeUser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_routeUser_routeUser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_routeUser_routeUser_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Profile) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Profile) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeUser_routeUser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_routeUser_routeUser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_routeUser_routeUser_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_routeUser_routeUser_proto protoreflect.FileDescriptor

var file_routeUser_routeUser_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x74, 0x65, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x21, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x41, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x34, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x1a, 0x10, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_routeUser_routeUser_proto_rawDescOnce sync.Once
	file_routeUser_routeUser_proto_rawDescData = file_routeUser_routeUser_proto_rawDesc
)

func file_routeUser_routeUser_proto_rawDescGZIP() []byte {
	file_routeUser_routeUser_proto_rawDescOnce.Do(func() {
		file_routeUser_routeUser_proto_rawDescData = protoimpl.X.CompressGZIP(file_routeUser_routeUser_proto_rawDescData)
	})
	return file_routeUser_routeUser_proto_rawDescData
}

var file_routeUser_routeUser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_routeUser_routeUser_proto_goTypes = []interface{}{
	(*Profile)(nil), // 0: routetest.Profile
	(*Reply)(nil),   // 1: routetest.Reply
}
var file_routeUser_routeUser_proto_depIdxs = []int32{
	0, // 0: routetest.RouteUser.createUser:input_type -> routetest.Profile
	1, // 1: routetest.RouteUser.createUser:output_type -> routetest.Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_routeUser_routeUser_proto_init() }
func file_routeUser_routeUser_proto_init() {
	if File_routeUser_routeUser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routeUser_routeUser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_routeUser_routeUser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_routeUser_routeUser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routeUser_routeUser_proto_goTypes,
		DependencyIndexes: file_routeUser_routeUser_proto_depIdxs,
		MessageInfos:      file_routeUser_routeUser_proto_msgTypes,
	}.Build()
	File_routeUser_routeUser_proto = out.File
	file_routeUser_routeUser_proto_rawDesc = nil
	file_routeUser_routeUser_proto_goTypes = nil
	file_routeUser_routeUser_proto_depIdxs = nil
}