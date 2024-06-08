// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: grpcApi/service.proto

package grpcApi

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key             string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	CryptoAlgorithm string `protobuf:"bytes,2,opt,name=crypto_algorithm,json=cryptoAlgorithm,proto3" json:"crypto_algorithm,omitempty"`
	Padding         string `protobuf:"bytes,3,opt,name=padding,proto3" json:"padding,omitempty"`
	CipherMode      string `protobuf:"bytes,4,opt,name=cipher_mode,json=cipherMode,proto3" json:"cipher_mode,omitempty"`
	Content         string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcApi_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_grpcApi_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_grpcApi_service_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Message) GetCryptoAlgorithm() string {
	if x != nil {
		return x.CryptoAlgorithm
	}
	return ""
}

func (x *Message) GetPadding() string {
	if x != nil {
		return x.Padding
	}
	return ""
}

func (x *Message) GetCipherMode() string {
	if x != nil {
		return x.CipherMode
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MessageLikeFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Format  string   `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *MessageLikeFile) Reset() {
	*x = MessageLikeFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcApi_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageLikeFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageLikeFile) ProtoMessage() {}

func (x *MessageLikeFile) ProtoReflect() protoreflect.Message {
	mi := &file_grpcApi_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageLikeFile.ProtoReflect.Descriptor instead.
func (*MessageLikeFile) Descriptor() ([]byte, []int) {
	return file_grpcApi_service_proto_rawDescGZIP(), []int{1}
}

func (x *MessageLikeFile) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *MessageLikeFile) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

type EncryptDecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EncryptDecryptRequest) Reset() {
	*x = EncryptDecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcApi_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptDecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptDecryptRequest) ProtoMessage() {}

func (x *EncryptDecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcApi_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptDecryptRequest.ProtoReflect.Descriptor instead.
func (*EncryptDecryptRequest) Descriptor() ([]byte, []int) {
	return file_grpcApi_service_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptDecryptRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type EncryptDecryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Message `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EncryptDecryptResponse) Reset() {
	*x = EncryptDecryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcApi_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptDecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptDecryptResponse) ProtoMessage() {}

func (x *EncryptDecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcApi_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptDecryptResponse.ProtoReflect.Descriptor instead.
func (*EncryptDecryptResponse) Descriptor() ([]byte, []int) {
	return file_grpcApi_service_proto_rawDescGZIP(), []int{3}
}

func (x *EncryptDecryptResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EncryptDecryptResponse) GetData() *Message {
	if x != nil {
		return x.Data
	}
	return nil
}

type EncryptDecryptFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *MessageLikeFile `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EncryptDecryptFileRequest) Reset() {
	*x = EncryptDecryptFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcApi_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptDecryptFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptDecryptFileRequest) ProtoMessage() {}

func (x *EncryptDecryptFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcApi_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptDecryptFileRequest.ProtoReflect.Descriptor instead.
func (*EncryptDecryptFileRequest) Descriptor() ([]byte, []int) {
	return file_grpcApi_service_proto_rawDescGZIP(), []int{4}
}

func (x *EncryptDecryptFileRequest) GetMessage() *MessageLikeFile {
	if x != nil {
		return x.Message
	}
	return nil
}

type EncryptDecryptFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string           `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    *MessageLikeFile `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EncryptDecryptFileResponse) Reset() {
	*x = EncryptDecryptFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcApi_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptDecryptFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptDecryptFileResponse) ProtoMessage() {}

func (x *EncryptDecryptFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcApi_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptDecryptFileResponse.ProtoReflect.Descriptor instead.
func (*EncryptDecryptFileResponse) Descriptor() ([]byte, []int) {
	return file_grpcApi_service_proto_rawDescGZIP(), []int{5}
}

func (x *EncryptDecryptFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EncryptDecryptFileResponse) GetData() *MessageLikeFile {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_grpcApi_service_proto protoreflect.FileDescriptor

var file_grpcApi_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x9b, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x55,
	0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x43, 0x0a, 0x15, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x58, 0x0a, 0x16, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x4f, 0x0a, 0x19, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44,
	0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x1a, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6b,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xf7, 0x02, 0x0a, 0x0d,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a,
	0x0b, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x65,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x65,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x0b, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x65,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x65,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a,
	0x17, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x65,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x62, 0x0a, 0x17, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x65,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x3b, 0x67,
	0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcApi_service_proto_rawDescOnce sync.Once
	file_grpcApi_service_proto_rawDescData = file_grpcApi_service_proto_rawDesc
)

func file_grpcApi_service_proto_rawDescGZIP() []byte {
	file_grpcApi_service_proto_rawDescOnce.Do(func() {
		file_grpcApi_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcApi_service_proto_rawDescData)
	})
	return file_grpcApi_service_proto_rawDescData
}

var file_grpcApi_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpcApi_service_proto_goTypes = []interface{}{
	(*Message)(nil),                    // 0: service.Message
	(*MessageLikeFile)(nil),            // 1: service.MessageLikeFile
	(*EncryptDecryptRequest)(nil),      // 2: service.EncryptDecryptRequest
	(*EncryptDecryptResponse)(nil),     // 3: service.EncryptDecryptResponse
	(*EncryptDecryptFileRequest)(nil),  // 4: service.EncryptDecryptFileRequest
	(*EncryptDecryptFileResponse)(nil), // 5: service.EncryptDecryptFileResponse
}
var file_grpcApi_service_proto_depIdxs = []int32{
	0, // 0: service.MessageLikeFile.message:type_name -> service.Message
	0, // 1: service.EncryptDecryptRequest.message:type_name -> service.Message
	0, // 2: service.EncryptDecryptResponse.data:type_name -> service.Message
	1, // 3: service.EncryptDecryptFileRequest.message:type_name -> service.MessageLikeFile
	1, // 4: service.EncryptDecryptFileResponse.data:type_name -> service.MessageLikeFile
	2, // 5: service.CryptoService.EncryptData:input_type -> service.EncryptDecryptRequest
	2, // 6: service.CryptoService.DecryptData:input_type -> service.EncryptDecryptRequest
	4, // 7: service.CryptoService.EncryptFileFromComputer:input_type -> service.EncryptDecryptFileRequest
	4, // 8: service.CryptoService.DecryptFileFromComputer:input_type -> service.EncryptDecryptFileRequest
	3, // 9: service.CryptoService.EncryptData:output_type -> service.EncryptDecryptResponse
	3, // 10: service.CryptoService.DecryptData:output_type -> service.EncryptDecryptResponse
	5, // 11: service.CryptoService.EncryptFileFromComputer:output_type -> service.EncryptDecryptFileResponse
	5, // 12: service.CryptoService.DecryptFileFromComputer:output_type -> service.EncryptDecryptFileResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpcApi_service_proto_init() }
func file_grpcApi_service_proto_init() {
	if File_grpcApi_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcApi_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_grpcApi_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageLikeFile); i {
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
		file_grpcApi_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptDecryptRequest); i {
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
		file_grpcApi_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptDecryptResponse); i {
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
		file_grpcApi_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptDecryptFileRequest); i {
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
		file_grpcApi_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptDecryptFileResponse); i {
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
			RawDescriptor: file_grpcApi_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcApi_service_proto_goTypes,
		DependencyIndexes: file_grpcApi_service_proto_depIdxs,
		MessageInfos:      file_grpcApi_service_proto_msgTypes,
	}.Build()
	File_grpcApi_service_proto = out.File
	file_grpcApi_service_proto_rawDesc = nil
	file_grpcApi_service_proto_goTypes = nil
	file_grpcApi_service_proto_depIdxs = nil
}