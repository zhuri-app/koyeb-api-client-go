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

// ListRegionalDeploymentsReply struct for ListRegionalDeploymentsReply
type ListRegionalDeploymentsReply struct {
	RegionalDeployments *[]RegionalDeploymentListItem `json:"regional_deployments,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewListRegionalDeploymentsReply instantiates a new ListRegionalDeploymentsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRegionalDeploymentsReply() *ListRegionalDeploymentsReply {
	this := ListRegionalDeploymentsReply{}
	return &this
}

// NewListRegionalDeploymentsReplyWithDefaults instantiates a new ListRegionalDeploymentsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRegionalDeploymentsReplyWithDefaults() *ListRegionalDeploymentsReply {
	this := ListRegionalDeploymentsReply{}
	return &this
}

// GetRegionalDeployments returns the RegionalDeployments field value if set, zero value otherwise.
func (o *ListRegionalDeploymentsReply) GetRegionalDeployments() []RegionalDeploymentListItem {
	if o == nil || o.RegionalDeployments == nil {
		var ret []RegionalDeploymentListItem
		return ret
	}
	return *o.RegionalDeployments
}

// GetRegionalDeploymentsOk returns a tuple with the RegionalDeployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegionalDeploymentsReply) GetRegionalDeploymentsOk() (*[]RegionalDeploymentListItem, bool) {
	if o == nil || o.RegionalDeployments == nil {
		return nil, false
	}
	return o.RegionalDeployments, true
}

// HasRegionalDeployments returns a boolean if a field has been set.
func (o *ListRegionalDeploymentsReply) HasRegionalDeployments() bool {
	if o != nil && o.RegionalDeployments != nil {
		return true
	}

	return false
}

// SetRegionalDeployments gets a reference to the given []RegionalDeploymentListItem and assigns it to the RegionalDeployments field.
func (o *ListRegionalDeploymentsReply) SetRegionalDeployments(v []RegionalDeploymentListItem) {
	o.RegionalDeployments = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListRegionalDeploymentsReply) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegionalDeploymentsReply) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListRegionalDeploymentsReply) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListRegionalDeploymentsReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListRegionalDeploymentsReply) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegionalDeploymentsReply) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListRegionalDeploymentsReply) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ListRegionalDeploymentsReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListRegionalDeploymentsReply) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegionalDeploymentsReply) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListRegionalDeploymentsReply) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListRegionalDeploymentsReply) SetCount(v int64) {
	o.Count = &v
}

func (o ListRegionalDeploymentsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RegionalDeployments != nil {
		toSerialize["regional_deployments"] = o.RegionalDeployments
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

type NullableListRegionalDeploymentsReply struct {
	value *ListRegionalDeploymentsReply
	isSet bool
}

func (v NullableListRegionalDeploymentsReply) Get() *ListRegionalDeploymentsReply {
	return v.value
}

func (v *NullableListRegionalDeploymentsReply) Set(val *ListRegionalDeploymentsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListRegionalDeploymentsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListRegionalDeploymentsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRegionalDeploymentsReply(val *ListRegionalDeploymentsReply) *NullableListRegionalDeploymentsReply {
	return &NullableListRegionalDeploymentsReply{value: val, isSet: true}
}

func (v NullableListRegionalDeploymentsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRegionalDeploymentsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


