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

// ListCatalogFunctionsReply struct for ListCatalogFunctionsReply
type ListCatalogFunctionsReply struct {
	CatalogFunctions *[]CatalogFunctionListItem `json:"catalog_functions,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewListCatalogFunctionsReply instantiates a new ListCatalogFunctionsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCatalogFunctionsReply() *ListCatalogFunctionsReply {
	this := ListCatalogFunctionsReply{}
	return &this
}

// NewListCatalogFunctionsReplyWithDefaults instantiates a new ListCatalogFunctionsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCatalogFunctionsReplyWithDefaults() *ListCatalogFunctionsReply {
	this := ListCatalogFunctionsReply{}
	return &this
}

// GetCatalogFunctions returns the CatalogFunctions field value if set, zero value otherwise.
func (o *ListCatalogFunctionsReply) GetCatalogFunctions() []CatalogFunctionListItem {
	if o == nil || o.CatalogFunctions == nil {
		var ret []CatalogFunctionListItem
		return ret
	}
	return *o.CatalogFunctions
}

// GetCatalogFunctionsOk returns a tuple with the CatalogFunctions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCatalogFunctionsReply) GetCatalogFunctionsOk() (*[]CatalogFunctionListItem, bool) {
	if o == nil || o.CatalogFunctions == nil {
		return nil, false
	}
	return o.CatalogFunctions, true
}

// HasCatalogFunctions returns a boolean if a field has been set.
func (o *ListCatalogFunctionsReply) HasCatalogFunctions() bool {
	if o != nil && o.CatalogFunctions != nil {
		return true
	}

	return false
}

// SetCatalogFunctions gets a reference to the given []CatalogFunctionListItem and assigns it to the CatalogFunctions field.
func (o *ListCatalogFunctionsReply) SetCatalogFunctions(v []CatalogFunctionListItem) {
	o.CatalogFunctions = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListCatalogFunctionsReply) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCatalogFunctionsReply) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListCatalogFunctionsReply) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListCatalogFunctionsReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListCatalogFunctionsReply) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCatalogFunctionsReply) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListCatalogFunctionsReply) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ListCatalogFunctionsReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListCatalogFunctionsReply) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCatalogFunctionsReply) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListCatalogFunctionsReply) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListCatalogFunctionsReply) SetCount(v int64) {
	o.Count = &v
}

func (o ListCatalogFunctionsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CatalogFunctions != nil {
		toSerialize["catalog_functions"] = o.CatalogFunctions
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableListCatalogFunctionsReply struct {
	value *ListCatalogFunctionsReply
	isSet bool
}

func (v NullableListCatalogFunctionsReply) Get() *ListCatalogFunctionsReply {
	return v.value
}

func (v *NullableListCatalogFunctionsReply) Set(val *ListCatalogFunctionsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListCatalogFunctionsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListCatalogFunctionsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCatalogFunctionsReply(val *ListCatalogFunctionsReply) *NullableListCatalogFunctionsReply {
	return &NullableListCatalogFunctionsReply{value: val, isSet: true}
}

func (v NullableListCatalogFunctionsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCatalogFunctionsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

