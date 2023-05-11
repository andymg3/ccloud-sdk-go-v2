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

// NetworkingV1PeeringStatus The status of the Peering
type NetworkingV1PeeringStatus struct {
	// The lifecycle phase of the peering:    PROVISIONING: peering provisioning is in progress;    PENDING_ACCEPT: peering connection request is pending acceptance by the customer;    READY:  peering is ready;    FAILED: peering is in a failed state;    DEPROVISIONING: peering deprovisioning is in progress;    DISCONNECTED: peering has been disconnected in the cloud provider by the customer;
	Phase string `json:"phase,omitempty"`
	// Error code if peering is in a failed state. May be used for programmatic error checking.
	ErrorCode *string `json:"error_code,omitempty"`
	// Displayable error message if peering is in a failed state
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewNetworkingV1PeeringStatus instantiates a new NetworkingV1PeeringStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1PeeringStatus(phase string) *NetworkingV1PeeringStatus {
	this := NetworkingV1PeeringStatus{}
	this.Phase = phase
	return &this
}

// NewNetworkingV1PeeringStatusWithDefaults instantiates a new NetworkingV1PeeringStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1PeeringStatusWithDefaults() *NetworkingV1PeeringStatus {
	this := NetworkingV1PeeringStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *NetworkingV1PeeringStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PeeringStatus) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *NetworkingV1PeeringStatus) SetPhase(v string) {
	o.Phase = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *NetworkingV1PeeringStatus) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PeeringStatus) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *NetworkingV1PeeringStatus) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *NetworkingV1PeeringStatus) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *NetworkingV1PeeringStatus) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PeeringStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *NetworkingV1PeeringStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *NetworkingV1PeeringStatus) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1PeeringStatus) Redact() {
	o.recurseRedact(&o.Phase)
	o.recurseRedact(o.ErrorCode)
	o.recurseRedact(o.ErrorMessage)
}

func (o *NetworkingV1PeeringStatus) recurseRedact(v interface{}) {
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

func (o NetworkingV1PeeringStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1PeeringStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["phase"] = o.Phase
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1PeeringStatus struct {
	value *NetworkingV1PeeringStatus
	isSet bool
}

func (v NullableNetworkingV1PeeringStatus) Get() *NetworkingV1PeeringStatus {
	return v.value
}

func (v *NullableNetworkingV1PeeringStatus) Set(val *NetworkingV1PeeringStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1PeeringStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1PeeringStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1PeeringStatus(val *NetworkingV1PeeringStatus) *NullableNetworkingV1PeeringStatus {
	return &NullableNetworkingV1PeeringStatus{value: val, isSet: true}
}

func (v NullableNetworkingV1PeeringStatus) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1PeeringStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
