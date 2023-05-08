# ServiceClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the service class. | [optional] 
**Type** | Pointer to **string** | The type of object for internal informational purposes. | [optional] 
**Name** | Pointer to **string** | The name of the service class. | [optional] 
**VpnConnections** | Pointer to **int32** | The maximum number of client connections for the service class. | [optional] 
**BrokerScalingTier** | Pointer to **string** | The underlying scaling tiers for the software event broker. | [optional] 
**VpnMaxSpoolSize** | Pointer to **int32** | The maximum message spool size of the service class, in gigabytes (GB). | [optional] 
**MaxNumberVpns** | Pointer to **int32** | The maximum number of Message VPNs for the service class. | [optional] 
**HighAvailabilityCapable** | Pointer to **bool** | Whether the service class supports High-Availability. A value of &#39;True&#39; indicates that high-availability is supported. | [optional] 

## Methods

### NewServiceClass

`func NewServiceClass() *ServiceClass`

NewServiceClass instantiates a new ServiceClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceClassWithDefaults

`func NewServiceClassWithDefaults() *ServiceClass`

NewServiceClassWithDefaults instantiates a new ServiceClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceClass) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ServiceClass) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceClass) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceClass) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceClass) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ServiceClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceClass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceClass) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVpnConnections

`func (o *ServiceClass) GetVpnConnections() int32`

GetVpnConnections returns the VpnConnections field if non-nil, zero value otherwise.

### GetVpnConnectionsOk

`func (o *ServiceClass) GetVpnConnectionsOk() (*int32, bool)`

GetVpnConnectionsOk returns a tuple with the VpnConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnections

`func (o *ServiceClass) SetVpnConnections(v int32)`

SetVpnConnections sets VpnConnections field to given value.

### HasVpnConnections

`func (o *ServiceClass) HasVpnConnections() bool`

HasVpnConnections returns a boolean if a field has been set.

### GetBrokerScalingTier

`func (o *ServiceClass) GetBrokerScalingTier() string`

GetBrokerScalingTier returns the BrokerScalingTier field if non-nil, zero value otherwise.

### GetBrokerScalingTierOk

`func (o *ServiceClass) GetBrokerScalingTierOk() (*string, bool)`

GetBrokerScalingTierOk returns a tuple with the BrokerScalingTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerScalingTier

`func (o *ServiceClass) SetBrokerScalingTier(v string)`

SetBrokerScalingTier sets BrokerScalingTier field to given value.

### HasBrokerScalingTier

`func (o *ServiceClass) HasBrokerScalingTier() bool`

HasBrokerScalingTier returns a boolean if a field has been set.

### GetVpnMaxSpoolSize

`func (o *ServiceClass) GetVpnMaxSpoolSize() int32`

GetVpnMaxSpoolSize returns the VpnMaxSpoolSize field if non-nil, zero value otherwise.

### GetVpnMaxSpoolSizeOk

`func (o *ServiceClass) GetVpnMaxSpoolSizeOk() (*int32, bool)`

GetVpnMaxSpoolSizeOk returns a tuple with the VpnMaxSpoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnMaxSpoolSize

`func (o *ServiceClass) SetVpnMaxSpoolSize(v int32)`

SetVpnMaxSpoolSize sets VpnMaxSpoolSize field to given value.

### HasVpnMaxSpoolSize

`func (o *ServiceClass) HasVpnMaxSpoolSize() bool`

HasVpnMaxSpoolSize returns a boolean if a field has been set.

### GetMaxNumberVpns

`func (o *ServiceClass) GetMaxNumberVpns() int32`

GetMaxNumberVpns returns the MaxNumberVpns field if non-nil, zero value otherwise.

### GetMaxNumberVpnsOk

`func (o *ServiceClass) GetMaxNumberVpnsOk() (*int32, bool)`

GetMaxNumberVpnsOk returns a tuple with the MaxNumberVpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberVpns

`func (o *ServiceClass) SetMaxNumberVpns(v int32)`

SetMaxNumberVpns sets MaxNumberVpns field to given value.

### HasMaxNumberVpns

`func (o *ServiceClass) HasMaxNumberVpns() bool`

HasMaxNumberVpns returns a boolean if a field has been set.

### GetHighAvailabilityCapable

`func (o *ServiceClass) GetHighAvailabilityCapable() bool`

GetHighAvailabilityCapable returns the HighAvailabilityCapable field if non-nil, zero value otherwise.

### GetHighAvailabilityCapableOk

`func (o *ServiceClass) GetHighAvailabilityCapableOk() (*bool, bool)`

GetHighAvailabilityCapableOk returns a tuple with the HighAvailabilityCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailabilityCapable

`func (o *ServiceClass) SetHighAvailabilityCapable(v bool)`

SetHighAvailabilityCapable sets HighAvailabilityCapable field to given value.

### HasHighAvailabilityCapable

`func (o *ServiceClass) HasHighAvailabilityCapable() bool`

HasHighAvailabilityCapable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


