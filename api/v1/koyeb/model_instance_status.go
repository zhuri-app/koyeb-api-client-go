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
	"fmt"
)

// InstanceStatus the model 'InstanceStatus'
type InstanceStatus string

// List of Instance.Status
const (
	INSTANCESTATUS_ALLOCATING InstanceStatus = "ALLOCATING"
	INSTANCESTATUS_STARTING InstanceStatus = "STARTING"
	INSTANCESTATUS_HEALTHY InstanceStatus = "HEALTHY"
	INSTANCESTATUS_UNHEALTHY InstanceStatus = "UNHEALTHY"
	INSTANCESTATUS_STOPPING InstanceStatus = "STOPPING"
	INSTANCESTATUS_STOPPED InstanceStatus = "STOPPED"
	INSTANCESTATUS_ERROR InstanceStatus = "ERROR"
)

var allowedInstanceStatusEnumValues = []InstanceStatus{
	"ALLOCATING",
	"STARTING",
	"HEALTHY",
	"UNHEALTHY",
	"STOPPING",
	"STOPPED",
	"ERROR",
}

func (v *InstanceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InstanceStatus(value)
	for _, existing := range allowedInstanceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstanceStatus", value)
}

// NewInstanceStatusFromValue returns a pointer to a valid InstanceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceStatusFromValue(v string) (*InstanceStatus, error) {
	ev := InstanceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceStatus: valid values are %v", v, allowedInstanceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceStatus) IsValid() bool {
	for _, existing := range allowedInstanceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Instance.Status value
func (v InstanceStatus) Ptr() *InstanceStatus {
	return &v
}

type NullableInstanceStatus struct {
	value *InstanceStatus
	isSet bool
}

func (v NullableInstanceStatus) Get() *InstanceStatus {
	return v.value
}

func (v *NullableInstanceStatus) Set(val *InstanceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceStatus(val *InstanceStatus) *NullableInstanceStatus {
	return &NullableInstanceStatus{value: val, isSet: true}
}

func (v NullableInstanceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

