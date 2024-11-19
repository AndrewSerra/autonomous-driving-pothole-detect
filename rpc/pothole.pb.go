// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: pothole.proto

package rpc

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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	mi := &file_pothole_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_pothole_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_pothole_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type RegisterVehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartLocation *Point `protobuf:"bytes,1,opt,name=start_location,json=startLocation,proto3" json:"start_location,omitempty"`
	ListeningAt   string `protobuf:"bytes,2,opt,name=listening_at,json=listeningAt,proto3" json:"listening_at,omitempty"`
}

func (x *RegisterVehicleRequest) Reset() {
	*x = RegisterVehicleRequest{}
	mi := &file_pothole_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterVehicleRequest) ProtoMessage() {}

func (x *RegisterVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pothole_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterVehicleRequest.ProtoReflect.Descriptor instead.
func (*RegisterVehicleRequest) Descriptor() ([]byte, []int) {
	return file_pothole_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterVehicleRequest) GetStartLocation() *Point {
	if x != nil {
		return x.StartLocation
	}
	return nil
}

func (x *RegisterVehicleRequest) GetListeningAt() string {
	if x != nil {
		return x.ListeningAt
	}
	return ""
}

type RegisterVehicleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId string `protobuf:"bytes,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
}

func (x *RegisterVehicleResponse) Reset() {
	*x = RegisterVehicleResponse{}
	mi := &file_pothole_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterVehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterVehicleResponse) ProtoMessage() {}

func (x *RegisterVehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pothole_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterVehicleResponse.ProtoReflect.Descriptor instead.
func (*RegisterVehicleResponse) Descriptor() ([]byte, []int) {
	return file_pothole_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterVehicleResponse) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

type PushLocationUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId   string   `protobuf:"bytes,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	CarLocation *Point   `protobuf:"bytes,2,opt,name=car_location,json=carLocation,proto3" json:"car_location,omitempty"`
	Potholes    []*Point `protobuf:"bytes,3,rep,name=potholes,proto3" json:"potholes,omitempty"`
}

func (x *PushLocationUpdateRequest) Reset() {
	*x = PushLocationUpdateRequest{}
	mi := &file_pothole_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushLocationUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushLocationUpdateRequest) ProtoMessage() {}

func (x *PushLocationUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pothole_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushLocationUpdateRequest.ProtoReflect.Descriptor instead.
func (*PushLocationUpdateRequest) Descriptor() ([]byte, []int) {
	return file_pothole_proto_rawDescGZIP(), []int{3}
}

func (x *PushLocationUpdateRequest) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

func (x *PushLocationUpdateRequest) GetCarLocation() *Point {
	if x != nil {
		return x.CarLocation
	}
	return nil
}

func (x *PushLocationUpdateRequest) GetPotholes() []*Point {
	if x != nil {
		return x.Potholes
	}
	return nil
}

type PushLocationUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *PushLocationUpdateResponse) Reset() {
	*x = PushLocationUpdateResponse{}
	mi := &file_pothole_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushLocationUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushLocationUpdateResponse) ProtoMessage() {}

