// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package purchaseproto

import (
	context "context"
	fmt "fmt"
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

type RequestHeader struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Product struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity             int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type RequestField struct {
	Vendor               int32      `protobuf:"varint,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	VendorBillNumber     string     `protobuf:"bytes,2,opt,name=vendorBillNumber,proto3" json:"vendorBillNumber,omitempty"`
	VendorBillDate       string     `protobuf:"bytes,3,opt,name=vendorBillDate,proto3" json:"vendorBillDate,omitempty"`
	Name                 string     `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string     `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	PaidAmount           float64    `protobuf:"fixed64,6,opt,name=paidAmount,proto3" json:"paidAmount,omitempty"`
	Products             []*Product `protobuf:"bytes,7,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RequestField) Reset()         { *m = RequestField{} }
func (m *RequestField) String() string { return proto.CompactTextString(m) }
func (*RequestField) ProtoMessage()    {}
func (*RequestField) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *RequestField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestField.Unmarshal(m, b)
}
func (m *RequestField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestField.Marshal(b, m, deterministic)
}
func (m *RequestField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestField.Merge(m, src)
}
func (m *RequestField) XXX_Size() int {
	return xxx_messageInfo_RequestField.Size(m)
}
func (m *RequestField) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestField.DiscardUnknown(m)
}

var xxx_messageInfo_RequestField proto.InternalMessageInfo

func (m *RequestField) GetVendor() int32 {
	if m != nil {
		return m.Vendor
	}
	return 0
}

func (m *RequestField) GetVendorBillNumber() string {
	if m != nil {
		return m.VendorBillNumber
	}
	return ""
}

func (m *RequestField) GetVendorBillDate() string {
	if m != nil {
		return m.VendorBillDate
	}
	return ""
}

func (m *RequestField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestField) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RequestField) GetPaidAmount() float64 {
	if m != nil {
		return m.PaidAmount
	}
	return 0
}

