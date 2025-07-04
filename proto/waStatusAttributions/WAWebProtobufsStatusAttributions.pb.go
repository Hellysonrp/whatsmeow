// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: waStatusAttributions/WAWebProtobufsStatusAttributions.proto

package waStatusAttributions

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusAttribution_Type int32

const (
	StatusAttribution_RESHARE        StatusAttribution_Type = 0
	StatusAttribution_EXTERNAL_SHARE StatusAttribution_Type = 1
	StatusAttribution_MUSIC          StatusAttribution_Type = 2
	StatusAttribution_STATUS_MENTION StatusAttribution_Type = 3
	StatusAttribution_GROUP_STATUS   StatusAttribution_Type = 4
)

// Enum value maps for StatusAttribution_Type.
var (
	StatusAttribution_Type_name = map[int32]string{
		0: "RESHARE",
		1: "EXTERNAL_SHARE",
		2: "MUSIC",
		3: "STATUS_MENTION",
		4: "GROUP_STATUS",
	}
	StatusAttribution_Type_value = map[string]int32{
		"RESHARE":        0,
		"EXTERNAL_SHARE": 1,
		"MUSIC":          2,
		"STATUS_MENTION": 3,
		"GROUP_STATUS":   4,
	}
)

func (x StatusAttribution_Type) Enum() *StatusAttribution_Type {
	p := new(StatusAttribution_Type)
	*p = x
	return p
}

func (x StatusAttribution_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusAttribution_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_enumTypes[0].Descriptor()
}

func (StatusAttribution_Type) Type() protoreflect.EnumType {
	return &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_enumTypes[0]
}

func (x StatusAttribution_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *StatusAttribution_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = StatusAttribution_Type(num)
	return nil
}

// Deprecated: Use StatusAttribution_Type.Descriptor instead.
func (StatusAttribution_Type) EnumDescriptor() ([]byte, []int) {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP(), []int{0, 0}
}

type StatusAttribution_ExternalShare_Source int32

const (
	StatusAttribution_ExternalShare_UNKNOWN   StatusAttribution_ExternalShare_Source = 0
	StatusAttribution_ExternalShare_INSTAGRAM StatusAttribution_ExternalShare_Source = 1
	StatusAttribution_ExternalShare_FACEBOOK  StatusAttribution_ExternalShare_Source = 2
	StatusAttribution_ExternalShare_MESSENGER StatusAttribution_ExternalShare_Source = 3
	StatusAttribution_ExternalShare_SPOTIFY   StatusAttribution_ExternalShare_Source = 4
	StatusAttribution_ExternalShare_YOUTUBE   StatusAttribution_ExternalShare_Source = 5
	StatusAttribution_ExternalShare_PINTEREST StatusAttribution_ExternalShare_Source = 6
)

// Enum value maps for StatusAttribution_ExternalShare_Source.
var (
	StatusAttribution_ExternalShare_Source_name = map[int32]string{
		0: "UNKNOWN",
		1: "INSTAGRAM",
		2: "FACEBOOK",
		3: "MESSENGER",
		4: "SPOTIFY",
		5: "YOUTUBE",
		6: "PINTEREST",
	}
	StatusAttribution_ExternalShare_Source_value = map[string]int32{
		"UNKNOWN":   0,
		"INSTAGRAM": 1,
		"FACEBOOK":  2,
		"MESSENGER": 3,
		"SPOTIFY":   4,
		"YOUTUBE":   5,
		"PINTEREST": 6,
	}
)

func (x StatusAttribution_ExternalShare_Source) Enum() *StatusAttribution_ExternalShare_Source {
	p := new(StatusAttribution_ExternalShare_Source)
	*p = x
	return p
}

func (x StatusAttribution_ExternalShare_Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusAttribution_ExternalShare_Source) Descriptor() protoreflect.EnumDescriptor {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_enumTypes[1].Descriptor()
}

func (StatusAttribution_ExternalShare_Source) Type() protoreflect.EnumType {
	return &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_enumTypes[1]
}

func (x StatusAttribution_ExternalShare_Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *StatusAttribution_ExternalShare_Source) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = StatusAttribution_ExternalShare_Source(num)
	return nil
}

// Deprecated: Use StatusAttribution_ExternalShare_Source.Descriptor instead.
func (StatusAttribution_ExternalShare_Source) EnumDescriptor() ([]byte, []int) {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP(), []int{0, 0, 0}
}

