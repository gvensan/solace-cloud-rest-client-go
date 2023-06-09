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

// checks if the ServiceClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceClass{}

// ServiceClass struct for ServiceClass
type ServiceClass struct {
	// The identifier of the service class.
	Id *string `json:"id,omitempty"`
	// The type of object for internal informational purposes.
	Type *string `json:"type,omitempty"`
	// The name of the service class.
	Name *string `json:"name,omitempty"`
	// The maximum number of client connections for the service class.
	VpnConnections *int32 `json:"vpnConnections,omitempty"`
	// The underlying scaling tiers for the software event broker.
	BrokerScalingTier *string `json:"brokerScalingTier,omitempty"`
	// The maximum message spool size of the service class, in gigabytes (GB).
	VpnMaxSpoolSize *int32 `json:"vpnMaxSpoolSize,omitempty"`
	// The maximum number of Message VPNs for the service class.
	MaxNumberVpns *int32 `json:"maxNumberVpns,omitempty"`
	// Whether the service class supports High-Availability. A value of 'True' indicates that high-availability is supported.
	HighAvailabilityCapable *bool `json:"highAvailabilityCapable,omitempty"`
}

// NewServiceClass instantiates a new ServiceClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceClass() *ServiceClass {
	this := ServiceClass{}
	return &this
}

// NewServiceClassWithDefaults instantiates a new ServiceClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceClassWithDefaults() *ServiceClass {
	this := ServiceClass{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceClass) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClass) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceClass) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceClass) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceClass) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClass) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceClass) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceClass) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceClass) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClass) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceClass) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceClass) SetName(v string) {
	o.Name = &v
}

// GetVpnConnections returns the VpnConnections field value if set, zero value otherwise.
func (o *ServiceClass) GetVpnConnections() int32 {
	if o == nil || IsNil(o.VpnConnections) {
		var ret int32
		return ret
	}
	return *o.VpnConnections
}

// GetVpnConnectionsOk returns a tuple with the VpnConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClass) GetVpnConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.VpnConnections) {
		return nil, false
	}
	return o.VpnConnections, true
}

// HasVpnConnections returns a boolean if a field has been set.
func (o *ServiceClass) HasVpnConnections() bool {
	if o != nil && !IsNil(o.VpnConnections) {
		return true
	}

	return false
}

// SetVpnConnections gets a reference to the given int32 and assigns it to the VpnConnections field.
func (o *ServiceClass) SetVpnConnections(v int32) {
	o.VpnConnections = &v
}

// GetBrokerScalingTier returns the BrokerScalingTier field value if set, zero value otherwise.
func (o *ServiceClass) GetBrokerScalingTier() string {
	if o == nil || IsNil(o.BrokerScalingTier) {
		var ret string
		return ret
	}
	return *o.BrokerScalingTier
}

// GetBrokerScalingTierOk returns a tuple with the BrokerScalingTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClass) GetBrokerScalingTierOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerScalingTier) {
		return nil, false
	}
	return o.BrokerScalingTier, true
}

// HasBrokerScalingTier returns a boolean if a field has been set.
func (o *ServiceClass) HasBrokerScalingTier() bool {
	if o != nil && !IsNil(o.BrokerScalingTier) {
		return true
	}

	return false
}

// SetBrokerScalingTier gets a reference to the given string and assigns it to the BrokerScalingTier field.
func (o *ServiceClass) SetBrokerScalingTier(v string) {
	o.BrokerScalingTier = &v
}

// GetVpnMaxSpoolSize returns the VpnMaxSpoolSize field value if set, zero value otherwise.
func (o *ServiceClass) GetVpnMaxSpoolSize() int32 {
	if o == nil || IsNil(o.VpnMaxSpoolSize) {
		var ret int32
		return ret
	}
	return *o.VpnMaxSpoolSize
}

