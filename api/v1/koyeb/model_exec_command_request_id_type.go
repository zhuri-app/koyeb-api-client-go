/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"fmt"
)

// ExecCommandRequestIdType the model 'ExecCommandRequestIdType'
type ExecCommandRequestIdType string

// List of ExecCommandRequest.IdType
const (
	EXECCOMMANDREQUESTIDTYPE_INVALID ExecCommandRequestIdType = "INVALID"
	EXECCOMMANDREQUESTIDTYPE_INSTANCE_ID ExecCommandRequestIdType = "INSTANCE_ID"
	EXECCOMMANDREQUESTIDTYPE_SERVICE_ID ExecCommandRequestIdType = "SERVICE_ID"
)

// All allowed values of ExecCommandRequestIdType enum
var AllowedExecCommandRequestIdTypeEnumValues = []ExecCommandRequestIdType{
	"INVALID",
	"INSTANCE_ID",
	"SERVICE_ID",
}

func (v *ExecCommandRequestIdType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExecCommandRequestIdType(value)
	for _, existing := range AllowedExecCommandRequestIdTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExecCommandRequestIdType", value)
}

// NewExecCommandRequestIdTypeFromValue returns a pointer to a valid ExecCommandRequestIdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExecCommandRequestIdTypeFromValue(v string) (*ExecCommandRequestIdType, error) {
	ev := ExecCommandRequestIdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExecCommandRequestIdType: valid values are %v", v, AllowedExecCommandRequestIdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExecCommandRequestIdType) IsValid() bool {
	for _, existing := range AllowedExecCommandRequestIdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExecCommandRequest.IdType value
func (v ExecCommandRequestIdType) Ptr() *ExecCommandRequestIdType {
	return &v
}

type NullableExecCommandRequestIdType struct {
	value *ExecCommandRequestIdType
	isSet bool
}

func (v NullableExecCommandRequestIdType) Get() *ExecCommandRequestIdType {
	return v.value
}

func (v *NullableExecCommandRequestIdType) Set(val *ExecCommandRequestIdType) {
	v.value = val
	v.isSet = true
}

func (v NullableExecCommandRequestIdType) IsSet() bool {
	return v.isSet
}

func (v *NullableExecCommandRequestIdType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecCommandRequestIdType(val *ExecCommandRequestIdType) *NullableExecCommandRequestIdType {
	return &NullableExecCommandRequestIdType{value: val, isSet: true}
}

func (v NullableExecCommandRequestIdType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecCommandRequestIdType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

