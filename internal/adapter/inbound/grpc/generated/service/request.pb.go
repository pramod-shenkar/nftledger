// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0--rc1
// source: service/request.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	model "nftledger/internal/adapter/inbound/grpc/generated/model"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_request_proto protoreflect.FileDescriptor

const file_service_request_proto_rawDesc = "" +
	"\n" +
	"\x15service/request.proto\x12\aservice\x1a\x13model/request.proto2\xf7\x02\n" +
	"\x0eRequestService\x12D\n" +
	"\vSaveRequest\x12\x19.model.SaveRequestRequest\x1a\x1a.model.SaveRequestResponse\x12A\n" +
	"\n" +
	"GetRequest\x12\x18.model.GetRequestRequest\x1a\x19.model.GetRequestResponse\x12D\n" +
	"\vGetRequests\x12\x19.model.GetRequestsRequest\x1a\x1a.model.GetRequestsResponse\x12J\n" +
	"\rUpdateRequest\x12\x1b.model.UpdateRequestRequest\x1a\x1c.model.UpdateRequestResponse\x12J\n" +
	"\rDeleteRequest\x12\x1b.model.DeleteRequestRequest\x1a\x1c.model.DeleteRequestResponseB=H\x01Z9nftledger/internal/adapter/inbound/grpc/generated/serviceb\x06proto3"

var file_service_request_proto_goTypes = []any{
	(*model.SaveRequestRequest)(nil),    // 0: model.SaveRequestRequest
	(*model.GetRequestRequest)(nil),     // 1: model.GetRequestRequest
	(*model.GetRequestsRequest)(nil),    // 2: model.GetRequestsRequest
	(*model.UpdateRequestRequest)(nil),  // 3: model.UpdateRequestRequest
	(*model.DeleteRequestRequest)(nil),  // 4: model.DeleteRequestRequest
	(*model.SaveRequestResponse)(nil),   // 5: model.SaveRequestResponse
	(*model.GetRequestResponse)(nil),    // 6: model.GetRequestResponse
	(*model.GetRequestsResponse)(nil),   // 7: model.GetRequestsResponse
	(*model.UpdateRequestResponse)(nil), // 8: model.UpdateRequestResponse
	(*model.DeleteRequestResponse)(nil), // 9: model.DeleteRequestResponse
}
var file_service_request_proto_depIdxs = []int32{
	0, // 0: service.RequestService.SaveRequest:input_type -> model.SaveRequestRequest
	1, // 1: service.RequestService.GetRequest:input_type -> model.GetRequestRequest
	2, // 2: service.RequestService.GetRequests:input_type -> model.GetRequestsRequest
	3, // 3: service.RequestService.UpdateRequest:input_type -> model.UpdateRequestRequest
	4, // 4: service.RequestService.DeleteRequest:input_type -> model.DeleteRequestRequest
	5, // 5: service.RequestService.SaveRequest:output_type -> model.SaveRequestResponse
	6, // 6: service.RequestService.GetRequest:output_type -> model.GetRequestResponse
	7, // 7: service.RequestService.GetRequests:output_type -> model.GetRequestsResponse
	8, // 8: service.RequestService.UpdateRequest:output_type -> model.UpdateRequestResponse
	9, // 9: service.RequestService.DeleteRequest:output_type -> model.DeleteRequestResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_request_proto_init() }
func file_service_request_proto_init() {
	if File_service_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_request_proto_rawDesc), len(file_service_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_request_proto_goTypes,
		DependencyIndexes: file_service_request_proto_depIdxs,
	}.Build()
	File_service_request_proto = out.File
	file_service_request_proto_goTypes = nil
	file_service_request_proto_depIdxs = nil
}
