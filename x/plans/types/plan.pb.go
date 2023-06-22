// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plans/plan.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// the enum below determines the pairing algorithm's behaviour with the selected providers feature
type SELECTED_PROVIDERS_MODE int32

const (
	SELECTED_PROVIDERS_MODE_ALLOWED   SELECTED_PROVIDERS_MODE = 0
	SELECTED_PROVIDERS_MODE_MIXED     SELECTED_PROVIDERS_MODE = 1
	SELECTED_PROVIDERS_MODE_EXCLUSIVE SELECTED_PROVIDERS_MODE = 2
	SELECTED_PROVIDERS_MODE_DISABLED  SELECTED_PROVIDERS_MODE = 3
)

var SELECTED_PROVIDERS_MODE_name = map[int32]string{
	0: "ALLOWED",
	1: "MIXED",
	2: "EXCLUSIVE",
	3: "DISABLED",
}

var SELECTED_PROVIDERS_MODE_value = map[string]int32{
	"ALLOWED":   0,
	"MIXED":     1,
	"EXCLUSIVE": 2,
	"DISABLED":  3,
}

func (x SELECTED_PROVIDERS_MODE) String() string {
	return proto.EnumName(SELECTED_PROVIDERS_MODE_name, int32(x))
}

func (SELECTED_PROVIDERS_MODE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5909a10cd0e3497, []int{0}
}

