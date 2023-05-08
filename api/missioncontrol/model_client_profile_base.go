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

// checks if the ClientProfileBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientProfileBase{}

// ClientProfileBase The client profile configured on the event broker service.
type ClientProfileBase struct {
	// The name of the client profile.
	Name string `json:"name"`
}

// NewClientProfileBase instantiates a new ClientProfileBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProfileBase(name string) *ClientProfileBase {
	this := ClientProfileBase{}
	this.Name = name
	return &this
}

// NewClientProfileBaseWithDefaults instantiates a new ClientProfileBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProfileBaseWithDefaults() *ClientProfileBase {
	this := ClientProfileBase{}
	return &this
}

// GetName returns the Name field value
func (o *ClientProfileBase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientProfileBase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientProfileBase) SetName(v string) {
	o.Name = v
}

func (o ClientProfileBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientProfileBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableClientProfileBase struct {
	value *ClientProfileBase
	isSet bool
}

func (v NullableClientProfileBase) Get() *ClientProfileBase {
	return v.value
}

func (v *NullableClientProfileBase) Set(val *ClientProfileBase) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProfileBase) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProfileBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProfileBase(val *ClientProfileBase) *NullableClientProfileBase {
	return &NullableClientProfileBase{value: val, isSet: true}
}

func (v NullableClientProfileBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProfileBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


