// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0--rc2
// source: contact_verification.proto

package protobuf

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

type RequestContactVerification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock     uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Challenge string `protobuf:"bytes,3,opt,name=challenge,proto3" json:"challenge,omitempty"`
}

func (x *RequestContactVerification) Reset() {
	*x = RequestContactVerification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_verification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestContactVerification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestContactVerification) ProtoMessage() {}

func (x *RequestContactVerification) ProtoReflect() protoreflect.Message {
	mi := &file_contact_verification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestContactVerification.ProtoReflect.Descriptor instead.
func (*RequestContactVerification) Descriptor() ([]byte, []int) {
	return file_contact_verification_proto_rawDescGZIP(), []int{0}
}

func (x *RequestContactVerification) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *RequestContactVerification) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

type AcceptContactVerification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock    uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Response string `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AcceptContactVerification) Reset() {
	*x = AcceptContactVerification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_verification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptContactVerification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptContactVerification) ProtoMessage() {}

func (x *AcceptContactVerification) ProtoReflect() protoreflect.Message {
	mi := &file_contact_verification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptContactVerification.ProtoReflect.Descriptor instead.
func (*AcceptContactVerification) Descriptor() ([]byte, []int) {
	return file_contact_verification_proto_rawDescGZIP(), []int{1}
}

func (x *AcceptContactVerification) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *AcceptContactVerification) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type ContactVerificationTrusted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (x *ContactVerificationTrusted) Reset() {
	*x = ContactVerificationTrusted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_verification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactVerificationTrusted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactVerificationTrusted) ProtoMessage() {}

func (x *ContactVerificationTrusted) ProtoReflect() protoreflect.Message {
	mi := &file_contact_verification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactVerificationTrusted.ProtoReflect.Descriptor instead.
func (*ContactVerificationTrusted) Descriptor() ([]byte, []int) {
	return file_contact_verification_proto_rawDescGZIP(), []int{2}
}

func (x *ContactVerificationTrusted) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

type DeclineContactVerification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (x *DeclineContactVerification) Reset() {
	*x = DeclineContactVerification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_verification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeclineContactVerification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclineContactVerification) ProtoMessage() {}

func (x *DeclineContactVerification) ProtoReflect() protoreflect.Message {
	mi := &file_contact_verification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclineContactVerification.ProtoReflect.Descriptor instead.
func (*DeclineContactVerification) Descriptor() ([]byte, []int) {
	return file_contact_verification_proto_rawDescGZIP(), []int{3}
}

func (x *DeclineContactVerification) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

var File_contact_verification_proto protoreflect.FileDescriptor

var file_contact_verification_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x50, 0x0a, 0x1a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x19, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x32, 0x0a, 0x1a, 0x44,
	0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contact_verification_proto_rawDescOnce sync.Once
	file_contact_verification_proto_rawDescData = file_contact_verification_proto_rawDesc
)

func file_contact_verification_proto_rawDescGZIP() []byte {
	file_contact_verification_proto_rawDescOnce.Do(func() {
		file_contact_verification_proto_rawDescData = protoimpl.X.CompressGZIP(file_contact_verification_proto_rawDescData)
	})
	return file_contact_verification_proto_rawDescData
}

var file_contact_verification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_contact_verification_proto_goTypes = []interface{}{
	(*RequestContactVerification)(nil), // 0: protobuf.RequestContactVerification
	(*AcceptContactVerification)(nil),  // 1: protobuf.AcceptContactVerification
	(*ContactVerificationTrusted)(nil), // 2: protobuf.ContactVerificationTrusted
	(*DeclineContactVerification)(nil), // 3: protobuf.DeclineContactVerification
}
var file_contact_verification_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contact_verification_proto_init() }
func file_contact_verification_proto_init() {
	if File_contact_verification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contact_verification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestContactVerification); i {
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
		file_contact_verification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptContactVerification); i {
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
		file_contact_verification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactVerificationTrusted); i {
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
		file_contact_verification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineContactVerification); i {
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
			RawDescriptor: file_contact_verification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contact_verification_proto_goTypes,
		DependencyIndexes: file_contact_verification_proto_depIdxs,
		MessageInfos:      file_contact_verification_proto_msgTypes,
	}.Build()
	File_contact_verification_proto = out.File
	file_contact_verification_proto_rawDesc = nil
	file_contact_verification_proto_goTypes = nil
	file_contact_verification_proto_depIdxs = nil
}