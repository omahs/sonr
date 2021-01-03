// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: data.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MIME_Type int32

const (
	MIME_application MIME_Type = 0
	MIME_audio       MIME_Type = 1
	MIME_image       MIME_Type = 2
	MIME_text        MIME_Type = 3
	MIME_video       MIME_Type = 4
)

// Enum value maps for MIME_Type.
var (
	MIME_Type_name = map[int32]string{
		0: "application",
		1: "audio",
		2: "image",
		3: "text",
		4: "video",
	}
	MIME_Type_value = map[string]int32{
		"application": 0,
		"audio":       1,
		"image":       2,
		"text":        3,
		"video":       4,
	}
)

func (x MIME_Type) Enum() *MIME_Type {
	p := new(MIME_Type)
	*p = x
	return p
}

func (x MIME_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MIME_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[0].Descriptor()
}

func (MIME_Type) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[0]
}

func (x MIME_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MIME_Type.Descriptor instead.
func (MIME_Type) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2, 0}
}

type Contact_SocialTile_Provider int32

const (
	Contact_SocialTile_Instagram Contact_SocialTile_Provider = 0 // Feed/Link Option
	Contact_SocialTile_TikTok    Contact_SocialTile_Provider = 1 // Feed/Link Option
	Contact_SocialTile_Facebook  Contact_SocialTile_Provider = 2 // Link Only
	Contact_SocialTile_YouTube   Contact_SocialTile_Provider = 3 // Feed/Link Option
	Contact_SocialTile_Spotify   Contact_SocialTile_Provider = 4 // Feed/Link Option
	Contact_SocialTile_Medium    Contact_SocialTile_Provider = 5 // Feed/Link Option
	Contact_SocialTile_Twitter   Contact_SocialTile_Provider = 6 // Feed/Link Option
	Contact_SocialTile_Snapchat  Contact_SocialTile_Provider = 7 // Feed/Link Option
	Contact_SocialTile_Github    Contact_SocialTile_Provider = 8 // Feed/Link Option
)

// Enum value maps for Contact_SocialTile_Provider.
var (
	Contact_SocialTile_Provider_name = map[int32]string{
		0: "Instagram",
		1: "TikTok",
		2: "Facebook",
		3: "YouTube",
		4: "Spotify",
		5: "Medium",
		6: "Twitter",
		7: "Snapchat",
		8: "Github",
	}
	Contact_SocialTile_Provider_value = map[string]int32{
		"Instagram": 0,
		"TikTok":    1,
		"Facebook":  2,
		"YouTube":   3,
		"Spotify":   4,
		"Medium":    5,
		"Twitter":   6,
		"Snapchat":  7,
		"Github":    8,
	}
)

func (x Contact_SocialTile_Provider) Enum() *Contact_SocialTile_Provider {
	p := new(Contact_SocialTile_Provider)
	*p = x
	return p
}

func (x Contact_SocialTile_Provider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Contact_SocialTile_Provider) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[1].Descriptor()
}

func (Contact_SocialTile_Provider) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[1]
}

func (x Contact_SocialTile_Provider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Contact_SocialTile_Provider.Descriptor instead.
func (Contact_SocialTile_Provider) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4, 0, 0}
}

type Peer_Proximity int32

const (
	Peer_NONE      Peer_Proximity = 0
	Peer_IMMEDIATE Peer_Proximity = 1
	Peer_NEAR      Peer_Proximity = 2
	Peer_FAR       Peer_Proximity = 3
)

// Enum value maps for Peer_Proximity.
var (
	Peer_Proximity_name = map[int32]string{
		0: "NONE",
		1: "IMMEDIATE",
		2: "NEAR",
		3: "FAR",
	}
	Peer_Proximity_value = map[string]int32{
		"NONE":      0,
		"IMMEDIATE": 1,
		"NEAR":      2,
		"FAR":       3,
	}
)

func (x Peer_Proximity) Enum() *Peer_Proximity {
	p := new(Peer_Proximity)
	*p = x
	return p
}

