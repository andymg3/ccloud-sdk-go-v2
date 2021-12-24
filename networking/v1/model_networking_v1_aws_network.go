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

// NetworkingV1AwsNetwork The AWS network details.
type NetworkingV1AwsNetwork struct {
	Kind string `json:"kind"`
	// The AWS VPC id for the network.
	Vpc string `json:"vpc"`
	// The AWS account id for the network.
	Account string `json:"account"`
	// The AWS VPC endpoint service for the network (used for PrivateLink) if available.
	PrivateLinkEndpointService *string `json:"private_link_endpoint_service,omitempty"`
}

// NewNetworkingV1AwsNetwork instantiates a new NetworkingV1AwsNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1AwsNetwork(kind string, vpc string, account string) *NetworkingV1AwsNetwork {
	this := NetworkingV1AwsNetwork{}
	this.Kind = kind
	this.Vpc = vpc
	this.Account = account
	return &this
}

// NewNetworkingV1AwsNetworkWithDefaults instantiates a new NetworkingV1AwsNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1AwsNetworkWithDefaults() *NetworkingV1AwsNetwork {
	this := NetworkingV1AwsNetwork{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1AwsNetwork) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsNetwork) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1AwsNetwork) SetKind(v string) {
	o.Kind = v
}

// GetVpc returns the Vpc field value
func (o *NetworkingV1AwsNetwork) GetVpc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vpc
}

// GetVpcOk returns a tuple with the Vpc field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsNetwork) GetVpcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vpc, true
}

// SetVpc sets field value
func (o *NetworkingV1AwsNetwork) SetVpc(v string) {
	o.Vpc = v
}

// GetAccount returns the Account field value
func (o *NetworkingV1AwsNetwork) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsNetwork) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NetworkingV1AwsNetwork) SetAccount(v string) {
	o.Account = v
}

// GetPrivateLinkEndpointService returns the PrivateLinkEndpointService field value if set, zero value otherwise.
func (o *NetworkingV1AwsNetwork) GetPrivateLinkEndpointService() string {
	if o == nil || o.PrivateLinkEndpointService == nil {
		var ret string
		return ret
	}
	return *o.PrivateLinkEndpointService
}

// GetPrivateLinkEndpointServiceOk returns a tuple with the PrivateLinkEndpointService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsNetwork) GetPrivateLinkEndpointServiceOk() (*string, bool) {
	if o == nil || o.PrivateLinkEndpointService == nil {
		return nil, false
	}
	return o.PrivateLinkEndpointService, true
}

// HasPrivateLinkEndpointService returns a boolean if a field has been set.
func (o *NetworkingV1AwsNetwork) HasPrivateLinkEndpointService() bool {
	if o != nil && o.PrivateLinkEndpointService != nil {
		return true
	}

	return false
}

// SetPrivateLinkEndpointService gets a reference to the given string and assigns it to the PrivateLinkEndpointService field.
func (o *NetworkingV1AwsNetwork) SetPrivateLinkEndpointService(v string) {
	o.PrivateLinkEndpointService = &v
}

func (o NetworkingV1AwsNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["vpc"] = o.Vpc
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if o.PrivateLinkEndpointService != nil {
		toSerialize["private_link_endpoint_service"] = o.PrivateLinkEndpointService
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1AwsNetwork struct {
	value *NetworkingV1AwsNetwork
	isSet bool
}

func (v NullableNetworkingV1AwsNetwork) Get() *NetworkingV1AwsNetwork {
	return v.value
}

func (v *NullableNetworkingV1AwsNetwork) Set(val *NetworkingV1AwsNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AwsNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AwsNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AwsNetwork(val *NetworkingV1AwsNetwork) *NullableNetworkingV1AwsNetwork {
	return &NullableNetworkingV1AwsNetwork{value: val, isSet: true}
}

func (v NullableNetworkingV1AwsNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1AwsNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


