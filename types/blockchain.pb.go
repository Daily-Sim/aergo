// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package types // import "github.com/aergoio/aergo/types"

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

type Block struct {
	Hash                 []byte       `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Header               *BlockHeader `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body                 *BlockBody   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_2f1b5e0575d45901, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetBody() *BlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type BlockHeader struct {
	PrevBlockHash        []byte   `protobuf:"bytes,1,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	BlockNo              uint64   `protobuf:"varint,2,opt,name=blockNo,proto3" json:"blockNo,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BlocksRootHash       []byte   `protobuf:"bytes,4,opt,name=blocksRootHash,proto3" json:"blocksRootHash,omitempty"`
	TxsRootHash          []byte   `protobuf:"bytes,5,opt,name=txsRootHash,proto3" json:"txsRootHash,omitempty"`
	PubKey               []byte   `protobuf:"bytes,6,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Sign                 []byte   `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_2f1b5e0575d45901, []int{1}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *BlockHeader) GetBlockNo() uint64 {
	if m != nil {
		return m.BlockNo
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetBlocksRootHash() []byte {
	if m != nil {
		return m.BlocksRootHash
	}
	return nil
}

func (m *BlockHeader) GetTxsRootHash() []byte {
	if m != nil {
		return m.TxsRootHash
	}
	return nil
}

func (m *BlockHeader) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *BlockHeader) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type BlockBody struct {
	Txs                  []*Tx    `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockBody) Reset()         { *m = BlockBody{} }
func (m *BlockBody) String() string { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()    {}
func (*BlockBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_2f1b5e0575d45901, []int{2}
}
func (m *BlockBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockBody.Unmarshal(m, b)
}
func (m *BlockBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockBody.Marshal(b, m, deterministic)
}
func (dst *BlockBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockBody.Merge(dst, src)
}
func (m *BlockBody) XXX_Size() int {
	return xxx_messageInfo_BlockBody.Size(m)
}
func (m *BlockBody) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockBody.DiscardUnknown(m)
}

var xxx_messageInfo_BlockBody proto.InternalMessageInfo

func (m *BlockBody) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

type TxList struct {
	Txs                  []*Tx    `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxList) Reset()         { *m = TxList{} }
func (m *TxList) String() string { return proto.CompactTextString(m) }
func (*TxList) ProtoMessage()    {}
func (*TxList) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_2f1b5e0575d45901, []int{3}
}
func (m *TxList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxList.Unmarshal(m, b)
}
func (m *TxList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxList.Marshal(b, m, deterministic)
}
func (dst *TxList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxList.Merge(dst, src)
}
func (m *TxList) XXX_Size() int {
	return xxx_messageInfo_TxList.Size(m)
}
func (m *TxList) XXX_DiscardUnknown() {
	xxx_messageInfo_TxList.DiscardUnknown(m)
}

var xxx_messageInfo_TxList proto.InternalMessageInfo

func (m *TxList) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Tx struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Body                 *TxBody  `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_2f1b5e0575d45901, []int{4}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tx.Unmarshal(m, b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
}
func (dst *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(dst, src)
}
func (m *Tx) XXX_Size() int {
	return xxx_messageInfo_Tx.Size(m)
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Tx) GetBody() *TxBody {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Tx) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type TxBody struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Account              []byte   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Recipient            []byte   `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount               uint64   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Limit                uint64   `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Price                uint64   `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	Sign                 []byte   `protobuf:"bytes,8,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxBody) Reset()         { *m = TxBody{} }
func (m *TxBody) String() string { return proto.CompactTextString(m) }
func (*TxBody) ProtoMessage()    {}
func (*TxBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_2f1b5e0575d45901, []int{5}
}
func (m *TxBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxBody.Unmarshal(m, b)
}
func (m *TxBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxBody.Marshal(b, m, deterministic)
}
func (dst *TxBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxBody.Merge(dst, src)
}
func (m *TxBody) XXX_Size() int {
	return xxx_messageInfo_TxBody.Size(m)
}
func (m *TxBody) XXX_DiscardUnknown() {
	xxx_messageInfo_TxBody.DiscardUnknown(m)
}

var xxx_messageInfo_TxBody proto.InternalMessageInfo

func (m *TxBody) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TxBody) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *TxBody) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *TxBody) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxBody) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TxBody) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TxBody) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TxBody) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type TxIdx struct {
	BlockHash            []byte   `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Idx                  int32    `protobuf:"varint,2,opt,name=idx,proto3" json:"idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxIdx) Reset()         { *m = TxIdx{} }
func (m *TxIdx) String() string { return proto.CompactTextString(m) }
func (*TxIdx) ProtoMessage()    {}
func (*TxIdx) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_2f1b5e0575d45901, []int{6}
}
func (m *TxIdx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxIdx.Unmarshal(m, b)
}
func (m *TxIdx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxIdx.Marshal(b, m, deterministic)
}
func (dst *TxIdx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxIdx.Merge(dst, src)
}
func (m *TxIdx) XXX_Size() int {
	return xxx_messageInfo_TxIdx.Size(m)
}
func (m *TxIdx) XXX_DiscardUnknown() {
	xxx_messageInfo_TxIdx.DiscardUnknown(m)
}

var xxx_messageInfo_TxIdx proto.InternalMessageInfo

func (m *TxIdx) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *TxIdx) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

type State struct {
	Account              []byte   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Nonce                uint64   `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Balance              uint64   `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_2f1b5e0575d45901, []int{7}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *State) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *State) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*Block)(nil), "types.Block")
	proto.RegisterType((*BlockHeader)(nil), "types.BlockHeader")
	proto.RegisterType((*BlockBody)(nil), "types.BlockBody")
	proto.RegisterType((*TxList)(nil), "types.TxList")
	proto.RegisterType((*Tx)(nil), "types.Tx")
	proto.RegisterType((*TxBody)(nil), "types.TxBody")
	proto.RegisterType((*TxIdx)(nil), "types.TxIdx")
	proto.RegisterType((*State)(nil), "types.State")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_blockchain_2f1b5e0575d45901) }

var fileDescriptor_blockchain_2f1b5e0575d45901 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5f, 0xab, 0xd3, 0x30,
	0x18, 0xc6, 0xe9, 0x9a, 0xf6, 0xb8, 0x77, 0x3b, 0x32, 0x82, 0x48, 0x41, 0x91, 0x5a, 0x8e, 0x32,
	0xbc, 0xd8, 0x81, 0xe3, 0x85, 0xf7, 0xbb, 0x52, 0x14, 0xc5, 0xb8, 0x2b, 0xef, 0xd2, 0x36, 0xac,
	0xc1, 0xb5, 0x09, 0x6d, 0x26, 0xa9, 0x9f, 0xd0, 0xaf, 0xe1, 0x37, 0x91, 0xbc, 0xe9, 0xd6, 0x6e,
	0x88, 0x57, 0xcb, 0xf3, 0xe4, 0xfd, 0xfb, 0x4b, 0x07, 0xab, 0xfc, 0xa0, 0x8a, 0x1f, 0x45, 0xc5,
	0x65, 0xb3, 0xd1, 0xad, 0x32, 0x8a, 0x46, 0xa6, 0xd7, 0xa2, 0xcb, 0x6a, 0x88, 0xb6, 0xee, 0x8a,
	0x52, 0x20, 0x15, 0xef, 0xaa, 0x24, 0x48, 0x83, 0xf5, 0x92, 0xe1, 0x99, 0xbe, 0x81, 0xb8, 0x12,
	0xbc, 0x14, 0x6d, 0x32, 0x4b, 0x83, 0xf5, 0xe2, 0x81, 0x6e, 0x30, 0x69, 0x83, 0x19, 0xef, 0xf1,
	0x86, 0x0d, 0x11, 0xf4, 0x0e, 0x48, 0xae, 0xca, 0x3e, 0x09, 0x31, 0x72, 0x35, 0x8d, 0xdc, 0xaa,
	0xb2, 0x67, 0x78, 0x9b, 0xfd, 0x09, 0x60, 0x31, 0xc9, 0xa6, 0x77, 0x70, 0xab, 0x5b, 0xf1, 0xd3,
	0x5b, 0x63, 0xfb, 0x4b, 0x93, 0x26, 0x70, 0x83, 0xf3, 0x7f, 0x56, 0x38, 0x08, 0x61, 0x27, 0x49,
	0x9f, 0xc3, 0xdc, 0xc8, 0x5a, 0x74, 0x86, 0xd7, 0x1a, 0x5b, 0x87, 0x6c, 0x34, 0xe8, 0x6b, 0x78,
	0x8c, 0x81, 0x1d, 0x53, 0xca, 0x60, 0x79, 0x82, 0xe5, 0xaf, 0x5c, 0x9a, 0xc2, 0xc2, 0xd8, 0x31,
	0x28, 0xc2, 0xa0, 0xa9, 0x45, 0x9f, 0x42, 0xac, 0x8f, 0xf9, 0x47, 0xd1, 0x27, 0x31, 0x5e, 0x0e,
	0xca, 0x51, 0xeb, 0xe4, 0xbe, 0x49, 0x6e, 0x3c, 0x35, 0x77, 0xce, 0xd6, 0x30, 0x3f, 0xaf, 0x4d,
	0x9f, 0x41, 0x68, 0x6c, 0x97, 0x04, 0x69, 0xb8, 0x5e, 0x3c, 0xcc, 0x07, 0x2a, 0x3b, 0xcb, 0x9c,
	0x9b, 0xbd, 0x82, 0x78, 0x67, 0x3f, 0xc9, 0xce, 0xfc, 0x3f, 0xec, 0x0b, 0xcc, 0x76, 0xf6, 0x9f,
	0x0f, 0xf4, 0x72, 0x80, 0xee, 0x9f, 0xe7, 0xf6, 0x9c, 0x37, 0x12, 0xf7, 0x13, 0xfe, 0x12, 0x08,
	0x87, 0x30, 0x3c, 0x67, 0xbf, 0x03, 0xd7, 0x18, 0xe7, 0x7b, 0x02, 0x51, 0xa3, 0x9a, 0x42, 0x60,
	0x59, 0xc2, 0xbc, 0x70, 0xc0, 0x79, 0x51, 0xa8, 0x63, 0x63, 0xb0, 0xf4, 0x92, 0x9d, 0xa4, 0x03,
	0xde, 0x8a, 0x42, 0x6a, 0x29, 0x1a, 0x83, 0x35, 0x97, 0x6c, 0x34, 0x1c, 0x26, 0x5e, 0x63, 0x1a,
	0xc1, 0x72, 0x83, 0x72, 0xf5, 0x34, 0xef, 0x0f, 0x8a, 0x97, 0x03, 0xdc, 0x93, 0x74, 0xfd, 0x0f,
	0xb2, 0x96, 0x06, 0xb9, 0x12, 0xe6, 0x85, 0x73, 0x75, 0x2b, 0x0b, 0x81, 0x5c, 0x09, 0xf3, 0xe2,
	0x0c, 0xfb, 0xd1, 0x04, 0xf6, 0x3b, 0x88, 0x76, 0xf6, 0x43, 0x69, 0xdd, 0x60, 0xf9, 0xd5, 0x57,
	0x34, 0x1a, 0x74, 0x05, 0xa1, 0x2c, 0x2d, 0x2e, 0x13, 0x31, 0x77, 0xcc, 0xbe, 0x42, 0xf4, 0xcd,
	0x70, 0x73, 0xb1, 0x6b, 0x70, 0xb9, 0xeb, 0x99, 0xcd, 0xec, 0x8a, 0x4d, 0xce, 0x0f, 0xdc, 0xf9,
	0xe1, 0xf0, 0x31, 0x7a, 0xb9, 0x4d, 0xbf, 0xbf, 0xd8, 0x4b, 0x53, 0x1d, 0xf3, 0x4d, 0xa1, 0xea,
	0x7b, 0x2e, 0xda, 0xbd, 0x92, 0xca, 0xff, 0xde, 0xe3, 0xcb, 0xe4, 0x31, 0xfe, 0xf7, 0xde, 0xfe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xeb, 0x31, 0xa4, 0x8f, 0x03, 0x00, 0x00,
}
