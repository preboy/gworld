// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: 4_quest.proto

package msg

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type QuestData struct {
	Key int32 `protobuf:"varint,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Val int32 `protobuf:"varint,2,opt,name=Val,proto3" json:"Val,omitempty"`
}

func (m *QuestData) Reset()                    { *m = QuestData{} }
func (m *QuestData) String() string            { return proto.CompactTextString(m) }
func (*QuestData) ProtoMessage()               {}
func (*QuestData) Descriptor() ([]byte, []int) { return fileDescriptor4Quest, []int{0} }

func (m *QuestData) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *QuestData) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

type QuestInfo struct {
	Id   uint32       `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Task uint32       `protobuf:"varint,2,opt,name=Task,proto3" json:"Task,omitempty"`
	Data []*QuestData `protobuf:"bytes,3,rep,name=Data" json:"Data,omitempty"`
}

func (m *QuestInfo) Reset()                    { *m = QuestInfo{} }
func (m *QuestInfo) String() string            { return proto.CompactTextString(m) }
func (*QuestInfo) ProtoMessage()               {}
func (*QuestInfo) Descriptor() ([]byte, []int) { return fileDescriptor4Quest, []int{1} }

func (m *QuestInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QuestInfo) GetTask() uint32 {
	if m != nil {
		return m.Task
	}
	return 0
}

func (m *QuestInfo) GetData() []*QuestData {
	if m != nil {
		return m.Data
	}
	return nil
}

type QuestListRequest struct {
}

func (m *QuestListRequest) Reset()                    { *m = QuestListRequest{} }
func (m *QuestListRequest) String() string            { return proto.CompactTextString(m) }
func (*QuestListRequest) ProtoMessage()               {}
func (*QuestListRequest) Descriptor() ([]byte, []int) { return fileDescriptor4Quest, []int{2} }

type QuestListResponse struct {
	Quests []*QuestInfo `protobuf:"bytes,1,rep,name=Quests" json:"Quests,omitempty"`
}

func (m *QuestListResponse) Reset()                    { *m = QuestListResponse{} }
func (m *QuestListResponse) String() string            { return proto.CompactTextString(m) }
func (*QuestListResponse) ProtoMessage()               {}
func (*QuestListResponse) Descriptor() ([]byte, []int) { return fileDescriptor4Quest, []int{3} }

func (m *QuestListResponse) GetQuests() []*QuestInfo {
	if m != nil {
		return m.Quests
	}
	return nil
}

// 任务操作
type QuestOpRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Op uint32 `protobuf:"varint,2,opt,name=Op,proto3" json:"Op,omitempty"`
	R  uint32 `protobuf:"varint,3,opt,name=R,proto3" json:"R,omitempty"`
}

func (m *QuestOpRequest) Reset()                    { *m = QuestOpRequest{} }
func (m *QuestOpRequest) String() string            { return proto.CompactTextString(m) }
func (*QuestOpRequest) ProtoMessage()               {}
func (*QuestOpRequest) Descriptor() ([]byte, []int) { return fileDescriptor4Quest, []int{4} }

func (m *QuestOpRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QuestOpRequest) GetOp() uint32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *QuestOpRequest) GetR() uint32 {
	if m != nil {
		return m.R
	}
	return 0
}

type QuestOpResponse struct {
	Id        uint32     `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Op        uint32     `protobuf:"varint,2,opt,name=Op,proto3" json:"Op,omitempty"`
	R         uint32     `protobuf:"varint,3,opt,name=R,proto3" json:"R,omitempty"`
	ErrorCode uint32     `protobuf:"varint,4,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
	Quest     *QuestInfo `protobuf:"bytes,5,opt,name=Quest" json:"Quest,omitempty"`
}

func (m *QuestOpResponse) Reset()                    { *m = QuestOpResponse{} }
func (m *QuestOpResponse) String() string            { return proto.CompactTextString(m) }
func (*QuestOpResponse) ProtoMessage()               {}
func (*QuestOpResponse) Descriptor() ([]byte, []int) { return fileDescriptor4Quest, []int{5} }

func (m *QuestOpResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QuestOpResponse) GetOp() uint32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *QuestOpResponse) GetR() uint32 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *QuestOpResponse) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *QuestOpResponse) GetQuest() *QuestInfo {
	if m != nil {
		return m.Quest
	}
	return nil
}

// 任务变更推送(由服务端产生变化时推送)
type QuestUpdate struct {
	Quests []*QuestInfo `protobuf:"bytes,1,rep,name=Quests" json:"Quests,omitempty"`
}

func (m *QuestUpdate) Reset()                    { *m = QuestUpdate{} }
func (m *QuestUpdate) String() string            { return proto.CompactTextString(m) }
func (*QuestUpdate) ProtoMessage()               {}
func (*QuestUpdate) Descriptor() ([]byte, []int) { return fileDescriptor4Quest, []int{6} }

