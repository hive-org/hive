// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

package inout

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateEmailConfirmationEventV1 struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEmailConfirmationEventV1) Reset()         { *m = CreateEmailConfirmationEventV1{} }
func (m *CreateEmailConfirmationEventV1) String() string { return proto.CompactTextString(m) }
func (*CreateEmailConfirmationEventV1) ProtoMessage()    {}
func (*CreateEmailConfirmationEventV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_c19bd04ed2a6f279, []int{0}
}
func (m *CreateEmailConfirmationEventV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEmailConfirmationEventV1.Unmarshal(m, b)
}
func (m *CreateEmailConfirmationEventV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEmailConfirmationEventV1.Marshal(b, m, deterministic)
}
func (dst *CreateEmailConfirmationEventV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEmailConfirmationEventV1.Merge(dst, src)
}
func (m *CreateEmailConfirmationEventV1) XXX_Size() int {
	return xxx_messageInfo_CreateEmailConfirmationEventV1.Size(m)
}
func (m *CreateEmailConfirmationEventV1) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEmailConfirmationEventV1.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEmailConfirmationEventV1 proto.InternalMessageInfo

func (m *CreateEmailConfirmationEventV1) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateEmailConfirmationEventV1) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type CreatePhoneConfirmationEventV1 struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePhoneConfirmationEventV1) Reset()         { *m = CreatePhoneConfirmationEventV1{} }
func (m *CreatePhoneConfirmationEventV1) String() string { return proto.CompactTextString(m) }
func (*CreatePhoneConfirmationEventV1) ProtoMessage()    {}
func (*CreatePhoneConfirmationEventV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_c19bd04ed2a6f279, []int{1}
}
func (m *CreatePhoneConfirmationEventV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePhoneConfirmationEventV1.Unmarshal(m, b)
}
func (m *CreatePhoneConfirmationEventV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePhoneConfirmationEventV1.Marshal(b, m, deterministic)
}
func (dst *CreatePhoneConfirmationEventV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePhoneConfirmationEventV1.Merge(dst, src)
}
func (m *CreatePhoneConfirmationEventV1) XXX_Size() int {
	return xxx_messageInfo_CreatePhoneConfirmationEventV1.Size(m)
}
func (m *CreatePhoneConfirmationEventV1) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePhoneConfirmationEventV1.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePhoneConfirmationEventV1 proto.InternalMessageInfo

func (m *CreatePhoneConfirmationEventV1) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CreatePhoneConfirmationEventV1) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ChangedUserViewsEventV1 struct {
	Identifiers          []string `protobuf:"bytes,1,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangedUserViewsEventV1) Reset()         { *m = ChangedUserViewsEventV1{} }
func (m *ChangedUserViewsEventV1) String() string { return proto.CompactTextString(m) }
func (*ChangedUserViewsEventV1) ProtoMessage()    {}
func (*ChangedUserViewsEventV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_c19bd04ed2a6f279, []int{2}
}
func (m *ChangedUserViewsEventV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangedUserViewsEventV1.Unmarshal(m, b)
}
func (m *ChangedUserViewsEventV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangedUserViewsEventV1.Marshal(b, m, deterministic)
}
func (dst *ChangedUserViewsEventV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangedUserViewsEventV1.Merge(dst, src)
}
func (m *ChangedUserViewsEventV1) XXX_Size() int {
	return xxx_messageInfo_ChangedUserViewsEventV1.Size(m)
}
func (m *ChangedUserViewsEventV1) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangedUserViewsEventV1.DiscardUnknown(m)
}

var xxx_messageInfo_ChangedUserViewsEventV1 proto.InternalMessageInfo

func (m *ChangedUserViewsEventV1) GetIdentifiers() []string {
	if m != nil {
		return m.Identifiers
	}
	return nil
}

type Event struct {
	Sender       string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Object       string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Action       string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	EventVersion string `protobuf:"bytes,4,opt,name=event_version,json=eventVersion,proto3" json:"event_version,omitempty"`
	DataVersion  string `protobuf:"bytes,5,opt,name=data_version,json=dataVersion,proto3" json:"data_version,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*Event_PhoneConfirmationEvent
	//	*Event_EmailConfirmationEvent
	//	*Event_ChangedUserViewsEvent
	Data                 isEvent_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_c19bd04ed2a6f279, []int{3}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Event) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *Event) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Event) GetEventVersion() string {
	if m != nil {
		return m.EventVersion
	}
	return ""
}

