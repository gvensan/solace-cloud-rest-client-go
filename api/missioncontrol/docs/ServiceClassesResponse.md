# ServiceClassesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ServiceClass**](ServiceClass.md) |  | [optional] 
**Meta** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewServiceClassesResponse

`func NewServiceClassesResponse() *ServiceClassesResponse`

NewServiceClassesResponse instantiates a new ServiceClassesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceClassesResponseWithDefaults

`func NewServiceClassesResponseWithDefaults() *ServiceClassesResponse`

NewServiceClassesResponseWithDefaults instantiates a new ServiceClassesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceClassesResponse) GetData() []ServiceClass`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceClassesResponse) GetDataOk() (*[]ServiceClass, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceClassesResponse) SetData(v []ServiceClass)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceClassesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ServiceClassesResponse) GetMeta() map[string]map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServiceClassesResponse) GetMetaOk() (*map[string]map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServiceClassesResponse) SetMeta(v map[string]map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ServiceClassesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


