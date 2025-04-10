// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	FirstName      string                 `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName     string                 `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName       string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	BornDate       int32                  `protobuf:"varint,4,opt,name=born_date,json=bornDate,proto3" json:"born_date,omitempty"`
	PhoneNumber    string                 `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email          string                 `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Password       string                 `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Specialization string                 `protobuf:"bytes,8,opt,name=specialization,proto3" json:"specialization,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RegisterRequest) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *RegisterRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *RegisterRequest) GetBornDate() int32 {
	if x != nil {
		return x.BornDate
	}
	return 0
}

func (x *RegisterRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetSpecialization() string {
	if x != nil {
		return x.Specialization
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NewLogin      string                 `protobuf:"bytes,3,opt,name=new_login,json=newLogin,proto3" json:"new_login,omitempty"`
	Role          string                 `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *RegisterResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RegisterResponse) GetNewLogin() string {
	if x != nil {
		return x.NewLogin
	}
	return ""
}

func (x *RegisterResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role          string                 `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Session       string                 `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	Csrf          string                 `protobuf:"bytes,5,opt,name=csrf,proto3" json:"csrf,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LoginResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *LoginResponse) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *LoginResponse) GetCsrf() string {
	if x != nil {
		return x.Csrf
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

const file_auth_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"auth.proto\"\x88\x02\n" +
	"\x0fRegisterRequest\x12\x1d\n" +
	"\n" +
	"first_name\x18\x01 \x01(\tR\tfirstName\x12\x1f\n" +
	"\vmiddle_name\x18\x02 \x01(\tR\n" +
	"middleName\x12\x1b\n" +
	"\tlast_name\x18\x03 \x01(\tR\blastName\x12\x1b\n" +
	"\tborn_date\x18\x04 \x01(\x05R\bbornDate\x12!\n" +
	"\fphone_number\x18\x05 \x01(\tR\vphoneNumber\x12\x14\n" +
	"\x05email\x18\x06 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\a \x01(\tR\bpassword\x12&\n" +
	"\x0especialization\x18\b \x01(\tR\x0especialization\"l\n" +
	"\x10RegisterResponse\x12\x0e\n" +
	"\x02ok\x18\x01 \x01(\bR\x02ok\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x05R\x06userId\x12\x1b\n" +
	"\tnew_login\x18\x03 \x01(\tR\bnewLogin\x12\x12\n" +
	"\x04role\x18\x04 \x01(\tR\x04role\">\n" +
	"\fLoginRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"~\n" +
	"\rLoginResponse\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x05R\x06userId\x12\x12\n" +
	"\x04role\x18\x03 \x01(\tR\x04role\x12\x18\n" +
	"\asession\x18\x04 \x01(\tR\asession\x12\x12\n" +
	"\x04csrf\x18\x05 \x01(\tR\x04csrf2_\n" +
	"\x04Auth\x12/\n" +
	"\bRegister\x12\x10.RegisterRequest\x1a\x11.RegisterResponse\x12&\n" +
	"\x05Login\x12\r.LoginRequest\x1a\x0e.LoginResponseB.Z,github.com/bigelle/student-portal/proto/authb\x06proto3"

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData []byte
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_proto_rawDesc), len(file_auth_proto_rawDesc)))
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_proto_goTypes = []any{
	(*RegisterRequest)(nil),  // 0: RegisterRequest
	(*RegisterResponse)(nil), // 1: RegisterResponse
	(*LoginRequest)(nil),     // 2: LoginRequest
	(*LoginResponse)(nil),    // 3: LoginResponse
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: Auth.Register:input_type -> RegisterRequest
	2, // 1: Auth.Login:input_type -> LoginRequest
	1, // 2: Auth.Register:output_type -> RegisterResponse
	3, // 3: Auth.Login:output_type -> LoginResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_proto_rawDesc), len(file_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
