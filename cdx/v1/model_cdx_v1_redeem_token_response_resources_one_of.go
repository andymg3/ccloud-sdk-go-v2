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
Stream Share APIs

# Introduction

API version: 0.1.0-alpha0
Contact: cdx@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// CdxV1RedeemTokenResponseResourcesOneOf - struct for CdxV1RedeemTokenResponseResourcesOneOf
type CdxV1RedeemTokenResponseResourcesOneOf struct {
	CdxV1SharedGroup *CdxV1SharedGroup
	CdxV1SharedTopic *CdxV1SharedTopic
}

// CdxV1SharedGroupAsCdxV1RedeemTokenResponseResourcesOneOf is a convenience function that returns CdxV1SharedGroup wrapped in CdxV1RedeemTokenResponseResourcesOneOf
func CdxV1SharedGroupAsCdxV1RedeemTokenResponseResourcesOneOf(v *CdxV1SharedGroup) CdxV1RedeemTokenResponseResourcesOneOf {
	return CdxV1RedeemTokenResponseResourcesOneOf{CdxV1SharedGroup: v}
}

// CdxV1SharedTopicAsCdxV1RedeemTokenResponseResourcesOneOf is a convenience function that returns CdxV1SharedTopic wrapped in CdxV1RedeemTokenResponseResourcesOneOf
func CdxV1SharedTopicAsCdxV1RedeemTokenResponseResourcesOneOf(v *CdxV1SharedTopic) CdxV1RedeemTokenResponseResourcesOneOf {
	return CdxV1RedeemTokenResponseResourcesOneOf{CdxV1SharedTopic: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CdxV1RedeemTokenResponseResourcesOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'Group'
	if jsonDict["kind"] == "Group" {
		// try to unmarshal JSON data into CdxV1SharedGroup
		err = json.Unmarshal(data, &dst.CdxV1SharedGroup)
		if err == nil {
			return nil // data stored in dst.CdxV1SharedGroup, return on the first match
		} else {
			dst.CdxV1SharedGroup = nil
			return fmt.Errorf("Failed to unmarshal CdxV1RedeemTokenResponseResourcesOneOf as CdxV1SharedGroup: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Topic'
	if jsonDict["kind"] == "Topic" {
		// try to unmarshal JSON data into CdxV1SharedTopic
		err = json.Unmarshal(data, &dst.CdxV1SharedTopic)
		if err == nil {
			return nil // data stored in dst.CdxV1SharedTopic, return on the first match
		} else {
			dst.CdxV1SharedTopic = nil
			return fmt.Errorf("Failed to unmarshal CdxV1RedeemTokenResponseResourcesOneOf as CdxV1SharedTopic: %s", err.Error())
		}
	}

	// check if the discriminator value is 'cdx.v1.SharedGroup'
	if jsonDict["kind"] == "cdx.v1.SharedGroup" {
		// try to unmarshal JSON data into CdxV1SharedGroup
		err = json.Unmarshal(data, &dst.CdxV1SharedGroup)
		if err == nil {
			return nil // data stored in dst.CdxV1SharedGroup, return on the first match
		} else {
			dst.CdxV1SharedGroup = nil
			return fmt.Errorf("Failed to unmarshal CdxV1RedeemTokenResponseResourcesOneOf as CdxV1SharedGroup: %s", err.Error())
		}
	}

	// check if the discriminator value is 'cdx.v1.SharedTopic'
	if jsonDict["kind"] == "cdx.v1.SharedTopic" {
		// try to unmarshal JSON data into CdxV1SharedTopic
		err = json.Unmarshal(data, &dst.CdxV1SharedTopic)
		if err == nil {
			return nil // data stored in dst.CdxV1SharedTopic, return on the first match
		} else {
			dst.CdxV1SharedTopic = nil
			return fmt.Errorf("Failed to unmarshal CdxV1RedeemTokenResponseResourcesOneOf as CdxV1SharedTopic: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CdxV1RedeemTokenResponseResourcesOneOf) MarshalJSON() ([]byte, error) {
	if src.CdxV1SharedGroup != nil {
		return json.Marshal(&src.CdxV1SharedGroup)
	}

	if src.CdxV1SharedTopic != nil {
		return json.Marshal(&src.CdxV1SharedTopic)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CdxV1RedeemTokenResponseResourcesOneOf) GetActualInstance() interface{} {
	if obj.CdxV1SharedGroup != nil {
		return obj.CdxV1SharedGroup
	}

	if obj.CdxV1SharedTopic != nil {
		return obj.CdxV1SharedTopic
	}

	// all schemas are nil
	return nil
}

type NullableCdxV1RedeemTokenResponseResourcesOneOf struct {
	value *CdxV1RedeemTokenResponseResourcesOneOf
	isSet bool
}

func (v NullableCdxV1RedeemTokenResponseResourcesOneOf) Get() *CdxV1RedeemTokenResponseResourcesOneOf {
	return v.value
}

func (v *NullableCdxV1RedeemTokenResponseResourcesOneOf) Set(val *CdxV1RedeemTokenResponseResourcesOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCdxV1RedeemTokenResponseResourcesOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCdxV1RedeemTokenResponseResourcesOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdxV1RedeemTokenResponseResourcesOneOf(val *CdxV1RedeemTokenResponseResourcesOneOf) *NullableCdxV1RedeemTokenResponseResourcesOneOf {
	return &NullableCdxV1RedeemTokenResponseResourcesOneOf{value: val, isSet: true}
}

func (v NullableCdxV1RedeemTokenResponseResourcesOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdxV1RedeemTokenResponseResourcesOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
