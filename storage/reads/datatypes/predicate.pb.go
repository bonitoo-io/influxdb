// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: predicate.proto

package datatypes

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Node_Type int32

const (
	NodeTypeLogicalExpression    Node_Type = 0
	NodeTypeComparisonExpression Node_Type = 1
	NodeTypeParenExpression      Node_Type = 2
	NodeTypeTagRef               Node_Type = 3
	NodeTypeLiteral              Node_Type = 4
	NodeTypeFieldRef             Node_Type = 5
)

var Node_Type_name = map[int32]string{
	0: "LOGICAL_EXPRESSION",
	1: "COMPARISON_EXPRESSION",
	2: "PAREN_EXPRESSION",
	3: "TAG_REF",
	4: "LITERAL",
	5: "FIELD_REF",
}

var Node_Type_value = map[string]int32{
	"LOGICAL_EXPRESSION":    0,
	"COMPARISON_EXPRESSION": 1,
	"PAREN_EXPRESSION":      2,
	"TAG_REF":               3,
	"LITERAL":               4,
	"FIELD_REF":             5,
}

func (x Node_Type) String() string {
	return proto.EnumName(Node_Type_name, int32(x))
}

func (Node_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87cba9804b436f42, []int{0, 0}
}

type Node_Comparison int32

const (
	ComparisonEqual        Node_Comparison = 0
	ComparisonNotEqual     Node_Comparison = 1
	ComparisonStartsWith   Node_Comparison = 2
	ComparisonRegex        Node_Comparison = 3
	ComparisonNotRegex     Node_Comparison = 4
	ComparisonLess         Node_Comparison = 5
	ComparisonLessEqual    Node_Comparison = 6
	ComparisonGreater      Node_Comparison = 7
	ComparisonGreaterEqual Node_Comparison = 8
)

var Node_Comparison_name = map[int32]string{
	0: "EQUAL",
	1: "NOT_EQUAL",
	2: "STARTS_WITH",
	3: "REGEX",
	4: "NOT_REGEX",
	5: "LT",
	6: "LTE",
	7: "GT",
	8: "GTE",
}

var Node_Comparison_value = map[string]int32{
	"EQUAL":       0,
	"NOT_EQUAL":   1,
	"STARTS_WITH": 2,
	"REGEX":       3,
	"NOT_REGEX":   4,
	"LT":          5,
	"LTE":         6,
	"GT":          7,
	"GTE":         8,
}

func (x Node_Comparison) String() string {
	return proto.EnumName(Node_Comparison_name, int32(x))
}

func (Node_Comparison) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87cba9804b436f42, []int{0, 1}
}

// Logical operators apply to boolean values and combine to produce a single boolean result.
type Node_Logical int32

const (
	LogicalAnd Node_Logical = 0
	LogicalOr  Node_Logical = 1
)

var Node_Logical_name = map[int32]string{
	0: "AND",
	1: "OR",
}

var Node_Logical_value = map[string]int32{
	"AND": 0,
	"OR":  1,
}

func (x Node_Logical) String() string {
	return proto.EnumName(Node_Logical_name, int32(x))
}

func (Node_Logical) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87cba9804b436f42, []int{0, 2}
}

type Node struct {
	NodeType Node_Type `protobuf:"varint,1,opt,name=node_type,json=nodeType,proto3,enum=influxdata.platform.storage.Node_Type" json:"nodeType"`
	Children []*Node   `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Node_StringValue
	//	*Node_BooleanValue
	//	*Node_IntegerValue
	//	*Node_UnsignedValue
	//	*Node_FloatValue
	//	*Node_RegexValue
	//	*Node_TagRefValue
	//	*Node_FieldRefValue
	//	*Node_Logical_
	//	*Node_Comparison_
	Value isNode_Value `protobuf_oneof:"value"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_87cba9804b436f42, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

type isNode_Value interface {
	isNode_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Node_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof" json:"string_value,omitempty"`
}
type Node_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof" json:"bool_value,omitempty"`
}
type Node_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,5,opt,name=int_value,json=intValue,proto3,oneof" json:"int_value,omitempty"`
}
type Node_UnsignedValue struct {
	UnsignedValue uint64 `protobuf:"varint,6,opt,name=uint_value,json=uintValue,proto3,oneof" json:"uint_value,omitempty"`
}
type Node_FloatValue struct {
	FloatValue float64 `protobuf:"fixed64,7,opt,name=float_value,json=floatValue,proto3,oneof" json:"float_value,omitempty"`
}
type Node_RegexValue struct {
	RegexValue string `protobuf:"bytes,8,opt,name=regex_value,json=regexValue,proto3,oneof" json:"regex_value,omitempty"`
}
type Node_TagRefValue struct {
	TagRefValue string `protobuf:"bytes,9,opt,name=tag_ref_value,json=tagRefValue,proto3,oneof" json:"tag_ref_value,omitempty"`
}
type Node_FieldRefValue struct {
	FieldRefValue string `protobuf:"bytes,10,opt,name=field_ref_value,json=fieldRefValue,proto3,oneof" json:"field_ref_value,omitempty"`
}
type Node_Logical_ struct {
	Logical Node_Logical `protobuf:"varint,11,opt,name=logical,proto3,enum=influxdata.platform.storage.Node_Logical,oneof" json:"logical,omitempty"`
}
type Node_Comparison_ struct {
	Comparison Node_Comparison `protobuf:"varint,12,opt,name=comparison,proto3,enum=influxdata.platform.storage.Node_Comparison,oneof" json:"comparison,omitempty"`
}

