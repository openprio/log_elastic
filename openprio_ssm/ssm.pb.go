// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssm.proto

package ssm

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

// SSM V2.1 specification PrioritizationResponseStatus values
type SignalStatusPackage_PrioritizationResponseStatus int32

const (
	SignalStatusPackage_UNKNOWN           SignalStatusPackage_PrioritizationResponseStatus = 0
	SignalStatusPackage_REQUESTED         SignalStatusPackage_PrioritizationResponseStatus = 1
	SignalStatusPackage_PROCESSING        SignalStatusPackage_PrioritizationResponseStatus = 2
	SignalStatusPackage_WATCHOTHERTRAFFIC SignalStatusPackage_PrioritizationResponseStatus = 3
	SignalStatusPackage_GRANTED           SignalStatusPackage_PrioritizationResponseStatus = 4
	SignalStatusPackage_REJECTED          SignalStatusPackage_PrioritizationResponseStatus = 5
	SignalStatusPackage_MAXPRESENCE       SignalStatusPackage_PrioritizationResponseStatus = 6
	SignalStatusPackage_RESERVICELOCKED   SignalStatusPackage_PrioritizationResponseStatus = 7
)

var SignalStatusPackage_PrioritizationResponseStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "REQUESTED",
	2: "PROCESSING",
	3: "WATCHOTHERTRAFFIC",
	4: "GRANTED",
	5: "REJECTED",
	6: "MAXPRESENCE",
	7: "RESERVICELOCKED",
}

var SignalStatusPackage_PrioritizationResponseStatus_value = map[string]int32{
	"UNKNOWN":           0,
	"REQUESTED":         1,
	"PROCESSING":        2,
	"WATCHOTHERTRAFFIC": 3,
	"GRANTED":           4,
	"REJECTED":          5,
	"MAXPRESENCE":       6,
	"RESERVICELOCKED":   7,
}

func (x SignalStatusPackage_PrioritizationResponseStatus) String() string {
	return proto.EnumName(SignalStatusPackage_PrioritizationResponseStatus_name, int32(x))
}

func (SignalStatusPackage_PrioritizationResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{4, 0}
}

// SSM V2.1 specification Vehicle Roles
type SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole int32

const (
	SignalStatusPackage_SignalRequesterInfo_RequestorType_BASICVEHICLE     SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole = 0
	SignalStatusPackage_SignalRequesterInfo_RequestorType_PUBLICTRANSPORT  SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole = 1
	SignalStatusPackage_SignalRequesterInfo_RequestorType_SPECIALTRANSPORT SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole = 2
	SignalStatusPackage_SignalRequesterInfo_RequestorType_DANGEROUSGOODS   SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole = 3
	SignalStatusPackage_SignalRequesterInfo_RequestorType_ROADWORK         SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole = 4
	SignalStatusPackage_SignalRequesterInfo_RequestorType_ROADRESCUE       SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole = 5
	SignalStatusPackage_SignalRequesterInfo_RequestorType_EMERGENCY        SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole = 6
	SignalStatusPackage_SignalRequesterInfo_RequestorType_SAFETYCAR        SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole = 7
)

var SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole_name = map[int32]string{
	0: "BASICVEHICLE",
	1: "PUBLICTRANSPORT",
	2: "SPECIALTRANSPORT",
	3: "DANGEROUSGOODS",
	4: "ROADWORK",
	5: "ROADRESCUE",
	6: "EMERGENCY",
	7: "SAFETYCAR",
}

var SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole_value = map[string]int32{
	"BASICVEHICLE":     0,
	"PUBLICTRANSPORT":  1,
	"SPECIALTRANSPORT": 2,
	"DANGEROUSGOODS":   3,
	"ROADWORK":         4,
	"ROADRESCUE":       5,
	"EMERGENCY":        6,
	"SAFETYCAR":        7,
}

func (x SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole) String() string {
	return proto.EnumName(SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole_name, int32(x))
}

func (SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{4, 0, 1, 0}
}

// SSM V2.1 specification SubRoles
type SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole int32

const (
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLEUNKNOWN  SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 0
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE1        SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 1
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE2        SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 2
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE3        SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 3
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE4        SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 4
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE5        SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 5
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE6        SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 6
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE7        SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 7
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE8        SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 8
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE9        SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 9
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE10       SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 10
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE11       SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 11
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE12       SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 12
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE13       SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 13
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLE14       SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 14
	SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLERESERVED SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole = 15
)

var SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole_name = map[int32]string{
	0:  "REQUESTSUBROLEUNKNOWN",
	1:  "REQUESTSUBROLE1",
	2:  "REQUESTSUBROLE2",
	3:  "REQUESTSUBROLE3",
	4:  "REQUESTSUBROLE4",
	5:  "REQUESTSUBROLE5",
	6:  "REQUESTSUBROLE6",
	7:  "REQUESTSUBROLE7",
	8:  "REQUESTSUBROLE8",
	9:  "REQUESTSUBROLE9",
	10: "REQUESTSUBROLE10",
	11: "REQUESTSUBROLE11",
	12: "REQUESTSUBROLE12",
	13: "REQUESTSUBROLE13",
	14: "REQUESTSUBROLE14",
	15: "REQUESTSUBROLERESERVED",
}

var SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole_value = map[string]int32{
	"REQUESTSUBROLEUNKNOWN":  0,
	"REQUESTSUBROLE1":        1,
	"REQUESTSUBROLE2":        2,
	"REQUESTSUBROLE3":        3,
	"REQUESTSUBROLE4":        4,
	"REQUESTSUBROLE5":        5,
	"REQUESTSUBROLE6":        6,
	"REQUESTSUBROLE7":        7,
	"REQUESTSUBROLE8":        8,
	"REQUESTSUBROLE9":        9,
	"REQUESTSUBROLE10":       10,
	"REQUESTSUBROLE11":       11,
	"REQUESTSUBROLE12":       12,
	"REQUESTSUBROLE13":       13,
	"REQUESTSUBROLE14":       14,
	"REQUESTSUBROLERESERVED": 15,
}

func (x SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole) String() string {
	return proto.EnumName(SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole_name, int32(x))
}

func (SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{4, 0, 1, 1}
}

// ExtendedSSM consists of the original SSM header and body as defined by the 2.1 specification with the addition of a VehicleDescriptor message for vehicle identification
type ExtendedSSM struct {
	Header               *Header            `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Ssm                  *SSM               `protobuf:"bytes,2,opt,name=ssm,proto3" json:"ssm,omitempty"`
	VehicleDescriptor    *VehicleDescriptor `protobuf:"bytes,3,opt,name=vehicleDescriptor,proto3" json:"vehicleDescriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExtendedSSM) Reset()         { *m = ExtendedSSM{} }
func (m *ExtendedSSM) String() string { return proto.CompactTextString(m) }
func (*ExtendedSSM) ProtoMessage()    {}
func (*ExtendedSSM) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{0}
}

func (m *ExtendedSSM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtendedSSM.Unmarshal(m, b)
}
func (m *ExtendedSSM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtendedSSM.Marshal(b, m, deterministic)
}
func (m *ExtendedSSM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedSSM.Merge(m, src)
}
func (m *ExtendedSSM) XXX_Size() int {
	return xxx_messageInfo_ExtendedSSM.Size(m)
}
func (m *ExtendedSSM) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedSSM.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedSSM proto.InternalMessageInfo

func (m *ExtendedSSM) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ExtendedSSM) GetSsm() *SSM {
	if m != nil {
		return m.Ssm
	}
	return nil
}

func (m *ExtendedSSM) GetVehicleDescriptor() *VehicleDescriptor {
	if m != nil {
		return m.VehicleDescriptor
	}
	return nil
}

