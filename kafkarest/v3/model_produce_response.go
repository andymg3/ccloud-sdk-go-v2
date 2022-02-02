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
	"time"
)

import (
	"reflect"
)

// ProduceResponse struct for ProduceResponse
type ProduceResponse struct {
	ClusterId string `json:"cluster_id"`
	TopicName string `json:"topic_name"`
	PartitionId int32 `json:"partition_id"`
	Offset int32 `json:"offset"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	Key NullableProduceResponseData `json:"key,omitempty"`
	Value NullableProduceResponseData `json:"value,omitempty"`
}

// NewProduceResponse instantiates a new ProduceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduceResponse(clusterId string, topicName string, partitionId int32, offset int32) *ProduceResponse {
	this := ProduceResponse{}
	this.ClusterId = clusterId
	this.TopicName = topicName
	this.PartitionId = partitionId
	this.Offset = offset
	return &this
}

// NewProduceResponseWithDefaults instantiates a new ProduceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduceResponseWithDefaults() *ProduceResponse {
	this := ProduceResponse{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *ProduceResponse) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ProduceResponse) GetClusterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ProduceResponse) SetClusterId(v string) {
	o.ClusterId = v
}

// GetTopicName returns the TopicName field value
func (o *ProduceResponse) GetTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value
// and a boolean to check if the value has been set.
func (o *ProduceResponse) GetTopicNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TopicName, true
}

// SetTopicName sets field value
func (o *ProduceResponse) SetTopicName(v string) {
	o.TopicName = v
}

// GetPartitionId returns the PartitionId field value
func (o *ProduceResponse) GetPartitionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value
// and a boolean to check if the value has been set.
func (o *ProduceResponse) GetPartitionIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PartitionId, true
}

// SetPartitionId sets field value
func (o *ProduceResponse) SetPartitionId(v int32) {
	o.PartitionId = v
}

// GetOffset returns the Offset field value
func (o *ProduceResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ProduceResponse) GetOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ProduceResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProduceResponse) GetTimestamp() time.Time {
	if o == nil || o.Timestamp.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProduceResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ProduceResponse) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *ProduceResponse) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *ProduceResponse) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *ProduceResponse) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProduceResponse) GetKey() ProduceResponseData {
	if o == nil || o.Key.Get() == nil {
		var ret ProduceResponseData
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProduceResponse) GetKeyOk() (*ProduceResponseData, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ProduceResponse) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableProduceResponseData and assigns it to the Key field.
func (o *ProduceResponse) SetKey(v ProduceResponseData) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *ProduceResponse) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ProduceResponse) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProduceResponse) GetValue() ProduceResponseData {
	if o == nil || o.Value.Get() == nil {
		var ret ProduceResponseData
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProduceResponse) GetValueOk() (*ProduceResponseData, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProduceResponse) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableProduceResponseData and assigns it to the Value field.
func (o *ProduceResponse) SetValue(v ProduceResponseData) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ProduceResponse) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProduceResponse) UnsetValue() {
	o.Value.Unset()
}

// Redact resets all sensitive fields to their zero value.
func (o *ProduceResponse) Redact() {
    o.recurseRedact(&o.ClusterId)
    o.recurseRedact(&o.TopicName)
    o.recurseRedact(&o.PartitionId)
    o.recurseRedact(&o.Offset)
    o.recurseRedact(o.Timestamp)
    o.recurseRedact(o.Key)
    o.recurseRedact(o.Value)
}

func (o *ProduceResponse) recurseRedact(v interface{}) {
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

func (o ProduceResponse) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ProduceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["topic_name"] = o.TopicName
	}
	if true {
		toSerialize["partition_id"] = o.PartitionId
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProduceResponse struct {
	value *ProduceResponse
	isSet bool
}

func (v NullableProduceResponse) Get() *ProduceResponse {
	return v.value
}

func (v *NullableProduceResponse) Set(val *ProduceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProduceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProduceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduceResponse(val *ProduceResponse) *NullableProduceResponse {
	return &NullableProduceResponse{value: val, isSet: true}
}

func (v NullableProduceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


