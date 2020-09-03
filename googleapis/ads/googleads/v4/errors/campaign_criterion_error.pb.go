// Copyright 2020 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/ads/googleads/v4/errors/campaign_criterion_error.proto

package errors

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Enum describing possible campaign criterion errors.
type CampaignCriterionErrorEnum_CampaignCriterionError int32

const (
	// Enum unspecified.
	CampaignCriterionErrorEnum_UNSPECIFIED CampaignCriterionErrorEnum_CampaignCriterionError = 0
	// The received error code is not known in this version.
	CampaignCriterionErrorEnum_UNKNOWN CampaignCriterionErrorEnum_CampaignCriterionError = 1
	// Concrete type of criterion (keyword v.s. placement) is required for
	// CREATE and UPDATE operations.
	CampaignCriterionErrorEnum_CONCRETE_TYPE_REQUIRED CampaignCriterionErrorEnum_CampaignCriterionError = 2
	// Invalid placement URL.
	CampaignCriterionErrorEnum_INVALID_PLACEMENT_URL CampaignCriterionErrorEnum_CampaignCriterionError = 3
	// Criteria type can not be excluded for the campaign by the customer. like
	// AOL account type cannot target site type criteria
	CampaignCriterionErrorEnum_CANNOT_EXCLUDE_CRITERIA_TYPE CampaignCriterionErrorEnum_CampaignCriterionError = 4
	// Cannot set the campaign criterion status for this criteria type.
	CampaignCriterionErrorEnum_CANNOT_SET_STATUS_FOR_CRITERIA_TYPE CampaignCriterionErrorEnum_CampaignCriterionError = 5
	// Cannot set the campaign criterion status for an excluded criteria.
	CampaignCriterionErrorEnum_CANNOT_SET_STATUS_FOR_EXCLUDED_CRITERIA CampaignCriterionErrorEnum_CampaignCriterionError = 6
	// Cannot target and exclude the same criterion.
	CampaignCriterionErrorEnum_CANNOT_TARGET_AND_EXCLUDE CampaignCriterionErrorEnum_CampaignCriterionError = 7
	// The mutate contained too many operations.
	CampaignCriterionErrorEnum_TOO_MANY_OPERATIONS CampaignCriterionErrorEnum_CampaignCriterionError = 8
	// This operator cannot be applied to a criterion of this type.
	CampaignCriterionErrorEnum_OPERATOR_NOT_SUPPORTED_FOR_CRITERION_TYPE CampaignCriterionErrorEnum_CampaignCriterionError = 9
	// The Shopping campaign sales country is not supported for
	// ProductSalesChannel targeting.
	CampaignCriterionErrorEnum_SHOPPING_CAMPAIGN_SALES_COUNTRY_NOT_SUPPORTED_FOR_SALES_CHANNEL CampaignCriterionErrorEnum_CampaignCriterionError = 10
	// The existing field can't be updated with CREATE operation. It can be
	// updated with UPDATE operation only.
	CampaignCriterionErrorEnum_CANNOT_ADD_EXISTING_FIELD CampaignCriterionErrorEnum_CampaignCriterionError = 11
	// Negative criteria are immutable, so updates are not allowed.
	CampaignCriterionErrorEnum_CANNOT_UPDATE_NEGATIVE_CRITERION CampaignCriterionErrorEnum_CampaignCriterionError = 12
)

