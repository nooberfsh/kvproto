// Code generated by protoc-gen-gogo.
// source: statpb.proto
// DO NOT EDIT!

/*
	Package statpb is a generated protocol buffer package.

	It is generated from these files:
		statpb.proto

	It has these top-level messages:
		PingReq
		PingResp
		RequestHeader
		ResponseHeader
		Request
		Response
*/
package statpb

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_Ping MessageType = 0
)

var MessageType_name = map[int32]string{
	0: "Ping",
}
var MessageType_value = map[string]int32{
	"Ping": 0,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptorStatpb, []int{0} }

type PingReq struct {
	FromStoreId      *uint64 `protobuf:"varint,1,opt,name=from_store_id,json=fromStoreId" json:"from_store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PingReq) Reset()                    { *m = PingReq{} }
func (m *PingReq) String() string            { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()               {}
func (*PingReq) Descriptor() ([]byte, []int) { return fileDescriptorStatpb, []int{0} }

func (m *PingReq) GetFromStoreId() uint64 {
	if m != nil && m.FromStoreId != nil {
		return *m.FromStoreId
	}
	return 0
}

type PingResp struct {
	FromStoreId      *uint64 `protobuf:"varint,1,opt,name=from_store_id,json=fromStoreId" json:"from_store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PingResp) Reset()                    { *m = PingResp{} }
func (m *PingResp) String() string            { return proto.CompactTextString(m) }
func (*PingResp) ProtoMessage()               {}
func (*PingResp) Descriptor() ([]byte, []int) { return fileDescriptorStatpb, []int{1} }

func (m *PingResp) GetFromStoreId() uint64 {
	if m != nil && m.FromStoreId != nil {
		return *m.FromStoreId
	}
	return 0
}

type RequestHeader struct {
	Id               *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestHeader) Reset()                    { *m = RequestHeader{} }
func (m *RequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()               {}
func (*RequestHeader) Descriptor() ([]byte, []int) { return fileDescriptorStatpb, []int{2} }

func (m *RequestHeader) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type ResponseHeader struct {
	Id               *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptorStatpb, []int{3} }

func (m *ResponseHeader) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type Request struct {
	Header           *RequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MsgType          MessageType    `protobuf:"varint,2,opt,name=msg_type,json=msgType,enum=statpb.MessageType" json:"msg_type"`
	PingReq          *PingReq       `protobuf:"bytes,3,opt,name=ping_req,json=pingReq" json:"ping_req,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorStatpb, []int{4} }

func (m *Request) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetMsgType() MessageType {
	if m != nil {
		return m.MsgType
	}
	return MessageType_Ping
}

func (m *Request) GetPingReq() *PingReq {
	if m != nil {
		return m.PingReq
	}
	return nil
}

type Response struct {
	Header           *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MsgType          MessageType     `protobuf:"varint,2,opt,name=msg_type,json=msgType,enum=statpb.MessageType" json:"msg_type"`
	PingResp         *PingResp       `protobuf:"bytes,3,opt,name=ping_resp,json=pingResp" json:"ping_resp,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorStatpb, []int{5} }

func (m *Response) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetMsgType() MessageType {
	if m != nil {
		return m.MsgType
	}
	return MessageType_Ping
}

func (m *Response) GetPingResp() *PingResp {
	if m != nil {
		return m.PingResp
	}
	return nil
}

func init() {
	proto.RegisterType((*PingReq)(nil), "statpb.PingReq")
	proto.RegisterType((*PingResp)(nil), "statpb.PingResp")
	proto.RegisterType((*RequestHeader)(nil), "statpb.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "statpb.ResponseHeader")
	proto.RegisterType((*Request)(nil), "statpb.Request")
	proto.RegisterType((*Response)(nil), "statpb.Response")
	proto.RegisterEnum("statpb.MessageType", MessageType_name, MessageType_value)
}
func (m *PingReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FromStoreId != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStatpb(dAtA, i, uint64(*m.FromStoreId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PingResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FromStoreId != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStatpb(dAtA, i, uint64(*m.FromStoreId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *RequestHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStatpb(dAtA, i, uint64(*m.Id))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ResponseHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStatpb(dAtA, i, uint64(*m.Id))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStatpb(dAtA, i, uint64(m.Header.Size()))
		n1, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	dAtA[i] = 0x10
	i++
	i = encodeVarintStatpb(dAtA, i, uint64(m.MsgType))
	if m.PingReq != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintStatpb(dAtA, i, uint64(m.PingReq.Size()))
		n2, err := m.PingReq.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStatpb(dAtA, i, uint64(m.Header.Size()))
		n3, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	dAtA[i] = 0x10
	i++
	i = encodeVarintStatpb(dAtA, i, uint64(m.MsgType))
	if m.PingResp != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintStatpb(dAtA, i, uint64(m.PingResp.Size()))
		n4, err := m.PingResp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Statpb(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Statpb(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStatpb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PingReq) Size() (n int) {
	var l int
	_ = l
	if m.FromStoreId != nil {
		n += 1 + sovStatpb(uint64(*m.FromStoreId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PingResp) Size() (n int) {
	var l int
	_ = l
	if m.FromStoreId != nil {
		n += 1 + sovStatpb(uint64(*m.FromStoreId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestHeader) Size() (n int) {
	var l int
	_ = l
	if m.Id != nil {
		n += 1 + sovStatpb(uint64(*m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseHeader) Size() (n int) {
	var l int
	_ = l
	if m.Id != nil {
		n += 1 + sovStatpb(uint64(*m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovStatpb(uint64(l))
	}
	n += 1 + sovStatpb(uint64(m.MsgType))
	if m.PingReq != nil {
		l = m.PingReq.Size()
		n += 1 + l + sovStatpb(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovStatpb(uint64(l))
	}
	n += 1 + sovStatpb(uint64(m.MsgType))
	if m.PingResp != nil {
		l = m.PingResp.Size()
		n += 1 + l + sovStatpb(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStatpb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStatpb(x uint64) (n int) {
	return sovStatpb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PingReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatpb
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
			return fmt.Errorf("proto: PingReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromStoreId", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FromStoreId = &v
		default:
			iNdEx = preIndex
			skippy, err := skipStatpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PingResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatpb
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
			return fmt.Errorf("proto: PingResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromStoreId", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FromStoreId = &v
		default:
			iNdEx = preIndex
			skippy, err := skipStatpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatpb
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
			return fmt.Errorf("proto: RequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Id = &v
		default:
			iNdEx = preIndex
			skippy, err := skipStatpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatpb
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
			return fmt.Errorf("proto: ResponseHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Id = &v
		default:
			iNdEx = preIndex
			skippy, err := skipStatpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatpb
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
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
				return ErrInvalidLengthStatpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &RequestHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			m.MsgType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgType |= (MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PingReq", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
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
				return ErrInvalidLengthStatpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PingReq == nil {
				m.PingReq = &PingReq{}
			}
			if err := m.PingReq.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatpb
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
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
				return ErrInvalidLengthStatpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			m.MsgType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgType |= (MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PingResp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatpb
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
				return ErrInvalidLengthStatpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PingResp == nil {
				m.PingResp = &PingResp{}
			}
			if err := m.PingResp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStatpb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatpb
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
					return 0, ErrIntOverflowStatpb
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
					return 0, ErrIntOverflowStatpb
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
				return 0, ErrInvalidLengthStatpb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStatpb
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
				next, err := skipStatpb(dAtA[start:])
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
	ErrInvalidLengthStatpb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatpb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("statpb.proto", fileDescriptorStatpb) }

var fileDescriptorStatpb = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x80, 0xbb, 0xb5, 0x34, 0xeb, 0xd4, 0xd6, 0xb2, 0xfe, 0x15, 0x0f, 0x69, 0xc9, 0xa9, 0x14,
	0x1a, 0xa1, 0xf8, 0x04, 0x3d, 0xe9, 0x41, 0x90, 0xd5, 0x7b, 0xa8, 0x64, 0x5c, 0x73, 0x68, 0x76,
	0x93, 0x5d, 0x0f, 0x7d, 0x13, 0xf1, 0xe2, 0xeb, 0xf4, 0xe8, 0x13, 0x88, 0xc4, 0x17, 0x91, 0x4d,
	0x36, 0xda, 0x0a, 0x82, 0xe0, 0x6d, 0x76, 0xe6, 0x9b, 0x99, 0x6f, 0x16, 0xf6, 0xb4, 0x59, 0x18,
	0x75, 0x17, 0xaa, 0x5c, 0x1a, 0xc9, 0xda, 0xd5, 0xeb, 0xf4, 0x50, 0x48, 0x21, 0xcb, 0xd4, 0x99,
	0x8d, 0xaa, 0x6a, 0x30, 0x05, 0xef, 0x3a, 0x49, 0x05, 0xc7, 0x8c, 0x05, 0xd0, 0xbd, 0xcf, 0xe5,
	0x32, 0xd2, 0x46, 0xe6, 0x18, 0x25, 0xf1, 0x80, 0x8c, 0xc8, 0xb8, 0xc5, 0x3b, 0x36, 0x79, 0x63,
	0x73, 0x97, 0x71, 0x10, 0x02, 0xad, 0x70, 0xad, 0xfe, 0xc4, 0x0f, 0xa1, 0xcb, 0x31, 0x7b, 0x44,
	0x6d, 0x2e, 0x70, 0x11, 0x63, 0xce, 0x7a, 0xd0, 0xfc, 0x22, 0x9b, 0x49, 0x1c, 0x8c, 0xa0, 0x67,
	0x87, 0xc9, 0x54, 0xe3, 0x2f, 0xc4, 0x33, 0x01, 0xcf, 0xcd, 0x60, 0x53, 0x68, 0x3f, 0x94, 0x54,
	0x59, 0xef, 0xcc, 0x8e, 0x42, 0x77, 0xea, 0xd6, 0x12, 0xee, 0x20, 0x76, 0x0e, 0x74, 0xa9, 0x45,
	0x64, 0x56, 0x0a, 0x07, 0xcd, 0x11, 0x19, 0xf7, 0x66, 0x07, 0x75, 0xc3, 0x15, 0x6a, 0xbd, 0x10,
	0x78, 0xbb, 0x52, 0x38, 0x6f, 0xad, 0xdf, 0x86, 0x0d, 0xee, 0x2d, 0xb5, 0xb0, 0x4f, 0x36, 0x01,
	0xaa, 0x92, 0x54, 0x44, 0x39, 0x66, 0x83, 0x9d, 0x72, 0xcd, 0x7e, 0xdd, 0xe5, 0xbe, 0x8a, 0x7b,
	0xaa, 0x0a, 0x82, 0x17, 0x02, 0xb4, 0xf6, 0x67, 0xe1, 0x0f, 0xbb, 0xe3, 0x6f, 0xbb, 0xcd, 0x0b,
	0xff, 0xa9, 0x37, 0x85, 0x5d, 0xa7, 0xa7, 0x95, 0xf3, 0xeb, 0x6f, 0xfb, 0x69, 0xc5, 0xa9, 0x72,
	0xd1, 0xe4, 0x04, 0x3a, 0x1b, 0xc3, 0x18, 0x85, 0x96, 0x85, 0xfa, 0x8d, 0x79, 0x7f, 0x5d, 0xf8,
	0xe4, 0xb5, 0xf0, 0xc9, 0x7b, 0xe1, 0x93, 0xa7, 0x0f, 0xbf, 0xf1, 0x19, 0x00, 0x00, 0xff, 0xff,
	0x42, 0x49, 0x72, 0xf7, 0x38, 0x02, 0x00, 0x00,
}
