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

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// NetworkingV1GcpPrivateServiceConnectAccess GCP Private Service Connect access configuration.
type NetworkingV1GcpPrivateServiceConnectAccess struct {
	// PrivateLink kind type.
	Kind string `json:"kind"`
	// GCP project to allow for Private Service Connect access.
	Project string `json:"project"`
}

// NewNetworkingV1GcpPrivateServiceConnectAccess instantiates a new NetworkingV1GcpPrivateServiceConnectAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1GcpPrivateServiceConnectAccess(kind string, project string) *NetworkingV1GcpPrivateServiceConnectAccess {
	this := NetworkingV1GcpPrivateServiceConnectAccess{}
	this.Kind = kind
	this.Project = project
	return &this
}

// NewNetworkingV1GcpPrivateServiceConnectAccessWithDefaults instantiates a new NetworkingV1GcpPrivateServiceConnectAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1GcpPrivateServiceConnectAccessWithDefaults() *NetworkingV1GcpPrivateServiceConnectAccess {
	this := NetworkingV1GcpPrivateServiceConnectAccess{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1GcpPrivateServiceConnectAccess) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1GcpPrivateServiceConnectAccess) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1GcpPrivateServiceConnectAccess) SetKind(v string) {
	o.Kind = v
}

// GetProject returns the Project field value
func (o *NetworkingV1GcpPrivateServiceConnectAccess) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1GcpPrivateServiceConnectAccess) GetProjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *NetworkingV1GcpPrivateServiceConnectAccess) SetProject(v string) {
	o.Project = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1GcpPrivateServiceConnectAccess) Redact() {
    o.recurseRedact(&o.Kind)
    o.recurseRedact(&o.Project)
}

func (o *NetworkingV1GcpPrivateServiceConnectAccess) recurseRedact(v interface{}) {
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

func (o NetworkingV1GcpPrivateServiceConnectAccess) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1GcpPrivateServiceConnectAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1GcpPrivateServiceConnectAccess struct {
	value *NetworkingV1GcpPrivateServiceConnectAccess
	isSet bool
}

func (v NullableNetworkingV1GcpPrivateServiceConnectAccess) Get() *NetworkingV1GcpPrivateServiceConnectAccess {
	return v.value
}

func (v *NullableNetworkingV1GcpPrivateServiceConnectAccess) Set(val *NetworkingV1GcpPrivateServiceConnectAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1GcpPrivateServiceConnectAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1GcpPrivateServiceConnectAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1GcpPrivateServiceConnectAccess(val *NetworkingV1GcpPrivateServiceConnectAccess) *NullableNetworkingV1GcpPrivateServiceConnectAccess {
	return &NullableNetworkingV1GcpPrivateServiceConnectAccess{value: val, isSet: true}
}

func (v NullableNetworkingV1GcpPrivateServiceConnectAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1GcpPrivateServiceConnectAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


