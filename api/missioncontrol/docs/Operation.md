# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the operation. | [optional] 
**Type** | Pointer to **string** | The type of object for informational purposes. | [optional] [readonly] 
**OperationType** | Pointer to **string** | The type of operation against the resource, such as &#39;createService&#39;, &#39;cloneService&#39;, &#39;deleteService&#39;, &#39;createClientProfile&#39;, &#39;updateClientProfile&#39;, and &#39;deleteClientProfile&#39;. | [optional] 
**CreatedBy** | Pointer to **string** | The unique identifier representing the user who created the operation. | [optional] 
**CreatedTime** | Pointer to **string** | The time the operation was created, in ISO 8601 date/time format. | [optional] 
**CompletedTime** | Pointer to **string** | The completion time, whether it was successful or failed, in ISO 8601 date/time format. | [optional] 
**ResourceId** | Pointer to **string** | The resource ID that the operation belongs to. | [optional] [readonly] 
**ResourceType** | Pointer to **string** | The resource type that the operation belongs to. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the operation. | [optional] [readonly] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewOperation

`func NewOperation() *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Operation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Operation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Operation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Operation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Operation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Operation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Operation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Operation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOperationType

`func (o *Operation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *Operation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *Operation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *Operation) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Operation) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Operation) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Operation) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Operation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Operation) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Operation) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Operation) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Operation) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCompletedTime

`func (o *Operation) GetCompletedTime() string`

GetCompletedTime returns the CompletedTime field if non-nil, zero value otherwise.

### GetCompletedTimeOk

`func (o *Operation) GetCompletedTimeOk() (*string, bool)`

GetCompletedTimeOk returns a tuple with the CompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTime

`func (o *Operation) SetCompletedTime(v string)`

SetCompletedTime sets CompletedTime field to given value.

### HasCompletedTime

`func (o *Operation) HasCompletedTime() bool`

HasCompletedTime returns a boolean if a field has been set.

### GetResourceId

`func (o *Operation) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Operation) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Operation) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Operation) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *Operation) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Operation) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Operation) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Operation) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *Operation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Operation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Operation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Operation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *Operation) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Operation) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Operation) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *Operation) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


