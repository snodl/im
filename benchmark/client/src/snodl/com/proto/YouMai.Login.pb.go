// Code generated by protoc-gen-go.
// source: YouMai.Login.proto
// DO NOT EDIT!

package youmai

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User_Login struct {
	UserId           *string      `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	AppId            *string      `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	DeviceType       *Device_Type `protobuf:"varint,3,opt,name=device_type,json=deviceType,enum=youmai.Device_Type" json:"device_type,omitempty"`
	DeviceId         *string      `protobuf:"bytes,4,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Version          *int32       `protobuf:"varint,8,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *User_Login) Reset()                    { *m = User_Login{} }
func (m *User_Login) String() string            { return proto.CompactTextString(m) }
func (*User_Login) ProtoMessage()               {}
func (*User_Login) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *User_Login) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *User_Login) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

func (m *User_Login) GetDeviceType() Device_Type {
	if m != nil && m.DeviceType != nil {
		return *m.DeviceType
	}
	return Device_Type_DeviceType_IPhone
}

func (m *User_Login) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

func (m *User_Login) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type User_Login_Ack struct {
	ErrerNo          *ERRNO_CODE `protobuf:"varint,1,opt,name=errer_no,json=errerNo,enum=youmai.ERRNO_CODE" json:"errer_no,omitempty"`
	Error            *string     `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *User_Login_Ack) Reset()                    { *m = User_Login_Ack{} }
func (m *User_Login_Ack) String() string            { return proto.CompactTextString(m) }
func (*User_Login_Ack) ProtoMessage()               {}
func (*User_Login_Ack) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *User_Login_Ack) GetErrerNo() ERRNO_CODE {
	if m != nil && m.ErrerNo != nil {
		return *m.ErrerNo
	}
	return ERRNO_CODE_ERRNO_CODE_OK
}

func (m *User_Login_Ack) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

type Multi_Device_Kicked_Notify struct {
	UserId           *string      `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Timestamp        *int32       `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	NewDevice        *Device_Type `protobuf:"varint,3,opt,name=new_device,json=newDevice,enum=youmai.Device_Type" json:"new_device,omitempty"`
	NewDeviceId      *string      `protobuf:"bytes,4,opt,name=new_device_id,json=newDeviceId" json:"new_device_id,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Multi_Device_Kicked_Notify) Reset()                    { *m = Multi_Device_Kicked_Notify{} }
func (m *Multi_Device_Kicked_Notify) String() string            { return proto.CompactTextString(m) }
func (*Multi_Device_Kicked_Notify) ProtoMessage()               {}
func (*Multi_Device_Kicked_Notify) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *Multi_Device_Kicked_Notify) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *Multi_Device_Kicked_Notify) GetTimestamp() int32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Multi_Device_Kicked_Notify) GetNewDevice() Device_Type {
	if m != nil && m.NewDevice != nil {
		return *m.NewDevice
	}
	return Device_Type_DeviceType_IPhone
}

func (m *Multi_Device_Kicked_Notify) GetNewDeviceId() string {
	if m != nil && m.NewDeviceId != nil {
		return *m.NewDeviceId
	}
	return ""
}

type User_LogOff struct {
	UserId           *string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	DeviceId         *string `protobuf:"bytes,2,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *User_LogOff) Reset()                    { *m = User_LogOff{} }
func (m *User_LogOff) String() string            { return proto.CompactTextString(m) }
func (*User_LogOff) ProtoMessage()               {}
func (*User_LogOff) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *User_LogOff) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *User_LogOff) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

type User_LogOff_Ack struct {
	ErrerNo          *ERRNO_CODE `protobuf:"varint,1,opt,name=errer_no,json=errerNo,enum=youmai.ERRNO_CODE" json:"errer_no,omitempty"`
	Error            *string     `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *User_LogOff_Ack) Reset()                    { *m = User_LogOff_Ack{} }
func (m *User_LogOff_Ack) String() string            { return proto.CompactTextString(m) }
func (*User_LogOff_Ack) ProtoMessage()               {}
func (*User_LogOff_Ack) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *User_LogOff_Ack) GetErrerNo() ERRNO_CODE {
	if m != nil && m.ErrerNo != nil {
		return *m.ErrerNo
	}
	return ERRNO_CODE_ERRNO_CODE_OK
}

func (m *User_LogOff_Ack) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*User_Login)(nil), "youmai.User_Login")
	proto.RegisterType((*User_Login_Ack)(nil), "youmai.User_Login_Ack")
	proto.RegisterType((*Multi_Device_Kicked_Notify)(nil), "youmai.Multi_Device_Kicked_Notify")
	proto.RegisterType((*User_LogOff)(nil), "youmai.User_LogOff")
	proto.RegisterType((*User_LogOff_Ack)(nil), "youmai.User_LogOff_Ack")
}

