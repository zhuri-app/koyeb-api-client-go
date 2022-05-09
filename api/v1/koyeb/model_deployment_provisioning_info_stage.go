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
	"time"
)

// DeploymentProvisioningInfoStage struct for DeploymentProvisioningInfoStage
type DeploymentProvisioningInfoStage struct {
	Name *string `json:"name,omitempty"`
	Status *DeploymentProvisioningInfoStageStatus `json:"status,omitempty"`
	Messages *[]string `json:"messages,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	BuildAttempts *[]DeploymentProvisioningInfoStageBuildAttempt `json:"build_attempts,omitempty"`
}

// NewDeploymentProvisioningInfoStage instantiates a new DeploymentProvisioningInfoStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentProvisioningInfoStage() *DeploymentProvisioningInfoStage {
	this := DeploymentProvisioningInfoStage{}
	var status DeploymentProvisioningInfoStageStatus = DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewDeploymentProvisioningInfoStageWithDefaults instantiates a new DeploymentProvisioningInfoStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentProvisioningInfoStageWithDefaults() *DeploymentProvisioningInfoStage {
	this := DeploymentProvisioningInfoStage{}
	var status DeploymentProvisioningInfoStageStatus = DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentProvisioningInfoStage) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetStatus() DeploymentProvisioningInfoStageStatus {
	if o == nil || o.Status == nil {
		var ret DeploymentProvisioningInfoStageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetStatusOk() (*DeploymentProvisioningInfoStageStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentProvisioningInfoStageStatus and assigns it to the Status field.
func (o *DeploymentProvisioningInfoStage) SetStatus(v DeploymentProvisioningInfoStageStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetMessages() []string {
	if o == nil || o.Messages == nil {
		var ret []string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *DeploymentProvisioningInfoStage) SetMessages(v []string) {
	o.Messages = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DeploymentProvisioningInfoStage) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasFinishedAt() bool {
	if o != nil && o.FinishedAt != nil {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DeploymentProvisioningInfoStage) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetBuildAttempts returns the BuildAttempts field value if set, zero value otherwise.
func (o *DeploymentProvisioningInfoStage) GetBuildAttempts() []DeploymentProvisioningInfoStageBuildAttempt {
	if o == nil || o.BuildAttempts == nil {
		var ret []DeploymentProvisioningInfoStageBuildAttempt
		return ret
	}
	return *o.BuildAttempts
}

// GetBuildAttemptsOk returns a tuple with the BuildAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentProvisioningInfoStage) GetBuildAttemptsOk() (*[]DeploymentProvisioningInfoStageBuildAttempt, bool) {
	if o == nil || o.BuildAttempts == nil {
		return nil, false
	}
	return o.BuildAttempts, true
}

// HasBuildAttempts returns a boolean if a field has been set.
func (o *DeploymentProvisioningInfoStage) HasBuildAttempts() bool {
	if o != nil && o.BuildAttempts != nil {
		return true
	}

	return false
}

// SetBuildAttempts gets a reference to the given []DeploymentProvisioningInfoStageBuildAttempt and assigns it to the BuildAttempts field.
func (o *DeploymentProvisioningInfoStage) SetBuildAttempts(v []DeploymentProvisioningInfoStageBuildAttempt) {
	o.BuildAttempts = &v
}

func (o DeploymentProvisioningInfoStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.FinishedAt != nil {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if o.BuildAttempts != nil {
		toSerialize["build_attempts"] = o.BuildAttempts
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentProvisioningInfoStage struct {
	value *DeploymentProvisioningInfoStage
	isSet bool
}

func (v NullableDeploymentProvisioningInfoStage) Get() *DeploymentProvisioningInfoStage {
	return v.value
}

func (v *NullableDeploymentProvisioningInfoStage) Set(val *DeploymentProvisioningInfoStage) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentProvisioningInfoStage) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentProvisioningInfoStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentProvisioningInfoStage(val *DeploymentProvisioningInfoStage) *NullableDeploymentProvisioningInfoStage {
	return &NullableDeploymentProvisioningInfoStage{value: val, isSet: true}
}

func (v NullableDeploymentProvisioningInfoStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentProvisioningInfoStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


