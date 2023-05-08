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

// checks if the ClientProfileSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientProfileSummary{}

// ClientProfileSummary The client profile configured on the event broker service.
type ClientProfileSummary struct {
	// The name of the client profile.
	Name string `json:"name"`
	// The identifier of the client profile.
	Id *string `json:"id,omitempty"`
	// The type of object for informational purposes.
	Type *string `json:"type,omitempty"`
	// Indicates whether clients assigned the client profile are allowed to publish guaranteed messages. The valid values are 'true' (allowed) and 'false' (not allowed). The default is 'false'.
	AllowGuaranteedMsgSendEnabled *bool `json:"allowGuaranteedMsgSendEnabled,omitempty"`
	// Indicates whether clients assigned to the client profile are allowed to bind to topic endpoints or queues to receive guaranteed messages. The valid values are 'true' (allowed) or 'false' (not allowed). The default is 'false'.
	AllowGuaranteedMsgReceiveEnabled *bool `json:"allowGuaranteedMsgReceiveEnabled,omitempty"`
	// Indicates whether clients assigned to the client profile are allowed to establish transacted sessions or XA sessions. The valid values are 'true' (allowed) and 'false' (not allowed). The default is 'true'.
	AllowTransactedSessionsEnabled *bool `json:"allowTransactedSessionsEnabled,omitempty"`
	// Indicates whether clients assigned to the client profile are allowed to establish Dynamic Messaging Routing (DMR) links (or bridge links) from the current Message VPN to another Message VPN in a separate event broker service. The valid values are 'true' (allowed) and 'false' (not allowed). The default is 'false'.
	AllowBridgeConnectionsEnabled *bool `json:"allowBridgeConnectionsEnabled,omitempty"`
	// Indicates whether clients assigned to the client profile are allowed to create queues or topic endpoints. The valid values are 'true' (allowed) and 'false' (not allowed). The default is 'false'.
	AllowGuaranteedEndpointCreateEnabled *bool `json:"allowGuaranteedEndpointCreateEnabled,omitempty"`
}

// NewClientProfileSummary instantiates a new ClientProfileSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProfileSummary(name string) *ClientProfileSummary {
	this := ClientProfileSummary{}
	this.Name = name
	return &this
}

// NewClientProfileSummaryWithDefaults instantiates a new ClientProfileSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProfileSummaryWithDefaults() *ClientProfileSummary {
	this := ClientProfileSummary{}
	return &this
}

// GetName returns the Name field value
func (o *ClientProfileSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientProfileSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientProfileSummary) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientProfileSummary) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfileSummary) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientProfileSummary) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientProfileSummary) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClientProfileSummary) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfileSummary) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClientProfileSummary) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClientProfileSummary) SetType(v string) {
	o.Type = &v
}

// GetAllowGuaranteedMsgSendEnabled returns the AllowGuaranteedMsgSendEnabled field value if set, zero value otherwise.
func (o *ClientProfileSummary) GetAllowGuaranteedMsgSendEnabled() bool {
	if o == nil || IsNil(o.AllowGuaranteedMsgSendEnabled) {
		var ret bool
		return ret
	}
	return *o.AllowGuaranteedMsgSendEnabled
}

// GetAllowGuaranteedMsgSendEnabledOk returns a tuple with the AllowGuaranteedMsgSendEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfileSummary) GetAllowGuaranteedMsgSendEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowGuaranteedMsgSendEnabled) {
		return nil, false
	}
	return o.AllowGuaranteedMsgSendEnabled, true
}

// HasAllowGuaranteedMsgSendEnabled returns a boolean if a field has been set.
func (o *ClientProfileSummary) HasAllowGuaranteedMsgSendEnabled() bool {
	if o != nil && !IsNil(o.AllowGuaranteedMsgSendEnabled) {
		return true
	}

	return false
}

// SetAllowGuaranteedMsgSendEnabled gets a reference to the given bool and assigns it to the AllowGuaranteedMsgSendEnabled field.
func (o *ClientProfileSummary) SetAllowGuaranteedMsgSendEnabled(v bool) {
	o.AllowGuaranteedMsgSendEnabled = &v
}

// GetAllowGuaranteedMsgReceiveEnabled returns the AllowGuaranteedMsgReceiveEnabled field value if set, zero value otherwise.
func (o *ClientProfileSummary) GetAllowGuaranteedMsgReceiveEnabled() bool {
	if o == nil || IsNil(o.AllowGuaranteedMsgReceiveEnabled) {
		var ret bool
		return ret
	}
	return *o.AllowGuaranteedMsgReceiveEnabled
}

// GetAllowGuaranteedMsgReceiveEnabledOk returns a tuple with the AllowGuaranteedMsgReceiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfileSummary) GetAllowGuaranteedMsgReceiveEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowGuaranteedMsgReceiveEnabled) {
		return nil, false
	}
	return o.AllowGuaranteedMsgReceiveEnabled, true
}

// HasAllowGuaranteedMsgReceiveEnabled returns a boolean if a field has been set.
func (o *ClientProfileSummary) HasAllowGuaranteedMsgReceiveEnabled() bool {
	if o != nil && !IsNil(o.AllowGuaranteedMsgReceiveEnabled) {
		return true
	}

	return false
}

// SetAllowGuaranteedMsgReceiveEnabled gets a reference to the given bool and assigns it to the AllowGuaranteedMsgReceiveEnabled field.
func (o *ClientProfileSummary) SetAllowGuaranteedMsgReceiveEnabled(v bool) {
	o.AllowGuaranteedMsgReceiveEnabled = &v
}