type Plan struct {
	Index                    string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Block                    uint64     `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
	Price                    types.Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
	AllowOveruse             bool       `protobuf:"varint,8,opt,name=allow_overuse,json=allowOveruse,proto3" json:"allow_overuse,omitempty"`
	OveruseRate              uint64     `protobuf:"varint,9,opt,name=overuse_rate,json=overuseRate,proto3" json:"overuse_rate,omitempty"`
	Description              string     `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Type                     string     `protobuf:"bytes,12,opt,name=type,proto3" json:"type,omitempty"`
	AnnualDiscountPercentage uint64     `protobuf:"varint,13,opt,name=annual_discount_percentage,json=annualDiscountPercentage,proto3" json:"annual_discount_percentage,omitempty"`
	PlanPolicy               Policy     `protobuf:"bytes,14,opt,name=plan_policy,json=planPolicy,proto3" json:"plan_policy"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5909a10cd0e3497, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Plan) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Plan) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

func (m *Plan) GetAllowOveruse() bool {
	if m != nil {
		return m.AllowOveruse
	}
	return false
}

func (m *Plan) GetOveruseRate() uint64 {
	if m != nil {
		return m.OveruseRate
	}
	return 0
}

func (m *Plan) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Plan) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Plan) GetAnnualDiscountPercentage() uint64 {
	if m != nil {
		return m.AnnualDiscountPercentage
	}
	return 0
}

func (m *Plan) GetPlanPolicy() Policy {
	if m != nil {
		return m.PlanPolicy
	}
	return Policy{}
}

// protobuf expected in YAML format: used "moretags" to simplify parsing
type Policy struct {
	ChainPolicies         []ChainPolicy           `protobuf:"bytes,1,rep,name=chain_policies,json=chainPolicies,proto3" json:"chain_policies" mapstructure:"chain_policies"`
	GeolocationProfile    uint64                  `protobuf:"varint,2,opt,name=geolocation_profile,json=geolocationProfile,proto3" json:"geolocation_profile" mapstructure:"geolocation_profile"`
	TotalCuLimit          uint64                  `protobuf:"varint,3,opt,name=total_cu_limit,json=totalCuLimit,proto3" json:"total_cu_limit" mapstructure:"total_cu_limit"`
	EpochCuLimit          uint64                  `protobuf:"varint,4,opt,name=epoch_cu_limit,json=epochCuLimit,proto3" json:"epoch_cu_limit" mapstructure:"epoch_cu_limit"`
	MaxProvidersToPair    uint64                  `protobuf:"varint,5,opt,name=max_providers_to_pair,json=maxProvidersToPair,proto3" json:"max_providers_to_pair" mapstructure:"max_providers_to_pair"`
	SelectedProvidersMode SELECTED_PROVIDERS_MODE `protobuf:"varint,6,opt,name=selected_providers_mode,json=selectedProvidersMode,proto3,enum=lavanet.lava.plans.SELECTED_PROVIDERS_MODE" json:"selected_providers_mode" mapstructure:"selected_providers_mode"`
	SelectedProviders     []string                `protobuf:"bytes,7,rep,name=selected_providers,json=selectedProviders,proto3" json:"selected_providers" mapstructure:"selected_providers"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5909a10cd0e3497, []int{1}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return m.Size()
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetChainPolicies() []ChainPolicy {
	if m != nil {
		return m.ChainPolicies
	}
	return nil
}

func (m *Policy) GetGeolocationProfile() uint64 {
	if m != nil {
		return m.GeolocationProfile
	}
	return 0
}

func (m *Policy) GetTotalCuLimit() uint64 {
	if m != nil {
		return m.TotalCuLimit
	}
	return 0
}

func (m *Policy) GetEpochCuLimit() uint64 {
	if m != nil {
		return m.EpochCuLimit
	}
	return 0
}

func (m *Policy) GetMaxProvidersToPair() uint64 {
	if m != nil {
		return m.MaxProvidersToPair
	}
	return 0
}

func (m *Policy) GetSelectedProvidersMode() SELECTED_PROVIDERS_MODE {
	if m != nil {
		return m.SelectedProvidersMode
	}
	return SELECTED_PROVIDERS_MODE_ALLOWED
}

func (m *Policy) GetSelectedProviders() []string {
	if m != nil {
		return m.SelectedProviders
	}
	return nil
}

type ChainPolicy struct {
	ChainId string   `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" mapstructure:"chain_id"`
	Apis    []string `protobuf:"bytes,2,rep,name=apis,proto3" json:"apis,omitempty" mapstructure:"apis"`
}

func (m *ChainPolicy) Reset()         { *m = ChainPolicy{} }
func (m *ChainPolicy) String() string { return proto.CompactTextString(m) }
func (*ChainPolicy) ProtoMessage()    {}
func (*ChainPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5909a10cd0e3497, []int{2}
}
func (m *ChainPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainPolicy.Merge(m, src)
}
func (m *ChainPolicy) XXX_Size() int {
	return m.Size()
}
func (m *ChainPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_ChainPolicy proto.InternalMessageInfo

func (m *ChainPolicy) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ChainPolicy) GetApis() []string {
	if m != nil {
		return m.Apis
	}
	return nil
}

func init() {
	proto.RegisterEnum("lavanet.lava.plans.SELECTED_PROVIDERS_MODE", SELECTED_PROVIDERS_MODE_name, SELECTED_PROVIDERS_MODE_value)
	proto.RegisterType((*Plan)(nil), "lavanet.lava.plans.Plan")
	proto.RegisterType((*Policy)(nil), "lavanet.lava.plans.Policy")
	proto.RegisterType((*ChainPolicy)(nil), "lavanet.lava.plans.ChainPolicy")
}

func init() { proto.RegisterFile("plans/plan.proto", fileDescriptor_e5909a10cd0e3497) }

var fileDescriptor_e5909a10cd0e3497 = []byte{
	// 831 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x4f, 0x6f, 0xdb, 0x36,
	0x1c, 0xb5, 0x62, 0xd9, 0x96, 0x29, 0x27, 0xd0, 0xd8, 0x16, 0xd1, 0x52, 0xcc, 0x72, 0xb5, 0x75,
	0x30, 0x36, 0x40, 0x42, 0x53, 0x6c, 0x87, 0xfd, 0x39, 0xc4, 0xb6, 0x0e, 0x36, 0x9c, 0xc5, 0x50,
	0xba, 0xae, 0xdb, 0x45, 0xa0, 0x29, 0xce, 0xe1, 0x26, 0x8b, 0x82, 0x44, 0x7b, 0xe9, 0x17, 0xd8,
	0x79, 0x1f, 0x63, 0x18, 0xf6, 0x41, 0x7a, 0xec, 0x71, 0x27, 0x61, 0x48, 0x6e, 0x3e, 0xfa, 0x13,
	0x0c, 0xa4, 0x94, 0x34, 0x9e, 0x1d, 0xf4, 0x62, 0x91, 0xef, 0x3d, 0xbe, 0x47, 0xfe, 0xf4, 0x13,
	0x0d, 0x8c, 0x24, 0x42, 0x71, 0xe6, 0x8a, 0x5f, 0x27, 0x49, 0x19, 0x67, 0x10, 0x46, 0x68, 0x89,
	0x62, 0xc2, 0x1d, 0xf1, 0x74, 0x24, 0x7d, 0xf4, 0x70, 0xc6, 0x66, 0x4c, 0xd2, 0xae, 0x18, 0x15,
	0xca, 0xa3, 0x36, 0x66, 0xd9, 0x9c, 0x65, 0xee, 0x14, 0x65, 0xc4, 0x5d, 0x3e, 0x9b, 0x12, 0x8e,
	0x9e, 0xb9, 0x98, 0xd1, 0xd2, 0xc9, 0xfe, 0xbd, 0x0a, 0xd4, 0x49, 0x84, 0x62, 0xf8, 0x10, 0xd4,
	0x68, 0x1c, 0x92, 0x4b, 0x53, 0xe9, 0x28, 0xdd, 0xa6, 0x5f, 0x4c, 0x04, 0x3a, 0x8d, 0x18, 0xfe,
	0xd5, 0xac, 0x76, 0x94, 0xae, 0xea, 0x17, 0x13, 0xf8, 0x05, 0xa8, 0x25, 0x29, 0xc5, 0xc4, 0x54,
	0x3b, 0x4a, 0x57, 0x3f, 0xfe, 0xd0, 0x29, 0x42, 0x1c, 0x11, 0xe2, 0x94, 0x21, 0x4e, 0x9f, 0xd1,
	0xb8, 0xa7, 0xbe, 0xc9, 0xad, 0x8a, 0x5f, 0xa8, 0xe1, 0xc7, 0x60, 0x1f, 0x45, 0x11, 0xfb, 0x2d,
	0x60, 0x4b, 0x92, 0x2e, 0x32, 0x62, 0x6a, 0x1d, 0xa5, 0xab, 0xf9, 0x2d, 0x09, 0x9e, 0x15, 0x18,
	0x7c, 0x02, 0x5a, 0x25, 0x1d, 0xa4, 0x88, 0x13, 0xb3, 0x29, 0x83, 0xf5, 0x12, 0xf3, 0x11, 0x27,
	0xb0, 0x03, 0xf4, 0x90, 0x64, 0x38, 0xa5, 0x09, 0xa7, 0x2c, 0x36, 0x75, 0xb9, 0xe1, 0xbb, 0x10,
	0x84, 0x40, 0xe5, 0xaf, 0x13, 0x62, 0xb6, 0x24, 0x25, 0xc7, 0xf0, 0x1b, 0x70, 0x84, 0xe2, 0x78,
	0x81, 0xa2, 0x20, 0xa4, 0x19, 0x66, 0x8b, 0x98, 0x07, 0x09, 0x49, 0x31, 0x89, 0x39, 0x9a, 0x11,
	0x73, 0x5f, 0xc6, 0x98, 0x85, 0x62, 0x50, 0x0a, 0x26, 0xb7, 0x3c, 0x3c, 0x01, 0xba, 0x28, 0x73,
	0x90, 0xb0, 0x88, 0xe2, 0xd7, 0xe6, 0x81, 0x3c, 0xf8, 0x91, 0xb3, 0xfd, 0x1e, 0x9c, 0x89, 0x54,
	0x94, 0x27, 0x07, 0x02, 0x2b, 0x90, 0x91, 0xaa, 0x01, 0x43, 0x1f, 0xa9, 0xda, 0x9e, 0x51, 0x1d,
	0xa9, 0x5a, 0xcd, 0xa8, 0x8f, 0x54, 0xad, 0x6e, 0x34, 0x46, 0xaa, 0xd6, 0x30, 0x34, 0xfb, 0xef,
	0x3a, 0xa8, 0x17, 0x42, 0x38, 0x07, 0x07, 0xf8, 0x02, 0xd1, 0x32, 0x8c, 0x92, 0xcc, 0x54, 0x3a,
	0xd5, 0xae, 0x7e, 0x6c, 0xed, 0x8a, 0xeb, 0x0b, 0x65, 0x99, 0xf9, 0x54, 0x64, 0xae, 0x73, 0xeb,
	0xa3, 0x39, 0x4a, 0x32, 0x9e, 0x2e, 0x30, 0x5f, 0xa4, 0xe4, 0x2b, 0x7b, 0xd3, 0xcc, 0xf6, 0xf7,
	0xf1, 0xed, 0x1a, 0x4a, 0x32, 0x18, 0x83, 0x07, 0x33, 0xc2, 0x22, 0x86, 0x91, 0xa8, 0x5d, 0x90,
	0xa4, 0xec, 0x67, 0x1a, 0x11, 0x73, 0x4f, 0x54, 0xa4, 0xf7, 0xed, 0x2a, 0xb7, 0x76, 0xd1, 0xeb,
	0xdc, 0xb2, 0x37, 0x53, 0x76, 0x88, 0x6c, 0x1f, 0xde, 0x41, 0x27, 0x05, 0x08, 0x7f, 0x04, 0x07,
	0x9c, 0x71, 0x14, 0x05, 0x78, 0x11, 0x44, 0x74, 0x4e, 0x79, 0xd1, 0x5c, 0xbd, 0xe7, 0xab, 0xdc,
	0xfa, 0x1f, 0xb3, 0x7d, 0x96, 0x4d, 0xde, 0xf6, 0x5b, 0x12, 0xe8, 0x2f, 0xc6, 0x62, 0x2a, 0xac,
	0x49, 0xc2, 0xf0, 0xc5, 0x3b, 0x6b, 0xf5, 0x9d, 0xf5, 0x26, 0xb3, 0x6d, 0xbd, 0xc9, 0xdb, 0x7e,
	0x4b, 0x02, 0x37, 0xd6, 0x1c, 0x3c, 0x9a, 0xa3, 0x4b, 0x71, 0xb2, 0x25, 0x0d, 0x49, 0x9a, 0x05,
	0x9c, 0x05, 0x09, 0xa2, 0xa9, 0x59, 0x93, 0x09, 0x27, 0xab, 0xdc, 0xda, 0x2d, 0x58, 0xe7, 0xd6,
	0x27, 0x9b, 0x41, 0x3b, 0x65, 0xb6, 0x0f, 0xe7, 0xe8, 0x72, 0x72, 0x03, 0xbf, 0x60, 0x13, 0x44,
	0x53, 0xf8, 0x97, 0x02, 0x0e, 0x33, 0x12, 0x11, 0xcc, 0x49, 0x78, 0x67, 0xcd, 0x9c, 0x85, 0xc4,
	0xac, 0x77, 0x94, 0xee, 0xc1, 0xf1, 0xe7, 0xbb, 0x9a, 0xe2, 0xdc, 0x1b, 0x7b, 0xfd, 0x17, 0xde,
	0x20, 0x98, 0xf8, 0x67, 0x2f, 0x87, 0x03, 0xcf, 0x3f, 0x0f, 0x4e, 0xcf, 0x06, 0x5e, 0xcf, 0x5b,
	0xe5, 0xd6, 0x7d, 0x7e, 0xeb, 0xdc, 0xfa, 0x74, 0x73, 0x9f, 0xf7, 0x08, 0x6d, 0xff, 0xd1, 0x0d,
	0x73, 0xbb, 0xdd, 0x53, 0x16, 0x12, 0xf8, 0x0b, 0x80, 0xdb, 0x4b, 0xcc, 0x46, 0xa7, 0xda, 0x6d,
	0xf6, 0xbe, 0x5e, 0xe5, 0xd6, 0x0e, 0x76, 0x9d, 0x5b, 0x4f, 0xde, 0x17, 0x6a, 0xfb, 0x1f, 0x6c,
	0xe5, 0xd9, 0x4b, 0xa0, 0xdf, 0xe9, 0x7c, 0xf8, 0x25, 0xd0, 0x8a, 0x2e, 0xa7, 0x61, 0x71, 0x81,
	0xf5, 0x1e, 0xaf, 0x73, 0xeb, 0x70, 0xd7, 0x77, 0x40, 0x43, 0xdb, 0x6f, 0xc8, 0xe1, 0x30, 0x84,
	0x2e, 0x50, 0x51, 0x42, 0x33, 0x73, 0x4f, 0x6e, 0xf2, 0x71, 0xf9, 0xfd, 0x3c, 0xd8, 0x5c, 0x27,
	0x14, 0xb6, 0x2f, 0x85, 0x9f, 0x7d, 0x07, 0x0e, 0xef, 0x29, 0x2e, 0xd4, 0x41, 0xe3, 0x64, 0x3c,
	0x3e, 0xfb, 0xc1, 0x1b, 0x18, 0x15, 0xd8, 0x04, 0xb5, 0xd3, 0xe1, 0x2b, 0x6f, 0x60, 0x28, 0x70,
	0x1f, 0x34, 0xbd, 0x57, 0xfd, 0xf1, 0xf7, 0xe7, 0xc3, 0x97, 0x9e, 0xb1, 0x07, 0x5b, 0x40, 0x1b,
	0x0c, 0xcf, 0x4f, 0x7a, 0x63, 0x6f, 0x60, 0x54, 0x7b, 0xfd, 0x3f, 0xaf, 0xda, 0xca, 0x9b, 0xab,
	0xb6, 0xf2, 0xf6, 0xaa, 0xad, 0xfc, 0x7b, 0xd5, 0x56, 0xfe, 0xb8, 0x6e, 0x57, 0xde, 0x5e, 0xb7,
	0x2b, 0xff, 0x5c, 0xb7, 0x2b, 0x3f, 0x3d, 0x9d, 0x51, 0x7e, 0xb1, 0x98, 0x3a, 0x98, 0xcd, 0xdd,
	0xf2, 0x35, 0xcb, 0xa7, 0x7b, 0xe9, 0x16, 0xff, 0x09, 0xe2, 0x66, 0xcb, 0xa6, 0x75, 0x79, 0x97,
	0x3f, 0xff, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xea, 0x8a, 0x65, 0x4e, 0x29, 0x06, 0x00, 0x00,
}

func (this *Plan) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Plan)
	if !ok {
		that2, ok := that.(Plan)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if this.Block != that1.Block {
		return false
	}
	if !this.Price.Equal(&that1.Price) {
		return false
	}
	if this.AllowOveruse != that1.AllowOveruse {
		return false
	}
	if this.OveruseRate != that1.OveruseRate {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.AnnualDiscountPercentage != that1.AnnualDiscountPercentage {
		return false
	}
	if !this.PlanPolicy.Equal(&that1.PlanPolicy) {
		return false
	}
	return true
}
func (this *Policy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Policy)
	if !ok {
		that2, ok := that.(Policy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ChainPolicies) != len(that1.ChainPolicies) {
		return false
	}
	for i := range this.ChainPolicies {
		if !this.ChainPolicies[i].Equal(&that1.ChainPolicies[i]) {
			return false
		}
	}
	if this.GeolocationProfile != that1.GeolocationProfile {
		return false
	}
	if this.TotalCuLimit != that1.TotalCuLimit {
		return false
	}
	if this.EpochCuLimit != that1.EpochCuLimit {
		return false
	}
	if this.MaxProvidersToPair != that1.MaxProvidersToPair {
		return false
	}
	if this.SelectedProvidersMode != that1.SelectedProvidersMode {
		return false
	}
	if len(this.SelectedProviders) != len(that1.SelectedProviders) {
		return false
	}
	for i := range this.SelectedProviders {
		if this.SelectedProviders[i] != that1.SelectedProviders[i] {
			return false
		}
	}
	return true
}
func (this *ChainPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChainPolicy)
	if !ok {
		that2, ok := that.(ChainPolicy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ChainId != that1.ChainId {
		return false
	}
	if len(this.Apis) != len(that1.Apis) {
		return false
	}
	for i := range this.Apis {
		if this.Apis[i] != that1.Apis[i] {
			return false
		}
	}
	return true
}
func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PlanPolicy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	if m.AnnualDiscountPercentage != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.AnnualDiscountPercentage))
		i--
		dAtA[i] = 0x68
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x5a
	}
	if m.OveruseRate != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.OveruseRate))
		i--
		dAtA[i] = 0x48
	}
	if m.AllowOveruse {
		i--
		if m.AllowOveruse {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Block != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SelectedProviders) > 0 {
		for iNdEx := len(m.SelectedProviders) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SelectedProviders[iNdEx])
			copy(dAtA[i:], m.SelectedProviders[iNdEx])
			i = encodeVarintPlan(dAtA, i, uint64(len(m.SelectedProviders[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.SelectedProvidersMode != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.SelectedProvidersMode))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxProvidersToPair != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.MaxProvidersToPair))
		i--
		dAtA[i] = 0x28
	}
	if m.EpochCuLimit != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.EpochCuLimit))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalCuLimit != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.TotalCuLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.GeolocationProfile != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.GeolocationProfile))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainPolicies) > 0 {
		for iNdEx := len(m.ChainPolicies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainPolicies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPlan(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChainPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Apis) > 0 {
		for iNdEx := len(m.Apis) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Apis[iNdEx])
			copy(dAtA[i:], m.Apis[iNdEx])
			i = encodeVarintPlan(dAtA, i, uint64(len(m.Apis[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovPlan(uint64(m.Block))
	}
	l = m.Price.Size()
	n += 1 + l + sovPlan(uint64(l))
	if m.AllowOveruse {
		n += 2
	}
	if m.OveruseRate != 0 {
		n += 1 + sovPlan(uint64(m.OveruseRate))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.AnnualDiscountPercentage != 0 {
		n += 1 + sovPlan(uint64(m.AnnualDiscountPercentage))
	}
	l = m.PlanPolicy.Size()
	n += 1 + l + sovPlan(uint64(l))
	return n
}

func (m *Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChainPolicies) > 0 {
		for _, e := range m.ChainPolicies {
			l = e.Size()
			n += 1 + l + sovPlan(uint64(l))
		}
	}
	if m.GeolocationProfile != 0 {
		n += 1 + sovPlan(uint64(m.GeolocationProfile))
	}
	if m.TotalCuLimit != 0 {
		n += 1 + sovPlan(uint64(m.TotalCuLimit))
	}
	if m.EpochCuLimit != 0 {
		n += 1 + sovPlan(uint64(m.EpochCuLimit))
	}
	if m.MaxProvidersToPair != 0 {
		n += 1 + sovPlan(uint64(m.MaxProvidersToPair))
	}
	if m.SelectedProvidersMode != 0 {
		n += 1 + sovPlan(uint64(m.SelectedProvidersMode))
	}
	if len(m.SelectedProviders) > 0 {
		for _, s := range m.SelectedProviders {
			l = len(s)
			n += 1 + l + sovPlan(uint64(l))
		}
	}
	return n
}

func (m *ChainPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if len(m.Apis) > 0 {
		for _, s := range m.Apis {
			l = len(s)
			n += 1 + l + sovPlan(uint64(l))
		}
	}
	return n
}

func sovPlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlan(x uint64) (n int) {
	return sovPlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowOveruse", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
			m.AllowOveruse = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OveruseRate", wireType)
			}
			m.OveruseRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OveruseRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualDiscountPercentage", wireType)
			}
			m.AnnualDiscountPercentage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnnualDiscountPercentage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PlanPolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
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
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainPolicies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainPolicies = append(m.ChainPolicies, ChainPolicy{})
			if err := m.ChainPolicies[len(m.ChainPolicies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GeolocationProfile", wireType)
			}
			m.GeolocationProfile = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GeolocationProfile |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCuLimit", wireType)
			}
			m.TotalCuLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalCuLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochCuLimit", wireType)
			}
			m.EpochCuLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochCuLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxProvidersToPair", wireType)
			}
			m.MaxProvidersToPair = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxProvidersToPair |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedProvidersMode", wireType)
			}
			m.SelectedProvidersMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SelectedProvidersMode |= SELECTED_PROVIDERS_MODE(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedProviders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SelectedProviders = append(m.SelectedProviders, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
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
func (m *ChainPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: ChainPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Apis", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Apis = append(m.Apis, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
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
func skipPlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
				return 0, ErrInvalidLengthPlan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlan = fmt.Errorf("proto: unexpected end of group")
)
