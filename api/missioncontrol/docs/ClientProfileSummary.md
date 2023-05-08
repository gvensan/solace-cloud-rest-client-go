# ClientProfileSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the client profile. | 
**Id** | Pointer to **string** | The identifier of the client profile. | [optional] 
**Type** | Pointer to **string** | The type of object for informational purposes. | [optional] 
**AllowGuaranteedMsgSendEnabled** | Pointer to **bool** | Indicates whether clients assigned the client profile are allowed to publish guaranteed messages. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**AllowGuaranteedMsgReceiveEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to bind to topic endpoints or queues to receive guaranteed messages. The valid values are &#39;true&#39; (allowed) or &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**AllowTransactedSessionsEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to establish transacted sessions or XA sessions. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;true&#39;. | [optional] 
**AllowBridgeConnectionsEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to establish Dynamic Messaging Routing (DMR) links (or bridge links) from the current Message VPN to another Message VPN in a separate event broker service. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**AllowGuaranteedEndpointCreateEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to create queues or topic endpoints. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 

## Methods

### NewClientProfileSummary

`func NewClientProfileSummary(name string, ) *ClientProfileSummary`

NewClientProfileSummary instantiates a new ClientProfileSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientProfileSummaryWithDefaults

`func NewClientProfileSummaryWithDefaults() *ClientProfileSummary`

NewClientProfileSummaryWithDefaults instantiates a new ClientProfileSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientProfileSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientProfileSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientProfileSummary) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *ClientProfileSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientProfileSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientProfileSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientProfileSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ClientProfileSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientProfileSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientProfileSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientProfileSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAllowGuaranteedMsgSendEnabled

`func (o *ClientProfileSummary) GetAllowGuaranteedMsgSendEnabled() bool`

GetAllowGuaranteedMsgSendEnabled returns the AllowGuaranteedMsgSendEnabled field if non-nil, zero value otherwise.

### GetAllowGuaranteedMsgSendEnabledOk

`func (o *ClientProfileSummary) GetAllowGuaranteedMsgSendEnabledOk() (*bool, bool)`

GetAllowGuaranteedMsgSendEnabledOk returns a tuple with the AllowGuaranteedMsgSendEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedMsgSendEnabled

`func (o *ClientProfileSummary) SetAllowGuaranteedMsgSendEnabled(v bool)`

SetAllowGuaranteedMsgSendEnabled sets AllowGuaranteedMsgSendEnabled field to given value.

### HasAllowGuaranteedMsgSendEnabled

`func (o *ClientProfileSummary) HasAllowGuaranteedMsgSendEnabled() bool`

HasAllowGuaranteedMsgSendEnabled returns a boolean if a field has been set.

### GetAllowGuaranteedMsgReceiveEnabled

`func (o *ClientProfileSummary) GetAllowGuaranteedMsgReceiveEnabled() bool`

GetAllowGuaranteedMsgReceiveEnabled returns the AllowGuaranteedMsgReceiveEnabled field if non-nil, zero value otherwise.

### GetAllowGuaranteedMsgReceiveEnabledOk

`func (o *ClientProfileSummary) GetAllowGuaranteedMsgReceiveEnabledOk() (*bool, bool)`

GetAllowGuaranteedMsgReceiveEnabledOk returns a tuple with the AllowGuaranteedMsgReceiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedMsgReceiveEnabled

`func (o *ClientProfileSummary) SetAllowGuaranteedMsgReceiveEnabled(v bool)`

SetAllowGuaranteedMsgReceiveEnabled sets AllowGuaranteedMsgReceiveEnabled field to given value.

### HasAllowGuaranteedMsgReceiveEnabled

`func (o *ClientProfileSummary) HasAllowGuaranteedMsgReceiveEnabled() bool`

HasAllowGuaranteedMsgReceiveEnabled returns a boolean if a field has been set.

### GetAllowTransactedSessionsEnabled

`func (o *ClientProfileSummary) GetAllowTransactedSessionsEnabled() bool`

GetAllowTransactedSessionsEnabled returns the AllowTransactedSessionsEnabled field if non-nil, zero value otherwise.

### GetAllowTransactedSessionsEnabledOk

`func (o *ClientProfileSummary) GetAllowTransactedSessionsEnabledOk() (*bool, bool)`

GetAllowTransactedSessionsEnabledOk returns a tuple with the AllowTransactedSessionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransactedSessionsEnabled

`func (o *ClientProfileSummary) SetAllowTransactedSessionsEnabled(v bool)`

SetAllowTransactedSessionsEnabled sets AllowTransactedSessionsEnabled field to given value.

### HasAllowTransactedSessionsEnabled

`func (o *ClientProfileSummary) HasAllowTransactedSessionsEnabled() bool`

HasAllowTransactedSessionsEnabled returns a boolean if a field has been set.

### GetAllowBridgeConnectionsEnabled

`func (o *ClientProfileSummary) GetAllowBridgeConnectionsEnabled() bool`

GetAllowBridgeConnectionsEnabled returns the AllowBridgeConnectionsEnabled field if non-nil, zero value otherwise.

### GetAllowBridgeConnectionsEnabledOk

`func (o *ClientProfileSummary) GetAllowBridgeConnectionsEnabledOk() (*bool, bool)`

GetAllowBridgeConnectionsEnabledOk returns a tuple with the AllowBridgeConnectionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBridgeConnectionsEnabled

`func (o *ClientProfileSummary) SetAllowBridgeConnectionsEnabled(v bool)`

SetAllowBridgeConnectionsEnabled sets AllowBridgeConnectionsEnabled field to given value.

### HasAllowBridgeConnectionsEnabled

`func (o *ClientProfileSummary) HasAllowBridgeConnectionsEnabled() bool`

HasAllowBridgeConnectionsEnabled returns a boolean if a field has been set.

### GetAllowGuaranteedEndpointCreateEnabled

`func (o *ClientProfileSummary) GetAllowGuaranteedEndpointCreateEnabled() bool`

GetAllowGuaranteedEndpointCreateEnabled returns the AllowGuaranteedEndpointCreateEnabled field if non-nil, zero value otherwise.

### GetAllowGuaranteedEndpointCreateEnabledOk

`func (o *ClientProfileSummary) GetAllowGuaranteedEndpointCreateEnabledOk() (*bool, bool)`

GetAllowGuaranteedEndpointCreateEnabledOk returns a tuple with the AllowGuaranteedEndpointCreateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedEndpointCreateEnabled

`func (o *ClientProfileSummary) SetAllowGuaranteedEndpointCreateEnabled(v bool)`

SetAllowGuaranteedEndpointCreateEnabled sets AllowGuaranteedEndpointCreateEnabled field to given value.

### HasAllowGuaranteedEndpointCreateEnabled

`func (o *ClientProfileSummary) HasAllowGuaranteedEndpointCreateEnabled() bool`

HasAllowGuaranteedEndpointCreateEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


