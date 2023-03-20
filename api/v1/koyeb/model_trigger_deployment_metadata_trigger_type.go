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

// TriggerDeploymentMetadataTriggerType the model 'TriggerDeploymentMetadataTriggerType'
type TriggerDeploymentMetadataTriggerType string

// List of TriggerDeploymentMetadata.TriggerType
const (
	TRIGGERDEPLOYMENTMETADATATRIGGERTYPE_UNKNOWN_TYPE TriggerDeploymentMetadataTriggerType = "UNKNOWN_TYPE"
	TRIGGERDEPLOYMENTMETADATATRIGGERTYPE_GIT TriggerDeploymentMetadataTriggerType = "GIT"
	TRIGGERDEPLOYMENTMETADATATRIGGERTYPE_RESUME TriggerDeploymentMetadataTriggerType = "RESUME"
)

// All allowed values of TriggerDeploymentMetadataTriggerType enum
var AllowedTriggerDeploymentMetadataTriggerTypeEnumValues = []TriggerDeploymentMetadataTriggerType{
	"UNKNOWN_TYPE",
	"GIT",
	"RESUME",
}

func (v *TriggerDeploymentMetadataTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TriggerDeploymentMetadataTriggerType(value)
	for _, existing := range AllowedTriggerDeploymentMetadataTriggerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TriggerDeploymentMetadataTriggerType", value)
}

// NewTriggerDeploymentMetadataTriggerTypeFromValue returns a pointer to a valid TriggerDeploymentMetadataTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggerDeploymentMetadataTriggerTypeFromValue(v string) (*TriggerDeploymentMetadataTriggerType, error) {
	ev := TriggerDeploymentMetadataTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TriggerDeploymentMetadataTriggerType: valid values are %v", v, AllowedTriggerDeploymentMetadataTriggerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TriggerDeploymentMetadataTriggerType) IsValid() bool {
	for _, existing := range AllowedTriggerDeploymentMetadataTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerDeploymentMetadata.TriggerType value
func (v TriggerDeploymentMetadataTriggerType) Ptr() *TriggerDeploymentMetadataTriggerType {
	return &v
}

type NullableTriggerDeploymentMetadataTriggerType struct {
	value *TriggerDeploymentMetadataTriggerType
	isSet bool
}

func (v NullableTriggerDeploymentMetadataTriggerType) Get() *TriggerDeploymentMetadataTriggerType {
	return v.value
}

func (v *NullableTriggerDeploymentMetadataTriggerType) Set(val *TriggerDeploymentMetadataTriggerType) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerDeploymentMetadataTriggerType) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerDeploymentMetadataTriggerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerDeploymentMetadataTriggerType(val *TriggerDeploymentMetadataTriggerType) *NullableTriggerDeploymentMetadataTriggerType {
	return &NullableTriggerDeploymentMetadataTriggerType{value: val, isSet: true}
}

func (v NullableTriggerDeploymentMetadataTriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerDeploymentMetadataTriggerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

