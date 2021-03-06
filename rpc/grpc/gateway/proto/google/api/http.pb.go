// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/http.proto

package google_api

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

type Http struct {
	Rules                []*HttpRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Http) Reset()         { *m = Http{} }
func (m *Http) String() string { return proto.CompactTextString(m) }
func (*Http) ProtoMessage()    {}
func (*Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9994be407cdcc9, []int{0}
}

func (m *Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http.Unmarshal(m, b)
}
func (m *Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http.Marshal(b, m, deterministic)
}
func (m *Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http.Merge(m, src)
}
func (m *Http) XXX_Size() int {
	return xxx_messageInfo_Http.Size(m)
}
func (m *Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Http proto.InternalMessageInfo

func (m *Http) GetRules() []*HttpRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type HttpRule struct {
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Types that are valid to be assigned to Pattern:
	//	*HttpRule_Get
	//	*HttpRule_Put
	//	*HttpRule_Post
	//	*HttpRule_Delete
	//	*HttpRule_Patch
	//	*HttpRule_Custom
	Pattern              isHttpRule_Pattern `protobuf_oneof:"pattern"`
	Body                 string             `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	AditionalBidings     []*HttpRule        `protobuf:"bytes,11,rep,name=aditional_bidings,json=aditionalBidings,proto3" json:"aditional_bidings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HttpRule) Reset()         { *m = HttpRule{} }
func (m *HttpRule) String() string { return proto.CompactTextString(m) }
func (*HttpRule) ProtoMessage()    {}
func (*HttpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9994be407cdcc9, []int{1}
}

func (m *HttpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRule.Unmarshal(m, b)
}
func (m *HttpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRule.Marshal(b, m, deterministic)
}
func (m *HttpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRule.Merge(m, src)
}
func (m *HttpRule) XXX_Size() int {
	return xxx_messageInfo_HttpRule.Size(m)
}
func (m *HttpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRule proto.InternalMessageInfo

func (m *HttpRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
}

type HttpRule_Get struct {
	Get string `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type HttpRule_Put struct {
	Put string `protobuf:"bytes,3,opt,name=put,proto3,oneof"`
}

type HttpRule_Post struct {
	Post string `protobuf:"bytes,4,opt,name=post,proto3,oneof"`
}

type HttpRule_Delete struct {
	Delete string `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}

type HttpRule_Patch struct {
	Patch string `protobuf:"bytes,6,opt,name=patch,proto3,oneof"`
}

type HttpRule_Custom struct {
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom,proto3,oneof"`
}

func (*HttpRule_Get) isHttpRule_Pattern() {}

func (*HttpRule_Put) isHttpRule_Pattern() {}

func (*HttpRule_Post) isHttpRule_Pattern() {}

func (*HttpRule_Delete) isHttpRule_Pattern() {}

func (*HttpRule_Patch) isHttpRule_Pattern() {}

func (*HttpRule_Custom) isHttpRule_Pattern() {}

func (m *HttpRule) GetPattern() isHttpRule_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *HttpRule) GetGet() string {
	if x, ok := m.GetPattern().(*HttpRule_Get); ok {
		return x.Get
	}
	return ""
}

func (m *HttpRule) GetPut() string {
	if x, ok := m.GetPattern().(*HttpRule_Put); ok {
		return x.Put
	}
	return ""
}

func (m *HttpRule) GetPost() string {
	if x, ok := m.GetPattern().(*HttpRule_Post); ok {
		return x.Post
	}
	return ""
}

func (m *HttpRule) GetDelete() string {
	if x, ok := m.GetPattern().(*HttpRule_Delete); ok {
		return x.Delete
	}
	return ""
}

func (m *HttpRule) GetPatch() string {
	if x, ok := m.GetPattern().(*HttpRule_Patch); ok {
		return x.Patch
	}
	return ""
}

func (m *HttpRule) GetCustom() *CustomHttpPattern {
	if x, ok := m.GetPattern().(*HttpRule_Custom); ok {
		return x.Custom
	}
	return nil
}

func (m *HttpRule) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HttpRule) GetAditionalBidings() []*HttpRule {
	if m != nil {
		return m.AditionalBidings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpRule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpRule_Get)(nil),
		(*HttpRule_Put)(nil),
		(*HttpRule_Post)(nil),
		(*HttpRule_Delete)(nil),
		(*HttpRule_Patch)(nil),
		(*HttpRule_Custom)(nil),
	}
}

