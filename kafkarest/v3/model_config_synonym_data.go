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

// ConfigSynonymData struct for ConfigSynonymData
type ConfigSynonymData struct {
	Name string `json:"name"`
	Value NullableString `json:"value,omitempty"`
	Source string `json:"source"`
}

// NewConfigSynonymData instantiates a new ConfigSynonymData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSynonymData(name string, source string) *ConfigSynonymData {
	this := ConfigSynonymData{}
	this.Name = name
	this.Source = source
	return &this
}

// NewConfigSynonymDataWithDefaults instantiates a new ConfigSynonymData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSynonymDataWithDefaults() *ConfigSynonymData {
	this := ConfigSynonymData{}
	return &this
}

// GetName returns the Name field value
func (o *ConfigSynonymData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigSynonymData) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigSynonymData) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSynonymData) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSynonymData) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigSynonymData) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ConfigSynonymData) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ConfigSynonymData) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ConfigSynonymData) UnsetValue() {
	o.Value.Unset()
}

// GetSource returns the Source field value
func (o *ConfigSynonymData) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ConfigSynonymData) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ConfigSynonymData) SetSource(v string) {
	o.Source = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConfigSynonymData) Redact() {
    o.recurseRedact(&o.Name)
    o.recurseRedact(o.Value)
    o.recurseRedact(&o.Source)
}

func (o *ConfigSynonymData) recurseRedact(v interface{}) {
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

func (o ConfigSynonymData) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ConfigSynonymData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableConfigSynonymData struct {
	value *ConfigSynonymData
	isSet bool
}

func (v NullableConfigSynonymData) Get() *ConfigSynonymData {
	return v.value
}

func (v *NullableConfigSynonymData) Set(val *ConfigSynonymData) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSynonymData) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSynonymData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSynonymData(val *ConfigSynonymData) *NullableConfigSynonymData {
	return &NullableConfigSynonymData{value: val, isSet: true}
}

func (v NullableConfigSynonymData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSynonymData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