type StatusAttribution_StatusReshare_Source int32

const (
	StatusAttribution_StatusReshare_UNKNOWN          StatusAttribution_StatusReshare_Source = 0
	StatusAttribution_StatusReshare_INTERNAL_RESHARE StatusAttribution_StatusReshare_Source = 1
	StatusAttribution_StatusReshare_MENTION_RESHARE  StatusAttribution_StatusReshare_Source = 2
	StatusAttribution_StatusReshare_CHANNEL_RESHARE  StatusAttribution_StatusReshare_Source = 3
)

// Enum value maps for StatusAttribution_StatusReshare_Source.
var (
	StatusAttribution_StatusReshare_Source_name = map[int32]string{
		0: "UNKNOWN",
		1: "INTERNAL_RESHARE",
		2: "MENTION_RESHARE",
		3: "CHANNEL_RESHARE",
	}
	StatusAttribution_StatusReshare_Source_value = map[string]int32{
		"UNKNOWN":          0,
		"INTERNAL_RESHARE": 1,
		"MENTION_RESHARE":  2,
		"CHANNEL_RESHARE":  3,
	}
)

func (x StatusAttribution_StatusReshare_Source) Enum() *StatusAttribution_StatusReshare_Source {
	p := new(StatusAttribution_StatusReshare_Source)
	*p = x
	return p
}

func (x StatusAttribution_StatusReshare_Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusAttribution_StatusReshare_Source) Descriptor() protoreflect.EnumDescriptor {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_enumTypes[2].Descriptor()
}

func (StatusAttribution_StatusReshare_Source) Type() protoreflect.EnumType {
	return &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_enumTypes[2]
}

func (x StatusAttribution_StatusReshare_Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *StatusAttribution_StatusReshare_Source) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = StatusAttribution_StatusReshare_Source(num)
	return nil
}

// Deprecated: Use StatusAttribution_StatusReshare_Source.Descriptor instead.
func (StatusAttribution_StatusReshare_Source) EnumDescriptor() ([]byte, []int) {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP(), []int{0, 1, 0}
}

