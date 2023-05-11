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

// NetworkingV1NetworkLinkServiceSpec The desired state of the Network Link Service
type NetworkingV1NetworkLinkServiceSpec struct {
	// The name of the network link service
	DisplayName *string `json:"display_name,omitempty"`
	// The description of the network link service
	Description *string `json:"description,omitempty"`
	// Network Link Service Accept policy
	Accept *NetworkingV1NetworkLinkServiceAcceptPolicy `json:"accept,omitempty"`
	// The environment to which this belongs.
	Environment *GlobalObjectReference `json:"environment,omitempty"`
	// The network to which this belongs.
	Network *EnvScopedObjectReference `json:"network,omitempty"`
}

// NewNetworkingV1NetworkLinkServiceSpec instantiates a new NetworkingV1NetworkLinkServiceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1NetworkLinkServiceSpec() *NetworkingV1NetworkLinkServiceSpec {
	this := NetworkingV1NetworkLinkServiceSpec{}
	return &this
}

// NewNetworkingV1NetworkLinkServiceSpecWithDefaults instantiates a new NetworkingV1NetworkLinkServiceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1NetworkLinkServiceSpecWithDefaults() *NetworkingV1NetworkLinkServiceSpec {
	this := NetworkingV1NetworkLinkServiceSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1NetworkLinkServiceSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkingV1NetworkLinkServiceSpec) SetDescription(v string) {
	o.Description = &v
}

// GetAccept returns the Accept field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceSpec) GetAccept() NetworkingV1NetworkLinkServiceAcceptPolicy {
	if o == nil || o.Accept == nil {
		var ret NetworkingV1NetworkLinkServiceAcceptPolicy
		return ret
	}
	return *o.Accept
}

// GetAcceptOk returns a tuple with the Accept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) GetAcceptOk() (*NetworkingV1NetworkLinkServiceAcceptPolicy, bool) {
	if o == nil || o.Accept == nil {
		return nil, false
	}
	return o.Accept, true
}

// HasAccept returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) HasAccept() bool {
	if o != nil && o.Accept != nil {
		return true
	}

	return false
}

// SetAccept gets a reference to the given NetworkingV1NetworkLinkServiceAcceptPolicy and assigns it to the Accept field.
func (o *NetworkingV1NetworkLinkServiceSpec) SetAccept(v NetworkingV1NetworkLinkServiceAcceptPolicy) {
	o.Accept = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceSpec) GetEnvironment() GlobalObjectReference {
	if o == nil || o.Environment == nil {
		var ret GlobalObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) GetEnvironmentOk() (*GlobalObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given GlobalObjectReference and assigns it to the Environment field.
func (o *NetworkingV1NetworkLinkServiceSpec) SetEnvironment(v GlobalObjectReference) {
	o.Environment = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkServiceSpec) GetNetwork() EnvScopedObjectReference {
	if o == nil || o.Network == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) GetNetworkOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkServiceSpec) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given EnvScopedObjectReference and assigns it to the Network field.
func (o *NetworkingV1NetworkLinkServiceSpec) SetNetwork(v EnvScopedObjectReference) {
	o.Network = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1NetworkLinkServiceSpec) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Description)
	o.recurseRedact(o.Accept)
	o.recurseRedact(o.Environment)
	o.recurseRedact(o.Network)
}

func (o *NetworkingV1NetworkLinkServiceSpec) recurseRedact(v interface{}) {
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

func (o NetworkingV1NetworkLinkServiceSpec) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1NetworkLinkServiceSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Accept != nil {
		toSerialize["accept"] = o.Accept
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1NetworkLinkServiceSpec struct {
	value *NetworkingV1NetworkLinkServiceSpec
	isSet bool
}

func (v NullableNetworkingV1NetworkLinkServiceSpec) Get() *NetworkingV1NetworkLinkServiceSpec {
	return v.value
}

func (v *NullableNetworkingV1NetworkLinkServiceSpec) Set(val *NetworkingV1NetworkLinkServiceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1NetworkLinkServiceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1NetworkLinkServiceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1NetworkLinkServiceSpec(val *NetworkingV1NetworkLinkServiceSpec) *NullableNetworkingV1NetworkLinkServiceSpec {
	return &NullableNetworkingV1NetworkLinkServiceSpec{value: val, isSet: true}
}

func (v NullableNetworkingV1NetworkLinkServiceSpec) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1NetworkLinkServiceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
