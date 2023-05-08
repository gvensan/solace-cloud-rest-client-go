# MsgVpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**AuthenticationBasicEnabled** | Pointer to **bool** | Whether basic authentication is enabled. | [optional] 
**AuthenticationBasicType** | Pointer to **string** | The authentication type. This can be &#39;none&#39;, &#39;internal&#39;, or &#39;ldap&#39;. | [optional] 
**AuthenticationClientCertEnabled** | Pointer to **bool** | Whether client certificate authentication is enabled. | [optional] 
**AuthenticationClientCertValidateDateEnabled** | Pointer to **bool** | Whether the validation of the &#39;Not Before&#39; and &#39;Not After&#39; dates in a client certificate is enabled. | [optional] 
**AuthenticationOauthEnabled** | Pointer to **bool** | Whether OAuth authentication is enabled for the Message VPN. | [optional] 
**ClientProfiles** | Pointer to [**[]ClientProfileBase**](ClientProfileBase.md) | The client profiles configured on the Message VPN. | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**EventLargeMsgThreshold** | Pointer to **int32** | The large message threshold generates events when the size of a message in a Message VPN exceeds a specified size. | [optional] 
**ManagementAdminLoginCredential** | Pointer to [**ManagementLoginCredential**](ManagementLoginCredential.md) |  | [optional] 
**ServiceLoginCredential** | Pointer to [**LoginCredential**](LoginCredential.md) |  | [optional] 
**MaxConnectionCount** | Pointer to **int32** | The maximum number of clients that are permitted to simultaneously connect to the Message VPN.  | [optional] 
**MaxEgressFlowCount** | Pointer to **int32** | The maximum number of egress flows (that is, Guaranteed message client receive flows or consumer flows) that can be created by a single client. | [optional] 
**MaxEndpointCount** | Pointer to **int32** | The maximum number of flows that can bind to a non-exclusive durable topic endpoint. | [optional] 
**MaxIngressFlowCount** | Pointer to **int32** | The total permitted number of ingress flows (that is, Guaranteed Message client publish flows) for a Message VPN. | [optional] 
**MaxMsgSpoolUsage** | Pointer to **int32** | The maximum message spool usage. | [optional] 
**MaxSubscriptionCount** | Pointer to **int32** | The maximum number of unique subscriptions. | [optional] 
**MaxTransactedSessionCount** | Pointer to **int32** | The maximum number of simultaneous transacted sessions and/or XA Sessions allowed for the given Message VPN. | [optional] 
**MaxTransactionCount** | Pointer to **int32** | The total number of simultaneous transactions (both local transactions and transactions within distributed/XA transaction branches) in a Message VPN. | [optional] 
**SempOverMessageBus** | Pointer to [**SEMPOverMsgBus**](SEMPOverMsgBus.md) |  | [optional] 
**SubDomainName** | Pointer to **string** | The generated hostname assigned for the Message VPN.  For example, &#39;mr54hcalmefac.messaging.solace.cloud&#39;. | [optional] 
**TruststoreUri** | Pointer to **string** | The URI for the SSL trust store. | [optional] 

## Methods

### NewMsgVpn

`func NewMsgVpn() *MsgVpn`

NewMsgVpn instantiates a new MsgVpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnWithDefaults

`func NewMsgVpnWithDefaults() *MsgVpn`

NewMsgVpnWithDefaults instantiates a new MsgVpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgVpnName

