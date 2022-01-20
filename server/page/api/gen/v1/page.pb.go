// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: page.proto

package pagepb

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Html     string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Tag      string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

func (x *Block) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Block) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type BlockEmtity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Block *Block `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *BlockEmtity) Reset() {
	*x = BlockEmtity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockEmtity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockEmtity) ProtoMessage() {}

func (x *BlockEmtity) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockEmtity.ProtoReflect.Descriptor instead.
func (*BlockEmtity) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{1}
}

func (x *BlockEmtity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BlockEmtity) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatorId string         `protobuf:"bytes,1,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Blocks    []*BlockEmtity `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
	CreatedAt int32          `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int32          `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{2}
}

func (x *Page) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Page) GetBlocks() []*BlockEmtity {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *Page) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Page) GetUpdatedAt() int32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type PageEmtity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Page *Page  `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *PageEmtity) Reset() {
	*x = PageEmtity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageEmtity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageEmtity) ProtoMessage() {}

func (x *PageEmtity) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageEmtity.ProtoReflect.Descriptor instead.
func (*PageEmtity) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{3}
}

func (x *PageEmtity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PageEmtity) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type GetPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPageRequest) Reset() {
	*x = GetPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageRequest) ProtoMessage() {}

func (x *GetPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageRequest.ProtoReflect.Descriptor instead.
func (*GetPageRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{4}
}

func (x *GetPageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPagesRequest) Reset() {
	*x = GetPagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPagesRequest) ProtoMessage() {}

func (x *GetPagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPagesRequest.ProtoReflect.Descriptor instead.
func (*GetPagesRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{5}
}

type GetPagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pages []*PageEmtity `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (x *GetPagesResponse) Reset() {
	*x = GetPagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPagesResponse) ProtoMessage() {}

func (x *GetPagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPagesResponse.ProtoReflect.Descriptor instead.
func (*GetPagesResponse) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{6}
}

func (x *GetPagesResponse) GetPages() []*PageEmtity {
	if x != nil {
		return x.Pages
	}
	return nil
}

type CreatePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Page *Page  `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *CreatePageRequest) Reset() {
	*x = CreatePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePageRequest) ProtoMessage() {}

func (x *CreatePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePageRequest.ProtoReflect.Descriptor instead.
func (*CreatePageRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{7}
}

func (x *CreatePageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePageRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type UpdatePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Page *Page  `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *UpdatePageRequest) Reset() {
	*x = UpdatePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePageRequest) ProtoMessage() {}

func (x *UpdatePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePageRequest.ProtoReflect.Descriptor instead.
func (*UpdatePageRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePageRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type DeletePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePageRequest) Reset() {
	*x = DeletePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageRequest) ProtoMessage() {}

func (x *DeletePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageRequest.ProtoReflect.Descriptor instead.
func (*DeletePageRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{9}
}

func (x *DeletePageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePageResponse) Reset() {
	*x = DeletePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageResponse) ProtoMessage() {}

func (x *DeletePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageResponse.ProtoReflect.Descriptor instead.
func (*DeletePageResponse) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{10}
}

var File_page_proto protoreflect.FileDescriptor

var file_page_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x4a, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74,
	0x6d, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x22, 0x43, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6d, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x91, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6d,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x0a, 0x50, 0x61,
	0x67, 0x65, 0x45, 0x6d, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x45, 0x6d, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x46, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xcc, 0x02, 0x0a, 0x0b, 0x50,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x45, 0x6d, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x18, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x45, 0x6d, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x45, 0x6d, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x45, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x6e, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_page_proto_rawDescOnce sync.Once
	file_page_proto_rawDescData = file_page_proto_rawDesc
)

func file_page_proto_rawDescGZIP() []byte {
	file_page_proto_rawDescOnce.Do(func() {
		file_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_page_proto_rawDescData)
	})
	return file_page_proto_rawDescData
}

var file_page_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_page_proto_goTypes = []interface{}{
	(*Block)(nil),              // 0: page.v1.Block
	(*BlockEmtity)(nil),        // 1: page.v1.BlockEmtity
	(*Page)(nil),               // 2: page.v1.Page
	(*PageEmtity)(nil),         // 3: page.v1.PageEmtity
	(*GetPageRequest)(nil),     // 4: page.v1.GetPageRequest
	(*GetPagesRequest)(nil),    // 5: page.v1.GetPagesRequest
	(*GetPagesResponse)(nil),   // 6: page.v1.GetPagesResponse
	(*CreatePageRequest)(nil),  // 7: page.v1.CreatePageRequest
	(*UpdatePageRequest)(nil),  // 8: page.v1.UpdatePageRequest
	(*DeletePageRequest)(nil),  // 9: page.v1.DeletePageRequest
	(*DeletePageResponse)(nil), // 10: page.v1.DeletePageResponse
}
var file_page_proto_depIdxs = []int32{
	0,  // 0: page.v1.BlockEmtity.block:type_name -> page.v1.Block
	1,  // 1: page.v1.Page.blocks:type_name -> page.v1.BlockEmtity
	2,  // 2: page.v1.PageEmtity.page:type_name -> page.v1.Page
	3,  // 3: page.v1.GetPagesResponse.pages:type_name -> page.v1.PageEmtity
	2,  // 4: page.v1.CreatePageRequest.page:type_name -> page.v1.Page
	2,  // 5: page.v1.UpdatePageRequest.page:type_name -> page.v1.Page
	4,  // 6: page.v1.PageService.GetPage:input_type -> page.v1.GetPageRequest
	5,  // 7: page.v1.PageService.GetPages:input_type -> page.v1.GetPagesRequest
	7,  // 8: page.v1.PageService.CreatePage:input_type -> page.v1.CreatePageRequest
	8,  // 9: page.v1.PageService.UpdatePage:input_type -> page.v1.UpdatePageRequest
	9,  // 10: page.v1.PageService.DeletePage:input_type -> page.v1.DeletePageRequest
	3,  // 11: page.v1.PageService.GetPage:output_type -> page.v1.PageEmtity
	6,  // 12: page.v1.PageService.GetPages:output_type -> page.v1.GetPagesResponse
	3,  // 13: page.v1.PageService.CreatePage:output_type -> page.v1.PageEmtity
	3,  // 14: page.v1.PageService.UpdatePage:output_type -> page.v1.PageEmtity
	10, // 15: page.v1.PageService.DeletePage:output_type -> page.v1.DeletePageResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_page_proto_init() }
func file_page_proto_init() {
	if File_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockEmtity); i {
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
		file_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageEmtity); i {
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
		file_page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageRequest); i {
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
		file_page_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPagesRequest); i {
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
		file_page_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPagesResponse); i {
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
		file_page_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePageRequest); i {
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
		file_page_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePageRequest); i {
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
		file_page_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePageRequest); i {
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
		file_page_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePageResponse); i {
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
			RawDescriptor: file_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_page_proto_goTypes,
		DependencyIndexes: file_page_proto_depIdxs,
		MessageInfos:      file_page_proto_msgTypes,
	}.Build()
	File_page_proto = out.File
	file_page_proto_rawDesc = nil
	file_page_proto_goTypes = nil
	file_page_proto_depIdxs = nil
}
