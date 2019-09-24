// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pair/pair.proto

package pair

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

// Pair is a trading pair
type Pair int32

const (
	Pair_BTCUSDT    Pair = 0
	Pair_ETHUSDT    Pair = 1
	Pair_BNBUSDT    Pair = 2
	Pair_BCCUSDT    Pair = 3
	Pair_NEOUSDT    Pair = 4
	Pair_LTCUSDT    Pair = 5
	Pair_QTUMUSDT   Pair = 6
	Pair_ADAUSDT    Pair = 7
	Pair_XRPUSDT    Pair = 8
	Pair_EOSUSDT    Pair = 9
	Pair_TUSDUSDT   Pair = 10
	Pair_IOTAUSDT   Pair = 11
	Pair_XLMUSDT    Pair = 12
	Pair_ONTUSDT    Pair = 13
	Pair_TRXUSDT    Pair = 14
	Pair_ETCUSDT    Pair = 15
	Pair_ICXUSDT    Pair = 16
	Pair_VENUSDT    Pair = 17
	Pair_NULSUSDT   Pair = 18
	Pair_VETUSDT    Pair = 19
	Pair_PAXUSDT    Pair = 20
	Pair_BCHABCUSDT Pair = 21
	Pair_BCHSVUSDT  Pair = 22
	Pair_USDCUSDT   Pair = 23
	Pair_LINKUSDT   Pair = 24
	Pair_WAVESUSDT  Pair = 25
	Pair_BTTUSDT    Pair = 26
	Pair_USDSUSDT   Pair = 27
	Pair_ONGUSDT    Pair = 28
	Pair_HOTUSDT    Pair = 29
	Pair_ZILUSDT    Pair = 30
	Pair_ZRXUSDT    Pair = 31
	Pair_FETUSDT    Pair = 32
	Pair_BATUSDT    Pair = 33
	Pair_XMRUSDT    Pair = 34
	Pair_ZECUSDT    Pair = 35
	Pair_IOSTUSDT   Pair = 36
	Pair_CELRUSDT   Pair = 37
	Pair_DASHUSDT   Pair = 38
	Pair_NANOUSDT   Pair = 39
	Pair_OMGUSDT    Pair = 40
	Pair_THETAUSDT  Pair = 41
	Pair_ENJUSDT    Pair = 42
	Pair_MITHUSDT   Pair = 43
	Pair_MATICUSDT  Pair = 44
	Pair_ATOMUSDT   Pair = 45
	Pair_TFUELUSDT  Pair = 46
	Pair_ONEUSDT    Pair = 47
	Pair_FTMUSDT    Pair = 48
	Pair_ALGOUSDT   Pair = 49
	Pair_USDSBUSDT  Pair = 50
	Pair_GTOUSDT    Pair = 51
	Pair_ERDUSDT    Pair = 52
	Pair_DOGEUSDT   Pair = 53
	Pair_DUSKUSDT   Pair = 54
	Pair_ANKRUSDT   Pair = 55
	Pair_WINUSDT    Pair = 56
	Pair_COSUSDT    Pair = 57
	Pair_NPXSUSDT   Pair = 58
	Pair_COCOSUSDT  Pair = 59
	Pair_MTLUSDT    Pair = 60
	Pair_TOMOUSDT   Pair = 61
	Pair_PERLUSDT   Pair = 62
	Pair_DENTUSDT   Pair = 63
	Pair_MFTUSDT    Pair = 64
	Pair_KEYUSDT    Pair = 65
	Pair_STORMUSDT  Pair = 66
	Pair_DOCKUSDT   Pair = 67
	Pair_WANUSDT    Pair = 68
	Pair_FUNUSDT    Pair = 69
	Pair_CVCUSDT    Pair = 70
	Pair_CHZUSDT    Pair = 71
	Pair_BANDUSDT   Pair = 72
)

