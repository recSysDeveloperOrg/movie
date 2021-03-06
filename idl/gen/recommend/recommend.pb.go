// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: idl/recommend.proto

package recommend

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

type RecommendSourceType int32

const (
	RecommendSourceType_RECOMMEND_SOURCE_TYPE_RATING RecommendSourceType = 0
	RecommendSourceType_RECOMMEND_SOURCE_TYPE_TAG    RecommendSourceType = 1
	RecommendSourceType_RECOMMEND_SOURCE_TYPE_LOG    RecommendSourceType = 2
	RecommendSourceType_RECOMMEND_SOURCE_TYPE_TOP_K  RecommendSourceType = 3
)

// Enum value maps for RecommendSourceType.
var (
	RecommendSourceType_name = map[int32]string{
		0: "RECOMMEND_SOURCE_TYPE_RATING",
		1: "RECOMMEND_SOURCE_TYPE_TAG",
		2: "RECOMMEND_SOURCE_TYPE_LOG",
		3: "RECOMMEND_SOURCE_TYPE_TOP_K",
	}
	RecommendSourceType_value = map[string]int32{
		"RECOMMEND_SOURCE_TYPE_RATING": 0,
		"RECOMMEND_SOURCE_TYPE_TAG":    1,
		"RECOMMEND_SOURCE_TYPE_LOG":    2,
		"RECOMMEND_SOURCE_TYPE_TOP_K":  3,
	}
)

func (x RecommendSourceType) Enum() *RecommendSourceType {
	p := new(RecommendSourceType)
	*p = x
	return p
}

func (x RecommendSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecommendSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_recommend_proto_enumTypes[0].Descriptor()
}

func (RecommendSourceType) Type() protoreflect.EnumType {
	return &file_idl_recommend_proto_enumTypes[0]
}

func (x RecommendSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecommendSourceType.Descriptor instead.
func (RecommendSourceType) EnumDescriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{0}
}

type FilterType int32

const (
	FilterType_FILTER_TYPE_MOVIE FilterType = 0
	FilterType_FILTER_TYPE_TAG   FilterType = 1
)

// Enum value maps for FilterType.
var (
	FilterType_name = map[int32]string{
		0: "FILTER_TYPE_MOVIE",
		1: "FILTER_TYPE_TAG",
	}
	FilterType_value = map[string]int32{
		"FILTER_TYPE_MOVIE": 0,
		"FILTER_TYPE_TAG":   1,
	}
)

func (x FilterType) Enum() *FilterType {
	p := new(FilterType)
	*p = x
	return p
}

func (x FilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_recommend_proto_enumTypes[1].Descriptor()
}

func (FilterType) Type() protoreflect.EnumType {
	return &file_idl_recommend_proto_enumTypes[1]
}

func (x FilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterType.Descriptor instead.
func (FilterType) EnumDescriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{1}
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_recommend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_recommend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaseResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RecommendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page   int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *RecommendReq) Reset() {
	*x = RecommendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_recommend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendReq) ProtoMessage() {}

func (x *RecommendReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_recommend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendReq.ProtoReflect.Descriptor instead.
func (*RecommendReq) Descriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{1}
}

func (x *RecommendReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RecommendReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *RecommendReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type RecommendEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId  string              `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	RsType   RecommendSourceType `protobuf:"varint,2,opt,name=rs_type,json=rsType,proto3,enum=recommend.RecommendSourceType" json:"rs_type,omitempty"`
	SourceId string              `protobuf:"bytes,3,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"` // ??????RATING???LOG?????????movie_id,??????TAG?????????tag_id,TOP_K??????
}

func (x *RecommendEntry) Reset() {
	*x = RecommendEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_recommend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendEntry) ProtoMessage() {}

func (x *RecommendEntry) ProtoReflect() protoreflect.Message {
	mi := &file_idl_recommend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendEntry.ProtoReflect.Descriptor instead.
func (*RecommendEntry) Descriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{2}
}

func (x *RecommendEntry) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

func (x *RecommendEntry) GetRsType() RecommendSourceType {
	if x != nil {
		return x.RsType
	}
	return RecommendSourceType_RECOMMEND_SOURCE_TYPE_RATING
}

func (x *RecommendEntry) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

type RecommendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp   *BaseResp         `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	Entry      []*RecommendEntry `protobuf:"bytes,2,rep,name=entry,proto3" json:"entry,omitempty"`
	NRecommend int64             `protobuf:"varint,3,opt,name=n_recommend,json=nRecommend,proto3" json:"n_recommend,omitempty"`
}

func (x *RecommendResp) Reset() {
	*x = RecommendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_recommend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendResp) ProtoMessage() {}

func (x *RecommendResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_recommend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendResp.ProtoReflect.Descriptor instead.
func (*RecommendResp) Descriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{3}
}

func (x *RecommendResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *RecommendResp) GetEntry() []*RecommendEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *RecommendResp) GetNRecommend() int64 {
	if x != nil {
		return x.NRecommend
	}
	return 0
}

type FilterRuleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FType    FilterType `protobuf:"varint,1,opt,name=f_type,json=fType,proto3,enum=recommend.FilterType" json:"f_type,omitempty"`
	SourceId string     `protobuf:"bytes,2,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	UserId   string     `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FilterRuleReq) Reset() {
	*x = FilterRuleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_recommend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterRuleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterRuleReq) ProtoMessage() {}

func (x *FilterRuleReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_recommend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterRuleReq.ProtoReflect.Descriptor instead.
func (*FilterRuleReq) Descriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{4}
}

func (x *FilterRuleReq) GetFType() FilterType {
	if x != nil {
		return x.FType
	}
	return FilterType_FILTER_TYPE_MOVIE
}

