// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capsule8/api/v0/types.proto

/*
Package capsule8_api_v0 is a generated protocol buffer package.

It is generated from these files:
	capsule8/api/v0/types.proto
	capsule8/api/v0/telemetry_event.proto
	capsule8/api/v0/telemetry_service.proto
	capsule8/api/v0/subscription.proto
	capsule8/api/v0/expression.proto

It has these top-level messages:
	IPv4Address
	IPv4AddressAndPort
	IPv6Address
	IPv6AddressAndPort
	NetworkAddress
	Credentials
	TelemetryEvent
	ChargenEvent
	TickerEvent
	ContainerEvent
	ProcessEvent
	SyscallEvent
	FileEvent
	Process
	KernelFunctionCallEvent
	NetworkEvent
	GetEventsRequest
	GetEventsResponse
	ReceivedTelemetryEvent
	Subscription
	ContainerFilter
	EventFilter
	SyscallEventFilter
	ProcessEventFilter
	FileEventFilter
	KernelFunctionCallFilter
	NetworkEventFilter
	ContainerEventFilter
	ChargenEventFilter
	TickerEventFilter
	Modifier
	ThrottleModifier
	LimitModifier
	Value
	BinaryOp
	Expression
*/
package capsule8_api_v0

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

// Supported network address families
type NetworkAddressFamily int32

const (
	// The network address family is unknown
	NetworkAddressFamily_NETWORK_ADDRESS_FAMILY_UNKNOWN NetworkAddressFamily = 0
	// AF_INET; IPv4 address formats
	NetworkAddressFamily_NETWORK_ADDRESS_FAMILY_INET NetworkAddressFamily = 1
	// AF_INET6; IPv6 address formats
	NetworkAddressFamily_NETWORK_ADDRESS_FAMILY_INET6 NetworkAddressFamily = 2
	// AF_LOCAL / AF_UNIX; local filesystem address formats
	NetworkAddressFamily_NETWORK_ADDRESS_FAMILY_LOCAL NetworkAddressFamily = 3
)

var NetworkAddressFamily_name = map[int32]string{
	0: "NETWORK_ADDRESS_FAMILY_UNKNOWN",
	1: "NETWORK_ADDRESS_FAMILY_INET",
	2: "NETWORK_ADDRESS_FAMILY_INET6",
	3: "NETWORK_ADDRESS_FAMILY_LOCAL",
}
var NetworkAddressFamily_value = map[string]int32{
	"NETWORK_ADDRESS_FAMILY_UNKNOWN": 0,
	"NETWORK_ADDRESS_FAMILY_INET":    1,
	"NETWORK_ADDRESS_FAMILY_INET6":   2,
	"NETWORK_ADDRESS_FAMILY_LOCAL":   3,
}

func (x NetworkAddressFamily) String() string {
	return proto.EnumName(NetworkAddressFamily_name, int32(x))
}
func (NetworkAddressFamily) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// An IPv4 address
type IPv4Address struct {
	// The IPv4 address is network byte order (big endian)
	Address uint32 `protobuf:"fixed32,1,opt,name=address" json:"address,omitempty"`
}

func (m *IPv4Address) Reset()                    { *m = IPv4Address{} }
func (m *IPv4Address) String() string            { return proto.CompactTextString(m) }
func (*IPv4Address) ProtoMessage()               {}
func (*IPv4Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IPv4Address) GetAddress() uint32 {
	if m != nil {
		return m.Address
	}
	return 0
}

// An IPv4 address and port
type IPv4AddressAndPort struct {
	// The IPv4 address
	Address *IPv4Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// The port
	Port uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *IPv4AddressAndPort) Reset()                    { *m = IPv4AddressAndPort{} }
func (m *IPv4AddressAndPort) String() string            { return proto.CompactTextString(m) }
func (*IPv4AddressAndPort) ProtoMessage()               {}
func (*IPv4AddressAndPort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IPv4AddressAndPort) GetAddress() *IPv4Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *IPv4AddressAndPort) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// An IPv6 address
type IPv6Address struct {
	// The high-order bytes of the IPv6 address
	High uint64 `protobuf:"fixed64,1,opt,name=high" json:"high,omitempty"`
	// The low-order bytes of the IPv6 address
	Low uint64 `protobuf:"fixed64,2,opt,name=low" json:"low,omitempty"`
}

func (m *IPv6Address) Reset()                    { *m = IPv6Address{} }
func (m *IPv6Address) String() string            { return proto.CompactTextString(m) }
func (*IPv6Address) ProtoMessage()               {}
func (*IPv6Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IPv6Address) GetHigh() uint64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *IPv6Address) GetLow() uint64 {
	if m != nil {
		return m.Low
	}
	return 0
}

