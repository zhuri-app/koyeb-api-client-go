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

// InstanceListItem struct for InstanceListItem
type InstanceListItem struct {
	Id *string `json:"id,omitempty"`
	Description *string `json:"description,omitempty"`
	Vcpu *int64 `json:"vcpu,omitempty"`
	Memory *string `json:"memory,omitempty"`
	Disk *string `json:"disk,omitempty"`
	PriceHourly *string `json:"price_hourly,omitempty"`
	PriceMonthly *string `json:"price_monthly,omitempty"`
	Regions *[]string `json:"regions,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewInstanceListItem instantiates a new InstanceListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceListItem() *InstanceListItem {
	this := InstanceListItem{}
	return &this
}

// NewInstanceListItemWithDefaults instantiates a new InstanceListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceListItemWithDefaults() *InstanceListItem {
	this := InstanceListItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceListItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceListItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstanceListItem) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstanceListItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InstanceListItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstanceListItem) SetDescription(v string) {
	o.Description = &v
}

// GetVcpu returns the Vcpu field value if set, zero value otherwise.
func (o *InstanceListItem) GetVcpu() int64 {
	if o == nil || o.Vcpu == nil {
		var ret int64
		return ret
	}
	return *o.Vcpu
}

// GetVcpuOk returns a tuple with the Vcpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetVcpuOk() (*int64, bool) {
	if o == nil || o.Vcpu == nil {
		return nil, false
	}
	return o.Vcpu, true
}

// HasVcpu returns a boolean if a field has been set.
func (o *InstanceListItem) HasVcpu() bool {
	if o != nil && o.Vcpu != nil {
		return true
	}

	return false
}

// SetVcpu gets a reference to the given int64 and assigns it to the Vcpu field.
func (o *InstanceListItem) SetVcpu(v int64) {
	o.Vcpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *InstanceListItem) GetMemory() string {
	if o == nil || o.Memory == nil {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetMemoryOk() (*string, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *InstanceListItem) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *InstanceListItem) SetMemory(v string) {
	o.Memory = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *InstanceListItem) GetDisk() string {
	if o == nil || o.Disk == nil {
		var ret string
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetDiskOk() (*string, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *InstanceListItem) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given string and assigns it to the Disk field.
func (o *InstanceListItem) SetDisk(v string) {
	o.Disk = &v
}

// GetPriceHourly returns the PriceHourly field value if set, zero value otherwise.
func (o *InstanceListItem) GetPriceHourly() string {
	if o == nil || o.PriceHourly == nil {
		var ret string
		return ret
	}
	return *o.PriceHourly
}

// GetPriceHourlyOk returns a tuple with the PriceHourly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetPriceHourlyOk() (*string, bool) {
	if o == nil || o.PriceHourly == nil {
		return nil, false
	}
	return o.PriceHourly, true
}

// HasPriceHourly returns a boolean if a field has been set.
func (o *InstanceListItem) HasPriceHourly() bool {
	if o != nil && o.PriceHourly != nil {
		return true
	}

	return false
}

// SetPriceHourly gets a reference to the given string and assigns it to the PriceHourly field.
func (o *InstanceListItem) SetPriceHourly(v string) {
	o.PriceHourly = &v
}

// GetPriceMonthly returns the PriceMonthly field value if set, zero value otherwise.
func (o *InstanceListItem) GetPriceMonthly() string {
	if o == nil || o.PriceMonthly == nil {
		var ret string
		return ret
	}
	return *o.PriceMonthly
}

// GetPriceMonthlyOk returns a tuple with the PriceMonthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetPriceMonthlyOk() (*string, bool) {
	if o == nil || o.PriceMonthly == nil {
		return nil, false
	}
	return o.PriceMonthly, true
}

// HasPriceMonthly returns a boolean if a field has been set.
func (o *InstanceListItem) HasPriceMonthly() bool {
	if o != nil && o.PriceMonthly != nil {
		return true
	}

	return false
}

// SetPriceMonthly gets a reference to the given string and assigns it to the PriceMonthly field.
func (o *InstanceListItem) SetPriceMonthly(v string) {
	o.PriceMonthly = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *InstanceListItem) GetRegions() []string {
	if o == nil || o.Regions == nil {
		var ret []string
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetRegionsOk() (*[]string, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *InstanceListItem) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *InstanceListItem) SetRegions(v []string) {
	o.Regions = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InstanceListItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InstanceListItem) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InstanceListItem) SetStatus(v string) {
	o.Status = &v
}

func (o InstanceListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Vcpu != nil {
		toSerialize["vcpu"] = o.Vcpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Disk != nil {
		toSerialize["disk"] = o.Disk
	}
	if o.PriceHourly != nil {
		toSerialize["price_hourly"] = o.PriceHourly
	}
	if o.PriceMonthly != nil {
		toSerialize["price_monthly"] = o.PriceMonthly
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceListItem struct {
	value *InstanceListItem
	isSet bool
}

func (v NullableInstanceListItem) Get() *InstanceListItem {
	return v.value
}

func (v *NullableInstanceListItem) Set(val *InstanceListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceListItem(val *InstanceListItem) *NullableInstanceListItem {
	return &NullableInstanceListItem{value: val, isSet: true}
}

func (v NullableInstanceListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

