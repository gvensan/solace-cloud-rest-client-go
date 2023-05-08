# Broker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version number for the event broker service. | [optional] 
**VersionFamily** | Pointer to **string** | Version Family, a truncated form of the version. | [optional] 
**ServicePackageId** | Pointer to **string** | The service package identifier of the event broker version. | [optional] 
**MaxSpoolUsage** | Pointer to **int32** | The maximum message spool usage allowed on the event broker service, in gigabytes (GB). | [optional] 
**DiskSize** | Pointer to **int32** | The disk size for the message spool, in gigabytes (GB). | [optional] 
**RedundancyGroupSslEnabled** | Pointer to **bool** | Whether redundancy group SSL is enabled. This is mate-link encryption. | [optional] 
**ConfigSyncSslEnabled** | Pointer to **bool** | Whether Config-Sync SSL is enabled. | [optional] 
**MonitoringMode** | Pointer to **string** | The monitoring mode. This can be &#39;basic&#39; or &#39;advanced&#39;. The value of basic is default monitoring and advanced means that monitoring of the event broker is enabled. | [optional] 
**ClientCertificateAuthorities** | Pointer to [**[]CertificateAuthority**](CertificateAuthority.md) | The list of client certificate authorities. | [optional] 
**DomainCertificateAuthorities** | Pointer to [**[]CertificateAuthority**](CertificateAuthority.md) | The list of domain certificate authorities. | [optional] 
**TlsStandardDomainCertificateAuthoritiesEnabled** | Pointer to **bool** | Whether TLS Standard Domain Certificate Authorities is enabled. | [optional] 
**LdapProfiles** | Pointer to [**[]LdapProfile**](LdapProfile.md) | The LDAP profiles configured for the event broker service. | [optional] 
**Cluster** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**ManagementReadOnlyLoginCredential** | Pointer to [**ManagementLoginCredential**](ManagementLoginCredential.md) |  | [optional] 
**MsgVpns** | Pointer to [**[]MsgVpn**](MsgVpn.md) | The list of Message VPNs configured on the event broker service. | [optional] 

## Methods

### NewBroker

`func NewBroker() *Broker`

NewBroker instantiates a new Broker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerWithDefaults

`func NewBrokerWithDefaults() *Broker`

NewBrokerWithDefaults instantiates a new Broker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Broker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Broker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Broker) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Broker) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionFamily

`func (o *Broker) GetVersionFamily() string`

GetVersionFamily returns the VersionFamily field if non-nil, zero value otherwise.

### GetVersionFamilyOk

`func (o *Broker) GetVersionFamilyOk() (*string, bool)`

GetVersionFamilyOk returns a tuple with the VersionFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionFamily

`func (o *Broker) SetVersionFamily(v string)`

SetVersionFamily sets VersionFamily field to given value.

### HasVersionFamily

`func (o *Broker) HasVersionFamily() bool`

HasVersionFamily returns a boolean if a field has been set.

### GetServicePackageId

`func (o *Broker) GetServicePackageId() string`

GetServicePackageId returns the ServicePackageId field if non-nil, zero value otherwise.

### GetServicePackageIdOk

`func (o *Broker) GetServicePackageIdOk() (*string, bool)`

GetServicePackageIdOk returns a tuple with the ServicePackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePackageId

`func (o *Broker) SetServicePackageId(v string)`

SetServicePackageId sets ServicePackageId field to given value.

### HasServicePackageId

`func (o *Broker) HasServicePackageId() bool`

HasServicePackageId returns a boolean if a field has been set.

### GetMaxSpoolUsage

`func (o *Broker) GetMaxSpoolUsage() int32`

GetMaxSpoolUsage returns the MaxSpoolUsage field if non-nil, zero value otherwise.

### GetMaxSpoolUsageOk

`func (o *Broker) GetMaxSpoolUsageOk() (*int32, bool)`

GetMaxSpoolUsageOk returns a tuple with the MaxSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpoolUsage

`func (o *Broker) SetMaxSpoolUsage(v int32)`

