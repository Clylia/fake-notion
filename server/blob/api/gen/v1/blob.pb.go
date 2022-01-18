// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: blob.proto

package blobpb

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

type CreateBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId           string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	UploadUrlTimeoutSec int32  `protobuf:"varint,2,opt,name=upload_url_timeout_sec,json=uploadUrlTimeoutSec,proto3" json:"upload_url_timeout_sec,omitempty"`
}

func (x *CreateBlobRequest) Reset() {
	*x = CreateBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlobRequest) ProtoMessage() {}

func (x *CreateBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlobRequest.ProtoReflect.Descriptor instead.
func (*CreateBlobRequest) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBlobRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CreateBlobRequest) GetUploadUrlTimeoutSec() int32 {
	if x != nil {
		return x.UploadUrlTimeoutSec
	}
	return 0
}

type CreateBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UploadUrl string `protobuf:"bytes,2,opt,name=upload_url,json=uploadUrl,proto3" json:"upload_url,omitempty"`
}

func (x *CreateBlobResponse) Reset() {
	*x = CreateBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlobResponse) ProtoMessage() {}

func (x *CreateBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlobResponse.ProtoReflect.Descriptor instead.
func (*CreateBlobResponse) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBlobResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateBlobResponse) GetUploadUrl() string {
	if x != nil {
		return x.UploadUrl
	}
	return ""
}

type GetBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBlobRequest) Reset() {
	*x = GetBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlobRequest) ProtoMessage() {}

func (x *GetBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlobRequest.ProtoReflect.Descriptor instead.
func (*GetBlobRequest) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetBlobResponse) Reset() {
	*x = GetBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlobResponse) ProtoMessage() {}

func (x *GetBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlobResponse.ProtoReflect.Descriptor instead.
func (*GetBlobResponse) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{3}
}

func (x *GetBlobResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetBlobURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TimeoutSec int32  `protobuf:"varint,2,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
}

func (x *GetBlobURLRequest) Reset() {
	*x = GetBlobURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlobURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlobURLRequest) ProtoMessage() {}

func (x *GetBlobURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlobURLRequest.ProtoReflect.Descriptor instead.
func (*GetBlobURLRequest) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{4}
}

func (x *GetBlobURLRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBlobURLRequest) GetTimeoutSec() int32 {
	if x != nil {
		return x.TimeoutSec
	}
	return 0
}

type GetBlobURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetBlobURLResponse) Reset() {
	*x = GetBlobURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlobURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlobURLResponse) ProtoMessage() {}

func (x *GetBlobURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlobURLResponse.ProtoReflect.Descriptor instead.
func (*GetBlobURLResponse) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{5}
}

func (x *GetBlobURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_blob_proto protoreflect.FileDescriptor

var file_blob_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6e, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x67, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63,
	0x22, 0x43, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x44,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73,
	0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x53, 0x65, 0x63, 0x22, 0x26, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xe5, 0x01, 0x0a,
	0x0b, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1c, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x62, 0x12, 0x19, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x62, 0x55, 0x52, 0x4c, 0x12, 0x1c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62,
	0x6c, 0x6f, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x62,
	0x6c, 0x6f, 0x62, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blob_proto_rawDescOnce sync.Once
	file_blob_proto_rawDescData = file_blob_proto_rawDesc
)

func file_blob_proto_rawDescGZIP() []byte {
	file_blob_proto_rawDescOnce.Do(func() {
		file_blob_proto_rawDescData = protoimpl.X.CompressGZIP(file_blob_proto_rawDescData)
	})
	return file_blob_proto_rawDescData
}

var file_blob_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_blob_proto_goTypes = []interface{}{
	(*CreateBlobRequest)(nil),  // 0: notion.v1.CreateBlobRequest
	(*CreateBlobResponse)(nil), // 1: notion.v1.CreateBlobResponse
	(*GetBlobRequest)(nil),     // 2: notion.v1.GetBlobRequest
	(*GetBlobResponse)(nil),    // 3: notion.v1.GetBlobResponse
	(*GetBlobURLRequest)(nil),  // 4: notion.v1.GetBlobURLRequest
	(*GetBlobURLResponse)(nil), // 5: notion.v1.GetBlobURLResponse
}
var file_blob_proto_depIdxs = []int32{
	0, // 0: notion.v1.BlobService.CreateBlob:input_type -> notion.v1.CreateBlobRequest
	2, // 1: notion.v1.BlobService.GetBlob:input_type -> notion.v1.GetBlobRequest
	4, // 2: notion.v1.BlobService.GetBlobURL:input_type -> notion.v1.GetBlobURLRequest
	1, // 3: notion.v1.BlobService.CreateBlob:output_type -> notion.v1.CreateBlobResponse
	3, // 4: notion.v1.BlobService.GetBlob:output_type -> notion.v1.GetBlobResponse
	5, // 5: notion.v1.BlobService.GetBlobURL:output_type -> notion.v1.GetBlobURLResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blob_proto_init() }
func file_blob_proto_init() {
	if File_blob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlobRequest); i {
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
		file_blob_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlobResponse); i {
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
		file_blob_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlobRequest); i {
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
		file_blob_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlobResponse); i {
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
		file_blob_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlobURLRequest); i {
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
		file_blob_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlobURLResponse); i {
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
			RawDescriptor: file_blob_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blob_proto_goTypes,
		DependencyIndexes: file_blob_proto_depIdxs,
		MessageInfos:      file_blob_proto_msgTypes,
	}.Build()
	File_blob_proto = out.File
	file_blob_proto_rawDesc = nil
	file_blob_proto_goTypes = nil
	file_blob_proto_depIdxs = nil
}
