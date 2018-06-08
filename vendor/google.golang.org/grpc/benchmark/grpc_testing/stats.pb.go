// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stats.proto

package grpc_testing

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

type ServerStats struct {
	// wall clock time change in seconds since last reset
	TimeElapsed float64 `protobuf:"fixed64,1,opt,name=time_elapsed,json=timeElapsed" json:"time_elapsed,omitempty"`
	// change in user time (in seconds) used by the server since last reset
	TimeUser float64 `protobuf:"fixed64,2,opt,name=time_user,json=timeUser" json:"time_user,omitempty"`
	// change in server time (in seconds) used by the server process and all
	// threads since last reset
	TimeSystem           float64  `protobuf:"fixed64,3,opt,name=time_system,json=timeSystem" json:"time_system,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStats) Reset()         { *m = ServerStats{} }
func (m *ServerStats) String() string { return proto.CompactTextString(m) }
func (*ServerStats) ProtoMessage()    {}
func (*ServerStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_stats_8ba831c0cb3c3440, []int{0}
}
func (m *ServerStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStats.Unmarshal(m, b)
}
func (m *ServerStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStats.Marshal(b, m, deterministic)
}
func (dst *ServerStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStats.Merge(dst, src)
}
func (m *ServerStats) XXX_Size() int {
	return xxx_messageInfo_ServerStats.Size(m)
}
func (m *ServerStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStats.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStats proto.InternalMessageInfo

func (m *ServerStats) GetTimeElapsed() float64 {
	if m != nil {
		return m.TimeElapsed
	}
	return 0
}

func (m *ServerStats) GetTimeUser() float64 {
	if m != nil {
		return m.TimeUser
	}
	return 0
}

func (m *ServerStats) GetTimeSystem() float64 {
	if m != nil {
		return m.TimeSystem
	}
	return 0
}

// Histogram params based on grpc/support/histogram.c
type HistogramParams struct {
	Resolution           float64  `protobuf:"fixed64,1,opt,name=resolution" json:"resolution,omitempty"`
	MaxPossible          float64  `protobuf:"fixed64,2,opt,name=max_possible,json=maxPossible" json:"max_possible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistogramParams) Reset()         { *m = HistogramParams{} }
func (m *HistogramParams) String() string { return proto.CompactTextString(m) }
func (*HistogramParams) ProtoMessage()    {}
func (*HistogramParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_stats_8ba831c0cb3c3440, []int{1}
}
func (m *HistogramParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistogramParams.Unmarshal(m, b)
}
func (m *HistogramParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistogramParams.Marshal(b, m, deterministic)
}
func (dst *HistogramParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistogramParams.Merge(dst, src)
}
func (m *HistogramParams) XXX_Size() int {
	return xxx_messageInfo_HistogramParams.Size(m)
}
func (m *HistogramParams) XXX_DiscardUnknown() {
	xxx_messageInfo_HistogramParams.DiscardUnknown(m)
}

var xxx_messageInfo_HistogramParams proto.InternalMessageInfo

func (m *HistogramParams) GetResolution() float64 {
	if m != nil {
		return m.Resolution
	}
	return 0
}

func (m *HistogramParams) GetMaxPossible() float64 {
	if m != nil {
		return m.MaxPossible
	}
	return 0
}

// Histogram data based on grpc/support/histogram.c
type HistogramData struct {
	Bucket               []uint32 `protobuf:"varint,1,rep,packed,name=bucket" json:"bucket,omitempty"`
	MinSeen              float64  `protobuf:"fixed64,2,opt,name=min_seen,json=minSeen" json:"min_seen,omitempty"`
	MaxSeen              float64  `protobuf:"fixed64,3,opt,name=max_seen,json=maxSeen" json:"max_seen,omitempty"`
	Sum                  float64  `protobuf:"fixed64,4,opt,name=sum" json:"sum,omitempty"`
	SumOfSquares         float64  `protobuf:"fixed64,5,opt,name=sum_of_squares,json=sumOfSquares" json:"sum_of_squares,omitempty"`
	Count                float64  `protobuf:"fixed64,6,opt,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistogramData) Reset()         { *m = HistogramData{} }
func (m *HistogramData) String() string { return proto.CompactTextString(m) }
func (*HistogramData) ProtoMessage()    {}
func (*HistogramData) Descriptor() ([]byte, []int) {
	return fileDescriptor_stats_8ba831c0cb3c3440, []int{2}
}
func (m *HistogramData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistogramData.Unmarshal(m, b)
}
func (m *HistogramData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistogramData.Marshal(b, m, deterministic)
}
func (dst *HistogramData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistogramData.Merge(dst, src)
}
func (m *HistogramData) XXX_Size() int {
	return xxx_messageInfo_HistogramData.Size(m)
}
func (m *HistogramData) XXX_DiscardUnknown() {
	xxx_messageInfo_HistogramData.DiscardUnknown(m)
}

var xxx_messageInfo_HistogramData proto.InternalMessageInfo

func (m *HistogramData) GetBucket() []uint32 {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *HistogramData) GetMinSeen() float64 {
	if m != nil {
		return m.MinSeen
	}
	return 0
}

func (m *HistogramData) GetMaxSeen() float64 {
	if m != nil {
		return m.MaxSeen
	}
	return 0
}

func (m *HistogramData) GetSum() float64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *HistogramData) GetSumOfSquares() float64 {
	if m != nil {
		return m.SumOfSquares
	}
	return 0
}

func (m *HistogramData) GetCount() float64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ClientStats struct {
	// Latency histogram. Data points are in nanoseconds.
	Latencies *HistogramData `protobuf:"bytes,1,opt,name=latencies" json:"latencies,omitempty"`
	// See ServerStats for details.
	TimeElapsed          float64  `protobuf:"fixed64,2,opt,name=time_elapsed,json=timeElapsed" json:"time_elapsed,omitempty"`
	TimeUser             float64  `protobuf:"fixed64,3,opt,name=time_user,json=timeUser" json:"time_user,omitempty"`
	TimeSystem           float64  `protobuf:"fixed64,4,opt,name=time_system,json=timeSystem" json:"time_system,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStats) Reset()         { *m = ClientStats{} }
func (m *ClientStats) String() string { return proto.CompactTextString(m) }
func (*ClientStats) ProtoMessage()    {}
func (*ClientStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_stats_8ba831c0cb3c3440, []int{3}
}
func (m *ClientStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStats.Unmarshal(m, b)
}
func (m *ClientStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStats.Marshal(b, m, deterministic)
}
func (dst *ClientStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStats.Merge(dst, src)
}
func (m *ClientStats) XXX_Size() int {
	return xxx_messageInfo_ClientStats.Size(m)
}
func (m *ClientStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStats.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStats proto.InternalMessageInfo

func (m *ClientStats) GetLatencies() *HistogramData {
	if m != nil {
		return m.Latencies
	}
	return nil
}

func (m *ClientStats) GetTimeElapsed() float64 {
	if m != nil {
		return m.TimeElapsed
	}
	return 0
}

func (m *ClientStats) GetTimeUser() float64 {
	if m != nil {
		return m.TimeUser
	}
	return 0
}

func (m *ClientStats) GetTimeSystem() float64 {
	if m != nil {
		return m.TimeSystem
	}
	return 0
}

func init() {
	proto.RegisterType((*ServerStats)(nil), "grpc.testing.ServerStats")
	proto.RegisterType((*HistogramParams)(nil), "grpc.testing.HistogramParams")
	proto.RegisterType((*HistogramData)(nil), "grpc.testing.HistogramData")
	proto.RegisterType((*ClientStats)(nil), "grpc.testing.ClientStats")
}

func init() { proto.RegisterFile("stats.proto", fileDescriptor_stats_8ba831c0cb3c3440) }

var fileDescriptor_stats_8ba831c0cb3c3440 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0xd3, 0xf6, 0xb6, 0x27, 0xed, 0xbd, 0x97, 0x41, 0x24, 0x52, 0xd0, 0x1a, 0x5c,
	0x74, 0x95, 0x85, 0xae, 0x5c, 0xab, 0xe0, 0xce, 0xd2, 0xe8, 0x3a, 0x4c, 0xe3, 0x69, 0x19, 0xcc,
	0xcc, 0xc4, 0x39, 0x33, 0x12, 0x1f, 0x49, 0x7c, 0x49, 0xc9, 0x24, 0x68, 0x55, 0xd0, 0x5d, 0xe6,
	0xfb, 0x7e, 0xe6, 0xe4, 0xe4, 0x0f, 0x44, 0x64, 0xb9, 0xa5, 0xb4, 0x32, 0xda, 0x6a, 0x36, 0xd9,
	0x9a, 0xaa, 0x48, 0x2d, 0x92, 0x15, 0x6a, 0x9b, 0x28, 0x88, 0x32, 0x34, 0x4f, 0x68, 0xb2, 0x26,
	0xc2, 0x8e, 0x61, 0x62, 0x85, 0xc4, 0x1c, 0x4b, 0x5e, 0x11, 0xde, 0xc7, 0xc1, 0x3c, 0x58, 0x04,
	0xab, 0xa8, 0x61, 0x57, 0x2d, 0x62, 0x33, 0x18, 0xfb, 0x88, 0x23, 0x34, 0x71, 0xcf, 0xfb, 0x51,
	0x03, 0xee, 0x08, 0x0d, 0x3b, 0x02, 0x9f, 0xcd, 0xe9, 0x99, 0x2c, 0xca, 0x38, 0xf4, 0x1a, 0x1a,
	0x94, 0x79, 0x92, 0xdc, 0xc2, 0xbf, 0x6b, 0x41, 0x56, 0x6f, 0x0d, 0x97, 0x4b, 0x6e, 0xb8, 0x24,
	0x76, 0x08, 0x60, 0x90, 0x74, 0xe9, 0xac, 0xd0, 0xaa, 0x9b, 0xb8, 0x43, 0x9a, 0x77, 0x92, 0xbc,
	0xce, 0x2b, 0x4d, 0x24, 0xd6, 0x25, 0x76, 0x33, 0x23, 0xc9, 0xeb, 0x65, 0x87, 0x92, 0xd7, 0x00,
	0xa6, 0xef, 0xd7, 0x5e, 0x72, 0xcb, 0xd9, 0x3e, 0x0c, 0xd7, 0xae, 0x78, 0x40, 0x1b, 0x07, 0xf3,
	0x70, 0x31, 0x5d, 0x75, 0x27, 0x76, 0x00, 0x23, 0x29, 0x54, 0x4e, 0x88, 0xaa, 0xbb, 0xe8, 0x8f,
	0x14, 0x2a, 0x43, 0x54, 0x5e, 0xf1, 0xba, 0x55, 0x61, 0xa7, 0x78, 0xed, 0xd5, 0x7f, 0x08, 0xc9,
	0xc9, 0xb8, 0xef, 0x69, 0xf3, 0xc8, 0x4e, 0xe0, 0x2f, 0x39, 0x99, 0xeb, 0x4d, 0x4e, 0x8f, 0x8e,
	0x1b, 0xa4, 0x78, 0xe0, 0xe5, 0x84, 0x9c, 0xbc, 0xd9, 0x64, 0x2d, 0x63, 0x7b, 0x30, 0x28, 0xb4,
	0x53, 0x36, 0x1e, 0x7a, 0xd9, 0x1e, 0x92, 0x97, 0x00, 0xa2, 0x8b, 0x52, 0xa0, 0xb2, 0xed, 0x47,
	0x3f, 0x87, 0x71, 0xc9, 0x2d, 0xaa, 0x42, 0x20, 0xf9, 0xfd, 0xa3, 0xd3, 0x59, 0xba, 0xdb, 0x52,
	0xfa, 0x69, 0xb7, 0xd5, 0x47, 0xfa, 0x5b, 0x5f, 0xbd, 0x5f, 0xfa, 0x0a, 0x7f, 0xee, 0xab, 0xff,
	0xb5, 0xaf, 0xf5, 0xd0, 0xff, 0x34, 0x67, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xea, 0x75, 0x34,
	0x90, 0x43, 0x02, 0x00, 0x00,
}
