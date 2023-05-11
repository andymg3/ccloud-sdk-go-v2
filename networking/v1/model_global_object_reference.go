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

// GlobalObjectReference ObjectReference provides information for you to locate the referred object
type GlobalObjectReference struct {
	// ID of the referred resource
	Id string `json:"id,omitempty"`
	// API URL for accessing or modifying the referred object
	Related string `json:"related,omitempty"`
	// CRN reference to the referred resource
	ResourceName string `json:"resource_name,omitempty"`
}

// NewGlobalObjectReference instantiates a new GlobalObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalObjectReference(id string, related string, resourceName string) *GlobalObjectReference {
	this := GlobalObjectReference{}
	this.Id = id
	this.Related = related
	this.ResourceName = resourceName
	return &this
}

// NewGlobalObjectReferenceWithDefaults instantiates a new GlobalObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalObjectReferenceWithDefaults() *GlobalObjectReference {
	this := GlobalObjectReference{}
	return &this
}

// GetId returns the Id field value
func (o *GlobalObjectReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GlobalObjectReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GlobalObjectReference) SetId(v string) {
	o.Id = v
}

// GetRelated returns the Related field value
func (o *GlobalObjectReference) GetRelated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *GlobalObjectReference) GetRelatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *GlobalObjectReference) SetRelated(v string) {
	o.Related = v
}

// GetResourceName returns the ResourceName field value
func (o *GlobalObjectReference) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *GlobalObjectReference) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *GlobalObjectReference) SetResourceName(v string) {
	o.ResourceName = v
}

// Redact resets all sensitive fields to their zero value.
func (o *GlobalObjectReference) Redact() {
	o.recurseRedact(&o.Id)
	o.recurseRedact(&o.Related)
	o.recurseRedact(&o.ResourceName)
}

func (o *GlobalObjectReference) recurseRedact(v interface{}) {
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

func (o GlobalObjectReference) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o GlobalObjectReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["related"] = o.Related
	}
	if true {
		toSerialize["resource_name"] = o.ResourceName
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableGlobalObjectReference struct {
	value *GlobalObjectReference
	isSet bool
}

func (v NullableGlobalObjectReference) Get() *GlobalObjectReference {
	return v.value
}

func (v *NullableGlobalObjectReference) Set(val *GlobalObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalObjectReference(val *GlobalObjectReference) *NullableGlobalObjectReference {
	return &NullableGlobalObjectReference{value: val, isSet: true}
}

func (v NullableGlobalObjectReference) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableGlobalObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
