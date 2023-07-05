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
	"bytes"
	"encoding/json"
	"time"
)

import (
	"reflect"
)

// AnyUnevenLoadData struct for AnyUnevenLoadData
type AnyUnevenLoadData struct {
	Kind           string           `json:"kind,omitempty"`
	Metadata       ResourceMetadata `json:"metadata,omitempty"`
	ClusterId      string           `json:"cluster_id,omitempty"`
	Status         string           `json:"status,omitempty"`
	PreviousStatus string           `json:"previous_status,omitempty"`
	// The date and time at which this task was created.
	StatusUpdatedAt time.Time `json:"status_updated_at,omitempty"`
	// The date and time at which this task was created.
	PreviousStatusUpdatedAt time.Time      `json:"previous_status_updated_at,omitempty"`
	ErrorCode               NullableInt32  `json:"error_code,omitempty"`
	ErrorMessage            NullableString `json:"error_message,omitempty"`
	BrokerTasks             Relationship   `json:"broker_tasks,omitempty"`
}

// NewAnyUnevenLoadData instantiates a new AnyUnevenLoadData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnyUnevenLoadData(kind string, metadata ResourceMetadata, clusterId string, status string, previousStatus string, statusUpdatedAt time.Time, previousStatusUpdatedAt time.Time, brokerTasks Relationship) *AnyUnevenLoadData {
	this := AnyUnevenLoadData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.Status = status
	this.PreviousStatus = previousStatus
	this.StatusUpdatedAt = statusUpdatedAt
	this.PreviousStatusUpdatedAt = previousStatusUpdatedAt
	this.BrokerTasks = brokerTasks
	return &this
}

// NewAnyUnevenLoadDataWithDefaults instantiates a new AnyUnevenLoadData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnyUnevenLoadDataWithDefaults() *AnyUnevenLoadData {
	this := AnyUnevenLoadData{}
	return &this
}

// GetKind returns the Kind field value
func (o *AnyUnevenLoadData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AnyUnevenLoadData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AnyUnevenLoadData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *AnyUnevenLoadData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *AnyUnevenLoadData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *AnyUnevenLoadData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *AnyUnevenLoadData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *AnyUnevenLoadData) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *AnyUnevenLoadData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetStatus returns the Status field value
func (o *AnyUnevenLoadData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AnyUnevenLoadData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AnyUnevenLoadData) SetStatus(v string) {
	o.Status = v
}

// GetPreviousStatus returns the PreviousStatus field value
func (o *AnyUnevenLoadData) GetPreviousStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousStatus
}

// GetPreviousStatusOk returns a tuple with the PreviousStatus field value
// and a boolean to check if the value has been set.
func (o *AnyUnevenLoadData) GetPreviousStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousStatus, true
}

// SetPreviousStatus sets field value
func (o *AnyUnevenLoadData) SetPreviousStatus(v string) {
	o.PreviousStatus = v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value
func (o *AnyUnevenLoadData) GetStatusUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AnyUnevenLoadData) GetStatusUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusUpdatedAt, true
}

// SetStatusUpdatedAt sets field value
func (o *AnyUnevenLoadData) SetStatusUpdatedAt(v time.Time) {
	o.StatusUpdatedAt = v
}

// GetPreviousStatusUpdatedAt returns the PreviousStatusUpdatedAt field value
func (o *AnyUnevenLoadData) GetPreviousStatusUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PreviousStatusUpdatedAt
}

// GetPreviousStatusUpdatedAtOk returns a tuple with the PreviousStatusUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AnyUnevenLoadData) GetPreviousStatusUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousStatusUpdatedAt, true
}

// SetPreviousStatusUpdatedAt sets field value
func (o *AnyUnevenLoadData) SetPreviousStatusUpdatedAt(v time.Time) {
	o.PreviousStatusUpdatedAt = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnyUnevenLoadData) GetErrorCode() int32 {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnyUnevenLoadData) GetErrorCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AnyUnevenLoadData) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableInt32 and assigns it to the ErrorCode field.
func (o *AnyUnevenLoadData) SetErrorCode(v int32) {
	o.ErrorCode.Set(&v)
}

// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *AnyUnevenLoadData) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *AnyUnevenLoadData) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnyUnevenLoadData) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnyUnevenLoadData) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AnyUnevenLoadData) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *AnyUnevenLoadData) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *AnyUnevenLoadData) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *AnyUnevenLoadData) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetBrokerTasks returns the BrokerTasks field value
func (o *AnyUnevenLoadData) GetBrokerTasks() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.BrokerTasks
}

// GetBrokerTasksOk returns a tuple with the BrokerTasks field value
// and a boolean to check if the value has been set.
func (o *AnyUnevenLoadData) GetBrokerTasksOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerTasks, true
}

// SetBrokerTasks sets field value
func (o *AnyUnevenLoadData) SetBrokerTasks(v Relationship) {
	o.BrokerTasks = v
}

// Redact resets all sensitive fields to their zero value.
func (o *AnyUnevenLoadData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.Status)
	o.recurseRedact(&o.PreviousStatus)
	o.recurseRedact(&o.StatusUpdatedAt)
	o.recurseRedact(&o.PreviousStatusUpdatedAt)
	o.recurseRedact(o.ErrorCode)
	o.recurseRedact(o.ErrorMessage)
	o.recurseRedact(&o.BrokerTasks)
}

func (o *AnyUnevenLoadData) recurseRedact(v interface{}) {
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

func (o AnyUnevenLoadData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o AnyUnevenLoadData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["previous_status"] = o.PreviousStatus
	}
	if true {
		toSerialize["status_updated_at"] = o.StatusUpdatedAt
	}
	if true {
		toSerialize["previous_status_updated_at"] = o.PreviousStatusUpdatedAt
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["error_message"] = o.ErrorMessage.Get()
	}
	if true {
		toSerialize["broker_tasks"] = o.BrokerTasks
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableAnyUnevenLoadData struct {
	value *AnyUnevenLoadData
	isSet bool
}

func (v NullableAnyUnevenLoadData) Get() *AnyUnevenLoadData {
	return v.value
}

func (v *NullableAnyUnevenLoadData) Set(val *AnyUnevenLoadData) {
	v.value = val
	v.isSet = true
}

func (v NullableAnyUnevenLoadData) IsSet() bool {
	return v.isSet
}

func (v *NullableAnyUnevenLoadData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnyUnevenLoadData(val *AnyUnevenLoadData) *NullableAnyUnevenLoadData {
	return &NullableAnyUnevenLoadData{value: val, isSet: true}
}

func (v NullableAnyUnevenLoadData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableAnyUnevenLoadData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
