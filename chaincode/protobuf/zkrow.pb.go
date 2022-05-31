// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zkrow.proto

//package main;

package zkrow_package

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

//zkrow represents a row in the public ledger
type Zkrow struct {
	Columns              map[string]*OrgColumn `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IsValidBalCor        bool                  `protobuf:"varint,2,opt,name=isValidBalCor,proto3" json:"isValidBalCor,omitempty"`
	IsValidAsset         bool                  `protobuf:"varint,3,opt,name=isValidAsset,proto3" json:"isValidAsset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Zkrow) Reset()         { *m = Zkrow{} }
func (m *Zkrow) String() string { return proto.CompactTextString(m) }
func (*Zkrow) ProtoMessage()    {}
func (*Zkrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_3687e745eb88f81a, []int{0}
}

func (m *Zkrow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Zkrow.Unmarshal(m, b)
}
func (m *Zkrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Zkrow.Marshal(b, m, deterministic)
}
func (m *Zkrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Zkrow.Merge(m, src)
}
func (m *Zkrow) XXX_Size() int {
	return xxx_messageInfo_Zkrow.Size(m)
}
func (m *Zkrow) XXX_DiscardUnknown() {
	xxx_messageInfo_Zkrow.DiscardUnknown(m)
}

var xxx_messageInfo_Zkrow proto.InternalMessageInfo

func (m *Zkrow) GetColumns() map[string]*OrgColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Zkrow) GetIsValidBalCor() bool {
	if m != nil {
		return m.IsValidBalCor
	}
	return false
}

func (m *Zkrow) GetIsValidAsset() bool {
	if m != nil {
		return m.IsValidAsset
	}
	return false
}

