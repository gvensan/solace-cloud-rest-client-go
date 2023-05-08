# GetAllServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetServices**](GetServices.md) |  | [optional] 
**Meta** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewGetAllServicesResponse

`func NewGetAllServicesResponse() *GetAllServicesResponse`

NewGetAllServicesResponse instantiates a new GetAllServicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllServicesResponseWithDefaults

`func NewGetAllServicesResponseWithDefaults() *GetAllServicesResponse`

NewGetAllServicesResponseWithDefaults instantiates a new GetAllServicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAllServicesResponse) GetData() []GetServices`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllServicesResponse) GetDataOk() (*[]GetServices, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllServicesResponse) SetData(v []GetServices)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllServicesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetAllServicesResponse) GetMeta() map[string]map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAllServicesResponse) GetMetaOk() (*map[string]map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAllServicesResponse) SetMeta(v map[string]map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetAllServicesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


