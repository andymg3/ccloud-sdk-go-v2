// Copyright 2021 Confluent Inc. All Rights Reserved.
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

/*
REST Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0
Contact: kafka-clients-proxy-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"fmt"
)

// AclResourceType the model 'AclResourceType'
type AclResourceType string

// List of AclResourceType
const (
	UNKNOWN AclResourceType = "UNKNOWN"
	ANY AclResourceType = "ANY"
	TOPIC AclResourceType = "TOPIC"
	GROUP AclResourceType = "GROUP"
	CLUSTER AclResourceType = "CLUSTER"
	TRANSACTIONAL_ID AclResourceType = "TRANSACTIONAL_ID"
	DELEGATION_TOKEN AclResourceType = "DELEGATION_TOKEN"
)

var allowedAclResourceTypeEnumValues = []AclResourceType{
	"UNKNOWN",
	"ANY",
	"TOPIC",
	"GROUP",
	"CLUSTER",
	"TRANSACTIONAL_ID",
	"DELEGATION_TOKEN",
}

func (v *AclResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AclResourceType(value)
	for _, existing := range allowedAclResourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AclResourceType", value)
}

// NewAclResourceTypeFromValue returns a pointer to a valid AclResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAclResourceTypeFromValue(v string) (*AclResourceType, error) {
	ev := AclResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AclResourceType: valid values are %v", v, allowedAclResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AclResourceType) IsValid() bool {
	for _, existing := range allowedAclResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AclResourceType value
func (v AclResourceType) Ptr() *AclResourceType {
	return &v
}

type NullableAclResourceType struct {
	value *AclResourceType
	isSet bool
}

func (v NullableAclResourceType) Get() *AclResourceType {
	return v.value
}

func (v *NullableAclResourceType) Set(val *AclResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableAclResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableAclResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclResourceType(val *AclResourceType) *NullableAclResourceType {
	return &NullableAclResourceType{value: val, isSet: true}
}

func (v NullableAclResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