// OrgColumn represents one organization
type OrgColumn struct {
	// transaction content
	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	AuditToken []byte `protobuf:"bytes,2,opt,name=auditToken,proto3" json:"auditToken,omitempty"`
	// two step validation state
	IsValidBalCor bool `protobuf:"varint,3,opt,name=isValidBalCor,proto3" json:"isValidBalCor,omitempty"`
	IsValidAsset  bool `protobuf:"varint,4,opt,name=isValidAsset,proto3" json:"isValidAsset,omitempty"`
	// auxiliary data for proofs
	TokenPrime       []byte      `protobuf:"bytes,5,opt,name=TokenPrime,proto3" json:"TokenPrime,omitempty"`
	TokenDoublePrime []byte      `protobuf:"bytes,6,opt,name=TokenDoublePrime,proto3" json:"TokenDoublePrime,omitempty"`
	Rp               *RangeProof `protobuf:"bytes,7,opt,name=rp,proto3" json:"rp,omitempty"`
	//  bytes rp = 7;
	Dzkp                 *DisjunctiveProof `protobuf:"bytes,8,opt,name=dzkp,proto3" json:"dzkp,omitempty"`
	S                    []byte            `protobuf:"bytes,9,opt,name=s,proto3" json:"s,omitempty"`
	T                    []byte            `protobuf:"bytes,10,opt,name=t,proto3" json:"t,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OrgColumn) Reset()         { *m = OrgColumn{} }
func (m *OrgColumn) String() string { return proto.CompactTextString(m) }
func (*OrgColumn) ProtoMessage()    {}
func (*OrgColumn) Descriptor() ([]byte, []int) {
	return fileDescriptor_3687e745eb88f81a, []int{1}
}

func (m *OrgColumn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrgColumn.Unmarshal(m, b)
}
func (m *OrgColumn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrgColumn.Marshal(b, m, deterministic)
}
func (m *OrgColumn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrgColumn.Merge(m, src)
}
func (m *OrgColumn) XXX_Size() int {
	return xxx_messageInfo_OrgColumn.Size(m)
}
func (m *OrgColumn) XXX_DiscardUnknown() {
	xxx_messageInfo_OrgColumn.DiscardUnknown(m)
}

var xxx_messageInfo_OrgColumn proto.InternalMessageInfo

func (m *OrgColumn) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *OrgColumn) GetAuditToken() []byte {
	if m != nil {
		return m.AuditToken
	}
	return nil
}

func (m *OrgColumn) GetIsValidBalCor() bool {
	if m != nil {
		return m.IsValidBalCor
	}
	return false
}

func (m *OrgColumn) GetIsValidAsset() bool {
	if m != nil {
		return m.IsValidAsset
	}
	return false
}

func (m *OrgColumn) GetTokenPrime() []byte {
	if m != nil {
		return m.TokenPrime
	}
	return nil
}

func (m *OrgColumn) GetTokenDoublePrime() []byte {
	if m != nil {
		return m.TokenDoublePrime
	}
	return nil
}

func (m *OrgColumn) GetRp() *RangeProof {
	if m != nil {
		return m.Rp
	}
	return nil
}

func (m *OrgColumn) GetDzkp() *DisjunctiveProof {
	if m != nil {
		return m.Dzkp
	}
	return nil
}

func (m *OrgColumn) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *OrgColumn) GetT() []byte {
	if m != nil {
		return m.T
	}
	return nil
}

type DisjunctiveProof struct {
	Proof                []byte   `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
	G1                   []byte   `protobuf:"bytes,2,opt,name=g1,proto3" json:"g1,omitempty"`
	Y1                   []byte   `protobuf:"bytes,3,opt,name=y1,proto3" json:"y1,omitempty"`
	G2                   []byte   `protobuf:"bytes,4,opt,name=g2,proto3" json:"g2,omitempty"`
	Y2                   []byte   `protobuf:"bytes,5,opt,name=y2,proto3" json:"y2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisjunctiveProof) Reset()         { *m = DisjunctiveProof{} }
func (m *DisjunctiveProof) String() string { return proto.CompactTextString(m) }
func (*DisjunctiveProof) ProtoMessage()    {}
func (*DisjunctiveProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_3687e745eb88f81a, []int{2}
}

func (m *DisjunctiveProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisjunctiveProof.Unmarshal(m, b)
}
func (m *DisjunctiveProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisjunctiveProof.Marshal(b, m, deterministic)
}
func (m *DisjunctiveProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisjunctiveProof.Merge(m, src)
}
func (m *DisjunctiveProof) XXX_Size() int {
	return xxx_messageInfo_DisjunctiveProof.Size(m)
}
func (m *DisjunctiveProof) XXX_DiscardUnknown() {
	xxx_messageInfo_DisjunctiveProof.DiscardUnknown(m)
}

var xxx_messageInfo_DisjunctiveProof proto.InternalMessageInfo

func (m *DisjunctiveProof) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *DisjunctiveProof) GetG1() []byte {
	if m != nil {
		return m.G1
	}
	return nil
}

func (m *DisjunctiveProof) GetY1() []byte {
	if m != nil {
		return m.Y1
	}
	return nil
}

func (m *DisjunctiveProof) GetG2() []byte {
	if m != nil {
		return m.G2
	}
	return nil
}

func (m *DisjunctiveProof) GetY2() []byte {
	if m != nil {
		return m.Y2
	}
	return nil
}

type RangeProof struct {
	Comm []byte        `protobuf:"bytes,1,opt,name=Comm,proto3" json:"Comm,omitempty"`
	A    []byte        `protobuf:"bytes,2,opt,name=A,proto3" json:"A,omitempty"`
	S    []byte        `protobuf:"bytes,3,opt,name=S,proto3" json:"S,omitempty"`
	T1   []byte        `protobuf:"bytes,4,opt,name=T1,proto3" json:"T1,omitempty"`
	T2   []byte        `protobuf:"bytes,5,opt,name=T2,proto3" json:"T2,omitempty"`
	Tau  []byte        `protobuf:"bytes,6,opt,name=Tau,proto3" json:"Tau,omitempty"`
	Th   []byte        `protobuf:"bytes,7,opt,name=Th,proto3" json:"Th,omitempty"`
	Mu   []byte        `protobuf:"bytes,8,opt,name=Mu,proto3" json:"Mu,omitempty"`
	IPP  *InnerProdArg `protobuf:"bytes,9,opt,name=IPP,proto3" json:"IPP,omitempty"`
	// challenges
	Cy                   []byte   `protobuf:"bytes,10,opt,name=Cy,proto3" json:"Cy,omitempty"`
	Cz                   []byte   `protobuf:"bytes,11,opt,name=Cz,proto3" json:"Cz,omitempty"`
	Cx                   []byte   `protobuf:"bytes,12,opt,name=Cx,proto3" json:"Cx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RangeProof) Reset()         { *m = RangeProof{} }
func (m *RangeProof) String() string { return proto.CompactTextString(m) }
func (*RangeProof) ProtoMessage()    {}
func (*RangeProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_3687e745eb88f81a, []int{3}
}

func (m *RangeProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RangeProof.Unmarshal(m, b)
}
func (m *RangeProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RangeProof.Marshal(b, m, deterministic)
}
func (m *RangeProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeProof.Merge(m, src)
}
func (m *RangeProof) XXX_Size() int {
	return xxx_messageInfo_RangeProof.Size(m)
}
func (m *RangeProof) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeProof.DiscardUnknown(m)
}