// GetAllowTransactedSessionsEnabled returns the AllowTransactedSessionsEnabled field value if set, zero value otherwise.
func (o *ClientProfileSummary) GetAllowTransactedSessionsEnabled() bool {
	if o == nil || IsNil(o.AllowTransactedSessionsEnabled) {
		var ret bool
		return ret
	}
	return *o.AllowTransactedSessionsEnabled
}

// GetAllowTransactedSessionsEnabledOk returns a tuple with the AllowTransactedSessionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfileSummary) GetAllowTransactedSessionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowTransactedSessionsEnabled) {
		return nil, false
	}
	return o.AllowTransactedSessionsEnabled, true
}

// HasAllowTransactedSessionsEnabled returns a boolean if a field has been set.
func (o *ClientProfileSummary) HasAllowTransactedSessionsEnabled() bool {
	if o != nil && !IsNil(o.AllowTransactedSessionsEnabled) {
		return true
	}

	return false
}

// SetAllowTransactedSessionsEnabled gets a reference to the given bool and assigns it to the AllowTransactedSessionsEnabled field.
func (o *ClientProfileSummary) SetAllowTransactedSessionsEnabled(v bool) {
	o.AllowTransactedSessionsEnabled = &v
}

// GetAllowBridgeConnectionsEnabled returns the AllowBridgeConnectionsEnabled field value if set, zero value otherwise.
func (o *ClientProfileSummary) GetAllowBridgeConnectionsEnabled() bool {
	if o == nil || IsNil(o.AllowBridgeConnectionsEnabled) {
		var ret bool
		return ret
	}
	return *o.AllowBridgeConnectionsEnabled
}

// GetAllowBridgeConnectionsEnabledOk returns a tuple with the AllowBridgeConnectionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfileSummary) GetAllowBridgeConnectionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowBridgeConnectionsEnabled) {
		return nil, false
	}
	return o.AllowBridgeConnectionsEnabled, true
}

// HasAllowBridgeConnectionsEnabled returns a boolean if a field has been set.
func (o *ClientProfileSummary) HasAllowBridgeConnectionsEnabled() bool {
	if o != nil && !IsNil(o.AllowBridgeConnectionsEnabled) {
		return true
	}

	return false
}

// SetAllowBridgeConnectionsEnabled gets a reference to the given bool and assigns it to the AllowBridgeConnectionsEnabled field.
func (o *ClientProfileSummary) SetAllowBridgeConnectionsEnabled(v bool) {
	o.AllowBridgeConnectionsEnabled = &v
}

// GetAllowGuaranteedEndpointCreateEnabled returns the AllowGuaranteedEndpointCreateEnabled field value if set, zero value otherwise.
func (o *ClientProfileSummary) GetAllowGuaranteedEndpointCreateEnabled() bool {
	if o == nil || IsNil(o.AllowGuaranteedEndpointCreateEnabled) {
		var ret bool
		return ret
	}
	return *o.AllowGuaranteedEndpointCreateEnabled
}

// GetAllowGuaranteedEndpointCreateEnabledOk returns a tuple with the AllowGuaranteedEndpointCreateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfileSummary) GetAllowGuaranteedEndpointCreateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowGuaranteedEndpointCreateEnabled) {
		return nil, false
	}
	return o.AllowGuaranteedEndpointCreateEnabled, true
}

// HasAllowGuaranteedEndpointCreateEnabled returns a boolean if a field has been set.
func (o *ClientProfileSummary) HasAllowGuaranteedEndpointCreateEnabled() bool {
	if o != nil && !IsNil(o.AllowGuaranteedEndpointCreateEnabled) {
		return true
	}

	return false
}

// SetAllowGuaranteedEndpointCreateEnabled gets a reference to the given bool and assigns it to the AllowGuaranteedEndpointCreateEnabled field.
func (o *ClientProfileSummary) SetAllowGuaranteedEndpointCreateEnabled(v bool) {
	o.AllowGuaranteedEndpointCreateEnabled = &v
}

func (o ClientProfileSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientProfileSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AllowGuaranteedMsgSendEnabled) {
		toSerialize["allowGuaranteedMsgSendEnabled"] = o.AllowGuaranteedMsgSendEnabled
	}
	if !IsNil(o.AllowGuaranteedMsgReceiveEnabled) {
		toSerialize["allowGuaranteedMsgReceiveEnabled"] = o.AllowGuaranteedMsgReceiveEnabled
	}
	if !IsNil(o.AllowTransactedSessionsEnabled) {
		toSerialize["allowTransactedSessionsEnabled"] = o.AllowTransactedSessionsEnabled
	}
	if !IsNil(o.AllowBridgeConnectionsEnabled) {
		toSerialize["allowBridgeConnectionsEnabled"] = o.AllowBridgeConnectionsEnabled
	}
	if !IsNil(o.AllowGuaranteedEndpointCreateEnabled) {
		toSerialize["allowGuaranteedEndpointCreateEnabled"] = o.AllowGuaranteedEndpointCreateEnabled
	}
	return toSerialize, nil
}

type NullableClientProfileSummary struct {
	value *ClientProfileSummary
	isSet bool
}

func (v NullableClientProfileSummary) Get() *ClientProfileSummary {
	return v.value
}

func (v *NullableClientProfileSummary) Set(val *ClientProfileSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProfileSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProfileSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProfileSummary(val *ClientProfileSummary) *NullableClientProfileSummary {
	return &NullableClientProfileSummary{value: val, isSet: true}
}

func (v NullableClientProfileSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProfileSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


