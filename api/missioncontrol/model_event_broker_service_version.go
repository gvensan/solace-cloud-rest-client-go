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

// checks if the EventBrokerServiceVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventBrokerServiceVersion{}

// EventBrokerServiceVersion The event broker service version.
type EventBrokerServiceVersion struct {
	// Event broker service version.
	Version string `json:"version"`
	// Supported service classes.
	SupportedServiceClasses []string `json:"supportedServiceClasses"`
	// Event broker service capabilities.
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
}

// NewEventBrokerServiceVersion instantiates a new EventBrokerServiceVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventBrokerServiceVersion(version string, supportedServiceClasses []string) *EventBrokerServiceVersion {
	this := EventBrokerServiceVersion{}
	this.Version = version
	this.SupportedServiceClasses = supportedServiceClasses
	return &this
}

// NewEventBrokerServiceVersionWithDefaults instantiates a new EventBrokerServiceVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventBrokerServiceVersionWithDefaults() *EventBrokerServiceVersion {
	this := EventBrokerServiceVersion{}
	return &this
}

// GetVersion returns the Version field value
func (o *EventBrokerServiceVersion) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EventBrokerServiceVersion) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EventBrokerServiceVersion) SetVersion(v string) {
	o.Version = v
}

// GetSupportedServiceClasses returns the SupportedServiceClasses field value
func (o *EventBrokerServiceVersion) GetSupportedServiceClasses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SupportedServiceClasses
}

// GetSupportedServiceClassesOk returns a tuple with the SupportedServiceClasses field value
// and a boolean to check if the value has been set.
func (o *EventBrokerServiceVersion) GetSupportedServiceClassesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedServiceClasses, true
}

// SetSupportedServiceClasses sets field value
func (o *EventBrokerServiceVersion) SetSupportedServiceClasses(v []string) {
	o.SupportedServiceClasses = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *EventBrokerServiceVersion) GetCapabilities() map[string]interface{} {
	if o == nil || IsNil(o.Capabilities) {
		var ret map[string]interface{}
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBrokerServiceVersion) GetCapabilitiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return map[string]interface{}{}, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *EventBrokerServiceVersion) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given map[string]interface{} and assigns it to the Capabilities field.
func (o *EventBrokerServiceVersion) SetCapabilities(v map[string]interface{}) {
	o.Capabilities = v
}

func (o EventBrokerServiceVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventBrokerServiceVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["supportedServiceClasses"] = o.SupportedServiceClasses
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	return toSerialize, nil
}

type NullableEventBrokerServiceVersion struct {
	value *EventBrokerServiceVersion
	isSet bool
}

func (v NullableEventBrokerServiceVersion) Get() *EventBrokerServiceVersion {
	return v.value
}

func (v *NullableEventBrokerServiceVersion) Set(val *EventBrokerServiceVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableEventBrokerServiceVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableEventBrokerServiceVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventBrokerServiceVersion(val *EventBrokerServiceVersion) *NullableEventBrokerServiceVersion {
	return &NullableEventBrokerServiceVersion{value: val, isSet: true}
}

func (v NullableEventBrokerServiceVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventBrokerServiceVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


