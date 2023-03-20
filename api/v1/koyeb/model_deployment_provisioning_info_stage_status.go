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

// DeploymentProvisioningInfoStageStatus the model 'DeploymentProvisioningInfoStageStatus'
type DeploymentProvisioningInfoStageStatus string

// List of DeploymentProvisioningInfo.Stage.Status
const (
	DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_UNKNOWN DeploymentProvisioningInfoStageStatus = "UNKNOWN"
	DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_RUNNING DeploymentProvisioningInfoStageStatus = "RUNNING"
	DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_FAILED DeploymentProvisioningInfoStageStatus = "FAILED"
	DEPLOYMENTPROVISIONINGINFOSTAGESTATUS_COMPLETED DeploymentProvisioningInfoStageStatus = "COMPLETED"
)

// All allowed values of DeploymentProvisioningInfoStageStatus enum
var AllowedDeploymentProvisioningInfoStageStatusEnumValues = []DeploymentProvisioningInfoStageStatus{
	"UNKNOWN",
	"RUNNING",
	"FAILED",
	"COMPLETED",
}

func (v *DeploymentProvisioningInfoStageStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeploymentProvisioningInfoStageStatus(value)
	for _, existing := range AllowedDeploymentProvisioningInfoStageStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeploymentProvisioningInfoStageStatus", value)
}

// NewDeploymentProvisioningInfoStageStatusFromValue returns a pointer to a valid DeploymentProvisioningInfoStageStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeploymentProvisioningInfoStageStatusFromValue(v string) (*DeploymentProvisioningInfoStageStatus, error) {
	ev := DeploymentProvisioningInfoStageStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeploymentProvisioningInfoStageStatus: valid values are %v", v, AllowedDeploymentProvisioningInfoStageStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeploymentProvisioningInfoStageStatus) IsValid() bool {
	for _, existing := range AllowedDeploymentProvisioningInfoStageStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentProvisioningInfo.Stage.Status value
func (v DeploymentProvisioningInfoStageStatus) Ptr() *DeploymentProvisioningInfoStageStatus {
	return &v
}

type NullableDeploymentProvisioningInfoStageStatus struct {
	value *DeploymentProvisioningInfoStageStatus
	isSet bool
}

func (v NullableDeploymentProvisioningInfoStageStatus) Get() *DeploymentProvisioningInfoStageStatus {
	return v.value
}

func (v *NullableDeploymentProvisioningInfoStageStatus) Set(val *DeploymentProvisioningInfoStageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentProvisioningInfoStageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentProvisioningInfoStageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentProvisioningInfoStageStatus(val *DeploymentProvisioningInfoStageStatus) *NullableDeploymentProvisioningInfoStageStatus {
	return &NullableDeploymentProvisioningInfoStageStatus{value: val, isSet: true}
}

func (v NullableDeploymentProvisioningInfoStageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentProvisioningInfoStageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

