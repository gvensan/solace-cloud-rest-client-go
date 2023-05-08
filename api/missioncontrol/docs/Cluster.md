# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the DMR cluster. | [optional] 
**Password** | Pointer to **string** | The password for the cluster. | [optional] 
**RemoteAddress** | Pointer to **string** | The address of the remote node in the cluster. | [optional] 
**PrimaryRouterName** | Pointer to **string** | The name of the primary event broker. | [optional] 
**BackupRouterName** | Pointer to **string** | The name of the backup event broker. | [optional] 
**MonitoringRouterName** | Pointer to **string** | The name of the monitoring node. | [optional] 
**SupportedAuthenticationMode** | Pointer to **[]string** | The authentication mode between the nodes in the DMR cluster. | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *Cluster) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Cluster) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Cluster) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Cluster) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *Cluster) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *Cluster) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *Cluster) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *Cluster) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetPrimaryRouterName

`func (o *Cluster) GetPrimaryRouterName() string`

GetPrimaryRouterName returns the PrimaryRouterName field if non-nil, zero value otherwise.

### GetPrimaryRouterNameOk

`func (o *Cluster) GetPrimaryRouterNameOk() (*string, bool)`

GetPrimaryRouterNameOk returns a tuple with the PrimaryRouterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRouterName

`func (o *Cluster) SetPrimaryRouterName(v string)`

SetPrimaryRouterName sets PrimaryRouterName field to given value.

### HasPrimaryRouterName

`func (o *Cluster) HasPrimaryRouterName() bool`

HasPrimaryRouterName returns a boolean if a field has been set.

### GetBackupRouterName

`func (o *Cluster) GetBackupRouterName() string`

GetBackupRouterName returns the BackupRouterName field if non-nil, zero value otherwise.

### GetBackupRouterNameOk

`func (o *Cluster) GetBackupRouterNameOk() (*string, bool)`

GetBackupRouterNameOk returns a tuple with the BackupRouterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRouterName

`func (o *Cluster) SetBackupRouterName(v string)`

SetBackupRouterName sets BackupRouterName field to given value.

### HasBackupRouterName

`func (o *Cluster) HasBackupRouterName() bool`

HasBackupRouterName returns a boolean if a field has been set.

### GetMonitoringRouterName

`func (o *Cluster) GetMonitoringRouterName() string`

GetMonitoringRouterName returns the MonitoringRouterName field if non-nil, zero value otherwise.

### GetMonitoringRouterNameOk

`func (o *Cluster) GetMonitoringRouterNameOk() (*string, bool)`

GetMonitoringRouterNameOk returns a tuple with the MonitoringRouterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringRouterName

`func (o *Cluster) SetMonitoringRouterName(v string)`

SetMonitoringRouterName sets MonitoringRouterName field to given value.

### HasMonitoringRouterName

`func (o *Cluster) HasMonitoringRouterName() bool`

HasMonitoringRouterName returns a boolean if a field has been set.

### GetSupportedAuthenticationMode

`func (o *Cluster) GetSupportedAuthenticationMode() []string`

GetSupportedAuthenticationMode returns the SupportedAuthenticationMode field if non-nil, zero value otherwise.

### GetSupportedAuthenticationModeOk

`func (o *Cluster) GetSupportedAuthenticationModeOk() (*[]string, bool)`

GetSupportedAuthenticationModeOk returns a tuple with the SupportedAuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthenticationMode

`func (o *Cluster) SetSupportedAuthenticationMode(v []string)`

SetSupportedAuthenticationMode sets SupportedAuthenticationMode field to given value.

### HasSupportedAuthenticationMode

`func (o *Cluster) HasSupportedAuthenticationMode() bool`

HasSupportedAuthenticationMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


