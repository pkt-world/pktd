// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autopilotrpc/autopilot.proto

package autopilotrpc

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

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b9dc347a92e084, []int{0}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusResponse struct {
	// Indicates whether the autopilot is active or not.
	Active               bool     `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b9dc347a92e084, []int{1}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type ModifyStatusRequest struct {
	// Whether the autopilot agent should be enabled or not.
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyStatusRequest) Reset()         { *m = ModifyStatusRequest{} }
func (m *ModifyStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyStatusRequest) ProtoMessage()    {}
func (*ModifyStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b9dc347a92e084, []int{2}
}

func (m *ModifyStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyStatusRequest.Unmarshal(m, b)
}
func (m *ModifyStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyStatusRequest.Marshal(b, m, deterministic)
}
func (m *ModifyStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyStatusRequest.Merge(m, src)
}
func (m *ModifyStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyStatusRequest.Size(m)
}
func (m *ModifyStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyStatusRequest proto.InternalMessageInfo

func (m *ModifyStatusRequest) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

type ModifyStatusResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyStatusResponse) Reset()         { *m = ModifyStatusResponse{} }
func (m *ModifyStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyStatusResponse) ProtoMessage()    {}
func (*ModifyStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b9dc347a92e084, []int{3}
}

func (m *ModifyStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyStatusResponse.Unmarshal(m, b)
}
func (m *ModifyStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyStatusResponse.Marshal(b, m, deterministic)
}
func (m *ModifyStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyStatusResponse.Merge(m, src)
}
func (m *ModifyStatusResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyStatusResponse.Size(m)
}
func (m *ModifyStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyStatusResponse proto.InternalMessageInfo

type QueryScoresRequest struct {
	Pubkeys []string `protobuf:"bytes,1,rep,name=pubkeys,proto3" json:"pubkeys,omitempty"`
	// If set, we will ignore the local channel state when calculating scores.
	IgnoreLocalState     bool     `protobuf:"varint,2,opt,name=ignore_local_state,json=ignoreLocalState,proto3" json:"ignore_local_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryScoresRequest) Reset()         { *m = QueryScoresRequest{} }
func (m *QueryScoresRequest) String() string { return proto.CompactTextString(m) }
func (*QueryScoresRequest) ProtoMessage()    {}
func (*QueryScoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b9dc347a92e084, []int{4}
}

func (m *QueryScoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryScoresRequest.Unmarshal(m, b)
}
func (m *QueryScoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryScoresRequest.Marshal(b, m, deterministic)
}
func (m *QueryScoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryScoresRequest.Merge(m, src)
}
func (m *QueryScoresRequest) XXX_Size() int {
	return xxx_messageInfo_QueryScoresRequest.Size(m)
}
func (m *QueryScoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryScoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryScoresRequest proto.InternalMessageInfo

func (m *QueryScoresRequest) GetPubkeys() []string {
	if m != nil {
		return m.Pubkeys
	}
	return nil
}

func (m *QueryScoresRequest) GetIgnoreLocalState() bool {
	if m != nil {
		return m.IgnoreLocalState
	}
	return false
}

type QueryScoresResponse struct {
	Results              []*QueryScoresResponse_HeuristicResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *QueryScoresResponse) Reset()         { *m = QueryScoresResponse{} }
func (m *QueryScoresResponse) String() string { return proto.CompactTextString(m) }
func (*QueryScoresResponse) ProtoMessage()    {}
func (*QueryScoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b9dc347a92e084, []int{5}
}

func (m *QueryScoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryScoresResponse.Unmarshal(m, b)
}
func (m *QueryScoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryScoresResponse.Marshal(b, m, deterministic)
}
func (m *QueryScoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryScoresResponse.Merge(m, src)
}
func (m *QueryScoresResponse) XXX_Size() int {
	return xxx_messageInfo_QueryScoresResponse.Size(m)
}
func (m *QueryScoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryScoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryScoresResponse proto.InternalMessageInfo

func (m *QueryScoresResponse) GetResults() []*QueryScoresResponse_HeuristicResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type QueryScoresResponse_HeuristicResult struct {
	Heuristic            string             `protobuf:"bytes,1,opt,name=heuristic,proto3" json:"heuristic,omitempty"`
	Scores               map[string]float64 `protobuf:"bytes,2,rep,name=scores,proto3" json:"scores,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *QueryScoresResponse_HeuristicResult) Reset()         { *m = QueryScoresResponse_HeuristicResult{} }
func (m *QueryScoresResponse_HeuristicResult) String() string { return proto.CompactTextString(m) }
func (*QueryScoresResponse_HeuristicResult) ProtoMessage()    {}
func (*QueryScoresResponse_HeuristicResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b9dc347a92e084, []int{5, 0}
}

func (m *QueryScoresResponse_HeuristicResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryScoresResponse_HeuristicResult.Unmarshal(m, b)
}
func (m *QueryScoresResponse_HeuristicResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryScoresResponse_HeuristicResult.Marshal(b, m, deterministic)
}
func (m *QueryScoresResponse_HeuristicResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryScoresResponse_HeuristicResult.Merge(m, src)
}
func (m *QueryScoresResponse_HeuristicResult) XXX_Size() int {
	return xxx_messageInfo_QueryScoresResponse_HeuristicResult.Size(m)
}
func (m *QueryScoresResponse_HeuristicResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryScoresResponse_HeuristicResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryScoresResponse_HeuristicResult proto.InternalMessageInfo

func (m *QueryScoresResponse_HeuristicResult) GetHeuristic() string {
	if m != nil {
		return m.Heuristic
	}
	return ""
}

func (m *QueryScoresResponse_HeuristicResult) GetScores() map[string]float64 {
	if m != nil {
		return m.Scores
	}
	return nil
}

type SetScoresRequest struct {
	// The name of the heuristic to provide scores to.
	Heuristic string `protobuf:"bytes,1,opt,name=heuristic,proto3" json:"heuristic,omitempty"`
	//
	//A map from hex-encoded public keys to scores. Scores must be in the range
	//[0.0, 1.0].
	Scores               map[string]float64 `protobuf:"bytes,2,rep,name=scores,proto3" json:"scores,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SetScoresRequest) Reset()         { *m = SetScoresRequest{} }
func (m *SetScoresRequest) String() string { return proto.CompactTextString(m) }
func (*SetScoresRequest) ProtoMessage()    {}
func (*SetScoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b9dc347a92e084, []int{6}
}

func (m *SetScoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetScoresRequest.Unmarshal(m, b)
}
func (m *SetScoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetScoresRequest.Marshal(b, m, deterministic)
}
func (m *SetScoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetScoresRequest.Merge(m, src)
}
func (m *SetScoresRequest) XXX_Size() int {
	return xxx_messageInfo_SetScoresRequest.Size(m)
}
func (m *SetScoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetScoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetScoresRequest proto.InternalMessageInfo

func (m *SetScoresRequest) GetHeuristic() string {
	if m != nil {
		return m.Heuristic
	}
	return ""
}

func (m *SetScoresRequest) GetScores() map[string]float64 {
	if m != nil {
		return m.Scores
	}
	return nil
}

type SetScoresResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetScoresResponse) Reset()         { *m = SetScoresResponse{} }
func (m *SetScoresResponse) String() string { return proto.CompactTextString(m) }
func (*SetScoresResponse) ProtoMessage()    {}
func (*SetScoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b9dc347a92e084, []int{7}
}

func (m *SetScoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetScoresResponse.Unmarshal(m, b)
}
func (m *SetScoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetScoresResponse.Marshal(b, m, deterministic)
}
func (m *SetScoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetScoresResponse.Merge(m, src)
}
func (m *SetScoresResponse) XXX_Size() int {
	return xxx_messageInfo_SetScoresResponse.Size(m)
}
func (m *SetScoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetScoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetScoresResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StatusRequest)(nil), "autopilotrpc.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "autopilotrpc.StatusResponse")
	proto.RegisterType((*ModifyStatusRequest)(nil), "autopilotrpc.ModifyStatusRequest")
	proto.RegisterType((*ModifyStatusResponse)(nil), "autopilotrpc.ModifyStatusResponse")
	proto.RegisterType((*QueryScoresRequest)(nil), "autopilotrpc.QueryScoresRequest")
	proto.RegisterType((*QueryScoresResponse)(nil), "autopilotrpc.QueryScoresResponse")
	proto.RegisterType((*QueryScoresResponse_HeuristicResult)(nil), "autopilotrpc.QueryScoresResponse.HeuristicResult")
	proto.RegisterMapType((map[string]float64)(nil), "autopilotrpc.QueryScoresResponse.HeuristicResult.ScoresEntry")
	proto.RegisterType((*SetScoresRequest)(nil), "autopilotrpc.SetScoresRequest")
	proto.RegisterMapType((map[string]float64)(nil), "autopilotrpc.SetScoresRequest.ScoresEntry")
	proto.RegisterType((*SetScoresResponse)(nil), "autopilotrpc.SetScoresResponse")
}

func init() { proto.RegisterFile("autopilotrpc/autopilot.proto", fileDescriptor_e0b9dc347a92e084) }

var fileDescriptor_e0b9dc347a92e084 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x49, 0x8a, 0x5b, 0xf3, 0xb6, 0xda, 0x75, 0xb6, 0x94, 0x10, 0x17, 0xdd, 0xce, 0x69,
	0x11, 0x9b, 0xa5, 0xf5, 0xa2, 0x82, 0x07, 0x2b, 0x82, 0x60, 0x3d, 0x38, 0x4b, 0x2f, 0x22, 0x94,
	0xc9, 0x74, 0xec, 0x86, 0x8d, 0x99, 0x38, 0x3f, 0x0a, 0xf9, 0x87, 0xbc, 0xfa, 0x37, 0x78, 0xf4,
	0xbf, 0x92, 0x64, 0x92, 0x98, 0x84, 0x35, 0x22, 0xf4, 0x10, 0x98, 0xf7, 0xbe, 0x6f, 0x3e, 0x6f,
	0xde, 0x9b, 0x97, 0x81, 0x19, 0x35, 0x5a, 0x64, 0x71, 0x22, 0xb4, 0xcc, 0xd8, 0xb2, 0x31, 0xc2,
	0x4c, 0x0a, 0x2d, 0xd0, 0x5e, 0x5b, 0xc5, 0xfb, 0x70, 0x6f, 0xa5, 0xa9, 0x36, 0x8a, 0xf0, 0x6f,
	0x86, 0x2b, 0x8d, 0x17, 0x70, 0xbf, 0x76, 0xa8, 0x4c, 0xa4, 0x8a, 0xa3, 0x43, 0x18, 0x51, 0xa6,
	0xe3, 0x1b, 0xee, 0x3b, 0x73, 0x67, 0x71, 0x97, 0x54, 0x16, 0x3e, 0x86, 0xe9, 0x07, 0x71, 0x15,
	0x7f, 0xc9, 0x3b, 0x80, 0x22, 0x9c, 0xa7, 0x34, 0x4a, 0x9a, 0x70, 0x6b, 0xe1, 0x43, 0x38, 0xe8,
	0x86, 0x5b, 0x3c, 0xfe, 0x0c, 0xe8, 0xa3, 0xe1, 0x32, 0x5f, 0x31, 0x21, 0x79, 0x43, 0xf1, 0x61,
	0x37, 0x33, 0xd1, 0x86, 0xe7, 0xca, 0x77, 0xe6, 0x3b, 0x0b, 0x8f, 0xd4, 0x26, 0x7a, 0x0a, 0x28,
	0xbe, 0x4e, 0x85, 0xe4, 0x97, 0x89, 0x60, 0x34, 0xb9, 0x54, 0x9a, 0x6a, 0xee, 0xbb, 0x65, 0xae,
	0x89, 0x55, 0xce, 0x0b, 0xa1, 0x48, 0xc3, 0xf1, 0x77, 0x17, 0xa6, 0x1d, 0x7c, 0x55, 0xd4, 0x7b,
	0xd8, 0x95, 0x5c, 0x99, 0x44, 0x5b, 0xfe, 0xf8, 0xf4, 0x24, 0x6c, 0xf7, 0x25, 0xdc, 0xb2, 0x27,
	0x7c, 0xc7, 0x8d, 0x8c, 0x95, 0x8e, 0x19, 0x29, 0x77, 0x92, 0x9a, 0x10, 0xfc, 0x74, 0x60, 0xbf,
	0x27, 0xa2, 0x19, 0x78, 0xeb, 0xda, 0x55, 0x76, 0xc2, 0x23, 0x7f, 0x1c, 0xe8, 0x02, 0x46, 0xaa,
	0x84, 0xfb, 0x6e, 0x99, 0xfd, 0xd5, 0x7f, 0x67, 0x0f, 0xad, 0xfc, 0x36, 0xd5, 0x32, 0x27, 0x15,
	0x2c, 0x78, 0x01, 0xe3, 0x96, 0x1b, 0x4d, 0x60, 0x67, 0xc3, 0xf3, 0x2a, 0x7b, 0xb1, 0x44, 0x07,
	0x70, 0xe7, 0x86, 0x26, 0xc6, 0xf6, 0xcb, 0x21, 0xd6, 0x78, 0xe9, 0x3e, 0x77, 0xf0, 0x0f, 0x07,
	0x26, 0x2b, 0xae, 0xbb, 0xb7, 0x30, 0x5c, 0xc4, 0x59, 0xaf, 0x88, 0x27, 0xdd, 0x22, 0xfa, 0xb4,
	0xdb, 0x3e, 0xf1, 0x14, 0x1e, 0xb4, 0x52, 0xd8, 0x2e, 0x9d, 0xfe, 0x72, 0xc1, 0x7b, 0x5d, 0x9f,
	0x02, 0xbd, 0x81, 0x91, 0x9d, 0x36, 0xf4, 0xb0, 0x77, 0xb6, 0xf6, 0xc8, 0x06, 0xb3, 0xed, 0x62,
	0x35, 0x2a, 0x17, 0xb0, 0xd7, 0x1e, 0x5c, 0x74, 0xd4, 0x8d, 0xde, 0xf2, 0x0f, 0x04, 0x78, 0x28,
	0xa4, 0xc2, 0x12, 0x18, 0xb7, 0xae, 0x19, 0xcd, 0x07, 0x26, 0xc0, 0x42, 0x8f, 0xfe, 0x39, 0x23,
	0xe8, 0x1c, 0xbc, 0xa6, 0x25, 0xe8, 0xd1, 0xf0, 0x75, 0x04, 0x8f, 0xff, 0xaa, 0x5b, 0xda, 0xd9,
	0xc9, 0xa7, 0xe5, 0x75, 0xac, 0xd7, 0x26, 0x0a, 0x99, 0xf8, 0xba, 0xcc, 0x36, 0xfa, 0x98, 0x51,
	0xb5, 0x2e, 0x16, 0x57, 0xcb, 0x24, 0x2d, 0xbe, 0xce, 0xfb, 0x22, 0x33, 0x16, 0x8d, 0xca, 0x37,
	0xe6, 0xd9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x70, 0x2a, 0x77, 0x83, 0x04, 0x00, 0x00,
}