// Enum value maps for CampaignCriterionErrorEnum_CampaignCriterionError.
var (
	CampaignCriterionErrorEnum_CampaignCriterionError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CONCRETE_TYPE_REQUIRED",
		3:  "INVALID_PLACEMENT_URL",
		4:  "CANNOT_EXCLUDE_CRITERIA_TYPE",
		5:  "CANNOT_SET_STATUS_FOR_CRITERIA_TYPE",
		6:  "CANNOT_SET_STATUS_FOR_EXCLUDED_CRITERIA",
		7:  "CANNOT_TARGET_AND_EXCLUDE",
		8:  "TOO_MANY_OPERATIONS",
		9:  "OPERATOR_NOT_SUPPORTED_FOR_CRITERION_TYPE",
		10: "SHOPPING_CAMPAIGN_SALES_COUNTRY_NOT_SUPPORTED_FOR_SALES_CHANNEL",
		11: "CANNOT_ADD_EXISTING_FIELD",
		12: "CANNOT_UPDATE_NEGATIVE_CRITERION",
	}
	CampaignCriterionErrorEnum_CampaignCriterionError_value = map[string]int32{
		"UNSPECIFIED":                               0,
		"UNKNOWN":                                   1,
		"CONCRETE_TYPE_REQUIRED":                    2,
		"INVALID_PLACEMENT_URL":                     3,
		"CANNOT_EXCLUDE_CRITERIA_TYPE":              4,
		"CANNOT_SET_STATUS_FOR_CRITERIA_TYPE":       5,
		"CANNOT_SET_STATUS_FOR_EXCLUDED_CRITERIA":   6,
		"CANNOT_TARGET_AND_EXCLUDE":                 7,
		"TOO_MANY_OPERATIONS":                       8,
		"OPERATOR_NOT_SUPPORTED_FOR_CRITERION_TYPE": 9,
		"SHOPPING_CAMPAIGN_SALES_COUNTRY_NOT_SUPPORTED_FOR_SALES_CHANNEL": 10,
		"CANNOT_ADD_EXISTING_FIELD":                                       11,
		"CANNOT_UPDATE_NEGATIVE_CRITERION":                                12,
	}
)

func (x CampaignCriterionErrorEnum_CampaignCriterionError) Enum() *CampaignCriterionErrorEnum_CampaignCriterionError {
	p := new(CampaignCriterionErrorEnum_CampaignCriterionError)
	*p = x
	return p
}

func (x CampaignCriterionErrorEnum_CampaignCriterionError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignCriterionErrorEnum_CampaignCriterionError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_enumTypes[0].Descriptor()
}

func (CampaignCriterionErrorEnum_CampaignCriterionError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_enumTypes[0]
}

func (x CampaignCriterionErrorEnum_CampaignCriterionError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignCriterionErrorEnum_CampaignCriterionError.Descriptor instead.
func (CampaignCriterionErrorEnum_CampaignCriterionError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible campaign criterion errors.
type CampaignCriterionErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignCriterionErrorEnum) Reset() {
	*x = CampaignCriterionErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignCriterionErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignCriterionErrorEnum) ProtoMessage() {}

func (x *CampaignCriterionErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignCriterionErrorEnum.ProtoReflect.Descriptor instead.
func (*CampaignCriterionErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v4_errors_campaign_criterion_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x03,
	0x0a, 0x1a, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xd6, 0x03, 0x0a,
	0x16, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x43, 0x52, 0x45, 0x54,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x4c, 0x41,
	0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x5f, 0x43,
	0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x04, 0x12, 0x27,
	0x0a, 0x23, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x41,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x05, 0x12, 0x2b, 0x0a, 0x27, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52,
	0x49, 0x41, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54,
	0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44,
	0x45, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x08, 0x12, 0x2d, 0x0a, 0x29,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45,
	0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x09, 0x12, 0x43, 0x0a, 0x3f, 0x53,
	0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e,
	0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x10, 0x0a,
	0x12, 0x1d, 0x0a, 0x19, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x45,
	0x58, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x0b, 0x12,
	0x24, 0x0a, 0x20, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x4e, 0x45, 0x47, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52,
	0x49, 0x4f, 0x4e, 0x10, 0x0c, 0x42, 0xf6, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x1b, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x34, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x34, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x34, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x34, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDescData = file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDesc
)

func file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDescData
}

var file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_goTypes = []interface{}{
	(CampaignCriterionErrorEnum_CampaignCriterionError)(0), // 0: google.ads.googleads.v4.errors.CampaignCriterionErrorEnum.CampaignCriterionError
	(*CampaignCriterionErrorEnum)(nil),                     // 1: google.ads.googleads.v4.errors.CampaignCriterionErrorEnum
}
var file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_init() }
func file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_init() {
	if File_google_ads_googleads_v4_errors_campaign_criterion_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignCriterionErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v4_errors_campaign_criterion_error_proto = out.File
	file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_rawDesc = nil
	file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_goTypes = nil
	file_google_ads_googleads_v4_errors_campaign_criterion_error_proto_depIdxs = nil
}
