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
	"time"
)

// InstanceListItem struct for InstanceListItem
type InstanceListItem struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	ServiceId *string `json:"service_id,omitempty"`
	DeploymentId *string `json:"deployment_id,omitempty"`
	RegionalDeploymentId *string `json:"regional_deployment_id,omitempty"`
	AllocationId *string `json:"allocation_id,omitempty"`
	Datacenter *string `json:"datacenter,omitempty"`
	Status *InstanceStatus `json:"status,omitempty"`
	Messages *[]string `json:"messages,omitempty"`
}

// NewInstanceListItem instantiates a new InstanceListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceListItem() *InstanceListItem {
	this := InstanceListItem{}
	var status InstanceStatus = INSTANCESTATUS_ALLOCATING
	this.Status = &status
	return &this
}

// NewInstanceListItemWithDefaults instantiates a new InstanceListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceListItemWithDefaults() *InstanceListItem {
	this := InstanceListItem{}
	var status InstanceStatus = INSTANCESTATUS_ALLOCATING
	this.Status = &status
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

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InstanceListItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InstanceListItem) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InstanceListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InstanceListItem) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InstanceListItem) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InstanceListItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InstanceListItem) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InstanceListItem) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InstanceListItem) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *InstanceListItem) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *InstanceListItem) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *InstanceListItem) SetAppId(v string) {
	o.AppId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *InstanceListItem) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *InstanceListItem) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *InstanceListItem) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *InstanceListItem) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *InstanceListItem) HasDeploymentId() bool {
	if o != nil && o.DeploymentId != nil {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *InstanceListItem) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetRegionalDeploymentId returns the RegionalDeploymentId field value if set, zero value otherwise.
func (o *InstanceListItem) GetRegionalDeploymentId() string {
	if o == nil || o.RegionalDeploymentId == nil {
		var ret string
		return ret
	}
	return *o.RegionalDeploymentId
}

// GetRegionalDeploymentIdOk returns a tuple with the RegionalDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetRegionalDeploymentIdOk() (*string, bool) {
	if o == nil || o.RegionalDeploymentId == nil {
		return nil, false
	}
	return o.RegionalDeploymentId, true
}

// HasRegionalDeploymentId returns a boolean if a field has been set.
func (o *InstanceListItem) HasRegionalDeploymentId() bool {
	if o != nil && o.RegionalDeploymentId != nil {
		return true
	}

	return false
}

// SetRegionalDeploymentId gets a reference to the given string and assigns it to the RegionalDeploymentId field.
func (o *InstanceListItem) SetRegionalDeploymentId(v string) {
	o.RegionalDeploymentId = &v
}

// GetAllocationId returns the AllocationId field value if set, zero value otherwise.
func (o *InstanceListItem) GetAllocationId() string {
	if o == nil || o.AllocationId == nil {
		var ret string
		return ret
	}
	return *o.AllocationId
}

// GetAllocationIdOk returns a tuple with the AllocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetAllocationIdOk() (*string, bool) {
	if o == nil || o.AllocationId == nil {
		return nil, false
	}
	return o.AllocationId, true
}

// HasAllocationId returns a boolean if a field has been set.
func (o *InstanceListItem) HasAllocationId() bool {
	if o != nil && o.AllocationId != nil {
		return true
	}

	return false
}

// SetAllocationId gets a reference to the given string and assigns it to the AllocationId field.
func (o *InstanceListItem) SetAllocationId(v string) {
	o.AllocationId = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *InstanceListItem) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *InstanceListItem) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *InstanceListItem) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InstanceListItem) GetStatus() InstanceStatus {
	if o == nil || o.Status == nil {
		var ret InstanceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetStatusOk() (*InstanceStatus, bool) {
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

// SetStatus gets a reference to the given InstanceStatus and assigns it to the Status field.
func (o *InstanceListItem) SetStatus(v InstanceStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *InstanceListItem) GetMessages() []string {
	if o == nil || o.Messages == nil {
		var ret []string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListItem) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *InstanceListItem) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *InstanceListItem) SetMessages(v []string) {
	o.Messages = &v
}

func (o InstanceListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.DeploymentId != nil {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if o.RegionalDeploymentId != nil {
		toSerialize["regional_deployment_id"] = o.RegionalDeploymentId
	}
	if o.AllocationId != nil {
		toSerialize["allocation_id"] = o.AllocationId
	}
	if o.Datacenter != nil {
		toSerialize["datacenter"] = o.Datacenter
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
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


