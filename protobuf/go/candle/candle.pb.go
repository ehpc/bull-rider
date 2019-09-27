// Code generated by protoc-gen-go. DO NOT EDIT.
// source: candle/candle.proto

package candle

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

// Candle is an exhange candle
type Candle struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Pair                 string   `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	Interval             string   `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	OpenTime             int64    `protobuf:"varint,4,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	CloseTime            int64    `protobuf:"varint,5,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	Open                 float64  `protobuf:"fixed64,6,opt,name=open,proto3" json:"open,omitempty"`
	Close                float64  `protobuf:"fixed64,7,opt,name=close,proto3" json:"close,omitempty"`
	High                 float64  `protobuf:"fixed64,8,opt,name=high,proto3" json:"high,omitempty"`
	Low                  float64  `protobuf:"fixed64,9,opt,name=low,proto3" json:"low,omitempty"`
	Volume               float64  `protobuf:"fixed64,10,opt,name=volume,proto3" json:"volume,omitempty"`
	QuoteVolume          float64  `protobuf:"fixed64,11,opt,name=quote_volume,json=quoteVolume,proto3" json:"quote_volume,omitempty"`
	TradesCount          int64    `protobuf:"varint,12,opt,name=trades_count,json=tradesCount,proto3" json:"trades_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candle) Reset()         { *m = Candle{} }
func (m *Candle) String() string { return proto.CompactTextString(m) }
func (*Candle) ProtoMessage()    {}
func (*Candle) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f877bc93b8156de, []int{0}
}

func (m *Candle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candle.Unmarshal(m, b)
}
func (m *Candle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candle.Marshal(b, m, deterministic)
}
func (m *Candle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candle.Merge(m, src)
}
func (m *Candle) XXX_Size() int {
	return xxx_messageInfo_Candle.Size(m)
}
func (m *Candle) XXX_DiscardUnknown() {
	xxx_messageInfo_Candle.DiscardUnknown(m)
}

var xxx_messageInfo_Candle proto.InternalMessageInfo

func (m *Candle) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Candle) GetPair() string {
	if m != nil {
		return m.Pair
	}
	return ""
}

func (m *Candle) GetInterval() string {
	if m != nil {
		return m.Interval
	}
	return ""
}

func (m *Candle) GetOpenTime() int64 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

func (m *Candle) GetCloseTime() int64 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func (m *Candle) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Candle) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Candle) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Candle) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Candle) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Candle) GetQuoteVolume() float64 {
	if m != nil {
		return m.QuoteVolume
	}
	return 0
}

func (m *Candle) GetTradesCount() int64 {
	if m != nil {
		return m.TradesCount
	}
	return 0
}

// Candles is an array of candles
type Candles struct {
	Candles              []*Candles `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Candles) Reset()         { *m = Candles{} }
func (m *Candles) String() string { return proto.CompactTextString(m) }
func (*Candles) ProtoMessage()    {}
func (*Candles) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f877bc93b8156de, []int{1}
}

func (m *Candles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candles.Unmarshal(m, b)
}
func (m *Candles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candles.Marshal(b, m, deterministic)
}
func (m *Candles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candles.Merge(m, src)
}
func (m *Candles) XXX_Size() int {
	return xxx_messageInfo_Candles.Size(m)
}
func (m *Candles) XXX_DiscardUnknown() {
	xxx_messageInfo_Candles.DiscardUnknown(m)
}

var xxx_messageInfo_Candles proto.InternalMessageInfo

func (m *Candles) GetCandles() []*Candles {
	if m != nil {
		return m.Candles
	}
	return nil
}

func init() {
	proto.RegisterType((*Candle)(nil), "candle.Candle")
	proto.RegisterType((*Candles)(nil), "candle.Candles")
}

func init() { proto.RegisterFile("candle/candle.proto", fileDescriptor_4f877bc93b8156de) }