func (x *PushLocationUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pothole_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushLocationUpdateResponse.ProtoReflect.Descriptor instead.
func (*PushLocationUpdateResponse) Descriptor() ([]byte, []int) {
	return file_pothole_proto_rawDescGZIP(), []int{4}
}

func (x *PushLocationUpdateResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

type ExtendUpcomingAreaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pothole *Point `protobuf:"bytes,1,opt,name=pothole,proto3" json:"pothole,omitempty"`
}

func (x *ExtendUpcomingAreaRequest) Reset() {
	*x = ExtendUpcomingAreaRequest{}
	mi := &file_pothole_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtendUpcomingAreaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendUpcomingAreaRequest) ProtoMessage() {}

func (x *ExtendUpcomingAreaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pothole_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendUpcomingAreaRequest.ProtoReflect.Descriptor instead.
func (*ExtendUpcomingAreaRequest) Descriptor() ([]byte, []int) {
	return file_pothole_proto_rawDescGZIP(), []int{5}
}

func (x *ExtendUpcomingAreaRequest) GetPothole() *Point {
	if x != nil {
		return x.Pothole
	}
	return nil
}

type ExtendUpcomingAreaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *ExtendUpcomingAreaResponse) Reset() {
	*x = ExtendUpcomingAreaResponse{}
	mi := &file_pothole_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtendUpcomingAreaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendUpcomingAreaResponse) ProtoMessage() {}

func (x *ExtendUpcomingAreaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pothole_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendUpcomingAreaResponse.ProtoReflect.Descriptor instead.
func (*ExtendUpcomingAreaResponse) Descriptor() ([]byte, []int) {
	return file_pothole_proto_rawDescGZIP(), []int{6}
}

func (x *ExtendUpcomingAreaResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

var File_pothole_proto protoreflect.FileDescriptor

var file_pothole_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6f, 0x74, 0x68, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x70, 0x63, 0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x6e, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x22, 0x38, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x64, 0x22, 0x91, 0x01, 0x0a, 0x19, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2d,
	0x0a, 0x0c, 0x63, 0x61, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x0b, 0x63, 0x61, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x08, 0x70, 0x6f, 0x74, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x70, 0x6f, 0x74,
	0x68, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x1a, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x22,
	0x41, 0x0a, 0x19, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e,
	0x67, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x07,
	0x70, 0x6f, 0x74, 0x68, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x6f, 0x74, 0x68, 0x6f,
	0x6c, 0x65, 0x22, 0x38, 0x0a, 0x1a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x55, 0x70, 0x63, 0x6f,
	0x6d, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x32, 0x9d, 0x02, 0x0a,
	0x1b, 0x50, 0x6f, 0x74, 0x68, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x50, 0x75,
	0x73, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x12, 0x57, 0x0a, 0x12, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x55, 0x70, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x72,
	0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x72,
	0x65, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x3e, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65,
	0x77, 0x53, 0x65, 0x72, 0x72, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6e, 0x6f, 0x6d, 0x6f, 0x75,
	0x73, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6f, 0x74, 0x68, 0x6f, 0x6c,
	0x65, 0x2d, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pothole_proto_rawDescOnce sync.Once
	file_pothole_proto_rawDescData = file_pothole_proto_rawDesc
)

func file_pothole_proto_rawDescGZIP() []byte {
	file_pothole_proto_rawDescOnce.Do(func() {
		file_pothole_proto_rawDescData = protoimpl.X.CompressGZIP(file_pothole_proto_rawDescData)
	})
	return file_pothole_proto_rawDescData
}

var file_pothole_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pothole_proto_goTypes = []any{
	(*Point)(nil),                      // 0: rpc.Point
	(*RegisterVehicleRequest)(nil),     // 1: rpc.RegisterVehicleRequest
	(*RegisterVehicleResponse)(nil),    // 2: rpc.RegisterVehicleResponse
	(*PushLocationUpdateRequest)(nil),  // 3: rpc.PushLocationUpdateRequest
	(*PushLocationUpdateResponse)(nil), // 4: rpc.PushLocationUpdateResponse
	(*ExtendUpcomingAreaRequest)(nil),  // 5: rpc.ExtendUpcomingAreaRequest
	(*ExtendUpcomingAreaResponse)(nil), // 6: rpc.ExtendUpcomingAreaResponse
}
var file_pothole_proto_depIdxs = []int32{
	0, // 0: rpc.RegisterVehicleRequest.start_location:type_name -> rpc.Point
	0, // 1: rpc.PushLocationUpdateRequest.car_location:type_name -> rpc.Point
	0, // 2: rpc.PushLocationUpdateRequest.potholes:type_name -> rpc.Point
	0, // 3: rpc.ExtendUpcomingAreaRequest.pothole:type_name -> rpc.Point
	1, // 4: rpc.PotholeCommunicationService.RegisterVehicle:input_type -> rpc.RegisterVehicleRequest
	3, // 5: rpc.PotholeCommunicationService.PushLocationUpdate:input_type -> rpc.PushLocationUpdateRequest
	5, // 6: rpc.PotholeCommunicationService.ExtendUpcomingArea:input_type -> rpc.ExtendUpcomingAreaRequest
	2, // 7: rpc.PotholeCommunicationService.RegisterVehicle:output_type -> rpc.RegisterVehicleResponse
	4, // 8: rpc.PotholeCommunicationService.PushLocationUpdate:output_type -> rpc.PushLocationUpdateResponse
	6, // 9: rpc.PotholeCommunicationService.ExtendUpcomingArea:output_type -> rpc.ExtendUpcomingAreaResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pothole_proto_init() }
func file_pothole_proto_init() {
	if File_pothole_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pothole_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pothole_proto_goTypes,
		DependencyIndexes: file_pothole_proto_depIdxs,
		MessageInfos:      file_pothole_proto_msgTypes,
	}.Build()
	File_pothole_proto = out.File
	file_pothole_proto_rawDesc = nil
	file_pothole_proto_goTypes = nil
	file_pothole_proto_depIdxs = nil
}
