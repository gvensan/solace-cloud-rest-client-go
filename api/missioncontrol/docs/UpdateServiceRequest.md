# UpdateServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new service name. The new service name must be unique within an organization. | [optional] 
**OwnedBy** | Pointer to **string** | The owner of the event broker service. The owner must belong to the same organization. | [optional] 
**Locked** | Pointer to **bool** | Lock (true) and unlock (false) the event broker service. If an event broker service is locked it&#39;s protected from being deleted. | [optional] 

## Methods

### NewUpdateServiceRequest

`func NewUpdateServiceRequest() *UpdateServiceRequest`

NewUpdateServiceRequest instantiates a new UpdateServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceRequestWithDefaults

`func NewUpdateServiceRequestWithDefaults() *UpdateServiceRequest`

NewUpdateServiceRequestWithDefaults instantiates a new UpdateServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServiceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServiceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnedBy

`func (o *UpdateServiceRequest) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *UpdateServiceRequest) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *UpdateServiceRequest) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *UpdateServiceRequest) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetLocked

`func (o *UpdateServiceRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UpdateServiceRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UpdateServiceRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UpdateServiceRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


