// Code generated by protoc-gen-go. DO NOT EDIT.
// source: net/redeemer.proto

package net

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Ticket struct {
	TicketParams         *TicketParams           `protobuf:"bytes,1,opt,name=ticket_params,json=ticketParams,proto3" json:"ticket_params,omitempty"`
	Sender               []byte                  `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	ExpirationParams     *TicketExpirationParams `protobuf:"bytes,3,opt,name=expiration_params,json=expirationParams,proto3" json:"expiration_params,omitempty"`
	SenderParams         *TicketSenderParams     `protobuf:"bytes,4,opt,name=sender_params,json=senderParams,proto3" json:"sender_params,omitempty"`
	RecipientRand        []byte                  `protobuf:"bytes,5,opt,name=recipient_rand,json=recipientRand,proto3" json:"recipient_rand,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Ticket) Reset()         { *m = Ticket{} }
func (m *Ticket) String() string { return proto.CompactTextString(m) }
func (*Ticket) ProtoMessage()    {}
func (*Ticket) Descriptor() ([]byte, []int) {
	return fileDescriptor_41a074e4ea0232f2, []int{0}
}

func (m *Ticket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ticket.Unmarshal(m, b)
}
func (m *Ticket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ticket.Marshal(b, m, deterministic)
}
func (m *Ticket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticket.Merge(m, src)
}
func (m *Ticket) XXX_Size() int {
	return xxx_messageInfo_Ticket.Size(m)
}
func (m *Ticket) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticket.DiscardUnknown(m)
}

var xxx_messageInfo_Ticket proto.InternalMessageInfo

func (m *Ticket) GetTicketParams() *TicketParams {
	if m != nil {
		return m.TicketParams
	}
	return nil
}

func (m *Ticket) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Ticket) GetExpirationParams() *TicketExpirationParams {
	if m != nil {
		return m.ExpirationParams
	}
	return nil
}

func (m *Ticket) GetSenderParams() *TicketSenderParams {
	if m != nil {
		return m.SenderParams
	}
	return nil
}

func (m *Ticket) GetRecipientRand() []byte {
	if m != nil {
		return m.RecipientRand
	}
	return nil
}

type QueueTicketRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueTicketRes) Reset()         { *m = QueueTicketRes{} }
func (m *QueueTicketRes) String() string { return proto.CompactTextString(m) }
func (*QueueTicketRes) ProtoMessage()    {}
func (*QueueTicketRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_41a074e4ea0232f2, []int{1}
}

func (m *QueueTicketRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueTicketRes.Unmarshal(m, b)
}
func (m *QueueTicketRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueTicketRes.Marshal(b, m, deterministic)
}
func (m *QueueTicketRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueTicketRes.Merge(m, src)
}
func (m *QueueTicketRes) XXX_Size() int {
	return xxx_messageInfo_QueueTicketRes.Size(m)
}
func (m *QueueTicketRes) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueTicketRes.DiscardUnknown(m)
}

var xxx_messageInfo_QueueTicketRes proto.InternalMessageInfo

type MaxFloatReq struct {
	Sender               []byte   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxFloatReq) Reset()         { *m = MaxFloatReq{} }
func (m *MaxFloatReq) String() string { return proto.CompactTextString(m) }
func (*MaxFloatReq) ProtoMessage()    {}
func (*MaxFloatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41a074e4ea0232f2, []int{2}
}

func (m *MaxFloatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxFloatReq.Unmarshal(m, b)
}
func (m *MaxFloatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxFloatReq.Marshal(b, m, deterministic)
}
func (m *MaxFloatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxFloatReq.Merge(m, src)
}
func (m *MaxFloatReq) XXX_Size() int {
	return xxx_messageInfo_MaxFloatReq.Size(m)
}
func (m *MaxFloatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxFloatReq.DiscardUnknown(m)
}

var xxx_messageInfo_MaxFloatReq proto.InternalMessageInfo

func (m *MaxFloatReq) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