func (m *Event) GetDataVersion() string {
	if m != nil {
		return m.DataVersion
	}
	return ""
}

type isEvent_Data interface {
	isEvent_Data()
}

type Event_PhoneConfirmationEvent struct {
	PhoneConfirmationEvent *CreatePhoneConfirmationEventV1 `protobuf:"bytes,6,opt,name=phoneConfirmationEvent,proto3,oneof"`
}

type Event_EmailConfirmationEvent struct {
	EmailConfirmationEvent *CreateEmailConfirmationEventV1 `protobuf:"bytes,7,opt,name=emailConfirmationEvent,proto3,oneof"`
}

type Event_ChangedUserViewsEvent struct {
	ChangedUserViewsEvent *ChangedUserViewsEventV1 `protobuf:"bytes,8,opt,name=changedUserViewsEvent,proto3,oneof"`
}

func (*Event_PhoneConfirmationEvent) isEvent_Data() {}

func (*Event_EmailConfirmationEvent) isEvent_Data() {}

func (*Event_ChangedUserViewsEvent) isEvent_Data() {}

func (m *Event) GetData() isEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetPhoneConfirmationEvent() *CreatePhoneConfirmationEventV1 {
	if x, ok := m.GetData().(*Event_PhoneConfirmationEvent); ok {
		return x.PhoneConfirmationEvent
	}
	return nil
}

func (m *Event) GetEmailConfirmationEvent() *CreateEmailConfirmationEventV1 {
	if x, ok := m.GetData().(*Event_EmailConfirmationEvent); ok {
		return x.EmailConfirmationEvent
	}
	return nil
}

func (m *Event) GetChangedUserViewsEvent() *ChangedUserViewsEventV1 {
	if x, ok := m.GetData().(*Event_ChangedUserViewsEvent); ok {
		return x.ChangedUserViewsEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_PhoneConfirmationEvent)(nil),
		(*Event_EmailConfirmationEvent)(nil),
		(*Event_ChangedUserViewsEvent)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// data
	switch x := m.Data.(type) {
	case *Event_PhoneConfirmationEvent:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PhoneConfirmationEvent); err != nil {
			return err
		}
	case *Event_EmailConfirmationEvent:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EmailConfirmationEvent); err != nil {
			return err
		}
	case *Event_ChangedUserViewsEvent:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ChangedUserViewsEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Data has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 6: // data.phoneConfirmationEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreatePhoneConfirmationEventV1)
		err := b.DecodeMessage(msg)
		m.Data = &Event_PhoneConfirmationEvent{msg}
		return true, err
	case 7: // data.emailConfirmationEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateEmailConfirmationEventV1)
		err := b.DecodeMessage(msg)
		m.Data = &Event_EmailConfirmationEvent{msg}
		return true, err
	case 8: // data.changedUserViewsEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ChangedUserViewsEventV1)
		err := b.DecodeMessage(msg)
		m.Data = &Event_ChangedUserViewsEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// data
	switch x := m.Data.(type) {
	case *Event_PhoneConfirmationEvent:
		s := proto.Size(x.PhoneConfirmationEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_EmailConfirmationEvent:
		s := proto.Size(x.EmailConfirmationEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_ChangedUserViewsEvent:
		s := proto.Size(x.ChangedUserViewsEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CreateEmailConfirmationEventV1)(nil), "inout.CreateEmailConfirmationEventV1")
	proto.RegisterType((*CreatePhoneConfirmationEventV1)(nil), "inout.CreatePhoneConfirmationEventV1")
	proto.RegisterType((*ChangedUserViewsEventV1)(nil), "inout.ChangedUserViewsEventV1")
	proto.RegisterType((*Event)(nil), "inout.Event")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_events_c19bd04ed2a6f279) }