SetMaxSpoolUsage sets MaxSpoolUsage field to given value.

### HasMaxSpoolUsage

`func (o *Broker) HasMaxSpoolUsage() bool`

HasMaxSpoolUsage returns a boolean if a field has been set.

### GetDiskSize

`func (o *Broker) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *Broker) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *Broker) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *Broker) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetRedundancyGroupSslEnabled

`func (o *Broker) GetRedundancyGroupSslEnabled() bool`

GetRedundancyGroupSslEnabled returns the RedundancyGroupSslEnabled field if non-nil, zero value otherwise.

### GetRedundancyGroupSslEnabledOk

`func (o *Broker) GetRedundancyGroupSslEnabledOk() (*bool, bool)`

GetRedundancyGroupSslEnabledOk returns a tuple with the RedundancyGroupSslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyGroupSslEnabled

`func (o *Broker) SetRedundancyGroupSslEnabled(v bool)`

SetRedundancyGroupSslEnabled sets RedundancyGroupSslEnabled field to given value.

### HasRedundancyGroupSslEnabled

`func (o *Broker) HasRedundancyGroupSslEnabled() bool`

HasRedundancyGroupSslEnabled returns a boolean if a field has been set.

### GetConfigSyncSslEnabled

`func (o *Broker) GetConfigSyncSslEnabled() bool`

GetConfigSyncSslEnabled returns the ConfigSyncSslEnabled field if non-nil, zero value otherwise.

### GetConfigSyncSslEnabledOk

`func (o *Broker) GetConfigSyncSslEnabledOk() (*bool, bool)`

GetConfigSyncSslEnabledOk returns a tuple with the ConfigSyncSslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncSslEnabled

`func (o *Broker) SetConfigSyncSslEnabled(v bool)`

SetConfigSyncSslEnabled sets ConfigSyncSslEnabled field to given value.

### HasConfigSyncSslEnabled

`func (o *Broker) HasConfigSyncSslEnabled() bool`

HasConfigSyncSslEnabled returns a boolean if a field has been set.

### GetMonitoringMode

`func (o *Broker) GetMonitoringMode() string`

GetMonitoringMode returns the MonitoringMode field if non-nil, zero value otherwise.

### GetMonitoringModeOk

`func (o *Broker) GetMonitoringModeOk() (*string, bool)`

GetMonitoringModeOk returns a tuple with the MonitoringMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringMode

`func (o *Broker) SetMonitoringMode(v string)`

SetMonitoringMode sets MonitoringMode field to given value.

### HasMonitoringMode

`func (o *Broker) HasMonitoringMode() bool`

HasMonitoringMode returns a boolean if a field has been set.

### GetClientCertificateAuthorities

`func (o *Broker) GetClientCertificateAuthorities() []CertificateAuthority`

GetClientCertificateAuthorities returns the ClientCertificateAuthorities field if non-nil, zero value otherwise.

### GetClientCertificateAuthoritiesOk

`func (o *Broker) GetClientCertificateAuthoritiesOk() (*[]CertificateAuthority, bool)`

GetClientCertificateAuthoritiesOk returns a tuple with the ClientCertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateAuthorities

`func (o *Broker) SetClientCertificateAuthorities(v []CertificateAuthority)`

SetClientCertificateAuthorities sets ClientCertificateAuthorities field to given value.

### HasClientCertificateAuthorities

`func (o *Broker) HasClientCertificateAuthorities() bool`

HasClientCertificateAuthorities returns a boolean if a field has been set.

### GetDomainCertificateAuthorities

`func (o *Broker) GetDomainCertificateAuthorities() []CertificateAuthority`

GetDomainCertificateAuthorities returns the DomainCertificateAuthorities field if non-nil, zero value otherwise.

### GetDomainCertificateAuthoritiesOk

`func (o *Broker) GetDomainCertificateAuthoritiesOk() (*[]CertificateAuthority, bool)`

GetDomainCertificateAuthoritiesOk returns a tuple with the DomainCertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCertificateAuthorities

`func (o *Broker) SetDomainCertificateAuthorities(v []CertificateAuthority)`

SetDomainCertificateAuthorities sets DomainCertificateAuthorities field to given value.

### HasDomainCertificateAuthorities

`func (o *Broker) HasDomainCertificateAuthorities() bool`

HasDomainCertificateAuthorities returns a boolean if a field has been set.

### GetTlsStandardDomainCertificateAuthoritiesEnabled

`func (o *Broker) GetTlsStandardDomainCertificateAuthoritiesEnabled() bool`

GetTlsStandardDomainCertificateAuthoritiesEnabled returns the TlsStandardDomainCertificateAuthoritiesEnabled field if non-nil, zero value otherwise.

### GetTlsStandardDomainCertificateAuthoritiesEnabledOk

`func (o *Broker) GetTlsStandardDomainCertificateAuthoritiesEnabledOk() (*bool, bool)`

GetTlsStandardDomainCertificateAuthoritiesEnabledOk returns a tuple with the TlsStandardDomainCertificateAuthoritiesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsStandardDomainCertificateAuthoritiesEnabled

`func (o *Broker) SetTlsStandardDomainCertificateAuthoritiesEnabled(v bool)`

SetTlsStandardDomainCertificateAuthoritiesEnabled sets TlsStandardDomainCertificateAuthoritiesEnabled field to given value.

### HasTlsStandardDomainCertificateAuthoritiesEnabled

`func (o *Broker) HasTlsStandardDomainCertificateAuthoritiesEnabled() bool`

HasTlsStandardDomainCertificateAuthoritiesEnabled returns a boolean if a field has been set.

### GetLdapProfiles

`func (o *Broker) GetLdapProfiles() []LdapProfile`

GetLdapProfiles returns the LdapProfiles field if non-nil, zero value otherwise.

### GetLdapProfilesOk

`func (o *Broker) GetLdapProfilesOk() (*[]LdapProfile, bool)`

GetLdapProfilesOk returns a tuple with the LdapProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProfiles

`func (o *Broker) SetLdapProfiles(v []LdapProfile)`

SetLdapProfiles sets LdapProfiles field to given value.

### HasLdapProfiles

`func (o *Broker) HasLdapProfiles() bool`

HasLdapProfiles returns a boolean if a field has been set.

### GetCluster

`func (o *Broker) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Broker) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Broker) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Broker) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetManagementReadOnlyLoginCredential

`func (o *Broker) GetManagementReadOnlyLoginCredential() ManagementLoginCredential`

GetManagementReadOnlyLoginCredential returns the ManagementReadOnlyLoginCredential field if non-nil, zero value otherwise.

### GetManagementReadOnlyLoginCredentialOk

`func (o *Broker) GetManagementReadOnlyLoginCredentialOk() (*ManagementLoginCredential, bool)`

GetManagementReadOnlyLoginCredentialOk returns a tuple with the ManagementReadOnlyLoginCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementReadOnlyLoginCredential

`func (o *Broker) SetManagementReadOnlyLoginCredential(v ManagementLoginCredential)`

SetManagementReadOnlyLoginCredential sets ManagementReadOnlyLoginCredential field to given value.

### HasManagementReadOnlyLoginCredential

`func (o *Broker) HasManagementReadOnlyLoginCredential() bool`

HasManagementReadOnlyLoginCredential returns a boolean if a field has been set.

### GetMsgVpns

`func (o *Broker) GetMsgVpns() []MsgVpn`

GetMsgVpns returns the MsgVpns field if non-nil, zero value otherwise.

### GetMsgVpnsOk

`func (o *Broker) GetMsgVpnsOk() (*[]MsgVpn, bool)`

GetMsgVpnsOk returns a tuple with the MsgVpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpns

`func (o *Broker) SetMsgVpns(v []MsgVpn)`

SetMsgVpns sets MsgVpns field to given value.

### HasMsgVpns

`func (o *Broker) HasMsgVpns() bool`

HasMsgVpns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


