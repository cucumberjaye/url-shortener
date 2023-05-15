// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: proto/shortener.proto

package pb

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

type CommonResponse_Status int32

const (
	CommonResponse_OK   CommonResponse_Status = 0
	CommonResponse_FAIL CommonResponse_Status = 1
)

// Enum value maps for CommonResponse_Status.
var (
	CommonResponse_Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	CommonResponse_Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x CommonResponse_Status) Enum() *CommonResponse_Status {
	p := new(CommonResponse_Status)
	*p = x
	return p
}

func (x CommonResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_shortener_proto_enumTypes[0].Descriptor()
}

func (CommonResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_shortener_proto_enumTypes[0]
}

func (x CommonResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonResponse_Status.Descriptor instead.
func (CommonResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{10, 0}
}

type Short struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *Short) Reset() {
	*x = Short{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Short) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Short) ProtoMessage() {}

func (x *Short) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Short.ProtoReflect.Descriptor instead.
func (*Short) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *Short) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type Original struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *Original) Reset() {
	*x = Original{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Original) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Original) ProtoMessage() {}

func (x *Original) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Original.ProtoReflect.Descriptor instead.
func (*Original) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *Original) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type BatchOriginal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *BatchOriginal) Reset() {
	*x = BatchOriginal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOriginal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOriginal) ProtoMessage() {}

func (x *BatchOriginal) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOriginal.ProtoReflect.Descriptor instead.
func (*BatchOriginal) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *BatchOriginal) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *BatchOriginal) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type BatchShort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl      string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *BatchShort) Reset() {
	*x = BatchShort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchShort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchShort) ProtoMessage() {}

func (x *BatchShort) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchShort.ProtoReflect.Descriptor instead.
func (*BatchShort) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *BatchShort) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *BatchShort) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type BatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrls []*BatchOriginal `protobuf:"bytes,1,rep,name=original_urls,json=originalUrls,proto3" json:"original_urls,omitempty"`
}

func (x *BatchRequest) Reset() {
	*x = BatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchRequest) ProtoMessage() {}

func (x *BatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchRequest.ProtoReflect.Descriptor instead.
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *BatchRequest) GetOriginalUrls() []*BatchOriginal {
	if x != nil {
		return x.OriginalUrls
	}
	return nil
}

type BatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrls []*BatchShort `protobuf:"bytes,1,rep,name=short_urls,json=shortUrls,proto3" json:"short_urls,omitempty"`
}

func (x *BatchResponse) Reset() {
	*x = BatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponse) ProtoMessage() {}

func (x *BatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponse.ProtoReflect.Descriptor instead.
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *BatchResponse) GetShortUrls() []*BatchShort {
	if x != nil {
		return x.ShortUrls
	}
	return nil
}

type URLs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *URLs) Reset() {
	*x = URLs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLs) ProtoMessage() {}

