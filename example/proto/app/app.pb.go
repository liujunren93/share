// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/app.proto

package app

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AppInfo struct {
	// @inject_tag: json:"id"
	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// @inject_tag: binding:"required" json:"app_name"
	AppName string `protobuf:"bytes,2,opt,name=appName,proto3" json:"appName,omitempty"`
	// @inject_tag: binding:"required" json:"version"
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// @inject_tag: binding:"required" json:"img"
	Img string `protobuf:"bytes,4,opt,name=img,proto3" json:"img,omitempty"`
	// @inject_tag: binding:"required" json:"icon"
	Icon string `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	// @inject_tag: binding:"required" json:"price"
	Price float64 `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	// @inject_tag: binding:"required,min=0,max=5" json:"unit"
	Unit int32 `protobuf:"varint,7,opt,name=unit,proto3" json:"unit,omitempty"`
	// @inject_tag: binding:"required,min=0,max=5" json:"type"
	Type int32 `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	//@inject_tag: json:"desc"
	Desc string `protobuf:"bytes,9,opt,name=desc,proto3" json:"desc,omitempty"`
	//@inject_tag: json:"content"
	Content string `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty"`
	//@inject_tag: json:"created_at,omitempty"
	CreatedAt string `protobuf:"bytes,11,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	//@inject_tag: json:"updated_at,omitempty"
	UpdatedAt            string   `protobuf:"bytes,12,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppInfo) Reset()         { *m = AppInfo{} }
func (m *AppInfo) String() string { return proto.CompactTextString(m) }
func (*AppInfo) ProtoMessage()    {}
func (*AppInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3850787927b6c91f, []int{0}
}

func (m *AppInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppInfo.Unmarshal(m, b)
}
func (m *AppInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppInfo.Marshal(b, m, deterministic)
}
func (m *AppInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppInfo.Merge(m, src)
}
func (m *AppInfo) XXX_Size() int {
	return xxx_messageInfo_AppInfo.Size(m)
}
func (m *AppInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AppInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AppInfo proto.InternalMessageInfo

func (m *AppInfo) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *AppInfo) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *AppInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AppInfo) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *AppInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *AppInfo) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *AppInfo) GetUnit() int32 {
	if m != nil {
		return m.Unit
	}
	return 0
}

func (m *AppInfo) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *AppInfo) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *AppInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *AppInfo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *AppInfo) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CreateReq struct {
	Data                 *AppInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReq) Reset()         { *m = CreateReq{} }