type CustomHttpPattern struct {
	Kind                 string   `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomHttpPattern) Reset()         { *m = CustomHttpPattern{} }
func (m *CustomHttpPattern) String() string { return proto.CompactTextString(m) }
func (*CustomHttpPattern) ProtoMessage()    {}
func (*CustomHttpPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9994be407cdcc9, []int{2}
}

func (m *CustomHttpPattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomHttpPattern.Unmarshal(m, b)
}
func (m *CustomHttpPattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomHttpPattern.Marshal(b, m, deterministic)
}
func (m *CustomHttpPattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomHttpPattern.Merge(m, src)
}
func (m *CustomHttpPattern) XXX_Size() int {
	return xxx_messageInfo_CustomHttpPattern.Size(m)
}
func (m *CustomHttpPattern) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomHttpPattern.DiscardUnknown(m)
}

var xxx_messageInfo_CustomHttpPattern proto.InternalMessageInfo

func (m *CustomHttpPattern) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CustomHttpPattern) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Http)(nil), "google.api.Http")
	proto.RegisterType((*HttpRule)(nil), "google.api.HttpRule")
	proto.RegisterType((*CustomHttpPattern)(nil), "google.api.CustomHttpPattern")
}

func init() { proto.RegisterFile("google/api/http.proto", fileDescriptor_ff9994be407cdcc9) }

var fileDescriptor_ff9994be407cdcc9 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x71, 0x93, 0xa6, 0xed, 0x75, 0xa1, 0xa7, 0x82, 0x2c, 0x24, 0xa4, 0xaa, 0x53, 0xc5,
	0x90, 0x4a, 0x65, 0x60, 0x60, 0x22, 0x2c, 0x1d, 0x51, 0x5e, 0x00, 0x39, 0x89, 0x95, 0x58, 0x98,
	0xd8, 0x4a, 0x2e, 0x03, 0x8f, 0xc5, 0xdb, 0x31, 0x22, 0xdb, 0x69, 0x41, 0x42, 0x62, 0xbb, 0xff,
	0xbb, 0x5f, 0xbe, 0xbb, 0xdf, 0x70, 0x55, 0x1b, 0x53, 0x6b, 0xb9, 0x17, 0x56, 0xed, 0x1b, 0x22,
	0x9b, 0xda, 0xce, 0x90, 0x41, 0x08, 0x38, 0x15, 0x56, 0x6d, 0x0f, 0x10, 0x1f, 0x89, 0x2c, 0xde,
	0xc1, 0xb4, 0x1b, 0xb4, 0xec, 0x39, 0xdb, 0x44, 0xbb, 0xe5, 0x61, 0x9d, 0xfe, 0x78, 0x52, 0x67,
	0xc8, 0x07, 0x2d, 0xf3, 0x60, 0xd9, 0x7e, 0x4e, 0x60, 0x7e, 0x62, 0x78, 0x03, 0xf3, 0x5e, 0x6a,
	0x59, 0x92, 0xe9, 0x38, 0xdb, 0xb0, 0xdd, 0x22, 0x3f, 0x6b, 0x44, 0x88, 0x6a, 0x49, 0x7c, 0xe2,
	0xf0, 0xf1, 0x22, 0x77, 0xc2, 0x31, 0x3b, 0x10, 0x8f, 0x4e, 0xcc, 0x0e, 0x84, 0x6b, 0x88, 0xad,
	0xe9, 0x89, 0xc7, 0x23, 0xf4, 0x0a, 0x39, 0x24, 0x95, 0xd4, 0x92, 0x24, 0x9f, 0x8e, 0x7c, 0xd4,
	0x78, 0x0d, 0x53, 0x2b, 0xa8, 0x6c, 0x78, 0x32, 0x36, 0x82, 0xc4, 0x07, 0x48, 0xca, 0xa1, 0x27,
	0xf3, 0xce, 0xe7, 0x1b, 0xb6, 0x5b, 0x1e, 0x6e, 0x7f, 0x5f, 0xf1, 0xec, 0x3b, 0x6e, 0xef, 0x17,
	0x41, 0x24, 0xbb, 0xd6, 0x3d, 0x18, 0xec, 0x88, 0x10, 0x17, 0xa6, 0xfa, 0xe0, 0x33, 0x7f, 0x80,
	0xaf, 0xf1, 0x09, 0x56, 0xa2, 0x52, 0xa4, 0x4c, 0x2b, 0xf4, 0x6b, 0xa1, 0x2a, 0xd5, 0xd6, 0x3d,
	0x5f, 0xfe, 0x93, 0xce, 0xe5, 0xd9, 0x9e, 0x05, 0x77, 0xb6, 0x80, 0x99, 0x0d, 0xb3, 0xb6, 0x8f,
	0xb0, 0xfa, 0xb3, 0x80, 0x1b, 0xfb, 0xa6, 0xda, 0x6a, 0xcc, 0xcd, 0xd7, 0x8e, 0x59, 0x41, 0x4d,
	0x08, 0x2d, 0xf7, 0x75, 0x16, 0x7d, 0x31, 0x56, 0x24, 0xfe, 0xf3, 0xee, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xff, 0xfd, 0x71, 0xc1, 0xd5, 0x01, 0x00, 0x00,
}
