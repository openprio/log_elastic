// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openprio_location/openprio_pt_position_data.proto

package openprio

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

// This enum describes the status of the door of a bus.
type DoorOpeningStatus int32

const (
	DoorOpeningStatus_CLOSED  DoorOpeningStatus = 0
	DoorOpeningStatus_OPEN    DoorOpeningStatus = 1
	DoorOpeningStatus_NO_DATA DoorOpeningStatus = 2
)

var DoorOpeningStatus_name = map[int32]string{
	0: "CLOSED",
	1: "OPEN",
	2: "NO_DATA",
}

var DoorOpeningStatus_value = map[string]int32{
	"CLOSED":  0,
	"OPEN":    1,
	"NO_DATA": 2,
}

func (x DoorOpeningStatus) String() string {
	return proto.EnumName(DoorOpeningStatus_name, int32(x))
}

func (DoorOpeningStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75d034d1a50152ca, []int{0}
}

// Message that vehicle sends every second.
type LocationMessage struct {
	VehicleDescriptor    *VehicleDescriptor `protobuf:"bytes,1,opt,name=vehicle_descriptor,json=vehicleDescriptor,proto3" json:"vehicle_descriptor,omitempty"`
	Position             *Position          `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Timestamp            int64              `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DoorStatus           DoorOpeningStatus  `protobuf:"varint,4,opt,name=door_status,json=doorStatus,proto3,enum=DoorOpeningStatus" json:"door_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LocationMessage) Reset()         { *m = LocationMessage{} }
func (m *LocationMessage) String() string { return proto.CompactTextString(m) }
func (*LocationMessage) ProtoMessage()    {}
func (*LocationMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_75d034d1a50152ca, []int{0}
}

func (m *LocationMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationMessage.Unmarshal(m, b)
}
func (m *LocationMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationMessage.Marshal(b, m, deterministic)
}
func (m *LocationMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationMessage.Merge(m, src)
}
func (m *LocationMessage) XXX_Size() int {
	return xxx_messageInfo_LocationMessage.Size(m)
}
func (m *LocationMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LocationMessage proto.InternalMessageInfo

func (m *LocationMessage) GetVehicleDescriptor() *VehicleDescriptor {
	if m != nil {
		return m.VehicleDescriptor
	}
	return nil
}

func (m *LocationMessage) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *LocationMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *LocationMessage) GetDoorStatus() DoorOpeningStatus {
	if m != nil {
		return m.DoorStatus
	}
	return DoorOpeningStatus_CLOSED
}

type Position struct {
	Latitude             float32  `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Accuracy             float32  `protobuf:"fixed32,3,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	Speed                float32  `protobuf:"fixed32,4,opt,name=speed,proto3" json:"speed,omitempty"`
	Bearing              float32  `protobuf:"fixed32,5,opt,name=bearing,proto3" json:"bearing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_75d034d1a50152ca, []int{1}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Position) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Position) GetAccuracy() float32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func (m *Position) GetSpeed() float32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *Position) GetBearing() float32 {
	if m != nil {
		return m.Bearing
	}
	return 0
}