// SSM V2.1 specification SSM Header definition
type Header struct {
	ProtocolVersion      int32    `protobuf:"varint,1,opt,name=protocolVersion,proto3" json:"protocolVersion,omitempty"`
	MessageID            int32    `protobuf:"varint,2,opt,name=messageID,proto3" json:"messageID,omitempty"`
	StationID            int64    `protobuf:"varint,3,opt,name=stationID,proto3" json:"stationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{1}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetProtocolVersion() int32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

func (m *Header) GetMessageID() int32 {
	if m != nil {
		return m.MessageID
	}
	return 0
}

func (m *Header) GetStationID() int64 {
	if m != nil {
		return m.StationID
	}
	return 0
}

// SSM V2.1 specification SSM definition
// Remarks:
// - Fields that are not used in the V2.1 specification are omitted from this protobuf definition
// - All types defined in the V2.1 spec which only contain one field with a scalar value or string (ex. MinuteOfTheYear, Dsecond...) are simplified and are not defined as separate types for simplification
type SSM struct {
	TimeStamp            int64           `protobuf:"varint,1,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	Second               int64           `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
	SequenceNumber       int32           `protobuf:"varint,3,opt,name=sequenceNumber,proto3" json:"sequenceNumber,omitempty"`
	Status               []*SignalStatus `protobuf:"bytes,4,rep,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SSM) Reset()         { *m = SSM{} }
func (m *SSM) String() string { return proto.CompactTextString(m) }
func (*SSM) ProtoMessage()    {}
func (*SSM) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{2}
}

func (m *SSM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSM.Unmarshal(m, b)
}
func (m *SSM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSM.Marshal(b, m, deterministic)
}
func (m *SSM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSM.Merge(m, src)
}
func (m *SSM) XXX_Size() int {
	return xxx_messageInfo_SSM.Size(m)
}
func (m *SSM) XXX_DiscardUnknown() {
	xxx_messageInfo_SSM.DiscardUnknown(m)
}

var xxx_messageInfo_SSM proto.InternalMessageInfo

func (m *SSM) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *SSM) GetSecond() int64 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *SSM) GetSequenceNumber() int32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *SSM) GetStatus() []*SignalStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// SSM V2.1 specification SignalStatus data frame
type SignalStatus struct {
	SequenceNumber       int32                                 `protobuf:"varint,1,opt,name=sequenceNumber,proto3" json:"sequenceNumber,omitempty"`
	Id                   *SignalStatus_IntersectionReferenceID `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	SigStatus            []*SignalStatusPackage                `protobuf:"bytes,3,rep,name=sigStatus,proto3" json:"sigStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *SignalStatus) Reset()         { *m = SignalStatus{} }
func (m *SignalStatus) String() string { return proto.CompactTextString(m) }
func (*SignalStatus) ProtoMessage()    {}
func (*SignalStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{3}
}

func (m *SignalStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalStatus.Unmarshal(m, b)
}
func (m *SignalStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalStatus.Marshal(b, m, deterministic)
}
func (m *SignalStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalStatus.Merge(m, src)
}
func (m *SignalStatus) XXX_Size() int {
	return xxx_messageInfo_SignalStatus.Size(m)
}
func (m *SignalStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SignalStatus proto.InternalMessageInfo

func (m *SignalStatus) GetSequenceNumber() int32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *SignalStatus) GetId() *SignalStatus_IntersectionReferenceID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SignalStatus) GetSigStatus() []*SignalStatusPackage {
	if m != nil {
		return m.SigStatus
	}
	return nil
}

// SSM V2.1 specification IntersectionReferenceID Type as defined at level 1.2
type SignalStatus_IntersectionReferenceID struct {
	Region               int64    `protobuf:"varint,1,opt,name=region,proto3" json:"region,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalStatus_IntersectionReferenceID) Reset()         { *m = SignalStatus_IntersectionReferenceID{} }
func (m *SignalStatus_IntersectionReferenceID) String() string { return proto.CompactTextString(m) }
func (*SignalStatus_IntersectionReferenceID) ProtoMessage()    {}
func (*SignalStatus_IntersectionReferenceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{3, 0}
}

func (m *SignalStatus_IntersectionReferenceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalStatus_IntersectionReferenceID.Unmarshal(m, b)
}
func (m *SignalStatus_IntersectionReferenceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalStatus_IntersectionReferenceID.Marshal(b, m, deterministic)
}
func (m *SignalStatus_IntersectionReferenceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalStatus_IntersectionReferenceID.Merge(m, src)
}
func (m *SignalStatus_IntersectionReferenceID) XXX_Size() int {
	return xxx_messageInfo_SignalStatus_IntersectionReferenceID.Size(m)
}
func (m *SignalStatus_IntersectionReferenceID) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalStatus_IntersectionReferenceID.DiscardUnknown(m)
}

var xxx_messageInfo_SignalStatus_IntersectionReferenceID proto.InternalMessageInfo

func (m *SignalStatus_IntersectionReferenceID) GetRegion() int64 {
	if m != nil {
		return m.Region
	}
	return 0
}