`func (o *MsgVpn) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpn) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpn) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpn) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetAuthenticationBasicEnabled

`func (o *MsgVpn) GetAuthenticationBasicEnabled() bool`

GetAuthenticationBasicEnabled returns the AuthenticationBasicEnabled field if non-nil, zero value otherwise.

### GetAuthenticationBasicEnabledOk

`func (o *MsgVpn) GetAuthenticationBasicEnabledOk() (*bool, bool)`

GetAuthenticationBasicEnabledOk returns a tuple with the AuthenticationBasicEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicEnabled

`func (o *MsgVpn) SetAuthenticationBasicEnabled(v bool)`

SetAuthenticationBasicEnabled sets AuthenticationBasicEnabled field to given value.

### HasAuthenticationBasicEnabled

`func (o *MsgVpn) HasAuthenticationBasicEnabled() bool`

HasAuthenticationBasicEnabled returns a boolean if a field has been set.

### GetAuthenticationBasicType

`func (o *MsgVpn) GetAuthenticationBasicType() string`

GetAuthenticationBasicType returns the AuthenticationBasicType field if non-nil, zero value otherwise.

### GetAuthenticationBasicTypeOk

`func (o *MsgVpn) GetAuthenticationBasicTypeOk() (*string, bool)`

GetAuthenticationBasicTypeOk returns a tuple with the AuthenticationBasicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicType

`func (o *MsgVpn) SetAuthenticationBasicType(v string)`

SetAuthenticationBasicType sets AuthenticationBasicType field to given value.

### HasAuthenticationBasicType

`func (o *MsgVpn) HasAuthenticationBasicType() bool`

HasAuthenticationBasicType returns a boolean if a field has been set.

### GetAuthenticationClientCertEnabled

`func (o *MsgVpn) GetAuthenticationClientCertEnabled() bool`

GetAuthenticationClientCertEnabled returns the AuthenticationClientCertEnabled field if non-nil, zero value otherwise.

### GetAuthenticationClientCertEnabledOk

`func (o *MsgVpn) GetAuthenticationClientCertEnabledOk() (*bool, bool)`

GetAuthenticationClientCertEnabledOk returns a tuple with the AuthenticationClientCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertEnabled

`func (o *MsgVpn) SetAuthenticationClientCertEnabled(v bool)`

SetAuthenticationClientCertEnabled sets AuthenticationClientCertEnabled field to given value.

### HasAuthenticationClientCertEnabled

`func (o *MsgVpn) HasAuthenticationClientCertEnabled() bool`

HasAuthenticationClientCertEnabled returns a boolean if a field has been set.

### GetAuthenticationClientCertValidateDateEnabled

`func (o *MsgVpn) GetAuthenticationClientCertValidateDateEnabled() bool`

GetAuthenticationClientCertValidateDateEnabled returns the AuthenticationClientCertValidateDateEnabled field if non-nil, zero value otherwise.

### GetAuthenticationClientCertValidateDateEnabledOk

`func (o *MsgVpn) GetAuthenticationClientCertValidateDateEnabledOk() (*bool, bool)`

GetAuthenticationClientCertValidateDateEnabledOk returns a tuple with the AuthenticationClientCertValidateDateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertValidateDateEnabled

`func (o *MsgVpn) SetAuthenticationClientCertValidateDateEnabled(v bool)`

SetAuthenticationClientCertValidateDateEnabled sets AuthenticationClientCertValidateDateEnabled field to given value.

### HasAuthenticationClientCertValidateDateEnabled

`func (o *MsgVpn) HasAuthenticationClientCertValidateDateEnabled() bool`

HasAuthenticationClientCertValidateDateEnabled returns a boolean if a field has been set.

### GetAuthenticationOauthEnabled

`func (o *MsgVpn) GetAuthenticationOauthEnabled() bool`

GetAuthenticationOauthEnabled returns the AuthenticationOauthEnabled field if non-nil, zero value otherwise.

### GetAuthenticationOauthEnabledOk

`func (o *MsgVpn) GetAuthenticationOauthEnabledOk() (*bool, bool)`

GetAuthenticationOauthEnabledOk returns a tuple with the AuthenticationOauthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthEnabled

`func (o *MsgVpn) SetAuthenticationOauthEnabled(v bool)`

SetAuthenticationOauthEnabled sets AuthenticationOauthEnabled field to given value.

### HasAuthenticationOauthEnabled

`func (o *MsgVpn) HasAuthenticationOauthEnabled() bool`

HasAuthenticationOauthEnabled returns a boolean if a field has been set.

### GetClientProfiles

`func (o *MsgVpn) GetClientProfiles() []ClientProfileBase`

GetClientProfiles returns the ClientProfiles field if non-nil, zero value otherwise.

### GetClientProfilesOk

`func (o *MsgVpn) GetClientProfilesOk() (*[]ClientProfileBase, bool)`

GetClientProfilesOk returns a tuple with the ClientProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfiles

`func (o *MsgVpn) SetClientProfiles(v []ClientProfileBase)`

SetClientProfiles sets ClientProfiles field to given value.

### HasClientProfiles

`func (o *MsgVpn) HasClientProfiles() bool`

HasClientProfiles returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpn) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpn) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpn) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpn) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventLargeMsgThreshold

`func (o *MsgVpn) GetEventLargeMsgThreshold() int32`

GetEventLargeMsgThreshold returns the EventLargeMsgThreshold field if non-nil, zero value otherwise.

### GetEventLargeMsgThresholdOk

`func (o *MsgVpn) GetEventLargeMsgThresholdOk() (*int32, bool)`

GetEventLargeMsgThresholdOk returns a tuple with the EventLargeMsgThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLargeMsgThreshold

`func (o *MsgVpn) SetEventLargeMsgThreshold(v int32)`

SetEventLargeMsgThreshold sets EventLargeMsgThreshold field to given value.

### HasEventLargeMsgThreshold

`func (o *MsgVpn) HasEventLargeMsgThreshold() bool`

HasEventLargeMsgThreshold returns a boolean if a field has been set.

### GetManagementAdminLoginCredential

`func (o *MsgVpn) GetManagementAdminLoginCredential() ManagementLoginCredential`

GetManagementAdminLoginCredential returns the ManagementAdminLoginCredential field if non-nil, zero value otherwise.

### GetManagementAdminLoginCredentialOk

`func (o *MsgVpn) GetManagementAdminLoginCredentialOk() (*ManagementLoginCredential, bool)`

GetManagementAdminLoginCredentialOk returns a tuple with the ManagementAdminLoginCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAdminLoginCredential

`func (o *MsgVpn) SetManagementAdminLoginCredential(v ManagementLoginCredential)`

SetManagementAdminLoginCredential sets ManagementAdminLoginCredential field to given value.

### HasManagementAdminLoginCredential

`func (o *MsgVpn) HasManagementAdminLoginCredential() bool`

HasManagementAdminLoginCredential returns a boolean if a field has been set.

### GetServiceLoginCredential

`func (o *MsgVpn) GetServiceLoginCredential() LoginCredential`

GetServiceLoginCredential returns the ServiceLoginCredential field if non-nil, zero value otherwise.

### GetServiceLoginCredentialOk

`func (o *MsgVpn) GetServiceLoginCredentialOk() (*LoginCredential, bool)`

GetServiceLoginCredentialOk returns a tuple with the ServiceLoginCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLoginCredential

`func (o *MsgVpn) SetServiceLoginCredential(v LoginCredential)`

SetServiceLoginCredential sets ServiceLoginCredential field to given value.

### HasServiceLoginCredential

`func (o *MsgVpn) HasServiceLoginCredential() bool`

HasServiceLoginCredential returns a boolean if a field has been set.

### GetMaxConnectionCount

`func (o *MsgVpn) GetMaxConnectionCount() int32`

GetMaxConnectionCount returns the MaxConnectionCount field if non-nil, zero value otherwise.

### GetMaxConnectionCountOk

`func (o *MsgVpn) GetMaxConnectionCountOk() (*int32, bool)`

GetMaxConnectionCountOk returns a tuple with the MaxConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionCount

`func (o *MsgVpn) SetMaxConnectionCount(v int32)`

SetMaxConnectionCount sets MaxConnectionCount field to given value.

### HasMaxConnectionCount

`func (o *MsgVpn) HasMaxConnectionCount() bool`

HasMaxConnectionCount returns a boolean if a field has been set.

### GetMaxEgressFlowCount

`func (o *MsgVpn) GetMaxEgressFlowCount() int32`

GetMaxEgressFlowCount returns the MaxEgressFlowCount field if non-nil, zero value otherwise.

### GetMaxEgressFlowCountOk

`func (o *MsgVpn) GetMaxEgressFlowCountOk() (*int32, bool)`

GetMaxEgressFlowCountOk returns a tuple with the MaxEgressFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEgressFlowCount

`func (o *MsgVpn) SetMaxEgressFlowCount(v int32)`

SetMaxEgressFlowCount sets MaxEgressFlowCount field to given value.

### HasMaxEgressFlowCount

`func (o *MsgVpn) HasMaxEgressFlowCount() bool`

HasMaxEgressFlowCount returns a boolean if a field has been set.

### GetMaxEndpointCount

`func (o *MsgVpn) GetMaxEndpointCount() int32`

GetMaxEndpointCount returns the MaxEndpointCount field if non-nil, zero value otherwise.

### GetMaxEndpointCountOk

`func (o *MsgVpn) GetMaxEndpointCountOk() (*int32, bool)`

GetMaxEndpointCountOk returns a tuple with the MaxEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndpointCount

`func (o *MsgVpn) SetMaxEndpointCount(v int32)`

SetMaxEndpointCount sets MaxEndpointCount field to given value.

### HasMaxEndpointCount

`func (o *MsgVpn) HasMaxEndpointCount() bool`

HasMaxEndpointCount returns a boolean if a field has been set.

### GetMaxIngressFlowCount

`func (o *MsgVpn) GetMaxIngressFlowCount() int32`

GetMaxIngressFlowCount returns the MaxIngressFlowCount field if non-nil, zero value otherwise.

### GetMaxIngressFlowCountOk

`func (o *MsgVpn) GetMaxIngressFlowCountOk() (*int32, bool)`

GetMaxIngressFlowCountOk returns a tuple with the MaxIngressFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIngressFlowCount

`func (o *MsgVpn) SetMaxIngressFlowCount(v int32)`

SetMaxIngressFlowCount sets MaxIngressFlowCount field to given value.

### HasMaxIngressFlowCount

`func (o *MsgVpn) HasMaxIngressFlowCount() bool`

HasMaxIngressFlowCount returns a boolean if a field has been set.

### GetMaxMsgSpoolUsage

`func (o *MsgVpn) GetMaxMsgSpoolUsage() int32`

GetMaxMsgSpoolUsage returns the MaxMsgSpoolUsage field if non-nil, zero value otherwise.

### GetMaxMsgSpoolUsageOk

`func (o *MsgVpn) GetMaxMsgSpoolUsageOk() (*int32, bool)`

GetMaxMsgSpoolUsageOk returns a tuple with the MaxMsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSpoolUsage

`func (o *MsgVpn) SetMaxMsgSpoolUsage(v int32)`

SetMaxMsgSpoolUsage sets MaxMsgSpoolUsage field to given value.

### HasMaxMsgSpoolUsage

`func (o *MsgVpn) HasMaxMsgSpoolUsage() bool`

HasMaxMsgSpoolUsage returns a boolean if a field has been set.

### GetMaxSubscriptionCount

`func (o *MsgVpn) GetMaxSubscriptionCount() int32`

GetMaxSubscriptionCount returns the MaxSubscriptionCount field if non-nil, zero value otherwise.

### GetMaxSubscriptionCountOk

`func (o *MsgVpn) GetMaxSubscriptionCountOk() (*int32, bool)`

GetMaxSubscriptionCountOk returns a tuple with the MaxSubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSubscriptionCount

`func (o *MsgVpn) SetMaxSubscriptionCount(v int32)`

SetMaxSubscriptionCount sets MaxSubscriptionCount field to given value.

### HasMaxSubscriptionCount

`func (o *MsgVpn) HasMaxSubscriptionCount() bool`

HasMaxSubscriptionCount returns a boolean if a field has been set.

### GetMaxTransactedSessionCount

`func (o *MsgVpn) GetMaxTransactedSessionCount() int32`

GetMaxTransactedSessionCount returns the MaxTransactedSessionCount field if non-nil, zero value otherwise.

### GetMaxTransactedSessionCountOk

`func (o *MsgVpn) GetMaxTransactedSessionCountOk() (*int32, bool)`

GetMaxTransactedSessionCountOk returns a tuple with the MaxTransactedSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactedSessionCount

`func (o *MsgVpn) SetMaxTransactedSessionCount(v int32)`

SetMaxTransactedSessionCount sets MaxTransactedSessionCount field to given value.

### HasMaxTransactedSessionCount

`func (o *MsgVpn) HasMaxTransactedSessionCount() bool`

HasMaxTransactedSessionCount returns a boolean if a field has been set.

### GetMaxTransactionCount

`func (o *MsgVpn) GetMaxTransactionCount() int32`

GetMaxTransactionCount returns the MaxTransactionCount field if non-nil, zero value otherwise.

### GetMaxTransactionCountOk

`func (o *MsgVpn) GetMaxTransactionCountOk() (*int32, bool)`

GetMaxTransactionCountOk returns a tuple with the MaxTransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactionCount

`func (o *MsgVpn) SetMaxTransactionCount(v int32)`

SetMaxTransactionCount sets MaxTransactionCount field to given value.

### HasMaxTransactionCount

`func (o *MsgVpn) HasMaxTransactionCount() bool`

HasMaxTransactionCount returns a boolean if a field has been set.

### GetSempOverMessageBus

`func (o *MsgVpn) GetSempOverMessageBus() SEMPOverMsgBus`

GetSempOverMessageBus returns the SempOverMessageBus field if non-nil, zero value otherwise.

### GetSempOverMessageBusOk

`func (o *MsgVpn) GetSempOverMessageBusOk() (*SEMPOverMsgBus, bool)`

GetSempOverMessageBusOk returns a tuple with the SempOverMessageBus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempOverMessageBus

`func (o *MsgVpn) SetSempOverMessageBus(v SEMPOverMsgBus)`

SetSempOverMessageBus sets SempOverMessageBus field to given value.

### HasSempOverMessageBus

`func (o *MsgVpn) HasSempOverMessageBus() bool`

HasSempOverMessageBus returns a boolean if a field has been set.

### GetSubDomainName

`func (o *MsgVpn) GetSubDomainName() string`

GetSubDomainName returns the SubDomainName field if non-nil, zero value otherwise.

### GetSubDomainNameOk

`func (o *MsgVpn) GetSubDomainNameOk() (*string, bool)`

GetSubDomainNameOk returns a tuple with the SubDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDomainName

`func (o *MsgVpn) SetSubDomainName(v string)`

SetSubDomainName sets SubDomainName field to given value.

### HasSubDomainName

`func (o *MsgVpn) HasSubDomainName() bool`

HasSubDomainName returns a boolean if a field has been set.

### GetTruststoreUri

`func (o *MsgVpn) GetTruststoreUri() string`

GetTruststoreUri returns the TruststoreUri field if non-nil, zero value otherwise.

### GetTruststoreUriOk

`func (o *MsgVpn) GetTruststoreUriOk() (*string, bool)`

GetTruststoreUriOk returns a tuple with the TruststoreUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststoreUri

`func (o *MsgVpn) SetTruststoreUri(v string)`

SetTruststoreUri sets TruststoreUri field to given value.

### HasTruststoreUri

`func (o *MsgVpn) HasTruststoreUri() bool`

HasTruststoreUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