func (x Peer_Proximity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Peer_Proximity) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[2].Descriptor()
}

func (Peer_Proximity) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[2]
}

func (x Peer_Proximity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Peer_Proximity.Descriptor instead.
func (Peer_Proximity) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{7, 0}
}

// Define Metadata Type: For Received Transfer
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path      string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Size      int32  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Mime      *MIME  `protobuf:"bytes,5,opt,name=mime,proto3" json:"mime,omitempty"`
	Thumbnail []byte `protobuf:"bytes,6,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Received  int32  `protobuf:"varint,7,opt,name=received,proto3" json:"received,omitempty"`
	Owner     *Peer  `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Metadata) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Metadata) GetMime() *MIME {
	if x != nil {
		return x.Mime
	}
	return nil
}

func (x *Metadata) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *Metadata) GetReceived() int32 {
	if x != nil {
		return x.Received
	}
	return 0
}

func (x *Metadata) GetOwner() *Peer {
	if x != nil {
		return x.Owner
	}
	return nil
}

// Define Preview: For File Transfer Request
type Preview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mime      *MIME  `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"`
	Path      string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Size      int32  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Thumbnail []byte `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *Preview) Reset() {
	*x = Preview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Preview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preview) ProtoMessage() {}

func (x *Preview) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preview.ProtoReflect.Descriptor instead.
func (*Preview) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1}
}

func (x *Preview) GetMime() *MIME {
	if x != nil {
		return x.Mime
	}
	return nil
}

func (x *Preview) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Preview) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Preview) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Preview) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

// Define MIME: Protobuf Version of Mime
type MIME struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    MIME_Type `protobuf:"varint,1,opt,name=type,proto3,enum=MIME_Type" json:"type,omitempty"`
	Subtype string    `protobuf:"bytes,2,opt,name=subtype,proto3" json:"subtype,omitempty"`
	Value   string    `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MIME) Reset() {
	*x = MIME{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MIME) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MIME) ProtoMessage() {}

func (x *MIME) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MIME.ProtoReflect.Descriptor instead.
func (*MIME) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2}
}

func (x *MIME) GetType() MIME_Type {
	if x != nil {
		return x.Type
	}
	return MIME_application
}

func (x *MIME) GetSubtype() string {
	if x != nil {
		return x.Subtype
	}
	return ""
}