var fileDescriptor_4f877bc93b8156de = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x91, 0xc1, 0x4e, 0xb3, 0x40,
	0x14, 0x85, 0x43, 0x69, 0x69, 0xb9, 0x34, 0xf9, 0xff, 0x5c, 0x8d, 0x99, 0x68, 0x4c, 0xb0, 0x2b,
	0x5c, 0xb4, 0x24, 0xea, 0x13, 0xd8, 0x37, 0x20, 0xc6, 0x85, 0x9b, 0x06, 0xe8, 0x15, 0x26, 0x19,
	0x18, 0x84, 0xa1, 0xfa, 0x74, 0x3e, 0x9b, 0x99, 0x3b, 0xb4, 0x2b, 0xce, 0xf9, 0xce, 0xb7, 0x20,
	0x77, 0xe0, 0xaa, 0xcc, 0xdb, 0xa3, 0xa2, 0xd4, 0x7d, 0x76, 0x5d, 0xaf, 0x8d, 0xc6, 0xc0, 0xb5,
	0xcd, 0xef, 0x0c, 0x82, 0x3d, 0x47, 0xbc, 0x85, 0x15, 0xfd, 0x94, 0x75, 0xde, 0x56, 0x24, 0xbc,
	0xd8, 0x4b, 0xc2, 0xec, 0xd2, 0x11, 0x61, 0xde, 0xe5, 0xb2, 0x17, 0x33, 0xe6, 0x9c, 0xad, 0x2f,
	0x5b, 0x43, 0xfd, 0x29, 0x57, 0xc2, 0x77, 0xfe, 0xb9, 0xe3, 0x1d, 0x84, 0xba, 0xa3, 0xf6, 0x60,
	0x64, 0x43, 0x62, 0x1e, 0x7b, 0x89, 0x9f, 0xad, 0x2c, 0x78, 0x93, 0x0d, 0xe1, 0x3d, 0x40, 0xa9,
	0xf4, 0x40, 0x6e, 0x5d, 0xf0, 0x1a, 0x32, 0xe1, 0x19, 0x61, 0x6e, 0x55, 0x11, 0xc4, 0x5e, 0xe2,
	0x65, 0x9c, 0xf1, 0x1a, 0x16, 0x2c, 0x88, 0x25, 0x43, 0x57, 0xac, 0x59, 0xcb, 0xaa, 0x16, 0x2b,
	0x67, 0xda, 0x8c, 0xff, 0xc1, 0x57, 0xfa, 0x5b, 0x84, 0x8c, 0x6c, 0xc4, 0x1b, 0x08, 0x4e, 0x5a,
	0x8d, 0x0d, 0x09, 0x60, 0x38, 0x35, 0x7c, 0x80, 0xf5, 0xd7, 0xa8, 0x0d, 0x1d, 0xa6, 0x35, 0xe2,
	0x35, 0x62, 0xf6, 0x7e, 0x51, 0x4c, 0x9f, 0x1f, 0x69, 0x38, 0x94, 0x7a, 0x6c, 0x8d, 0x58, 0xf3,
	0xbf, 0x46, 0x8e, 0xed, 0x2d, 0xda, 0xbc, 0xc0, 0xd2, 0xdd, 0x6f, 0xc0, 0x47, 0x58, 0xba, 0xab,
	0x0e, 0xc2, 0x8b, 0xfd, 0x24, 0x7a, 0xfa, 0xb7, 0x9b, 0x6e, 0x3e, 0x19, 0xd9, 0x79, 0x7f, 0x4d,
	0x3f, 0xb6, 0x95, 0x34, 0xf5, 0x58, 0xec, 0x4a, 0xdd, 0xa4, 0x54, 0x77, 0x65, 0x5a, 0x8c, 0x4a,
	0x6d, 0x7b, 0x79, 0xa4, 0x3e, 0xe5, 0x27, 0x2a, 0xc6, 0xcf, 0xb4, 0xd2, 0xd3, 0xab, 0x15, 0x01,
	0xb3, 0xe7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0x6d, 0x5b, 0xd9, 0xcd, 0x01, 0x00, 0x00,
}
