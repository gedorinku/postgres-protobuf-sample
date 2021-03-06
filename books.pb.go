// Code generated by protoc-gen-go. DO NOT EDIT.
// source: books.proto

package main

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

type Book struct {
	Title                string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Authors              []*Author `protobuf:"bytes,3,rep,name=authors,proto3" json:"authors,omitempty"`
	PageCount            int32     `protobuf:"varint,4,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_7d86882b54f3bfc2, []int{0}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (dst *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(dst, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetAuthors() []*Author {
	if m != nil {
		return m.Authors
	}
	return nil
}

func (m *Book) GetPageCount() int32 {
	if m != nil {
		return m.PageCount
	}
	return 0
}

type Author struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_7d86882b54f3bfc2, []int{1}
}
func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (dst *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(dst, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Book)(nil), "main.Book")
	proto.RegisterType((*Author)(nil), "main.Author")
}

func init() { proto.RegisterFile("books.proto", fileDescriptor_books_7d86882b54f3bfc2) }

var fileDescriptor_books_7d86882b54f3bfc2 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xca, 0xcf, 0xcf,
	0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x4a, 0xe6,
	0x62, 0x71, 0xca, 0xcf, 0xcf, 0x16, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2c, 0xc9, 0x49, 0x95, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xd4, 0xb8, 0xd8, 0x13, 0x4b, 0x4b, 0x32, 0xf2,
	0x8b, 0x8a, 0x25, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x78, 0xf4, 0x40, 0xba, 0xf4, 0x1c, 0xc1,
	0x82, 0x41, 0x30, 0x49, 0x21, 0x59, 0x2e, 0xae, 0x82, 0xc4, 0xf4, 0xd4, 0xf8, 0xe4, 0xfc, 0xd2,
	0xbc, 0x12, 0x09, 0x16, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x4e, 0x90, 0x88, 0x33, 0x48, 0x40, 0x49,
	0x86, 0x8b, 0x0d, 0xa2, 0x43, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x11, 0x6c,
	0x0b, 0x98, 0x9d, 0xc4, 0x06, 0x76, 0x8f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x04, 0xeb, 0x04,
	0x1f, 0x9e, 0x00, 0x00, 0x00,
}
