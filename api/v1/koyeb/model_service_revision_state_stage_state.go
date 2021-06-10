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

// ServiceRevisionStateStageState struct for ServiceRevisionStateStageState
type ServiceRevisionStateStageState struct {
	Name *string `json:"name,omitempty"`
	StatusMessage *string `json:"status_message,omitempty"`
	Status *ServiceRevisionStateStageStateStatus `json:"status,omitempty"`
}

// NewServiceRevisionStateStageState instantiates a new ServiceRevisionStateStageState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRevisionStateStageState() *ServiceRevisionStateStageState {
	this := ServiceRevisionStateStageState{}
	var status ServiceRevisionStateStageStateStatus = SERVICEREVISIONSTATESTAGESTATESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewServiceRevisionStateStageStateWithDefaults instantiates a new ServiceRevisionStateStageState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRevisionStateStageStateWithDefaults() *ServiceRevisionStateStageState {
	this := ServiceRevisionStateStageState{}
	var status ServiceRevisionStateStageStateStatus = SERVICEREVISIONSTATESTAGESTATESTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceRevisionStateStageState) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRevisionStateStageState) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceRevisionStateStageState) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceRevisionStateStageState) SetName(v string) {
	o.Name = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ServiceRevisionStateStageState) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRevisionStateStageState) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ServiceRevisionStateStageState) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ServiceRevisionStateStageState) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceRevisionStateStageState) GetStatus() ServiceRevisionStateStageStateStatus {
	if o == nil || o.Status == nil {
		var ret ServiceRevisionStateStageStateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRevisionStateStageState) GetStatusOk() (*ServiceRevisionStateStageStateStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceRevisionStateStageState) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ServiceRevisionStateStageStateStatus and assigns it to the Status field.
func (o *ServiceRevisionStateStageState) SetStatus(v ServiceRevisionStateStageStateStatus) {
	o.Status = &v
}

func (o ServiceRevisionStateStageState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StatusMessage != nil {
		toSerialize["status_message"] = o.StatusMessage
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableServiceRevisionStateStageState struct {
	value *ServiceRevisionStateStageState
	isSet bool
}

func (v NullableServiceRevisionStateStageState) Get() *ServiceRevisionStateStageState {
	return v.value
}

func (v *NullableServiceRevisionStateStageState) Set(val *ServiceRevisionStateStageState) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRevisionStateStageState) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRevisionStateStageState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRevisionStateStageState(val *ServiceRevisionStateStageState) *NullableServiceRevisionStateStageState {
	return &NullableServiceRevisionStateStageState{value: val, isSet: true}
}

func (v NullableServiceRevisionStateStageState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceRevisionStateStageState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


