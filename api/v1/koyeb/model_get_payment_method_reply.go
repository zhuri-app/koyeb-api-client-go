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

// GetPaymentMethodReply struct for GetPaymentMethodReply
type GetPaymentMethodReply struct {
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
}

// NewGetPaymentMethodReply instantiates a new GetPaymentMethodReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPaymentMethodReply() *GetPaymentMethodReply {
	this := GetPaymentMethodReply{}
	return &this
}

// NewGetPaymentMethodReplyWithDefaults instantiates a new GetPaymentMethodReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPaymentMethodReplyWithDefaults() *GetPaymentMethodReply {
	this := GetPaymentMethodReply{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *GetPaymentMethodReply) GetPaymentMethod() PaymentMethod {
	if o == nil || isNil(o.PaymentMethod) {
		var ret PaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentMethodReply) GetPaymentMethodOk() (*PaymentMethod, bool) {
	if o == nil || isNil(o.PaymentMethod) {
    return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *GetPaymentMethodReply) HasPaymentMethod() bool {
	if o != nil && !isNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given PaymentMethod and assigns it to the PaymentMethod field.
func (o *GetPaymentMethodReply) SetPaymentMethod(v PaymentMethod) {
	o.PaymentMethod = &v
}

func (o GetPaymentMethodReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	return json.Marshal(toSerialize)
}

type NullableGetPaymentMethodReply struct {
	value *GetPaymentMethodReply
	isSet bool
}

func (v NullableGetPaymentMethodReply) Get() *GetPaymentMethodReply {
	return v.value
}

func (v *NullableGetPaymentMethodReply) Set(val *GetPaymentMethodReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPaymentMethodReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPaymentMethodReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPaymentMethodReply(val *GetPaymentMethodReply) *NullableGetPaymentMethodReply {
	return &NullableGetPaymentMethodReply{value: val, isSet: true}
}

func (v NullableGetPaymentMethodReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPaymentMethodReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


