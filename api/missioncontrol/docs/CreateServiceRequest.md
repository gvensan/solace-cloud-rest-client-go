# CreateServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The event broker service name. | 
**EventBrokerVersion** | Pointer to **string** | The event broker version. A default version is provided when this is not specified. | [optional] 
**MsgVpnName** | Pointer to **string** | The message VPN name. A default message VPN name is provided when this is not specified. | [optional] 
**MaxSpoolUsage** | Pointer to **int32** | The message spool size, in gigabytes (GB). A default message spool size is provided if this is not specified. | [optional] 
**ServiceClassId** | **string** | The identifier of the service class. | 
**DatacenterId** | **string** | The identifier of the datacenter. | 
**ClusterName** | Pointer to **string** | The name of the DMR cluster. | [optional] 
**Locked** | Pointer to **bool** | Lock service after creation. A locked service cannot be deleted. | [optional] [default to false]
**RedundancyGroupSslEnabled** | Pointer to **bool** | Whether redundancy group SSL is enabled. This is mate-link encryption. | [optional] [default to false]
**ServiceConnectionEndpoints** | Pointer to [**[]ConnectionEndpoint**](ConnectionEndpoint.md) | The collection of service connection endpoints. | [optional] 

## Methods

### NewCreateServiceRequest

`func NewCreateServiceRequest(name string, serviceClassId string, datacenterId string, ) *CreateServiceRequest`

NewCreateServiceRequest instantiates a new CreateServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceRequestWithDefaults

`func NewCreateServiceRequestWithDefaults() *CreateServiceRequest`

NewCreateServiceRequestWithDefaults instantiates a new CreateServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEventBrokerVersion

`func (o *CreateServiceRequest) GetEventBrokerVersion() string`

GetEventBrokerVersion returns the EventBrokerVersion field if non-nil, zero value otherwise.

### GetEventBrokerVersionOk

`func (o *CreateServiceRequest) GetEventBrokerVersionOk() (*string, bool)`

GetEventBrokerVersionOk returns a tuple with the EventBrokerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBrokerVersion

`func (o *CreateServiceRequest) SetEventBrokerVersion(v string)`

SetEventBrokerVersion sets EventBrokerVersion field to given value.

### HasEventBrokerVersion

`func (o *CreateServiceRequest) HasEventBrokerVersion() bool`

HasEventBrokerVersion returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *CreateServiceRequest) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *CreateServiceRequest) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *CreateServiceRequest) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *CreateServiceRequest) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetMaxSpoolUsage

`func (o *CreateServiceRequest) GetMaxSpoolUsage() int32`

GetMaxSpoolUsage returns the MaxSpoolUsage field if non-nil, zero value otherwise.

### GetMaxSpoolUsageOk

`func (o *CreateServiceRequest) GetMaxSpoolUsageOk() (*int32, bool)`

GetMaxSpoolUsageOk returns a tuple with the MaxSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpoolUsage

`func (o *CreateServiceRequest) SetMaxSpoolUsage(v int32)`

SetMaxSpoolUsage sets MaxSpoolUsage field to given value.

### HasMaxSpoolUsage

`func (o *CreateServiceRequest) HasMaxSpoolUsage() bool`

HasMaxSpoolUsage returns a boolean if a field has been set.

### GetServiceClassId

`func (o *CreateServiceRequest) GetServiceClassId() string`

GetServiceClassId returns the ServiceClassId field if non-nil, zero value otherwise.

### GetServiceClassIdOk

`func (o *CreateServiceRequest) GetServiceClassIdOk() (*string, bool)`

GetServiceClassIdOk returns a tuple with the ServiceClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClassId

`func (o *CreateServiceRequest) SetServiceClassId(v string)`

SetServiceClassId sets ServiceClassId field to given value.


### GetDatacenterId

`func (o *CreateServiceRequest) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *CreateServiceRequest) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *CreateServiceRequest) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetClusterName

`func (o *CreateServiceRequest) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CreateServiceRequest) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CreateServiceRequest) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *CreateServiceRequest) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetLocked

`func (o *CreateServiceRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CreateServiceRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CreateServiceRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CreateServiceRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetRedundancyGroupSslEnabled

`func (o *CreateServiceRequest) GetRedundancyGroupSslEnabled() bool`

GetRedundancyGroupSslEnabled returns the RedundancyGroupSslEnabled field if non-nil, zero value otherwise.

### GetRedundancyGroupSslEnabledOk

`func (o *CreateServiceRequest) GetRedundancyGroupSslEnabledOk() (*bool, bool)`

GetRedundancyGroupSslEnabledOk returns a tuple with the RedundancyGroupSslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyGroupSslEnabled

`func (o *CreateServiceRequest) SetRedundancyGroupSslEnabled(v bool)`

SetRedundancyGroupSslEnabled sets RedundancyGroupSslEnabled field to given value.

### HasRedundancyGroupSslEnabled

`func (o *CreateServiceRequest) HasRedundancyGroupSslEnabled() bool`

HasRedundancyGroupSslEnabled returns a boolean if a field has been set.

### GetServiceConnectionEndpoints

`func (o *CreateServiceRequest) GetServiceConnectionEndpoints() []ConnectionEndpoint`

GetServiceConnectionEndpoints returns the ServiceConnectionEndpoints field if non-nil, zero value otherwise.

### GetServiceConnectionEndpointsOk

`func (o *CreateServiceRequest) GetServiceConnectionEndpointsOk() (*[]ConnectionEndpoint, bool)`

GetServiceConnectionEndpointsOk returns a tuple with the ServiceConnectionEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnectionEndpoints

`func (o *CreateServiceRequest) SetServiceConnectionEndpoints(v []ConnectionEndpoint)`

SetServiceConnectionEndpoints sets ServiceConnectionEndpoints field to given value.

### HasServiceConnectionEndpoints

`func (o *CreateServiceRequest) HasServiceConnectionEndpoints() bool`

HasServiceConnectionEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


