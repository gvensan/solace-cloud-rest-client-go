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

// checks if the DatacentersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatacentersResponse{}

// DatacentersResponse struct for DatacentersResponse
type DatacentersResponse struct {
	Data []Datacenter `json:"data,omitempty"`
	Meta map[string]map[string]interface{} `json:"meta,omitempty"`
}

// NewDatacentersResponse instantiates a new DatacentersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatacentersResponse() *DatacentersResponse {
	this := DatacentersResponse{}
	return &this
}

// NewDatacentersResponseWithDefaults instantiates a new DatacentersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatacentersResponseWithDefaults() *DatacentersResponse {
	this := DatacentersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DatacentersResponse) GetData() []Datacenter {
	if o == nil || IsNil(o.Data) {
		var ret []Datacenter
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacentersResponse) GetDataOk() ([]Datacenter, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DatacentersResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Datacenter and assigns it to the Data field.
func (o *DatacentersResponse) SetData(v []Datacenter) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DatacentersResponse) GetMeta() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacentersResponse) GetMetaOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DatacentersResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]map[string]interface{} and assigns it to the Meta field.
func (o *DatacentersResponse) SetMeta(v map[string]map[string]interface{}) {
	o.Meta = v
}

func (o DatacentersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatacentersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableDatacentersResponse struct {
	value *DatacentersResponse
	isSet bool
}

func (v NullableDatacentersResponse) Get() *DatacentersResponse {
	return v.value
}

func (v *NullableDatacentersResponse) Set(val *DatacentersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatacentersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatacentersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatacentersResponse(val *DatacentersResponse) *NullableDatacentersResponse {
	return &NullableDatacentersResponse{value: val, isSet: true}
}

func (v NullableDatacentersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatacentersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