func (m *RequestField) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type RawRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RawRequest) Reset()         { *m = RawRequest{} }
func (m *RawRequest) String() string { return proto.CompactTextString(m) }
func (*RawRequest) ProtoMessage()    {}
func (*RawRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *RawRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawRequest.Unmarshal(m, b)
}
func (m *RawRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawRequest.Marshal(b, m, deterministic)
}
func (m *RawRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawRequest.Merge(m, src)
}
func (m *RawRequest) XXX_Size() int {
	return xxx_messageInfo_RawRequest.Size(m)
}
func (m *RawRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RawRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RawRequest proto.InternalMessageInfo

func (m *RawRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type GetRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Field                *RequestField  `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CreateRequest) GetField() *RequestField {
	if m != nil {
		return m.Field
	}
	return nil
}

type UpdateRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Field                *RequestField  `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UpdateRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetField() *RequestField {
	if m != nil {
		return m.Field
	}
	return nil
}

type ActionRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRequest.Unmarshal(m, b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
}
func (m *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(m, src)
}
func (m *ActionRequest) XXX_Size() int {
	return xxx_messageInfo_ActionRequest.Size(m)
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ActionRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DataResponse struct {
	StatusCode           int32              `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *DataResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataResponse) Reset()         { *m = DataResponse{} }
func (m *DataResponse) String() string { return proto.CompactTextString(m) }
func (*DataResponse) ProtoMessage()    {}
func (*DataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *DataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse.Unmarshal(m, b)
}
func (m *DataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse.Marshal(b, m, deterministic)
}
func (m *DataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse.Merge(m, src)
}
func (m *DataResponse) XXX_Size() int {
	return xxx_messageInfo_DataResponse.Size(m)
}
func (m *DataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse proto.InternalMessageInfo

func (m *DataResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DataResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DataResponse) GetData() *DataResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DataResponse_Data struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	VendorID             int32             `protobuf:"varint,4,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
	Vendor               string            `protobuf:"bytes,5,opt,name=vendor,proto3" json:"vendor,omitempty"`
	VendorBillNumber     string            `protobuf:"bytes,6,opt,name=vendorBillNumber,proto3" json:"vendorBillNumber,omitempty"`
	VendorBillDate       string            `protobuf:"bytes,7,opt,name=vendorBillDate,proto3" json:"vendorBillDate,omitempty"`
	Products             []int32           `protobuf:"varint,8,rep,packed,name=products,proto3" json:"products,omitempty"`
	TotalNetAmount       float64           `protobuf:"fixed64,9,opt,name=totalNetAmount,proto3" json:"totalNetAmount,omitempty"`
	TotalLevyAmount      float64           `protobuf:"fixed64,10,opt,name=totalLevyAmount,proto3" json:"totalLevyAmount,omitempty"`
	TotalFinalAmount     float64           `protobuf:"fixed64,11,opt,name=totalFinalAmount,proto3" json:"totalFinalAmount,omitempty"`
	GrandTotal           float64           `protobuf:"fixed64,12,opt,name=grandTotal,proto3" json:"grandTotal,omitempty"`
	TotalPaidAmount      float64           `protobuf:"fixed64,13,opt,name=totalPaidAmount,proto3" json:"totalPaidAmount,omitempty"`
	CreatedDate          string            `protobuf:"bytes,14,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	CreatedBy            string            `protobuf:"bytes,15,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	LastModifiedBy       string            `protobuf:"bytes,16,opt,name=lastModifiedBy,proto3" json:"lastModifiedBy,omitempty"`
	LastModifiedDate     string            `protobuf:"bytes,17,opt,name=lastModifiedDate,proto3" json:"lastModifiedDate,omitempty"`
	RecordState          string            `protobuf:"bytes,18,opt,name=recordState,proto3" json:"recordState,omitempty"`
	ReferencedBy         int64             `protobuf:"varint,19,opt,name=referencedBy,proto3" json:"referencedBy,omitempty"`
	Currency             string            `protobuf:"bytes,20,opt,name=currency,proto3" json:"currency,omitempty"`
	Quantity             map[string]string `protobuf:"bytes,21,rep,name=quantity,proto3" json:"quantity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DataResponse_Data) Reset()         { *m = DataResponse_Data{} }
func (m *DataResponse_Data) String() string { return proto.CompactTextString(m) }
func (*DataResponse_Data) ProtoMessage()    {}
func (*DataResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8, 0}
}

func (m *DataResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse_Data.Unmarshal(m, b)
}
func (m *DataResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse_Data.Marshal(b, m, deterministic)
}
func (m *DataResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse_Data.Merge(m, src)
}
func (m *DataResponse_Data) XXX_Size() int {
	return xxx_messageInfo_DataResponse_Data.Size(m)
}
func (m *DataResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse_Data proto.InternalMessageInfo

func (m *DataResponse_Data) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataResponse_Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataResponse_Data) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DataResponse_Data) GetVendorID() int32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *DataResponse_Data) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *DataResponse_Data) GetVendorBillNumber() string {
	if m != nil {
		return m.VendorBillNumber
	}
	return ""
}

func (m *DataResponse_Data) GetVendorBillDate() string {
	if m != nil {
		return m.VendorBillDate
	}
	return ""
}

func (m *DataResponse_Data) GetProducts() []int32 {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *DataResponse_Data) GetTotalNetAmount() float64 {
	if m != nil {
		return m.TotalNetAmount
	}
	return 0
}

func (m *DataResponse_Data) GetTotalLevyAmount() float64 {
	if m != nil {
		return m.TotalLevyAmount
	}
	return 0
}

func (m *DataResponse_Data) GetTotalFinalAmount() float64 {
	if m != nil {
		return m.TotalFinalAmount
	}
	return 0
}

func (m *DataResponse_Data) GetGrandTotal() float64 {
	if m != nil {
		return m.GrandTotal
	}
	return 0
}

func (m *DataResponse_Data) GetTotalPaidAmount() float64 {
	if m != nil {
		return m.TotalPaidAmount
	}
	return 0
}

func (m *DataResponse_Data) GetCreatedDate() string {
	if m != nil {
		return m.CreatedDate
	}
	return ""
}

func (m *DataResponse_Data) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *DataResponse_Data) GetLastModifiedBy() string {
	if m != nil {
		return m.LastModifiedBy
	}
	return ""
}

func (m *DataResponse_Data) GetLastModifiedDate() string {
	if m != nil {
		return m.LastModifiedDate
	}
	return ""
}

func (m *DataResponse_Data) GetRecordState() string {
	if m != nil {
		return m.RecordState
	}
	return ""
}

func (m *DataResponse_Data) GetReferencedBy() int64 {
	if m != nil {
		return m.ReferencedBy
	}
	return 0
}

func (m *DataResponse_Data) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *DataResponse_Data) GetQuantity() map[string]string {
	if m != nil {
		return m.Quantity
	}
	return nil
}

type RawResponse struct {
	StatusCode           int32             `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *RawResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RawResponse) Reset()         { *m = RawResponse{} }
func (m *RawResponse) String() string { return proto.CompactTextString(m) }
func (*RawResponse) ProtoMessage()    {}
func (*RawResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *RawResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawResponse.Unmarshal(m, b)
}
func (m *RawResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawResponse.Marshal(b, m, deterministic)
}
func (m *RawResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawResponse.Merge(m, src)
}
func (m *RawResponse) XXX_Size() int {
	return xxx_messageInfo_RawResponse.Size(m)
}
func (m *RawResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RawResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RawResponse proto.InternalMessageInfo

func (m *RawResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *RawResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RawResponse) GetData() *RawResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type RawResponse_Data struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawResponse_Data) Reset()         { *m = RawResponse_Data{} }
func (m *RawResponse_Data) String() string { return proto.CompactTextString(m) }
func (*RawResponse_Data) ProtoMessage()    {}
func (*RawResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9, 0}
}

func (m *RawResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawResponse_Data.Unmarshal(m, b)
}
func (m *RawResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawResponse_Data.Marshal(b, m, deterministic)
}
func (m *RawResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawResponse_Data.Merge(m, src)
}
func (m *RawResponse_Data) XXX_Size() int {
	return xxx_messageInfo_RawResponse_Data.Size(m)
}
func (m *RawResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_RawResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_RawResponse_Data proto.InternalMessageInfo

func (m *RawResponse_Data) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "purchaseproto.RequestHeader")
	proto.RegisterType((*Product)(nil), "purchaseproto.Product")
	proto.RegisterType((*RequestField)(nil), "purchaseproto.RequestField")
	proto.RegisterType((*RawRequest)(nil), "purchaseproto.RawRequest")
	proto.RegisterType((*GetRequest)(nil), "purchaseproto.GetRequest")
	proto.RegisterType((*CreateRequest)(nil), "purchaseproto.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "purchaseproto.UpdateRequest")
	proto.RegisterType((*ActionRequest)(nil), "purchaseproto.ActionRequest")
	proto.RegisterType((*DataResponse)(nil), "purchaseproto.DataResponse")
	proto.RegisterType((*DataResponse_Data)(nil), "purchaseproto.DataResponse.Data")
	proto.RegisterMapType((map[string]string)(nil), "purchaseproto.DataResponse.Data.QuantityEntry")
	proto.RegisterType((*RawResponse)(nil), "purchaseproto.RawResponse")
	proto.RegisterType((*RawResponse_Data)(nil), "purchaseproto.RawResponse.Data")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0xed, 0xc4, 0x49, 0x26, 0x71, 0xdb, 0x7f, 0xff, 0x52, 0x19, 0xb7, 0x82, 0xc8, 0x12,
	0x28, 0xe2, 0x22, 0x12, 0x69, 0x91, 0x10, 0x87, 0x8b, 0x36, 0xa5, 0x05, 0x04, 0x55, 0x31, 0xf0,
	0x00, 0x5b, 0x7b, 0xdb, 0x5a, 0x4d, 0xec, 0x74, 0xbd, 0x4e, 0xc9, 0x3d, 0x17, 0x3c, 0x06, 0x6f,
	0x01, 0x4f, 0xc2, 0xf3, 0xa0, 0x3d, 0x24, 0x3e, 0xa4, 0x4d, 0x2b, 0x85, 0x3b, 0xcf, 0xb7, 0xe3,
	0x6f, 0x66, 0x67, 0xf6, 0x9b, 0x01, 0x2b, 0x21, 0x74, 0x1c, 0xfa, 0xa4, 0x3b, 0xa2, 0x31, 0x8b,
	0x91, 0x35, 0x4a, 0xa9, 0x7f, 0x8e, 0x13, 0x22, 0x4c, 0xf7, 0x11, 0x58, 0x1e, 0xb9, 0x4c, 0x49,
	0xc2, 0xde, 0x12, 0x1c, 0x10, 0x8a, 0xd6, 0xa1, 0xca, 0xe2, 0x0b, 0x12, 0xd9, 0x5a, 0x5b, 0xeb,
	0x34, 0x3c, 0x69, 0xb8, 0xcf, 0xa0, 0x76, 0x4c, 0xe3, 0x20, 0xf5, 0x19, 0x5a, 0x01, 0x3d, 0x0c,
	0xc4, 0x69, 0xd5, 0xd3, 0xc3, 0x00, 0x39, 0x50, 0xbf, 0x4c, 0x71, 0xc4, 0x42, 0x36, 0xb1, 0x75,
	0x81, 0xce, 0x6c, 0xf7, 0xbb, 0x0e, 0x2d, 0x45, 0x7f, 0x10, 0x92, 0x41, 0x80, 0x36, 0xc0, 0x1c,
	0x93, 0x28, 0x88, 0xa9, 0x22, 0x50, 0x16, 0x7a, 0x02, 0x6b, 0xf2, 0x6b, 0x2f, 0x1c, 0x0c, 0x8e,
	0xd2, 0xe1, 0x09, 0xa1, 0x82, 0xac, 0xe1, 0xcd, 0xe1, 0xe8, 0x31, 0xac, 0x64, 0xd8, 0x3e, 0x66,
	0xc4, 0x36, 0x84, 0x67, 0x09, 0x45, 0x08, 0x2a, 0x11, 0x1e, 0x12, 0xbb, 0x22, 0x4e, 0xc5, 0x37,
	0x6a, 0x43, 0x33, 0x20, 0x89, 0x4f, 0xc3, 0x11, 0x0b, 0xe3, 0xc8, 0xae, 0x8a, 0xa3, 0x3c, 0x84,
	0x1e, 0x00, 0x8c, 0x70, 0x18, 0xec, 0x0e, 0xe3, 0x34, 0x62, 0xb6, 0xd9, 0xd6, 0x3a, 0x9a, 0x97,
	0x43, 0x50, 0x0f, 0xea, 0x23, 0x59, 0x89, 0xc4, 0xae, 0xb5, 0x8d, 0x4e, 0xb3, 0xb7, 0xd1, 0x2d,
	0x94, 0xb4, 0xab, 0x0a, 0xe5, 0xcd, 0xfc, 0xdc, 0x3d, 0x00, 0x0f, 0x5f, 0xa9, 0x42, 0xa0, 0x1d,
	0x30, 0xcf, 0x45, 0xad, 0x45, 0x0d, 0x9a, 0xbd, 0xad, 0xd2, 0xff, 0x85, 0x7e, 0x78, 0xca, 0xd7,
	0xf5, 0x00, 0x0e, 0x09, 0x5b, 0x8a, 0x43, 0xb5, 0x4e, 0x9f, 0xb6, 0xce, 0xfd, 0x06, 0x56, 0x9f,
	0x12, 0xcc, 0xc8, 0x72, 0xb4, 0x4f, 0xa1, 0x7a, 0xca, 0xbb, 0x2b, 0x98, 0x9b, 0xbd, 0xcd, 0xeb,
	0x7f, 0x12, 0x0f, 0xc0, 0x93, 0x9e, 0xee, 0x0f, 0x0d, 0xac, 0xaf, 0xa3, 0x60, 0xe9, 0xd0, 0xa5,
	0x1b, 0x65, 0xa9, 0x18, 0x77, 0x4e, 0x25, 0x04, 0x6b, 0xd7, 0xe7, 0xad, 0xff, 0xb7, 0x99, 0x4c,
	0x5f, 0x9f, 0x91, 0xbd, 0x3e, 0xf7, 0x77, 0x0d, 0x5a, 0xfb, 0x98, 0x61, 0x8f, 0x24, 0xa3, 0x38,
	0x4a, 0x08, 0x7f, 0x6c, 0x09, 0xc3, 0x2c, 0x4d, 0xfa, 0x71, 0x40, 0x94, 0x24, 0x72, 0x08, 0xb2,
	0xa1, 0x36, 0x24, 0x49, 0x82, 0xcf, 0x88, 0x52, 0xc3, 0xd4, 0x44, 0x3b, 0x50, 0x09, 0x30, 0xc3,
	0xea, 0x9e, 0xed, 0x52, 0x8a, 0xf9, 0x20, 0xd2, 0x10, 0xde, 0xce, 0x2f, 0x13, 0x2a, 0xdc, 0xcc,
	0x89, 0xb8, 0x51, 0xc8, 0x56, 0xbf, 0x59, 0x2b, 0xc6, 0xbc, 0x56, 0x1c, 0xa8, 0x4b, 0xcd, 0xbd,
	0xdb, 0x17, 0x2a, 0xab, 0x7a, 0x33, 0x3b, 0xa7, 0x74, 0x29, 0xb2, 0x45, 0x4a, 0x37, 0xef, 0xac,
	0xf4, 0xda, 0xb5, 0x4a, 0x77, 0x72, 0x9a, 0xac, 0xb7, 0x0d, 0x9e, 0xc7, 0xd4, 0xe6, 0x1c, 0x2c,
	0x66, 0x78, 0x70, 0x44, 0x98, 0xd2, 0x74, 0x43, 0x68, 0xba, 0x84, 0xa2, 0x0e, 0xac, 0x0a, 0xe4,
	0x03, 0x19, 0x4f, 0x94, 0x23, 0x08, 0xc7, 0x32, 0xcc, 0x6f, 0x20, 0xa0, 0x83, 0x30, 0xc2, 0x03,
	0xe5, 0xda, 0x14, 0xae, 0x73, 0x38, 0x6f, 0xf0, 0x19, 0xc5, 0x51, 0xf0, 0x85, 0x1f, 0xd8, 0x2d,
	0x39, 0x4d, 0x32, 0x64, 0x16, 0xf5, 0x38, 0x1b, 0x39, 0x56, 0x2e, 0x6a, 0x06, 0xf3, 0x6e, 0xf8,
	0x42, 0xab, 0x81, 0x28, 0xc4, 0x8a, 0xec, 0x46, 0x0e, 0x42, 0x5b, 0xd0, 0x50, 0xe6, 0xde, 0xc4,
	0x5e, 0x15, 0xe7, 0x19, 0xc0, 0xeb, 0x30, 0xc0, 0x09, 0xfb, 0x18, 0x07, 0xe1, 0x69, 0x28, 0x5c,
	0xd6, 0x64, 0x2d, 0x8b, 0x28, 0xbf, 0x5d, 0x1e, 0x11, 0xc1, 0xfe, 0x93, 0xfd, 0x29, 0xe3, 0x3c,
	0x27, 0x4a, 0xfc, 0x98, 0x06, 0x9f, 0x19, 0x77, 0x43, 0x32, 0xa7, 0x1c, 0x84, 0x5c, 0x68, 0x51,
	0x72, 0x4a, 0x28, 0x89, 0x7c, 0x11, 0xf3, 0xff, 0xb6, 0xd6, 0x31, 0xbc, 0x02, 0xc6, 0xbb, 0xe7,
	0xa7, 0x94, 0x9b, 0x13, 0x7b, 0x5d, 0x50, 0xcc, 0x6c, 0xf4, 0x3e, 0xb7, 0x5c, 0xee, 0x89, 0x69,
	0xdb, 0xbd, 0xed, 0xa9, 0x77, 0x3f, 0xa9, 0x1f, 0xde, 0x44, 0x8c, 0x4e, 0xb2, 0x65, 0xe4, 0xbc,
	0x04, 0xab, 0x70, 0x84, 0xd6, 0xc0, 0xb8, 0x20, 0x13, 0xa5, 0x02, 0xfe, 0xc9, 0x97, 0xdf, 0x18,
	0x0f, 0xd2, 0xa9, 0x0e, 0xa4, 0xf1, 0x42, 0x7f, 0xae, 0xb9, 0x3f, 0x35, 0x68, 0x8a, 0x19, 0xbe,
	0xb4, 0x72, 0xb7, 0x0b, 0xca, 0x7d, 0x58, 0x1e, 0x2e, 0x59, 0x8c, 0xbc, 0x70, 0x5d, 0xa5, 0xdb,
	0x7c, 0xad, 0xb4, 0x62, 0xad, 0x7a, 0x7f, 0x74, 0xa8, 0x1f, 0x2b, 0x32, 0xf4, 0x0a, 0x0c, 0x0f,
	0x5f, 0xa1, 0xfb, 0xd7, 0xd1, 0x8b, 0xf1, 0xe5, 0x38, 0x37, 0x47, 0x46, 0x7d, 0x30, 0xe5, 0x62,
	0x40, 0xe5, 0xe1, 0x57, 0xd8, 0x17, 0xce, 0xe6, 0x82, 0x66, 0xa0, 0xd7, 0x60, 0x1c, 0x12, 0x36,
	0x97, 0x42, 0xb6, 0xc5, 0x16, 0xff, 0xde, 0x07, 0x53, 0x6e, 0x88, 0xb9, 0x1c, 0x0a, 0x8b, 0xe3,
	0x56, 0x12, 0x39, 0xdc, 0xe7, 0x48, 0x0a, 0x33, 0x7f, 0x21, 0xc9, 0x89, 0x29, 0xb0, 0xed, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0x60, 0xcd, 0xdb, 0x4a, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PurchaseClient is the client API for Purchase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PurchaseClient interface {
	Raw(ctx context.Context, in *RawRequest, opts ...grpc.CallOption) (*RawResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*DataResponse, error)
}

type purchaseClient struct {
	cc *grpc.ClientConn
}

func NewPurchaseClient(cc *grpc.ClientConn) PurchaseClient {
	return &purchaseClient{cc}
}

func (c *purchaseClient) Raw(ctx context.Context, in *RawRequest, opts ...grpc.CallOption) (*RawResponse, error) {
	out := new(RawResponse)
	err := c.cc.Invoke(ctx, "/purchaseproto.Purchase/Raw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/purchaseproto.Purchase/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/purchaseproto.Purchase/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/purchaseproto.Purchase/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseClient) Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/purchaseproto.Purchase/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PurchaseServer is the server API for Purchase service.
type PurchaseServer interface {
	Raw(context.Context, *RawRequest) (*RawResponse, error)
	Create(context.Context, *CreateRequest) (*DataResponse, error)
	Get(context.Context, *GetRequest) (*DataResponse, error)
	Update(context.Context, *UpdateRequest) (*DataResponse, error)
	Action(context.Context, *ActionRequest) (*DataResponse, error)
}

// UnimplementedPurchaseServer can be embedded to have forward compatible implementations.
type UnimplementedPurchaseServer struct {
}

func (*UnimplementedPurchaseServer) Raw(ctx context.Context, req *RawRequest) (*RawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Raw not implemented")
}
func (*UnimplementedPurchaseServer) Create(ctx context.Context, req *CreateRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedPurchaseServer) Get(ctx context.Context, req *GetRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedPurchaseServer) Update(ctx context.Context, req *UpdateRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPurchaseServer) Action(ctx context.Context, req *ActionRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}

func RegisterPurchaseServer(s *grpc.Server, srv PurchaseServer) {
	s.RegisterService(&_Purchase_serviceDesc, srv)
}

func _Purchase_Raw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServer).Raw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseproto.Purchase/Raw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServer).Raw(ctx, req.(*RawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purchase_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseproto.Purchase/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purchase_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseproto.Purchase/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purchase_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseproto.Purchase/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purchase_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/purchaseproto.Purchase/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServer).Action(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Purchase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "purchaseproto.Purchase",
	HandlerType: (*PurchaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Raw",
			Handler:    _Purchase_Raw_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Purchase_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Purchase_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Purchase_Update_Handler,
		},
		{
			MethodName: "Action",
			Handler:    _Purchase_Action_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
