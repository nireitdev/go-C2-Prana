// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: src/proto/grpc.proto

package proto

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

type AdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
}

func (x *AdminRequest) Reset() {
	*x = AdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRequest) ProtoMessage() {}

func (x *AdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRequest.ProtoReflect.Descriptor instead.
func (*AdminRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *AdminRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type AdminResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *AdminResult) Reset() {
	*x = AdminResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminResult) ProtoMessage() {}

func (x *AdminResult) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminResult.ProtoReflect.Descriptor instead.
func (*AdminResult) Descriptor() ([]byte, []int) {
	return file_src_proto_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *AdminResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ImplantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
}

func (x *ImplantRequest) Reset() {
	*x = ImplantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImplantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImplantRequest) ProtoMessage() {}

func (x *ImplantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImplantRequest.ProtoReflect.Descriptor instead.
func (*ImplantRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *ImplantRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type ImplantResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ImplantResult) Reset() {
	*x = ImplantResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImplantResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImplantResult) ProtoMessage() {}

func (x *ImplantResult) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImplantResult.ProtoReflect.Descriptor instead.
func (*ImplantResult) Descriptor() ([]byte, []int) {
	return file_src_proto_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *ImplantResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *PushRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type PushResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *PushResult) Reset() {
	*x = PushResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResult) ProtoMessage() {}

func (x *PushResult) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResult.ProtoReflect.Descriptor instead.
func (*PushResult) Descriptor() ([]byte, []int) {
	return file_src_proto_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *PushResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_src_proto_grpc_proto protoreflect.FileDescriptor

var file_src_proto_grpc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x25, 0x0a, 0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x49, 0x6d, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x27, 0x0a, 0x0b,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x60, 0x0a, 0x06, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x08, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6d,
	0x64, 0x12, 0x0d, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d,
	0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x43, 0x6d, 0x64, 0x12, 0x0f, 0x2e, 0x49,
	0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x33, 0x0a,
	0x0b, 0x50, 0x75, 0x73, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x07,
	0x50, 0x75, 0x73, 0x68, 0x43, 0x6d, 0x64, 0x12, 0x0c, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x6e, 0x69, 0x72, 0x65, 0x69, 0x74, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x43, 0x32,
	0x2d, 0x50, 0x72, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_grpc_proto_rawDescOnce sync.Once
	file_src_proto_grpc_proto_rawDescData = file_src_proto_grpc_proto_rawDesc
)

func file_src_proto_grpc_proto_rawDescGZIP() []byte {
	file_src_proto_grpc_proto_rawDescOnce.Do(func() {
		file_src_proto_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_grpc_proto_rawDescData)
	})
	return file_src_proto_grpc_proto_rawDescData
}

var file_src_proto_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_src_proto_grpc_proto_goTypes = []interface{}{
	(*AdminRequest)(nil),   // 0: AdminRequest
	(*AdminResult)(nil),    // 1: AdminResult
	(*ImplantRequest)(nil), // 2: ImplantRequest
	(*ImplantResult)(nil),  // 3: ImplantResult
	(*PushRequest)(nil),    // 4: PushRequest
	(*PushResult)(nil),     // 5: PushResult
}
var file_src_proto_grpc_proto_depIdxs = []int32{
	0, // 0: Server.AdminCmd:input_type -> AdminRequest
	2, // 1: Server.ImplantCmd:input_type -> ImplantRequest
	4, // 2: PushCommand.PushCmd:input_type -> PushRequest
	1, // 3: Server.AdminCmd:output_type -> AdminResult
	3, // 4: Server.ImplantCmd:output_type -> ImplantResult
	5, // 5: PushCommand.PushCmd:output_type -> PushResult
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_proto_grpc_proto_init() }
func file_src_proto_grpc_proto_init() {
	if File_src_proto_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRequest); i {
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
		file_src_proto_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminResult); i {
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
		file_src_proto_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImplantRequest); i {
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
		file_src_proto_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImplantResult); i {
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
		file_src_proto_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_src_proto_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResult); i {
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
			RawDescriptor: file_src_proto_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_src_proto_grpc_proto_goTypes,
		DependencyIndexes: file_src_proto_grpc_proto_depIdxs,
		MessageInfos:      file_src_proto_grpc_proto_msgTypes,
	}.Build()
	File_src_proto_grpc_proto = out.File
	file_src_proto_grpc_proto_rawDesc = nil
	file_src_proto_grpc_proto_goTypes = nil
	file_src_proto_grpc_proto_depIdxs = nil
}