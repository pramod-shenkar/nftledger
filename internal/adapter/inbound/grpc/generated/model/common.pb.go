// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0--rc1
// source: model/common.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_created   Status = 0
	Status_approved  Status = 1
	Status_rejected  Status = 2
	Status_processed Status = 3
	Status_failed    Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "created",
		1: "approved",
		2: "rejected",
		3: "processed",
		4: "failed",
	}
	Status_value = map[string]int32{
		"created":   0,
		"approved":  1,
		"rejected":  2,
		"processed": 3,
		"failed":    4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_model_common_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_model_common_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_model_common_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status        Status                 `protobuf:"varint,3,opt,name=status,proto3,enum=model.Status" json:"status,omitempty"`
	Filecontent   string                 `protobuf:"bytes,4,opt,name=filecontent,proto3" json:"filecontent,omitempty"` // note: later make it bytes
	Filetype      string                 `protobuf:"bytes,5,opt,name=filetype,proto3" json:"filetype,omitempty"`
	Hash          string                 `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_model_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_model_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_model_common_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_created
}

func (x *Request) GetFilecontent() string {
	if x != nil {
		return x.Filecontent
	}
	return ""
}

func (x *Request) GetFiletype() string {
	if x != nil {
		return x.Filetype
	}
	return ""
}

func (x *Request) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Request) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Request) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Request) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_model_common_proto protoreflect.FileDescriptor

const file_model_common_proto_rawDesc = "" +
	"\n" +
	"\x12model/common.proto\x12\x05model\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd7\x02\n" +
	"\aRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12%\n" +
	"\x06status\x18\x03 \x01(\x0e2\r.model.statusR\x06status\x12 \n" +
	"\vfilecontent\x18\x04 \x01(\tR\vfilecontent\x12\x1a\n" +
	"\bfiletype\x18\x05 \x01(\tR\bfiletype\x12\x12\n" +
	"\x04hash\x18\x06 \x01(\tR\x04hash\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x129\n" +
	"\n" +
	"deleted_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\tdeletedAt*L\n" +
	"\x06status\x12\v\n" +
	"\acreated\x10\x00\x12\f\n" +
	"\bapproved\x10\x01\x12\f\n" +
	"\brejected\x10\x02\x12\r\n" +
	"\tprocessed\x10\x03\x12\n" +
	"\n" +
	"\x06failed\x10\x04B;H\x01Z7nftledger/internal/adapter/inbound/grpc/generated/modelb\x06proto3"

var (
	file_model_common_proto_rawDescOnce sync.Once
	file_model_common_proto_rawDescData []byte
)

func file_model_common_proto_rawDescGZIP() []byte {
	file_model_common_proto_rawDescOnce.Do(func() {
		file_model_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_common_proto_rawDesc), len(file_model_common_proto_rawDesc)))
	})
	return file_model_common_proto_rawDescData
}

var file_model_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_common_proto_goTypes = []any{
	(Status)(0),                   // 0: model.status
	(*Request)(nil),               // 1: model.Request
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_model_common_proto_depIdxs = []int32{
	0, // 0: model.Request.status:type_name -> model.status
	2, // 1: model.Request.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: model.Request.updated_at:type_name -> google.protobuf.Timestamp
	2, // 3: model.Request.deleted_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_model_common_proto_init() }
func file_model_common_proto_init() {
	if File_model_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_common_proto_rawDesc), len(file_model_common_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_common_proto_goTypes,
		DependencyIndexes: file_model_common_proto_depIdxs,
		EnumInfos:         file_model_common_proto_enumTypes,
		MessageInfos:      file_model_common_proto_msgTypes,
	}.Build()
	File_model_common_proto = out.File
	file_model_common_proto_goTypes = nil
	file_model_common_proto_depIdxs = nil
}