var Pair_name = map[int32]string{
	0:  "BTCUSDT",
	1:  "ETHUSDT",
	2:  "BNBUSDT",
	3:  "BCCUSDT",
	4:  "NEOUSDT",
	5:  "LTCUSDT",
	6:  "QTUMUSDT",
	7:  "ADAUSDT",
	8:  "XRPUSDT",
	9:  "EOSUSDT",
	10: "TUSDUSDT",
	11: "IOTAUSDT",
	12: "XLMUSDT",
	13: "ONTUSDT",
	14: "TRXUSDT",
	15: "ETCUSDT",
	16: "ICXUSDT",
	17: "VENUSDT",
	18: "NULSUSDT",
	19: "VETUSDT",
	20: "PAXUSDT",
	21: "BCHABCUSDT",
	22: "BCHSVUSDT",
	23: "USDCUSDT",
	24: "LINKUSDT",
	25: "WAVESUSDT",
	26: "BTTUSDT",
	27: "USDSUSDT",
	28: "ONGUSDT",
	29: "HOTUSDT",
	30: "ZILUSDT",
	31: "ZRXUSDT",
	32: "FETUSDT",
	33: "BATUSDT",
	34: "XMRUSDT",
	35: "ZECUSDT",
	36: "IOSTUSDT",
	37: "CELRUSDT",
	38: "DASHUSDT",
	39: "NANOUSDT",
	40: "OMGUSDT",
	41: "THETAUSDT",
	42: "ENJUSDT",
	43: "MITHUSDT",
	44: "MATICUSDT",
	45: "ATOMUSDT",
	46: "TFUELUSDT",
	47: "ONEUSDT",
	48: "FTMUSDT",
	49: "ALGOUSDT",
	50: "USDSBUSDT",
	51: "GTOUSDT",
	52: "ERDUSDT",
	53: "DOGEUSDT",
	54: "DUSKUSDT",
	55: "ANKRUSDT",
	56: "WINUSDT",
	57: "COSUSDT",
	58: "NPXSUSDT",
	59: "COCOSUSDT",
	60: "MTLUSDT",
	61: "TOMOUSDT",
	62: "PERLUSDT",
	63: "DENTUSDT",
	64: "MFTUSDT",
	65: "KEYUSDT",
	66: "STORMUSDT",
	67: "DOCKUSDT",
	68: "WANUSDT",
	69: "FUNUSDT",
	70: "CVCUSDT",
	71: "CHZUSDT",
	72: "BANDUSDT",
}

var Pair_value = map[string]int32{
	"BTCUSDT":    0,
	"ETHUSDT":    1,
	"BNBUSDT":    2,
	"BCCUSDT":    3,
	"NEOUSDT":    4,
	"LTCUSDT":    5,
	"QTUMUSDT":   6,
	"ADAUSDT":    7,
	"XRPUSDT":    8,
	"EOSUSDT":    9,
	"TUSDUSDT":   10,
	"IOTAUSDT":   11,
	"XLMUSDT":    12,
	"ONTUSDT":    13,
	"TRXUSDT":    14,
	"ETCUSDT":    15,
	"ICXUSDT":    16,
	"VENUSDT":    17,
	"NULSUSDT":   18,
	"VETUSDT":    19,
	"PAXUSDT":    20,
	"BCHABCUSDT": 21,
	"BCHSVUSDT":  22,
	"USDCUSDT":   23,
	"LINKUSDT":   24,
	"WAVESUSDT":  25,
	"BTTUSDT":    26,
	"USDSUSDT":   27,
	"ONGUSDT":    28,
	"HOTUSDT":    29,
	"ZILUSDT":    30,
	"ZRXUSDT":    31,
	"FETUSDT":    32,
	"BATUSDT":    33,
	"XMRUSDT":    34,
	"ZECUSDT":    35,
	"IOSTUSDT":   36,
	"CELRUSDT":   37,
	"DASHUSDT":   38,
	"NANOUSDT":   39,
	"OMGUSDT":    40,
	"THETAUSDT":  41,
	"ENJUSDT":    42,
	"MITHUSDT":   43,
	"MATICUSDT":  44,
	"ATOMUSDT":   45,
	"TFUELUSDT":  46,
	"ONEUSDT":    47,
	"FTMUSDT":    48,
	"ALGOUSDT":   49,
	"USDSBUSDT":  50,
	"GTOUSDT":    51,
	"ERDUSDT":    52,
	"DOGEUSDT":   53,
	"DUSKUSDT":   54,
	"ANKRUSDT":   55,
	"WINUSDT":    56,
	"COSUSDT":    57,
	"NPXSUSDT":   58,
	"COCOSUSDT":  59,
	"MTLUSDT":    60,
	"TOMOUSDT":   61,
	"PERLUSDT":   62,
	"DENTUSDT":   63,
	"MFTUSDT":    64,
	"KEYUSDT":    65,
	"STORMUSDT":  66,
	"DOCKUSDT":   67,
	"WANUSDT":    68,
	"FUNUSDT":    69,
	"CVCUSDT":    70,
	"CHZUSDT":    71,
	"BANDUSDT":   72,
}