var fileDescriptor_events_c19bd04ed2a6f279 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x4f, 0xb3, 0x40,
	0x14, 0x2c, 0x5f, 0x0b, 0x9f, 0x7d, 0xd4, 0xcb, 0x46, 0xeb, 0x9e, 0x1a, 0xac, 0x31, 0xe1, 0x44,
	0xa2, 0x1e, 0xbd, 0x49, 0x9a, 0x34, 0x9e, 0x0c, 0x89, 0x5c, 0x9b, 0x2d, 0xbc, 0xda, 0x35, 0x76,
	0x97, 0xec, 0xae, 0xf5, 0xb7, 0xf9, 0xef, 0xcc, 0x2e, 0xa0, 0x1e, 0x16, 0x6f, 0xcc, 0xbc, 0x61,
	0x86, 0xcc, 0x00, 0x33, 0x3c, 0xa2, 0x30, 0x3a, 0x6b, 0x94, 0x34, 0x92, 0x84, 0x5c, 0xc8, 0x77,
	0xb3, 0x7c, 0x84, 0x45, 0xae, 0x90, 0x19, 0x5c, 0x1d, 0x18, 0x7f, 0xcb, 0xa5, 0xd8, 0x71, 0x75,
	0x60, 0x86, 0x4b, 0xb1, 0xb2, 0xea, 0xf2, 0x86, 0x9c, 0x41, 0x88, 0xf6, 0x46, 0x83, 0x24, 0x48,
	0xa7, 0x45, 0x0b, 0x08, 0x81, 0x49, 0x25, 0x6b, 0xa4, 0xff, 0x1c, 0xe9, 0x9e, 0x7f, 0xbc, 0x9e,
	0xf6, 0x52, 0xe0, 0x80, 0x57, 0x63, 0x6f, 0xbd, 0x97, 0x03, 0x5e, 0xaf, 0x7b, 0xb8, 0xc8, 0xf7,
	0x4c, 0xbc, 0x60, 0xfd, 0xac, 0x51, 0x95, 0x1c, 0x3f, 0x74, 0x6f, 0x92, 0x40, 0xcc, 0x6b, 0x14,
	0x86, 0xef, 0x38, 0x2a, 0x4d, 0x83, 0x64, 0x9c, 0x4e, 0x8b, 0xdf, 0xd4, 0xf2, 0x73, 0x0c, 0xa1,
	0x53, 0x93, 0x39, 0x44, 0x1a, 0x45, 0x8d, 0xaa, 0x4b, 0xec, 0x90, 0xe5, 0xe5, 0xf6, 0x15, 0x2b,
	0xd3, 0x85, 0x76, 0xc8, 0xf2, 0xac, 0xb2, 0x5f, 0x4c, 0xc7, 0x2d, 0xdf, 0x22, 0x72, 0x05, 0xa7,
	0xae, 0xbd, 0xcd, 0x11, 0x95, 0xb6, 0xe7, 0x89, 0x3b, 0xb7, 0x95, 0x96, 0x2d, 0x47, 0x2e, 0x61,
	0x56, 0x33, 0xc3, 0xbe, 0x35, 0xa1, 0xd3, 0xc4, 0x96, 0xeb, 0x25, 0x1b, 0x98, 0x37, 0xde, 0x72,
	0x68, 0x94, 0x04, 0x69, 0x7c, 0x7b, 0x9d, 0xb9, 0x59, 0xb2, 0xbf, 0x7b, 0x5c, 0x8f, 0x8a, 0x01,
	0x1b, 0x1b, 0x80, 0xde, 0x25, 0xe9, 0x7f, 0x4f, 0xc0, 0xd0, 0xe8, 0x36, 0xc0, 0x6f, 0x43, 0x4a,
	0x38, 0xaf, 0x7c, 0xc3, 0xd0, 0x13, 0xe7, 0xbf, 0xe8, 0xfd, 0xfd, 0xe3, 0xad, 0x47, 0x85, 0xff,
	0xf5, 0x87, 0x08, 0x26, 0xb6, 0xa8, 0x6d, 0xe4, 0x7e, 0xcf, 0xbb, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x54, 0x4c, 0x0d, 0x46, 0xae, 0x02, 0x00, 0x00,
}