func (m *SignalStatus_IntersectionReferenceID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// SSM V2.1 specification SignalStatusPackage data frame
type SignalStatusPackage struct {
	Requestor *SignalStatusPackage_SignalRequesterInfo `protobuf:"bytes,1,opt,name=requestor,proto3" json:"requestor,omitempty"`
	// inboundOn should either be approach or connection based on the SRM, [IntersectionAccessPoint], required
	//
	// Types that are valid to be assigned to InboundOn:
	//	*SignalStatusPackage_Approach
	//	*SignalStatusPackage_Connection
	InboundOn            isSignalStatusPackage_InboundOn                  `protobuf_oneof:"inboundOn"`
	Minute               int64                                            `protobuf:"varint,4,opt,name=minute,proto3" json:"minute,omitempty"`
	Second               int64                                            `protobuf:"varint,5,opt,name=second,proto3" json:"second,omitempty"`
	Duration             int64                                            `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Status               SignalStatusPackage_PrioritizationResponseStatus `protobuf:"varint,7,opt,name=status,proto3,enum=SignalStatusPackage_PrioritizationResponseStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *SignalStatusPackage) Reset()         { *m = SignalStatusPackage{} }
func (m *SignalStatusPackage) String() string { return proto.CompactTextString(m) }
func (*SignalStatusPackage) ProtoMessage()    {}
func (*SignalStatusPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{4}
}

func (m *SignalStatusPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalStatusPackage.Unmarshal(m, b)
}
func (m *SignalStatusPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalStatusPackage.Marshal(b, m, deterministic)
}
func (m *SignalStatusPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalStatusPackage.Merge(m, src)
}
func (m *SignalStatusPackage) XXX_Size() int {
	return xxx_messageInfo_SignalStatusPackage.Size(m)
}
func (m *SignalStatusPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalStatusPackage.DiscardUnknown(m)
}

var xxx_messageInfo_SignalStatusPackage proto.InternalMessageInfo

func (m *SignalStatusPackage) GetRequestor() *SignalStatusPackage_SignalRequesterInfo {
	if m != nil {
		return m.Requestor
	}
	return nil
}

type isSignalStatusPackage_InboundOn interface {
	isSignalStatusPackage_InboundOn()
}

type SignalStatusPackage_Approach struct {
	Approach int64 `protobuf:"varint,2,opt,name=approach,proto3,oneof"`
}

type SignalStatusPackage_Connection struct {
	Connection int64 `protobuf:"varint,3,opt,name=connection,proto3,oneof"`
}

func (*SignalStatusPackage_Approach) isSignalStatusPackage_InboundOn() {}

func (*SignalStatusPackage_Connection) isSignalStatusPackage_InboundOn() {}

func (m *SignalStatusPackage) GetInboundOn() isSignalStatusPackage_InboundOn {
	if m != nil {
		return m.InboundOn
	}
	return nil
}

func (m *SignalStatusPackage) GetApproach() int64 {
	if x, ok := m.GetInboundOn().(*SignalStatusPackage_Approach); ok {
		return x.Approach
	}
	return 0
}

func (m *SignalStatusPackage) GetConnection() int64 {
	if x, ok := m.GetInboundOn().(*SignalStatusPackage_Connection); ok {
		return x.Connection
	}
	return 0
}

func (m *SignalStatusPackage) GetMinute() int64 {
	if m != nil {
		return m.Minute
	}
	return 0
}

func (m *SignalStatusPackage) GetSecond() int64 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *SignalStatusPackage) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *SignalStatusPackage) GetStatus() SignalStatusPackage_PrioritizationResponseStatus {
	if m != nil {
		return m.Status
	}
	return SignalStatusPackage_UNKNOWN
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignalStatusPackage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignalStatusPackage_Approach)(nil),
		(*SignalStatusPackage_Connection)(nil),
	}
}

// SSM V2.1 specification SignalRequesterInfo Type as defined at level 2.1
type SignalStatusPackage_SignalRequesterInfo struct {
	Id                   *SignalStatusPackage_SignalRequesterInfo_VehicleID     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Request              string                                                 `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	SequenceNumber       int32                                                  `protobuf:"varint,3,opt,name=sequenceNumber,proto3" json:"sequenceNumber,omitempty"`
	TypeData             *SignalStatusPackage_SignalRequesterInfo_RequestorType `protobuf:"bytes,4,opt,name=typeData,proto3" json:"typeData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                               `json:"-"`
	XXX_unrecognized     []byte                                                 `json:"-"`
	XXX_sizecache        int32                                                  `json:"-"`
}

func (m *SignalStatusPackage_SignalRequesterInfo) Reset() {
	*m = SignalStatusPackage_SignalRequesterInfo{}
}
func (m *SignalStatusPackage_SignalRequesterInfo) String() string { return proto.CompactTextString(m) }
func (*SignalStatusPackage_SignalRequesterInfo) ProtoMessage()    {}
func (*SignalStatusPackage_SignalRequesterInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{4, 0}
}

func (m *SignalStatusPackage_SignalRequesterInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo.Unmarshal(m, b)
}
func (m *SignalStatusPackage_SignalRequesterInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo.Marshal(b, m, deterministic)
}
func (m *SignalStatusPackage_SignalRequesterInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo.Merge(m, src)
}
func (m *SignalStatusPackage_SignalRequesterInfo) XXX_Size() int {
	return xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo.Size(m)
}
func (m *SignalStatusPackage_SignalRequesterInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo proto.InternalMessageInfo

func (m *SignalStatusPackage_SignalRequesterInfo) GetId() *SignalStatusPackage_SignalRequesterInfo_VehicleID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SignalStatusPackage_SignalRequesterInfo) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *SignalStatusPackage_SignalRequesterInfo) GetSequenceNumber() int32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *SignalStatusPackage_SignalRequesterInfo) GetTypeData() *SignalStatusPackage_SignalRequesterInfo_RequestorType {
	if m != nil {
		return m.TypeData
	}
	return nil
}

