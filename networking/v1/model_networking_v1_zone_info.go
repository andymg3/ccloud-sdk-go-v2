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
Network API

Network API

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// NetworkingV1ZoneInfo Cloud provider zone metadata.
type NetworkingV1ZoneInfo struct {
	// Cloud provider zone id
	ZoneId *string `json:"zone_id,omitempty"`
	// The IPv4 [CIDR block](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) to used for this network. Must be a `/27`. Required for VPC peering and AWS TransitGateway.
	Cidr *string `json:"cidr,omitempty"`
}

// NewNetworkingV1ZoneInfo instantiates a new NetworkingV1ZoneInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1ZoneInfo() *NetworkingV1ZoneInfo {
	this := NetworkingV1ZoneInfo{}
	return &this
}

// NewNetworkingV1ZoneInfoWithDefaults instantiates a new NetworkingV1ZoneInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1ZoneInfoWithDefaults() *NetworkingV1ZoneInfo {
	this := NetworkingV1ZoneInfo{}
	return &this
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *NetworkingV1ZoneInfo) GetZoneId() string {
	if o == nil || o.ZoneId == nil {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1ZoneInfo) GetZoneIdOk() (*string, bool) {
	if o == nil || o.ZoneId == nil {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *NetworkingV1ZoneInfo) HasZoneId() bool {
	if o != nil && o.ZoneId != nil {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *NetworkingV1ZoneInfo) SetZoneId(v string) {
	o.ZoneId = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *NetworkingV1ZoneInfo) GetCidr() string {
	if o == nil || o.Cidr == nil {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1ZoneInfo) GetCidrOk() (*string, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *NetworkingV1ZoneInfo) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *NetworkingV1ZoneInfo) SetCidr(v string) {
	o.Cidr = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1ZoneInfo) Redact() {
	o.recurseRedact(o.ZoneId)
	o.recurseRedact(o.Cidr)
}

func (o *NetworkingV1ZoneInfo) recurseRedact(v interface{}) {
	type redactor interface {
		Redact()
	}
	if r, ok := v.(redactor); ok {
		r.Redact()
	} else {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < val.Len(); i++ {
				// support data types declared without pointers
				o.recurseRedact(val.Index(i).Interface())
				// ... and data types that were declared without but need pointers (for Redact)
				if val.Index(i).CanAddr() {
					o.recurseRedact(val.Index(i).Addr().Interface())
				}
			}
		}
	}
}

func (o NetworkingV1ZoneInfo) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1ZoneInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneId != nil {
		toSerialize["zone_id"] = o.ZoneId
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1ZoneInfo struct {
	value *NetworkingV1ZoneInfo
	isSet bool
}

func (v NullableNetworkingV1ZoneInfo) Get() *NetworkingV1ZoneInfo {
	return v.value
}

func (v *NullableNetworkingV1ZoneInfo) Set(val *NetworkingV1ZoneInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1ZoneInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1ZoneInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1ZoneInfo(val *NetworkingV1ZoneInfo) *NullableNetworkingV1ZoneInfo {
	return &NullableNetworkingV1ZoneInfo{value: val, isSet: true}
}

func (v NullableNetworkingV1ZoneInfo) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1ZoneInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
