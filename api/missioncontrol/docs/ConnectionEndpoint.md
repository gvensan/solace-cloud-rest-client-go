# ConnectionEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the connection endpoint. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of object for informational purposes. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the connection endpoint. | [optional] 
**Description** | Pointer to **string** | The description for the connection endpoint. | [optional] 
**AccessType** | Pointer to **string** | The connectivity for the connection endpoint. This can be through private IP addresses (Private) or public Internet (Public). | [optional] 
**K8sServiceType** | Pointer to **string** | The connectivity configuration that is used in the Kubernetes cluster. | [optional] [readonly] 
**K8sServiceId** | Pointer to **string** | The identifier for the Kubernetes service. | [optional] [readonly] 
**HostNames** | Pointer to **[]string** | The hostnames assigned to the connection endpoint. | [optional] [readonly] 
**Ports** | Pointer to **map[string]int32** | The names and port numbers for the connection endpoint. | [optional] 

## Methods

### NewConnectionEndpoint

`func NewConnectionEndpoint() *ConnectionEndpoint`

NewConnectionEndpoint instantiates a new ConnectionEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionEndpointWithDefaults

`func NewConnectionEndpointWithDefaults() *ConnectionEndpoint`

NewConnectionEndpointWithDefaults instantiates a new ConnectionEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ConnectionEndpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionEndpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionEndpoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectionEndpoint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ConnectionEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectionEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionEndpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessType

`func (o *ConnectionEndpoint) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *ConnectionEndpoint) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *ConnectionEndpoint) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *ConnectionEndpoint) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetK8sServiceType

`func (o *ConnectionEndpoint) GetK8sServiceType() string`

GetK8sServiceType returns the K8sServiceType field if non-nil, zero value otherwise.

### GetK8sServiceTypeOk

`func (o *ConnectionEndpoint) GetK8sServiceTypeOk() (*string, bool)`

GetK8sServiceTypeOk returns a tuple with the K8sServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sServiceType

`func (o *ConnectionEndpoint) SetK8sServiceType(v string)`

SetK8sServiceType sets K8sServiceType field to given value.

### HasK8sServiceType

`func (o *ConnectionEndpoint) HasK8sServiceType() bool`

HasK8sServiceType returns a boolean if a field has been set.

### GetK8sServiceId

`func (o *ConnectionEndpoint) GetK8sServiceId() string`

GetK8sServiceId returns the K8sServiceId field if non-nil, zero value otherwise.

### GetK8sServiceIdOk

`func (o *ConnectionEndpoint) GetK8sServiceIdOk() (*string, bool)`

GetK8sServiceIdOk returns a tuple with the K8sServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sServiceId

`func (o *ConnectionEndpoint) SetK8sServiceId(v string)`

SetK8sServiceId sets K8sServiceId field to given value.

### HasK8sServiceId

`func (o *ConnectionEndpoint) HasK8sServiceId() bool`

HasK8sServiceId returns a boolean if a field has been set.

### GetHostNames

`func (o *ConnectionEndpoint) GetHostNames() []string`

GetHostNames returns the HostNames field if non-nil, zero value otherwise.

### GetHostNamesOk

`func (o *ConnectionEndpoint) GetHostNamesOk() (*[]string, bool)`

GetHostNamesOk returns a tuple with the HostNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNames

`func (o *ConnectionEndpoint) SetHostNames(v []string)`

SetHostNames sets HostNames field to given value.

### HasHostNames

`func (o *ConnectionEndpoint) HasHostNames() bool`

HasHostNames returns a boolean if a field has been set.

### GetPorts

`func (o *ConnectionEndpoint) GetPorts() map[string]int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ConnectionEndpoint) GetPortsOk() (*map[string]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ConnectionEndpoint) SetPorts(v map[string]int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ConnectionEndpoint) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


