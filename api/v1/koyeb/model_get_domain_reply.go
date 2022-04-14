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

// GetDomainReply struct for GetDomainReply
type GetDomainReply struct {
	Domain *Domain `json:"domain,omitempty"`
}

// NewGetDomainReply instantiates a new GetDomainReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDomainReply() *GetDomainReply {
	this := GetDomainReply{}
	return &this
}

// NewGetDomainReplyWithDefaults instantiates a new GetDomainReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDomainReplyWithDefaults() *GetDomainReply {
	this := GetDomainReply{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *GetDomainReply) GetDomain() Domain {
	if o == nil || o.Domain == nil {
		var ret Domain
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDomainReply) GetDomainOk() (*Domain, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *GetDomainReply) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given Domain and assigns it to the Domain field.
func (o *GetDomainReply) SetDomain(v Domain) {
	o.Domain = &v
}

func (o GetDomainReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	return json.Marshal(toSerialize)
}

type NullableGetDomainReply struct {
	value *GetDomainReply
	isSet bool
}

func (v NullableGetDomainReply) Get() *GetDomainReply {
	return v.value
}

func (v *NullableGetDomainReply) Set(val *GetDomainReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDomainReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDomainReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDomainReply(val *GetDomainReply) *NullableGetDomainReply {
	return &NullableGetDomainReply{value: val, isSet: true}
}

func (v NullableGetDomainReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDomainReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