var xxx_messageInfo_RangeProof proto.InternalMessageInfo

func (m *RangeProof) GetComm() []byte {
	if m != nil {
		return m.Comm
	}
	return nil
}

func (m *RangeProof) GetA() []byte {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *RangeProof) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *RangeProof) GetT1() []byte {
	if m != nil {
		return m.T1
	}
	return nil
}

func (m *RangeProof) GetT2() []byte {
	if m != nil {
		return m.T2
	}
	return nil
}

func (m *RangeProof) GetTau() []byte {
	if m != nil {
		return m.Tau
	}
	return nil
}

func (m *RangeProof) GetTh() []byte {
	if m != nil {
		return m.Th
	}
	return nil
}

func (m *RangeProof) GetMu() []byte {
	if m != nil {
		return m.Mu
	}
	return nil
}

func (m *RangeProof) GetIPP() *InnerProdArg {
	if m != nil {
		return m.IPP
	}
	return nil
}

func (m *RangeProof) GetCy() []byte {
	if m != nil {
		return m.Cy
	}
	return nil
}

func (m *RangeProof) GetCz() []byte {
	if m != nil {
		return m.Cz
	}
	return nil
}

func (m *RangeProof) GetCx() []byte {
	if m != nil {
		return m.Cx
	}
	return nil
}

