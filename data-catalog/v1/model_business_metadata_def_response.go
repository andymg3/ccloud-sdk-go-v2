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
Confluent Data Catalog

REST API for the Data Catalog

API version: 1.0.0
Contact: data-governance@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// BusinessMetadataDefResponse struct for BusinessMetadataDefResponse
type BusinessMetadataDefResponse struct {
	// The category
	Category *string `json:"category,omitempty"`
	// The internal guid
	Guid *string `json:"guid,omitempty"`
	// The creator
	CreatedBy *string `json:"createdBy,omitempty"`
	// The updater
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// The create time
	CreateTime *int64 `json:"createTime,omitempty"`
	// The update time
	UpdateTime *int64 `json:"updateTime,omitempty"`
	// The version
	Version *int32 `json:"version,omitempty"`
	// The name
	Name *string `json:"name,omitempty"`
	// The description
	Description *string `json:"description,omitempty"`
	// The type version
	TypeVersion *string `json:"typeVersion,omitempty"`
	// The service type
	ServiceType *string `json:"serviceType,omitempty"`
	// The options
	Options *map[string]string `json:"options,omitempty"`
	// The attribute definitions
	AttributeDefs *[]AttributeDef `json:"attributeDefs,omitempty"`
	Error         *ErrorMessage   `json:"error,omitempty"`
}

// NewBusinessMetadataDefResponse instantiates a new BusinessMetadataDefResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessMetadataDefResponse() *BusinessMetadataDefResponse {
	this := BusinessMetadataDefResponse{}
	return &this
}

// NewBusinessMetadataDefResponseWithDefaults instantiates a new BusinessMetadataDefResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessMetadataDefResponseWithDefaults() *BusinessMetadataDefResponse {
	this := BusinessMetadataDefResponse{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *BusinessMetadataDefResponse) SetCategory(v string) {
	o.Category = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *BusinessMetadataDefResponse) SetGuid(v string) {
	o.Guid = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *BusinessMetadataDefResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *BusinessMetadataDefResponse) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetCreateTime() int64 {
	if o == nil || o.CreateTime == nil {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *BusinessMetadataDefResponse) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetUpdateTime() int64 {
	if o == nil || o.UpdateTime == nil {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *BusinessMetadataDefResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BusinessMetadataDefResponse) SetVersion(v int32) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BusinessMetadataDefResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BusinessMetadataDefResponse) SetDescription(v string) {
	o.Description = &v
}

// GetTypeVersion returns the TypeVersion field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetTypeVersion() string {
	if o == nil || o.TypeVersion == nil {
		var ret string
		return ret
	}
	return *o.TypeVersion
}

// GetTypeVersionOk returns a tuple with the TypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetTypeVersionOk() (*string, bool) {
	if o == nil || o.TypeVersion == nil {
		return nil, false
	}
	return o.TypeVersion, true
}

// HasTypeVersion returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasTypeVersion() bool {
	if o != nil && o.TypeVersion != nil {
		return true
	}

	return false
}

// SetTypeVersion gets a reference to the given string and assigns it to the TypeVersion field.
func (o *BusinessMetadataDefResponse) SetTypeVersion(v string) {
	o.TypeVersion = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *BusinessMetadataDefResponse) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetOptions() map[string]string {
	if o == nil || o.Options == nil {
		var ret map[string]string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetOptionsOk() (*map[string]string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]string and assigns it to the Options field.
func (o *BusinessMetadataDefResponse) SetOptions(v map[string]string) {
	o.Options = &v
}

// GetAttributeDefs returns the AttributeDefs field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetAttributeDefs() []AttributeDef {
	if o == nil || o.AttributeDefs == nil {
		var ret []AttributeDef
		return ret
	}
	return *o.AttributeDefs
}

// GetAttributeDefsOk returns a tuple with the AttributeDefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetAttributeDefsOk() (*[]AttributeDef, bool) {
	if o == nil || o.AttributeDefs == nil {
		return nil, false
	}
	return o.AttributeDefs, true
}

// HasAttributeDefs returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasAttributeDefs() bool {
	if o != nil && o.AttributeDefs != nil {
		return true
	}

	return false
}

// SetAttributeDefs gets a reference to the given []AttributeDef and assigns it to the AttributeDefs field.
func (o *BusinessMetadataDefResponse) SetAttributeDefs(v []AttributeDef) {
	o.AttributeDefs = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BusinessMetadataDefResponse) GetError() ErrorMessage {
	if o == nil || o.Error == nil {
		var ret ErrorMessage
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessMetadataDefResponse) GetErrorOk() (*ErrorMessage, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BusinessMetadataDefResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given ErrorMessage and assigns it to the Error field.
func (o *BusinessMetadataDefResponse) SetError(v ErrorMessage) {
	o.Error = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *BusinessMetadataDefResponse) Redact() {
	o.recurseRedact(o.Category)
	o.recurseRedact(o.Guid)
	o.recurseRedact(o.CreatedBy)
	o.recurseRedact(o.UpdatedBy)
	o.recurseRedact(o.CreateTime)
	o.recurseRedact(o.UpdateTime)
	o.recurseRedact(o.Version)
	o.recurseRedact(o.Name)
	o.recurseRedact(o.Description)
	o.recurseRedact(o.TypeVersion)
	o.recurseRedact(o.ServiceType)
	o.recurseRedact(o.Options)
	o.recurseRedact(o.AttributeDefs)
	o.recurseRedact(o.Error)
}

func (o *BusinessMetadataDefResponse) recurseRedact(v interface{}) {
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

func (o BusinessMetadataDefResponse) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o BusinessMetadataDefResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TypeVersion != nil {
		toSerialize["typeVersion"] = o.TypeVersion
	}
	if o.ServiceType != nil {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.AttributeDefs != nil {
		toSerialize["attributeDefs"] = o.AttributeDefs
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableBusinessMetadataDefResponse struct {
	value *BusinessMetadataDefResponse
	isSet bool
}

func (v NullableBusinessMetadataDefResponse) Get() *BusinessMetadataDefResponse {
	return v.value
}

func (v *NullableBusinessMetadataDefResponse) Set(val *BusinessMetadataDefResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessMetadataDefResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessMetadataDefResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessMetadataDefResponse(val *BusinessMetadataDefResponse) *NullableBusinessMetadataDefResponse {
	return &NullableBusinessMetadataDefResponse{value: val, isSet: true}
}

func (v NullableBusinessMetadataDefResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessMetadataDefResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
