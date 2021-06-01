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

// GoogleProtobufNullValue `NullValue` is a singleton enumeration to represent the null value for the `Value` type union.   The JSON representation for `NullValue` is JSON `null`.   - NULL_VALUE: Null value.
type GoogleProtobufNullValue string

// List of google.protobuf.NullValue
const (
	GOOGLEPROTOBUFNULLVALUE_NULL_VALUE GoogleProtobufNullValue = "NULL_VALUE"
)

var allowedGoogleProtobufNullValueEnumValues = []GoogleProtobufNullValue{
	"NULL_VALUE",
}

func (v *GoogleProtobufNullValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GoogleProtobufNullValue(value)
	for _, existing := range allowedGoogleProtobufNullValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GoogleProtobufNullValue", value)
}

// NewGoogleProtobufNullValueFromValue returns a pointer to a valid GoogleProtobufNullValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGoogleProtobufNullValueFromValue(v string) (*GoogleProtobufNullValue, error) {
	ev := GoogleProtobufNullValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GoogleProtobufNullValue: valid values are %v", v, allowedGoogleProtobufNullValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GoogleProtobufNullValue) IsValid() bool {
	for _, existing := range allowedGoogleProtobufNullValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to google.protobuf.NullValue value
func (v GoogleProtobufNullValue) Ptr() *GoogleProtobufNullValue {
	return &v
}

type NullableGoogleProtobufNullValue struct {
	value *GoogleProtobufNullValue
	isSet bool
}

func (v NullableGoogleProtobufNullValue) Get() *GoogleProtobufNullValue {
	return v.value
}

func (v *NullableGoogleProtobufNullValue) Set(val *GoogleProtobufNullValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleProtobufNullValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleProtobufNullValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleProtobufNullValue(val *GoogleProtobufNullValue) *NullableGoogleProtobufNullValue {
	return &NullableGoogleProtobufNullValue{value: val, isSet: true}
}

func (v NullableGoogleProtobufNullValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleProtobufNullValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
