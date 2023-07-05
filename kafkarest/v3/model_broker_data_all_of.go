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

// BrokerDataAllOf struct for BrokerDataAllOf
type BrokerDataAllOf struct {
	ClusterId         string         `json:"cluster_id,omitempty"`
	BrokerId          int32          `json:"broker_id,omitempty"`
	Host              NullableString `json:"host,omitempty"`
	Port              NullableInt32  `json:"port,omitempty"`
	Rack              NullableString `json:"rack,omitempty"`
	Configs           Relationship   `json:"configs,omitempty"`
	PartitionReplicas Relationship   `json:"partition_replicas,omitempty"`
}

// NewBrokerDataAllOf instantiates a new BrokerDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerDataAllOf(clusterId string, brokerId int32, configs Relationship, partitionReplicas Relationship) *BrokerDataAllOf {
	this := BrokerDataAllOf{}
	this.ClusterId = clusterId
	this.BrokerId = brokerId
	this.Configs = configs
	this.PartitionReplicas = partitionReplicas
	return &this
}

// NewBrokerDataAllOfWithDefaults instantiates a new BrokerDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerDataAllOfWithDefaults() *BrokerDataAllOf {
	this := BrokerDataAllOf{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *BrokerDataAllOf) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *BrokerDataAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *BrokerDataAllOf) SetClusterId(v string) {
	o.ClusterId = v
}

// GetBrokerId returns the BrokerId field value
func (o *BrokerDataAllOf) GetBrokerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value
// and a boolean to check if the value has been set.
func (o *BrokerDataAllOf) GetBrokerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerId, true
}

// SetBrokerId sets field value
func (o *BrokerDataAllOf) SetBrokerId(v int32) {
	o.BrokerId = v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrokerDataAllOf) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrokerDataAllOf) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *BrokerDataAllOf) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *BrokerDataAllOf) SetHost(v string) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *BrokerDataAllOf) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *BrokerDataAllOf) UnsetHost() {
	o.Host.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrokerDataAllOf) GetPort() int32 {
	if o == nil || o.Port.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrokerDataAllOf) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *BrokerDataAllOf) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *BrokerDataAllOf) SetPort(v int32) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *BrokerDataAllOf) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *BrokerDataAllOf) UnsetPort() {
	o.Port.Unset()
}

// GetRack returns the Rack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrokerDataAllOf) GetRack() string {
	if o == nil || o.Rack.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rack.Get()
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrokerDataAllOf) GetRackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rack.Get(), o.Rack.IsSet()
}

// HasRack returns a boolean if a field has been set.
func (o *BrokerDataAllOf) HasRack() bool {
	if o != nil && o.Rack.IsSet() {
		return true
	}

	return false
}

// SetRack gets a reference to the given NullableString and assigns it to the Rack field.
func (o *BrokerDataAllOf) SetRack(v string) {
	o.Rack.Set(&v)
}

// SetRackNil sets the value for Rack to be an explicit nil
func (o *BrokerDataAllOf) SetRackNil() {
	o.Rack.Set(nil)
}

// UnsetRack ensures that no value is present for Rack, not even an explicit nil
func (o *BrokerDataAllOf) UnsetRack() {
	o.Rack.Unset()
}

// GetConfigs returns the Configs field value
func (o *BrokerDataAllOf) GetConfigs() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value
// and a boolean to check if the value has been set.
func (o *BrokerDataAllOf) GetConfigsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configs, true
}

// SetConfigs sets field value
func (o *BrokerDataAllOf) SetConfigs(v Relationship) {
	o.Configs = v
}

// GetPartitionReplicas returns the PartitionReplicas field value
func (o *BrokerDataAllOf) GetPartitionReplicas() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.PartitionReplicas
}

// GetPartitionReplicasOk returns a tuple with the PartitionReplicas field value
// and a boolean to check if the value has been set.
func (o *BrokerDataAllOf) GetPartitionReplicasOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionReplicas, true
}

// SetPartitionReplicas sets field value
func (o *BrokerDataAllOf) SetPartitionReplicas(v Relationship) {
	o.PartitionReplicas = v
}

// Redact resets all sensitive fields to their zero value.
func (o *BrokerDataAllOf) Redact() {
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.BrokerId)
	o.recurseRedact(o.Host)
	o.recurseRedact(o.Port)
	o.recurseRedact(o.Rack)
	o.recurseRedact(&o.Configs)
	o.recurseRedact(&o.PartitionReplicas)
}

func (o *BrokerDataAllOf) recurseRedact(v interface{}) {
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

func (o BrokerDataAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o BrokerDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["broker_id"] = o.BrokerId
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Rack.IsSet() {
		toSerialize["rack"] = o.Rack.Get()
	}
	if true {
		toSerialize["configs"] = o.Configs
	}
	if true {
		toSerialize["partition_replicas"] = o.PartitionReplicas
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableBrokerDataAllOf struct {
	value *BrokerDataAllOf
	isSet bool
}

func (v NullableBrokerDataAllOf) Get() *BrokerDataAllOf {
	return v.value
}

func (v *NullableBrokerDataAllOf) Set(val *BrokerDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerDataAllOf(val *BrokerDataAllOf) *NullableBrokerDataAllOf {
	return &NullableBrokerDataAllOf{value: val, isSet: true}
}

func (v NullableBrokerDataAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableBrokerDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
