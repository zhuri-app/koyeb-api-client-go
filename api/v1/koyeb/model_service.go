/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// Service struct for Service
type Service struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	SucceededAt *time.Time `json:"succeeded_at,omitempty"`
	PausedAt *time.Time `json:"paused_at,omitempty"`
	ResumedAt *time.Time `json:"resumed_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
	Name *string `json:"name,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	Status *ServiceStatus `json:"status,omitempty"`
	Messages []string `json:"messages,omitempty"`
	Version *string `json:"version,omitempty"`
	ActiveDeploymentId *string `json:"active_deployment_id,omitempty"`
	LatestDeploymentId *string `json:"latest_deployment_id,omitempty"`
	State *ServiceState `json:"state,omitempty"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService() *Service {
	this := Service{}
	var status ServiceStatus = SERVICESTATUS_STARTING
	this.Status = &status
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	var status ServiceStatus = SERVICESTATUS_STARTING
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Service) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Service) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Service) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Service) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Service) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Service) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Service) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Service) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Service) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *Service) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *Service) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *Service) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetSucceededAt returns the SucceededAt field value if set, zero value otherwise.
func (o *Service) GetSucceededAt() time.Time {
	if o == nil || isNil(o.SucceededAt) {
		var ret time.Time
		return ret
	}
	return *o.SucceededAt
}

// GetSucceededAtOk returns a tuple with the SucceededAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSucceededAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.SucceededAt) {
    return nil, false
	}
	return o.SucceededAt, true
}

// HasSucceededAt returns a boolean if a field has been set.
func (o *Service) HasSucceededAt() bool {
	if o != nil && !isNil(o.SucceededAt) {
		return true
	}

	return false
}

// SetSucceededAt gets a reference to the given time.Time and assigns it to the SucceededAt field.
func (o *Service) SetSucceededAt(v time.Time) {
	o.SucceededAt = &v
}

// GetPausedAt returns the PausedAt field value if set, zero value otherwise.
func (o *Service) GetPausedAt() time.Time {
	if o == nil || isNil(o.PausedAt) {
		var ret time.Time
		return ret
	}
	return *o.PausedAt
}

// GetPausedAtOk returns a tuple with the PausedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPausedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.PausedAt) {
    return nil, false
	}
	return o.PausedAt, true
}

// HasPausedAt returns a boolean if a field has been set.
func (o *Service) HasPausedAt() bool {
	if o != nil && !isNil(o.PausedAt) {
		return true
	}

	return false
}

// SetPausedAt gets a reference to the given time.Time and assigns it to the PausedAt field.
func (o *Service) SetPausedAt(v time.Time) {
	o.PausedAt = &v
}

// GetResumedAt returns the ResumedAt field value if set, zero value otherwise.
func (o *Service) GetResumedAt() time.Time {
	if o == nil || isNil(o.ResumedAt) {
		var ret time.Time
		return ret
	}
	return *o.ResumedAt
}

// GetResumedAtOk returns a tuple with the ResumedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetResumedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ResumedAt) {
    return nil, false
	}
	return o.ResumedAt, true
}

// HasResumedAt returns a boolean if a field has been set.
func (o *Service) HasResumedAt() bool {
	if o != nil && !isNil(o.ResumedAt) {
		return true
	}

	return false
}

// SetResumedAt gets a reference to the given time.Time and assigns it to the ResumedAt field.
func (o *Service) SetResumedAt(v time.Time) {
	o.ResumedAt = &v
}

// GetTerminatedAt returns the TerminatedAt field value if set, zero value otherwise.
func (o *Service) GetTerminatedAt() time.Time {
	if o == nil || isNil(o.TerminatedAt) {
		var ret time.Time
		return ret
	}
	return *o.TerminatedAt
}

// GetTerminatedAtOk returns a tuple with the TerminatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTerminatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.TerminatedAt) {
    return nil, false
	}
	return o.TerminatedAt, true
}

// HasTerminatedAt returns a boolean if a field has been set.
func (o *Service) HasTerminatedAt() bool {
	if o != nil && !isNil(o.TerminatedAt) {
		return true
	}

	return false
}

// SetTerminatedAt gets a reference to the given time.Time and assigns it to the TerminatedAt field.
func (o *Service) SetTerminatedAt(v time.Time) {
	o.TerminatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Service) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Service) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Service) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *Service) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Service) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *Service) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *Service) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *Service) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *Service) SetAppId(v string) {
	o.AppId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Service) GetStatus() ServiceStatus {
	if o == nil || isNil(o.Status) {
		var ret ServiceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetStatusOk() (*ServiceStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Service) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ServiceStatus and assigns it to the Status field.
func (o *Service) SetStatus(v ServiceStatus) {
	o.Status = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *Service) GetMessages() []string {
	if o == nil || isNil(o.Messages) {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetMessagesOk() ([]string, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *Service) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *Service) SetMessages(v []string) {
	o.Messages = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Service) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Service) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Service) SetVersion(v string) {
	o.Version = &v
}

// GetActiveDeploymentId returns the ActiveDeploymentId field value if set, zero value otherwise.
func (o *Service) GetActiveDeploymentId() string {
	if o == nil || isNil(o.ActiveDeploymentId) {
		var ret string
		return ret
	}
	return *o.ActiveDeploymentId
}

// GetActiveDeploymentIdOk returns a tuple with the ActiveDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetActiveDeploymentIdOk() (*string, bool) {
	if o == nil || isNil(o.ActiveDeploymentId) {
    return nil, false
	}
	return o.ActiveDeploymentId, true
}

// HasActiveDeploymentId returns a boolean if a field has been set.
func (o *Service) HasActiveDeploymentId() bool {
	if o != nil && !isNil(o.ActiveDeploymentId) {
		return true
	}

	return false
}

// SetActiveDeploymentId gets a reference to the given string and assigns it to the ActiveDeploymentId field.
func (o *Service) SetActiveDeploymentId(v string) {
	o.ActiveDeploymentId = &v
}

// GetLatestDeploymentId returns the LatestDeploymentId field value if set, zero value otherwise.
func (o *Service) GetLatestDeploymentId() string {
	if o == nil || isNil(o.LatestDeploymentId) {
		var ret string
		return ret
	}
	return *o.LatestDeploymentId
}

// GetLatestDeploymentIdOk returns a tuple with the LatestDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetLatestDeploymentIdOk() (*string, bool) {
	if o == nil || isNil(o.LatestDeploymentId) {
    return nil, false
	}
	return o.LatestDeploymentId, true
}

// HasLatestDeploymentId returns a boolean if a field has been set.
func (o *Service) HasLatestDeploymentId() bool {
	if o != nil && !isNil(o.LatestDeploymentId) {
		return true
	}

	return false
}

// SetLatestDeploymentId gets a reference to the given string and assigns it to the LatestDeploymentId field.
func (o *Service) SetLatestDeploymentId(v string) {
	o.LatestDeploymentId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Service) GetState() ServiceState {
	if o == nil || isNil(o.State) {
		var ret ServiceState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetStateOk() (*ServiceState, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Service) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ServiceState and assigns it to the State field.
func (o *Service) SetState(v ServiceState) {
	o.State = &v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !isNil(o.SucceededAt) {
		toSerialize["succeeded_at"] = o.SucceededAt
	}
	if !isNil(o.PausedAt) {
		toSerialize["paused_at"] = o.PausedAt
	}
	if !isNil(o.ResumedAt) {
		toSerialize["resumed_at"] = o.ResumedAt
	}
	if !isNil(o.TerminatedAt) {
		toSerialize["terminated_at"] = o.TerminatedAt
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.ActiveDeploymentId) {
		toSerialize["active_deployment_id"] = o.ActiveDeploymentId
	}
	if !isNil(o.LatestDeploymentId) {
		toSerialize["latest_deployment_id"] = o.LatestDeploymentId
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


