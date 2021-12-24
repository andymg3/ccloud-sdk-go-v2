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

// NetworkingV1AwsPrivateLinkAccess AWS PrivateLink access configuration.
type NetworkingV1AwsPrivateLinkAccess struct {
	Kind string `json:"kind"`
	// AWS account to allow for PrivateLink access.
	Account string `json:"account"`
}

// NewNetworkingV1AwsPrivateLinkAccess instantiates a new NetworkingV1AwsPrivateLinkAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1AwsPrivateLinkAccess(kind string, account string) *NetworkingV1AwsPrivateLinkAccess {
	this := NetworkingV1AwsPrivateLinkAccess{}
	this.Kind = kind
	this.Account = account
	return &this
}

// NewNetworkingV1AwsPrivateLinkAccessWithDefaults instantiates a new NetworkingV1AwsPrivateLinkAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1AwsPrivateLinkAccessWithDefaults() *NetworkingV1AwsPrivateLinkAccess {
	this := NetworkingV1AwsPrivateLinkAccess{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1AwsPrivateLinkAccess) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsPrivateLinkAccess) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1AwsPrivateLinkAccess) SetKind(v string) {
	o.Kind = v
}

// GetAccount returns the Account field value
func (o *NetworkingV1AwsPrivateLinkAccess) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsPrivateLinkAccess) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NetworkingV1AwsPrivateLinkAccess) SetAccount(v string) {
	o.Account = v
}

func (o NetworkingV1AwsPrivateLinkAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1AwsPrivateLinkAccess struct {
	value *NetworkingV1AwsPrivateLinkAccess
	isSet bool
}

func (v NullableNetworkingV1AwsPrivateLinkAccess) Get() *NetworkingV1AwsPrivateLinkAccess {
	return v.value
}

func (v *NullableNetworkingV1AwsPrivateLinkAccess) Set(val *NetworkingV1AwsPrivateLinkAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AwsPrivateLinkAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AwsPrivateLinkAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AwsPrivateLinkAccess(val *NetworkingV1AwsPrivateLinkAccess) *NullableNetworkingV1AwsPrivateLinkAccess {
	return &NullableNetworkingV1AwsPrivateLinkAccess{value: val, isSet: true}
}

func (v NullableNetworkingV1AwsPrivateLinkAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1AwsPrivateLinkAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