// SSM V2.1 specification Vehicle ID as defined at level 3
type SignalStatusPackage_SignalRequesterInfo_VehicleID struct {
	StationID            int64    `protobuf:"varint,1,opt,name=stationID,proto3" json:"stationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalStatusPackage_SignalRequesterInfo_VehicleID) Reset() {
	*m = SignalStatusPackage_SignalRequesterInfo_VehicleID{}
}
func (m *SignalStatusPackage_SignalRequesterInfo_VehicleID) String() string {
	return proto.CompactTextString(m)
}
func (*SignalStatusPackage_SignalRequesterInfo_VehicleID) ProtoMessage() {}
func (*SignalStatusPackage_SignalRequesterInfo_VehicleID) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{4, 0, 0}
}

func (m *SignalStatusPackage_SignalRequesterInfo_VehicleID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_VehicleID.Unmarshal(m, b)
}
func (m *SignalStatusPackage_SignalRequesterInfo_VehicleID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_VehicleID.Marshal(b, m, deterministic)
}
func (m *SignalStatusPackage_SignalRequesterInfo_VehicleID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_VehicleID.Merge(m, src)
}
func (m *SignalStatusPackage_SignalRequesterInfo_VehicleID) XXX_Size() int {
	return xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_VehicleID.Size(m)
}
func (m *SignalStatusPackage_SignalRequesterInfo_VehicleID) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_VehicleID.DiscardUnknown(m)
}

var xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_VehicleID proto.InternalMessageInfo

func (m *SignalStatusPackage_SignalRequesterInfo_VehicleID) GetStationID() int64 {
	if m != nil {
		return m.StationID
	}
	return 0
}

// SSM V2.1 specification Requestor Type as defined at level 4
type SignalStatusPackage_SignalRequesterInfo_RequestorType struct {
	Role                 SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole `protobuf:"varint,1,opt,name=role,proto3,enum=SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole" json:"role,omitempty"`
	SubRole              SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole   `protobuf:"varint,2,opt,name=subRole,proto3,enum=SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole" json:"subRole,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                               `json:"-"`
	XXX_unrecognized     []byte                                                                 `json:"-"`
	XXX_sizecache        int32                                                                  `json:"-"`
}

func (m *SignalStatusPackage_SignalRequesterInfo_RequestorType) Reset() {
	*m = SignalStatusPackage_SignalRequesterInfo_RequestorType{}
}
func (m *SignalStatusPackage_SignalRequesterInfo_RequestorType) String() string {
	return proto.CompactTextString(m)
}
func (*SignalStatusPackage_SignalRequesterInfo_RequestorType) ProtoMessage() {}
func (*SignalStatusPackage_SignalRequesterInfo_RequestorType) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1ea30807afcb2c, []int{4, 0, 1}
}

func (m *SignalStatusPackage_SignalRequesterInfo_RequestorType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_RequestorType.Unmarshal(m, b)
}
func (m *SignalStatusPackage_SignalRequesterInfo_RequestorType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_RequestorType.Marshal(b, m, deterministic)
}
func (m *SignalStatusPackage_SignalRequesterInfo_RequestorType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_RequestorType.Merge(m, src)
}
func (m *SignalStatusPackage_SignalRequesterInfo_RequestorType) XXX_Size() int {
	return xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_RequestorType.Size(m)
}
func (m *SignalStatusPackage_SignalRequesterInfo_RequestorType) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_RequestorType.DiscardUnknown(m)
}

var xxx_messageInfo_SignalStatusPackage_SignalRequesterInfo_RequestorType proto.InternalMessageInfo

func (m *SignalStatusPackage_SignalRequesterInfo_RequestorType) GetRole() SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole {
	if m != nil {
		return m.Role
	}
	return SignalStatusPackage_SignalRequesterInfo_RequestorType_BASICVEHICLE
}

func (m *SignalStatusPackage_SignalRequesterInfo_RequestorType) GetSubRole() SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole {
	if m != nil {
		return m.SubRole
	}
	return SignalStatusPackage_SignalRequesterInfo_RequestorType_REQUESTSUBROLEUNKNOWN
}

// Fields to identify vehicle
type VehicleDescriptor struct {
	DataOwnerCode string `protobuf:"bytes,1,opt,name=data_owner_code,json=dataOwnerCode,proto3" json:"data_owner_code,omitempty"`
	// One of block_code and vehicle_number is
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
	return fileDescriptor_4f1ea30807afcb2c, []int{5}
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
	proto.RegisterEnum("SignalStatusPackage_PrioritizationResponseStatus", SignalStatusPackage_PrioritizationResponseStatus_name, SignalStatusPackage_PrioritizationResponseStatus_value)
	proto.RegisterEnum("SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole", SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole_name, SignalStatusPackage_SignalRequesterInfo_RequestorType_BasicVehicleRole_value)
	proto.RegisterEnum("SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole", SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole_name, SignalStatusPackage_SignalRequesterInfo_RequestorType_RequestSubRole_value)
	proto.RegisterType((*ExtendedSSM)(nil), "ExtendedSSM")
	proto.RegisterType((*Header)(nil), "Header")
	proto.RegisterType((*SSM)(nil), "SSM")
	proto.RegisterType((*SignalStatus)(nil), "SignalStatus")
	proto.RegisterType((*SignalStatus_IntersectionReferenceID)(nil), "SignalStatus.IntersectionReferenceID")
	proto.RegisterType((*SignalStatusPackage)(nil), "SignalStatusPackage")
	proto.RegisterType((*SignalStatusPackage_SignalRequesterInfo)(nil), "SignalStatusPackage.SignalRequesterInfo")
	proto.RegisterType((*SignalStatusPackage_SignalRequesterInfo_VehicleID)(nil), "SignalStatusPackage.SignalRequesterInfo.VehicleID")
	proto.RegisterType((*SignalStatusPackage_SignalRequesterInfo_RequestorType)(nil), "SignalStatusPackage.SignalRequesterInfo.RequestorType")
	proto.RegisterType((*VehicleDescriptor)(nil), "VehicleDescriptor")
}

