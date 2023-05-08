# Datacenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the datacenter. | [optional] 
**Type** | Pointer to **string** | The type of object for informational purposes. | [optional] 
**Name** | **string** | The name of the datacenter. | 
**DatacenterType** | **string** | The type of the datacenter, in terms of ownership. The valus can be a Public Region( &#39;SolacePublic&#39;), a Dedicated Region(&#39;SolaceDedicated&#39;), or a Customer-Controlled Region (&#39;CustomerCloud&#39; or &#39;CustomerOnPrem&#39;). | 
**Provider** | **string** | The name of the cloud provider for the datacenter. | 
**OperState** | **string** | The operational state of the datacenter. The values can be &#39;up&#39; or &#39;down&#39;. | 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**CreatedBy** | Pointer to **string** | The unique identifier representing the user who created the datacenter. | [optional] 
**CreatedTime** | Pointer to **time.Time** | The time the datacenter was created, in ISO 8601 date/time format. | [optional] 
**UpdatedBy** | Pointer to **string** | The unique identifier representing the user who last updated the datacenter. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time of the last update was performed on the datacenter, in ISO 8601 date/time format. | [optional] 
**Available** | **bool** | Whether the datacenter is available. | 
**SupportedServiceClasses** | Pointer to **[]string** | The list of supported service classes in the datacenter. | [optional] 
**CloudAgentVersion** | Pointer to **string** | The version of the Mission Control Agent. | [optional] 
**K8sServiceType** | Pointer to **string** | The type of the Kubernetes service. The values can be &#39;LoadBalancer&#39; or &#39;NodePort&#39;. | [optional] 
**NumSupportedPrivateEndpoints** | Pointer to **int32** | The number of supported private endpoints. | [optional] 
**NumSupportedPublicEndpoints** | Pointer to **int32** | The number of supported public endpoints. | [optional] 
**OrganizationId** | Pointer to **string** | The identifier of the datacenter&#39;s organization. | [optional] 
**SpoolScaleUpCapabilityInfo** | Pointer to [**SpoolScaleUpCapabilityInfo**](SpoolScaleUpCapabilityInfo.md) |  | [optional] 

## Methods

### NewDatacenter

`func NewDatacenter(name string, datacenterType string, provider string, operState string, available bool, ) *Datacenter`

NewDatacenter instantiates a new Datacenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterWithDefaults

`func NewDatacenterWithDefaults() *Datacenter`

NewDatacenterWithDefaults instantiates a new Datacenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Datacenter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Datacenter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Datacenter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Datacenter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Datacenter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Datacenter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Datacenter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Datacenter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Datacenter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Datacenter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Datacenter) SetName(v string)`

SetName sets Name field to given value.


### GetDatacenterType

`func (o *Datacenter) GetDatacenterType() string`

GetDatacenterType returns the DatacenterType field if non-nil, zero value otherwise.

### GetDatacenterTypeOk

`func (o *Datacenter) GetDatacenterTypeOk() (*string, bool)`

GetDatacenterTypeOk returns a tuple with the DatacenterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterType

`func (o *Datacenter) SetDatacenterType(v string)`

SetDatacenterType sets DatacenterType field to given value.


### GetProvider

`func (o *Datacenter) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Datacenter) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Datacenter) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOperState

`func (o *Datacenter) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *Datacenter) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *Datacenter) SetOperState(v string)`

SetOperState sets OperState field to given value.


### GetLocation

`func (o *Datacenter) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Datacenter) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Datacenter) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Datacenter) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Datacenter) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Datacenter) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Datacenter) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Datacenter) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Datacenter) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Datacenter) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Datacenter) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Datacenter) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Datacenter) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Datacenter) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Datacenter) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Datacenter) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *Datacenter) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Datacenter) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Datacenter) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *Datacenter) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetAvailable

`func (o *Datacenter) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Datacenter) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Datacenter) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetSupportedServiceClasses

`func (o *Datacenter) GetSupportedServiceClasses() []string`

GetSupportedServiceClasses returns the SupportedServiceClasses field if non-nil, zero value otherwise.

### GetSupportedServiceClassesOk

`func (o *Datacenter) GetSupportedServiceClassesOk() (*[]string, bool)`

GetSupportedServiceClassesOk returns a tuple with the SupportedServiceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedServiceClasses

`func (o *Datacenter) SetSupportedServiceClasses(v []string)`

SetSupportedServiceClasses sets SupportedServiceClasses field to given value.

### HasSupportedServiceClasses