func (x *URLs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLs.ProtoReflect.Descriptor instead.
func (*URLs) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *URLs) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *URLs) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type URLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*URLs `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *URLsResponse) Reset() {
	*x = URLsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLsResponse) ProtoMessage() {}

func (x *URLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLsResponse.ProtoReflect.Descriptor instead.
func (*URLsResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *URLsResponse) GetUrls() []*URLs {
	if x != nil {
		return x.Urls
	}
	return nil
}

type BatchDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrls []*Short `protobuf:"bytes,1,rep,name=short_urls,json=shortUrls,proto3" json:"short_urls,omitempty"`
}

func (x *BatchDelete) Reset() {
	*x = BatchDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDelete) ProtoMessage() {}

func (x *BatchDelete) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDelete.ProtoReflect.Descriptor instead.
func (*BatchDelete) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *BatchDelete) GetShortUrls() []*Short {
	if x != nil {
		return x.ShortUrls
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{9}
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status CommonResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=internal.app.pb.CommonResponse_Status" json:"status,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *CommonResponse) GetStatus() CommonResponse_Status {
	if x != nil {
		return x.Status
	}
	return CommonResponse_OK
}

type StatsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlsCount  int32 `protobuf:"varint,1,opt,name=urls_count,json=urlsCount,proto3" json:"urls_count,omitempty"`
	UsersCount int32 `protobuf:"varint,2,opt,name=users_count,json=usersCount,proto3" json:"users_count,omitempty"`
}

func (x *StatsInfo) Reset() {
	*x = StatsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsInfo) ProtoMessage() {}

func (x *StatsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsInfo.ProtoReflect.Descriptor instead.
func (*StatsInfo) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{11}
}

func (x *StatsInfo) GetUrlsCount() int32 {
	if x != nil {
		return x.UrlsCount
	}
	return 0
}

func (x *StatsInfo) GetUsersCount() int32 {
	if x != nil {
		return x.UsersCount
	}
	return 0
}

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{12}
}

func (x *AuthToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_proto_shortener_proto protoreflect.FileDescriptor

var file_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x22, 0x24, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x2d,
	0x0a, 0x08, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x59, 0x0a,
	0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x50, 0x0a, 0x0a, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x53, 0x0a, 0x0c, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0d, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x73, 0x22,
	0x4b, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x46, 0x0a, 0x04,
	0x55, 0x52, 0x4c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x72, 0x6c, 0x22, 0x39, 0x0a, 0x0c, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22,
	0x44, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x35,
	0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6c,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x26, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x22, 0x4b, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x72, 0x6c,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75,
	0x72, 0x6c, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x21, 0x0a, 0x09, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xbd, 0x04, 0x0a,
	0x0f, 0x53, 0x68, 0x6f, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c,
	0x6c, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x1a, 0x19, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x1a,
	0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c,
	0x12, 0x1c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x1f,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_shortener_proto_rawDescOnce sync.Once
	file_proto_shortener_proto_rawDescData = file_proto_shortener_proto_rawDesc
)

func file_proto_shortener_proto_rawDescGZIP() []byte {
	file_proto_shortener_proto_rawDescOnce.Do(func() {
		file_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shortener_proto_rawDescData)
	})
	return file_proto_shortener_proto_rawDescData
}

var file_proto_shortener_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_shortener_proto_goTypes = []interface{}{
	(CommonResponse_Status)(0), // 0: internal.app.pb.CommonResponse.Status
	(*Short)(nil),              // 1: internal.app.pb.Short
	(*Original)(nil),           // 2: internal.app.pb.Original
	(*BatchOriginal)(nil),      // 3: internal.app.pb.BatchOriginal
	(*BatchShort)(nil),         // 4: internal.app.pb.BatchShort
	(*BatchRequest)(nil),       // 5: internal.app.pb.BatchRequest
	(*BatchResponse)(nil),      // 6: internal.app.pb.BatchResponse
	(*URLs)(nil),               // 7: internal.app.pb.URLs
	(*URLsResponse)(nil),       // 8: internal.app.pb.URLsResponse
	(*BatchDelete)(nil),        // 9: internal.app.pb.BatchDelete
	(*Empty)(nil),              // 10: internal.app.pb.Empty
	(*CommonResponse)(nil),     // 11: internal.app.pb.CommonResponse
	(*StatsInfo)(nil),          // 12: internal.app.pb.StatsInfo
	(*AuthToken)(nil),          // 13: internal.app.pb.AuthToken
}
var file_proto_shortener_proto_depIdxs = []int32{
	3,  // 0: internal.app.pb.BatchRequest.original_urls:type_name -> internal.app.pb.BatchOriginal
	4,  // 1: internal.app.pb.BatchResponse.short_urls:type_name -> internal.app.pb.BatchShort
	7,  // 2: internal.app.pb.URLsResponse.urls:type_name -> internal.app.pb.URLs
	1,  // 3: internal.app.pb.BatchDelete.short_urls:type_name -> internal.app.pb.Short
	0,  // 4: internal.app.pb.CommonResponse.status:type_name -> internal.app.pb.CommonResponse.Status
	10, // 5: internal.app.pb.ShotenerService.Authentication:input_type -> internal.app.pb.Empty
	1,  // 6: internal.app.pb.ShotenerService.GetFullURL:input_type -> internal.app.pb.Short
	2,  // 7: internal.app.pb.ShotenerService.Shortener:input_type -> internal.app.pb.Original
	10, // 8: internal.app.pb.ShotenerService.Ping:input_type -> internal.app.pb.Empty
	5,  // 9: internal.app.pb.ShotenerService.BatchShortnener:input_type -> internal.app.pb.BatchRequest
	10, // 10: internal.app.pb.ShotenerService.GetUserURL:input_type -> internal.app.pb.Empty
	9,  // 11: internal.app.pb.ShotenerService.DeleteUserURL:input_type -> internal.app.pb.BatchDelete
	10, // 12: internal.app.pb.ShotenerService.Stats:input_type -> internal.app.pb.Empty
	13, // 13: internal.app.pb.ShotenerService.Authentication:output_type -> internal.app.pb.AuthToken
	2,  // 14: internal.app.pb.ShotenerService.GetFullURL:output_type -> internal.app.pb.Original
	1,  // 15: internal.app.pb.ShotenerService.Shortener:output_type -> internal.app.pb.Short
	11, // 16: internal.app.pb.ShotenerService.Ping:output_type -> internal.app.pb.CommonResponse
	6,  // 17: internal.app.pb.ShotenerService.BatchShortnener:output_type -> internal.app.pb.BatchResponse
	8,  // 18: internal.app.pb.ShotenerService.GetUserURL:output_type -> internal.app.pb.URLsResponse
	11, // 19: internal.app.pb.ShotenerService.DeleteUserURL:output_type -> internal.app.pb.CommonResponse
	12, // 20: internal.app.pb.ShotenerService.Stats:output_type -> internal.app.pb.StatsInfo
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_shortener_proto_init() }
func file_proto_shortener_proto_init() {
	if File_proto_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Short); i {
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
		file_proto_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Original); i {
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
		file_proto_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOriginal); i {
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
		file_proto_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchShort); i {
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
		file_proto_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchRequest); i {
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
		file_proto_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchResponse); i {
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
		file_proto_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLs); i {
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
		file_proto_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLsResponse); i {
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
		file_proto_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDelete); i {
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
		file_proto_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_shortener_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
		file_proto_shortener_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsInfo); i {
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
		file_proto_shortener_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
			RawDescriptor: file_proto_shortener_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shortener_proto_goTypes,
		DependencyIndexes: file_proto_shortener_proto_depIdxs,
		EnumInfos:         file_proto_shortener_proto_enumTypes,
		MessageInfos:      file_proto_shortener_proto_msgTypes,
	}.Build()
	File_proto_shortener_proto = out.File
	file_proto_shortener_proto_rawDesc = nil
	file_proto_shortener_proto_goTypes = nil
	file_proto_shortener_proto_depIdxs = nil
}
