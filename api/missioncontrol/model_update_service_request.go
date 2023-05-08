/*
MISSION CONTROL

Documentation for Mission Control API

API version: V2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServiceRequest{}

// UpdateServiceRequest struct for UpdateServiceRequest
type UpdateServiceRequest struct {
	// The new service name. The new service name must be unique within an organization.
	Name *string `json:"name,omitempty"`
	// The owner of the event broker service. The owner must belong to the same organization.
	OwnedBy *string `json:"ownedBy,omitempty"`
	// Lock (true) and unlock (false) the event broker service. If an event broker service is locked it's protected from being deleted.
	Locked *bool `json:"locked,omitempty"`
}

// NewUpdateServiceRequest instantiates a new UpdateServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServiceRequest() *UpdateServiceRequest {
	this := UpdateServiceRequest{}
	return &this
}

// NewUpdateServiceRequestWithDefaults instantiates a new UpdateServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceRequestWithDefaults() *UpdateServiceRequest {
	this := UpdateServiceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateServiceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateServiceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateServiceRequest) SetName(v string) {
	o.Name = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *UpdateServiceRequest) GetOwnedBy() string {
	if o == nil || IsNil(o.OwnedBy) {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceRequest) GetOwnedByOk() (*string, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *UpdateServiceRequest) HasOwnedBy() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *UpdateServiceRequest) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *UpdateServiceRequest) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServiceRequest) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *UpdateServiceRequest) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *UpdateServiceRequest) SetLocked(v bool) {
	o.Locked = &v
}

func (o UpdateServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OwnedBy) {
		toSerialize["ownedBy"] = o.OwnedBy
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	return toSerialize, nil
}

type NullableUpdateServiceRequest struct {
	value *UpdateServiceRequest
	isSet bool
}

func (v NullableUpdateServiceRequest) Get() *UpdateServiceRequest {
	return v.value
}

func (v *NullableUpdateServiceRequest) Set(val *UpdateServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServiceRequest(val *UpdateServiceRequest) *NullableUpdateServiceRequest {
	return &NullableUpdateServiceRequest{value: val, isSet: true}
}

func (v NullableUpdateServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


