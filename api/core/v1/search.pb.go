// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/core/v1/search.proto

package v1

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type IndexObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Obj *_struct.Value `protobuf:"bytes,1,opt,name=obj,proto3" json:"obj"`
}

func (x *IndexObject) Reset() {
	*x = IndexObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexObject) ProtoMessage() {}

func (x *IndexObject) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexObject.ProtoReflect.Descriptor instead.
func (*IndexObject) Descriptor() ([]byte, []int) {
	return file_api_core_v1_search_proto_rawDescGZIP(), []int{0}
}

func (x *IndexObject) GetObj() *_struct.Value {
	if x != nil {
		return x.Obj
	}
	return nil
}

type IndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status"`
}

func (x *IndexResponse) Reset() {
	*x = IndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexResponse) ProtoMessage() {}

func (x *IndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexResponse.ProtoReflect.Descriptor instead.
func (*IndexResponse) Descriptor() ([]byte, []int) {
	return file_api_core_v1_search_proto_rawDescGZIP(), []int{1}
}

func (x *IndexResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type SearchCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field    string         `protobuf:"bytes,1,opt,name=field,proto3" json:"field"`
	Operator string         `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator"`
	Value    *_struct.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
}

func (x *SearchCondition) Reset() {
	*x = SearchCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCondition) ProtoMessage() {}

func (x *SearchCondition) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCondition.ProtoReflect.Descriptor instead.
func (*SearchCondition) Descriptor() ([]byte, []int) {
	return file_api_core_v1_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchCondition) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SearchCondition) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *SearchCondition) GetValue() *_struct.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source       string             `protobuf:"bytes,1,opt,name=source,proto3" json:"source"`
	Owner        string             `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner"`
	Query        string             `protobuf:"bytes,3,opt,name=query,proto3" json:"query"`
	Condition    []*SearchCondition `protobuf:"bytes,5,rep,name=condition,proto3" json:"condition"`
	PageNum      int32              `protobuf:"varint,7,opt,name=page_num,json=pageNum,proto3" json:"page_num"`
	PageSize     int32              `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	OrderBy      string             `protobuf:"bytes,9,opt,name=order_by,json=orderBy,proto3" json:"order_by"`
	IsDescending bool               `protobuf:"varint,10,opt,name=is_descending,json=isDescending,proto3" json:"is_descending"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_api_core_v1_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *SearchRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetCondition() []*SearchCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *SearchRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *SearchRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *SearchRequest) GetIsDescending() bool {
	if x != nil {
		return x.IsDescending
	}
	return false
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64            `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	PageNum  int32            `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num"`
	PageSize int32            `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Items    []*_struct.Value `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_api_core_v1_search_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchResponse) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *SearchResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchResponse) GetItems() []*_struct.Value {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
	Owner  string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner"`
}

func (x *DeleteByIDRequest) Reset() {
	*x = DeleteByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByIDRequest) ProtoMessage() {}

func (x *DeleteByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_core_v1_search_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteByIDRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *DeleteByIDRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type DeleteByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteByIDResponse) Reset() {
	*x = DeleteByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_search_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByIDResponse) ProtoMessage() {}

