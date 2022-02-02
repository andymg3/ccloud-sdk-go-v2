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

// ListLinksResponseDataAllOf struct for ListLinksResponseDataAllOf
type ListLinksResponseDataAllOf struct {
	SourceClusterId string `json:"source_cluster_id"`
	LinkName string `json:"link_name"`
	LinkId string `json:"link_id"`
	TopicsNames *[]string `json:"topics_names,omitempty"`
}

// NewListLinksResponseDataAllOf instantiates a new ListLinksResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLinksResponseDataAllOf(sourceClusterId string, linkName string, linkId string) *ListLinksResponseDataAllOf {
	this := ListLinksResponseDataAllOf{}
	this.SourceClusterId = sourceClusterId
	this.LinkName = linkName
	this.LinkId = linkId
	return &this
}

// NewListLinksResponseDataAllOfWithDefaults instantiates a new ListLinksResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLinksResponseDataAllOfWithDefaults() *ListLinksResponseDataAllOf {
	this := ListLinksResponseDataAllOf{}
	return &this
}

// GetSourceClusterId returns the SourceClusterId field value
func (o *ListLinksResponseDataAllOf) GetSourceClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceClusterId
}

// GetSourceClusterIdOk returns a tuple with the SourceClusterId field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseDataAllOf) GetSourceClusterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceClusterId, true
}

// SetSourceClusterId sets field value
func (o *ListLinksResponseDataAllOf) SetSourceClusterId(v string) {
	o.SourceClusterId = v
}

// GetLinkName returns the LinkName field value
func (o *ListLinksResponseDataAllOf) GetLinkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkName
}

// GetLinkNameOk returns a tuple with the LinkName field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseDataAllOf) GetLinkNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkName, true
}

// SetLinkName sets field value
func (o *ListLinksResponseDataAllOf) SetLinkName(v string) {
	o.LinkName = v
}

// GetLinkId returns the LinkId field value
func (o *ListLinksResponseDataAllOf) GetLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseDataAllOf) GetLinkIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkId, true
}

// SetLinkId sets field value
func (o *ListLinksResponseDataAllOf) SetLinkId(v string) {
	o.LinkId = v
}

// GetTopicsNames returns the TopicsNames field value if set, zero value otherwise.
func (o *ListLinksResponseDataAllOf) GetTopicsNames() []string {
	if o == nil || o.TopicsNames == nil {
		var ret []string
		return ret
	}
	return *o.TopicsNames
}

// GetTopicsNamesOk returns a tuple with the TopicsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLinksResponseDataAllOf) GetTopicsNamesOk() (*[]string, bool) {
	if o == nil || o.TopicsNames == nil {
		return nil, false
	}
	return o.TopicsNames, true
}

// HasTopicsNames returns a boolean if a field has been set.
func (o *ListLinksResponseDataAllOf) HasTopicsNames() bool {
	if o != nil && o.TopicsNames != nil {
		return true
	}

	return false
}

// SetTopicsNames gets a reference to the given []string and assigns it to the TopicsNames field.
func (o *ListLinksResponseDataAllOf) SetTopicsNames(v []string) {
	o.TopicsNames = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ListLinksResponseDataAllOf) Redact() {
    o.recurseRedact(&o.SourceClusterId)
    o.recurseRedact(&o.LinkName)
    o.recurseRedact(&o.LinkId)
    o.recurseRedact(o.TopicsNames)
}

func (o *ListLinksResponseDataAllOf) recurseRedact(v interface{}) {
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

func (o ListLinksResponseDataAllOf) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ListLinksResponseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source_cluster_id"] = o.SourceClusterId
	}
	if true {
		toSerialize["link_name"] = o.LinkName
	}
	if true {
		toSerialize["link_id"] = o.LinkId
	}
	if o.TopicsNames != nil {
		toSerialize["topics_names"] = o.TopicsNames
	}
	return json.Marshal(toSerialize)
}

type NullableListLinksResponseDataAllOf struct {
	value *ListLinksResponseDataAllOf
	isSet bool
}

func (v NullableListLinksResponseDataAllOf) Get() *ListLinksResponseDataAllOf {
	return v.value
}

func (v *NullableListLinksResponseDataAllOf) Set(val *ListLinksResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListLinksResponseDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListLinksResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLinksResponseDataAllOf(val *ListLinksResponseDataAllOf) *NullableListLinksResponseDataAllOf {
	return &NullableListLinksResponseDataAllOf{value: val, isSet: true}
}

func (v NullableListLinksResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLinksResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