type StatusAttribution struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to AttributionData:
	//
	//	*StatusAttribution_StatusReshare_
	//	*StatusAttribution_ExternalShare_
	//	*StatusAttribution_Music_
	//	*StatusAttribution_GroupStatus_
	AttributionData isStatusAttribution_AttributionData `protobuf_oneof:"attributionData"`
	Type            *StatusAttribution_Type             `protobuf:"varint,1,opt,name=type,enum=WAWebProtobufsStatusAttributions.StatusAttribution_Type" json:"type,omitempty"`
	ActionURL       *string                             `protobuf:"bytes,2,opt,name=actionURL" json:"actionURL,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *StatusAttribution) Reset() {
	*x = StatusAttribution{}
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusAttribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAttribution) ProtoMessage() {}

func (x *StatusAttribution) ProtoReflect() protoreflect.Message {
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAttribution.ProtoReflect.Descriptor instead.
func (*StatusAttribution) Descriptor() ([]byte, []int) {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP(), []int{0}
}

func (x *StatusAttribution) GetAttributionData() isStatusAttribution_AttributionData {
	if x != nil {
		return x.AttributionData
	}
	return nil
}

func (x *StatusAttribution) GetStatusReshare() *StatusAttribution_StatusReshare {
	if x != nil {
		if x, ok := x.AttributionData.(*StatusAttribution_StatusReshare_); ok {
			return x.StatusReshare
		}
	}
	return nil
}

func (x *StatusAttribution) GetExternalShare() *StatusAttribution_ExternalShare {
	if x != nil {
		if x, ok := x.AttributionData.(*StatusAttribution_ExternalShare_); ok {
			return x.ExternalShare
		}
	}
	return nil
}

func (x *StatusAttribution) GetMusic() *StatusAttribution_Music {
	if x != nil {
		if x, ok := x.AttributionData.(*StatusAttribution_Music_); ok {
			return x.Music
		}
	}
	return nil
}

func (x *StatusAttribution) GetGroupStatus() *StatusAttribution_GroupStatus {
	if x != nil {
		if x, ok := x.AttributionData.(*StatusAttribution_GroupStatus_); ok {
			return x.GroupStatus
		}
	}
	return nil
}

func (x *StatusAttribution) GetType() StatusAttribution_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return StatusAttribution_RESHARE
}

func (x *StatusAttribution) GetActionURL() string {
	if x != nil && x.ActionURL != nil {
		return *x.ActionURL
	}
	return ""
}

type isStatusAttribution_AttributionData interface {
	isStatusAttribution_AttributionData()
}

type StatusAttribution_StatusReshare_ struct {
	StatusReshare *StatusAttribution_StatusReshare `protobuf:"bytes,3,opt,name=statusReshare,oneof"`
}

type StatusAttribution_ExternalShare_ struct {
	ExternalShare *StatusAttribution_ExternalShare `protobuf:"bytes,4,opt,name=externalShare,oneof"`
}

type StatusAttribution_Music_ struct {
	Music *StatusAttribution_Music `protobuf:"bytes,5,opt,name=music,oneof"`
}

type StatusAttribution_GroupStatus_ struct {
	GroupStatus *StatusAttribution_GroupStatus `protobuf:"bytes,6,opt,name=groupStatus,oneof"`
}

func (*StatusAttribution_StatusReshare_) isStatusAttribution_AttributionData() {}

func (*StatusAttribution_ExternalShare_) isStatusAttribution_AttributionData() {}

func (*StatusAttribution_Music_) isStatusAttribution_AttributionData() {}

func (*StatusAttribution_GroupStatus_) isStatusAttribution_AttributionData() {}

type StatusAttribution_ExternalShare struct {
	state             protoimpl.MessageState                  `protogen:"open.v1"`
	ActionURL         *string                                 `protobuf:"bytes,1,opt,name=actionURL" json:"actionURL,omitempty"`
	Source            *StatusAttribution_ExternalShare_Source `protobuf:"varint,2,opt,name=source,enum=WAWebProtobufsStatusAttributions.StatusAttribution_ExternalShare_Source" json:"source,omitempty"`
	Duration          *int32                                  `protobuf:"varint,3,opt,name=duration" json:"duration,omitempty"`
	ActionFallbackURL *string                                 `protobuf:"bytes,4,opt,name=actionFallbackURL" json:"actionFallbackURL,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *StatusAttribution_ExternalShare) Reset() {
	*x = StatusAttribution_ExternalShare{}
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusAttribution_ExternalShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAttribution_ExternalShare) ProtoMessage() {}