type MaxFloatUpdate struct {
	MaxFloat             []byte   `protobuf:"bytes,1,opt,name=max_float,json=maxFloat,proto3" json:"max_float,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxFloatUpdate) Reset()         { *m = MaxFloatUpdate{} }
func (m *MaxFloatUpdate) String() string { return proto.CompactTextString(m) }
func (*MaxFloatUpdate) ProtoMessage()    {}
func (*MaxFloatUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_41a074e4ea0232f2, []int{3}
}

func (m *MaxFloatUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxFloatUpdate.Unmarshal(m, b)
}
func (m *MaxFloatUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxFloatUpdate.Marshal(b, m, deterministic)
}
func (m *MaxFloatUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxFloatUpdate.Merge(m, src)
}
func (m *MaxFloatUpdate) XXX_Size() int {
	return xxx_messageInfo_MaxFloatUpdate.Size(m)
}
func (m *MaxFloatUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxFloatUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MaxFloatUpdate proto.InternalMessageInfo

func (m *MaxFloatUpdate) GetMaxFloat() []byte {
	if m != nil {
		return m.MaxFloat
	}
	return nil
}

func init() {
	proto.RegisterType((*Ticket)(nil), "net.Ticket")
	proto.RegisterType((*QueueTicketRes)(nil), "net.QueueTicketRes")
	proto.RegisterType((*MaxFloatReq)(nil), "net.MaxFloatReq")
	proto.RegisterType((*MaxFloatUpdate)(nil), "net.MaxFloatUpdate")
}

func init() {
	proto.RegisterFile("net/redeemer.proto", fileDescriptor_41a074e4ea0232f2)
}

var fileDescriptor_41a074e4ea0232f2 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x29, 0x08, 0xe2, 0x14, 0x6a, 0x19, 0x13, 0x25, 0x70, 0x21, 0x4d, 0x48, 0xb8, 0x58,
	0x14, 0x12, 0x4f, 0x9c, 0x4c, 0x34, 0x5e, 0x48, 0xb4, 0xea, 0xc5, 0x4b, 0xb3, 0xd2, 0x31, 0x69,
	0xa4, 0xdb, 0xba, 0x1d, 0x12, 0xde, 0xc0, 0x37, 0xf2, 0xf9, 0x0c, 0xdb, 0x96, 0xac, 0xdc, 0xbc,
	0x75, 0xfe, 0xf9, 0xe7, 0x9b, 0x7f, 0xd2, 0x05, 0x94, 0xc4, 0x53, 0x45, 0x11, 0x51, 0x42, 0xca,
	0xcf, 0x54, 0xca, 0x29, 0x36, 0x24, 0xf1, 0xc0, 0xdd, 0x35, 0xd6, 0x59, 0xa8, 0xb2, 0x55, 0x21,
	0x7b, 0xdf, 0x75, 0x68, 0xbd, 0xc4, 0xab, 0x4f, 0x62, 0xbc, 0x81, 0x2e, 0xeb, 0xaf, 0x30, 0x13,
	0x4a, 0x24, 0x79, 0xdf, 0x1a, 0x59, 0x13, 0x7b, 0xd6, 0xf3, 0x25, 0xb1, 0x5f, 0x78, 0x1e, 0x75,
	0x23, 0xe8, 0xb0, 0x51, 0xe1, 0x39, 0xb4, 0x72, 0x92, 0x11, 0xa9, 0x7e, 0x7d, 0x64, 0x4d, 0x3a,
	0x41, 0x59, 0xe1, 0x03, 0xf4, 0x68, 0x9b, 0xc5, 0x4a, 0x70, 0x9c, 0xca, 0x8a, 0xd9, 0xd0, 0xcc,
	0xa1, 0xc1, 0xbc, 0xdb, 0x7b, 0x4a, 0xba, 0x4b, 0x07, 0x0a, 0x2e, 0xa0, 0x5b, 0x30, 0x2b, 0xca,
	0x91, 0xa6, 0x5c, 0x18, 0x94, 0x67, 0xdd, 0xaf, 0xf2, 0xe5, 0x46, 0x85, 0x63, 0x70, 0x14, 0xad,
	0xe2, 0x2c, 0x26, 0xc9, 0xa1, 0x12, 0x32, 0xea, 0x37, 0x75, 0xce, 0xee, 0x5e, 0x0d, 0x84, 0x8c,
	0x3c, 0x17, 0x9c, 0xa7, 0x0d, 0x6d, 0xa8, 0xe0, 0x05, 0x94, 0x7b, 0x63, 0xb0, 0x97, 0x62, 0x7b,
	0xbf, 0x4e, 0x05, 0x07, 0xf4, 0x65, 0xdc, 0x69, 0x99, 0x77, 0x7a, 0x97, 0xe0, 0x54, 0xb6, 0xd7,
	0x2c, 0x12, 0x4c, 0x38, 0x84, 0x93, 0x44, 0x6c, 0xc3, 0x8f, 0x9d, 0x54, 0x9a, 0xdb, 0x49, 0x69,
	0x99, 0xfd, 0x58, 0xe0, 0x54, 0x3b, 0x8a, 0x3f, 0x84, 0xd7, 0x60, 0x1b, 0xab, 0xd1, 0x36, 0xee,
	0x1a, 0x9c, 0xe9, 0xe2, 0x20, 0x59, 0x0d, 0xe7, 0xd0, 0xae, 0x96, 0xa2, 0xab, 0x2d, 0x46, 0xd4,
	0x72, 0xe8, 0x6f, 0x2a, 0xaf, 0x86, 0x0b, 0x38, 0x5d, 0xa6, 0x32, 0xe6, 0x54, 0xfd, 0x7b, 0xf6,
	0xca, 0xba, 0x3d, 0x7e, 0x6b, 0xfa, 0x53, 0x49, 0xfc, 0xde, 0xd2, 0x4f, 0x67, 0xfe, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0xc0, 0x86, 0xc9, 0x67, 0x02, 0x00, 0x00,
}