func (m *QuestUpdate) GetQuests() []*QuestInfo {
	if m != nil {
		return m.Quests
	}
	return nil
}

func init() {
	proto.RegisterType((*QuestData)(nil), "msg.QuestData")
	proto.RegisterType((*QuestInfo)(nil), "msg.QuestInfo")
	proto.RegisterType((*QuestListRequest)(nil), "msg.QuestListRequest")
	proto.RegisterType((*QuestListResponse)(nil), "msg.QuestListResponse")
	proto.RegisterType((*QuestOpRequest)(nil), "msg.QuestOpRequest")
	proto.RegisterType((*QuestOpResponse)(nil), "msg.QuestOpResponse")
	proto.RegisterType((*QuestUpdate)(nil), "msg.QuestUpdate")
}
func (m *QuestData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.Key))
	}
	if m.Val != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.Val))
	}
	return i, nil
}

func (m *QuestInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.Id))
	}
	if m.Task != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.Task))
	}
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			dAtA[i] = 0x1a
			i++
			i = encodeVarint4Quest(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *QuestListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestListRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *QuestListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestListResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Quests) > 0 {
		for _, msg := range m.Quests {
			dAtA[i] = 0xa
			i++
			i = encodeVarint4Quest(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *QuestOpRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestOpRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.Id))
	}
	if m.Op != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.Op))
	}
	if m.R != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.R))
	}
	return i, nil
}

func (m *QuestOpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestOpResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.Id))
	}
	if m.Op != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.Op))
	}
	if m.R != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.R))
	}
	if m.ErrorCode != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.ErrorCode))
	}
	if m.Quest != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarint4Quest(dAtA, i, uint64(m.Quest.Size()))
		n1, err := m.Quest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *QuestUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestUpdate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Quests) > 0 {
		for _, msg := range m.Quests {
			dAtA[i] = 0xa
			i++
			i = encodeVarint4Quest(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarint4Quest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *QuestData) Size() (n int) {
	var l int
	_ = l
	if m.Key != 0 {
		n += 1 + sov4Quest(uint64(m.Key))
	}
	if m.Val != 0 {
		n += 1 + sov4Quest(uint64(m.Val))
	}
	return n
}

func (m *QuestInfo) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sov4Quest(uint64(m.Id))
	}
	if m.Task != 0 {
		n += 1 + sov4Quest(uint64(m.Task))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sov4Quest(uint64(l))
		}
	}
	return n
}

func (m *QuestListRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *QuestListResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Quests) > 0 {
		for _, e := range m.Quests {
			l = e.Size()
			n += 1 + l + sov4Quest(uint64(l))
		}
	}
	return n
}

func (m *QuestOpRequest) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sov4Quest(uint64(m.Id))
	}
	if m.Op != 0 {
		n += 1 + sov4Quest(uint64(m.Op))
	}
	if m.R != 0 {
		n += 1 + sov4Quest(uint64(m.R))
	}
	return n
}

func (m *QuestOpResponse) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sov4Quest(uint64(m.Id))
	}
	if m.Op != 0 {
		n += 1 + sov4Quest(uint64(m.Op))
	}
	if m.R != 0 {
		n += 1 + sov4Quest(uint64(m.R))
	}
	if m.ErrorCode != 0 {
		n += 1 + sov4Quest(uint64(m.ErrorCode))
	}
	if m.Quest != nil {
		l = m.Quest.Size()
		n += 1 + l + sov4Quest(uint64(l))
	}
	return n
}

func (m *QuestUpdate) Size() (n int) {
	var l int
	_ = l
	if len(m.Quests) > 0 {
		for _, e := range m.Quests {
			l = e.Size()
			n += 1 + l + sov4Quest(uint64(l))
		}
	}
	return n
}

