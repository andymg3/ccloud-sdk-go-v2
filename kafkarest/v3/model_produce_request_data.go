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

// ProduceRequestData struct for ProduceRequestData
type ProduceRequestData struct {
	Type *string `json:"type,omitempty"`
	Subject NullableString `json:"subject,omitempty"`
	SubjectNameStrategy NullableString `json:"subject_name_strategy,omitempty"`
	SchemaId NullableInt32 `json:"schema_id,omitempty"`
	SchemaVersion NullableInt32 `json:"schema_version,omitempty"`
	RawSchema NullableString `json:"raw_schema,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// NewProduceRequestData instantiates a new ProduceRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduceRequestData() *ProduceRequestData {
	this := ProduceRequestData{}
	return &this
}

// NewProduceRequestDataWithDefaults instantiates a new ProduceRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduceRequestDataWithDefaults() *ProduceRequestData {
	this := ProduceRequestData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProduceRequestData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProduceRequestData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProduceRequestData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProduceRequestData) SetType(v string) {
	o.Type = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProduceRequestData) GetSubject() string {
	if o == nil || o.Subject.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProduceRequestData) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *ProduceRequestData) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *ProduceRequestData) SetSubject(v string) {
	o.Subject.Set(&v)
}
// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *ProduceRequestData) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *ProduceRequestData) UnsetSubject() {
	o.Subject.Unset()
}

// GetSubjectNameStrategy returns the SubjectNameStrategy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProduceRequestData) GetSubjectNameStrategy() string {
	if o == nil || o.SubjectNameStrategy.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubjectNameStrategy.Get()
}

// GetSubjectNameStrategyOk returns a tuple with the SubjectNameStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProduceRequestData) GetSubjectNameStrategyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubjectNameStrategy.Get(), o.SubjectNameStrategy.IsSet()
}

// HasSubjectNameStrategy returns a boolean if a field has been set.
func (o *ProduceRequestData) HasSubjectNameStrategy() bool {
	if o != nil && o.SubjectNameStrategy.IsSet() {
		return true
	}

	return false
}

// SetSubjectNameStrategy gets a reference to the given NullableString and assigns it to the SubjectNameStrategy field.
func (o *ProduceRequestData) SetSubjectNameStrategy(v string) {
	o.SubjectNameStrategy.Set(&v)
}
// SetSubjectNameStrategyNil sets the value for SubjectNameStrategy to be an explicit nil
func (o *ProduceRequestData) SetSubjectNameStrategyNil() {
	o.SubjectNameStrategy.Set(nil)
}

// UnsetSubjectNameStrategy ensures that no value is present for SubjectNameStrategy, not even an explicit nil
func (o *ProduceRequestData) UnsetSubjectNameStrategy() {
	o.SubjectNameStrategy.Unset()
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProduceRequestData) GetSchemaId() int32 {
	if o == nil || o.SchemaId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SchemaId.Get()
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProduceRequestData) GetSchemaIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SchemaId.Get(), o.SchemaId.IsSet()
}

// HasSchemaId returns a boolean if a field has been set.
func (o *ProduceRequestData) HasSchemaId() bool {
	if o != nil && o.SchemaId.IsSet() {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given NullableInt32 and assigns it to the SchemaId field.
func (o *ProduceRequestData) SetSchemaId(v int32) {
	o.SchemaId.Set(&v)
}
// SetSchemaIdNil sets the value for SchemaId to be an explicit nil
func (o *ProduceRequestData) SetSchemaIdNil() {
	o.SchemaId.Set(nil)
}

// UnsetSchemaId ensures that no value is present for SchemaId, not even an explicit nil
func (o *ProduceRequestData) UnsetSchemaId() {
	o.SchemaId.Unset()
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProduceRequestData) GetSchemaVersion() int32 {
	if o == nil || o.SchemaVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SchemaVersion.Get()
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProduceRequestData) GetSchemaVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SchemaVersion.Get(), o.SchemaVersion.IsSet()
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *ProduceRequestData) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion.IsSet() {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given NullableInt32 and assigns it to the SchemaVersion field.
func (o *ProduceRequestData) SetSchemaVersion(v int32) {
	o.SchemaVersion.Set(&v)
}
// SetSchemaVersionNil sets the value for SchemaVersion to be an explicit nil
func (o *ProduceRequestData) SetSchemaVersionNil() {
	o.SchemaVersion.Set(nil)
}

// UnsetSchemaVersion ensures that no value is present for SchemaVersion, not even an explicit nil
func (o *ProduceRequestData) UnsetSchemaVersion() {
	o.SchemaVersion.Unset()
}

// GetRawSchema returns the RawSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProduceRequestData) GetRawSchema() string {
	if o == nil || o.RawSchema.Get() == nil {
		var ret string
		return ret
	}
	return *o.RawSchema.Get()
}

// GetRawSchemaOk returns a tuple with the RawSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProduceRequestData) GetRawSchemaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RawSchema.Get(), o.RawSchema.IsSet()
}

// HasRawSchema returns a boolean if a field has been set.
func (o *ProduceRequestData) HasRawSchema() bool {
	if o != nil && o.RawSchema.IsSet() {
		return true
	}

	return false
}

// SetRawSchema gets a reference to the given NullableString and assigns it to the RawSchema field.
func (o *ProduceRequestData) SetRawSchema(v string) {
	o.RawSchema.Set(&v)
}
// SetRawSchemaNil sets the value for RawSchema to be an explicit nil
func (o *ProduceRequestData) SetRawSchemaNil() {
	o.RawSchema.Set(nil)
}

// UnsetRawSchema ensures that no value is present for RawSchema, not even an explicit nil
func (o *ProduceRequestData) UnsetRawSchema() {
	o.RawSchema.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProduceRequestData) GetData() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProduceRequestData) GetDataOk() (*interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProduceRequestData) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *ProduceRequestData) SetData(v interface{}) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ProduceRequestData) Redact() {
    o.recurseRedact(o.Type)
    o.recurseRedact(o.Subject)
    o.recurseRedact(o.SubjectNameStrategy)
    o.recurseRedact(o.SchemaId)
    o.recurseRedact(o.SchemaVersion)
    o.recurseRedact(o.RawSchema)
    o.recurseRedact(o.Data)
}

func (o *ProduceRequestData) recurseRedact(v interface{}) {
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

func (o ProduceRequestData) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ProduceRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.SubjectNameStrategy.IsSet() {
		toSerialize["subject_name_strategy"] = o.SubjectNameStrategy.Get()
	}
	if o.SchemaId.IsSet() {
		toSerialize["schema_id"] = o.SchemaId.Get()
	}
	if o.SchemaVersion.IsSet() {
		toSerialize["schema_version"] = o.SchemaVersion.Get()
	}
	if o.RawSchema.IsSet() {
		toSerialize["raw_schema"] = o.RawSchema.Get()
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableProduceRequestData struct {
	value *ProduceRequestData
	isSet bool
}

func (v NullableProduceRequestData) Get() *ProduceRequestData {
	return v.value
}

func (v *NullableProduceRequestData) Set(val *ProduceRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableProduceRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableProduceRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduceRequestData(val *ProduceRequestData) *NullableProduceRequestData {
	return &NullableProduceRequestData{value: val, isSet: true}
}

func (v NullableProduceRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduceRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


