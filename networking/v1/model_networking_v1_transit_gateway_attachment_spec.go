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

// NetworkingV1TransitGatewayAttachmentSpec The desired state of the Transit Gateway Attachment
type NetworkingV1TransitGatewayAttachmentSpec struct {
	// The name of the TGW attachment
	DisplayName *string `json:"display_name,omitempty"`
	// The full AWS Resource Name (ARN) for the AWS Resource Access Manager (RAM) Share of the Transit Gateways that you want Confluent Cloud attached to
	RamShareArn *string `json:"ram_share_arn,omitempty"`
	// The environment to which this belongs.
	Environment *ObjectReference `json:"environment,omitempty"`
	// The network to which this belongs.
	Network *ObjectReference `json:"network,omitempty"`
}

// NewNetworkingV1TransitGatewayAttachmentSpec instantiates a new NetworkingV1TransitGatewayAttachmentSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1TransitGatewayAttachmentSpec() *NetworkingV1TransitGatewayAttachmentSpec {
	this := NetworkingV1TransitGatewayAttachmentSpec{}
	return &this
}

// NewNetworkingV1TransitGatewayAttachmentSpecWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1TransitGatewayAttachmentSpecWithDefaults() *NetworkingV1TransitGatewayAttachmentSpec {
	this := NetworkingV1TransitGatewayAttachmentSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1TransitGatewayAttachmentSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1TransitGatewayAttachmentSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetRamShareArn returns the RamShareArn field value if set, zero value otherwise.
func (o *NetworkingV1TransitGatewayAttachmentSpec) GetRamShareArn() string {
	if o == nil || o.RamShareArn == nil {
		var ret string
		return ret
	}
	return *o.RamShareArn
}

// GetRamShareArnOk returns a tuple with the RamShareArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpec) GetRamShareArnOk() (*string, bool) {
	if o == nil || o.RamShareArn == nil {
		return nil, false
	}
	return o.RamShareArn, true
}

// HasRamShareArn returns a boolean if a field has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpec) HasRamShareArn() bool {
	if o != nil && o.RamShareArn != nil {
		return true
	}

	return false
}

// SetRamShareArn gets a reference to the given string and assigns it to the RamShareArn field.
func (o *NetworkingV1TransitGatewayAttachmentSpec) SetRamShareArn(v string) {
	o.RamShareArn = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1TransitGatewayAttachmentSpec) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpec) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *NetworkingV1TransitGatewayAttachmentSpec) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworkingV1TransitGatewayAttachmentSpec) GetNetwork() ObjectReference {
	if o == nil || o.Network == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpec) GetNetworkOk() (*ObjectReference, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpec) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ObjectReference and assigns it to the Network field.
func (o *NetworkingV1TransitGatewayAttachmentSpec) SetNetwork(v ObjectReference) {
	o.Network = &v
}

func (o NetworkingV1TransitGatewayAttachmentSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.RamShareArn != nil {
		toSerialize["ram_share_arn"] = o.RamShareArn
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1TransitGatewayAttachmentSpec struct {
	value *NetworkingV1TransitGatewayAttachmentSpec
	isSet bool
}

func (v NullableNetworkingV1TransitGatewayAttachmentSpec) Get() *NetworkingV1TransitGatewayAttachmentSpec {
	return v.value
}

func (v *NullableNetworkingV1TransitGatewayAttachmentSpec) Set(val *NetworkingV1TransitGatewayAttachmentSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1TransitGatewayAttachmentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1TransitGatewayAttachmentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1TransitGatewayAttachmentSpec(val *NetworkingV1TransitGatewayAttachmentSpec) *NullableNetworkingV1TransitGatewayAttachmentSpec {
	return &NullableNetworkingV1TransitGatewayAttachmentSpec{value: val, isSet: true}
}

func (v NullableNetworkingV1TransitGatewayAttachmentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1TransitGatewayAttachmentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