func (x *DeleteByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_search_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByIDResponse.ProtoReflect.Descriptor instead.
func (*DeleteByIDResponse) Descriptor() ([]byte, []int) {
	return file_api_core_v1_search_proto_rawDescGZIP(), []int{6}
}

var File_api_core_v1_search_proto protoreflect.FileDescriptor

var file_api_core_v1_search_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x22, 0x27, 0x0a, 0x0d,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0x92, 0x41, 0x15, 0x32, 0x13, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x20, 0x6b,
	0x65, 0x79, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0x92, 0x41, 0x21,
	0x32, 0x1f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x20, 0x24, 0x67, 0x74, 0x20, 0x24,
	0x67, 0x74, 0x65, 0x20, 0x24, 0x65, 0x71, 0x20, 0x24, 0x6c, 0x74, 0x20, 0x24, 0x6c, 0x74, 0x65,
	0x20, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x15, 0x92, 0x41, 0x12, 0x32, 0x10, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6b, 0x65, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xcc, 0x03, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20,
	0x69, 0x64, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x92, 0x41, 0x0a, 0x32, 0x08,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x20, 0x69, 0x64, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x29, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13,
	0x92, 0x41, 0x10, 0x32, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x50, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x14, 0x92, 0x41, 0x11,
	0x32, 0x0f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17,
	0x92, 0x41, 0x14, 0x32, 0x12, 0xe8, 0xae, 0xb0, 0xe5, 0xbd, 0x95, 0xe5, 0xbc, 0x80, 0xe5, 0xa7,
	0x8b, 0xe4, 0xbd, 0x8d, 0xe7, 0xbd, 0xae, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x34, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32, 0x12, 0xe6, 0xaf, 0x8f, 0xe9, 0xa1, 0xb5,
	0xe9, 0x99, 0x90, 0xe5, 0x88, 0xb6, 0xe6, 0x9d, 0xa1, 0xe6, 0x95, 0xb0, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c, 0xe6,
	0x8e, 0x92, 0xe5, 0xba, 0x8f, 0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x12, 0x59, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x42, 0x34, 0x92, 0x41, 0x31,
	0x32, 0x2f, 0xe6, 0x98, 0xaf, 0xe5, 0x90, 0xa6, 0xe9, 0x80, 0x86, 0xe5, 0xba, 0x8f, 0xef, 0xbc,
	0x8c, 0x20, 0x66, 0x61, 0x6c, 0x73, 0x65, 0xef, 0xbc, 0x9a, 0xe4, 0xb8, 0x8d, 0xe9, 0x80, 0x86,
	0xe5, 0xba, 0x8f, 0xef, 0xbc, 0x8c, 0x74, 0x72, 0x75, 0x65, 0x3a, 0xe9, 0x80, 0x86, 0xe5, 0xba,
	0x8f, 0x52, 0x0c, 0x69, 0x73, 0x44, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22,
	0xf4, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x19, 0x92, 0x41, 0x16, 0x32, 0x14, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32, 0x12, 0xe8, 0xae, 0xb0, 0xe5,
	0xbd, 0x95, 0xe5, 0xbc, 0x80, 0xe5, 0xa7, 0x8b, 0xe4, 0xbd, 0x8d, 0xe7, 0xbd, 0xae, 0x52, 0x07,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32,
	0x12, 0xe6, 0xaf, 0x8f, 0xe9, 0xa1, 0xb5, 0xe9, 0x99, 0x90, 0xe5, 0x88, 0xb6, 0xe6, 0x9d, 0xa1,
	0xe6, 0x95, 0xb0, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x47, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x19, 0x92, 0x41, 0x16, 0x32, 0x14, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41,
	0x0b, 0x32, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x64, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0d, 0x92, 0x41, 0x0a, 0x32, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x20,
	0x69, 0x64, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xca, 0x03, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x87, 0x01, 0x0a, 0x05, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x22, 0x06, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x03, 0x6f, 0x62, 0x6a, 0x92,
	0x41, 0x32, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x20, 0x61, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2a, 0x0b, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04,
	0x0a, 0x02, 0x4f, 0x4b, 0x12, 0x97, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c,
	0x22, 0x07, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3f, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x20, 0x62, 0x79, 0x20, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x2a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x12, 0x9b,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x2a, 0x07, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x92,
	0x41, 0x3a, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64,
	0x2a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4a,
	0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x42, 0x38, 0x0a, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2d,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_core_v1_search_proto_rawDescOnce sync.Once
	file_api_core_v1_search_proto_rawDescData = file_api_core_v1_search_proto_rawDesc
)

func file_api_core_v1_search_proto_rawDescGZIP() []byte {
	file_api_core_v1_search_proto_rawDescOnce.Do(func() {
		file_api_core_v1_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_core_v1_search_proto_rawDescData)
	})
	return file_api_core_v1_search_proto_rawDescData
}

var file_api_core_v1_search_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_core_v1_search_proto_goTypes = []interface{}{
	(*IndexObject)(nil),        // 0: api.core.v1.IndexObject
	(*IndexResponse)(nil),      // 1: api.core.v1.IndexResponse
	(*SearchCondition)(nil),    // 2: api.core.v1.SearchCondition
	(*SearchRequest)(nil),      // 3: api.core.v1.SearchRequest
	(*SearchResponse)(nil),     // 4: api.core.v1.SearchResponse
	(*DeleteByIDRequest)(nil),  // 5: api.core.v1.DeleteByIDRequest
	(*DeleteByIDResponse)(nil), // 6: api.core.v1.DeleteByIDResponse
	(*_struct.Value)(nil),      // 7: google.protobuf.Value
}
var file_api_core_v1_search_proto_depIdxs = []int32{
	7, // 0: api.core.v1.IndexObject.obj:type_name -> google.protobuf.Value
	7, // 1: api.core.v1.SearchCondition.value:type_name -> google.protobuf.Value
	2, // 2: api.core.v1.SearchRequest.condition:type_name -> api.core.v1.SearchCondition
	7, // 3: api.core.v1.SearchResponse.items:type_name -> google.protobuf.Value
	0, // 4: api.core.v1.Search.Index:input_type -> api.core.v1.IndexObject
	3, // 5: api.core.v1.Search.Search:input_type -> api.core.v1.SearchRequest
	5, // 6: api.core.v1.Search.DeleteByID:input_type -> api.core.v1.DeleteByIDRequest
	1, // 7: api.core.v1.Search.Index:output_type -> api.core.v1.IndexResponse
	4, // 8: api.core.v1.Search.Search:output_type -> api.core.v1.SearchResponse
	6, // 9: api.core.v1.Search.DeleteByID:output_type -> api.core.v1.DeleteByIDResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_core_v1_search_proto_init() }
func file_api_core_v1_search_proto_init() {
	if File_api_core_v1_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_core_v1_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexObject); i {
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
		file_api_core_v1_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexResponse); i {
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
		file_api_core_v1_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCondition); i {
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
		file_api_core_v1_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_api_core_v1_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_api_core_v1_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByIDRequest); i {
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
		file_api_core_v1_search_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByIDResponse); i {
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
			RawDescriptor: file_api_core_v1_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_core_v1_search_proto_goTypes,
		DependencyIndexes: file_api_core_v1_search_proto_depIdxs,
		MessageInfos:      file_api_core_v1_search_proto_msgTypes,
	}.Build()
	File_api_core_v1_search_proto = out.File
	file_api_core_v1_search_proto_rawDesc = nil
	file_api_core_v1_search_proto_goTypes = nil
	file_api_core_v1_search_proto_depIdxs = nil
}
