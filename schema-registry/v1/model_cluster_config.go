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
Confluent Schema Registry APIs

REST API for the Schema Registry

API version: 1.0.0
Contact: data-governance@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// ClusterConfig Cluster Config
type ClusterConfig struct {
	// Maximum number of registered schemas allowed
	MaxSchemas *int32 `json:"maxSchemas,omitempty"`
	// Maximum number of allowed requests per second
	MaxRequestsPerSec *int32 `json:"maxRequestsPerSec,omitempty"`
}

// NewClusterConfig instantiates a new ClusterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfig() *ClusterConfig {
	this := ClusterConfig{}
	return &this
}

// NewClusterConfigWithDefaults instantiates a new ClusterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigWithDefaults() *ClusterConfig {
	this := ClusterConfig{}
	return &this
}

// GetMaxSchemas returns the MaxSchemas field value if set, zero value otherwise.
func (o *ClusterConfig) GetMaxSchemas() int32 {
	if o == nil || o.MaxSchemas == nil {
		var ret int32
		return ret
	}
	return *o.MaxSchemas
}

// GetMaxSchemasOk returns a tuple with the MaxSchemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfig) GetMaxSchemasOk() (*int32, bool) {
	if o == nil || o.MaxSchemas == nil {
		return nil, false
	}
	return o.MaxSchemas, true
}

// HasMaxSchemas returns a boolean if a field has been set.
func (o *ClusterConfig) HasMaxSchemas() bool {
	if o != nil && o.MaxSchemas != nil {
		return true
	}

	return false
}

// SetMaxSchemas gets a reference to the given int32 and assigns it to the MaxSchemas field.
func (o *ClusterConfig) SetMaxSchemas(v int32) {
	o.MaxSchemas = &v
}

// GetMaxRequestsPerSec returns the MaxRequestsPerSec field value if set, zero value otherwise.
func (o *ClusterConfig) GetMaxRequestsPerSec() int32 {
	if o == nil || o.MaxRequestsPerSec == nil {
		var ret int32
		return ret
	}
	return *o.MaxRequestsPerSec
}

// GetMaxRequestsPerSecOk returns a tuple with the MaxRequestsPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfig) GetMaxRequestsPerSecOk() (*int32, bool) {
	if o == nil || o.MaxRequestsPerSec == nil {
		return nil, false
	}
	return o.MaxRequestsPerSec, true
}

// HasMaxRequestsPerSec returns a boolean if a field has been set.
func (o *ClusterConfig) HasMaxRequestsPerSec() bool {
	if o != nil && o.MaxRequestsPerSec != nil {
		return true
	}

	return false
}

// SetMaxRequestsPerSec gets a reference to the given int32 and assigns it to the MaxRequestsPerSec field.
func (o *ClusterConfig) SetMaxRequestsPerSec(v int32) {
	o.MaxRequestsPerSec = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ClusterConfig) Redact() {
	o.recurseRedact(o.MaxSchemas)
	o.recurseRedact(o.MaxRequestsPerSec)
}

func (o *ClusterConfig) recurseRedact(v interface{}) {
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

func (o ClusterConfig) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ClusterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxSchemas != nil {
		toSerialize["maxSchemas"] = o.MaxSchemas
	}
	if o.MaxRequestsPerSec != nil {
		toSerialize["maxRequestsPerSec"] = o.MaxRequestsPerSec
	}
	return json.Marshal(toSerialize)
}

type NullableClusterConfig struct {
	value *ClusterConfig
	isSet bool
}

func (v NullableClusterConfig) Get() *ClusterConfig {
	return v.value
}

func (v *NullableClusterConfig) Set(val *ClusterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterConfig(val *ClusterConfig) *NullableClusterConfig {
	return &NullableClusterConfig{value: val, isSet: true}
}

func (v NullableClusterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
