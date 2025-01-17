/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// Scaling struct for Scaling
type Scaling struct {
	Min *int64 `json:"min,omitempty"`
	Max *int64 `json:"max,omitempty"`
}

// NewScaling instantiates a new Scaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScaling() *Scaling {
	this := Scaling{}
	return &this
}

// NewScalingWithDefaults instantiates a new Scaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalingWithDefaults() *Scaling {
	this := Scaling{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Scaling) GetMin() int64 {
	if o == nil || isNil(o.Min) {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scaling) GetMinOk() (*int64, bool) {
	if o == nil || isNil(o.Min) {
    return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Scaling) HasMin() bool {
	if o != nil && !isNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *Scaling) SetMin(v int64) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Scaling) GetMax() int64 {
	if o == nil || isNil(o.Max) {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scaling) GetMaxOk() (*int64, bool) {
	if o == nil || isNil(o.Max) {
    return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Scaling) HasMax() bool {
	if o != nil && !isNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *Scaling) SetMax(v int64) {
	o.Max = &v
}

func (o Scaling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !isNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableScaling struct {
	value *Scaling
	isSet bool
}

func (v NullableScaling) Get() *Scaling {
	return v.value
}

func (v *NullableScaling) Set(val *Scaling) {
	v.value = val
	v.isSet = true
}

func (v NullableScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScaling(val *Scaling) *NullableScaling {
	return &NullableScaling{value: val, isSet: true}
}

func (v NullableScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


