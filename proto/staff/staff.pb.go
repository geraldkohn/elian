// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: proto/staff/staff.proto

package staff

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

type CreateStaffRequset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password     string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Hospital     string `protobuf:"bytes,3,opt,name=hospital,proto3" json:"hospital,omitempty"`
	Department   string `protobuf:"bytes,4,opt,name=department,proto3" json:"department,omitempty"`
	JobNumber    string `protobuf:"bytes,5,opt,name=job_number,json=jobNumber,proto3" json:"job_number,omitempty"`
	IdCardNumber string `protobuf:"bytes,6,opt,name=id_card_number,json=idCardNumber,proto3" json:"id_card_number,omitempty"`
}

func (x *CreateStaffRequset) Reset() {
	*x = CreateStaffRequset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_staff_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffRequset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffRequset) ProtoMessage() {}

func (x *CreateStaffRequset) ProtoReflect() protoreflect.Message {
	mi := &file_proto_staff_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffRequset.ProtoReflect.Descriptor instead.
func (*CreateStaffRequset) Descriptor() ([]byte, []int) {
	return file_proto_staff_staff_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStaffRequset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStaffRequset) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateStaffRequset) GetHospital() string {
	if x != nil {
		return x.Hospital
	}
	return ""
}

func (x *CreateStaffRequset) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *CreateStaffRequset) GetJobNumber() string {
	if x != nil {
		return x.JobNumber
	}
	return ""
}

func (x *CreateStaffRequset) GetIdCardNumber() string {
	if x != nil {
		return x.IdCardNumber
	}
	return ""
}

type CreateStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg       string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	ErrorCode int64  `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateStaffResponse) Reset() {
	*x = CreateStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_staff_staff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffResponse) ProtoMessage() {}

func (x *CreateStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_staff_staff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffResponse.ProtoReflect.Descriptor instead.
func (*CreateStaffResponse) Descriptor() ([]byte, []int) {
	return file_proto_staff_staff_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStaffResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateStaffResponse) GetErrorCode() int64 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *CreateStaffResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type StaffLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IdCardNumber string `protobuf:"bytes,2,opt,name=id_card_number,json=idCardNumber,proto3" json:"id_card_number,omitempty"`
	Password     string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *StaffLoginRequest) Reset() {
	*x = StaffLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_staff_staff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffLoginRequest) ProtoMessage() {}

func (x *StaffLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_staff_staff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffLoginRequest.ProtoReflect.Descriptor instead.
func (*StaffLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_staff_staff_proto_rawDescGZIP(), []int{2}
}

func (x *StaffLoginRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StaffLoginRequest) GetIdCardNumber() string {
	if x != nil {
		return x.IdCardNumber
	}
	return ""
}

func (x *StaffLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type StaffLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg       string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	ErrorCode int64  `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *StaffLoginResponse) Reset() {
	*x = StaffLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_staff_staff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffLoginResponse) ProtoMessage() {}

func (x *StaffLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_staff_staff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffLoginResponse.ProtoReflect.Descriptor instead.
func (*StaffLoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_staff_staff_proto_rawDescGZIP(), []int{3}
}

func (x *StaffLoginResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *StaffLoginResponse) GetErrorCode() int64 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *StaffLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_proto_staff_staff_proto protoreflect.FileDescriptor

var file_proto_staff_staff_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x22, 0xc5, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x70, 0x69,
	0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x70, 0x69,
	0x74, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6b, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x64, 0x43, 0x61, 0x72,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x5b, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x32, 0x9b, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x12, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x73, 0x65, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_staff_staff_proto_rawDescOnce sync.Once
	file_proto_staff_staff_proto_rawDescData = file_proto_staff_staff_proto_rawDesc
)

func file_proto_staff_staff_proto_rawDescGZIP() []byte {
	file_proto_staff_staff_proto_rawDescOnce.Do(func() {
		file_proto_staff_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_staff_staff_proto_rawDescData)
	})
	return file_proto_staff_staff_proto_rawDescData
}

var file_proto_staff_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_staff_staff_proto_goTypes = []interface{}{
	(*CreateStaffRequset)(nil),  // 0: staff.CreateStaffRequset
	(*CreateStaffResponse)(nil), // 1: staff.CreateStaffResponse
	(*StaffLoginRequest)(nil),   // 2: staff.StaffLoginRequest
	(*StaffLoginResponse)(nil),  // 3: staff.StaffLoginResponse
}
var file_proto_staff_staff_proto_depIdxs = []int32{
	0, // 0: staff.StaffService.CreateStaff:input_type -> staff.CreateStaffRequset
	2, // 1: staff.StaffService.StaffLogin:input_type -> staff.StaffLoginRequest
	1, // 2: staff.StaffService.CreateStaff:output_type -> staff.CreateStaffResponse
	3, // 3: staff.StaffService.StaffLogin:output_type -> staff.StaffLoginResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_staff_staff_proto_init() }
func file_proto_staff_staff_proto_init() {
	if File_proto_staff_staff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_staff_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffRequset); i {
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
		file_proto_staff_staff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffResponse); i {
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
		file_proto_staff_staff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffLoginRequest); i {
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
		file_proto_staff_staff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffLoginResponse); i {
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
			RawDescriptor: file_proto_staff_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_staff_staff_proto_goTypes,
		DependencyIndexes: file_proto_staff_staff_proto_depIdxs,
		MessageInfos:      file_proto_staff_staff_proto_msgTypes,
	}.Build()
	File_proto_staff_staff_proto = out.File
	file_proto_staff_staff_proto_rawDesc = nil
	file_proto_staff_staff_proto_goTypes = nil
	file_proto_staff_staff_proto_depIdxs = nil
}
