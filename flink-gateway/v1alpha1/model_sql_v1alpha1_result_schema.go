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
SQL API v1alpha1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// SqlV1alpha1ResultSchema The table columns of the results schema.
type SqlV1alpha1ResultSchema struct {
	// The properties of each SQL column in the schema.
	Columns *[]ColumnDetails `json:"columns,omitempty"`
}

// NewSqlV1alpha1ResultSchema instantiates a new SqlV1alpha1ResultSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlV1alpha1ResultSchema() *SqlV1alpha1ResultSchema {
	this := SqlV1alpha1ResultSchema{}
	return &this
}

// NewSqlV1alpha1ResultSchemaWithDefaults instantiates a new SqlV1alpha1ResultSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlV1alpha1ResultSchemaWithDefaults() *SqlV1alpha1ResultSchema {
	this := SqlV1alpha1ResultSchema{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *SqlV1alpha1ResultSchema) GetColumns() []ColumnDetails {
	if o == nil || o.Columns == nil {
		var ret []ColumnDetails
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1alpha1ResultSchema) GetColumnsOk() (*[]ColumnDetails, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *SqlV1alpha1ResultSchema) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []ColumnDetails and assigns it to the Columns field.
func (o *SqlV1alpha1ResultSchema) SetColumns(v []ColumnDetails) {
	o.Columns = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *SqlV1alpha1ResultSchema) Redact() {
    o.recurseRedact(o.Columns)
}

func (o *SqlV1alpha1ResultSchema) recurseRedact(v interface{}) {
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

func (o SqlV1alpha1ResultSchema) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o SqlV1alpha1ResultSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableSqlV1alpha1ResultSchema struct {
	value *SqlV1alpha1ResultSchema
	isSet bool
}

func (v NullableSqlV1alpha1ResultSchema) Get() *SqlV1alpha1ResultSchema {
	return v.value
}

func (v *NullableSqlV1alpha1ResultSchema) Set(val *SqlV1alpha1ResultSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSqlV1alpha1ResultSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSqlV1alpha1ResultSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqlV1alpha1ResultSchema(val *SqlV1alpha1ResultSchema) *NullableSqlV1alpha1ResultSchema {
	return &NullableSqlV1alpha1ResultSchema{value: val, isSet: true}
}

func (v NullableSqlV1alpha1ResultSchema) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableSqlV1alpha1ResultSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


