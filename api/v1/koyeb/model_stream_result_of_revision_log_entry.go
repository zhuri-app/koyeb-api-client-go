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

// StreamResultOfRevisionLogEntry struct for StreamResultOfRevisionLogEntry
type StreamResultOfRevisionLogEntry struct {
	Result *RevisionLogEntry `json:"result,omitempty"`
	Error *GoogleRpcStatus `json:"error,omitempty"`
}

// NewStreamResultOfRevisionLogEntry instantiates a new StreamResultOfRevisionLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamResultOfRevisionLogEntry() *StreamResultOfRevisionLogEntry {
	this := StreamResultOfRevisionLogEntry{}
	return &this
}

// NewStreamResultOfRevisionLogEntryWithDefaults instantiates a new StreamResultOfRevisionLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamResultOfRevisionLogEntryWithDefaults() *StreamResultOfRevisionLogEntry {
	this := StreamResultOfRevisionLogEntry{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *StreamResultOfRevisionLogEntry) GetResult() RevisionLogEntry {
	if o == nil || o.Result == nil {
		var ret RevisionLogEntry
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamResultOfRevisionLogEntry) GetResultOk() (*RevisionLogEntry, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *StreamResultOfRevisionLogEntry) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given RevisionLogEntry and assigns it to the Result field.
func (o *StreamResultOfRevisionLogEntry) SetResult(v RevisionLogEntry) {
	o.Result = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StreamResultOfRevisionLogEntry) GetError() GoogleRpcStatus {
	if o == nil || o.Error == nil {
		var ret GoogleRpcStatus
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamResultOfRevisionLogEntry) GetErrorOk() (*GoogleRpcStatus, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StreamResultOfRevisionLogEntry) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given GoogleRpcStatus and assigns it to the Error field.
func (o *StreamResultOfRevisionLogEntry) SetError(v GoogleRpcStatus) {
	o.Error = &v
}

func (o StreamResultOfRevisionLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableStreamResultOfRevisionLogEntry struct {
	value *StreamResultOfRevisionLogEntry
	isSet bool
}

func (v NullableStreamResultOfRevisionLogEntry) Get() *StreamResultOfRevisionLogEntry {
	return v.value
}

func (v *NullableStreamResultOfRevisionLogEntry) Set(val *StreamResultOfRevisionLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamResultOfRevisionLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamResultOfRevisionLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamResultOfRevisionLogEntry(val *StreamResultOfRevisionLogEntry) *NullableStreamResultOfRevisionLogEntry {
	return &NullableStreamResultOfRevisionLogEntry{value: val, isSet: true}
}

func (v NullableStreamResultOfRevisionLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamResultOfRevisionLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

