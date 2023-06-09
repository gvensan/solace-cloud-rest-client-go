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

// checks if the ConnectionEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionEndpoint{}

// ConnectionEndpoint The connection endpoint.
type ConnectionEndpoint struct {
	// The identifier of the connection endpoint.
	Id *string `json:"id,omitempty"`
	// The type of object for informational purposes.
	Type *string `json:"type,omitempty"`
	// The name of the connection endpoint.
	Name *string `json:"name,omitempty"`
	// The description for the connection endpoint.
	Description *string `json:"description,omitempty"`
	// The connectivity for the connection endpoint. This can be through private IP addresses (Private) or public Internet (Public).
	AccessType *string `json:"accessType,omitempty"`
	// The connectivity configuration that is used in the Kubernetes cluster.
	K8sServiceType *string `json:"k8sServiceType,omitempty"`
	// The identifier for the Kubernetes service.
	K8sServiceId *string `json:"k8sServiceId,omitempty"`
	// The hostnames assigned to the connection endpoint.
	HostNames []string `json:"hostNames,omitempty"`
	// The names and port numbers for the connection endpoint.
	Ports *map[string]int32 `json:"ports,omitempty"`
}

// NewConnectionEndpoint instantiates a new ConnectionEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionEndpoint() *ConnectionEndpoint {
	this := ConnectionEndpoint{}
	return &this
}

// NewConnectionEndpointWithDefaults instantiates a new ConnectionEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionEndpointWithDefaults() *ConnectionEndpoint {
	this := ConnectionEndpoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectionEndpoint) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEndpoint) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectionEndpoint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectionEndpoint) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConnectionEndpoint) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEndpoint) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConnectionEndpoint) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConnectionEndpoint) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectionEndpoint) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEndpoint) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectionEndpoint) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectionEndpoint) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectionEndpoint) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEndpoint) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectionEndpoint) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectionEndpoint) SetDescription(v string) {
	o.Description = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *ConnectionEndpoint) GetAccessType() string {
	if o == nil || IsNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEndpoint) GetAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *ConnectionEndpoint) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *ConnectionEndpoint) SetAccessType(v string) {
	o.AccessType = &v
}

// GetK8sServiceType returns the K8sServiceType field value if set, zero value otherwise.
func (o *ConnectionEndpoint) GetK8sServiceType() string {
	if o == nil || IsNil(o.K8sServiceType) {
		var ret string
		return ret
	}
	return *o.K8sServiceType
}

// GetK8sServiceTypeOk returns a tuple with the K8sServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEndpoint) GetK8sServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.K8sServiceType) {
		return nil, false
	}
	return o.K8sServiceType, true
}

// HasK8sServiceType returns a boolean if a field has been set.
func (o *ConnectionEndpoint) HasK8sServiceType() bool {
	if o != nil && !IsNil(o.K8sServiceType) {
		return true
	}

	return false
}

// SetK8sServiceType gets a reference to the given string and assigns it to the K8sServiceType field.
func (o *ConnectionEndpoint) SetK8sServiceType(v string) {
	o.K8sServiceType = &v
}

// GetK8sServiceId returns the K8sServiceId field value if set, zero value otherwise.
func (o *ConnectionEndpoint) GetK8sServiceId() string {
	if o == nil || IsNil(o.K8sServiceId) {
		var ret string
		return ret
	}
	return *o.K8sServiceId
}

// GetK8sServiceIdOk returns a tuple with the K8sServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEndpoint) GetK8sServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.K8sServiceId) {
		return nil, false
	}
	return o.K8sServiceId, true
}

// HasK8sServiceId returns a boolean if a field has been set.
func (o *ConnectionEndpoint) HasK8sServiceId() bool {
	if o != nil && !IsNil(o.K8sServiceId) {
		return true
	}

	return false
}

// SetK8sServiceId gets a reference to the given string and assigns it to the K8sServiceId field.
func (o *ConnectionEndpoint) SetK8sServiceId(v string) {
	o.K8sServiceId = &v
}

// GetHostNames returns the HostNames field value if set, zero value otherwise.
func (o *ConnectionEndpoint) GetHostNames() []string {
	if o == nil || IsNil(o.HostNames) {
		var ret []string
		return ret
	}
	return o.HostNames
}

// GetHostNamesOk returns a tuple with the HostNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEndpoint) GetHostNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.HostNames) {
		return nil, false
	}
	return o.HostNames, true
}

// HasHostNames returns a boolean if a field has been set.
func (o *ConnectionEndpoint) HasHostNames() bool {
	if o != nil && !IsNil(o.HostNames) {
		return true
	}

	return false
}

// SetHostNames gets a reference to the given []string and assigns it to the HostNames field.
func (o *ConnectionEndpoint) SetHostNames(v []string) {
	o.HostNames = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ConnectionEndpoint) GetPorts() map[string]int32 {
	if o == nil || IsNil(o.Ports) {
		var ret map[string]int32
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEndpoint) GetPortsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ConnectionEndpoint) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given map[string]int32 and assigns it to the Ports field.
func (o *ConnectionEndpoint) SetPorts(v map[string]int32) {
	o.Ports = &v
}

func (o ConnectionEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: type is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	// skip: k8sServiceType is readOnly
	// skip: k8sServiceId is readOnly
	// skip: hostNames is readOnly
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

type NullableConnectionEndpoint struct {
	value *ConnectionEndpoint
	isSet bool
}

func (v NullableConnectionEndpoint) Get() *ConnectionEndpoint {
	return v.value
}

func (v *NullableConnectionEndpoint) Set(val *ConnectionEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionEndpoint(val *ConnectionEndpoint) *NullableConnectionEndpoint {
	return &NullableConnectionEndpoint{value: val, isSet: true}
}

func (v NullableConnectionEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


