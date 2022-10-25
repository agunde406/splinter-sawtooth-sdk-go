// Copyright 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: protobuf/transaction_receipt_pb2/transaction_receipt.proto

package txn_receipt_pb2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	events_pb2 "github.com/hyperledger/sawtooth-sdk-go/protobuf/events_pb2"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StateChange_Type int32

const (
	StateChange_TYPE_UNSET StateChange_Type = 0
	StateChange_SET        StateChange_Type = 1
	StateChange_DELETE     StateChange_Type = 2
)

// Enum value maps for StateChange_Type.
var (
	StateChange_Type_name = map[int32]string{
		0: "TYPE_UNSET",
		1: "SET",
		2: "DELETE",
	}
	StateChange_Type_value = map[string]int32{
		"TYPE_UNSET": 0,
		"SET":        1,
		"DELETE":     2,
	}
)

func (x StateChange_Type) Enum() *StateChange_Type {
	p := new(StateChange_Type)
	*p = x
	return p
}

func (x StateChange_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateChange_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_enumTypes[0].Descriptor()
}

func (StateChange_Type) Type() protoreflect.EnumType {
	return &file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_enumTypes[0]
}

func (x StateChange_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateChange_Type.Descriptor instead.
func (StateChange_Type) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescGZIP(), []int{1, 0}
}

type TransactionReceipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// State changes made by this transaction
	// StateChange is defined in protos/transaction_receipt.proto
	StateChanges []*StateChange `protobuf:"bytes,1,rep,name=state_changes,json=stateChanges,proto3" json:"state_changes,omitempty"`
	// Events fired by this transaction
	Events []*events_pb2.Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	// Transaction family defined data
	Data          [][]byte `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	TransactionId string   `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *TransactionReceipt) Reset() {
	*x = TransactionReceipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionReceipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionReceipt) ProtoMessage() {}

func (x *TransactionReceipt) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionReceipt.ProtoReflect.Descriptor instead.
func (*TransactionReceipt) Descriptor() ([]byte, []int) {
	return file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionReceipt) GetStateChanges() []*StateChange {
	if x != nil {
		return x.StateChanges
	}
	return nil
}

func (x *TransactionReceipt) GetEvents() []*events_pb2.Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *TransactionReceipt) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TransactionReceipt) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

//  StateChange objects have the type of SET, which is either an insert or
//  update, or DELETE. Items marked as a DELETE will have no byte value.
type StateChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string           `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Value   []byte           `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type    StateChange_Type `protobuf:"varint,3,opt,name=type,proto3,enum=StateChange_Type" json:"type,omitempty"`
}

func (x *StateChange) Reset() {
	*x = StateChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateChange) ProtoMessage() {}

func (x *StateChange) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateChange.ProtoReflect.Descriptor instead.
func (*StateChange) Descriptor() ([]byte, []int) {
	return file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescGZIP(), []int{1}
}

func (x *StateChange) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *StateChange) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *StateChange) GetType() StateChange_Type {
	if x != nil {
		return x.Type
	}
	return StateChange_TYPE_UNSET
}

// A collection of state changes.
type StateChangeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateChanges []*StateChange `protobuf:"bytes,1,rep,name=state_changes,json=stateChanges,proto3" json:"state_changes,omitempty"`
}

func (x *StateChangeList) Reset() {
	*x = StateChangeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateChangeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateChangeList) ProtoMessage() {}

func (x *StateChangeList) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateChangeList.ProtoReflect.Descriptor instead.
func (*StateChangeList) Descriptor() ([]byte, []int) {
	return file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescGZIP(), []int{2}
}

func (x *StateChangeList) GetStateChanges() []*StateChange {
	if x != nil {
		return x.StateChanges
	}
	return nil
}

var File_protobuf_transaction_receipt_pb2_transaction_receipt_proto protoreflect.FileDescriptor

var file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x70,
	0x62, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x62,
	0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2,
	0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x31, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x22, 0x44, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0d, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0x33, 0x0a,
	0x15, 0x73, 0x61, 0x77, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x5a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x78, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x70,
	0x62, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescOnce sync.Once
	file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescData = file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDesc
)

func file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescGZIP() []byte {
	file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescOnce.Do(func() {
		file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescData)
	})
	return file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDescData
}

var file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_goTypes = []interface{}{
	(StateChange_Type)(0),      // 0: StateChange.Type
	(*TransactionReceipt)(nil), // 1: TransactionReceipt
	(*StateChange)(nil),        // 2: StateChange
	(*StateChangeList)(nil),    // 3: StateChangeList
	(*events_pb2.Event)(nil),   // 4: Event
}
var file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_depIdxs = []int32{
	2, // 0: TransactionReceipt.state_changes:type_name -> StateChange
	4, // 1: TransactionReceipt.events:type_name -> Event
	0, // 2: StateChange.type:type_name -> StateChange.Type
	2, // 3: StateChangeList.state_changes:type_name -> StateChange
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_init() }
func file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_init() {
	if File_protobuf_transaction_receipt_pb2_transaction_receipt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionReceipt); i {
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
		file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateChange); i {
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
		file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateChangeList); i {
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
			RawDescriptor: file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_goTypes,
		DependencyIndexes: file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_depIdxs,
		EnumInfos:         file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_enumTypes,
		MessageInfos:      file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_msgTypes,
	}.Build()
	File_protobuf_transaction_receipt_pb2_transaction_receipt_proto = out.File
	file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_rawDesc = nil
	file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_goTypes = nil
	file_protobuf_transaction_receipt_pb2_transaction_receipt_proto_depIdxs = nil
}