// An IPv6 address and port
type IPv6AddressAndPort struct {
	// The IPv6 address
	Address *IPv6Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// The port
	Port uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *IPv6AddressAndPort) Reset()                    { *m = IPv6AddressAndPort{} }
func (m *IPv6AddressAndPort) String() string            { return proto.CompactTextString(m) }
func (*IPv6AddressAndPort) ProtoMessage()               {}
func (*IPv6AddressAndPort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IPv6AddressAndPort) GetAddress() *IPv6Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *IPv6AddressAndPort) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// A network address
type NetworkAddress struct {
	// The address family that specifies which address format is in use
	Family NetworkAddressFamily `protobuf:"varint,1,opt,name=family,enum=capsule8.api.v0.NetworkAddressFamily" json:"family,omitempty"`
	// Types that are valid to be assigned to Address:
	//	*NetworkAddress_Ipv4Address
	//	*NetworkAddress_Ipv6Address
	//	*NetworkAddress_LocalAddress
	Address isNetworkAddress_Address `protobuf_oneof:"address"`
}

func (m *NetworkAddress) Reset()                    { *m = NetworkAddress{} }
func (m *NetworkAddress) String() string            { return proto.CompactTextString(m) }
func (*NetworkAddress) ProtoMessage()               {}
func (*NetworkAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isNetworkAddress_Address interface {
	isNetworkAddress_Address()
}

type NetworkAddress_Ipv4Address struct {
	Ipv4Address *IPv4AddressAndPort `protobuf:"bytes,10,opt,name=ipv4_address,json=ipv4Address,oneof"`
}
type NetworkAddress_Ipv6Address struct {
	Ipv6Address *IPv6AddressAndPort `protobuf:"bytes,20,opt,name=ipv6_address,json=ipv6Address,oneof"`
}
type NetworkAddress_LocalAddress struct {
	LocalAddress string `protobuf:"bytes,30,opt,name=local_address,json=localAddress,oneof"`
}

func (*NetworkAddress_Ipv4Address) isNetworkAddress_Address()  {}
func (*NetworkAddress_Ipv6Address) isNetworkAddress_Address()  {}
func (*NetworkAddress_LocalAddress) isNetworkAddress_Address() {}

func (m *NetworkAddress) GetAddress() isNetworkAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *NetworkAddress) GetFamily() NetworkAddressFamily {
	if m != nil {
		return m.Family
	}
	return NetworkAddressFamily_NETWORK_ADDRESS_FAMILY_UNKNOWN
}

func (m *NetworkAddress) GetIpv4Address() *IPv4AddressAndPort {
	if x, ok := m.GetAddress().(*NetworkAddress_Ipv4Address); ok {
		return x.Ipv4Address
	}
	return nil
}

func (m *NetworkAddress) GetIpv6Address() *IPv6AddressAndPort {
	if x, ok := m.GetAddress().(*NetworkAddress_Ipv6Address); ok {
		return x.Ipv6Address
	}
	return nil
}

func (m *NetworkAddress) GetLocalAddress() string {
	if x, ok := m.GetAddress().(*NetworkAddress_LocalAddress); ok {
		return x.LocalAddress
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NetworkAddress) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NetworkAddress_OneofMarshaler, _NetworkAddress_OneofUnmarshaler, _NetworkAddress_OneofSizer, []interface{}{
		(*NetworkAddress_Ipv4Address)(nil),
		(*NetworkAddress_Ipv6Address)(nil),
		(*NetworkAddress_LocalAddress)(nil),
	}
}

func _NetworkAddress_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NetworkAddress)
	// address
	switch x := m.Address.(type) {
	case *NetworkAddress_Ipv4Address:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ipv4Address); err != nil {
			return err
		}
	case *NetworkAddress_Ipv6Address:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ipv6Address); err != nil {
			return err
		}
	case *NetworkAddress_LocalAddress:
		b.EncodeVarint(30<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.LocalAddress)
	case nil:
	default:
		return fmt.Errorf("NetworkAddress.Address has unexpected type %T", x)
	}
	return nil
}

func _NetworkAddress_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NetworkAddress)
	switch tag {
	case 10: // address.ipv4_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IPv4AddressAndPort)
		err := b.DecodeMessage(msg)
		m.Address = &NetworkAddress_Ipv4Address{msg}
		return true, err
	case 20: // address.ipv6_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IPv6AddressAndPort)
		err := b.DecodeMessage(msg)
		m.Address = &NetworkAddress_Ipv6Address{msg}
		return true, err
	case 30: // address.local_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Address = &NetworkAddress_LocalAddress{x}
		return true, err
	default:
		return false, nil
	}
}