`func (o *Datacenter) HasSupportedServiceClasses() bool`

HasSupportedServiceClasses returns a boolean if a field has been set.

### GetCloudAgentVersion

`func (o *Datacenter) GetCloudAgentVersion() string`

GetCloudAgentVersion returns the CloudAgentVersion field if non-nil, zero value otherwise.

### GetCloudAgentVersionOk

`func (o *Datacenter) GetCloudAgentVersionOk() (*string, bool)`

GetCloudAgentVersionOk returns a tuple with the CloudAgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAgentVersion

`func (o *Datacenter) SetCloudAgentVersion(v string)`

SetCloudAgentVersion sets CloudAgentVersion field to given value.

### HasCloudAgentVersion

`func (o *Datacenter) HasCloudAgentVersion() bool`

HasCloudAgentVersion returns a boolean if a field has been set.

### GetK8sServiceType

`func (o *Datacenter) GetK8sServiceType() string`

GetK8sServiceType returns the K8sServiceType field if non-nil, zero value otherwise.

### GetK8sServiceTypeOk

`func (o *Datacenter) GetK8sServiceTypeOk() (*string, bool)`

GetK8sServiceTypeOk returns a tuple with the K8sServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sServiceType

`func (o *Datacenter) SetK8sServiceType(v string)`

SetK8sServiceType sets K8sServiceType field to given value.

### HasK8sServiceType

`func (o *Datacenter) HasK8sServiceType() bool`

HasK8sServiceType returns a boolean if a field has been set.

### GetNumSupportedPrivateEndpoints

`func (o *Datacenter) GetNumSupportedPrivateEndpoints() int32`

GetNumSupportedPrivateEndpoints returns the NumSupportedPrivateEndpoints field if non-nil, zero value otherwise.

### GetNumSupportedPrivateEndpointsOk

`func (o *Datacenter) GetNumSupportedPrivateEndpointsOk() (*int32, bool)`

GetNumSupportedPrivateEndpointsOk returns a tuple with the NumSupportedPrivateEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSupportedPrivateEndpoints

`func (o *Datacenter) SetNumSupportedPrivateEndpoints(v int32)`

SetNumSupportedPrivateEndpoints sets NumSupportedPrivateEndpoints field to given value.

### HasNumSupportedPrivateEndpoints

`func (o *Datacenter) HasNumSupportedPrivateEndpoints() bool`

HasNumSupportedPrivateEndpoints returns a boolean if a field has been set.

### GetNumSupportedPublicEndpoints

`func (o *Datacenter) GetNumSupportedPublicEndpoints() int32`

GetNumSupportedPublicEndpoints returns the NumSupportedPublicEndpoints field if non-nil, zero value otherwise.

### GetNumSupportedPublicEndpointsOk

`func (o *Datacenter) GetNumSupportedPublicEndpointsOk() (*int32, bool)`

GetNumSupportedPublicEndpointsOk returns a tuple with the NumSupportedPublicEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSupportedPublicEndpoints

`func (o *Datacenter) SetNumSupportedPublicEndpoints(v int32)`

SetNumSupportedPublicEndpoints sets NumSupportedPublicEndpoints field to given value.

### HasNumSupportedPublicEndpoints

`func (o *Datacenter) HasNumSupportedPublicEndpoints() bool`

HasNumSupportedPublicEndpoints returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Datacenter) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Datacenter) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Datacenter) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Datacenter) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSpoolScaleUpCapabilityInfo

`func (o *Datacenter) GetSpoolScaleUpCapabilityInfo() SpoolScaleUpCapabilityInfo`

GetSpoolScaleUpCapabilityInfo returns the SpoolScaleUpCapabilityInfo field if non-nil, zero value otherwise.

### GetSpoolScaleUpCapabilityInfoOk

`func (o *Datacenter) GetSpoolScaleUpCapabilityInfoOk() (*SpoolScaleUpCapabilityInfo, bool)`

GetSpoolScaleUpCapabilityInfoOk returns a tuple with the SpoolScaleUpCapabilityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoolScaleUpCapabilityInfo

`func (o *Datacenter) SetSpoolScaleUpCapabilityInfo(v SpoolScaleUpCapabilityInfo)`

SetSpoolScaleUpCapabilityInfo sets SpoolScaleUpCapabilityInfo field to given value.

### HasSpoolScaleUpCapabilityInfo

`func (o *Datacenter) HasSpoolScaleUpCapabilityInfo() bool`

HasSpoolScaleUpCapabilityInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


