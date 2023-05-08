# GetServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the event broker service. | [optional] 
**Type** | Pointer to **string** | The type of object for informational purposes. | [optional] 
**Name** | Pointer to **string** | The name of the event broker service. | [optional] 
**CreatedBy** | Pointer to **string** | The unique identifier representing the user who created the event broker service. | [optional] 
**CreatedTime** | Pointer to **time.Time** | The time the event broker service was created, in ISO 8601 date/time format. | [optional] 
**UpdatedBy** | Pointer to **string** | The unique identifier representing the user who last updated the event broker service. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time of the last update was performed on the event broker service, in ISO 8601 date/time format. | [optional] 
**OwnedBy** | Pointer to **string** | The unique identifier representing the user who owns the event broker service. | [optional] 
**InfrastructureId** | Pointer to **string** | A unique identifier representing for the infrastructure of the event broker service. | [optional] 
**DatacenterId** | Pointer to **string** | The identifier of the datacenter. | [optional] 
**ServiceClassId** | Pointer to **string** | The service class of the event broker service. | [optional] 
**EventMeshId** | Pointer to **string** | The identifier of the event mesh for which the event broker service belongs, if applicable. | [optional] 
**OngoingOperationIds** | Pointer to **[]string** | The operation identifiers for an ongoing operation on the event broker service. | [optional] 
**AdminState** | Pointer to **string** | The administration state of the event broker service, which can be &#39;initial&#39;, &#39;start&#39;, &#39;stop&#39; , or &#39;destroy&#39;. | [optional] 
**CreationState** | Pointer to **string** | The creation state of the event broker service, one of &#39;pending&#39;, &#39;inProgress&#39; , &#39; completed&#39; , or &#39;failed&#39;. | [optional] 
**Locked** | Pointer to **bool** | Whether the event broker service is locked. Locked event broker services cannot be deleted. | [optional] 

## Methods

### NewGetServices

`func NewGetServices() *GetServices`

NewGetServices instantiates a new GetServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServicesWithDefaults

`func NewGetServicesWithDefaults() *GetServices`

NewGetServicesWithDefaults instantiates a new GetServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetServices) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServices) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServices) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetServices) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetServices) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetServices) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetServices) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetServices) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetServices) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetServices) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetServices) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetServices) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetServices) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetServices) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetServices) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetServices) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedTime

`func (o *GetServices) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GetServices) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GetServices) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *GetServices) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetServices) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetServices) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetServices) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetServices) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *GetServices) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *GetServices) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *GetServices) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *GetServices) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetOwnedBy

`func (o *GetServices) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *GetServices) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *GetServices) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *GetServices) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *GetServices) GetInfrastructureId() string`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *GetServices) GetInfrastructureIdOk() (*string, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *GetServices) SetInfrastructureId(v string)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *GetServices) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### GetDatacenterId

`func (o *GetServices) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *GetServices) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *GetServices) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *GetServices) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetServiceClassId

`func (o *GetServices) GetServiceClassId() string`

GetServiceClassId returns the ServiceClassId field if non-nil, zero value otherwise.

### GetServiceClassIdOk

`func (o *GetServices) GetServiceClassIdOk() (*string, bool)`

GetServiceClassIdOk returns a tuple with the ServiceClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClassId

`func (o *GetServices) SetServiceClassId(v string)`

SetServiceClassId sets ServiceClassId field to given value.

### HasServiceClassId

`func (o *GetServices) HasServiceClassId() bool`

HasServiceClassId returns a boolean if a field has been set.

### GetEventMeshId

`func (o *GetServices) GetEventMeshId() string`

GetEventMeshId returns the EventMeshId field if non-nil, zero value otherwise.

### GetEventMeshIdOk

`func (o *GetServices) GetEventMeshIdOk() (*string, bool)`

GetEventMeshIdOk returns a tuple with the EventMeshId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMeshId

`func (o *GetServices) SetEventMeshId(v string)`

SetEventMeshId sets EventMeshId field to given value.

### HasEventMeshId

`func (o *GetServices) HasEventMeshId() bool`

HasEventMeshId returns a boolean if a field has been set.

### GetOngoingOperationIds

`func (o *GetServices) GetOngoingOperationIds() []string`

GetOngoingOperationIds returns the OngoingOperationIds field if non-nil, zero value otherwise.

### GetOngoingOperationIdsOk

`func (o *GetServices) GetOngoingOperationIdsOk() (*[]string, bool)`

GetOngoingOperationIdsOk returns a tuple with the OngoingOperationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoingOperationIds

`func (o *GetServices) SetOngoingOperationIds(v []string)`

SetOngoingOperationIds sets OngoingOperationIds field to given value.

### HasOngoingOperationIds

`func (o *GetServices) HasOngoingOperationIds() bool`

HasOngoingOperationIds returns a boolean if a field has been set.

### GetAdminState

`func (o *GetServices) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *GetServices) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *GetServices) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *GetServices) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetCreationState

`func (o *GetServices) GetCreationState() string`

GetCreationState returns the CreationState field if non-nil, zero value otherwise.

### GetCreationStateOk

`func (o *GetServices) GetCreationStateOk() (*string, bool)`

GetCreationStateOk returns a tuple with the CreationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationState

`func (o *GetServices) SetCreationState(v string)`

SetCreationState sets CreationState field to given value.

### HasCreationState

`func (o *GetServices) HasCreationState() bool`

HasCreationState returns a boolean if a field has been set.

### GetLocked

`func (o *GetServices) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *GetServices) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *GetServices) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *GetServices) HasLocked() bool`

HasLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


