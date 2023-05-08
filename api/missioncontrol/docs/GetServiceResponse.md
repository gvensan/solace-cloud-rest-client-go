# GetServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetService**](GetService.md) |  | [optional] 
**Meta** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewGetServiceResponse

`func NewGetServiceResponse() *GetServiceResponse`

NewGetServiceResponse instantiates a new GetServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceResponseWithDefaults

`func NewGetServiceResponseWithDefaults() *GetServiceResponse`

NewGetServiceResponseWithDefaults instantiates a new GetServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetServiceResponse) GetData() GetService`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetServiceResponse) GetDataOk() (*GetService, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetServiceResponse) SetData(v GetService)`

SetData sets Data field to given value.

### HasData

`func (o *GetServiceResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetServiceResponse) GetMeta() map[string]map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetServiceResponse) GetMetaOk() (*map[string]map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetServiceResponse) SetMeta(v map[string]map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetServiceResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