func (m *CreateReq) String() string { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()    {}
func (*CreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3850787927b6c91f, []int{1}
}

func (m *CreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReq.Unmarshal(m, b)
}
func (m *CreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReq.Marshal(b, m, deterministic)
}
func (m *CreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReq.Merge(m, src)
}
func (m *CreateReq) XXX_Size() int {
	return xxx_messageInfo_CreateReq.Size(m)
}
func (m *CreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReq proto.InternalMessageInfo

func (m *CreateReq) GetData() *AppInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdateReq struct {
	AppID                int32    `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Data                 *AppInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3850787927b6c91f, []int{2}
}

func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (m *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(m, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetAppID() int32 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func (m *UpdateReq) GetData() *AppInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type AppListReq struct {
	// @inject_tag: form:"app_name"
	AppName string `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	// @inject_tag: form:"version"
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	//@inject_tag: json:"page"
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	//@inject_tag: json:"page_size"
	PageSize             int32    `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppListReq) Reset()         { *m = AppListReq{} }
func (m *AppListReq) String() string { return proto.CompactTextString(m) }
func (*AppListReq) ProtoMessage()    {}
func (*AppListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3850787927b6c91f, []int{3}
}

func (m *AppListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppListReq.Unmarshal(m, b)
}
func (m *AppListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppListReq.Marshal(b, m, deterministic)
}
func (m *AppListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppListReq.Merge(m, src)
}
func (m *AppListReq) XXX_Size() int {
	return xxx_messageInfo_AppListReq.Size(m)
}
func (m *AppListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppListReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppListReq proto.InternalMessageInfo

func (m *AppListReq) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *AppListReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AppListReq) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AppListReq) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type AppInfoReq struct {
	//@inject_tag: json:"id"
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppInfoReq) Reset()         { *m = AppInfoReq{} }
func (m *AppInfoReq) String() string { return proto.CompactTextString(m) }
func (*AppInfoReq) ProtoMessage()    {}
func (*AppInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3850787927b6c91f, []int{4}
}

func (m *AppInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppInfoReq.Unmarshal(m, b)
}
func (m *AppInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppInfoReq.Marshal(b, m, deterministic)
}
func (m *AppInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppInfoReq.Merge(m, src)
}
func (m *AppInfoReq) XXX_Size() int {
	return xxx_messageInfo_AppInfoReq.Size(m)
}
func (m *AppInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppInfoReq proto.InternalMessageInfo

func (m *AppInfoReq) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type AppListData struct {
	//@inject_tag: json:"list"
	List []*AppInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	//@inject_tag: json:"total"
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppListData) Reset()         { *m = AppListData{} }
func (m *AppListData) String() string { return proto.CompactTextString(m) }
func (*AppListData) ProtoMessage()    {}
func (*AppListData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3850787927b6c91f, []int{5}
}

func (m *AppListData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppListData.Unmarshal(m, b)
}
func (m *AppListData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppListData.Marshal(b, m, deterministic)
}
func (m *AppListData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppListData.Merge(m, src)
}
func (m *AppListData) XXX_Size() int {
	return xxx_messageInfo_AppListData.Size(m)
}
func (m *AppListData) XXX_DiscardUnknown() {
	xxx_messageInfo_AppListData.DiscardUnknown(m)
}

var xxx_messageInfo_AppListData proto.InternalMessageInfo

func (m *AppListData) GetList() []*AppInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *AppListData) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type AppListRes struct {
	Code                 int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *AppListData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AppListRes) Reset()         { *m = AppListRes{} }
func (m *AppListRes) String() string { return proto.CompactTextString(m) }
func (*AppListRes) ProtoMessage()    {}
func (*AppListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3850787927b6c91f, []int{6}
}

func (m *AppListRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppListRes.Unmarshal(m, b)
}
func (m *AppListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppListRes.Marshal(b, m, deterministic)
}
func (m *AppListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppListRes.Merge(m, src)
}
func (m *AppListRes) XXX_Size() int {
	return xxx_messageInfo_AppListRes.Size(m)
}
func (m *AppListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AppListRes.DiscardUnknown(m)
}

var xxx_messageInfo_AppListRes proto.InternalMessageInfo

func (m *AppListRes) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AppListRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AppListRes) GetData() *AppListData {
	if m != nil {
		return m.Data
	}
	return nil
}

type AppInfoRes struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *AppInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppInfoRes) Reset()         { *m = AppInfoRes{} }
func (m *AppInfoRes) String() string { return proto.CompactTextString(m) }
func (*AppInfoRes) ProtoMessage()    {}
func (*AppInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3850787927b6c91f, []int{7}
}

func (m *AppInfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppInfoRes.Unmarshal(m, b)
}
func (m *AppInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppInfoRes.Marshal(b, m, deterministic)
}
func (m *AppInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppInfoRes.Merge(m, src)
}
func (m *AppInfoRes) XXX_Size() int {
	return xxx_messageInfo_AppInfoRes.Size(m)
}
func (m *AppInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AppInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_AppInfoRes proto.InternalMessageInfo

func (m *AppInfoRes) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AppInfoRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AppInfoRes) GetData() *AppInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteReq struct {
	// @inject_tag: binding:"required" json:"app_id"
	AppID                int32    `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReq) Reset()         { *m = DeleteReq{} }
func (m *DeleteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()    {}
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3850787927b6c91f, []int{8}
}

func (m *DeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReq.Unmarshal(m, b)
}
func (m *DeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReq.Marshal(b, m, deterministic)
}
func (m *DeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReq.Merge(m, src)
}
func (m *DeleteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteReq.Size(m)
}
func (m *DeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReq proto.InternalMessageInfo

func (m *DeleteReq) GetAppID() int32 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func init() {
	proto.RegisterType((*AppInfo)(nil), "app.AppInfo")
	proto.RegisterType((*CreateReq)(nil), "app.CreateReq")
	proto.RegisterType((*UpdateReq)(nil), "app.UpdateReq")
	proto.RegisterType((*AppListReq)(nil), "app.appListReq")
	proto.RegisterType((*AppInfoReq)(nil), "app.appInfoReq")
	proto.RegisterType((*AppListData)(nil), "app.appListData")
	proto.RegisterType((*AppListRes)(nil), "app.appListRes")
	proto.RegisterType((*AppInfoRes)(nil), "app.appInfoRes")
	proto.RegisterType((*DeleteReq)(nil), "app.DeleteReq")
}

func init() { proto.RegisterFile("app/app.proto", fileDescriptor_3850787927b6c91f) }

var fileDescriptor_3850787927b6c91f = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x7f, 0x49, 0x33, 0x29, 0x69, 0x59, 0x21, 0xb1, 0x8a, 0x7a, 0x70, 0xad, 0x1e, 0x9c,
	0x03, 0x89, 0x54, 0x8e, 0x9c, 0x9a, 0x86, 0x43, 0x24, 0xc4, 0xc1, 0x80, 0x84, 0xb8, 0x6d, 0xe3,
	0x6d, 0x64, 0x29, 0xb1, 0xa7, 0xf1, 0xb6, 0x12, 0xbc, 0x1b, 0x2f, 0xc2, 0xa3, 0xf4, 0x84, 0x66,
	0x76, 0xed, 0x98, 0x46, 0xfc, 0xe4, 0x92, 0x99, 0xef, 0x1b, 0xcf, 0xcc, 0x37, 0x33, 0x36, 0x3c,
	0x57, 0x88, 0x33, 0x85, 0x38, 0xc5, 0x5d, 0x65, 0x2a, 0x11, 0x28, 0xc4, 0xf1, 0x88, 0xb0, 0x1b,
	0x55, 0x6b, 0x0b, 0x8e, 0x5f, 0x3d, 0xa8, 0x4d, 0x91, 0x2b, 0xa3, 0x67, 0x8d, 0x61, 0x89, 0xe4,
	0x87, 0x0f, 0xfd, 0x2b, 0xc4, 0x65, 0x79, 0x5b, 0x89, 0x11, 0xf8, 0xcb, 0x85, 0xf4, 0x62, 0x2f,
	0x8d, 0x32, 0x7f, 0xb9, 0x10, 0xe7, 0xd0, 0x57, 0x88, 0x1f, 0xd4, 0x56, 0x4b, 0x3f, 0xf6, 0xd2,
	0xc1, 0xbc, 0xff, 0x38, 0x0f, 0x77, 0x7e, 0xec, 0x65, 0x0d, 0x4e, 0x21, 0x0f, 0x7a, 0x57, 0x17,
	0x55, 0x29, 0x83, 0x4e, 0x48, 0xea, 0x65, 0x0d, 0x2e, 0x4e, 0x21, 0x28, 0xb6, 0x6b, 0x19, 0x12,
	0x9d, 0x91, 0x29, 0x04, 0x84, 0xc5, 0xaa, 0x2a, 0x65, 0xc4, 0x10, 0xdb, 0xe2, 0x02, 0x22, 0xdc,
	0x15, 0x2b, 0x2d, 0x7b, 0xb1, 0x97, 0x7a, 0xf3, 0xd1, 0xe3, 0x7c, 0x28, 0x06, 0x93, 0x67, 0xee,
	0x97, 0x59, 0x92, 0x9e, 0xbc, 0x2f, 0x0b, 0x23, 0xfb, 0xdc, 0x23, 0xdb, 0x84, 0x99, 0x6f, 0xa8,
	0xe5, 0x91, 0xc5, 0xc8, 0x26, 0x2c, 0xd7, 0xf5, 0x4a, 0x0e, 0x6c, 0x05, 0xb2, 0x85, 0x84, 0xfe,
	0xaa, 0x2a, 0x8d, 0x2e, 0x8d, 0x04, 0x86, 0x1b, 0x57, 0x9c, 0xc1, 0xe0, 0x7a, 0xa7, 0x95, 0xd1,
	0xf9, 0x95, 0x91, 0x43, 0xe6, 0xf6, 0x00, 0xb1, 0x9f, 0x31, 0x77, 0xec, 0xb1, 0x65, 0x5b, 0x20,
	0x79, 0xdd, 0x3c, 0x9b, 0xe9, 0x3b, 0x11, 0x43, 0x98, 0x2b, 0xa3, 0x78, 0x84, 0xc3, 0xcb, 0xe3,
	0x29, 0x2d, 0xc5, 0x0d, 0x37, 0x63, 0x26, 0xb9, 0x6e, 0x92, 0x51, 0xf8, 0x4b, 0x88, 0x14, 0x62,
	0x3b, 0x72, 0xeb, 0xb4, 0x49, 0x82, 0x3f, 0x26, 0x41, 0x00, 0x85, 0xf8, 0xbe, 0xa8, 0x0d, 0x65,
	0x91, 0xfb, 0x2d, 0x79, 0x56, 0x57, 0xb3, 0x1c, 0xb9, 0x5f, 0x8e, 0x6f, 0x99, 0x66, 0x27, 0x02,
	0x42, 0x54, 0x6b, 0xcd, 0x35, 0xa2, 0x8c, 0x6d, 0x31, 0x86, 0x23, 0xfa, 0xff, 0x58, 0x7c, 0xd7,
	0xbc, 0xac, 0x28, 0x6b, 0xfd, 0xe4, 0x8c, 0x2b, 0x72, 0x0b, 0xfa, 0xee, 0xe9, 0x9d, 0x24, 0xef,
	0x60, 0xe8, 0xfa, 0x59, 0x28, 0xa3, 0x48, 0xc0, 0xa6, 0xa8, 0x8d, 0xf4, 0xe2, 0xe0, 0x50, 0x00,
	0x31, 0x24, 0xdc, 0x54, 0x46, 0x6d, 0xb8, 0xad, 0x20, 0xb3, 0x4e, 0xf2, 0xa5, 0x23, 0xab, 0xa6,
	0x16, 0x57, 0x55, 0xae, 0x5d, 0x19, 0xb6, 0xe9, 0x94, 0xb6, 0xf5, 0xda, 0x89, 0x21, 0x53, 0x5c,
	0xfc, 0x36, 0xac, 0x53, 0xae, 0xd5, 0xe9, 0xc5, 0x0d, 0xec, 0x53, 0xa7, 0xfd, 0xff, 0xcd, 0xfc,
	0xef, 0x35, 0x9c, 0xc3, 0x60, 0xa1, 0x37, 0xfa, 0x2f, 0xbb, 0xbc, 0xfc, 0xe9, 0x01, 0xbd, 0x8e,
	0x62, 0x02, 0x3d, 0x7b, 0x25, 0x62, 0xc4, 0x89, 0xda, 0x93, 0x19, 0x9f, 0xb0, 0x9f, 0xeb, 0x5b,
	0x75, 0xbf, 0x61, 0xdd, 0x13, 0xe8, 0xd9, 0x0b, 0x71, 0xa1, 0xed, 0xb9, 0x1c, 0x86, 0xa6, 0x10,
	0x92, 0x50, 0x71, 0xd2, 0x95, 0xbd, 0x8f, 0xec, 0x0c, 0x33, 0x85, 0x90, 0xdf, 0xf0, 0x96, 0x70,
	0xab, 0x1c, 0x3f, 0x01, 0xb8, 0xbc, 0x15, 0xe5, 0xca, 0xb7, 0x0a, 0x0f, 0xca, 0xcf, 0x5f, 0x7c,
	0x3d, 0x99, 0xce, 0xf8, 0x2b, 0x42, 0x5f, 0x9f, 0xb7, 0x0a, 0xf1, 0xa6, 0xc7, 0xee, 0x9b, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0xb1, 0x17, 0xab, 0x93, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*DefaultRes, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*DefaultRes, error)
	List(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*AppListRes, error)
	Info(ctx context.Context, in *AppInfoReq, opts ...grpc.CallOption) (*AppInfoRes, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DefaultRes, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*DefaultRes, error) {
	out := new(DefaultRes)
	err := c.cc.Invoke(ctx, "/app.app/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*DefaultRes, error) {
	out := new(DefaultRes)
	err := c.cc.Invoke(ctx, "/app.app/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) List(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*AppListRes, error) {
	out := new(AppListRes)
	err := c.cc.Invoke(ctx, "/app.app/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Info(ctx context.Context, in *AppInfoReq, opts ...grpc.CallOption) (*AppInfoRes, error) {
	out := new(AppInfoRes)
	err := c.cc.Invoke(ctx, "/app.app/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DefaultRes, error) {
	out := new(DefaultRes)
	err := c.cc.Invoke(ctx, "/app.app/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
type AppServer interface {
	Create(context.Context, *CreateReq) (*DefaultRes, error)
	Update(context.Context, *UpdateReq) (*DefaultRes, error)
	List(context.Context, *AppListReq) (*AppListRes, error)
	Info(context.Context, *AppInfoReq) (*AppInfoRes, error)
	Delete(context.Context, *DeleteReq) (*DefaultRes, error)
}

// UnimplementedAppServer can be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (*UnimplementedAppServer) Create(ctx context.Context, req *CreateReq) (*DefaultRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAppServer) Update(ctx context.Context, req *UpdateReq) (*DefaultRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAppServer) List(ctx context.Context, req *AppListReq) (*AppListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAppServer) Info(ctx context.Context, req *AppInfoReq) (*AppInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedAppServer) Delete(ctx context.Context, req *DeleteReq) (*DefaultRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.app/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.app/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.app/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).List(ctx, req.(*AppListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.app/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Info(ctx, req.(*AppInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.app/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.app",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _App_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _App_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _App_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _App_Info_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _App_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/app.proto",
}
