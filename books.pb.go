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
	BookId               uint32    `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Title                string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Authors              []*Author `protobuf:"bytes,3,rep,name=authors,proto3" json:"authors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_8f179872c876e166, []int{0}
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

func (m *Book) GetBookId() uint32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

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
	return fileDescriptor_books_8f179872c876e166, []int{1}
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

func init() { proto.RegisterFile("books.proto", fileDescriptor_books_8f179872c876e166) }

var fileDescriptor_books_8f179872c876e166 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xca, 0xcf, 0xcf,
	0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x8a, 0xe5,
	0x62, 0x71, 0xca, 0xcf, 0xcf, 0x16, 0x12, 0xe7, 0x62, 0x07, 0x49, 0xc6, 0x67, 0xa6, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0xf0, 0x06, 0xb1, 0x81, 0xb8, 0x9e, 0x29, 0x42, 0x22, 0x5c, 0xac, 0x25, 0x99,
	0x25, 0x39, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x1a, 0x17, 0x7b,
	0x62, 0x69, 0x49, 0x46, 0x7e, 0x51, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x11, 0x8f, 0x1e,
	0xc8, 0x38, 0x3d, 0x47, 0xb0, 0x60, 0x10, 0x4c, 0x52, 0x49, 0x86, 0x8b, 0x0d, 0x22, 0x24, 0x24,
	0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x0a, 0x36, 0x9d, 0x33, 0x08, 0xcc, 0x4e, 0x62, 0x03, 0xbb,
	0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xca, 0xa2, 0x40, 0x52, 0x98, 0x00, 0x00, 0x00,
}