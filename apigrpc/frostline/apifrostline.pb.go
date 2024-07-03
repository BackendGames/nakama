// Copyright 2018 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//*
// The Nakama server RPC protocol for games and apps.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.4
// source: apifrostline.proto

package apigrpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	DisplayName string `protobuf:"bytes,4,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Lang        string `protobuf:"bytes,5,opt,name=lang,proto3" json:"lang,omitempty"`
	AvatarUrl   string `protobuf:"bytes,6,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	Location    string `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	Timezone    string `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Coins       string `protobuf:"bytes,9,opt,name=coins,proto3" json:"coins,omitempty"`
	Gems        string `protobuf:"bytes,10,opt,name=gems,proto3" json:"gems,omitempty"`
	Xp          string `protobuf:"bytes,11,opt,name=xp,proto3" json:"xp,omitempty"`
	Level       string `protobuf:"bytes,12,opt,name=level,proto3" json:"level,omitempty"`
	LastLogin   string `protobuf:"bytes,13,opt,name=lastLogin,proto3" json:"lastLogin,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apifrostline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apifrostline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_apifrostline_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfoResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfoResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfoResponse) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *UserInfoResponse) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *UserInfoResponse) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserInfoResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UserInfoResponse) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *UserInfoResponse) GetCoins() string {
	if x != nil {
		return x.Coins
	}
	return ""
}

func (x *UserInfoResponse) GetGems() string {
	if x != nil {
		return x.Gems
	}
	return ""
}

func (x *UserInfoResponse) GetXp() string {
	if x != nil {
		return x.Xp
	}
	return ""
}

func (x *UserInfoResponse) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *UserInfoResponse) GetLastLogin() string {
	if x != nil {
		return x.LastLogin
	}
	return ""
}

type GetUserInventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items []*InventoryItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetUserInventory) Reset() {
	*x = GetUserInventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apifrostline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInventory) ProtoMessage() {}

func (x *GetUserInventory) ProtoReflect() protoreflect.Message {
	mi := &file_apifrostline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInventory.ProtoReflect.Descriptor instead.
func (*GetUserInventory) Descriptor() ([]byte, []int) {
	return file_apifrostline_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInventory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetUserInventory) GetItems() []*InventoryItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type InventoryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Currency    string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	ImageUrl    string `protobuf:"bytes,6,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
}

func (x *InventoryItem) Reset() {
	*x = InventoryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apifrostline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryItem) ProtoMessage() {}

func (x *InventoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_apifrostline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryItem.ProtoReflect.Descriptor instead.
func (*InventoryItem) Descriptor() ([]byte, []int) {
	return file_apifrostline_proto_rawDescGZIP(), []int{2}
}

func (x *InventoryItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InventoryItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InventoryItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InventoryItem) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *InventoryItem) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *InventoryItem) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type GetShopData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items []*ShopItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetShopData) Reset() {
	*x = GetShopData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apifrostline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShopData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopData) ProtoMessage() {}