func init() { proto.RegisterFile("ssm.proto", fileDescriptor_4f1ea30807afcb2c) }

var fileDescriptor_4f1ea30807afcb2c = []byte{
	// 1008 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xdf, 0x6f, 0xe2, 0x46,
	0x10, 0xc7, 0x63, 0xcc, 0x2f, 0x0f, 0x81, 0x6c, 0xf6, 0xee, 0x52, 0x8a, 0x52, 0x35, 0x42, 0xca,
	0x29, 0x7d, 0x41, 0x0d, 0xb9, 0xbb, 0xb6, 0x6f, 0x35, 0xf6, 0x06, 0xdc, 0x24, 0x98, 0xae, 0x21,
	0xe9, 0xa9, 0x0f, 0x91, 0x63, 0xef, 0x11, 0xeb, 0xc0, 0xa6, 0xb6, 0x69, 0x7b, 0x7d, 0xbb, 0xb7,
	0xaa, 0xff, 0x44, 0xa5, 0xfe, 0x6d, 0xf7, 0xd4, 0xc7, 0xfe, 0x05, 0xd5, 0xae, 0x0d, 0x04, 0x4c,
	0xab, 0xe8, 0x1e, 0xe7, 0x33, 0xb3, 0x33, 0xdf, 0x99, 0x9d, 0x35, 0x80, 0x12, 0x45, 0xd3, 0xd6,
	0x2c, 0x0c, 0xe2, 0xa0, 0xf9, 0xbb, 0x04, 0x15, 0xf2, 0x6b, 0xcc, 0x7c, 0x97, 0xb9, 0x96, 0x75,
	0x85, 0x3f, 0x87, 0xe2, 0x3d, 0xb3, 0x5d, 0x16, 0xd6, 0xa5, 0x23, 0xe9, 0xa4, 0xd2, 0x2e, 0xb5,
	0x7a, 0xc2, 0xa4, 0x29, 0xc6, 0x07, 0x20, 0x47, 0xd1, 0xb4, 0x9e, 0x13, 0xde, 0x7c, 0xcb, 0xb2,
	0xae, 0x28, 0x07, 0xf8, 0x5b, 0xd8, 0xff, 0x99, 0xdd, 0x7b, 0xce, 0x84, 0xe9, 0x2c, 0x72, 0x42,
	0x6f, 0x16, 0x07, 0x61, 0x5d, 0x16, 0x51, 0xb8, 0x75, 0xbd, 0xe9, 0xa1, 0xd9, 0xe0, 0xa6, 0x0f,
	0xc5, 0xa4, 0x16, 0x3e, 0x81, 0x3d, 0xa1, 0xce, 0x09, 0x26, 0xd7, 0x2c, 0x8c, 0xbc, 0xc0, 0x17,
	0x6a, 0x0a, 0x74, 0x13, 0xe3, 0x43, 0x50, 0xa6, 0x2c, 0x8a, 0xec, 0x31, 0x33, 0x74, 0xa1, 0xa9,
	0x40, 0x57, 0x80, 0x7b, 0xa3, 0xd8, 0x8e, 0xbd, 0xc0, 0x37, 0x74, 0xa1, 0x45, 0xa6, 0x2b, 0xd0,
	0xfc, 0x43, 0x02, 0x99, 0xb7, 0x7c, 0x08, 0x4a, 0xec, 0x4d, 0x99, 0x15, 0xdb, 0xd3, 0x99, 0xa8,
	0x23, 0xd3, 0x15, 0xc0, 0x07, 0x50, 0x8c, 0x98, 0x13, 0xf8, 0xae, 0x48, 0x2f, 0xd3, 0xd4, 0xc2,
	0xcf, 0xa1, 0x16, 0xb1, 0x9f, 0xe6, 0xcc, 0x77, 0x58, 0x7f, 0x3e, 0xbd, 0x63, 0x49, 0xb3, 0x05,
	0xba, 0x41, 0xf1, 0x31, 0x14, 0x79, 0xc9, 0x79, 0x54, 0xcf, 0x1f, 0xc9, 0x27, 0x95, 0x76, 0xb5,
	0x65, 0x79, 0x63, 0xdf, 0x9e, 0x58, 0x02, 0xd2, 0xd4, 0xd9, 0xfc, 0x20, 0xc1, 0xee, 0x43, 0xc7,
	0x96, 0xfc, 0xd2, 0xd6, 0xfc, 0x2f, 0x21, 0xe7, 0xb9, 0xe9, 0x75, 0x1c, 0xaf, 0xe5, 0x6e, 0x19,
	0x7e, 0xcc, 0xc2, 0x88, 0x39, 0xbc, 0x5f, 0xca, 0xde, 0xb0, 0x90, 0x1f, 0x33, 0x74, 0x9a, 0xf3,
	0x5c, 0xdc, 0x06, 0x25, 0xf2, 0xc6, 0x49, 0x60, 0x5d, 0x16, 0xca, 0x9e, 0xae, 0x9d, 0x1e, 0xd8,
	0xce, 0x5b, 0x7b, 0xcc, 0xe8, 0x2a, 0xac, 0xa1, 0xc2, 0x27, 0xff, 0x91, 0x92, 0x4f, 0x29, 0x64,
	0xe3, 0xc5, 0x45, 0xc9, 0x34, 0xb5, 0x70, 0x6d, 0xa9, 0x4e, 0xe6, 0x65, 0x9b, 0x7f, 0x57, 0xe0,
	0xc9, 0x96, 0x2a, 0xf8, 0x1c, 0x94, 0x90, 0xf7, 0x15, 0xf1, 0xad, 0x49, 0x36, 0xef, 0x64, 0x9b,
	0x9c, 0x94, 0xd1, 0x24, 0x96, 0x85, 0x86, 0xff, 0x26, 0xa0, 0xab, 0xa3, 0xf8, 0x10, 0xca, 0xf6,
	0x6c, 0x16, 0x06, 0xb6, 0x73, 0x9f, 0x54, 0xed, 0xed, 0xd0, 0x25, 0xc1, 0x47, 0x00, 0x4e, 0xe0,
	0xfb, 0x89, 0xfc, 0x64, 0x21, 0x7a, 0x3b, 0xf4, 0x01, 0xe3, 0x7d, 0x4c, 0x3d, 0x7f, 0x1e, 0xb3,
	0x7a, 0x3e, 0xe9, 0x23, 0xb1, 0x1e, 0x6c, 0x41, 0x61, 0x6d, 0x0b, 0x1a, 0x50, 0x76, 0xe7, 0xa1,
	0xd8, 0xa8, 0x7a, 0x51, 0x78, 0x96, 0x36, 0x36, 0x96, 0x37, 0x5f, 0x3a, 0x92, 0x4e, 0x6a, 0xed,
	0xd3, 0xad, 0x0d, 0x0d, 0x42, 0x2f, 0x08, 0xbd, 0xd8, 0xfb, 0xcd, 0x4e, 0x66, 0x1a, 0xcd, 0x02,
	0x3f, 0x62, 0xeb, 0xdb, 0xd1, 0x78, 0x5f, 0x5e, 0x8c, 0x6d, 0xad, 0x73, 0xdc, 0x11, 0xe3, 0x4d,
	0xe6, 0xd5, 0x7e, 0xec, 0xbc, 0x16, 0xaf, 0x31, 0xdd, 0x84, 0x3a, 0x94, 0xd2, 0xf9, 0x89, 0x89,
	0x29, 0x74, 0x61, 0x3e, 0x7a, 0xc5, 0x29, 0x94, 0xe3, 0x77, 0x33, 0xa6, 0xdb, 0xb1, 0x2d, 0xc6,
	0x56, 0x69, 0xbf, 0x7a, 0xb4, 0x16, 0xba, 0xb8, 0xba, 0xe1, 0xbb, 0x19, 0xa3, 0xcb, 0x3c, 0x8d,
	0x2f, 0x40, 0x59, 0xca, 0x5c, 0x7f, 0xc7, 0xd2, 0xc6, 0x3b, 0x6e, 0x7c, 0x28, 0x40, 0x75, 0x2d,
	0x0d, 0xfe, 0x11, 0xf2, 0x61, 0x30, 0x61, 0x22, 0xb4, 0xd6, 0xee, 0x7e, 0x9c, 0x98, 0x56, 0xc7,
	0x8e, 0x3c, 0x27, 0x15, 0x41, 0x83, 0x09, 0xa3, 0x22, 0x29, 0xbe, 0x85, 0x52, 0x34, 0xbf, 0xe3,
	0x40, 0xcc, 0xab, 0xd6, 0x26, 0x1f, 0x99, 0x3f, 0xb5, 0xac, 0x24, 0x19, 0x5d, 0x64, 0x6d, 0xfe,
	0x29, 0x01, 0xda, 0xac, 0x8d, 0x11, 0xec, 0x76, 0x54, 0xcb, 0xd0, 0xae, 0x49, 0xcf, 0xd0, 0x2e,
	0x09, 0xda, 0xc1, 0x4f, 0x60, 0x6f, 0x30, 0xea, 0x5c, 0x1a, 0xda, 0x90, 0xaa, 0x7d, 0x6b, 0x60,
	0xd2, 0x21, 0x92, 0xf0, 0x53, 0x40, 0xd6, 0x80, 0x68, 0x86, 0x7a, 0xb9, 0xa2, 0x39, 0x8c, 0xa1,
	0xa6, 0xab, 0xfd, 0x2e, 0xa1, 0xe6, 0xc8, 0xea, 0x9a, 0xa6, 0x6e, 0x21, 0x19, 0xef, 0x42, 0x99,
	0x9a, 0xaa, 0x7e, 0x63, 0xd2, 0x0b, 0x94, 0xc7, 0x35, 0x00, 0x6e, 0x51, 0x62, 0x69, 0x23, 0x82,
	0x0a, 0xb8, 0x0a, 0x0a, 0xb9, 0x22, 0xb4, 0x4b, 0xfa, 0xda, 0x6b, 0x54, 0xe4, 0xa6, 0xa5, 0x9e,
	0x93, 0xe1, 0x6b, 0x4d, 0xa5, 0xa8, 0xd4, 0xfc, 0x27, 0x07, 0xb5, 0x75, 0xf5, 0xf8, 0x53, 0x78,
	0x46, 0xc9, 0xf7, 0x23, 0x62, 0x0d, 0xad, 0x51, 0x87, 0x9a, 0x97, 0x64, 0xd4, 0xbf, 0xe8, 0x9b,
	0x37, 0xfd, 0x44, 0xe8, 0xba, 0xeb, 0x14, 0x49, 0x59, 0xd8, 0x46, 0xb9, 0x2c, 0x3c, 0x43, 0x72,
	0x16, 0xbe, 0x40, 0xf9, 0x2c, 0x7c, 0x89, 0x0a, 0x59, 0xf8, 0x0a, 0x15, 0xb3, 0xf0, 0x2b, 0x54,
	0xca, 0xc2, 0xaf, 0x51, 0x39, 0x0b, 0xbf, 0x41, 0x0a, 0x1f, 0xe8, 0x86, 0xf8, 0x2f, 0x11, 0x6c,
	0xa1, 0xa7, 0xa8, 0xb2, 0x85, 0xb6, 0xd1, 0xee, 0x16, 0x7a, 0x86, 0xaa, 0x5b, 0xe8, 0x0b, 0x54,
	0xc3, 0x0d, 0x38, 0x58, 0xa7, 0x94, 0x58, 0x84, 0x5e, 0x13, 0x1d, 0xed, 0x35, 0xff, 0x92, 0xe0,
	0xf0, 0xff, 0x3e, 0x16, 0xb8, 0x02, 0xa5, 0xd5, 0xd0, 0xab, 0xa0, 0xa4, 0x99, 0x88, 0x8e, 0x24,
	0x7e, 0xbf, 0x03, 0x6a, 0x6a, 0xc4, 0xb2, 0x8c, 0x7e, 0x17, 0xe5, 0xf0, 0x33, 0xd8, 0xbf, 0x51,
	0x87, 0x5a, 0xcf, 0x1c, 0xf6, 0x08, 0x1d, 0x52, 0xf5, 0xfc, 0xdc, 0xd0, 0x90, 0xcc, 0x53, 0x74,
	0xa9, 0xda, 0xe7, 0x67, 0xf2, 0x62, 0x43, 0xc8, 0x77, 0x44, 0xe3, 0x56, 0x01, 0xef, 0x41, 0xe5,
	0x4a, 0xfd, 0x61, 0xc0, 0x05, 0xf5, 0x35, 0xb2, 0x18, 0x2c, 0x57, 0x67, 0x68, 0xe4, 0xd2, 0xd4,
	0x2e, 0x88, 0x8e, 0x4a, 0x9d, 0x0a, 0x28, 0x9e, 0x7f, 0x17, 0xcc, 0x7d, 0xd7, 0xf4, 0x9b, 0xef,
	0x25, 0xd8, 0xcf, 0xfc, 0xf2, 0xe3, 0xe7, 0xb0, 0xe7, 0xda, 0xb1, 0x7d, 0x1b, 0xfc, 0xe2, 0xb3,
	0xf0, 0xd6, 0x09, 0xdc, 0xe4, 0x9d, 0x2a, 0xb4, 0xca, 0xb1, 0xc9, 0xa9, 0x16, 0xb8, 0x0c, 0x7f,
	0x06, 0x70, 0x37, 0x09, 0x9c, 0xb7, 0x49, 0x48, 0xfa, 0xdb, 0x2e, 0x88, 0x70, 0x1f, 0x43, 0x2d,
	0xfd, 0x0b, 0x71, 0xeb, 0x3f, 0xfc, 0x38, 0x55, 0x53, 0x9a, 0x7c, 0x9b, 0xee, 0x8a, 0xe2, 0x1f,
	0xc3, 0xd9, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14, 0x17, 0x28, 0xa4, 0xf3, 0x08, 0x00, 0x00,
}