func (x Pair) String() string {
	return proto.EnumName(Pair_name, int32(x))
}

func (Pair) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71c081736a82b062, []int{0}
}

func init() {
	proto.RegisterEnum("pair.Pair", Pair_name, Pair_value)
}

func init() { proto.RegisterFile("pair/pair.proto", fileDescriptor_71c081736a82b062) }

var fileDescriptor_71c081736a82b062 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x94, 0x79, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x39, 0x42, 0xd3, 0xba, 0x27, 0xe6, 0xbe, 0x8f, 0x52, 0x8e, 0x42, 0x13, 0xa0, 0xdc,
	0xf7, 0x7a, 0xbd, 0x89, 0x4d, 0xec, 0xdd, 0x60, 0xaf, 0xd3, 0x90, 0xff, 0x1a, 0x28, 0x10, 0xa9,
	0x52, 0xa2, 0x88, 0x7e, 0x64, 0xbe, 0x07, 0x9a, 0xf7, 0xdc, 0x7f, 0xaa, 0xfe, 0x32, 0x33, 0x6f,
	0xde, 0xcc, 0x8e, 0x1c, 0xac, 0xcf, 0xf6, 0x27, 0xf3, 0xb6, 0xfc, 0x69, 0xcd, 0xe6, 0xd3, 0xbf,
	0xd3, 0xb0, 0x21, 0xff, 0x6f, 0xff, 0x6b, 0x06, 0x8d, 0xfe, 0xfe, 0x64, 0x1e, 0x2e, 0x07, 0xcd,
	0xc8, 0xeb, 0xaa, 0x8c, 0xfd, 0xc6, 0x09, 0x01, 0xe3, 0x13, 0xc0, 0x49, 0x44, 0x6c, 0x04, 0x38,
	0x05, 0xd0, 0x4c, 0x3b, 0x2d, 0x60, 0x8d, 0x03, 0x34, 0x04, 0xb2, 0x5a, 0xe0, 0x4c, 0xb8, 0x12,
	0x2c, 0x7e, 0xf3, 0x55, 0x0e, 0x5a, 0x90, 0x90, 0x8a, 0x15, 0xa0, 0x29, 0x30, 0x2c, 0xfa, 0x80,
	0x45, 0x34, 0x72, 0x25, 0x60, 0x49, 0x8a, 0x7c, 0x55, 0xc6, 0xa0, 0x40, 0x28, 0x75, 0x9e, 0x55,
	0xcb, 0xa8, 0xca, 0xa8, 0xb7, 0x22, 0xe0, 0xac, 0x07, 0xac, 0x0a, 0xf8, 0x62, 0x08, 0x58, 0xa3,
	0x71, 0x9a, 0x58, 0x17, 0x48, 0x35, 0x23, 0x1b, 0x02, 0x03, 0x63, 0x01, 0x67, 0x45, 0xdb, 0x56,
	0x19, 0xfb, 0x86, 0x0c, 0x51, 0xee, 0x9c, 0x40, 0x5f, 0xb1, 0xe8, 0x7c, 0xb8, 0x16, 0x04, 0x91,
	0x4e, 0x54, 0x44, 0xc5, 0x0b, 0xe1, 0x6a, 0xb0, 0x14, 0xe9, 0xa4, 0x1c, 0x00, 0x2f, 0x8a, 0x4c,
	0x55, 0xc6, 0x0c, 0x5e, 0x12, 0xca, 0x52, 0xdb, 0x03, 0x5d, 0x96, 0xd4, 0x3d, 0x35, 0x30, 0xec,
	0x71, 0x85, 0xeb, 0x65, 0x8f, 0xab, 0x75, 0x1d, 0x43, 0xd7, 0x38, 0x4d, 0x17, 0x70, 0x5d, 0x20,
	0x71, 0xcc, 0xbb, 0x21, 0x30, 0x4a, 0x33, 0xc0, 0x4d, 0x40, 0x3d, 0xe7, 0x2d, 0x81, 0x4e, 0x6d,
	0xf9, 0x36, 0xb4, 0x15, 0xe1, 0x0e, 0x16, 0x95, 0x17, 0x80, 0xbb, 0xa8, 0x31, 0xf4, 0xb7, 0xc9,
	0x85, 0x96, 0xcc, 0xbb, 0x27, 0xa4, 0x4d, 0xc6, 0xc4, 0x2d, 0xa1, 0x58, 0x95, 0x7c, 0xf1, 0xfb,
	0x58, 0x8f, 0xb2, 0x7c, 0xd8, 0x07, 0xf0, 0x97, 0xd3, 0xdf, 0x43, 0x19, 0xcb, 0x27, 0xa6, 0x7e,
	0x96, 0x47, 0xd8, 0xb7, 0xfd, 0x0a, 0xd8, 0x96, 0xb2, 0x3c, 0xad, 0xcf, 0xe6, 0xb1, 0x64, 0xe6,
	0xca, 0xa7, 0xec, 0xfe, 0x44, 0x82, 0xca, 0x3b, 0xbe, 0xe0, 0x0e, 0x64, 0x3a, 0x95, 0xe1, 0x6c,
	0x2d, 0xae, 0xc0, 0x00, 0xda, 0x98, 0xcd, 0x33, 0xf1, 0x29, 0xca, 0xb2, 0x2e, 0xad, 0x3c, 0x93,
	0x32, 0x59, 0x1c, 0x8f, 0xf1, 0xb9, 0x64, 0x76, 0x3d, 0x63, 0xbb, 0xb0, 0x52, 0xf0, 0x78, 0x5e,
	0x60, 0x1e, 0xd7, 0xa5, 0xe2, 0x4b, 0x50, 0x55, 0xf2, 0x65, 0x5e, 0x41, 0xd2, 0xf6, 0x38, 0xf9,
	0x6b, 0x29, 0xdb, 0x4b, 0x79, 0x17, 0x6f, 0x04, 0x74, 0x7d, 0x8e, 0x6f, 0xb1, 0x85, 0xfe, 0x90,
	0xf4, 0x4e, 0x5a, 0x6b, 0x77, 0x1c, 0x7c, 0x2f, 0x99, 0xb9, 0xa7, 0xfd, 0x0f, 0x38, 0x5c, 0x97,
	0xd3, 0xc8, 0x47, 0xa1, 0xbe, 0x29, 0x18, 0xfb, 0x84, 0xde, 0xa6, 0x3e, 0xd6, 0xcf, 0x28, 0xeb,
	0x10, 0xbe, 0x08, 0xf4, 0xcc, 0x77, 0x80, 0x12, 0xfd, 0xd2, 0xbb, 0x82, 0x73, 0x47, 0x1c, 0x40,
	0xd3, 0xb2, 0x86, 0x49, 0x45, 0x93, 0x31, 0xf6, 0x53, 0x11, 0x0c, 0x1c, 0x0f, 0xb8, 0xe3, 0x0e,
	0x20, 0x19, 0x01, 0xba, 0xa2, 0x10, 0x29, 0xcb, 0x85, 0x24, 0xd1, 0xd6, 0x68, 0xf3, 0xe0, 0xcf,
	0xec, 0x47, 0x6b, 0x32, 0x6d, 0x8f, 0x8f, 0x0e, 0x0f, 0x77, 0xe6, 0x93, 0x9f, 0x07, 0xf3, 0x36,
	0xbe, 0x04, 0xe3, 0xa3, 0x5f, 0xed, 0xdf, 0x53, 0x7c, 0x1a, 0xc6, 0x0b, 0xf8, 0x65, 0xf7, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xbc, 0xb9, 0x9e, 0x2e, 0x04, 0x00, 0x00,
}