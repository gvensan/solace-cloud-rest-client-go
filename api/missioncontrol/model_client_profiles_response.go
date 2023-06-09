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

// checks if the ClientProfilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientProfilesResponse{}

// ClientProfilesResponse struct for ClientProfilesResponse
type ClientProfilesResponse struct {
	Data []ClientProfileSummary `json:"data,omitempty"`
	Meta map[string]map[string]interface{} `json:"meta,omitempty"`
}

// NewClientProfilesResponse instantiates a new ClientProfilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProfilesResponse() *ClientProfilesResponse {
	this := ClientProfilesResponse{}
	return &this
}

// NewClientProfilesResponseWithDefaults instantiates a new ClientProfilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProfilesResponseWithDefaults() *ClientProfilesResponse {
	this := ClientProfilesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ClientProfilesResponse) GetData() []ClientProfileSummary {
	if o == nil || IsNil(o.Data) {
		var ret []ClientProfileSummary
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfilesResponse) GetDataOk() ([]ClientProfileSummary, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ClientProfilesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ClientProfileSummary and assigns it to the Data field.
func (o *ClientProfilesResponse) SetData(v []ClientProfileSummary) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ClientProfilesResponse) GetMeta() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfilesResponse) GetMetaOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ClientProfilesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]map[string]interface{} and assigns it to the Meta field.
func (o *ClientProfilesResponse) SetMeta(v map[string]map[string]interface{}) {
	o.Meta = v
}

func (o ClientProfilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientProfilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableClientProfilesResponse struct {
	value *ClientProfilesResponse
	isSet bool
}

func (v NullableClientProfilesResponse) Get() *ClientProfilesResponse {
	return v.value
}

func (v *NullableClientProfilesResponse) Set(val *ClientProfilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProfilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProfilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProfilesResponse(val *ClientProfilesResponse) *NullableClientProfilesResponse {
	return &NullableClientProfilesResponse{value: val, isSet: true}
}

func (v NullableClientProfilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProfilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