// GetVpnMaxSpoolSizeOk returns a tuple with the VpnMaxSpoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClass) GetVpnMaxSpoolSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.VpnMaxSpoolSize) {
		return nil, false
	}
	return o.VpnMaxSpoolSize, true
}

// HasVpnMaxSpoolSize returns a boolean if a field has been set.
func (o *ServiceClass) HasVpnMaxSpoolSize() bool {
	if o != nil && !IsNil(o.VpnMaxSpoolSize) {
		return true
	}

	return false
}

// SetVpnMaxSpoolSize gets a reference to the given int32 and assigns it to the VpnMaxSpoolSize field.
func (o *ServiceClass) SetVpnMaxSpoolSize(v int32) {
	o.VpnMaxSpoolSize = &v
}

// GetMaxNumberVpns returns the MaxNumberVpns field value if set, zero value otherwise.
func (o *ServiceClass) GetMaxNumberVpns() int32 {
	if o == nil || IsNil(o.MaxNumberVpns) {
		var ret int32
		return ret
	}
	return *o.MaxNumberVpns
}

// GetMaxNumberVpnsOk returns a tuple with the MaxNumberVpns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClass) GetMaxNumberVpnsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumberVpns) {
		return nil, false
	}
	return o.MaxNumberVpns, true
}

// HasMaxNumberVpns returns a boolean if a field has been set.
func (o *ServiceClass) HasMaxNumberVpns() bool {
	if o != nil && !IsNil(o.MaxNumberVpns) {
		return true
	}

	return false
}

// SetMaxNumberVpns gets a reference to the given int32 and assigns it to the MaxNumberVpns field.
func (o *ServiceClass) SetMaxNumberVpns(v int32) {
	o.MaxNumberVpns = &v
}

// GetHighAvailabilityCapable returns the HighAvailabilityCapable field value if set, zero value otherwise.
func (o *ServiceClass) GetHighAvailabilityCapable() bool {
	if o == nil || IsNil(o.HighAvailabilityCapable) {
		var ret bool
		return ret
	}
	return *o.HighAvailabilityCapable
}

// GetHighAvailabilityCapableOk returns a tuple with the HighAvailabilityCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceClass) GetHighAvailabilityCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.HighAvailabilityCapable) {
		return nil, false
	}
	return o.HighAvailabilityCapable, true
}

// HasHighAvailabilityCapable returns a boolean if a field has been set.
func (o *ServiceClass) HasHighAvailabilityCapable() bool {
	if o != nil && !IsNil(o.HighAvailabilityCapable) {
		return true
	}

	return false
}

// SetHighAvailabilityCapable gets a reference to the given bool and assigns it to the HighAvailabilityCapable field.
func (o *ServiceClass) SetHighAvailabilityCapable(v bool) {
	o.HighAvailabilityCapable = &v
}

func (o ServiceClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.VpnConnections) {
		toSerialize["vpnConnections"] = o.VpnConnections
	}
	if !IsNil(o.BrokerScalingTier) {
		toSerialize["brokerScalingTier"] = o.BrokerScalingTier
	}
	if !IsNil(o.VpnMaxSpoolSize) {
		toSerialize["vpnMaxSpoolSize"] = o.VpnMaxSpoolSize
	}
	if !IsNil(o.MaxNumberVpns) {
		toSerialize["maxNumberVpns"] = o.MaxNumberVpns
	}
	if !IsNil(o.HighAvailabilityCapable) {
		toSerialize["highAvailabilityCapable"] = o.HighAvailabilityCapable
	}
	return toSerialize, nil
}

type NullableServiceClass struct {
	value *ServiceClass
	isSet bool
}

func (v NullableServiceClass) Get() *ServiceClass {
	return v.value
}

func (v *NullableServiceClass) Set(val *ServiceClass) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceClass) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceClass(val *ServiceClass) *NullableServiceClass {
	return &NullableServiceClass{value: val, isSet: true}
}

func (v NullableServiceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