func (x *GetShopData) ProtoReflect() protoreflect.Message {
	mi := &file_apifrostline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopData.ProtoReflect.Descriptor instead.
func (*GetShopData) Descriptor() ([]byte, []int) {
	return file_apifrostline_proto_rawDescGZIP(), []int{3}
}

func (x *GetShopData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetShopData) GetItems() []*ShopItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ShopItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Currency    string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	ImageUrl    string `protobuf:"bytes,6,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
}

func (x *ShopItem) Reset() {
	*x = ShopItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apifrostline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopItem) ProtoMessage() {}

func (x *ShopItem) ProtoReflect() protoreflect.Message {
	mi := &file_apifrostline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopItem.ProtoReflect.Descriptor instead.
func (*ShopItem) Descriptor() ([]byte, []int) {
	return file_apifrostline_proto_rawDescGZIP(), []int{4}
}

func (x *ShopItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShopItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShopItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ShopItem) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ShopItem) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ShopItem) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type GetTransactionsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetTransactionsList) Reset() {
	*x = GetTransactionsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apifrostline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsList) ProtoMessage() {}

func (x *GetTransactionsList) ProtoReflect() protoreflect.Message {
	mi := &file_apifrostline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsList.ProtoReflect.Descriptor instead.
func (*GetTransactionsList) Descriptor() ([]byte, []int) {
	return file_apifrostline_proto_rawDescGZIP(), []int{5}
}

func (x *GetTransactionsList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetTransactionsList) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apifrostline_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_apifrostline_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_apifrostline_proto_rawDescGZIP(), []int{6}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Transaction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetTransactionResponse) Reset() {
	*x = GetTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apifrostline_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionResponse) ProtoMessage() {}

func (x *GetTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apifrostline_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionResponse) Descriptor() ([]byte, []int) {
	return file_apifrostline_proto_rawDescGZIP(), []int{7}
}

func (x *GetTransactionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetTransactionResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTransactionResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type TransactionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TransactionId) Reset() {
	*x = TransactionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apifrostline_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionId) ProtoMessage() {}

func (x *TransactionId) ProtoReflect() protoreflect.Message {
	mi := &file_apifrostline_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionId.ProtoReflect.Descriptor instead.
func (*TransactionId) Descriptor() ([]byte, []int) {
	return file_apifrostline_proto_rawDescGZIP(), []int{8}
}

func (x *TransactionId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apifrostline_proto protoreflect.FileDescriptor

var file_apifrostline_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce,
	0x02, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x65, 0x6d, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x65, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x78,
	0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x78, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22,
	0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x4c, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x72,
	0x6f, 0x73, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x08,
	0x53, 0x68, 0x6f, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x65, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x72, 0x6f, 0x73,
	0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x53, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1f, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x93, 0x04, 0x0a, 0x06, 0x4e, 0x61,
	0x6b, 0x61, 0x6d, 0x61, 0x12, 0x5f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x66, 0x72,
	0x6f, 0x73, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x65, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x56, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1a, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x6e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x22, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x32,
	0x2f, 0x67, 0x65, 0x74, 0x75, 0x73, 0x65, 0x72, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x79, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x66, 0x72, 0x6f,
	0x73, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x25, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65,
	0x72, 0x6f, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2f,
	0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apifrostline_proto_rawDescOnce sync.Once
	file_apifrostline_proto_rawDescData = file_apifrostline_proto_rawDesc
)

func file_apifrostline_proto_rawDescGZIP() []byte {
	file_apifrostline_proto_rawDescOnce.Do(func() {
		file_apifrostline_proto_rawDescData = protoimpl.X.CompressGZIP(file_apifrostline_proto_rawDescData)
	})
	return file_apifrostline_proto_rawDescData
}

var file_apifrostline_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apifrostline_proto_goTypes = []interface{}{
	(*UserInfoResponse)(nil),       // 0: frostline.api.UserInfoResponse
	(*GetUserInventory)(nil),       // 1: frostline.api.GetUserInventory
	(*InventoryItem)(nil),          // 2: frostline.api.InventoryItem
	(*GetShopData)(nil),            // 3: frostline.api.GetShopData
	(*ShopItem)(nil),               // 4: frostline.api.ShopItem
	(*GetTransactionsList)(nil),    // 5: frostline.api.GetTransactionsList
	(*Transaction)(nil),            // 6: frostline.api.Transaction
	(*GetTransactionResponse)(nil), // 7: frostline.api.GetTransactionResponse
	(*TransactionId)(nil),          // 8: frostline.api.TransactionId
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_apifrostline_proto_depIdxs = []int32{
	2, // 0: frostline.api.GetUserInventory.items:type_name -> frostline.api.InventoryItem
	4, // 1: frostline.api.GetShopData.items:type_name -> frostline.api.ShopItem
	6, // 2: frostline.api.GetTransactionsList.transactions:type_name -> frostline.api.Transaction
	9, // 3: frostline.api.Nakama.GetUserInfo:input_type -> google.protobuf.Empty
	9, // 4: frostline.api.Nakama.GetInventory:input_type -> google.protobuf.Empty
	9, // 5: frostline.api.Nakama.GetShop:input_type -> google.protobuf.Empty
	9, // 6: frostline.api.Nakama.GetTransactions:input_type -> google.protobuf.Empty
	8, // 7: frostline.api.Nakama.GetTransactionInfo:input_type -> frostline.api.TransactionId
	0, // 8: frostline.api.Nakama.GetUserInfo:output_type -> frostline.api.UserInfoResponse
	1, // 9: frostline.api.Nakama.GetInventory:output_type -> frostline.api.GetUserInventory
	3, // 10: frostline.api.Nakama.GetShop:output_type -> frostline.api.GetShopData
	5, // 11: frostline.api.Nakama.GetTransactions:output_type -> frostline.api.GetTransactionsList
	7, // 12: frostline.api.Nakama.GetTransactionInfo:output_type -> frostline.api.GetTransactionResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apifrostline_proto_init() }
func file_apifrostline_proto_init() {
	if File_apifrostline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apifrostline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
		file_apifrostline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInventory); i {
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
		file_apifrostline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryItem); i {
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
		file_apifrostline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShopData); i {
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
		file_apifrostline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopItem); i {
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
		file_apifrostline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsList); i {
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
		file_apifrostline_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_apifrostline_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionResponse); i {
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
		file_apifrostline_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionId); i {
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
			RawDescriptor: file_apifrostline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apifrostline_proto_goTypes,
		DependencyIndexes: file_apifrostline_proto_depIdxs,
		MessageInfos:      file_apifrostline_proto_msgTypes,
	}.Build()
	File_apifrostline_proto = out.File
	file_apifrostline_proto_rawDesc = nil
	file_apifrostline_proto_goTypes = nil
	file_apifrostline_proto_depIdxs = nil
}
