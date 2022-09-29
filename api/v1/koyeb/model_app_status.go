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

// AppStatus the model 'AppStatus'
type AppStatus string

// List of App.Status
const (
	APPSTATUS_STARTING AppStatus = "STARTING"
	APPSTATUS_HEALTHY AppStatus = "HEALTHY"
	APPSTATUS_DEGRADED AppStatus = "DEGRADED"
	APPSTATUS_UNHEALTHY AppStatus = "UNHEALTHY"
	APPSTATUS_DELETING AppStatus = "DELETING"
	APPSTATUS_DELETED AppStatus = "DELETED"
	APPSTATUS_PAUSING AppStatus = "PAUSING"
	APPSTATUS_PAUSED AppStatus = "PAUSED"
	APPSTATUS_RESUMING AppStatus = "RESUMING"
)

var allowedAppStatusEnumValues = []AppStatus{
	"STARTING",
	"HEALTHY",
	"DEGRADED",
	"UNHEALTHY",
	"DELETING",
	"DELETED",
	"PAUSING",
	"PAUSED",
	"RESUMING",
}

func (v *AppStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppStatus(value)
	for _, existing := range allowedAppStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppStatus", value)
}

// NewAppStatusFromValue returns a pointer to a valid AppStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppStatusFromValue(v string) (*AppStatus, error) {
	ev := AppStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppStatus: valid values are %v", v, allowedAppStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppStatus) IsValid() bool {
	for _, existing := range allowedAppStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to App.Status value
func (v AppStatus) Ptr() *AppStatus {
	return &v
}

type NullableAppStatus struct {
	value *AppStatus
	isSet bool
}

func (v NullableAppStatus) Get() *AppStatus {
	return v.value
}

func (v *NullableAppStatus) Set(val *AppStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStatus(val *AppStatus) *NullableAppStatus {
	return &NullableAppStatus{value: val, isSet: true}
}

func (v NullableAppStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
