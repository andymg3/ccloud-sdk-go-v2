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

// TopicConfigDataAllOf struct for TopicConfigDataAllOf
type TopicConfigDataAllOf struct {
	TopicName string `json:"topic_name"`
}

// NewTopicConfigDataAllOf instantiates a new TopicConfigDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicConfigDataAllOf(topicName string) *TopicConfigDataAllOf {
	this := TopicConfigDataAllOf{}
	this.TopicName = topicName
	return &this
}

// NewTopicConfigDataAllOfWithDefaults instantiates a new TopicConfigDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicConfigDataAllOfWithDefaults() *TopicConfigDataAllOf {
	this := TopicConfigDataAllOf{}
	return &this
}

// GetTopicName returns the TopicName field value
func (o *TopicConfigDataAllOf) GetTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value
// and a boolean to check if the value has been set.
func (o *TopicConfigDataAllOf) GetTopicNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TopicName, true
}

// SetTopicName sets field value
func (o *TopicConfigDataAllOf) SetTopicName(v string) {
	o.TopicName = v
}

// Redact resets all sensitive fields to their zero value.
func (o *TopicConfigDataAllOf) Redact() {
    o.recurseRedact(&o.TopicName)
}

func (o *TopicConfigDataAllOf) recurseRedact(v interface{}) {
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

func (o TopicConfigDataAllOf) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o TopicConfigDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["topic_name"] = o.TopicName
	}
	return json.Marshal(toSerialize)
}

type NullableTopicConfigDataAllOf struct {
	value *TopicConfigDataAllOf
	isSet bool
}

func (v NullableTopicConfigDataAllOf) Get() *TopicConfigDataAllOf {
	return v.value
}

func (v *NullableTopicConfigDataAllOf) Set(val *TopicConfigDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicConfigDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicConfigDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicConfigDataAllOf(val *TopicConfigDataAllOf) *NullableTopicConfigDataAllOf {
	return &NullableTopicConfigDataAllOf{value: val, isSet: true}
}

func (v NullableTopicConfigDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicConfigDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