func _NetworkAddress_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NetworkAddress)
	// address
	switch x := m.Address.(type) {
	case *NetworkAddress_Ipv4Address:
		s := proto.Size(x.Ipv4Address)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *NetworkAddress_Ipv6Address:
		s := proto.Size(x.Ipv6Address)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *NetworkAddress_LocalAddress:
		n += proto.SizeVarint(30<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.LocalAddress)))
		n += len(x.LocalAddress)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Process credentials (uid, gid, etc.)
type Credentials struct {
	// The real user ID
	Uid uint32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	// The real group ID
	Gid uint32 `protobuf:"varint,2,opt,name=gid" json:"gid,omitempty"`
	// The effective user ID
	Euid uint32 `protobuf:"varint,3,opt,name=euid" json:"euid,omitempty"`
	// The effective group ID
	Egid uint32 `protobuf:"varint,4,opt,name=egid" json:"egid,omitempty"`
	// The saved user ID
	Suid uint32 `protobuf:"varint,5,opt,name=suid" json:"suid,omitempty"`
	// The saved group ID
	Sgid uint32 `protobuf:"varint,6,opt,name=sgid" json:"sgid,omitempty"`
	// The user ID for filesystem operations
	Fsuid uint32 `protobuf:"varint,7,opt,name=fsuid" json:"fsuid,omitempty"`
	// The group ID for filesystem operations
	Fsgid uint32 `protobuf:"varint,8,opt,name=fsgid" json:"fsgid,omitempty"`
}

func (m *Credentials) Reset()                    { *m = Credentials{} }
func (m *Credentials) String() string            { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()               {}
func (*Credentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Credentials) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Credentials) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *Credentials) GetEuid() uint32 {
	if m != nil {
		return m.Euid
	}
	return 0
}

func (m *Credentials) GetEgid() uint32 {
	if m != nil {
		return m.Egid
	}
	return 0
}

func (m *Credentials) GetSuid() uint32 {
	if m != nil {
		return m.Suid
	}
	return 0
}

func (m *Credentials) GetSgid() uint32 {
	if m != nil {
		return m.Sgid
	}
	return 0
}

func (m *Credentials) GetFsuid() uint32 {
	if m != nil {
		return m.Fsuid
	}
	return 0
}

func (m *Credentials) GetFsgid() uint32 {
	if m != nil {
		return m.Fsgid
	}
	return 0
}

func init() {
	proto.RegisterType((*IPv4Address)(nil), "capsule8.api.v0.IPv4Address")
	proto.RegisterType((*IPv4AddressAndPort)(nil), "capsule8.api.v0.IPv4AddressAndPort")
	proto.RegisterType((*IPv6Address)(nil), "capsule8.api.v0.IPv6Address")
	proto.RegisterType((*IPv6AddressAndPort)(nil), "capsule8.api.v0.IPv6AddressAndPort")
	proto.RegisterType((*NetworkAddress)(nil), "capsule8.api.v0.NetworkAddress")
	proto.RegisterType((*Credentials)(nil), "capsule8.api.v0.Credentials")
	proto.RegisterEnum("capsule8.api.v0.NetworkAddressFamily", NetworkAddressFamily_name, NetworkAddressFamily_value)
}