func (x *StatusAttribution_ExternalShare) ProtoReflect() protoreflect.Message {
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAttribution_ExternalShare.ProtoReflect.Descriptor instead.
func (*StatusAttribution_ExternalShare) Descriptor() ([]byte, []int) {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StatusAttribution_ExternalShare) GetActionURL() string {
	if x != nil && x.ActionURL != nil {
		return *x.ActionURL
	}
	return ""
}

func (x *StatusAttribution_ExternalShare) GetSource() StatusAttribution_ExternalShare_Source {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return StatusAttribution_ExternalShare_UNKNOWN
}

func (x *StatusAttribution_ExternalShare) GetDuration() int32 {
	if x != nil && x.Duration != nil {
		return *x.Duration
	}
	return 0
}

func (x *StatusAttribution_ExternalShare) GetActionFallbackURL() string {
	if x != nil && x.ActionFallbackURL != nil {
		return *x.ActionFallbackURL
	}
	return ""
}

type StatusAttribution_StatusReshare struct {
	state         protoimpl.MessageState                    `protogen:"open.v1"`
	Source        *StatusAttribution_StatusReshare_Source   `protobuf:"varint,1,opt,name=source,enum=WAWebProtobufsStatusAttributions.StatusAttribution_StatusReshare_Source" json:"source,omitempty"`
	Metadata      *StatusAttribution_StatusReshare_Metadata `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatusAttribution_StatusReshare) Reset() {
	*x = StatusAttribution_StatusReshare{}
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusAttribution_StatusReshare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAttribution_StatusReshare) ProtoMessage() {}

func (x *StatusAttribution_StatusReshare) ProtoReflect() protoreflect.Message {
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAttribution_StatusReshare.ProtoReflect.Descriptor instead.
func (*StatusAttribution_StatusReshare) Descriptor() ([]byte, []int) {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP(), []int{0, 1}
}

func (x *StatusAttribution_StatusReshare) GetSource() StatusAttribution_StatusReshare_Source {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return StatusAttribution_StatusReshare_UNKNOWN
}

func (x *StatusAttribution_StatusReshare) GetMetadata() *StatusAttribution_StatusReshare_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type StatusAttribution_GroupStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AuthorJID     *string                `protobuf:"bytes,1,opt,name=authorJID" json:"authorJID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatusAttribution_GroupStatus) Reset() {
	*x = StatusAttribution_GroupStatus{}
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusAttribution_GroupStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAttribution_GroupStatus) ProtoMessage() {}

func (x *StatusAttribution_GroupStatus) ProtoReflect() protoreflect.Message {
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAttribution_GroupStatus.ProtoReflect.Descriptor instead.
func (*StatusAttribution_GroupStatus) Descriptor() ([]byte, []int) {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP(), []int{0, 2}
}

func (x *StatusAttribution_GroupStatus) GetAuthorJID() string {
	if x != nil && x.AuthorJID != nil {
		return *x.AuthorJID
	}
	return ""
}

type StatusAttribution_Music struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	AuthorName        *string                `protobuf:"bytes,1,opt,name=authorName" json:"authorName,omitempty"`
	SongID            *string                `protobuf:"bytes,2,opt,name=songID" json:"songID,omitempty"`
	Title             *string                `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Author            *string                `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
	ArtistAttribution *string                `protobuf:"bytes,5,opt,name=artistAttribution" json:"artistAttribution,omitempty"`
	IsExplicit        *bool                  `protobuf:"varint,6,opt,name=isExplicit" json:"isExplicit,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *StatusAttribution_Music) Reset() {
	*x = StatusAttribution_Music{}
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusAttribution_Music) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAttribution_Music) ProtoMessage() {}

func (x *StatusAttribution_Music) ProtoReflect() protoreflect.Message {
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAttribution_Music.ProtoReflect.Descriptor instead.
func (*StatusAttribution_Music) Descriptor() ([]byte, []int) {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP(), []int{0, 3}
}

func (x *StatusAttribution_Music) GetAuthorName() string {
	if x != nil && x.AuthorName != nil {
		return *x.AuthorName
	}
	return ""
}

func (x *StatusAttribution_Music) GetSongID() string {
	if x != nil && x.SongID != nil {
		return *x.SongID
	}
	return ""
}

func (x *StatusAttribution_Music) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *StatusAttribution_Music) GetAuthor() string {
	if x != nil && x.Author != nil {
		return *x.Author
	}
	return ""
}

func (x *StatusAttribution_Music) GetArtistAttribution() string {
	if x != nil && x.ArtistAttribution != nil {
		return *x.ArtistAttribution
	}
	return ""
}

func (x *StatusAttribution_Music) GetIsExplicit() bool {
	if x != nil && x.IsExplicit != nil {
		return *x.IsExplicit
	}
	return false
}

type StatusAttribution_StatusReshare_Metadata struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Duration            *int32                 `protobuf:"varint,1,opt,name=duration" json:"duration,omitempty"`
	ChannelJID          *string                `protobuf:"bytes,2,opt,name=channelJID" json:"channelJID,omitempty"`
	ChannelMessageID    *int32                 `protobuf:"varint,3,opt,name=channelMessageID" json:"channelMessageID,omitempty"`
	HasMultipleReshares *bool                  `protobuf:"varint,4,opt,name=hasMultipleReshares" json:"hasMultipleReshares,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *StatusAttribution_StatusReshare_Metadata) Reset() {
	*x = StatusAttribution_StatusReshare_Metadata{}
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusAttribution_StatusReshare_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAttribution_StatusReshare_Metadata) ProtoMessage() {}

func (x *StatusAttribution_StatusReshare_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAttribution_StatusReshare_Metadata.ProtoReflect.Descriptor instead.
func (*StatusAttribution_StatusReshare_Metadata) Descriptor() ([]byte, []int) {
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *StatusAttribution_StatusReshare_Metadata) GetDuration() int32 {
	if x != nil && x.Duration != nil {
		return *x.Duration
	}
	return 0
}

func (x *StatusAttribution_StatusReshare_Metadata) GetChannelJID() string {
	if x != nil && x.ChannelJID != nil {
		return *x.ChannelJID
	}
	return ""
}

func (x *StatusAttribution_StatusReshare_Metadata) GetChannelMessageID() int32 {
	if x != nil && x.ChannelMessageID != nil {
		return *x.ChannelMessageID
	}
	return 0
}

func (x *StatusAttribution_StatusReshare_Metadata) GetHasMultipleReshares() bool {
	if x != nil && x.HasMultipleReshares != nil {
		return *x.HasMultipleReshares
	}
	return false
}

var File_waStatusAttributions_WAWebProtobufsStatusAttributions_proto protoreflect.FileDescriptor

const file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDesc = "" +
	"\n" +
	";waStatusAttributions/WAWebProtobufsStatusAttributions.proto\x12 WAWebProtobufsStatusAttributions\"\x87\r\n" +
	"\x11StatusAttribution\x12i\n" +
	"\rstatusReshare\x18\x03 \x01(\v2A.WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshareH\x00R\rstatusReshare\x12i\n" +
	"\rexternalShare\x18\x04 \x01(\v2A.WAWebProtobufsStatusAttributions.StatusAttribution.ExternalShareH\x00R\rexternalShare\x12Q\n" +
	"\x05music\x18\x05 \x01(\v29.WAWebProtobufsStatusAttributions.StatusAttribution.MusicH\x00R\x05music\x12c\n" +
	"\vgroupStatus\x18\x06 \x01(\v2?.WAWebProtobufsStatusAttributions.StatusAttribution.GroupStatusH\x00R\vgroupStatus\x12L\n" +
	"\x04type\x18\x01 \x01(\x0e28.WAWebProtobufsStatusAttributions.StatusAttribution.TypeR\x04type\x12\x1c\n" +
	"\tactionURL\x18\x02 \x01(\tR\tactionURL\x1a\xc5\x02\n" +
	"\rExternalShare\x12\x1c\n" +
	"\tactionURL\x18\x01 \x01(\tR\tactionURL\x12`\n" +
	"\x06source\x18\x02 \x01(\x0e2H.WAWebProtobufsStatusAttributions.StatusAttribution.ExternalShare.SourceR\x06source\x12\x1a\n" +
	"\bduration\x18\x03 \x01(\x05R\bduration\x12,\n" +
	"\x11actionFallbackURL\x18\x04 \x01(\tR\x11actionFallbackURL\"j\n" +
	"\x06Source\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\r\n" +
	"\tINSTAGRAM\x10\x01\x12\f\n" +
	"\bFACEBOOK\x10\x02\x12\r\n" +
	"\tMESSENGER\x10\x03\x12\v\n" +
	"\aSPOTIFY\x10\x04\x12\v\n" +
	"\aYOUTUBE\x10\x05\x12\r\n" +
	"\tPINTEREST\x10\x06\x1a\xd7\x03\n" +
	"\rStatusReshare\x12`\n" +
	"\x06source\x18\x01 \x01(\x0e2H.WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare.SourceR\x06source\x12f\n" +
	"\bmetadata\x18\x02 \x01(\v2J.WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare.MetadataR\bmetadata\x1a\xa4\x01\n" +
	"\bMetadata\x12\x1a\n" +
	"\bduration\x18\x01 \x01(\x05R\bduration\x12\x1e\n" +
	"\n" +
	"channelJID\x18\x02 \x01(\tR\n" +
	"channelJID\x12*\n" +
	"\x10channelMessageID\x18\x03 \x01(\x05R\x10channelMessageID\x120\n" +
	"\x13hasMultipleReshares\x18\x04 \x01(\bR\x13hasMultipleReshares\"U\n" +
	"\x06Source\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\x14\n" +
	"\x10INTERNAL_RESHARE\x10\x01\x12\x13\n" +
	"\x0fMENTION_RESHARE\x10\x02\x12\x13\n" +
	"\x0fCHANNEL_RESHARE\x10\x03\x1a+\n" +
	"\vGroupStatus\x12\x1c\n" +
	"\tauthorJID\x18\x01 \x01(\tR\tauthorJID\x1a\xbb\x01\n" +
	"\x05Music\x12\x1e\n" +
	"\n" +
	"authorName\x18\x01 \x01(\tR\n" +
	"authorName\x12\x16\n" +
	"\x06songID\x18\x02 \x01(\tR\x06songID\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x16\n" +
	"\x06author\x18\x04 \x01(\tR\x06author\x12,\n" +
	"\x11artistAttribution\x18\x05 \x01(\tR\x11artistAttribution\x12\x1e\n" +
	"\n" +
	"isExplicit\x18\x06 \x01(\bR\n" +
	"isExplicit\"X\n" +
	"\x04Type\x12\v\n" +
	"\aRESHARE\x10\x00\x12\x12\n" +
	"\x0eEXTERNAL_SHARE\x10\x01\x12\t\n" +
	"\x05MUSIC\x10\x02\x12\x12\n" +
	"\x0eSTATUS_MENTION\x10\x03\x12\x10\n" +
	"\fGROUP_STATUS\x10\x04B\x11\n" +
	"\x0fattributionDataB0Z.go.mau.fi/whatsmeow/proto/waStatusAttributions"

var (
	file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescOnce sync.Once
	file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescData []byte
)

func file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescGZIP() []byte {
	file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescOnce.Do(func() {
		file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDesc), len(file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDesc)))
	})
	return file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDescData
}

