# DatacenterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Datacenter**](Datacenter.md) |  | [optional] 
**Meta** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewDatacenterResponse

`func NewDatacenterResponse() *DatacenterResponse`

NewDatacenterResponse instantiates a new DatacenterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterResponseWithDefaults

`func NewDatacenterResponseWithDefaults() *DatacenterResponse`

NewDatacenterResponseWithDefaults instantiates a new DatacenterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DatacenterResponse) GetData() Datacenter`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatacenterResponse) GetDataOk() (*Datacenter, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatacenterResponse) SetData(v Datacenter)`

SetData sets Data field to given value.

### HasData

`func (o *DatacenterResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *DatacenterResponse) GetMeta() map[string]map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DatacenterResponse) GetMetaOk() (*map[string]map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DatacenterResponse) SetMeta(v map[string]map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DatacenterResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