// Fields to identify vehicle
type VehicleDescriptor struct {
	DataOwnerCode string `protobuf:"bytes,1,opt,name=data_owner_code,json=dataOwnerCode,proto3" json:"data_owner_code,omitempty"`
	// One of block_code and vehicle_number is required
	BlockCode            int32    `protobuf:"varint,2,opt,name=block_code,json=blockCode,proto3" json:"block_code,omitempty"`
	VehicleNumber        int32    `protobuf:"varint,3,opt,name=vehicle_number,json=vehicleNumber,proto3" json:"vehicle_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VehicleDescriptor) Reset()         { *m = VehicleDescriptor{} }
func (m *VehicleDescriptor) String() string { return proto.CompactTextString(m) }
func (*VehicleDescriptor) ProtoMessage()    {}
func (*VehicleDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_75d034d1a50152ca, []int{2}
}

func (m *VehicleDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehicleDescriptor.Unmarshal(m, b)
}
func (m *VehicleDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehicleDescriptor.Marshal(b, m, deterministic)
}
func (m *VehicleDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehicleDescriptor.Merge(m, src)
}
func (m *VehicleDescriptor) XXX_Size() int {
	return xxx_messageInfo_VehicleDescriptor.Size(m)
}
func (m *VehicleDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_VehicleDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_VehicleDescriptor proto.InternalMessageInfo

func (m *VehicleDescriptor) GetDataOwnerCode() string {
	if m != nil {
		return m.DataOwnerCode
	}
	return ""
}

func (m *VehicleDescriptor) GetBlockCode() int32 {
	if m != nil {
		return m.BlockCode
	}
	return 0
}

func (m *VehicleDescriptor) GetVehicleNumber() int32 {
	if m != nil {
		return m.VehicleNumber
	}
	return 0
}

func init() {
	proto.RegisterEnum("DoorOpeningStatus", DoorOpeningStatus_name, DoorOpeningStatus_value)
	proto.RegisterType((*LocationMessage)(nil), "LocationMessage")
	proto.RegisterType((*Position)(nil), "Position")
	proto.RegisterType((*VehicleDescriptor)(nil), "VehicleDescriptor")
}

func init() {
	proto.RegisterFile("openprio_location/openprio_pt_position_data.proto", fileDescriptor_75d034d1a50152ca)
}

var fileDescriptor_75d034d1a50152ca = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6a, 0x1b, 0x31,
	0x10, 0xc6, 0xab, 0x4d, 0x9c, 0xd8, 0x63, 0x92, 0xd8, 0xa2, 0x87, 0xa5, 0xb4, 0x60, 0x0c, 0x29,
	0xa6, 0x07, 0x97, 0x26, 0xd0, 0xbb, 0x89, 0x73, 0x4b, 0xbd, 0x41, 0x29, 0xbd, 0x0a, 0x59, 0x2b,
	0xb6, 0xa2, 0x6b, 0x8d, 0x90, 0xb4, 0x29, 0x3d, 0xf6, 0x0d, 0xfa, 0x64, 0x7d, 0xa6, 0x22, 0xed,
	0x9f, 0x40, 0x7d, 0xd3, 0xf7, 0xfb, 0x66, 0x98, 0x8f, 0x19, 0xc1, 0x27, 0xb4, 0xca, 0x58, 0xa7,
	0x91, 0xd7, 0x28, 0x45, 0xd0, 0x68, 0x3e, 0x0e, 0xc4, 0x06, 0x6e, 0xd1, 0xeb, 0x08, 0x79, 0x29,
	0x82, 0x58, 0x5b, 0x87, 0x01, 0x97, 0x7f, 0x09, 0x5c, 0x3d, 0x74, 0xc5, 0x5f, 0x94, 0xf7, 0xa2,
	0x52, 0x74, 0x03, 0xf4, 0x59, 0x7d, 0xd7, 0xb2, 0x56, 0xbc, 0x54, 0x5e, 0x3a, 0x6d, 0x03, 0xba,
	0x9c, 0x2c, 0xc8, 0x6a, 0x7a, 0x43, 0xd7, 0xdf, 0x5a, 0x6b, 0x3b, 0x38, 0x6c, 0xfe, 0xfc, 0x3f,
	0xa2, 0xd7, 0x30, 0xee, 0xa7, 0xe5, 0x59, 0x6a, 0x9c, 0xac, 0x1f, 0x3b, 0xc0, 0x06, 0x8b, 0xbe,
	0x85, 0x49, 0xd0, 0x07, 0xe5, 0x83, 0x38, 0xd8, 0xfc, 0x64, 0x41, 0x56, 0x27, 0xec, 0x05, 0xd0,
	0x5b, 0x98, 0x96, 0x88, 0x8e, 0xfb, 0x20, 0x42, 0xe3, 0xf3, 0xd3, 0x05, 0x59, 0x5d, 0xde, 0xd0,
	0xf5, 0x16, 0xd1, 0x15, 0x56, 0x19, 0x6d, 0xaa, 0xa7, 0xe4, 0x30, 0x88, 0x65, 0xed, 0x7b, 0xf9,
	0x87, 0xc0, 0xb8, 0x9f, 0x44, 0xdf, 0xc0, 0xb8, 0x16, 0x41, 0x87, 0xa6, 0x54, 0x29, 0x7f, 0xc6,
	0x06, 0x1d, 0x67, 0xd7, 0x68, 0xaa, 0xd6, 0xcc, 0x92, 0xf9, 0x02, 0x62, 0xa7, 0x90, 0xb2, 0x71,
	0x42, 0xfe, 0x4a, 0xc1, 0x32, 0x36, 0x68, 0xfa, 0x1a, 0x46, 0xde, 0x2a, 0x55, 0xa6, 0x44, 0x19,
	0x6b, 0x05, 0xcd, 0xe1, 0x7c, 0xaf, 0x84, 0xd3, 0xa6, 0xca, 0x47, 0x89, 0xf7, 0x72, 0xf9, 0x9b,
	0xc0, 0xfc, 0x68, 0x6b, 0xf4, 0x3d, 0x5c, 0xc5, 0x3b, 0x70, 0xfc, 0x69, 0x94, 0xe3, 0x12, 0xbb,
	0x88, 0x13, 0x76, 0x11, 0x71, 0x11, 0xe9, 0x1d, 0x96, 0x8a, 0xbe, 0x03, 0xd8, 0xd7, 0x28, 0x7f,
	0xb4, 0x25, 0x31, 0xe8, 0x88, 0x4d, 0x12, 0x49, 0xf6, 0x35, 0x5c, 0xf6, 0xc7, 0x32, 0xcd, 0x61,
	0xaf, 0x5c, 0x8a, 0x3b, 0x62, 0x17, 0x1d, 0xdd, 0x25, 0xf8, 0xe1, 0x33, 0xcc, 0x8f, 0xf6, 0x46,
	0x01, 0xce, 0xee, 0x1e, 0x8a, 0xa7, 0xfb, 0xed, 0xec, 0x15, 0x1d, 0xc3, 0x69, 0xf1, 0x78, 0xbf,
	0x9b, 0x11, 0x3a, 0x85, 0xf3, 0x5d, 0xc1, 0xb7, 0x9b, 0xaf, 0x9b, 0x59, 0xb6, 0x3f, 0x4b, 0xdf,
	0xe4, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x52, 0x79, 0xec, 0x5b, 0x02, 0x00, 0x00,
}