func init() { proto.RegisterFile("capsule8/api/v0/types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x51, 0x6f, 0x9a, 0x50,
	0x1c, 0xc5, 0x45, 0x5b, 0x5c, 0xff, 0xd4, 0x8e, 0xdc, 0xf8, 0x60, 0xd2, 0xa6, 0x23, 0x2c, 0xcd,
	0xcc, 0x1e, 0xb0, 0x69, 0x1b, 0xb2, 0x97, 0x3d, 0xb0, 0xd6, 0x46, 0x53, 0x87, 0xe6, 0xea, 0x62,
	0xf6, 0xc4, 0x98, 0xa0, 0xde, 0x8c, 0x09, 0xe1, 0x22, 0xc6, 0x0f, 0xb2, 0xe7, 0x7d, 0x82, 0x7d,
	0xc7, 0xe5, 0x5e, 0x2e, 0x4c, 0x37, 0x35, 0x4b, 0xdf, 0xfe, 0xf7, 0xf0, 0xfb, 0x1f, 0x0e, 0xe7,
	0x06, 0x38, 0x9f, 0xb8, 0x11, 0x5d, 0x06, 0xfe, 0xbb, 0x96, 0x1b, 0x91, 0x56, 0x7a, 0xdd, 0x4a,
	0xd6, 0x91, 0x4f, 0x8d, 0x28, 0x0e, 0x93, 0x10, 0xbd, 0xcc, 0x1f, 0x1a, 0x6e, 0x44, 0x8c, 0xf4,
	0x5a, 0x7f, 0x03, 0x4a, 0x77, 0x90, 0xde, 0x59, 0x9e, 0x17, 0xfb, 0x94, 0xa2, 0x06, 0x54, 0xdd,
	0x6c, 0x6c, 0x48, 0x9a, 0xd4, 0xac, 0xe2, 0xfc, 0xa8, 0x7f, 0x01, 0xb4, 0x01, 0x5a, 0x0b, 0x6f,
	0x10, 0xc6, 0x09, 0x32, 0xb7, 0x79, 0xe5, 0xe6, 0xc2, 0xf8, 0xeb, 0x0d, 0xc6, 0xc6, 0x56, 0xe1,
	0x86, 0x10, 0x1c, 0x45, 0x61, 0x9c, 0x34, 0xca, 0x9a, 0xd4, 0xac, 0x61, 0x3e, 0xeb, 0xb7, 0x3c,
	0x8a, 0x69, 0xfd, 0x41, 0xe6, 0x64, 0x36, 0xe7, 0xbe, 0x32, 0xe6, 0x33, 0x52, 0xa1, 0x12, 0x84,
	0x2b, 0xbe, 0x25, 0x63, 0x36, 0x8a, 0x58, 0xe6, 0xb3, 0x62, 0x99, 0xff, 0x15, 0xeb, 0x47, 0x19,
	0xce, 0x6c, 0x3f, 0x59, 0x85, 0xf1, 0xb7, 0x3c, 0xda, 0x7b, 0x90, 0xa7, 0xee, 0x77, 0x12, 0xac,
	0xb9, 0xfb, 0xd9, 0xcd, 0xd5, 0x3f, 0xee, 0xdb, 0x0b, 0x8f, 0x1c, 0xc6, 0x62, 0x09, 0x75, 0xe0,
	0x94, 0x44, 0xe9, 0x9d, 0x93, 0x47, 0x04, 0x1e, 0xf1, 0xf5, 0xa1, 0xe6, 0xc4, 0x87, 0x75, 0x4a,
	0x58, 0x61, 0xab, 0x79, 0x90, 0xcc, 0xc9, 0x2c, 0x9c, 0xea, 0xfb, 0x9d, 0xcc, 0x9d, 0x4e, 0x45,
	0xdb, 0x57, 0x50, 0x0b, 0xc2, 0x89, 0x1b, 0x14, 0x56, 0x97, 0x9a, 0xd4, 0x3c, 0xe9, 0x94, 0xf0,
	0x29, 0x97, 0x05, 0xf6, 0xe1, 0xa4, 0x28, 0x56, 0xff, 0x25, 0x81, 0x72, 0x1f, 0xfb, 0x9e, 0xbf,
	0x48, 0x88, 0x1b, 0x50, 0x76, 0x37, 0x4b, 0xe2, 0xf1, 0x46, 0x6a, 0x98, 0x8d, 0x4c, 0x99, 0x11,
	0x4f, 0x94, 0xc9, 0x46, 0xd6, 0xaf, 0xcf, 0xa0, 0x4a, 0xd6, 0x2f, 0x9b, 0xb9, 0xc6, 0xb0, 0x23,
	0xa1, 0x09, 0x8e, 0x32, 0xee, 0x38, 0xd3, 0xa8, 0xe0, 0x28, 0xe3, 0x64, 0xa1, 0x31, 0xae, 0x0e,
	0xc7, 0x53, 0x0e, 0x56, 0xb9, 0x98, 0x1d, 0x32, 0x95, 0xa1, 0x2f, 0x72, 0x75, 0x46, 0xbc, 0xb7,
	0x3f, 0x25, 0xa8, 0xef, 0xba, 0x16, 0xa4, 0xc3, 0xa5, 0xdd, 0x1e, 0x8d, 0xfb, 0xf8, 0xc9, 0xb1,
	0x1e, 0x1e, 0x70, 0x7b, 0x38, 0x74, 0x1e, 0xad, 0x8f, 0xdd, 0xde, 0x67, 0xe7, 0x93, 0xfd, 0x64,
	0xf7, 0xc7, 0xb6, 0x5a, 0x42, 0xaf, 0xe0, 0x7c, 0x0f, 0xd3, 0xb5, 0xdb, 0x23, 0x55, 0x42, 0x1a,
	0x5c, 0x1c, 0x00, 0x4c, 0xb5, 0x7c, 0x80, 0xe8, 0xf5, 0xef, 0xad, 0x9e, 0x5a, 0xf9, 0x2a, 0xf3,
	0x7f, 0xf4, 0xf6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xf4, 0xc6, 0xf8, 0xc2, 0x03, 0x00,
	0x00,
}