func (x *FilterRuleReq) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *FilterRuleReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FilterRuleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *FilterRuleResp) Reset() {
	*x = FilterRuleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_recommend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterRuleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterRuleResp) ProtoMessage() {}

func (x *FilterRuleResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_recommend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterRuleResp.ProtoReflect.Descriptor instead.
func (*FilterRuleResp) Descriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{5}
}

func (x *FilterRuleResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type ViewLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieId string `protobuf:"bytes,2,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
}

func (x *ViewLogReq) Reset() {
	*x = ViewLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_recommend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewLogReq) ProtoMessage() {}

func (x *ViewLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_recommend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewLogReq.ProtoReflect.Descriptor instead.
func (*ViewLogReq) Descriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{6}
}

func (x *ViewLogReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ViewLogReq) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

type ViewLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *ViewLogResp) Reset() {
	*x = ViewLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_recommend_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewLogResp) ProtoMessage() {}

func (x *ViewLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_recommend_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewLogResp.ProtoReflect.Descriptor instead.
func (*ViewLogResp) Descriptor() ([]byte, []int) {
	return file_idl_recommend_proto_rawDescGZIP(), []int{7}
}

func (x *ViewLogResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_idl_recommend_proto protoreflect.FileDescriptor

var file_idl_recommend_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x64, 0x6c, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x22, 0x30, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x53, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x07, 0x72, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x72, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x0d,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a,
	0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2f, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x22, 0x73, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x66, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x40, 0x0a, 0x0a, 0x56, 0x69,
	0x65, 0x77, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x0b,
	0x56, 0x69, 0x65, 0x77, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x2a, 0x96, 0x01,
	0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45,
	0x4e, 0x44, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x43, 0x4f, 0x4d,
	0x4d, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x41, 0x47, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d,
	0x45, 0x4e, 0x44, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4c, 0x4f, 0x47, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45,
	0x4e, 0x44, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x4f, 0x50, 0x5f, 0x4b, 0x10, 0x03, 0x2a, 0x38, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x56, 0x49, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x01,
	0x32, 0xd6, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x40, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x56, 0x69, 0x65, 0x77, 0x4c, 0x6f, 0x67, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x56, 0x69, 0x65, 0x77,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x69,
	0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_recommend_proto_rawDescOnce sync.Once
	file_idl_recommend_proto_rawDescData = file_idl_recommend_proto_rawDesc
)

func file_idl_recommend_proto_rawDescGZIP() []byte {
	file_idl_recommend_proto_rawDescOnce.Do(func() {
		file_idl_recommend_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_recommend_proto_rawDescData)
	})
	return file_idl_recommend_proto_rawDescData
}

var file_idl_recommend_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_idl_recommend_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_idl_recommend_proto_goTypes = []interface{}{
	(RecommendSourceType)(0), // 0: recommend.RecommendSourceType
	(FilterType)(0),          // 1: recommend.FilterType
	(*BaseResp)(nil),         // 2: recommend.BaseResp
	(*RecommendReq)(nil),     // 3: recommend.RecommendReq
	(*RecommendEntry)(nil),   // 4: recommend.RecommendEntry
	(*RecommendResp)(nil),    // 5: recommend.RecommendResp
	(*FilterRuleReq)(nil),    // 6: recommend.FilterRuleReq
	(*FilterRuleResp)(nil),   // 7: recommend.FilterRuleResp
	(*ViewLogReq)(nil),       // 8: recommend.ViewLogReq
	(*ViewLogResp)(nil),      // 9: recommend.ViewLogResp
}
var file_idl_recommend_proto_depIdxs = []int32{
	0, // 0: recommend.RecommendEntry.rs_type:type_name -> recommend.RecommendSourceType
	2, // 1: recommend.RecommendResp.base_resp:type_name -> recommend.BaseResp
	4, // 2: recommend.RecommendResp.entry:type_name -> recommend.RecommendEntry
	1, // 3: recommend.FilterRuleReq.f_type:type_name -> recommend.FilterType
	2, // 4: recommend.FilterRuleResp.base_resp:type_name -> recommend.BaseResp
	2, // 5: recommend.ViewLogResp.base_resp:type_name -> recommend.BaseResp
	3, // 6: recommend.Recommender.Recommend:input_type -> recommend.RecommendReq
	6, // 7: recommend.Recommender.AddFilterRule:input_type -> recommend.FilterRuleReq
	8, // 8: recommend.Recommender.AddViewLog:input_type -> recommend.ViewLogReq
	5, // 9: recommend.Recommender.Recommend:output_type -> recommend.RecommendResp
	7, // 10: recommend.Recommender.AddFilterRule:output_type -> recommend.FilterRuleResp
	9, // 11: recommend.Recommender.AddViewLog:output_type -> recommend.ViewLogResp
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_idl_recommend_proto_init() }
func file_idl_recommend_proto_init() {
	if File_idl_recommend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_recommend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_idl_recommend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendReq); i {
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
		file_idl_recommend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendEntry); i {
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
		file_idl_recommend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendResp); i {
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
		file_idl_recommend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterRuleReq); i {
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
		file_idl_recommend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterRuleResp); i {
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
		file_idl_recommend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewLogReq); i {
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
		file_idl_recommend_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewLogResp); i {
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
			RawDescriptor: file_idl_recommend_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_recommend_proto_goTypes,
		DependencyIndexes: file_idl_recommend_proto_depIdxs,
		EnumInfos:         file_idl_recommend_proto_enumTypes,
		MessageInfos:      file_idl_recommend_proto_msgTypes,
	}.Build()
	File_idl_recommend_proto = out.File
	file_idl_recommend_proto_rawDesc = nil
	file_idl_recommend_proto_goTypes = nil
	file_idl_recommend_proto_depIdxs = nil
}
