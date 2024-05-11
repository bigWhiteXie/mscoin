// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: ucenter.proto

package login

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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Token         string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	MemberLevel   string `protobuf:"bytes,3,opt,name=memberLevel,proto3" json:"memberLevel,omitempty"`
	RealName      string `protobuf:"bytes,4,opt,name=realName,proto3" json:"realName,omitempty"`
	Country       string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Avatar        string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	PromotionCode string `protobuf:"bytes,7,opt,name=promotionCode,proto3" json:"promotionCode,omitempty"`
	Id            int64  `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
	LoginCount    int32  `protobuf:"varint,9,opt,name=loginCount,proto3" json:"loginCount,omitempty"`
	SuperPartner  string `protobuf:"bytes,10,opt,name=superPartner,proto3" json:"superPartner,omitempty"`
	MemberRate    int32  `protobuf:"varint,11,opt,name=memberRate,proto3" json:"memberRate,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRes) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginRes) GetMemberLevel() string {
	if x != nil {
		return x.MemberLevel
	}
	return ""
}

func (x *LoginRes) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *LoginRes) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *LoginRes) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *LoginRes) GetPromotionCode() string {
	if x != nil {
		return x.PromotionCode
	}
	return ""
}

func (x *LoginRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginRes) GetLoginCount() int32 {
	if x != nil {
		return x.LoginCount
	}
	return 0
}

func (x *LoginRes) GetSuperPartner() string {
	if x != nil {
		return x.SuperPartner
	}
	return ""
}

func (x *LoginRes) GetMemberRate() int32 {
	if x != nil {
		return x.MemberRate
	}
	return 0
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x22, 0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xc6, 0x02, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x61, 0x74,
	0x65, 0x32, 0x32, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x75, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2d, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_login_proto_goTypes = []interface{}{
	(*LoginReq)(nil), // 0: ucenter.LoginReq
	(*LoginRes)(nil), // 1: ucenter.LoginRes
}
var file_login_proto_depIdxs = []int32{
	0, // 0: ucenter.Login.ucenter:input_type -> ucenter.LoginReq
	1, // 1: ucenter.Login.ucenter:output_type -> ucenter.LoginRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
