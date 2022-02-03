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

// DesiredDeploymentGroup struct for DesiredDeploymentGroup
type DesiredDeploymentGroup struct {
	Name *string `json:"name,omitempty"`
	RevisionIds *[]string `json:"revision_ids,omitempty"`
	DeploymentIds *[]string `json:"deployment_ids,omitempty"`
}

// NewDesiredDeploymentGroup instantiates a new DesiredDeploymentGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesiredDeploymentGroup() *DesiredDeploymentGroup {
	this := DesiredDeploymentGroup{}
	return &this
}

// NewDesiredDeploymentGroupWithDefaults instantiates a new DesiredDeploymentGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesiredDeploymentGroupWithDefaults() *DesiredDeploymentGroup {
	this := DesiredDeploymentGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DesiredDeploymentGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesiredDeploymentGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DesiredDeploymentGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DesiredDeploymentGroup) SetName(v string) {
	o.Name = &v
}

// GetRevisionIds returns the RevisionIds field value if set, zero value otherwise.
func (o *DesiredDeploymentGroup) GetRevisionIds() []string {
	if o == nil || o.RevisionIds == nil {
		var ret []string
		return ret
	}
	return *o.RevisionIds
}

// GetRevisionIdsOk returns a tuple with the RevisionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesiredDeploymentGroup) GetRevisionIdsOk() (*[]string, bool) {
	if o == nil || o.RevisionIds == nil {
		return nil, false
	}
	return o.RevisionIds, true
}

// HasRevisionIds returns a boolean if a field has been set.
func (o *DesiredDeploymentGroup) HasRevisionIds() bool {
	if o != nil && o.RevisionIds != nil {
		return true
	}

	return false
}

// SetRevisionIds gets a reference to the given []string and assigns it to the RevisionIds field.
func (o *DesiredDeploymentGroup) SetRevisionIds(v []string) {
	o.RevisionIds = &v
}

// GetDeploymentIds returns the DeploymentIds field value if set, zero value otherwise.
func (o *DesiredDeploymentGroup) GetDeploymentIds() []string {
	if o == nil || o.DeploymentIds == nil {
		var ret []string
		return ret
	}
	return *o.DeploymentIds
}

// GetDeploymentIdsOk returns a tuple with the DeploymentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesiredDeploymentGroup) GetDeploymentIdsOk() (*[]string, bool) {
	if o == nil || o.DeploymentIds == nil {
		return nil, false
	}
	return o.DeploymentIds, true
}

// HasDeploymentIds returns a boolean if a field has been set.
func (o *DesiredDeploymentGroup) HasDeploymentIds() bool {
	if o != nil && o.DeploymentIds != nil {
		return true
	}

	return false
}

// SetDeploymentIds gets a reference to the given []string and assigns it to the DeploymentIds field.
func (o *DesiredDeploymentGroup) SetDeploymentIds(v []string) {
	o.DeploymentIds = &v
}

func (o DesiredDeploymentGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RevisionIds != nil {
		toSerialize["revision_ids"] = o.RevisionIds
	}
	if o.DeploymentIds != nil {
		toSerialize["deployment_ids"] = o.DeploymentIds
	}
	return json.Marshal(toSerialize)
}

type NullableDesiredDeploymentGroup struct {
	value *DesiredDeploymentGroup
	isSet bool
}

func (v NullableDesiredDeploymentGroup) Get() *DesiredDeploymentGroup {
	return v.value
}

func (v *NullableDesiredDeploymentGroup) Set(val *DesiredDeploymentGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableDesiredDeploymentGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableDesiredDeploymentGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesiredDeploymentGroup(val *DesiredDeploymentGroup) *NullableDesiredDeploymentGroup {
	return &NullableDesiredDeploymentGroup{value: val, isSet: true}
}

func (v NullableDesiredDeploymentGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesiredDeploymentGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