func init() { proto.RegisterFile("YouMai.Login.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0x53, 0xee, 0x2d, 0xd0, 0x43, 0x2e, 0x37, 0x19, 0x35, 0x36, 0xe8, 0x82, 0x74, 0xc5,
	0xc6, 0x2e, 0x88, 0x2f, 0xa0, 0xc0, 0xa2, 0x51, 0x20, 0x99, 0x88, 0x89, 0xab, 0x49, 0xd3, 0x39,
	0x98, 0x09, 0xd2, 0x99, 0x4c, 0x5b, 0x48, 0x1f, 0xc9, 0xf8, 0x92, 0xa6, 0x33, 0xd4, 0xca, 0x02,
	0x57, 0xee, 0xe6, 0xfb, 0xce, 0x9f, 0xf9, 0xce, 0x0f, 0xc8, 0x8b, 0x2c, 0xe6, 0xb1, 0x08, 0x1f,
	0xe5, 0xab, 0x48, 0x43, 0xa5, 0x65, 0x2e, 0x49, 0xbb, 0x94, 0xc5, 0x36, 0x16, 0x83, 0xba, 0x76,
	0x1f, 0x67, 0x22, 0xb1, 0xb5, 0xe0, 0xdd, 0x01, 0x58, 0x65, 0xa8, 0x99, 0x19, 0x20, 0x97, 0xd0,
	0x29, 0x2a, 0x25, 0xb8, 0xef, 0x0c, 0x9d, 0x91, 0x47, 0xdb, 0x95, 0x8c, 0x38, 0xb9, 0x80, 0x76,
	0xac, 0x54, 0xe5, 0xb7, 0x8c, 0xef, 0xc6, 0x4a, 0x45, 0x9c, 0xdc, 0x42, 0x8f, 0xe3, 0x4e, 0x24,
	0xc8, 0xf2, 0x52, 0xa1, 0xff, 0x67, 0xe8, 0x8c, 0xfa, 0xe3, 0xb3, 0xd0, 0x7e, 0x18, 0x4e, 0x6d,
	0xe9, 0xa9, 0x54, 0x48, 0xc1, 0xf6, 0x55, 0x6f, 0x72, 0x05, 0xde, 0x61, 0x4a, 0x70, 0xff, 0xaf,
	0xd9, 0xd7, 0xb5, 0x46, 0xc4, 0x89, 0x0f, 0x9d, 0x1d, 0xea, 0x4c, 0xc8, 0xd4, 0xef, 0x0e, 0x9d,
	0x91, 0x4b, 0x6b, 0x19, 0xac, 0xa0, 0xdf, 0x44, 0x65, 0x77, 0xc9, 0x86, 0xdc, 0x40, 0x17, 0xb5,
	0x46, 0xcd, 0x52, 0x69, 0xf2, 0xf6, 0xc7, 0xa4, 0xfe, 0x7b, 0x46, 0xe9, 0x62, 0xc9, 0x26, 0xcb,
	0xe9, 0x8c, 0x76, 0x4c, 0xcf, 0x42, 0x92, 0x73, 0x70, 0x51, 0x6b, 0xa9, 0xeb, 0x1b, 0x8c, 0x08,
	0x3e, 0x1c, 0x18, 0xcc, 0x8b, 0xb7, 0x5c, 0xb0, 0x43, 0xde, 0x07, 0x91, 0x6c, 0x90, 0xb3, 0x85,
	0xcc, 0xc5, 0xba, 0x3c, 0x8d, 0xe4, 0x1a, 0xbc, 0x5c, 0x6c, 0x31, 0xcb, 0xe3, 0xad, 0x32, 0x1b,
	0x5d, 0xda, 0x18, 0x64, 0x0c, 0x90, 0xe2, 0x9e, 0xd9, 0xb3, 0x7e, 0x02, 0xe3, 0xa5, 0xb8, 0xb7,
	0x9a, 0x04, 0xf0, 0xaf, 0x99, 0x69, 0xd8, 0xf4, 0xbe, 0x3a, 0x22, 0x1e, 0x4c, 0xa0, 0x57, 0x43,
	0x58, 0xae, 0xd7, 0xa7, 0xd3, 0x1d, 0x31, 0x6e, 0x1d, 0x33, 0x0e, 0x9e, 0xe1, 0xff, 0xb7, 0x25,
	0xbf, 0x86, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x07, 0xa1, 0x01, 0xde, 0x7e, 0x02, 0x00, 0x00,
}