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
)

import (
	"reflect"
)

// PartitionData struct for PartitionData
type PartitionData struct {
	Kind string `json:"kind"`
	Metadata ResourceMetadata `json:"metadata"`
	ClusterId string `json:"cluster_id"`
	TopicName string `json:"topic_name"`
	PartitionId int32 `json:"partition_id"`
	Leader *Relationship `json:"leader,omitempty"`
	Replicas Relationship `json:"replicas"`
	Reassignment Relationship `json:"reassignment"`
}

// NewPartitionData instantiates a new PartitionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartitionData(kind string, metadata ResourceMetadata, clusterId string, topicName string, partitionId int32, replicas Relationship, reassignment Relationship) *PartitionData {
	this := PartitionData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.TopicName = topicName
	this.PartitionId = partitionId
	this.Replicas = replicas
	this.Reassignment = reassignment
	return &this
}

// NewPartitionDataWithDefaults instantiates a new PartitionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionDataWithDefaults() *PartitionData {
	this := PartitionData{}
	return &this
}

// GetKind returns the Kind field value
func (o *PartitionData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *PartitionData) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *PartitionData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *PartitionData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *PartitionData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *PartitionData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *PartitionData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *PartitionData) GetClusterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *PartitionData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetTopicName returns the TopicName field value
func (o *PartitionData) GetTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value
// and a boolean to check if the value has been set.
func (o *PartitionData) GetTopicNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TopicName, true
}

// SetTopicName sets field value
func (o *PartitionData) SetTopicName(v string) {
	o.TopicName = v
}

// GetPartitionId returns the PartitionId field value
func (o *PartitionData) GetPartitionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value
// and a boolean to check if the value has been set.
func (o *PartitionData) GetPartitionIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PartitionId, true
}

// SetPartitionId sets field value
func (o *PartitionData) SetPartitionId(v int32) {
	o.PartitionId = v
}

// GetLeader returns the Leader field value if set, zero value otherwise.
func (o *PartitionData) GetLeader() Relationship {
	if o == nil || o.Leader == nil {
		var ret Relationship
		return ret
	}
	return *o.Leader
}

// GetLeaderOk returns a tuple with the Leader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionData) GetLeaderOk() (*Relationship, bool) {
	if o == nil || o.Leader == nil {
		return nil, false
	}
	return o.Leader, true
}

// HasLeader returns a boolean if a field has been set.
func (o *PartitionData) HasLeader() bool {
	if o != nil && o.Leader != nil {
		return true
	}

	return false
}

// SetLeader gets a reference to the given Relationship and assigns it to the Leader field.
func (o *PartitionData) SetLeader(v Relationship) {
	o.Leader = &v
}

// GetReplicas returns the Replicas field value
func (o *PartitionData) GetReplicas() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *PartitionData) GetReplicasOk() (*Relationship, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value
func (o *PartitionData) SetReplicas(v Relationship) {
	o.Replicas = v
}

// GetReassignment returns the Reassignment field value
func (o *PartitionData) GetReassignment() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Reassignment
}

// GetReassignmentOk returns a tuple with the Reassignment field value
// and a boolean to check if the value has been set.
func (o *PartitionData) GetReassignmentOk() (*Relationship, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reassignment, true
}

// SetReassignment sets field value
func (o *PartitionData) SetReassignment(v Relationship) {
	o.Reassignment = v
}

// Redact resets all sensitive fields to their zero value.
func (o *PartitionData) Redact() {
    o.recurseRedact(&o.Kind)
    o.recurseRedact(&o.Metadata)
    o.recurseRedact(&o.ClusterId)
    o.recurseRedact(&o.TopicName)
    o.recurseRedact(&o.PartitionId)
    o.recurseRedact(o.Leader)
    o.recurseRedact(&o.Replicas)
    o.recurseRedact(&o.Reassignment)
}

func (o *PartitionData) recurseRedact(v interface{}) {
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

func (o PartitionData) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o PartitionData) MarshalJSON() ([]byte, error) {
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
		toSerialize["topic_name"] = o.TopicName
	}
	if true {
		toSerialize["partition_id"] = o.PartitionId
	}
	if o.Leader != nil {
		toSerialize["leader"] = o.Leader
	}
	if true {
		toSerialize["replicas"] = o.Replicas
	}
	if true {
		toSerialize["reassignment"] = o.Reassignment
	}
	return json.Marshal(toSerialize)
}

type NullablePartitionData struct {
	value *PartitionData
	isSet bool
}

func (v NullablePartitionData) Get() *PartitionData {
	return v.value
}

func (v *NullablePartitionData) Set(val *PartitionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitionData(val *PartitionData) *NullablePartitionData {
	return &NullablePartitionData{value: val, isSet: true}
}

func (v NullablePartitionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


