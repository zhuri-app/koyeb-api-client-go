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

// AutoReleaseGroup struct for AutoReleaseGroup
type AutoReleaseGroup struct {
	Name *string `json:"name,omitempty"`
	Repository *string `json:"repository,omitempty"`
	GitRef *string `json:"git_ref,omitempty"`
	LatestSha *string `json:"latest_sha,omitempty"`
}

// NewAutoReleaseGroup instantiates a new AutoReleaseGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoReleaseGroup() *AutoReleaseGroup {
	this := AutoReleaseGroup{}
	return &this
}

// NewAutoReleaseGroupWithDefaults instantiates a new AutoReleaseGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoReleaseGroupWithDefaults() *AutoReleaseGroup {
	this := AutoReleaseGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AutoReleaseGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReleaseGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AutoReleaseGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AutoReleaseGroup) SetName(v string) {
	o.Name = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *AutoReleaseGroup) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReleaseGroup) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *AutoReleaseGroup) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *AutoReleaseGroup) SetRepository(v string) {
	o.Repository = &v
}

// GetGitRef returns the GitRef field value if set, zero value otherwise.
func (o *AutoReleaseGroup) GetGitRef() string {
	if o == nil || o.GitRef == nil {
		var ret string
		return ret
	}
	return *o.GitRef
}

// GetGitRefOk returns a tuple with the GitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReleaseGroup) GetGitRefOk() (*string, bool) {
	if o == nil || o.GitRef == nil {
		return nil, false
	}
	return o.GitRef, true
}

// HasGitRef returns a boolean if a field has been set.
func (o *AutoReleaseGroup) HasGitRef() bool {
	if o != nil && o.GitRef != nil {
		return true
	}

	return false
}

// SetGitRef gets a reference to the given string and assigns it to the GitRef field.
func (o *AutoReleaseGroup) SetGitRef(v string) {
	o.GitRef = &v
}

// GetLatestSha returns the LatestSha field value if set, zero value otherwise.
func (o *AutoReleaseGroup) GetLatestSha() string {
	if o == nil || o.LatestSha == nil {
		var ret string
		return ret
	}
	return *o.LatestSha
}

// GetLatestShaOk returns a tuple with the LatestSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReleaseGroup) GetLatestShaOk() (*string, bool) {
	if o == nil || o.LatestSha == nil {
		return nil, false
	}
	return o.LatestSha, true
}

// HasLatestSha returns a boolean if a field has been set.
func (o *AutoReleaseGroup) HasLatestSha() bool {
	if o != nil && o.LatestSha != nil {
		return true
	}

	return false
}

// SetLatestSha gets a reference to the given string and assigns it to the LatestSha field.
func (o *AutoReleaseGroup) SetLatestSha(v string) {
	o.LatestSha = &v
}

func (o AutoReleaseGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.GitRef != nil {
		toSerialize["git_ref"] = o.GitRef
	}
	if o.LatestSha != nil {
		toSerialize["latest_sha"] = o.LatestSha
	}
	return json.Marshal(toSerialize)
}

type NullableAutoReleaseGroup struct {
	value *AutoReleaseGroup
	isSet bool
}

func (v NullableAutoReleaseGroup) Get() *AutoReleaseGroup {
	return v.value
}

func (v *NullableAutoReleaseGroup) Set(val *AutoReleaseGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoReleaseGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoReleaseGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoReleaseGroup(val *AutoReleaseGroup) *NullableAutoReleaseGroup {
	return &NullableAutoReleaseGroup{value: val, isSet: true}
}

func (v NullableAutoReleaseGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoReleaseGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

