/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// CreateS3Credential struct for CreateS3Credential
type CreateS3Credential struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewCreateS3Credential instantiates a new CreateS3Credential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateS3Credential() *CreateS3Credential {
	this := CreateS3Credential{}
	return &this
}

// NewCreateS3CredentialWithDefaults instantiates a new CreateS3Credential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateS3CredentialWithDefaults() *CreateS3Credential {
	this := CreateS3Credential{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateS3Credential) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3Credential) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateS3Credential) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateS3Credential) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateS3Credential) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateS3Credential) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateS3Credential) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateS3Credential) SetDescription(v string) {
	o.Description = &v
}

func (o CreateS3Credential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCreateS3Credential struct {
	value *CreateS3Credential
	isSet bool
}

func (v NullableCreateS3Credential) Get() *CreateS3Credential {
	return v.value
}

func (v *NullableCreateS3Credential) Set(val *CreateS3Credential) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateS3Credential) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateS3Credential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateS3Credential(val *CreateS3Credential) *NullableCreateS3Credential {
	return &NullableCreateS3Credential{value: val, isSet: true}
}

func (v NullableCreateS3Credential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateS3Credential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