func (x *MIME) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Define Lobby Type: For Info about Lobby
type Lobby struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string           `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Size  int32            `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Peers map[string]*Peer `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Lobby) Reset() {
	*x = Lobby{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lobby) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lobby) ProtoMessage() {}

func (x *Lobby) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lobby.ProtoReflect.Descriptor instead.
func (*Lobby) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{3}
}

func (x *Lobby) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Lobby) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Lobby) GetPeers() map[string]*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

// Define Contact Type: Will be CSV in Future
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Default
	FirstName  string                `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName   string                `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Phone      string                `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Website    string                `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	Email      string                `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	ProfilePic []byte                `protobuf:"bytes,6,opt,name=profilePic,proto3" json:"profilePic,omitempty"`
	Header     string                `protobuf:"bytes,7,opt,name=header,proto3" json:"header,omitempty"`
	Socials    []*Contact_SocialTile `protobuf:"bytes,8,rep,name=socials,proto3" json:"socials,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4}
}

func (x *Contact) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Contact) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Contact) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Contact) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Contact) GetProfilePic() []byte {
	if x != nil {
		return x.ProfilePic
	}
	return nil
}

func (x *Contact) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *Contact) GetSocials() []*Contact_SocialTile {
	if x != nil {
		return x.Socials
	}
	return nil
}

// Define Device Type: Information about device
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Model    string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version  string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Sdk      int32  `protobuf:"varint,5,opt,name=sdk,proto3" json:"sdk,omitempty"` // For Android
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{5}
}

func (x *Device) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Device) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Device) GetSdk() int32 {
	if x != nil {
		return x.Sdk
	}
	return 0
}

// Define Directories Type: Where Data can be stored
type Directories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Documents string `protobuf:"bytes,1,opt,name=documents,proto3" json:"documents,omitempty"`
	Library   string `protobuf:"bytes,2,opt,name=library,proto3" json:"library,omitempty"`
	Temporary string `protobuf:"bytes,3,opt,name=temporary,proto3" json:"temporary,omitempty"`
}

func (x *Directories) Reset() {
	*x = Directories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Directories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Directories) ProtoMessage() {}

func (x *Directories) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Directories.ProtoReflect.Descriptor instead.
func (*Directories) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{6}
}

func (x *Directories) GetDocuments() string {
	if x != nil {
		return x.Documents
	}
	return ""
}

func (x *Directories) GetLibrary() string {
	if x != nil {
		return x.Library
	}
	return ""
}

func (x *Directories) GetTemporary() string {
	if x != nil {
		return x.Temporary
	}
	return ""
}

// Define Peer: Basic Info Sent to Peers
type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Device     *Device        `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	Discovery  string         `protobuf:"bytes,3,opt,name=discovery,proto3" json:"discovery,omitempty"`
	Username   string         `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	FirstName  string         `protobuf:"bytes,5,opt,name=firstName,proto3" json:"firstName,omitempty"`
	ProfilePic []byte         `protobuf:"bytes,6,opt,name=profilePic,proto3" json:"profilePic,omitempty"`
	Direction  float64        `protobuf:"fixed64,7,opt,name=direction,proto3" json:"direction,omitempty"`
	Difference float64        `protobuf:"fixed64,8,opt,name=difference,proto3" json:"difference,omitempty"`
	Distance   float64        `protobuf:"fixed64,9,opt,name=distance,proto3" json:"distance,omitempty"`
	Proximity  Peer_Proximity `protobuf:"varint,10,opt,name=proximity,proto3,enum=Peer_Proximity" json:"proximity,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{7}
}

func (x *Peer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Peer) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *Peer) GetDiscovery() string {
	if x != nil {
		return x.Discovery
	}
	return ""
}

func (x *Peer) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Peer) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Peer) GetProfilePic() []byte {
	if x != nil {
		return x.ProfilePic
	}
	return nil
}

func (x *Peer) GetDirection() float64 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *Peer) GetDifference() float64 {
	if x != nil {
		return x.Difference
	}
	return 0
}

func (x *Peer) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Peer) GetProximity() Peer_Proximity {
	if x != nil {
		return x.Proximity
	}
	return Peer_NONE
}

// Extended Social Media Integration
type Contact_SocialTile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider Contact_SocialTile_Provider `protobuf:"varint,1,opt,name=provider,proto3,enum=Contact_SocialTile_Provider" json:"provider,omitempty"`
	Position int32                       `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"` // 0-7
	Username string                      `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Showcase string                      `protobuf:"bytes,4,opt,name=showcase,proto3" json:"showcase,omitempty"`
	Feed     string                      `protobuf:"bytes,5,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (x *Contact_SocialTile) Reset() {
	*x = Contact_SocialTile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact_SocialTile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact_SocialTile) ProtoMessage() {}

func (x *Contact_SocialTile) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact_SocialTile.ProtoReflect.Descriptor instead.
func (*Contact_SocialTile) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Contact_SocialTile) GetProvider() Contact_SocialTile_Provider {
	if x != nil {
		return x.Provider
	}
	return Contact_SocialTile_Instagram
}

func (x *Contact_SocialTile) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Contact_SocialTile) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Contact_SocialTile) GetShowcase() string {
	if x != nil {
		return x.Showcase
	}
	return ""
}

func (x *Contact_SocialTile) GetFeed() string {
	if x != nil {
		return x.Feed
	}
	return ""
}

var File_data_proto protoreflect.FileDescriptor

