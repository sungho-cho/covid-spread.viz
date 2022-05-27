// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: covid_data.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_covid_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_covid_data_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

type GetActiveCasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Date    *Date  `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetActiveCasesRequest) Reset() {
	*x = GetActiveCasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActiveCasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActiveCasesRequest) ProtoMessage() {}

func (x *GetActiveCasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_covid_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActiveCasesRequest.ProtoReflect.Descriptor instead.
func (*GetActiveCasesRequest) Descriptor() ([]byte, []int) {
	return file_covid_data_proto_rawDescGZIP(), []int{1}
}

func (x *GetActiveCasesRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetActiveCasesRequest) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

type GetActiveCasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumCases int32 `protobuf:"varint,1,opt,name=num_cases,json=numCases,proto3" json:"num_cases,omitempty"`
}

func (x *GetActiveCasesResponse) Reset() {
	*x = GetActiveCasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActiveCasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActiveCasesResponse) ProtoMessage() {}

func (x *GetActiveCasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_covid_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActiveCasesResponse.ProtoReflect.Descriptor instead.
func (*GetActiveCasesResponse) Descriptor() ([]byte, []int) {
	return file_covid_data_proto_rawDescGZIP(), []int{2}
}

func (x *GetActiveCasesResponse) GetNumCases() int32 {
	if x != nil {
		return x.NumCases
	}
	return 0
}

type GetCountriesDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *Date `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetCountriesDataRequest) Reset() {
	*x = GetCountriesDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountriesDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountriesDataRequest) ProtoMessage() {}

func (x *GetCountriesDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_covid_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountriesDataRequest.ProtoReflect.Descriptor instead.
func (*GetCountriesDataRequest) Descriptor() ([]byte, []int) {
	return file_covid_data_proto_rawDescGZIP(), []int{3}
}

func (x *GetCountriesDataRequest) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

type GetCountriesDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountriesData *CountriesData `protobuf:"bytes,1,opt,name=countries_data,json=countriesData,proto3" json:"countries_data,omitempty"`
}

func (x *GetCountriesDataResponse) Reset() {
	*x = GetCountriesDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountriesDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountriesDataResponse) ProtoMessage() {}

func (x *GetCountriesDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_covid_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountriesDataResponse.ProtoReflect.Descriptor instead.
func (*GetCountriesDataResponse) Descriptor() ([]byte, []int) {
	return file_covid_data_proto_rawDescGZIP(), []int{4}
}

func (x *GetCountriesDataResponse) GetCountriesData() *CountriesData {
	if x != nil {
		return x.CountriesData
	}
	return nil
}

type CountryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country   string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Date      *Date  `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Confirmed int32  `protobuf:"varint,3,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	Recovered int32  `protobuf:"varint,4,opt,name=recovered,proto3" json:"recovered,omitempty"`
	Deaths    int32  `protobuf:"varint,5,opt,name=deaths,proto3" json:"deaths,omitempty"`
}

func (x *CountryData) Reset() {
	*x = CountryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountryData) ProtoMessage() {}

func (x *CountryData) ProtoReflect() protoreflect.Message {
	mi := &file_covid_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountryData.ProtoReflect.Descriptor instead.
func (*CountryData) Descriptor() ([]byte, []int) {
	return file_covid_data_proto_rawDescGZIP(), []int{5}
}

func (x *CountryData) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CountryData) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *CountryData) GetConfirmed() int32 {
	if x != nil {
		return x.Confirmed
	}
	return 0
}

func (x *CountryData) GetRecovered() int32 {
	if x != nil {
		return x.Recovered
	}
	return 0
}

func (x *CountryData) GetDeaths() int32 {
	if x != nil {
		return x.Deaths
	}
	return 0
}

type CountriesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date      *Date          `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Countries []*CountryData `protobuf:"bytes,2,rep,name=countries,proto3" json:"countries,omitempty"`
}

func (x *CountriesData) Reset() {
	*x = CountriesData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesData) ProtoMessage() {}

func (x *CountriesData) ProtoReflect() protoreflect.Message {
	mi := &file_covid_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesData.ProtoReflect.Descriptor instead.
func (*CountriesData) Descriptor() ([]byte, []int) {
	return file_covid_data_proto_rawDescGZIP(), []int{6}
}

func (x *CountriesData) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *CountriesData) GetCountries() []*CountryData {
	if x != nil {
		return x.Countries
	}
	return nil
}

var File_covid_data_proto protoreflect.FileDescriptor

var file_covid_data_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x42, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x22, 0x53,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x22, 0x35, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x61, 0x73, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x58, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x65, 0x61, 0x74, 0x68,
	0x73, 0x22, 0x64, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x32, 0xb7, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x76, 0x69,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x75, 0x6e, 0x67, 0x68, 0x6f, 0x2d, 0x63, 0x68, 0x6f, 0x2f, 0x63, 0x6f, 0x76, 0x69, 0x64,
	0x2d, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x76, 0x69, 0x7a, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_covid_data_proto_rawDescOnce sync.Once
	file_covid_data_proto_rawDescData = file_covid_data_proto_rawDesc
)

func file_covid_data_proto_rawDescGZIP() []byte {
	file_covid_data_proto_rawDescOnce.Do(func() {
		file_covid_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_covid_data_proto_rawDescData)
	})
	return file_covid_data_proto_rawDescData
}

var file_covid_data_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_covid_data_proto_goTypes = []interface{}{
	(*Date)(nil),                     // 0: protos.Date
	(*GetActiveCasesRequest)(nil),    // 1: protos.GetActiveCasesRequest
	(*GetActiveCasesResponse)(nil),   // 2: protos.GetActiveCasesResponse
	(*GetCountriesDataRequest)(nil),  // 3: protos.GetCountriesDataRequest
	(*GetCountriesDataResponse)(nil), // 4: protos.GetCountriesDataResponse
	(*CountryData)(nil),              // 5: protos.CountryData
	(*CountriesData)(nil),            // 6: protos.CountriesData
}
var file_covid_data_proto_depIdxs = []int32{
	0, // 0: protos.GetActiveCasesRequest.date:type_name -> protos.Date
	0, // 1: protos.GetCountriesDataRequest.date:type_name -> protos.Date
	6, // 2: protos.GetCountriesDataResponse.countries_data:type_name -> protos.CountriesData
	0, // 3: protos.CountryData.date:type_name -> protos.Date
	0, // 4: protos.CountriesData.date:type_name -> protos.Date
	5, // 5: protos.CountriesData.countries:type_name -> protos.CountryData
	1, // 6: protos.CovidData.GetActiveCases:input_type -> protos.GetActiveCasesRequest
	3, // 7: protos.CovidData.GetCountriesData:input_type -> protos.GetCountriesDataRequest
	2, // 8: protos.CovidData.GetActiveCases:output_type -> protos.GetActiveCasesResponse
	4, // 9: protos.CovidData.GetCountriesData:output_type -> protos.GetCountriesDataResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_covid_data_proto_init() }
func file_covid_data_proto_init() {
	if File_covid_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_covid_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_covid_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActiveCasesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_covid_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActiveCasesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_covid_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountriesDataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_covid_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountriesDataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_covid_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountryData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_covid_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_covid_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_covid_data_proto_goTypes,
		DependencyIndexes: file_covid_data_proto_depIdxs,
		MessageInfos:      file_covid_data_proto_msgTypes,
	}.Build()
	File_covid_data_proto = out.File
	file_covid_data_proto_rawDesc = nil
	file_covid_data_proto_goTypes = nil
	file_covid_data_proto_depIdxs = nil
}