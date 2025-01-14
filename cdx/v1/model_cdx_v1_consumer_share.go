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
Stream Share APIs

# Introduction

API version: 0.1.0-alpha0
Contact: cdx@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

import (
	"reflect"
)

// CdxV1ConsumerShare Resources accessible by the consumer   ## The Consumer Shared Resources Model <SchemaDefinition schemaRef=\"#/components/schemas/cdx.v1.ConsumerSharedResource\" />
type CdxV1ConsumerShare struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// Name of the consumer
	ConsumerUserName *string `json:"consumer_user_name,omitempty"`
	// Consumer organization name
	ConsumerOrganizationName *string `json:"consumer_organization_name,omitempty"`
	// Name of the provider
	ProviderUserName *string `json:"provider_user_name,omitempty"`
	// Status of share
	Status *string `json:"status,omitempty"`
	// The date and time at which the invitation will expire. Only for invited shares
	InviteExpiresAt *time.Time `json:"invite_expires_at,omitempty"`
	// The shared_resource to which this belongs.
	SharedResource *ObjectReference `json:"shared_resource,omitempty"`
}

// NewCdxV1ConsumerShare instantiates a new CdxV1ConsumerShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdxV1ConsumerShare() *CdxV1ConsumerShare {
	this := CdxV1ConsumerShare{}
	return &this
}

// NewCdxV1ConsumerShareWithDefaults instantiates a new CdxV1ConsumerShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdxV1ConsumerShareWithDefaults() *CdxV1ConsumerShare {
	this := CdxV1ConsumerShare{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CdxV1ConsumerShare) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CdxV1ConsumerShare) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CdxV1ConsumerShare) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *CdxV1ConsumerShare) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetConsumerUserName returns the ConsumerUserName field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetConsumerUserName() string {
	if o == nil || o.ConsumerUserName == nil {
		var ret string
		return ret
	}
	return *o.ConsumerUserName
}

// GetConsumerUserNameOk returns a tuple with the ConsumerUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetConsumerUserNameOk() (*string, bool) {
	if o == nil || o.ConsumerUserName == nil {
		return nil, false
	}
	return o.ConsumerUserName, true
}

// HasConsumerUserName returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasConsumerUserName() bool {
	if o != nil && o.ConsumerUserName != nil {
		return true
	}

	return false
}

// SetConsumerUserName gets a reference to the given string and assigns it to the ConsumerUserName field.
func (o *CdxV1ConsumerShare) SetConsumerUserName(v string) {
	o.ConsumerUserName = &v
}

// GetConsumerOrganizationName returns the ConsumerOrganizationName field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetConsumerOrganizationName() string {
	if o == nil || o.ConsumerOrganizationName == nil {
		var ret string
		return ret
	}
	return *o.ConsumerOrganizationName
}

// GetConsumerOrganizationNameOk returns a tuple with the ConsumerOrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetConsumerOrganizationNameOk() (*string, bool) {
	if o == nil || o.ConsumerOrganizationName == nil {
		return nil, false
	}
	return o.ConsumerOrganizationName, true
}

// HasConsumerOrganizationName returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasConsumerOrganizationName() bool {
	if o != nil && o.ConsumerOrganizationName != nil {
		return true
	}

	return false
}

// SetConsumerOrganizationName gets a reference to the given string and assigns it to the ConsumerOrganizationName field.
func (o *CdxV1ConsumerShare) SetConsumerOrganizationName(v string) {
	o.ConsumerOrganizationName = &v
}

// GetProviderUserName returns the ProviderUserName field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetProviderUserName() string {
	if o == nil || o.ProviderUserName == nil {
		var ret string
		return ret
	}
	return *o.ProviderUserName
}

// GetProviderUserNameOk returns a tuple with the ProviderUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetProviderUserNameOk() (*string, bool) {
	if o == nil || o.ProviderUserName == nil {
		return nil, false
	}
	return o.ProviderUserName, true
}

// HasProviderUserName returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasProviderUserName() bool {
	if o != nil && o.ProviderUserName != nil {
		return true
	}

	return false
}

