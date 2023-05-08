/*
MISSION CONTROL

Documentation for Mission Control API

API version: V2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the GetServices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServices{}

// GetServices struct for GetServices
type GetServices struct {
	// The identifier of the event broker service.
	Id *string `json:"id,omitempty"`
	// The type of object for informational purposes.
	Type *string `json:"type,omitempty"`
	// The name of the event broker service.
	Name *string `json:"name,omitempty"`
	// The unique identifier representing the user who created the event broker service.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The time the event broker service was created, in ISO 8601 date/time format.
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	// The unique identifier representing the user who last updated the event broker service.
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// The time of the last update was performed on the event broker service, in ISO 8601 date/time format.
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
	// The unique identifier representing the user who owns the event broker service.
	OwnedBy *string `json:"ownedBy,omitempty"`
	// A unique identifier representing for the infrastructure of the event broker service.
	InfrastructureId *string `json:"infrastructureId,omitempty"`
	// The identifier of the datacenter.
	DatacenterId *string `json:"datacenterId,omitempty"`
	// The service class of the event broker service.
	ServiceClassId *string `json:"serviceClassId,omitempty"`
	// The identifier of the event mesh for which the event broker service belongs, if applicable.
	EventMeshId *string `json:"eventMeshId,omitempty"`
	// The operation identifiers for an ongoing operation on the event broker service.
	OngoingOperationIds []string `json:"ongoingOperationIds,omitempty"`
	// The administration state of the event broker service, which can be 'initial', 'start', 'stop' , or 'destroy'.
	AdminState *string `json:"adminState,omitempty"`
	// The creation state of the event broker service, one of 'pending', 'inProgress' , ' completed' , or 'failed'.
	CreationState *string `json:"creationState,omitempty"`
	// Whether the event broker service is locked. Locked event broker services cannot be deleted.
	Locked *bool `json:"locked,omitempty"`
}

// NewGetServices instantiates a new GetServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServices() *GetServices {
	this := GetServices{}
	return &this
}

// NewGetServicesWithDefaults instantiates a new GetServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServicesWithDefaults() *GetServices {
	this := GetServices{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetServices) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetServices) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetServices) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetServices) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetServices) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetServices) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetServices) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetServices) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetServices) SetName(v string) {
	o.Name = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *GetServices) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GetServices) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *GetServices) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *GetServices) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *GetServices) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *GetServices) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *GetServices) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *GetServices) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *GetServices) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *GetServices) GetUpdatedTime() time.Time {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *GetServices) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given time.Time and assigns it to the UpdatedTime field.
func (o *GetServices) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *GetServices) GetOwnedBy() string {
	if o == nil || IsNil(o.OwnedBy) {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetOwnedByOk() (*string, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *GetServices) HasOwnedBy() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *GetServices) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

// GetInfrastructureId returns the InfrastructureId field value if set, zero value otherwise.
func (o *GetServices) GetInfrastructureId() string {
	if o == nil || IsNil(o.InfrastructureId) {
		var ret string
		return ret
	}
	return *o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetInfrastructureIdOk() (*string, bool) {
	if o == nil || IsNil(o.InfrastructureId) {
		return nil, false
	}
	return o.InfrastructureId, true
}

// HasInfrastructureId returns a boolean if a field has been set.
func (o *GetServices) HasInfrastructureId() bool {
	if o != nil && !IsNil(o.InfrastructureId) {
		return true
	}

	return false
}

// SetInfrastructureId gets a reference to the given string and assigns it to the InfrastructureId field.
func (o *GetServices) SetInfrastructureId(v string) {
	o.InfrastructureId = &v
}

// GetDatacenterId returns the DatacenterId field value if set, zero value otherwise.
func (o *GetServices) GetDatacenterId() string {
	if o == nil || IsNil(o.DatacenterId) {
		var ret string
		return ret
	}
	return *o.DatacenterId
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetDatacenterIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatacenterId) {
		return nil, false
	}
	return o.DatacenterId, true
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *GetServices) HasDatacenterId() bool {
	if o != nil && !IsNil(o.DatacenterId) {
		return true
	}

	return false
}

// SetDatacenterId gets a reference to the given string and assigns it to the DatacenterId field.
func (o *GetServices) SetDatacenterId(v string) {
	o.DatacenterId = &v
}

// GetServiceClassId returns the ServiceClassId field value if set, zero value otherwise.
func (o *GetServices) GetServiceClassId() string {
	if o == nil || IsNil(o.ServiceClassId) {
		var ret string
		return ret
	}
	return *o.ServiceClassId
}

// GetServiceClassIdOk returns a tuple with the ServiceClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetServiceClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceClassId) {
		return nil, false
	}
	return o.ServiceClassId, true
}

// HasServiceClassId returns a boolean if a field has been set.
func (o *GetServices) HasServiceClassId() bool {
	if o != nil && !IsNil(o.ServiceClassId) {
		return true
	}

	return false
}

// SetServiceClassId gets a reference to the given string and assigns it to the ServiceClassId field.
func (o *GetServices) SetServiceClassId(v string) {
	o.ServiceClassId = &v
}

// GetEventMeshId returns the EventMeshId field value if set, zero value otherwise.
func (o *GetServices) GetEventMeshId() string {
	if o == nil || IsNil(o.EventMeshId) {
		var ret string
		return ret
	}
	return *o.EventMeshId
}

// GetEventMeshIdOk returns a tuple with the EventMeshId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetEventMeshIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventMeshId) {
		return nil, false
	}
	return o.EventMeshId, true
}

// HasEventMeshId returns a boolean if a field has been set.
func (o *GetServices) HasEventMeshId() bool {
	if o != nil && !IsNil(o.EventMeshId) {
		return true
	}

	return false
}

// SetEventMeshId gets a reference to the given string and assigns it to the EventMeshId field.
func (o *GetServices) SetEventMeshId(v string) {
	o.EventMeshId = &v
}

// GetOngoingOperationIds returns the OngoingOperationIds field value if set, zero value otherwise.
func (o *GetServices) GetOngoingOperationIds() []string {
	if o == nil || IsNil(o.OngoingOperationIds) {
		var ret []string
		return ret
	}
	return o.OngoingOperationIds
}

// GetOngoingOperationIdsOk returns a tuple with the OngoingOperationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetOngoingOperationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OngoingOperationIds) {
		return nil, false
	}
	return o.OngoingOperationIds, true
}

// HasOngoingOperationIds returns a boolean if a field has been set.
func (o *GetServices) HasOngoingOperationIds() bool {
	if o != nil && !IsNil(o.OngoingOperationIds) {
		return true
	}

	return false
}

// SetOngoingOperationIds gets a reference to the given []string and assigns it to the OngoingOperationIds field.
func (o *GetServices) SetOngoingOperationIds(v []string) {
	o.OngoingOperationIds = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *GetServices) GetAdminState() string {
	if o == nil || IsNil(o.AdminState) {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetAdminStateOk() (*string, bool) {
	if o == nil || IsNil(o.AdminState) {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *GetServices) HasAdminState() bool {
	if o != nil && !IsNil(o.AdminState) {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *GetServices) SetAdminState(v string) {
	o.AdminState = &v
}

// GetCreationState returns the CreationState field value if set, zero value otherwise.
func (o *GetServices) GetCreationState() string {
	if o == nil || IsNil(o.CreationState) {
		var ret string
		return ret
	}
	return *o.CreationState
}

// GetCreationStateOk returns a tuple with the CreationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetCreationStateOk() (*string, bool) {
	if o == nil || IsNil(o.CreationState) {
		return nil, false
	}
	return o.CreationState, true
}

// HasCreationState returns a boolean if a field has been set.
func (o *GetServices) HasCreationState() bool {
	if o != nil && !IsNil(o.CreationState) {
		return true
	}

	return false
}

// SetCreationState gets a reference to the given string and assigns it to the CreationState field.
func (o *GetServices) SetCreationState(v string) {
	o.CreationState = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *GetServices) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServices) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *GetServices) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *GetServices) SetLocked(v bool) {
	o.Locked = &v
}

func (o GetServices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServices) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updatedTime"] = o.UpdatedTime
	}
	if !IsNil(o.OwnedBy) {
		toSerialize["ownedBy"] = o.OwnedBy
	}
	if !IsNil(o.InfrastructureId) {
		toSerialize["infrastructureId"] = o.InfrastructureId
	}
	if !IsNil(o.DatacenterId) {
		toSerialize["datacenterId"] = o.DatacenterId
	}
	if !IsNil(o.ServiceClassId) {
		toSerialize["serviceClassId"] = o.ServiceClassId
	}
	if !IsNil(o.EventMeshId) {
		toSerialize["eventMeshId"] = o.EventMeshId
	}
	if !IsNil(o.OngoingOperationIds) {
		toSerialize["ongoingOperationIds"] = o.OngoingOperationIds
	}
	if !IsNil(o.AdminState) {
		toSerialize["adminState"] = o.AdminState
	}
	if !IsNil(o.CreationState) {
		toSerialize["creationState"] = o.CreationState
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	return toSerialize, nil
}

type NullableGetServices struct {
	value *GetServices
	isSet bool
}

func (v NullableGetServices) Get() *GetServices {
	return v.value
}

func (v *NullableGetServices) Set(val *GetServices) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServices) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServices(val *GetServices) *NullableGetServices {
	return &NullableGetServices{value: val, isSet: true}
}

func (v NullableGetServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


