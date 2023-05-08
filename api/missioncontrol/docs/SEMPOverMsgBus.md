# SEMPOverMsgBus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SempOverMsgBusEnabled** | Pointer to **bool** | Whether SEMP over Message Bus is enabled.  When enabled for a Message VPN, clients have access to a limited subset of the event broker management commands. | [optional] 
**SempAccessToShowCmdsEnabled** | Pointer to **bool** | Whether access to show SEMP commands is enabled. | [optional] 
**SempAccessToAdminCmdsEnabled** | Pointer to **bool** | Whether access to SEMP commands with the admin user access-level is enabled. | [optional] 
**SempAccessToClientAdminCmdsEnabled** | Pointer to **bool** | Whether access to SEMP Client-Admin commands is enabled. | [optional] 
**SempAccessToCacheCmdsEnabled** | Pointer to **bool** | SEMP Access to Cache Commands Enabled | [optional] 

## Methods

### NewSEMPOverMsgBus

`func NewSEMPOverMsgBus() *SEMPOverMsgBus`

NewSEMPOverMsgBus instantiates a new SEMPOverMsgBus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSEMPOverMsgBusWithDefaults

`func NewSEMPOverMsgBusWithDefaults() *SEMPOverMsgBus`

NewSEMPOverMsgBusWithDefaults instantiates a new SEMPOverMsgBus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSempOverMsgBusEnabled

`func (o *SEMPOverMsgBus) GetSempOverMsgBusEnabled() bool`

GetSempOverMsgBusEnabled returns the SempOverMsgBusEnabled field if non-nil, zero value otherwise.

### GetSempOverMsgBusEnabledOk

`func (o *SEMPOverMsgBus) GetSempOverMsgBusEnabledOk() (*bool, bool)`

GetSempOverMsgBusEnabledOk returns a tuple with the SempOverMsgBusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempOverMsgBusEnabled

`func (o *SEMPOverMsgBus) SetSempOverMsgBusEnabled(v bool)`

SetSempOverMsgBusEnabled sets SempOverMsgBusEnabled field to given value.

### HasSempOverMsgBusEnabled

`func (o *SEMPOverMsgBus) HasSempOverMsgBusEnabled() bool`

HasSempOverMsgBusEnabled returns a boolean if a field has been set.

### GetSempAccessToShowCmdsEnabled

`func (o *SEMPOverMsgBus) GetSempAccessToShowCmdsEnabled() bool`

GetSempAccessToShowCmdsEnabled returns the SempAccessToShowCmdsEnabled field if non-nil, zero value otherwise.

### GetSempAccessToShowCmdsEnabledOk

`func (o *SEMPOverMsgBus) GetSempAccessToShowCmdsEnabledOk() (*bool, bool)`

GetSempAccessToShowCmdsEnabledOk returns a tuple with the SempAccessToShowCmdsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempAccessToShowCmdsEnabled

`func (o *SEMPOverMsgBus) SetSempAccessToShowCmdsEnabled(v bool)`

SetSempAccessToShowCmdsEnabled sets SempAccessToShowCmdsEnabled field to given value.

### HasSempAccessToShowCmdsEnabled

`func (o *SEMPOverMsgBus) HasSempAccessToShowCmdsEnabled() bool`

HasSempAccessToShowCmdsEnabled returns a boolean if a field has been set.

### GetSempAccessToAdminCmdsEnabled

`func (o *SEMPOverMsgBus) GetSempAccessToAdminCmdsEnabled() bool`

GetSempAccessToAdminCmdsEnabled returns the SempAccessToAdminCmdsEnabled field if non-nil, zero value otherwise.

### GetSempAccessToAdminCmdsEnabledOk

`func (o *SEMPOverMsgBus) GetSempAccessToAdminCmdsEnabledOk() (*bool, bool)`

GetSempAccessToAdminCmdsEnabledOk returns a tuple with the SempAccessToAdminCmdsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempAccessToAdminCmdsEnabled

`func (o *SEMPOverMsgBus) SetSempAccessToAdminCmdsEnabled(v bool)`

SetSempAccessToAdminCmdsEnabled sets SempAccessToAdminCmdsEnabled field to given value.

### HasSempAccessToAdminCmdsEnabled

`func (o *SEMPOverMsgBus) HasSempAccessToAdminCmdsEnabled() bool`

HasSempAccessToAdminCmdsEnabled returns a boolean if a field has been set.

### GetSempAccessToClientAdminCmdsEnabled

`func (o *SEMPOverMsgBus) GetSempAccessToClientAdminCmdsEnabled() bool`

GetSempAccessToClientAdminCmdsEnabled returns the SempAccessToClientAdminCmdsEnabled field if non-nil, zero value otherwise.

### GetSempAccessToClientAdminCmdsEnabledOk

`func (o *SEMPOverMsgBus) GetSempAccessToClientAdminCmdsEnabledOk() (*bool, bool)`

GetSempAccessToClientAdminCmdsEnabledOk returns a tuple with the SempAccessToClientAdminCmdsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempAccessToClientAdminCmdsEnabled

`func (o *SEMPOverMsgBus) SetSempAccessToClientAdminCmdsEnabled(v bool)`

SetSempAccessToClientAdminCmdsEnabled sets SempAccessToClientAdminCmdsEnabled field to given value.

### HasSempAccessToClientAdminCmdsEnabled

`func (o *SEMPOverMsgBus) HasSempAccessToClientAdminCmdsEnabled() bool`

HasSempAccessToClientAdminCmdsEnabled returns a boolean if a field has been set.

### GetSempAccessToCacheCmdsEnabled

`func (o *SEMPOverMsgBus) GetSempAccessToCacheCmdsEnabled() bool`

GetSempAccessToCacheCmdsEnabled returns the SempAccessToCacheCmdsEnabled field if non-nil, zero value otherwise.

### GetSempAccessToCacheCmdsEnabledOk

`func (o *SEMPOverMsgBus) GetSempAccessToCacheCmdsEnabledOk() (*bool, bool)`

GetSempAccessToCacheCmdsEnabledOk returns a tuple with the SempAccessToCacheCmdsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempAccessToCacheCmdsEnabled

`func (o *SEMPOverMsgBus) SetSempAccessToCacheCmdsEnabled(v bool)`

SetSempAccessToCacheCmdsEnabled sets SempAccessToCacheCmdsEnabled field to given value.

### HasSempAccessToCacheCmdsEnabled

`func (o *SEMPOverMsgBus) HasSempAccessToCacheCmdsEnabled() bool`

HasSempAccessToCacheCmdsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