// SetProviderUserName gets a reference to the given string and assigns it to the ProviderUserName field.
func (o *CdxV1ConsumerShare) SetProviderUserName(v string) {
	o.ProviderUserName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CdxV1ConsumerShare) SetStatus(v string) {
	o.Status = &v
}

// GetInviteExpiresAt returns the InviteExpiresAt field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetInviteExpiresAt() time.Time {
	if o == nil || o.InviteExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.InviteExpiresAt
}

// GetInviteExpiresAtOk returns a tuple with the InviteExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetInviteExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.InviteExpiresAt == nil {
		return nil, false
	}
	return o.InviteExpiresAt, true
}

// HasInviteExpiresAt returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasInviteExpiresAt() bool {
	if o != nil && o.InviteExpiresAt != nil {
		return true
	}

	return false
}

// SetInviteExpiresAt gets a reference to the given time.Time and assigns it to the InviteExpiresAt field.
func (o *CdxV1ConsumerShare) SetInviteExpiresAt(v time.Time) {
	o.InviteExpiresAt = &v
}

// GetSharedResource returns the SharedResource field value if set, zero value otherwise.
func (o *CdxV1ConsumerShare) GetSharedResource() ObjectReference {
	if o == nil || o.SharedResource == nil {
		var ret ObjectReference
		return ret
	}
	return *o.SharedResource
}

// GetSharedResourceOk returns a tuple with the SharedResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1ConsumerShare) GetSharedResourceOk() (*ObjectReference, bool) {
	if o == nil || o.SharedResource == nil {
		return nil, false
	}
	return o.SharedResource, true
}

// HasSharedResource returns a boolean if a field has been set.
func (o *CdxV1ConsumerShare) HasSharedResource() bool {
	if o != nil && o.SharedResource != nil {
		return true
	}

	return false
}

// SetSharedResource gets a reference to the given ObjectReference and assigns it to the SharedResource field.
func (o *CdxV1ConsumerShare) SetSharedResource(v ObjectReference) {
	o.SharedResource = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CdxV1ConsumerShare) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.ConsumerUserName)
	o.recurseRedact(o.ConsumerOrganizationName)
	o.recurseRedact(o.ProviderUserName)
	o.recurseRedact(o.Status)
	o.recurseRedact(o.InviteExpiresAt)
	o.recurseRedact(o.SharedResource)
}

func (o *CdxV1ConsumerShare) recurseRedact(v interface{}) {
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

func (o CdxV1ConsumerShare) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CdxV1ConsumerShare) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ConsumerUserName != nil {
		toSerialize["consumer_user_name"] = o.ConsumerUserName
	}
	if o.ConsumerOrganizationName != nil {
		toSerialize["consumer_organization_name"] = o.ConsumerOrganizationName
	}
	if o.ProviderUserName != nil {
		toSerialize["provider_user_name"] = o.ProviderUserName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.InviteExpiresAt != nil {
		toSerialize["invite_expires_at"] = o.InviteExpiresAt
	}
	if o.SharedResource != nil {
		toSerialize["shared_resource"] = o.SharedResource
	}
	return json.Marshal(toSerialize)
}

type NullableCdxV1ConsumerShare struct {
	value *CdxV1ConsumerShare
	isSet bool
}

func (v NullableCdxV1ConsumerShare) Get() *CdxV1ConsumerShare {
	return v.value
}

func (v *NullableCdxV1ConsumerShare) Set(val *CdxV1ConsumerShare) {
	v.value = val
	v.isSet = true
}

func (v NullableCdxV1ConsumerShare) IsSet() bool {
	return v.isSet
}

func (v *NullableCdxV1ConsumerShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdxV1ConsumerShare(val *CdxV1ConsumerShare) *NullableCdxV1ConsumerShare {
	return &NullableCdxV1ConsumerShare{value: val, isSet: true}
}

func (v NullableCdxV1ConsumerShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdxV1ConsumerShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
