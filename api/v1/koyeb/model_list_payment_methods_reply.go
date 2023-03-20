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

// ListPaymentMethodsReply struct for ListPaymentMethodsReply
type ListPaymentMethodsReply struct {
	PaymentMethods []PaymentMethod `json:"payment_methods,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewListPaymentMethodsReply instantiates a new ListPaymentMethodsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPaymentMethodsReply() *ListPaymentMethodsReply {
	this := ListPaymentMethodsReply{}
	return &this
}

// NewListPaymentMethodsReplyWithDefaults instantiates a new ListPaymentMethodsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPaymentMethodsReplyWithDefaults() *ListPaymentMethodsReply {
	this := ListPaymentMethodsReply{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *ListPaymentMethodsReply) GetPaymentMethods() []PaymentMethod {
	if o == nil || isNil(o.PaymentMethods) {
		var ret []PaymentMethod
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPaymentMethodsReply) GetPaymentMethodsOk() ([]PaymentMethod, bool) {
	if o == nil || isNil(o.PaymentMethods) {
    return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *ListPaymentMethodsReply) HasPaymentMethods() bool {
	if o != nil && !isNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []PaymentMethod and assigns it to the PaymentMethods field.
func (o *ListPaymentMethodsReply) SetPaymentMethods(v []PaymentMethod) {
	o.PaymentMethods = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListPaymentMethodsReply) GetLimit() int64 {
	if o == nil || isNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPaymentMethodsReply) GetLimitOk() (*int64, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListPaymentMethodsReply) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListPaymentMethodsReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListPaymentMethodsReply) GetOffset() int64 {
	if o == nil || isNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPaymentMethodsReply) GetOffsetOk() (*int64, bool) {
	if o == nil || isNil(o.Offset) {
    return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListPaymentMethodsReply) HasOffset() bool {
	if o != nil && !isNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ListPaymentMethodsReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListPaymentMethodsReply) GetCount() int64 {
	if o == nil || isNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPaymentMethodsReply) GetCountOk() (*int64, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListPaymentMethodsReply) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListPaymentMethodsReply) SetCount(v int64) {
	o.Count = &v
}

func (o ListPaymentMethodsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableListPaymentMethodsReply struct {
	value *ListPaymentMethodsReply
	isSet bool
}

func (v NullableListPaymentMethodsReply) Get() *ListPaymentMethodsReply {
	return v.value
}

func (v *NullableListPaymentMethodsReply) Set(val *ListPaymentMethodsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListPaymentMethodsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListPaymentMethodsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPaymentMethodsReply(val *ListPaymentMethodsReply) *NullableListPaymentMethodsReply {
	return &NullableListPaymentMethodsReply{value: val, isSet: true}
}

func (v NullableListPaymentMethodsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPaymentMethodsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