var file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_goTypes = []any{
	(StatusAttribution_Type)(0),                      // 0: WAWebProtobufsStatusAttributions.StatusAttribution.Type
	(StatusAttribution_ExternalShare_Source)(0),      // 1: WAWebProtobufsStatusAttributions.StatusAttribution.ExternalShare.Source
	(StatusAttribution_StatusReshare_Source)(0),      // 2: WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare.Source
	(*StatusAttribution)(nil),                        // 3: WAWebProtobufsStatusAttributions.StatusAttribution
	(*StatusAttribution_ExternalShare)(nil),          // 4: WAWebProtobufsStatusAttributions.StatusAttribution.ExternalShare
	(*StatusAttribution_StatusReshare)(nil),          // 5: WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare
	(*StatusAttribution_GroupStatus)(nil),            // 6: WAWebProtobufsStatusAttributions.StatusAttribution.GroupStatus
	(*StatusAttribution_Music)(nil),                  // 7: WAWebProtobufsStatusAttributions.StatusAttribution.Music
	(*StatusAttribution_StatusReshare_Metadata)(nil), // 8: WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare.Metadata
}
var file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_depIdxs = []int32{
	5, // 0: WAWebProtobufsStatusAttributions.StatusAttribution.statusReshare:type_name -> WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare
	4, // 1: WAWebProtobufsStatusAttributions.StatusAttribution.externalShare:type_name -> WAWebProtobufsStatusAttributions.StatusAttribution.ExternalShare
	7, // 2: WAWebProtobufsStatusAttributions.StatusAttribution.music:type_name -> WAWebProtobufsStatusAttributions.StatusAttribution.Music
	6, // 3: WAWebProtobufsStatusAttributions.StatusAttribution.groupStatus:type_name -> WAWebProtobufsStatusAttributions.StatusAttribution.GroupStatus
	0, // 4: WAWebProtobufsStatusAttributions.StatusAttribution.type:type_name -> WAWebProtobufsStatusAttributions.StatusAttribution.Type
	1, // 5: WAWebProtobufsStatusAttributions.StatusAttribution.ExternalShare.source:type_name -> WAWebProtobufsStatusAttributions.StatusAttribution.ExternalShare.Source
	2, // 6: WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare.source:type_name -> WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare.Source
	8, // 7: WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare.metadata:type_name -> WAWebProtobufsStatusAttributions.StatusAttribution.StatusReshare.Metadata
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_init() }
func file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_init() {
	if File_waStatusAttributions_WAWebProtobufsStatusAttributions_proto != nil {
		return
	}
	file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes[0].OneofWrappers = []any{
		(*StatusAttribution_StatusReshare_)(nil),
		(*StatusAttribution_ExternalShare_)(nil),
		(*StatusAttribution_Music_)(nil),
		(*StatusAttribution_GroupStatus_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDesc), len(file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_goTypes,
		DependencyIndexes: file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_depIdxs,
		EnumInfos:         file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_enumTypes,
		MessageInfos:      file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_msgTypes,
	}.Build()
	File_waStatusAttributions_WAWebProtobufsStatusAttributions_proto = out.File
	file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_goTypes = nil
	file_waStatusAttributions_WAWebProtobufsStatusAttributions_proto_depIdxs = nil
}