type InnerProdArg struct {
	L []byte `protobuf:"bytes,1,opt,name=L,proto3" json:"L,omitempty"`
	R []byte `protobuf:"bytes,2,opt,name=R,proto3" json:"R,omitempty"`
	A []byte `protobuf:"bytes,3,opt,name=A,proto3" json:"A,omitempty"`
	B []byte `protobuf:"bytes,4,opt,name=B,proto3" json:"B,omitempty"`
	// 只要使用repeated标记类型定义，就表示数组类型。 https://www.tizi365.com/archives/378.html
	Challenges           [][]byte `protobuf:"bytes,5,rep,name=Challenges,proto3" json:"Challenges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerProdArg) Reset()         { *m = InnerProdArg{} }
func (m *InnerProdArg) String() string { return proto.CompactTextString(m) }
func (*InnerProdArg) ProtoMessage()    {}
func (*InnerProdArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_3687e745eb88f81a, []int{4}
}

func (m *InnerProdArg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerProdArg.Unmarshal(m, b)
}
func (m *InnerProdArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerProdArg.Marshal(b, m, deterministic)
}
func (m *InnerProdArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerProdArg.Merge(m, src)
}
func (m *InnerProdArg) XXX_Size() int {
	return xxx_messageInfo_InnerProdArg.Size(m)
}
func (m *InnerProdArg) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerProdArg.DiscardUnknown(m)
}

var xxx_messageInfo_InnerProdArg proto.InternalMessageInfo

func (m *InnerProdArg) GetL() []byte {
	if m != nil {
		return m.L
	}
	return nil
}

func (m *InnerProdArg) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *InnerProdArg) GetA() []byte {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *InnerProdArg) GetB() []byte {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *InnerProdArg) GetChallenges() [][]byte {
	if m != nil {
		return m.Challenges
	}
	return nil
}

func init() {
	proto.RegisterType((*Zkrow)(nil), "zkrow_package.zkrow")
	proto.RegisterMapType((map[string]*OrgColumn)(nil), "zkrow_package.zkrow.ColumnsEntry")
	proto.RegisterType((*OrgColumn)(nil), "zkrow_package.OrgColumn")
	proto.RegisterType((*DisjunctiveProof)(nil), "zkrow_package.DisjunctiveProof")
	proto.RegisterType((*RangeProof)(nil), "zkrow_package.RangeProof")
	proto.RegisterType((*InnerProdArg)(nil), "zkrow_package.InnerProdArg")
}

func init() { proto.RegisterFile("zkrow.proto", fileDescriptor_3687e745eb88f81a) }

var fileDescriptor_3687e745eb88f81a = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0xd4, 0xae, 0x9b, 0x7e, 0xbc, 0xb8, 0x28, 0x5a, 0x71, 0x58, 0x40, 0x82, 0x60, 0x71, 0x08,
	0x48, 0x44, 0x8a, 0x7b, 0x41, 0x70, 0x4a, 0x5c, 0x0e, 0x95, 0x5a, 0x11, 0x6d, 0x2d, 0xae, 0xc8,
	0x4d, 0xb6, 0x8e, 0x89, 0xed, 0xb5, 0xd6, 0x76, 0xa9, 0xf3, 0x5f, 0xb9, 0x71, 0xe5, 0x3f, 0xa0,
	0x7d, 0xde, 0x34, 0x5f, 0x17, 0x6e, 0x3b, 0xb3, 0xf3, 0xde, 0x78, 0x46, 0x2b, 0x43, 0x77, 0xb5,
	0xd4, 0xea, 0xd7, 0xb0, 0xd0, 0xaa, 0x52, 0xec, 0x1c, 0xc1, 0x8f, 0x22, 0x9a, 0x2d, 0xa3, 0x58,
	0x7a, 0x7f, 0x08, 0x74, 0x90, 0x61, 0x5f, 0xe0, 0x64, 0xa6, 0xd2, 0x3a, 0xcb, 0x4b, 0x4e, 0xfa,
	0xce, 0xa0, 0xeb, 0xbf, 0x1d, 0xee, 0x48, 0x5b, 0x34, 0x0c, 0x5a, 0xcd, 0xd7, 0xbc, 0xd2, 0x8d,
	0x58, 0x4f, 0xb0, 0x77, 0x70, 0x9e, 0x94, 0xdf, 0xa3, 0x34, 0x99, 0x4f, 0xa2, 0x34, 0x50, 0x9a,
	0xd3, 0x3e, 0x19, 0x9c, 0x8a, 0x5d, 0x92, 0x79, 0xe0, 0x5a, 0x62, 0x5c, 0x96, 0xb2, 0xe2, 0x0e,
	0x8a, 0x76, 0xb8, 0x97, 0x21, 0xb8, 0xdb, 0x16, 0xac, 0x07, 0xce, 0x52, 0x36, 0x9c, 0xf4, 0xc9,
	0xe0, 0x4c, 0x98, 0x23, 0x1b, 0x42, 0xe7, 0x21, 0x4a, 0x6b, 0x89, 0x1e, 0x5d, 0x9f, 0xef, 0x7d,
	0xe6, 0x37, 0x1d, 0xb7, 0x0b, 0x44, 0x2b, 0xfb, 0x4c, 0x3f, 0x11, 0xef, 0x37, 0x85, 0xb3, 0xa7,
	0x0b, 0xf6, 0x1a, 0x60, 0xa6, 0xb2, 0x2c, 0xa9, 0x32, 0x99, 0x57, 0xb8, 0xda, 0x15, 0x5b, 0x8c,
	0xb9, 0x8f, 0xea, 0x79, 0x52, 0x85, 0x6a, 0x29, 0x73, 0xb4, 0x71, 0xc5, 0x16, 0x73, 0x98, 0xd6,
	0xf9, 0x9f, 0xb4, 0x47, 0x87, 0x69, 0x8d, 0x13, 0xae, 0x9c, 0xea, 0x24, 0x93, 0xbc, 0xd3, 0x3a,
	0x6d, 0x18, 0xf6, 0x01, 0x7a, 0x88, 0x2e, 0x55, 0x7d, 0x97, 0xca, 0x56, 0x75, 0x8c, 0xaa, 0x03,
	0x9e, 0xbd, 0x07, 0xaa, 0x0b, 0x7e, 0x82, 0xa5, 0xbc, 0xd8, 0x2b, 0x45, 0x44, 0x79, 0x2c, 0xa7,
	0x5a, 0xa9, 0x7b, 0x41, 0x75, 0xc1, 0x2e, 0xe0, 0x68, 0xbe, 0x5a, 0x16, 0xfc, 0x14, 0xc5, 0x6f,
	0xf6, 0xc4, 0x97, 0x49, 0xf9, 0xb3, 0xce, 0x67, 0x55, 0xf2, 0x60, 0x47, 0x50, 0xcc, 0x5c, 0x20,
	0x25, 0x3f, 0x43, 0x73, 0x52, 0x1a, 0x54, 0x71, 0x68, 0x51, 0xe5, 0x2d, 0xa0, 0xb7, 0x3f, 0xc5,
	0x9e, 0x43, 0xa7, 0x30, 0x07, 0x5b, 0x70, 0x0b, 0xd8, 0x33, 0xa0, 0xf1, 0xc8, 0x76, 0x4a, 0xe3,
	0x91, 0xc1, 0xcd, 0x08, 0x0b, 0x74, 0x05, 0x6d, 0x10, 0xc7, 0x3e, 0x76, 0x65, 0xee, 0x7d, 0xbc,
	0xf7, 0x6d, 0x33, 0xb4, 0xf1, 0xbd, 0xbf, 0x04, 0x60, 0x93, 0x86, 0x31, 0x38, 0x0a, 0x54, 0x96,
	0x59, 0x0f, 0x3c, 0x9b, 0x4f, 0x1b, 0x5b, 0x07, 0x32, 0x36, 0xe8, 0xd6, 0xee, 0x27, 0xb7, 0x66,
	0x5d, 0x38, 0x5a, 0xaf, 0x0f, 0xd1, 0x2e, 0x7c, 0x5a, 0x1f, 0xfa, 0xe6, 0xb9, 0x85, 0x51, 0x6d,
	0x3b, 0x36, 0x47, 0x54, 0x2c, 0xb0, 0x56, 0xa3, 0x58, 0x18, 0x7c, 0x53, 0x63, 0x73, 0xae, 0xa0,
	0x37, 0x35, 0xfb, 0x08, 0xce, 0xd5, 0x74, 0x8a, 0xc5, 0x74, 0xfd, 0x57, 0x7b, 0x55, 0x5e, 0xe5,
	0xb9, 0xd4, 0x53, 0xad, 0xe6, 0x63, 0x1d, 0x0b, 0xa3, 0x33, 0xe3, 0x41, 0x63, 0x8b, 0xa3, 0x41,
	0x83, 0x78, 0xc5, 0xbb, 0x16, 0xaf, 0x10, 0x3f, 0x72, 0xd7, 0xe2, 0x47, 0xef, 0x1e, 0xdc, 0xed,
	0x25, 0x26, 0xce, 0xb5, 0x4d, 0x4b, 0xae, 0x0d, 0x12, 0xeb, 0xa8, 0xa2, 0x0d, 0xee, 0x6c, 0x05,
	0x9f, 0xd8, 0xa4, 0x64, 0x62, 0x5e, 0x5a, 0xb0, 0x88, 0xd2, 0x54, 0xe6, 0xb1, 0x2c, 0x79, 0xa7,
	0xef, 0x98, 0x97, 0xb6, 0x61, 0xee, 0x8e, 0xf1, 0xf7, 0x70, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x7c, 0x6a, 0x97, 0xbb, 0x2d, 0x04, 0x00, 0x00,
}
