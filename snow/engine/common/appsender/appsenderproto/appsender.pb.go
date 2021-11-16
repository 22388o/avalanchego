// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: appsender.proto

package appsenderproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendAppRequestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The nodes to send this request to
	NodeIDs [][]byte `protobuf:"bytes,1,rep,name=nodeIDs,proto3" json:"nodeIDs,omitempty"`
	// The ID of this request
	RequestID uint32 `protobuf:"varint,2,opt,name=requestID,proto3" json:"requestID,omitempty"`
	// The request body
	Request []byte `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *SendAppRequestMsg) Reset() {
	*x = SendAppRequestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appsender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAppRequestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAppRequestMsg) ProtoMessage() {}

func (x *SendAppRequestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_appsender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAppRequestMsg.ProtoReflect.Descriptor instead.
func (*SendAppRequestMsg) Descriptor() ([]byte, []int) {
	return file_appsender_proto_rawDescGZIP(), []int{0}
}

func (x *SendAppRequestMsg) GetNodeIDs() [][]byte {
	if x != nil {
		return x.NodeIDs
	}
	return nil
}

func (x *SendAppRequestMsg) GetRequestID() uint32 {
	if x != nil {
		return x.RequestID
	}
	return 0
}

func (x *SendAppRequestMsg) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

type SendAppResponseMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node to send a response to
	NodeID []byte `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	// ID of this request
	RequestID uint32 `protobuf:"varint,2,opt,name=requestID,proto3" json:"requestID,omitempty"`
	// The response body
	Response []byte `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SendAppResponseMsg) Reset() {
	*x = SendAppResponseMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appsender_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAppResponseMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAppResponseMsg) ProtoMessage() {}

func (x *SendAppResponseMsg) ProtoReflect() protoreflect.Message {
	mi := &file_appsender_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAppResponseMsg.ProtoReflect.Descriptor instead.
func (*SendAppResponseMsg) Descriptor() ([]byte, []int) {
	return file_appsender_proto_rawDescGZIP(), []int{1}
}

func (x *SendAppResponseMsg) GetNodeID() []byte {
	if x != nil {
		return x.NodeID
	}
	return nil
}

func (x *SendAppResponseMsg) GetRequestID() uint32 {
	if x != nil {
		return x.RequestID
	}
	return 0
}

func (x *SendAppResponseMsg) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

type SendAppGossipMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message body
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendAppGossipMsg) Reset() {
	*x = SendAppGossipMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appsender_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAppGossipMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAppGossipMsg) ProtoMessage() {}

func (x *SendAppGossipMsg) ProtoReflect() protoreflect.Message {
	mi := &file_appsender_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAppGossipMsg.ProtoReflect.Descriptor instead.
func (*SendAppGossipMsg) Descriptor() ([]byte, []int) {
	return file_appsender_proto_rawDescGZIP(), []int{2}
}

func (x *SendAppGossipMsg) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

type SendAppGossipSpecificMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The nodes to send this request to
	NodeIDs [][]byte `protobuf:"bytes,1,rep,name=nodeIDs,proto3" json:"nodeIDs,omitempty"`
	// The message body
	Msg []byte `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendAppGossipSpecificMsg) Reset() {
	*x = SendAppGossipSpecificMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appsender_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAppGossipSpecificMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAppGossipSpecificMsg) ProtoMessage() {}

func (x *SendAppGossipSpecificMsg) ProtoReflect() protoreflect.Message {
	mi := &file_appsender_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAppGossipSpecificMsg.ProtoReflect.Descriptor instead.
func (*SendAppGossipSpecificMsg) Descriptor() ([]byte, []int) {
	return file_appsender_proto_rawDescGZIP(), []int{3}
}

func (x *SendAppGossipSpecificMsg) GetNodeIDs() [][]byte {
	if x != nil {
		return x.NodeIDs
	}
	return nil
}

func (x *SendAppGossipSpecificMsg) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_appsender_proto protoreflect.FileDescriptor

var file_appsender_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65,
	0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a,
	0x10, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x4d, 0x73,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x46, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x47, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x4d, 0x73, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xcd, 0x02, 0x0a, 0x09,
	0x41, 0x70, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x0e, 0x53, 0x65, 0x6e,
	0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70,
	0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x47,
	0x6f, 0x73, 0x73, 0x69, 0x70, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x59, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41,
	0x70, 0x70, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x4d, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x61, 0x2d, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x67, 0x6f, 0x2f, 0x73,
	0x6e, 0x6f, 0x77, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_appsender_proto_rawDescOnce sync.Once
	file_appsender_proto_rawDescData = file_appsender_proto_rawDesc
)

func file_appsender_proto_rawDescGZIP() []byte {
	file_appsender_proto_rawDescOnce.Do(func() {
		file_appsender_proto_rawDescData = protoimpl.X.CompressGZIP(file_appsender_proto_rawDescData)
	})
	return file_appsender_proto_rawDescData
}

var file_appsender_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_appsender_proto_goTypes = []interface{}{
	(*SendAppRequestMsg)(nil),        // 0: appsenderproto.SendAppRequestMsg
	(*SendAppResponseMsg)(nil),       // 1: appsenderproto.SendAppResponseMsg
	(*SendAppGossipMsg)(nil),         // 2: appsenderproto.SendAppGossipMsg
	(*SendAppGossipSpecificMsg)(nil), // 3: appsenderproto.SendAppGossipSpecificMsg
	(*emptypb.Empty)(nil),            // 4: google.protobuf.Empty
}
var file_appsender_proto_depIdxs = []int32{
	0, // 0: appsenderproto.AppSender.SendAppRequest:input_type -> appsenderproto.SendAppRequestMsg
	1, // 1: appsenderproto.AppSender.SendAppResponse:input_type -> appsenderproto.SendAppResponseMsg
	2, // 2: appsenderproto.AppSender.SendAppGossip:input_type -> appsenderproto.SendAppGossipMsg
	3, // 3: appsenderproto.AppSender.SendAppGossipSpecific:input_type -> appsenderproto.SendAppGossipSpecificMsg
	4, // 4: appsenderproto.AppSender.SendAppRequest:output_type -> google.protobuf.Empty
	4, // 5: appsenderproto.AppSender.SendAppResponse:output_type -> google.protobuf.Empty
	4, // 6: appsenderproto.AppSender.SendAppGossip:output_type -> google.protobuf.Empty
	4, // 7: appsenderproto.AppSender.SendAppGossipSpecific:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_appsender_proto_init() }
func file_appsender_proto_init() {
	if File_appsender_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appsender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAppRequestMsg); i {
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
		file_appsender_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAppResponseMsg); i {
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
		file_appsender_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAppGossipMsg); i {
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
		file_appsender_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAppGossipSpecificMsg); i {
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
			RawDescriptor: file_appsender_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appsender_proto_goTypes,
		DependencyIndexes: file_appsender_proto_depIdxs,
		MessageInfos:      file_appsender_proto_msgTypes,
	}.Build()
	File_appsender_proto = out.File
	file_appsender_proto_rawDesc = nil
	file_appsender_proto_goTypes = nil
	file_appsender_proto_depIdxs = nil
}
