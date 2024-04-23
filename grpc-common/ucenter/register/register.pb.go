// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: ucenter.proto

package register

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

type RegReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string      `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password     string      `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Captcha      *CaptchaReq `protobuf:"bytes,3,opt,name=captcha,proto3" json:"captcha,omitempty"`
	Phone        string      `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Promotion    string      `protobuf:"bytes,5,opt,name=promotion,proto3" json:"promotion,omitempty"`
	Code         string      `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`
	Country      string      `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	SuperPartner string      `protobuf:"bytes,8,opt,name=superPartner,proto3" json:"superPartner,omitempty"`
}

func (x *RegReq) Reset() {
	*x = RegReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegReq) ProtoMessage() {}

func (x *RegReq) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegReq.ProtoReflect.Descriptor instead.
func (*RegReq) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{0}
}

func (x *RegReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegReq) GetCaptcha() *CaptchaReq {
	if x != nil {
		return x.Captcha
	}
	return nil
}

func (x *RegReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegReq) GetPromotion() string {
	if x != nil {
		return x.Promotion
	}
	return ""
}

func (x *RegReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RegReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *RegReq) GetSuperPartner() string {
	if x != nil {
		return x.SuperPartner
	}
	return ""
}

type CaptchaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CaptchaReq) Reset() {
	*x = CaptchaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptchaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptchaReq) ProtoMessage() {}

func (x *CaptchaReq) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptchaReq.ProtoReflect.Descriptor instead.
func (*CaptchaReq) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{1}
}

func (x *CaptchaReq) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *CaptchaReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RegRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegRes) Reset() {
	*x = RegRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegRes) ProtoMessage() {}

func (x *RegRes) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegRes.ProtoReflect.Descriptor instead.
func (*RegRes) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{2}
}

var File_register_proto protoreflect.FileDescriptor

var file_register_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0xf6, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2e, 0x0a,
	0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x52, 0x65, 0x71, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x0a, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x08, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x32, 0x41, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_register_proto_rawDescOnce sync.Once
	file_register_proto_rawDescData = file_register_proto_rawDesc
)

func file_register_proto_rawDescGZIP() []byte {
	file_register_proto_rawDescOnce.Do(func() {
		file_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_proto_rawDescData)
	})
	return file_register_proto_rawDescData
}

var file_register_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_register_proto_goTypes = []interface{}{
	(*RegReq)(nil),     // 0: ucenter.RegReq
	(*CaptchaReq)(nil), // 1: ucenter.CaptchaReq
	(*RegRes)(nil),     // 2: ucenter.RegRes
}
var file_register_proto_depIdxs = []int32{
	1, // 0: ucenter.RegReq.captcha:type_name -> ucenter.CaptchaReq
	0, // 1: ucenter.Register.registerByPhone:input_type -> ucenter.RegReq
	2, // 2: ucenter.Register.registerByPhone:output_type -> ucenter.RegRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_register_proto_init() }
func file_register_proto_init() {
	if File_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegReq); i {
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
		file_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptchaReq); i {
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
		file_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegRes); i {
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
			RawDescriptor: file_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_register_proto_goTypes,
		DependencyIndexes: file_register_proto_depIdxs,
		MessageInfos:      file_register_proto_msgTypes,
	}.Build()
	File_register_proto = out.File
	file_register_proto_rawDesc = nil
	file_register_proto_goTypes = nil
	file_register_proto_depIdxs = nil
}