var file_data_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a,
	0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4d, 0x49, 0x4d, 0x45, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x7e, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x19, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x4d, 0x49, 0x4d, 0x45, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x22, 0x9a, 0x01, 0x0a, 0x04, 0x4d, 0x49, 0x4d, 0x45,
	0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x4d, 0x49, 0x4d, 0x45, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x42, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x10, 0x04, 0x22, 0x99, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x1a,
	0x3f, 0x0a, 0x0a, 0x50, 0x65, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xa4, 0x04, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x73, 0x1a, 0xb1, 0x02, 0x0a, 0x0a, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x54, 0x69,
	0x6c, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x65, 0x65, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x54, 0x69, 0x6b, 0x54, 0x6f, 0x6b, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x46, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x59, 0x6f,
	0x75, 0x54, 0x75, 0x62, 0x65, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x70, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x10, 0x05,
	0x12, 0x0b, 0x0a, 0x07, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x10, 0x06, 0x12, 0x0c, 0x0a,
	0x08, 0x53, 0x6e, 0x61, 0x70, 0x63, 0x68, 0x61, 0x74, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x47,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x10, 0x08, 0x22, 0x7a, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x64, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x73, 0x64, 0x6b, 0x22, 0x63, 0x0a, 0x0b, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x22, 0xf1, 0x02, 0x0a, 0x04, 0x50, 0x65, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x69,
	0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x69, 0x74,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x78, 0x69, 0x6d, 0x69, 0x74, 0x79, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d,
	0x69, 0x74, 0x79, 0x22, 0x37, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x69, 0x74, 0x79,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d,
	0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x41,
	0x52, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x41, 0x52, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_proto_rawDescOnce sync.Once
	file_data_proto_rawDescData = file_data_proto_rawDesc
)

func file_data_proto_rawDescGZIP() []byte {
	file_data_proto_rawDescOnce.Do(func() {
		file_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_proto_rawDescData)
	})
	return file_data_proto_rawDescData
}

var file_data_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_data_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_data_proto_goTypes = []interface{}{
	(MIME_Type)(0),                   // 0: MIME.Type
	(Contact_SocialTile_Provider)(0), // 1: Contact.SocialTile.Provider
	(Peer_Proximity)(0),              // 2: Peer.Proximity
	(*Metadata)(nil),                 // 3: Metadata
	(*Preview)(nil),                  // 4: Preview
	(*MIME)(nil),                     // 5: MIME
	(*Lobby)(nil),                    // 6: Lobby
	(*Contact)(nil),                  // 7: Contact
	(*Device)(nil),                   // 8: Device
	(*Directories)(nil),              // 9: Directories
	(*Peer)(nil),                     // 10: Peer
	nil,                              // 11: Lobby.PeersEntry
	(*Contact_SocialTile)(nil),       // 12: Contact.SocialTile
}
var file_data_proto_depIdxs = []int32{
	5,  // 0: Metadata.mime:type_name -> MIME
	10, // 1: Metadata.owner:type_name -> Peer
	5,  // 2: Preview.mime:type_name -> MIME
	0,  // 3: MIME.type:type_name -> MIME.Type
	11, // 4: Lobby.peers:type_name -> Lobby.PeersEntry
	12, // 5: Contact.socials:type_name -> Contact.SocialTile
	8,  // 6: Peer.device:type_name -> Device
	2,  // 7: Peer.proximity:type_name -> Peer.Proximity
	10, // 8: Lobby.PeersEntry.value:type_name -> Peer
	1,  // 9: Contact.SocialTile.provider:type_name -> Contact.SocialTile.Provider
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_data_proto_init() }
func file_data_proto_init() {
	if File_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Preview); i {
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
		file_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MIME); i {
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
		file_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lobby); i {
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
		file_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Directories); i {
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
		file_data_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
		file_data_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact_SocialTile); i {
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
			RawDescriptor: file_data_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_proto_goTypes,
		DependencyIndexes: file_data_proto_depIdxs,
		EnumInfos:         file_data_proto_enumTypes,
		MessageInfos:      file_data_proto_msgTypes,
	}.Build()
	File_data_proto = out.File
	file_data_proto_rawDesc = nil
	file_data_proto_goTypes = nil
	file_data_proto_depIdxs = nil
}
