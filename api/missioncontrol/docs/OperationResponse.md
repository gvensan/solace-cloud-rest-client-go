# OperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Meta** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOperationResponse

`func NewOperationResponse() *OperationResponse`

NewOperationResponse instantiates a new OperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseWithDefaults

`func NewOperationResponseWithDefaults() *OperationResponse`

NewOperationResponseWithDefaults instantiates a new OperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OperationResponse) GetData() Operation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OperationResponse) GetDataOk() (*Operation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OperationResponse) SetData(v Operation)`

SetData sets Data field to given value.

### HasData

`func (o *OperationResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *OperationResponse) GetMeta() map[string]map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OperationResponse) GetMetaOk() (*map[string]map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OperationResponse) SetMeta(v map[string]map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OperationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


