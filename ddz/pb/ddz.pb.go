// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ddz.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		ddz.proto

	It has these top-level messages:
		RegisterRequest
		RegisterResponse
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RegisterRequest struct {
	OP   *int32 `protobuf:"varint,1,req,name=OP,def=10001" json:"OP,omitempty"`
	Name string `protobuf:"bytes,2,req,name=name" json:"name"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptorDdz, []int{0} }

const Default_RegisterRequest_OP int32 = 10001

func (m *RegisterRequest) GetOP() int32 {
	if m != nil && m.OP != nil {
		return *m.OP
	}
	return Default_RegisterRequest_OP
}

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RegisterResponse struct {
	OP  *int32 `protobuf:"varint,1,req,name=OP,def=10002" json:"OP,omitempty"`
	Ret int32  `protobuf:"varint,2,req,name=ret" json:"ret"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptorDdz, []int{1} }

const Default_RegisterResponse_OP int32 = 10002

func (m *RegisterResponse) GetOP() int32 {
	if m != nil && m.OP != nil {
		return *m.OP
	}
	return Default_RegisterResponse_OP
}

func (m *RegisterResponse) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "pb.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "pb.RegisterResponse")
}
func (m *RegisterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OP == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("OP")
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDdz(dAtA, i, uint64(*m.OP))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintDdz(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	return i, nil
}

func (m *RegisterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OP == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("OP")
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDdz(dAtA, i, uint64(*m.OP))
	}
	dAtA[i] = 0x10
	i++
	i = encodeVarintDdz(dAtA, i, uint64(m.Ret))
	return i, nil
}

func encodeVarintDdz(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RegisterRequest) Size() (n int) {
	var l int
	_ = l
	if m.OP != nil {
		n += 1 + sovDdz(uint64(*m.OP))
	}
	l = len(m.Name)
	n += 1 + l + sovDdz(uint64(l))
	return n
}

func (m *RegisterResponse) Size() (n int) {
	var l int
	_ = l
	if m.OP != nil {
		n += 1 + sovDdz(uint64(*m.OP))
	}
	n += 1 + sovDdz(uint64(m.Ret))
	return n
}

func sovDdz(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDdz(x uint64) (n int) {
	return sovDdz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterRequest) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDdz
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OP", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OP = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDdz
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipDdz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDdz
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("OP")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterResponse) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDdz
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OP", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OP = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			m.Ret = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ret |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipDdz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDdz
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("OP")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("ret")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDdz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDdz
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
					return 0, ErrIntOverflowDdz
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDdz
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDdz
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDdz
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDdz(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDdz = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDdz   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ddz.proto", fileDescriptorDdz) }

var fileDescriptorDdz = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x49, 0xa9, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe2, 0xe2, 0x0f, 0x4a, 0x4d,
	0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe5, 0x62,
	0xf2, 0x0f, 0x90, 0x60, 0x54, 0x60, 0xd2, 0x60, 0xb5, 0x62, 0x35, 0x34, 0x30, 0x30, 0x30, 0x0c,
	0x62, 0xf2, 0x0f, 0x10, 0x92, 0xe0, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd2,
	0xe0, 0x74, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0x2c, 0xa2, 0xe4, 0xc8, 0x25, 0x80, 0x30,
	0xa3, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xd3, 0x10, 0x23, 0xb0, 0x21, 0x62, 0x5c, 0xcc, 0x45,
	0xa9, 0x25, 0x60, 0x33, 0x58, 0xa1, 0x66, 0x80, 0x04, 0x9c, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x00, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x9d, 0x21, 0x23, 0x0f, 0xa8, 0x00, 0x00, 0x00,
}
