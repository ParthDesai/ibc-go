// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/connection/v1/genesis.proto

package types

import (
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

// GenesisState defines the ibc connection submodule's genesis state.
type GenesisState struct {
	Connections           []IdentifiedConnection `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections"`
	ClientConnectionPaths []ConnectionPaths      `protobuf:"bytes,2,rep,name=client_connection_paths,json=clientConnectionPaths,proto3" json:"client_connection_paths" yaml:"client_connection_paths"`
	// the sequence for the next generated connection identifier
	NextConnectionSequence uint64 `protobuf:"varint,3,opt,name=next_connection_sequence,json=nextConnectionSequence,proto3" json:"next_connection_sequence,omitempty" yaml:"next_connection_sequence"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1879d34bc6ac3cd7, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetConnections() []IdentifiedConnection {
	if m != nil {
		return m.Connections
	}
	return nil
}

func (m *GenesisState) GetClientConnectionPaths() []ConnectionPaths {
	if m != nil {
		return m.ClientConnectionPaths
	}
	return nil
}

func (m *GenesisState) GetNextConnectionSequence() uint64 {
	if m != nil {
		return m.NextConnectionSequence
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.core.connection.v1.GenesisState")
}

func init() {
	proto.RegisterFile("ibc/core/connection/v1/genesis.proto", fileDescriptor_1879d34bc6ac3cd7)
}

var fileDescriptor_1879d34bc6ac3cd7 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4f, 0xc2, 0x40,
	0x18, 0x86, 0x5b, 0x20, 0x0e, 0xc5, 0xa9, 0x51, 0x6c, 0x18, 0xae, 0xa4, 0x1a, 0x61, 0x90, 0x3b,
	0x91, 0x49, 0xc7, 0x3a, 0x18, 0x37, 0x03, 0x4e, 0x26, 0x86, 0xd0, 0xe3, 0xb3, 0x5c, 0x42, 0xef,
	0x43, 0xee, 0x20, 0xf2, 0x0b, 0x5c, 0xfd, 0x59, 0x2c, 0x26, 0x8c, 0x4e, 0xc4, 0xc0, 0x3f, 0xe0,
	0x17, 0x98, 0xb6, 0xc6, 0xa2, 0xb1, 0xdb, 0xe5, 0xbe, 0xe7, 0x7d, 0xde, 0xe1, 0xb5, 0x4e, 0x44,
	0xc0, 0x19, 0xc7, 0x09, 0x30, 0x8e, 0x52, 0x02, 0xd7, 0x02, 0x25, 0x9b, 0xb5, 0x58, 0x08, 0x12,
	0x94, 0x50, 0x74, 0x3c, 0x41, 0x8d, 0x76, 0x45, 0x04, 0x9c, 0xc6, 0x14, 0xcd, 0x28, 0x3a, 0x6b,
	0x55, 0x0f, 0x42, 0x0c, 0x31, 0x41, 0x58, 0xfc, 0x4a, 0xe9, 0x6a, 0x3d, 0xc7, 0xb9, 0x93, 0x4d,
	0x40, 0xef, 0xbd, 0x60, 0xed, 0xdf, 0xa4, 0x45, 0x5d, 0xdd, 0xd7, 0x60, 0xdf, 0x5b, 0xe5, 0x0c,
	0x52, 0x8e, 0x59, 0x2b, 0x36, 0xca, 0x17, 0x67, 0xf4, 0xff, 0x76, 0x7a, 0x3b, 0x00, 0xa9, 0xc5,
	0x93, 0x80, 0xc1, 0xf5, 0xcf, 0xbf, 0x5f, 0x5a, 0xac, 0x5c, 0xa3, 0xb3, 0xab, 0xb1, 0x5f, 0x4d,
	0xeb, 0x88, 0x8f, 0x04, 0x48, 0xdd, 0xcb, 0xbe, 0x7b, 0xe3, 0xbe, 0x1e, 0x2a, 0xa7, 0x90, 0x54,
	0xd4, 0xf3, 0x2a, 0x32, 0xf1, 0x5d, 0x8c, 0xfb, 0xa7, 0xb1, 0x7d, 0xbb, 0x72, 0xc9, 0xbc, 0x1f,
	0x8d, 0xae, 0xbc, 0x1c, 0xab, 0xd7, 0x39, 0x4c, 0x2f, 0x7f, 0xe2, 0xf6, 0xa3, 0xe5, 0x48, 0x78,
	0xf9, 0x15, 0x50, 0xf0, 0x3c, 0x05, 0xc9, 0xc1, 0x29, 0xd6, 0xcc, 0x46, 0xc9, 0x3f, 0xde, 0xae,
	0x5c, 0x37, 0x95, 0xe7, 0x91, 0x5e, 0xa7, 0x12, 0x9f, 0x32, 0x77, 0xf7, 0xfb, 0xe0, 0x77, 0x17,
	0x6b, 0x62, 0x2e, 0xd7, 0xc4, 0xfc, 0x5c, 0x13, 0xf3, 0x6d, 0x43, 0x8c, 0xe5, 0x86, 0x18, 0x1f,
	0x1b, 0x62, 0x3c, 0x5c, 0x86, 0x42, 0x0f, 0xa7, 0x01, 0xe5, 0x18, 0x31, 0x8e, 0x2a, 0x42, 0xc5,
	0x44, 0xc0, 0x9b, 0x21, 0xb2, 0x08, 0x07, 0xd3, 0x11, 0xa8, 0x74, 0xaf, 0xf3, 0x76, 0x73, 0x67,
	0x32, 0x3d, 0x1f, 0x83, 0x0a, 0xf6, 0x92, 0xad, 0xda, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85,
	0x66, 0x6f, 0x29, 0x2a, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextConnectionSequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextConnectionSequence))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ClientConnectionPaths) > 0 {
		for iNdEx := len(m.ClientConnectionPaths) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientConnectionPaths[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Connections) > 0 {
		for iNdEx := len(m.Connections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Connections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Connections) > 0 {
		for _, e := range m.Connections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientConnectionPaths) > 0 {
		for _, e := range m.ClientConnectionPaths {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextConnectionSequence != 0 {
		n += 1 + sovGenesis(uint64(m.NextConnectionSequence))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Connections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Connections = append(m.Connections, IdentifiedConnection{})
			if err := m.Connections[len(m.Connections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientConnectionPaths", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientConnectionPaths = append(m.ClientConnectionPaths, ConnectionPaths{})
			if err := m.ClientConnectionPaths[len(m.ClientConnectionPaths)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextConnectionSequence", wireType)
			}
			m.NextConnectionSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextConnectionSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)