func sov4Quest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func soz4Quest(x uint64) (n int) {
	return sov4Quest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuestData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow4Quest
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
			return fmt.Errorf("proto: QuestData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			m.Key = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Key |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			m.Val = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Val |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip4Quest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLength4Quest
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
func (m *QuestInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow4Quest
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
			return fmt.Errorf("proto: QuestInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Task", wireType)
			}
			m.Task = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Task |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength4Quest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &QuestData{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip4Quest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLength4Quest
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
func (m *QuestListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow4Quest
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
			return fmt.Errorf("proto: QuestListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skip4Quest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLength4Quest
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
func (m *QuestListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow4Quest
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
			return fmt.Errorf("proto: QuestListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength4Quest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quests = append(m.Quests, &QuestInfo{})
			if err := m.Quests[len(m.Quests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip4Quest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLength4Quest
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
func (m *QuestOpRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow4Quest
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
			return fmt.Errorf("proto: QuestOpRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestOpRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			m.R = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.R |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip4Quest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLength4Quest
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
func (m *QuestOpResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow4Quest
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
			return fmt.Errorf("proto: QuestOpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestOpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			m.R = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.R |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength4Quest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Quest == nil {
				m.Quest = &QuestInfo{}
			}
			if err := m.Quest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip4Quest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLength4Quest
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
func (m *QuestUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow4Quest
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
			return fmt.Errorf("proto: QuestUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow4Quest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength4Quest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quests = append(m.Quests, &QuestInfo{})
			if err := m.Quests[len(m.Quests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip4Quest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLength4Quest
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
func skip4Quest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow4Quest
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
					return 0, ErrIntOverflow4Quest
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
					return 0, ErrIntOverflow4Quest
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
				return 0, ErrInvalidLength4Quest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflow4Quest
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
				next, err := skip4Quest(dAtA[start:])
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
	ErrInvalidLength4Quest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow4Quest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("4_quest.proto", fileDescriptor4Quest) }

var fileDescriptor4Quest = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbd, 0x4a, 0xc4, 0x40,
	0x14, 0x85, 0x9d, 0xfc, 0x2c, 0xec, 0x5d, 0x13, 0xe3, 0xad, 0xa6, 0x90, 0x10, 0x06, 0x91, 0xad,
	0x22, 0xf8, 0x53, 0x09, 0x16, 0xfe, 0x14, 0x41, 0x21, 0x38, 0xea, 0xb6, 0x32, 0x92, 0xb8, 0x88,
	0xbb, 0x3b, 0x63, 0x26, 0x16, 0xd6, 0xbe, 0x80, 0x8f, 0x65, 0xe9, 0x23, 0x48, 0x7c, 0x11, 0xc9,
	0xdd, 0xb0, 0x8a, 0x5a, 0x68, 0x77, 0xf2, 0x9d, 0x9b, 0x73, 0xcf, 0x65, 0x20, 0xd8, 0xb9, 0xba,
	0x7f, 0x28, 0x6d, 0x9d, 0x9a, 0x4a, 0xd7, 0x1a, 0xdd, 0xa9, 0x1d, 0x8b, 0x4d, 0xe8, 0x9f, 0xb5,
	0xec, 0x48, 0xd5, 0x0a, 0x23, 0x70, 0x4f, 0xca, 0x47, 0xce, 0x12, 0x36, 0xf4, 0x65, 0x2b, 0x5b,
	0x32, 0x52, 0x13, 0xee, 0xcc, 0xc9, 0x48, 0x4d, 0xc4, 0x79, 0xf7, 0x43, 0x36, 0xbb, 0xd1, 0x18,
	0x82, 0x93, 0x15, 0x34, 0x1f, 0x48, 0x27, 0x2b, 0x10, 0xc1, 0xbb, 0x50, 0xf6, 0x8e, 0xe6, 0x03,
	0x49, 0x1a, 0x05, 0x78, 0x6d, 0x38, 0x77, 0x13, 0x77, 0x38, 0xd8, 0x0a, 0xd3, 0xa9, 0x1d, 0xa7,
	0x8b, 0x95, 0x92, 0x3c, 0x81, 0x10, 0x11, 0x3a, 0xbd, 0xb5, 0xb5, 0x2c, 0xa9, 0xa4, 0xd8, 0x83,
	0xd5, 0x2f, 0xcc, 0x1a, 0x3d, 0xb3, 0x25, 0x6e, 0x40, 0x8f, 0xa0, 0xe5, 0xec, 0x7b, 0x5c, 0x5b,
	0x48, 0x76, 0xae, 0xd8, 0x87, 0x90, 0x54, 0x6e, 0xba, 0xb8, 0x1f, 0x55, 0x43, 0x70, 0x72, 0xd3,
	0x15, 0x75, 0x72, 0x83, 0xcb, 0xc0, 0x24, 0x77, 0xe9, 0x93, 0x49, 0xf1, 0xc4, 0x60, 0x65, 0x11,
	0xd0, 0xed, 0xfe, 0x57, 0x02, 0xae, 0x41, 0xff, 0xb8, 0xaa, 0x74, 0x75, 0xa8, 0x8b, 0x92, 0x7b,
	0x44, 0x3f, 0x01, 0xae, 0x83, 0x4f, 0xf1, 0xdc, 0x4f, 0xd8, 0x2f, 0x67, 0xcc, 0x4d, 0xb1, 0x0b,
	0x03, 0x12, 0x97, 0xa6, 0x50, 0xf5, 0x9f, 0x8f, 0x3f, 0x88, 0x5e, 0x9a, 0x98, 0xbd, 0x36, 0x31,
	0x7b, 0x6b, 0x62, 0xf6, 0xfc, 0x1e, 0x2f, 0x5d, 0xf7, 0xe8, 0xc5, 0xb7, 0x3f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x8a, 0xd3, 0x92, 0x7b, 0x02, 0x02, 0x00, 0x00,
}
