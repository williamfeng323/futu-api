// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: Qot_GetReference.proto

package qotgetreference

import (
	_ "github.com/williamfeng323/futu-api/protos/pb/common"
	qotcommon "github.com/williamfeng323/futu-api/protos/pb/qotcommon"
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

type ReferenceType int32

const (
	ReferenceType_ReferenceType_Unknow  ReferenceType = 0
	ReferenceType_ReferenceType_Warrant ReferenceType = 1 //正股相关的窝轮
	ReferenceType_ReferenceType_Future  ReferenceType = 2 //期货主连的相关合约
)

// Enum value maps for ReferenceType.
var (
	ReferenceType_name = map[int32]string{
		0: "ReferenceType_Unknow",
		1: "ReferenceType_Warrant",
		2: "ReferenceType_Future",
	}
	ReferenceType_value = map[string]int32{
		"ReferenceType_Unknow":  0,
		"ReferenceType_Warrant": 1,
		"ReferenceType_Future":  2,
	}
)

func (x ReferenceType) Enum() *ReferenceType {
	p := new(ReferenceType)
	*p = x
	return p
}

func (x ReferenceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReferenceType) Descriptor() protoreflect.EnumDescriptor {
	return file_Qot_GetReference_proto_enumTypes[0].Descriptor()
}

func (ReferenceType) Type() protoreflect.EnumType {
	return &file_Qot_GetReference_proto_enumTypes[0]
}

func (x ReferenceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ReferenceType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ReferenceType(num)
	return nil
}

// Deprecated: Use ReferenceType.Descriptor instead.
func (ReferenceType) EnumDescriptor() ([]byte, []int) {
	return file_Qot_GetReference_proto_rawDescGZIP(), []int{0}
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security      *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`            //股票
	ReferenceType *int32              `protobuf:"varint,2,req,name=referenceType" json:"referenceType,omitempty"` // ReferenceType, 相关类型
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetReference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetReference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Qot_GetReference_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *C2S) GetReferenceType() int32 {
	if x != nil && x.ReferenceType != nil {
		return *x.ReferenceType
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaticInfoList []*qotcommon.SecurityStaticInfo `protobuf:"bytes,2,rep,name=staticInfoList" json:"staticInfoList,omitempty"` //相关股票列表
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetReference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetReference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Qot_GetReference_proto_rawDescGZIP(), []int{1}
}

func (x *S2C) GetStaticInfoList() []*qotcommon.SecurityStaticInfo {
	if x != nil {
		return x.StaticInfoList
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetReference_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetReference_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Qot_GetReference_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetReference_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetReference_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Qot_GetReference_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Qot_GetReference_proto protoreflect.FileDescriptor

var file_Qot_GetReference_proto_rawDesc = []byte{
	0x0a, 0x16, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x03, 0x43, 0x32,
	0x53, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x03, 0x53, 0x32, 0x43,
	0x12, 0x46, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x85, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x73,
	0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x32, 0x43, 0x52,
	0x03, 0x73, 0x32, 0x63, 0x2a, 0x5e, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x57, 0x61, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x46, 0x75, 0x74, 0x75,
	0x72, 0x65, 0x10, 0x02, 0x42, 0x4c, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x66,
	0x65, 0x6e, 0x67, 0x33, 0x32, 0x33, 0x2f, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x62, 0x2f, 0x71, 0x6f, 0x74, 0x67, 0x65, 0x74, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65,
}

var (
	file_Qot_GetReference_proto_rawDescOnce sync.Once
	file_Qot_GetReference_proto_rawDescData = file_Qot_GetReference_proto_rawDesc
)

func file_Qot_GetReference_proto_rawDescGZIP() []byte {
	file_Qot_GetReference_proto_rawDescOnce.Do(func() {
		file_Qot_GetReference_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetReference_proto_rawDescData)
	})
	return file_Qot_GetReference_proto_rawDescData
}

var file_Qot_GetReference_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Qot_GetReference_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Qot_GetReference_proto_goTypes = []interface{}{
	(ReferenceType)(0),                   // 0: Qot_GetReference.ReferenceType
	(*C2S)(nil),                          // 1: Qot_GetReference.C2S
	(*S2C)(nil),                          // 2: Qot_GetReference.S2C
	(*Request)(nil),                      // 3: Qot_GetReference.Request
	(*Response)(nil),                     // 4: Qot_GetReference.Response
	(*qotcommon.Security)(nil),           // 5: Qot_Common.Security
	(*qotcommon.SecurityStaticInfo)(nil), // 6: Qot_Common.SecurityStaticInfo
}
var file_Qot_GetReference_proto_depIdxs = []int32{
	5, // 0: Qot_GetReference.C2S.security:type_name -> Qot_Common.Security
	6, // 1: Qot_GetReference.S2C.staticInfoList:type_name -> Qot_Common.SecurityStaticInfo
	1, // 2: Qot_GetReference.Request.c2s:type_name -> Qot_GetReference.C2S
	2, // 3: Qot_GetReference.Response.s2c:type_name -> Qot_GetReference.S2C
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_Qot_GetReference_proto_init() }
func file_Qot_GetReference_proto_init() {
	if File_Qot_GetReference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetReference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_Qot_GetReference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_Qot_GetReference_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Qot_GetReference_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_Qot_GetReference_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetReference_proto_goTypes,
		DependencyIndexes: file_Qot_GetReference_proto_depIdxs,
		EnumInfos:         file_Qot_GetReference_proto_enumTypes,
		MessageInfos:      file_Qot_GetReference_proto_msgTypes,
	}.Build()
	File_Qot_GetReference_proto = out.File
	file_Qot_GetReference_proto_rawDesc = nil
	file_Qot_GetReference_proto_goTypes = nil
	file_Qot_GetReference_proto_depIdxs = nil
}