func (*Node_StringValue) isNode_Value()   {}
func (*Node_BooleanValue) isNode_Value()  {}
func (*Node_IntegerValue) isNode_Value()  {}
func (*Node_UnsignedValue) isNode_Value() {}
func (*Node_FloatValue) isNode_Value()    {}
func (*Node_RegexValue) isNode_Value()    {}
func (*Node_TagRefValue) isNode_Value()   {}
func (*Node_FieldRefValue) isNode_Value() {}
func (*Node_Logical_) isNode_Value()      {}
func (*Node_Comparison_) isNode_Value()   {}

func (m *Node) GetValue() isNode_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Node) GetNodeType() Node_Type {
	if m != nil {
		return m.NodeType
	}
	return NodeTypeLogicalExpression
}

func (m *Node) GetChildren() []*Node {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Node) GetStringValue() string {
	if x, ok := m.GetValue().(*Node_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Node) GetBooleanValue() bool {
	if x, ok := m.GetValue().(*Node_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Node) GetIntegerValue() int64 {
	if x, ok := m.GetValue().(*Node_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (m *Node) GetUnsignedValue() uint64 {
	if x, ok := m.GetValue().(*Node_UnsignedValue); ok {
		return x.UnsignedValue
	}
	return 0
}

func (m *Node) GetFloatValue() float64 {
	if x, ok := m.GetValue().(*Node_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *Node) GetRegexValue() string {
	if x, ok := m.GetValue().(*Node_RegexValue); ok {
		return x.RegexValue
	}
	return ""
}

func (m *Node) GetTagRefValue() string {
	if x, ok := m.GetValue().(*Node_TagRefValue); ok {
		return x.TagRefValue
	}
	return ""
}

func (m *Node) GetFieldRefValue() string {
	if x, ok := m.GetValue().(*Node_FieldRefValue); ok {
		return x.FieldRefValue
	}
	return ""
}

func (m *Node) GetLogical() Node_Logical {
	if x, ok := m.GetValue().(*Node_Logical_); ok {
		return x.Logical
	}
	return LogicalAnd
}

func (m *Node) GetComparison() Node_Comparison {
	if x, ok := m.GetValue().(*Node_Comparison_); ok {
		return x.Comparison
	}
	return ComparisonEqual
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Node) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Node_StringValue)(nil),
		(*Node_BooleanValue)(nil),
		(*Node_IntegerValue)(nil),
		(*Node_UnsignedValue)(nil),
		(*Node_FloatValue)(nil),
		(*Node_RegexValue)(nil),
		(*Node_TagRefValue)(nil),
		(*Node_FieldRefValue)(nil),
		(*Node_Logical_)(nil),
		(*Node_Comparison_)(nil),
	}
}

type Predicate struct {
	Root *Node `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *Predicate) Reset()         { *m = Predicate{} }
func (m *Predicate) String() string { return proto.CompactTextString(m) }
func (*Predicate) ProtoMessage()    {}
func (*Predicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_87cba9804b436f42, []int{1}
}
func (m *Predicate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Predicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Predicate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Predicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Predicate.Merge(m, src)
}
func (m *Predicate) XXX_Size() int {
	return m.Size()
}
func (m *Predicate) XXX_DiscardUnknown() {
	xxx_messageInfo_Predicate.DiscardUnknown(m)
}

var xxx_messageInfo_Predicate proto.InternalMessageInfo

func (m *Predicate) GetRoot() *Node {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterEnum("influxdata.platform.storage.Node_Type", Node_Type_name, Node_Type_value)
	proto.RegisterEnum("influxdata.platform.storage.Node_Comparison", Node_Comparison_name, Node_Comparison_value)
	proto.RegisterEnum("influxdata.platform.storage.Node_Logical", Node_Logical_name, Node_Logical_value)
	proto.RegisterType((*Node)(nil), "influxdata.platform.storage.Node")
	proto.RegisterType((*Predicate)(nil), "influxdata.platform.storage.Predicate")
}

func init() { proto.RegisterFile("predicate.proto", fileDescriptor_87cba9804b436f42) }

var fileDescriptor_87cba9804b436f42 = []byte{
	// 869 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0xc7, 0x45, 0x7d, 0x58, 0xe2, 0xc8, 0x1f, 0xcc, 0xc6, 0x8e, 0x55, 0xa6, 0x91, 0xb6, 0x36,
	0x5a, 0x28, 0x40, 0xa1, 0xc0, 0x6e, 0x7d, 0x0a, 0x8a, 0x82, 0x74, 0x68, 0x59, 0x00, 0x2b, 0xa9,
	0x14, 0xd3, 0x04, 0x05, 0x0a, 0x81, 0xb1, 0x56, 0x0c, 0x01, 0x86, 0xab, 0x2e, 0x57, 0x85, 0xf3,
	0x06, 0x05, 0x4f, 0xbd, 0x07, 0x3c, 0xf5, 0x65, 0x7a, 0x29, 0x90, 0x63, 0x4f, 0x42, 0x21, 0xdf,
	0xfa, 0x14, 0x05, 0xbf, 0xe5, 0xb6, 0xa8, 0x73, 0xdb, 0x99, 0xfd, 0xff, 0xfe, 0x3b, 0x3b, 0x1c,
	0xad, 0x60, 0x6f, 0xc1, 0xc8, 0xcc, 0xb9, 0xb2, 0x38, 0xe9, 0x2d, 0x18, 0xe5, 0x14, 0x3d, 0x74,
	0xbc, 0xb9, 0xbb, 0xbc, 0x9e, 0x59, 0xdc, 0xea, 0x2d, 0x5c, 0x8b, 0xcf, 0x29, 0x7b, 0xd3, 0xf3,
	0x39, 0x65, 0x96, 0x4d, 0xe4, 0x7d, 0x9b, 0xda, 0x34, 0xd6, 0x3d, 0x89, 0x56, 0x09, 0x72, 0xf4,
	0xae, 0x09, 0xd5, 0x21, 0x9d, 0x11, 0xf4, 0x03, 0x88, 0x1e, 0x9d, 0x91, 0x29, 0x7f, 0xbb, 0x20,
	0x2d, 0x01, 0x0b, 0xdd, 0xdd, 0xd3, 0xcf, 0x7a, 0xff, 0xe3, 0xd7, 0x8b, 0xa8, 0x9e, 0xf9, 0x76,
	0x41, 0xd4, 0xd6, 0x7a, 0xd5, 0x69, 0x44, 0x61, 0x14, 0xfd, 0xb5, 0xea, 0x34, 0xbc, 0x74, 0x6d,
	0xe4, 0x2b, 0xf4, 0x15, 0x34, 0xae, 0x5e, 0x3b, 0xee, 0x8c, 0x11, 0xaf, 0x55, 0xc6, 0x95, 0x6e,
	0xf3, 0xf4, 0x93, 0x3b, 0xdd, 0x8d, 0x1c, 0x41, 0x5f, 0xc2, 0xb6, 0xcf, 0x99, 0xe3, 0xd9, 0xd3,
	0x9f, 0x2c, 0x77, 0x49, 0x5a, 0x15, 0x2c, 0x74, 0x45, 0x75, 0x6f, 0xbd, 0xea, 0x34, 0x27, 0x71,
	0xfe, 0xbb, 0x28, 0x7d, 0x59, 0x32, 0x9a, 0x7e, 0x11, 0xa2, 0x13, 0x80, 0x57, 0x94, 0xba, 0x29,
	0x53, 0xc5, 0x42, 0xb7, 0xa1, 0x4a, 0xeb, 0x55, 0x67, 0x5b, 0xa5, 0xd4, 0x25, 0x96, 0x97, 0x41,
	0x62, 0xa4, 0x4a, 0x90, 0x27, 0x20, 0x3a, 0x1e, 0x4f, 0x89, 0x1a, 0x16, 0xba, 0x95, 0x84, 0x18,
	0x78, 0x9c, 0xd8, 0x84, 0x65, 0x44, 0xc3, 0xf1, 0x78, 0x02, 0x9c, 0x02, 0x2c, 0x0b, 0x62, 0x0b,
	0x0b, 0xdd, 0xaa, 0x7a, 0x6f, 0xbd, 0xea, 0xec, 0x3c, 0xf7, 0x7c, 0xc7, 0xf6, 0xc8, 0x2c, 0x3f,
	0x64, 0x99, 0x33, 0x27, 0xd0, 0x9c, 0xbb, 0xd4, 0xca, 0xa0, 0x3a, 0x16, 0xba, 0x82, 0xba, 0xbb,
	0x5e, 0x75, 0xe0, 0x22, 0x4a, 0x67, 0x04, 0xcc, 0xf3, 0x28, 0x42, 0x18, 0xb1, 0xc9, 0x75, 0x8a,
	0x34, 0xe2, 0xfb, 0xc7, 0x88, 0x11, 0xa5, 0x73, 0x84, 0xe5, 0x11, 0x3a, 0x83, 0x1d, 0x6e, 0xd9,
	0x53, 0x46, 0xe6, 0x29, 0x24, 0x16, 0x4d, 0x33, 0x2d, 0xdb, 0x20, 0xf3, 0xbc, 0x69, 0xbc, 0x08,
	0xd1, 0x53, 0xd8, 0x9b, 0x3b, 0xc4, 0x9d, 0x6d, 0x80, 0x10, 0x83, 0xf1, 0xad, 0x2e, 0xa2, 0xad,
	0x0d, 0x74, 0x67, 0xbe, 0x99, 0x40, 0x1a, 0xd4, 0x5d, 0x6a, 0x3b, 0x57, 0x96, 0xdb, 0x6a, 0xc6,
	0x33, 0xf4, 0xf8, 0xee, 0x19, 0xd2, 0x13, 0xe0, 0xb2, 0x64, 0x64, 0x2c, 0x1a, 0x02, 0x5c, 0xd1,
	0x37, 0x0b, 0x8b, 0x39, 0x3e, 0xf5, 0x5a, 0xdb, 0xb1, 0xd3, 0xe7, 0x77, 0x3b, 0x9d, 0xe7, 0x4c,
	0xd4, 0x8a, 0xc2, 0xe1, 0xe8, 0x5d, 0x19, 0xaa, 0xf1, 0x18, 0x9e, 0x01, 0xd2, 0x47, 0xfd, 0xc1,
	0xb9, 0xa2, 0x4f, 0xb5, 0x97, 0x63, 0x43, 0x9b, 0x4c, 0x06, 0xa3, 0xa1, 0x54, 0x92, 0x1f, 0x05,
	0x21, 0xfe, 0x28, 0x1b, 0xe1, 0xb4, 0x20, 0xed, 0x7a, 0xc1, 0x88, 0xef, 0x3b, 0xd4, 0x43, 0x4f,
	0xe1, 0xe0, 0x7c, 0xf4, 0xcd, 0x58, 0x31, 0x06, 0x93, 0xd1, 0x70, 0x93, 0x14, 0x64, 0x1c, 0x84,
	0xf8, 0xe3, 0x8c, 0x2c, 0x0a, 0xd8, 0x80, 0x4f, 0x40, 0x1a, 0x2b, 0x86, 0x76, 0x8b, 0x2b, 0xcb,
	0x0f, 0x83, 0x10, 0x1f, 0x66, 0xdc, 0xd8, 0x62, 0x64, 0x13, 0xe9, 0x40, 0xdd, 0x54, 0xfa, 0x53,
	0x43, 0xbb, 0x90, 0x2a, 0x32, 0x0a, 0x42, 0xbc, 0x9b, 0x29, 0x93, 0x0f, 0x87, 0x30, 0xd4, 0xf5,
	0x81, 0xa9, 0x19, 0x8a, 0x2e, 0x55, 0xe5, 0xfb, 0x41, 0x88, 0xf7, 0xf2, 0xe2, 0x1d, 0x4e, 0x98,
	0xe5, 0xa2, 0x63, 0x10, 0x2f, 0x06, 0x9a, 0xfe, 0x2c, 0x36, 0xa9, 0xc9, 0xfb, 0x41, 0x88, 0xa5,
	0x4c, 0x93, 0x7d, 0x44, 0xb9, 0xfa, 0xf3, 0xaf, 0xed, 0xd2, 0xd1, 0xef, 0x65, 0x80, 0xa2, 0x72,
	0xd4, 0x86, 0x9a, 0xf6, 0xed, 0x73, 0x45, 0x97, 0x4a, 0x89, 0xf3, 0xc6, 0xa5, 0x7e, 0x5c, 0x5a,
	0x2e, 0xfa, 0x14, 0xc4, 0xe1, 0xc8, 0x9c, 0x26, 0x1a, 0x41, 0x7e, 0x10, 0x84, 0x18, 0x15, 0x9a,
	0x21, 0xe5, 0x89, 0xec, 0x31, 0x34, 0x27, 0xa6, 0x62, 0x98, 0x93, 0xe9, 0x8b, 0x81, 0x79, 0x29,
	0x95, 0xe5, 0x56, 0x10, 0xe2, 0xfd, 0x42, 0x38, 0xe1, 0x16, 0xe3, 0xfe, 0x0b, 0x87, 0xbf, 0x8e,
	0x4e, 0x34, 0xb4, 0xbe, 0xf6, 0x52, 0xaa, 0xfc, 0xf3, 0xc4, 0x78, 0xb8, 0xb3, 0x13, 0x13, 0x4d,
	0xf5, 0x3f, 0x4e, 0x4c, 0x64, 0x32, 0x94, 0x75, 0x53, 0xaa, 0x25, 0x0d, 0x2b, 0xf6, 0x75, 0xe2,
	0xfb, 0x08, 0x43, 0x45, 0x37, 0x35, 0x69, 0x4b, 0x3e, 0x0c, 0x42, 0x7c, 0xff, 0xf6, 0x66, 0x52,
	0xef, 0x23, 0x28, 0xf7, 0x4d, 0xa9, 0x2e, 0x1f, 0x04, 0x21, 0xbe, 0x57, 0x08, 0xfa, 0x8c, 0x58,
	0x9c, 0x30, 0x74, 0x0c, 0x95, 0xbe, 0xa9, 0x49, 0x0d, 0x59, 0x0e, 0x42, 0xfc, 0xe0, 0x5f, 0xfb,
	0xb1, 0x47, 0xda, 0xcf, 0xaf, 0xa1, 0x9e, 0x8e, 0x10, 0x3a, 0x84, 0x8a, 0x32, 0x7c, 0x26, 0x95,
	0xe4, 0xdd, 0x20, 0xc4, 0x90, 0x66, 0x15, 0x6f, 0x86, 0x0e, 0xa0, 0x3c, 0x32, 0x24, 0x41, 0xde,
	0x09, 0x42, 0x2c, 0xa6, 0xf9, 0x11, 0x4b, 0x0c, 0xd4, 0x3a, 0xd4, 0xe2, 0x1f, 0xde, 0x91, 0x0a,
	0xe2, 0x38, 0x7b, 0xe3, 0xd1, 0x19, 0x54, 0x19, 0xa5, 0x3c, 0x7e, 0x9c, 0x3f, 0xe8, 0xf9, 0x8c,
	0xe5, 0xea, 0xf1, 0x6f, 0xeb, 0xb6, 0xf0, 0x7e, 0xdd, 0x16, 0xfe, 0x5c, 0xb7, 0x85, 0x5f, 0x6e,
	0xda, 0xa5, 0xf7, 0x37, 0xed, 0xd2, 0x1f, 0x37, 0xed, 0xd2, 0xf7, 0x62, 0xc4, 0x46, 0xef, 0xbd,
	0xff, 0x6a, 0x2b, 0xfe, 0x37, 0xf8, 0xe2, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x63, 0xc0,
	0x0c, 0x53, 0x06, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Children) > 0 {
		for iNdEx := len(m.Children) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Children[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPredicate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.NodeType != 0 {
		i = encodeVarintPredicate(dAtA, i, uint64(m.NodeType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Node_StringValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_StringValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.StringValue)
	copy(dAtA[i:], m.StringValue)
	i = encodeVarintPredicate(dAtA, i, uint64(len(m.StringValue)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *Node_BooleanValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_BooleanValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.BooleanValue {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x20
	return len(dAtA) - i, nil
}
func (m *Node_IntegerValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_IntegerValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintPredicate(dAtA, i, uint64(m.IntegerValue))
	i--
	dAtA[i] = 0x28
	return len(dAtA) - i, nil
}
func (m *Node_UnsignedValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_UnsignedValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintPredicate(dAtA, i, uint64(m.UnsignedValue))
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}
func (m *Node_FloatValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_FloatValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.FloatValue))))
	i--
	dAtA[i] = 0x39
	return len(dAtA) - i, nil
}
func (m *Node_RegexValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_RegexValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.RegexValue)
	copy(dAtA[i:], m.RegexValue)
	i = encodeVarintPredicate(dAtA, i, uint64(len(m.RegexValue)))
	i--
	dAtA[i] = 0x42
	return len(dAtA) - i, nil
}
func (m *Node_TagRefValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_TagRefValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.TagRefValue)
	copy(dAtA[i:], m.TagRefValue)
	i = encodeVarintPredicate(dAtA, i, uint64(len(m.TagRefValue)))
	i--
	dAtA[i] = 0x4a
	return len(dAtA) - i, nil
}
func (m *Node_FieldRefValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_FieldRefValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.FieldRefValue)
	copy(dAtA[i:], m.FieldRefValue)
	i = encodeVarintPredicate(dAtA, i, uint64(len(m.FieldRefValue)))
	i--
	dAtA[i] = 0x52
	return len(dAtA) - i, nil
}
func (m *Node_Logical_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_Logical_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintPredicate(dAtA, i, uint64(m.Logical))
	i--
	dAtA[i] = 0x58
	return len(dAtA) - i, nil
}
func (m *Node_Comparison_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node_Comparison_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintPredicate(dAtA, i, uint64(m.Comparison))
	i--
	dAtA[i] = 0x60
	return len(dAtA) - i, nil
}
func (m *Predicate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Predicate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Predicate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Root != nil {
		{
			size, err := m.Root.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPredicate(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPredicate(dAtA []byte, offset int, v uint64) int {
	offset -= sovPredicate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeType != 0 {
		n += 1 + sovPredicate(uint64(m.NodeType))
	}
	if len(m.Children) > 0 {
		for _, e := range m.Children {
			l = e.Size()
			n += 1 + l + sovPredicate(uint64(l))
		}
	}
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *Node_StringValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StringValue)
	n += 1 + l + sovPredicate(uint64(l))
	return n
}
func (m *Node_BooleanValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *Node_IntegerValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovPredicate(uint64(m.IntegerValue))
	return n
}
func (m *Node_UnsignedValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovPredicate(uint64(m.UnsignedValue))
	return n
}
func (m *Node_FloatValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 9
	return n
}
func (m *Node_RegexValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RegexValue)
	n += 1 + l + sovPredicate(uint64(l))
	return n
}
func (m *Node_TagRefValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TagRefValue)
	n += 1 + l + sovPredicate(uint64(l))
	return n
}
func (m *Node_FieldRefValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FieldRefValue)
	n += 1 + l + sovPredicate(uint64(l))
	return n
}
func (m *Node_Logical_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovPredicate(uint64(m.Logical))
	return n
}
func (m *Node_Comparison_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovPredicate(uint64(m.Comparison))
	return n
}
func (m *Predicate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Root != nil {
		l = m.Root.Size()
		n += 1 + l + sovPredicate(uint64(l))
	}
	return n
}

func sovPredicate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPredicate(x uint64) (n int) {
	return sovPredicate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPredicate
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeType", wireType)
			}
			m.NodeType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeType |= Node_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPredicate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, &Node{})
			if err := m.Children[len(m.Children)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPredicate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Node_StringValue{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BooleanValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Value = &Node_BooleanValue{b}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntegerValue", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Node_IntegerValue{v}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnsignedValue", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Node_UnsignedValue{v}
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatValue", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = &Node_FloatValue{float64(math.Float64frombits(v))}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegexValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPredicate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Node_RegexValue{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagRefValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPredicate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Node_TagRefValue{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldRefValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPredicate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Node_FieldRefValue{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logical", wireType)
			}
			var v Node_Logical
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= Node_Logical(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Node_Logical_{v}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comparison", wireType)
			}
			var v Node_Comparison
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= Node_Comparison(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Node_Comparison_{v}
		default:
			iNdEx = preIndex
			skippy, err := skipPredicate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPredicate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPredicate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Predicate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPredicate
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Predicate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Predicate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPredicate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPredicate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Root == nil {
				m.Root = &Node{}
			}
			if err := m.Root.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPredicate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPredicate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPredicate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPredicate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPredicate
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPredicate
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPredicate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPredicate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPredicate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPredicate        = fmt.Errorf("proto: negative length found during unmarshalling")
	ErrIntOverflowPredicate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPredicate = fmt.Errorf("proto: unexpected end of group")
)
