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
)

import (
	"reflect"
)

// ReplicaData struct for ReplicaData
type ReplicaData struct {
	Kind        string           `json:"kind,omitempty"`
	Metadata    ResourceMetadata `json:"metadata,omitempty"`
	ClusterId   string           `json:"cluster_id,omitempty"`
	TopicName   string           `json:"topic_name,omitempty"`
	PartitionId int32            `json:"partition_id,omitempty"`
	BrokerId    int32            `json:"broker_id,omitempty"`
	IsLeader    bool             `json:"is_leader,omitempty"`
	IsInSync    bool             `json:"is_in_sync,omitempty"`
	Broker      Relationship     `json:"broker,omitempty"`
}

// NewReplicaData instantiates a new ReplicaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicaData(kind string, metadata ResourceMetadata, clusterId string, topicName string, partitionId int32, brokerId int32, isLeader bool, isInSync bool, broker Relationship) *ReplicaData {
	this := ReplicaData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.TopicName = topicName
	this.PartitionId = partitionId
	this.BrokerId = brokerId
	this.IsLeader = isLeader
	this.IsInSync = isInSync
	this.Broker = broker
	return &this
}

// NewReplicaDataWithDefaults instantiates a new ReplicaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaDataWithDefaults() *ReplicaData {
	this := ReplicaData{}
	return &this
}

// GetKind returns the Kind field value
func (o *ReplicaData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ReplicaData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ReplicaData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ReplicaData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ReplicaData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ReplicaData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *ReplicaData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ReplicaData) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ReplicaData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetTopicName returns the TopicName field value
func (o *ReplicaData) GetTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value
// and a boolean to check if the value has been set.
func (o *ReplicaData) GetTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicName, true
}

// SetTopicName sets field value
func (o *ReplicaData) SetTopicName(v string) {
	o.TopicName = v
}

// GetPartitionId returns the PartitionId field value
func (o *ReplicaData) GetPartitionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value
// and a boolean to check if the value has been set.
func (o *ReplicaData) GetPartitionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionId, true
}

// SetPartitionId sets field value
func (o *ReplicaData) SetPartitionId(v int32) {
	o.PartitionId = v
}

// GetBrokerId returns the BrokerId field value
func (o *ReplicaData) GetBrokerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value
// and a boolean to check if the value has been set.
func (o *ReplicaData) GetBrokerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerId, true
}

// SetBrokerId sets field value
func (o *ReplicaData) SetBrokerId(v int32) {
	o.BrokerId = v
}

// GetIsLeader returns the IsLeader field value
func (o *ReplicaData) GetIsLeader() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLeader
}

// GetIsLeaderOk returns a tuple with the IsLeader field value
// and a boolean to check if the value has been set.
func (o *ReplicaData) GetIsLeaderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLeader, true
}

// SetIsLeader sets field value
func (o *ReplicaData) SetIsLeader(v bool) {
	o.IsLeader = v
}

// GetIsInSync returns the IsInSync field value
func (o *ReplicaData) GetIsInSync() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInSync
}

// GetIsInSyncOk returns a tuple with the IsInSync field value
// and a boolean to check if the value has been set.
func (o *ReplicaData) GetIsInSyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInSync, true
}

// SetIsInSync sets field value
func (o *ReplicaData) SetIsInSync(v bool) {
	o.IsInSync = v
}

// GetBroker returns the Broker field value
func (o *ReplicaData) GetBroker() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Broker
}

// GetBrokerOk returns a tuple with the Broker field value
// and a boolean to check if the value has been set.
func (o *ReplicaData) GetBrokerOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Broker, true
}

// SetBroker sets field value
func (o *ReplicaData) SetBroker(v Relationship) {
	o.Broker = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ReplicaData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.TopicName)
	o.recurseRedact(&o.PartitionId)
	o.recurseRedact(&o.BrokerId)
	o.recurseRedact(&o.IsLeader)
	o.recurseRedact(&o.IsInSync)
	o.recurseRedact(&o.Broker)
}

func (o *ReplicaData) recurseRedact(v interface{}) {
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

func (o ReplicaData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ReplicaData) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["broker_id"] = o.BrokerId
	}
	if true {
		toSerialize["is_leader"] = o.IsLeader
	}
	if true {
		toSerialize["is_in_sync"] = o.IsInSync
	}
	if true {
		toSerialize["broker"] = o.Broker
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableReplicaData struct {
	value *ReplicaData
	isSet bool
}

func (v NullableReplicaData) Get() *ReplicaData {
	return v.value
}

func (v *NullableReplicaData) Set(val *ReplicaData) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaData) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaData(val *ReplicaData) *NullableReplicaData {
	return &NullableReplicaData{value: val, isSet: true}
}

func (v NullableReplicaData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableReplicaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
