# ClientProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AllowGuaranteedMsgSendEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to publish guaranteed messages. The valid values are &#39;true&#39; (allowed) or &#39;false&#39; (not allowed) . The default is &#39;false&#39;. | [optional] 
**AllowGuaranteedMsgReceiveEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to bind to topic endpoints or queues to receive guaranteed messages. The valid values are &#39;true&#39; (allowed) or &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**CompressionEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to transfer data using compression. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;true&#39;. | [optional] 
**ReplicationAllowClientConnectWhenStandbyEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to remain connected to the Message VPN when its replication state is Standby. This situation may occur when the Message VPN Replication state of the event broker service changes from Active to Standby. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**AllowTransactedSessionsEnabled** | Pointer to **bool** | Indicates whether client applications client applications assigned to the client profile are allowed to establish transacted sessions or XA sessions. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**AllowBridgeConnectionsEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to establish Dynamic Messaging Routing (DMR) links (or bridge links) from the current Message VPN to another Message VPN in a separate event broker service. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**AllowGuaranteedEndpointCreateEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to create queues or topic endpoints. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**AllowSharedSubscriptionsEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to use shared subscriptions. The valid values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**ApiQueueManagementCopyFromOnCreateName** | Pointer to **string** | The name of a queue template to copy settings from when a new queue is created by the client using the client profile. If the specified queue template does not exist, creation fails. Deprecated since 2.14. This attribute has been replaced with apiQueueManagementCopyFromOnCreateTemplateName. | [optional] 
**ApiQueueManagementCopyFromOnCreateTemplateName** | Pointer to **string** | The name of a queue template to copy settings from when a new queue is created by the client using the client profile. If the specified queue template does not exist, creation fails.  | [optional] 
**ApiTopicEndpointManagementCopyFromOnCreateName** | Pointer to **string** | The name of a topic endpoint to copy settings from when a new topic endpoint is created by the client using the client profile. If the specified topic endpoint does not exist, creation fails. Deprecated since 2.14. This attribute has been replaced with apiTopicEndpointManagementCopyFromOnCreateTemplateName. | [optional] 
**ApiTopicEndpointManagementCopyFromOnCreateTemplateName** | Pointer to **string** | The name of a topic endpoint to copy settings from when a new topic endpoint is created by the client using the client profile. If the specified topic endpoint does not exist, creation fails. | [optional] 
**ServiceMinKeepaliveTimeout** | Pointer to **int32** | The minimum period of time (in seconds) that the event broker service will tolerate inactivity on the client connection. This keepalive value is also enforced for MQTT and SMF (Solace Message Format) connections. The keepalive timeout value is calculated based on the client provided timeout interval (3 x the keepalive interval for SMF, 1.5 x the keepalive interval for MQTT). The default is 30 and valid ranges are 3–3600. | [optional] 
**ServiceSmfMinKeepaliveEnabled** | Pointer to **bool** | Indicates whether clients using the client profile have the minimum keep-alive timeout enabled for SMF (Solace Message Format) connections. The valid values are &#39;true&#39; (enabled) and &#39;false&#39; (not enabled). The default is &#39;false&#39;. | [optional] 
**ServiceWebInactiveTimeout** | Pointer to **int32** | The number of seconds a Web client has to send a request before its session times out and be terminated for being inactive. The default value is 30 seconds. | [optional] 
**ServiceWebMaxPayload** | Pointer to **int32** | The maximum transport payload size (in bytes) before fragmentation occurs for clients using the client profile. The size of the header is not included.  Solace Message Format (SMF) messages that are sent to a consuming Web client are contained within a Web transport message that the event broker sends in its HTTP response to that client. Each Web transport message that is sent can contain multiple SMF messages or partial SMF messages. The maximum Web payload value sets the maximum number of bytes allowed in a single Web transport message (not including the header). This value determines the number of SMF messages that can be sent in a single HTTP response and the size of the Web transport message sent in the HTTP response. The value range is 300 to 10000000, in bytes. The default is 1000000 bytes.  SMF messages that are sent to a consuming Web client are contained within a Web transport message that the event broker sends in its HTTP response to that client. Each Web transport message that is sent can contain multiple SMF messages or partial SMF messages.  The maximum Web payload value sets the maximum number of bytes allowed in a single Web transport message (not including the header). This value determines the number of SMF messages that can be sent in a single HTTP response and the size of the Web transport message sent in the HTTP response. Note that large SMF messages can be fragmented across Web transport messages to respect the value set for the maximum possible Web payload. | [optional] 
**MaxConnectionCountPerClientUsername** | Pointer to **int32** | The maximum permitted number of simultaneous Web transport client connections to the event broker service that can be made using the same client username account. The default is the maximum value supported by the platform. | [optional] 
**ServiceSmfMaxConnectionCountPerClientUsername** | Pointer to **int32** | The maximum permitted number of simultaneous Solace Message Format (SMF) client connections to the event broker that can be made using the same client username account. The default is the maximum value supported by the platform. | [optional] 
**ServiceWebMaxConnectionCountPerClientUsername** | Pointer to **int32** | The maximum permitted number of simultaneous Web transport client connections to the event broker service that can be made using the same client username account. The default is the maximum value supported by the platform. | [optional] 
**MaxEndpointCountPerClientUsername** | Pointer to **int32** | The maximum number of durable and non-durable queues and topic endpoints that can be owned by the clients using the same client username within a client profile. The default is 1000. | [optional] 
**MaxEgressFlowCount** | Pointer to **int32** | The maximum number of egress flows (that is, Guaranteed Message client receive flows or consumer flows) that can be created by a single client associated with this client profile. The default is 1000. | [optional] 
**MaxIngressFlowCount** | Pointer to **int32** | The maximum number of ingress flows (that is, Guaranteed Message client publish flows) that can be created by a single client associated with this client profile. The default is 1000. | [optional] 
**MaxSubscriptionCount** | Pointer to **int32** | The maximum number of subscriptions for a single client in the client profile. When you set this option, consider the total maximum number of permitted topic subscriptions and the total maximum number of client connections for the type of event broker that is used. That is, to ensure reliable system performance, you must find the right balance for your network, while staying within the system limits mentioned. The balance is generally between allowing the creation of many client applications and allowing each client to add a large number of topic subscriptions. The default varies by platform. | [optional] 
**MaxTransactedSessionCount** | Pointer to **int32** | The maximum number of simultaneous transacted sessions and/or XA sessions allowed for a single client associated with the client profile. The default is 10. | [optional] 
**MaxTransactionCount** | Pointer to **int32** | The total maximum number of simultaneous transactions (both local transactions and transactions within the XA transaction branches) allowed for a single client associated with the client profile. The default varies by platform. | [optional] 
**QueueGuaranteed1MaxDepth** | Pointer to **int32** | The egress queue maximum depth for Guaranteed Messages that represents the number of work units for the client priority queues. The valid range is 2 to 262144. The default is 20000. | [optional] 
**QueueGuaranteed1MinMsgBurst** | Pointer to **int32** | The minimum number of messages that must be on the Guaranteed Message queue before the queue’s depth is checked against the maximum depth setting (thereby allowing the queue to absorb a burst of large messages that exceeds the number of allowed work units). A valid range is 0 to 262144 with the default of 255. The value of 255 is recommended for memory usage optimized configurations, such as message applications; a value of 66000 is for WAN optimized configurations. | [optional] 
**QueueDirect1MaxDepth** | Pointer to **int32** | The egress queue maximum depth for Direct Messages 1 [Class of Service (COS) 1] that represents the number of work units for the client priority queues. The valid range is 2 to 262144. The default is 20000. | [optional] 
**QueueDirect1MinMsgBurst** | Pointer to **int32** | The minimum number of messages that must be on the Direct 1 (COS 1) queue before the queue’s depth is checked against the maximum depth setting (thereby allowing the queue to absorb a burst of large messages that exceeds the number of allowed work units). A valid range is 0 to 262144 with the default of 4. | [optional] 
**QueueDirect2MaxDepth** | Pointer to **int32** | The egress queue maximum depth for Direct Messages 2 (COS 2) that represents the number of work units for the client priority queues. The valid range is 2 to 262144. The default is 20000. | [optional] 
**QueueDirect2MinMsgBurst** | Pointer to **int32** | The minimum number of messages that must be on the Direct 2 (COS 2) queue before the queue’s depth is checked against the maximum depth setting (thereby allowing the queue to absorb a burst of large messages that exceeds the number of allowed work units). A valid range is 0 to 262144 with the default of 4. | [optional] 
**QueueDirect3MaxDepth** | Pointer to **int32** | The egress queue maximum depth for Direct Messages 3 [Class of Service (COS) 3] that represents the number of work units for the client priority queues. The valid range is 2 to 262144. The default is 20000. | [optional] 
**QueueDirect3MinMsgBurst** | Pointer to **int32** | The minimum number of messages that must be on the Direct 3 (COS 3) queue before the queue’s depth is checked against the maximum depth setting (thereby allowing the queue to absorb a burst of large messages that exceeds the number of allowed work units). A valid range is 0 to 262144 with the default of 4. | [optional] 
**QueueControl1MaxDepth** | Pointer to **int32** | The egress queue maximum depth for Control 1 that represents the number of work units for the client priority queues. The valid range is 2 to 262144. The default is 20000. | [optional] 
**QueueControl1MinMsgBurst** | Pointer to **int32** | The minimum number of messages that must be on the Direct 1 (COS 1) queue before the queue’s depth is checked against the maximum depth setting (thereby allowing the queue to absorb a burst of large messages that exceeds the number of allowed work units). A valid range is 0 to 262144 with the default of 4. | [optional] 
**TcpCongestionWindowSize** | Pointer to **int32** | The TCP initial congestion window size is the number of segments that TCP sends before waiting for an acknowledgment from the peer. The TCP initial congestion window size is used when starting up a TCP connection or on recovery from idle (that is, no traffic). Larger values of the initial window allows a connection to come up to speed more quickly. For further details, refer to RFC 2581. Changing the TCP initial congestion window size from its default of 2 results in non-compliance with RFC 2581. Further, if this setting is set to a value too high, it may cause congestion in the network. Contact support before you attempt to change this TCP parameter. | [optional] 
**TcpKeepaliveCount** | Pointer to **int32** | The maximum number of keepalive probes (from 2 to 5 ) that TCP should send before dropping the connection. The default is 5. | [optional] 
**TcpKeepaliveIdleTime** | Pointer to **int32** | The time (from 3 to 120 seconds) a connection must remain idle before TCP begins sending keepalive probes. The default is 3. | [optional] 
**TcpKeepaliveInterval** | Pointer to **int32** | The time (from 1 to 30 seconds) to set as the interval between individual keepalive probes. The default is 1. | [optional] 
**TcpMaxSegmentSize** | Pointer to **int32** | The TCP maximum segment size (MSS) used for client to the event broker service. The default is 1460. | [optional] 
**TcpMaxWindowSize** | Pointer to **int32** | The TCP window size between the event broker service and the client. If the maximum window size is set to less than the bandwidth-delay product, then the TCP connection operates below its maximum potential throughput. If the maximum window is set to less than about twice the bandwidth-delay product, then occasional packet loss causes the TCP connection to operate below its maximum potential throughput as it handles the missing acknowledgments and retransmissions. The default is 256.  Alternately, if the TCP maximum window size is set too large, in the presence of a high offered load, TCP gradually increases its congestion window size until either the congestion window size reaches the maximum window size, or packet loss occurs in the network.  Initially, when the TCP congestion window size is small, the physical bandwidth-delay of the network acts as a memory buffer for packets in flight. But as the congestion window crosses the bandwidth-delay product, the buffering of in-flight packets moves to queues in event broker services and other equipment throughout the network. As the TCP congestion window continues to increase in size, these various equipment queues overflow, causing packet loss and TCP backoff. | [optional] 
**ElidingDelay** | Pointer to **int32** | The amount of time to delay the delivery of messages to a client after the initial message has been delivered. You can specify a value from 0-6000 milliseconds. The default is 0. | [optional] 
**ElidingEnabled** | Pointer to **bool** | Indicates whether clients assigned to the client profile are allowed to use eliding. Eliding allows clients to define a custom per-topic rate for client applications so they can effectively consume relevant messages, rather than queuing up outdated messages. For example, when eliding is configured, clients could receive direct messages for their topic subscriptions at a rate of at most five per second, per topic–even though the source is publishing updates at a much higher rate. The valid values are &#39;true&#39; (enabled) and &#39;false&#39; (not enabled). The default is &#39;false&#39;. | [optional] 
**ElidingMaxTopicCount** | Pointer to **int32** | The maximum number of topics the event broker service that can track for performing the eliding function on each client connection. You can specify a value from 1-320000. The default is 256. | [optional] 
**RejectMsgToSenderOnNoSubscriptionMatchEnabled** | Pointer to **bool** | Indicates whether clients assigned the client profile are allowed to return NACKs (negative acknowledgements) for guaranteed messages that do not have a that do not have a matching guaranteed message subscription. The value values are &#39;true&#39; (allowed) and &#39;false&#39; (not allowed). The default is &#39;false&#39;. | [optional] 
**TlsAllowDowngradeToPlainTextEnabled** | Pointer to **bool** | Indicates whether clients assigned the client profile are allowed to use TLS/SSL encryption to protect the clients&#39; credentials. This setting doesn&#39;t encrypt the data that is transmitted after the clients are authenticated and authorized. This also allows connecting client applications to request to downgrade of their TLS/SSL connections to the Message VPN to a plain-text connection, and if that Message VPN allows TLS/SSL connection downgrades, after the clientsʼ login handshake are finished, their connections are downgraded. This means that the clientsʼ authentication data is still encrypted, but the subsequent application data that is transmitted is sent as non-encrypted plain-text. The valid values are &#39;true (allowed) or &#39;false (not allowed). The default is &#39;true&#39;. | [optional] 
**EventClientProvisionedEndpointSpoolUsageThreshold** | Pointer to [**ProvisionedEndpointSpoolUsageAlertThresholds**](ProvisionedEndpointSpoolUsageAlertThresholds.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewClientProfileRequest

`func NewClientProfileRequest() *ClientProfileRequest`

NewClientProfileRequest instantiates a new ClientProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientProfileRequestWithDefaults

`func NewClientProfileRequestWithDefaults() *ClientProfileRequest`

NewClientProfileRequestWithDefaults instantiates a new ClientProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllowGuaranteedMsgSendEnabled

`func (o *ClientProfileRequest) GetAllowGuaranteedMsgSendEnabled() bool`

GetAllowGuaranteedMsgSendEnabled returns the AllowGuaranteedMsgSendEnabled field if non-nil, zero value otherwise.

### GetAllowGuaranteedMsgSendEnabledOk

`func (o *ClientProfileRequest) GetAllowGuaranteedMsgSendEnabledOk() (*bool, bool)`

GetAllowGuaranteedMsgSendEnabledOk returns a tuple with the AllowGuaranteedMsgSendEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedMsgSendEnabled

`func (o *ClientProfileRequest) SetAllowGuaranteedMsgSendEnabled(v bool)`

SetAllowGuaranteedMsgSendEnabled sets AllowGuaranteedMsgSendEnabled field to given value.

### HasAllowGuaranteedMsgSendEnabled

`func (o *ClientProfileRequest) HasAllowGuaranteedMsgSendEnabled() bool`

HasAllowGuaranteedMsgSendEnabled returns a boolean if a field has been set.

### GetAllowGuaranteedMsgReceiveEnabled

`func (o *ClientProfileRequest) GetAllowGuaranteedMsgReceiveEnabled() bool`

GetAllowGuaranteedMsgReceiveEnabled returns the AllowGuaranteedMsgReceiveEnabled field if non-nil, zero value otherwise.

### GetAllowGuaranteedMsgReceiveEnabledOk

`func (o *ClientProfileRequest) GetAllowGuaranteedMsgReceiveEnabledOk() (*bool, bool)`

GetAllowGuaranteedMsgReceiveEnabledOk returns a tuple with the AllowGuaranteedMsgReceiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedMsgReceiveEnabled

`func (o *ClientProfileRequest) SetAllowGuaranteedMsgReceiveEnabled(v bool)`

SetAllowGuaranteedMsgReceiveEnabled sets AllowGuaranteedMsgReceiveEnabled field to given value.

### HasAllowGuaranteedMsgReceiveEnabled

`func (o *ClientProfileRequest) HasAllowGuaranteedMsgReceiveEnabled() bool`

HasAllowGuaranteedMsgReceiveEnabled returns a boolean if a field has been set.

### GetCompressionEnabled

`func (o *ClientProfileRequest) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *ClientProfileRequest) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *ClientProfileRequest) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *ClientProfileRequest) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### GetReplicationAllowClientConnectWhenStandbyEnabled

`func (o *ClientProfileRequest) GetReplicationAllowClientConnectWhenStandbyEnabled() bool`

GetReplicationAllowClientConnectWhenStandbyEnabled returns the ReplicationAllowClientConnectWhenStandbyEnabled field if non-nil, zero value otherwise.

### GetReplicationAllowClientConnectWhenStandbyEnabledOk

`func (o *ClientProfileRequest) GetReplicationAllowClientConnectWhenStandbyEnabledOk() (*bool, bool)`

GetReplicationAllowClientConnectWhenStandbyEnabledOk returns a tuple with the ReplicationAllowClientConnectWhenStandbyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationAllowClientConnectWhenStandbyEnabled

`func (o *ClientProfileRequest) SetReplicationAllowClientConnectWhenStandbyEnabled(v bool)`

SetReplicationAllowClientConnectWhenStandbyEnabled sets ReplicationAllowClientConnectWhenStandbyEnabled field to given value.

### HasReplicationAllowClientConnectWhenStandbyEnabled

`func (o *ClientProfileRequest) HasReplicationAllowClientConnectWhenStandbyEnabled() bool`

HasReplicationAllowClientConnectWhenStandbyEnabled returns a boolean if a field has been set.

### GetAllowTransactedSessionsEnabled

`func (o *ClientProfileRequest) GetAllowTransactedSessionsEnabled() bool`

GetAllowTransactedSessionsEnabled returns the AllowTransactedSessionsEnabled field if non-nil, zero value otherwise.

### GetAllowTransactedSessionsEnabledOk

`func (o *ClientProfileRequest) GetAllowTransactedSessionsEnabledOk() (*bool, bool)`

GetAllowTransactedSessionsEnabledOk returns a tuple with the AllowTransactedSessionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransactedSessionsEnabled

`func (o *ClientProfileRequest) SetAllowTransactedSessionsEnabled(v bool)`

SetAllowTransactedSessionsEnabled sets AllowTransactedSessionsEnabled field to given value.

### HasAllowTransactedSessionsEnabled

`func (o *ClientProfileRequest) HasAllowTransactedSessionsEnabled() bool`

HasAllowTransactedSessionsEnabled returns a boolean if a field has been set.

### GetAllowBridgeConnectionsEnabled

`func (o *ClientProfileRequest) GetAllowBridgeConnectionsEnabled() bool`

GetAllowBridgeConnectionsEnabled returns the AllowBridgeConnectionsEnabled field if non-nil, zero value otherwise.

### GetAllowBridgeConnectionsEnabledOk

`func (o *ClientProfileRequest) GetAllowBridgeConnectionsEnabledOk() (*bool, bool)`

GetAllowBridgeConnectionsEnabledOk returns a tuple with the AllowBridgeConnectionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBridgeConnectionsEnabled

`func (o *ClientProfileRequest) SetAllowBridgeConnectionsEnabled(v bool)`

SetAllowBridgeConnectionsEnabled sets AllowBridgeConnectionsEnabled field to given value.

### HasAllowBridgeConnectionsEnabled

`func (o *ClientProfileRequest) HasAllowBridgeConnectionsEnabled() bool`

HasAllowBridgeConnectionsEnabled returns a boolean if a field has been set.

### GetAllowGuaranteedEndpointCreateEnabled

`func (o *ClientProfileRequest) GetAllowGuaranteedEndpointCreateEnabled() bool`

GetAllowGuaranteedEndpointCreateEnabled returns the AllowGuaranteedEndpointCreateEnabled field if non-nil, zero value otherwise.

### GetAllowGuaranteedEndpointCreateEnabledOk

`func (o *ClientProfileRequest) GetAllowGuaranteedEndpointCreateEnabledOk() (*bool, bool)`

GetAllowGuaranteedEndpointCreateEnabledOk returns a tuple with the AllowGuaranteedEndpointCreateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedEndpointCreateEnabled

`func (o *ClientProfileRequest) SetAllowGuaranteedEndpointCreateEnabled(v bool)`

SetAllowGuaranteedEndpointCreateEnabled sets AllowGuaranteedEndpointCreateEnabled field to given value.

### HasAllowGuaranteedEndpointCreateEnabled

`func (o *ClientProfileRequest) HasAllowGuaranteedEndpointCreateEnabled() bool`

HasAllowGuaranteedEndpointCreateEnabled returns a boolean if a field has been set.

### GetAllowSharedSubscriptionsEnabled

`func (o *ClientProfileRequest) GetAllowSharedSubscriptionsEnabled() bool`

GetAllowSharedSubscriptionsEnabled returns the AllowSharedSubscriptionsEnabled field if non-nil, zero value otherwise.

### GetAllowSharedSubscriptionsEnabledOk

`func (o *ClientProfileRequest) GetAllowSharedSubscriptionsEnabledOk() (*bool, bool)`

GetAllowSharedSubscriptionsEnabledOk returns a tuple with the AllowSharedSubscriptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSharedSubscriptionsEnabled

`func (o *ClientProfileRequest) SetAllowSharedSubscriptionsEnabled(v bool)`

SetAllowSharedSubscriptionsEnabled sets AllowSharedSubscriptionsEnabled field to given value.

### HasAllowSharedSubscriptionsEnabled

`func (o *ClientProfileRequest) HasAllowSharedSubscriptionsEnabled() bool`

HasAllowSharedSubscriptionsEnabled returns a boolean if a field has been set.

### GetApiQueueManagementCopyFromOnCreateName

`func (o *ClientProfileRequest) GetApiQueueManagementCopyFromOnCreateName() string`

GetApiQueueManagementCopyFromOnCreateName returns the ApiQueueManagementCopyFromOnCreateName field if non-nil, zero value otherwise.

### GetApiQueueManagementCopyFromOnCreateNameOk

`func (o *ClientProfileRequest) GetApiQueueManagementCopyFromOnCreateNameOk() (*string, bool)`

GetApiQueueManagementCopyFromOnCreateNameOk returns a tuple with the ApiQueueManagementCopyFromOnCreateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiQueueManagementCopyFromOnCreateName

`func (o *ClientProfileRequest) SetApiQueueManagementCopyFromOnCreateName(v string)`

SetApiQueueManagementCopyFromOnCreateName sets ApiQueueManagementCopyFromOnCreateName field to given value.

### HasApiQueueManagementCopyFromOnCreateName

`func (o *ClientProfileRequest) HasApiQueueManagementCopyFromOnCreateName() bool`

HasApiQueueManagementCopyFromOnCreateName returns a boolean if a field has been set.

### GetApiQueueManagementCopyFromOnCreateTemplateName

`func (o *ClientProfileRequest) GetApiQueueManagementCopyFromOnCreateTemplateName() string`

GetApiQueueManagementCopyFromOnCreateTemplateName returns the ApiQueueManagementCopyFromOnCreateTemplateName field if non-nil, zero value otherwise.

### GetApiQueueManagementCopyFromOnCreateTemplateNameOk

`func (o *ClientProfileRequest) GetApiQueueManagementCopyFromOnCreateTemplateNameOk() (*string, bool)`

GetApiQueueManagementCopyFromOnCreateTemplateNameOk returns a tuple with the ApiQueueManagementCopyFromOnCreateTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiQueueManagementCopyFromOnCreateTemplateName

`func (o *ClientProfileRequest) SetApiQueueManagementCopyFromOnCreateTemplateName(v string)`

SetApiQueueManagementCopyFromOnCreateTemplateName sets ApiQueueManagementCopyFromOnCreateTemplateName field to given value.

### HasApiQueueManagementCopyFromOnCreateTemplateName

`func (o *ClientProfileRequest) HasApiQueueManagementCopyFromOnCreateTemplateName() bool`

HasApiQueueManagementCopyFromOnCreateTemplateName returns a boolean if a field has been set.

### GetApiTopicEndpointManagementCopyFromOnCreateName

`func (o *ClientProfileRequest) GetApiTopicEndpointManagementCopyFromOnCreateName() string`

GetApiTopicEndpointManagementCopyFromOnCreateName returns the ApiTopicEndpointManagementCopyFromOnCreateName field if non-nil, zero value otherwise.

### GetApiTopicEndpointManagementCopyFromOnCreateNameOk

`func (o *ClientProfileRequest) GetApiTopicEndpointManagementCopyFromOnCreateNameOk() (*string, bool)`

GetApiTopicEndpointManagementCopyFromOnCreateNameOk returns a tuple with the ApiTopicEndpointManagementCopyFromOnCreateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTopicEndpointManagementCopyFromOnCreateName

`func (o *ClientProfileRequest) SetApiTopicEndpointManagementCopyFromOnCreateName(v string)`

SetApiTopicEndpointManagementCopyFromOnCreateName sets ApiTopicEndpointManagementCopyFromOnCreateName field to given value.

### HasApiTopicEndpointManagementCopyFromOnCreateName

`func (o *ClientProfileRequest) HasApiTopicEndpointManagementCopyFromOnCreateName() bool`

HasApiTopicEndpointManagementCopyFromOnCreateName returns a boolean if a field has been set.

### GetApiTopicEndpointManagementCopyFromOnCreateTemplateName

`func (o *ClientProfileRequest) GetApiTopicEndpointManagementCopyFromOnCreateTemplateName() string`

GetApiTopicEndpointManagementCopyFromOnCreateTemplateName returns the ApiTopicEndpointManagementCopyFromOnCreateTemplateName field if non-nil, zero value otherwise.

### GetApiTopicEndpointManagementCopyFromOnCreateTemplateNameOk

`func (o *ClientProfileRequest) GetApiTopicEndpointManagementCopyFromOnCreateTemplateNameOk() (*string, bool)`

GetApiTopicEndpointManagementCopyFromOnCreateTemplateNameOk returns a tuple with the ApiTopicEndpointManagementCopyFromOnCreateTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTopicEndpointManagementCopyFromOnCreateTemplateName

`func (o *ClientProfileRequest) SetApiTopicEndpointManagementCopyFromOnCreateTemplateName(v string)`

SetApiTopicEndpointManagementCopyFromOnCreateTemplateName sets ApiTopicEndpointManagementCopyFromOnCreateTemplateName field to given value.

### HasApiTopicEndpointManagementCopyFromOnCreateTemplateName

`func (o *ClientProfileRequest) HasApiTopicEndpointManagementCopyFromOnCreateTemplateName() bool`

HasApiTopicEndpointManagementCopyFromOnCreateTemplateName returns a boolean if a field has been set.

### GetServiceMinKeepaliveTimeout

`func (o *ClientProfileRequest) GetServiceMinKeepaliveTimeout() int32`

GetServiceMinKeepaliveTimeout returns the ServiceMinKeepaliveTimeout field if non-nil, zero value otherwise.

### GetServiceMinKeepaliveTimeoutOk

`func (o *ClientProfileRequest) GetServiceMinKeepaliveTimeoutOk() (*int32, bool)`

GetServiceMinKeepaliveTimeoutOk returns a tuple with the ServiceMinKeepaliveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMinKeepaliveTimeout

`func (o *ClientProfileRequest) SetServiceMinKeepaliveTimeout(v int32)`

SetServiceMinKeepaliveTimeout sets ServiceMinKeepaliveTimeout field to given value.

### HasServiceMinKeepaliveTimeout

`func (o *ClientProfileRequest) HasServiceMinKeepaliveTimeout() bool`

HasServiceMinKeepaliveTimeout returns a boolean if a field has been set.

### GetServiceSmfMinKeepaliveEnabled

`func (o *ClientProfileRequest) GetServiceSmfMinKeepaliveEnabled() bool`

GetServiceSmfMinKeepaliveEnabled returns the ServiceSmfMinKeepaliveEnabled field if non-nil, zero value otherwise.

### GetServiceSmfMinKeepaliveEnabledOk

`func (o *ClientProfileRequest) GetServiceSmfMinKeepaliveEnabledOk() (*bool, bool)`

GetServiceSmfMinKeepaliveEnabledOk returns a tuple with the ServiceSmfMinKeepaliveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfMinKeepaliveEnabled

`func (o *ClientProfileRequest) SetServiceSmfMinKeepaliveEnabled(v bool)`

SetServiceSmfMinKeepaliveEnabled sets ServiceSmfMinKeepaliveEnabled field to given value.

### HasServiceSmfMinKeepaliveEnabled

`func (o *ClientProfileRequest) HasServiceSmfMinKeepaliveEnabled() bool`

HasServiceSmfMinKeepaliveEnabled returns a boolean if a field has been set.

### GetServiceWebInactiveTimeout

`func (o *ClientProfileRequest) GetServiceWebInactiveTimeout() int32`

GetServiceWebInactiveTimeout returns the ServiceWebInactiveTimeout field if non-nil, zero value otherwise.

### GetServiceWebInactiveTimeoutOk

`func (o *ClientProfileRequest) GetServiceWebInactiveTimeoutOk() (*int32, bool)`

GetServiceWebInactiveTimeoutOk returns a tuple with the ServiceWebInactiveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebInactiveTimeout

`func (o *ClientProfileRequest) SetServiceWebInactiveTimeout(v int32)`

SetServiceWebInactiveTimeout sets ServiceWebInactiveTimeout field to given value.

### HasServiceWebInactiveTimeout

`func (o *ClientProfileRequest) HasServiceWebInactiveTimeout() bool`

HasServiceWebInactiveTimeout returns a boolean if a field has been set.

### GetServiceWebMaxPayload

`func (o *ClientProfileRequest) GetServiceWebMaxPayload() int32`

GetServiceWebMaxPayload returns the ServiceWebMaxPayload field if non-nil, zero value otherwise.

### GetServiceWebMaxPayloadOk

`func (o *ClientProfileRequest) GetServiceWebMaxPayloadOk() (*int32, bool)`

GetServiceWebMaxPayloadOk returns a tuple with the ServiceWebMaxPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebMaxPayload

`func (o *ClientProfileRequest) SetServiceWebMaxPayload(v int32)`

SetServiceWebMaxPayload sets ServiceWebMaxPayload field to given value.

### HasServiceWebMaxPayload

`func (o *ClientProfileRequest) HasServiceWebMaxPayload() bool`

HasServiceWebMaxPayload returns a boolean if a field has been set.

### GetMaxConnectionCountPerClientUsername

`func (o *ClientProfileRequest) GetMaxConnectionCountPerClientUsername() int32`

GetMaxConnectionCountPerClientUsername returns the MaxConnectionCountPerClientUsername field if non-nil, zero value otherwise.

### GetMaxConnectionCountPerClientUsernameOk

`func (o *ClientProfileRequest) GetMaxConnectionCountPerClientUsernameOk() (*int32, bool)`

GetMaxConnectionCountPerClientUsernameOk returns a tuple with the MaxConnectionCountPerClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionCountPerClientUsername

`func (o *ClientProfileRequest) SetMaxConnectionCountPerClientUsername(v int32)`

SetMaxConnectionCountPerClientUsername sets MaxConnectionCountPerClientUsername field to given value.

### HasMaxConnectionCountPerClientUsername

`func (o *ClientProfileRequest) HasMaxConnectionCountPerClientUsername() bool`

HasMaxConnectionCountPerClientUsername returns a boolean if a field has been set.

### GetServiceSmfMaxConnectionCountPerClientUsername

`func (o *ClientProfileRequest) GetServiceSmfMaxConnectionCountPerClientUsername() int32`

GetServiceSmfMaxConnectionCountPerClientUsername returns the ServiceSmfMaxConnectionCountPerClientUsername field if non-nil, zero value otherwise.

### GetServiceSmfMaxConnectionCountPerClientUsernameOk

`func (o *ClientProfileRequest) GetServiceSmfMaxConnectionCountPerClientUsernameOk() (*int32, bool)`

GetServiceSmfMaxConnectionCountPerClientUsernameOk returns a tuple with the ServiceSmfMaxConnectionCountPerClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfMaxConnectionCountPerClientUsername

`func (o *ClientProfileRequest) SetServiceSmfMaxConnectionCountPerClientUsername(v int32)`

SetServiceSmfMaxConnectionCountPerClientUsername sets ServiceSmfMaxConnectionCountPerClientUsername field to given value.

### HasServiceSmfMaxConnectionCountPerClientUsername

`func (o *ClientProfileRequest) HasServiceSmfMaxConnectionCountPerClientUsername() bool`

HasServiceSmfMaxConnectionCountPerClientUsername returns a boolean if a field has been set.

### GetServiceWebMaxConnectionCountPerClientUsername

`func (o *ClientProfileRequest) GetServiceWebMaxConnectionCountPerClientUsername() int32`

GetServiceWebMaxConnectionCountPerClientUsername returns the ServiceWebMaxConnectionCountPerClientUsername field if non-nil, zero value otherwise.

### GetServiceWebMaxConnectionCountPerClientUsernameOk

`func (o *ClientProfileRequest) GetServiceWebMaxConnectionCountPerClientUsernameOk() (*int32, bool)`

GetServiceWebMaxConnectionCountPerClientUsernameOk returns a tuple with the ServiceWebMaxConnectionCountPerClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebMaxConnectionCountPerClientUsername

`func (o *ClientProfileRequest) SetServiceWebMaxConnectionCountPerClientUsername(v int32)`

SetServiceWebMaxConnectionCountPerClientUsername sets ServiceWebMaxConnectionCountPerClientUsername field to given value.

### HasServiceWebMaxConnectionCountPerClientUsername

`func (o *ClientProfileRequest) HasServiceWebMaxConnectionCountPerClientUsername() bool`

HasServiceWebMaxConnectionCountPerClientUsername returns a boolean if a field has been set.

### GetMaxEndpointCountPerClientUsername

`func (o *ClientProfileRequest) GetMaxEndpointCountPerClientUsername() int32`

GetMaxEndpointCountPerClientUsername returns the MaxEndpointCountPerClientUsername field if non-nil, zero value otherwise.

### GetMaxEndpointCountPerClientUsernameOk

`func (o *ClientProfileRequest) GetMaxEndpointCountPerClientUsernameOk() (*int32, bool)`

GetMaxEndpointCountPerClientUsernameOk returns a tuple with the MaxEndpointCountPerClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndpointCountPerClientUsername

`func (o *ClientProfileRequest) SetMaxEndpointCountPerClientUsername(v int32)`

SetMaxEndpointCountPerClientUsername sets MaxEndpointCountPerClientUsername field to given value.

### HasMaxEndpointCountPerClientUsername

`func (o *ClientProfileRequest) HasMaxEndpointCountPerClientUsername() bool`

HasMaxEndpointCountPerClientUsername returns a boolean if a field has been set.

### GetMaxEgressFlowCount

`func (o *ClientProfileRequest) GetMaxEgressFlowCount() int32`

GetMaxEgressFlowCount returns the MaxEgressFlowCount field if non-nil, zero value otherwise.

### GetMaxEgressFlowCountOk

`func (o *ClientProfileRequest) GetMaxEgressFlowCountOk() (*int32, bool)`

GetMaxEgressFlowCountOk returns a tuple with the MaxEgressFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEgressFlowCount

`func (o *ClientProfileRequest) SetMaxEgressFlowCount(v int32)`

SetMaxEgressFlowCount sets MaxEgressFlowCount field to given value.

### HasMaxEgressFlowCount

`func (o *ClientProfileRequest) HasMaxEgressFlowCount() bool`

HasMaxEgressFlowCount returns a boolean if a field has been set.

### GetMaxIngressFlowCount

`func (o *ClientProfileRequest) GetMaxIngressFlowCount() int32`

GetMaxIngressFlowCount returns the MaxIngressFlowCount field if non-nil, zero value otherwise.

### GetMaxIngressFlowCountOk

`func (o *ClientProfileRequest) GetMaxIngressFlowCountOk() (*int32, bool)`

GetMaxIngressFlowCountOk returns a tuple with the MaxIngressFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIngressFlowCount

`func (o *ClientProfileRequest) SetMaxIngressFlowCount(v int32)`

SetMaxIngressFlowCount sets MaxIngressFlowCount field to given value.

### HasMaxIngressFlowCount

`func (o *ClientProfileRequest) HasMaxIngressFlowCount() bool`

HasMaxIngressFlowCount returns a boolean if a field has been set.

### GetMaxSubscriptionCount

`func (o *ClientProfileRequest) GetMaxSubscriptionCount() int32`

GetMaxSubscriptionCount returns the MaxSubscriptionCount field if non-nil, zero value otherwise.

### GetMaxSubscriptionCountOk

`func (o *ClientProfileRequest) GetMaxSubscriptionCountOk() (*int32, bool)`

GetMaxSubscriptionCountOk returns a tuple with the MaxSubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSubscriptionCount

`func (o *ClientProfileRequest) SetMaxSubscriptionCount(v int32)`

SetMaxSubscriptionCount sets MaxSubscriptionCount field to given value.

### HasMaxSubscriptionCount

`func (o *ClientProfileRequest) HasMaxSubscriptionCount() bool`

HasMaxSubscriptionCount returns a boolean if a field has been set.

### GetMaxTransactedSessionCount

`func (o *ClientProfileRequest) GetMaxTransactedSessionCount() int32`

GetMaxTransactedSessionCount returns the MaxTransactedSessionCount field if non-nil, zero value otherwise.

### GetMaxTransactedSessionCountOk

`func (o *ClientProfileRequest) GetMaxTransactedSessionCountOk() (*int32, bool)`

GetMaxTransactedSessionCountOk returns a tuple with the MaxTransactedSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactedSessionCount

`func (o *ClientProfileRequest) SetMaxTransactedSessionCount(v int32)`

SetMaxTransactedSessionCount sets MaxTransactedSessionCount field to given value.

### HasMaxTransactedSessionCount

`func (o *ClientProfileRequest) HasMaxTransactedSessionCount() bool`

HasMaxTransactedSessionCount returns a boolean if a field has been set.

### GetMaxTransactionCount

`func (o *ClientProfileRequest) GetMaxTransactionCount() int32`

GetMaxTransactionCount returns the MaxTransactionCount field if non-nil, zero value otherwise.

### GetMaxTransactionCountOk

`func (o *ClientProfileRequest) GetMaxTransactionCountOk() (*int32, bool)`

GetMaxTransactionCountOk returns a tuple with the MaxTransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactionCount

`func (o *ClientProfileRequest) SetMaxTransactionCount(v int32)`

SetMaxTransactionCount sets MaxTransactionCount field to given value.

### HasMaxTransactionCount

`func (o *ClientProfileRequest) HasMaxTransactionCount() bool`

HasMaxTransactionCount returns a boolean if a field has been set.

### GetQueueGuaranteed1MaxDepth

`func (o *ClientProfileRequest) GetQueueGuaranteed1MaxDepth() int32`

GetQueueGuaranteed1MaxDepth returns the QueueGuaranteed1MaxDepth field if non-nil, zero value otherwise.

### GetQueueGuaranteed1MaxDepthOk

`func (o *ClientProfileRequest) GetQueueGuaranteed1MaxDepthOk() (*int32, bool)`

GetQueueGuaranteed1MaxDepthOk returns a tuple with the QueueGuaranteed1MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueGuaranteed1MaxDepth

`func (o *ClientProfileRequest) SetQueueGuaranteed1MaxDepth(v int32)`

SetQueueGuaranteed1MaxDepth sets QueueGuaranteed1MaxDepth field to given value.

### HasQueueGuaranteed1MaxDepth

`func (o *ClientProfileRequest) HasQueueGuaranteed1MaxDepth() bool`

HasQueueGuaranteed1MaxDepth returns a boolean if a field has been set.

### GetQueueGuaranteed1MinMsgBurst

`func (o *ClientProfileRequest) GetQueueGuaranteed1MinMsgBurst() int32`

GetQueueGuaranteed1MinMsgBurst returns the QueueGuaranteed1MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueGuaranteed1MinMsgBurstOk

`func (o *ClientProfileRequest) GetQueueGuaranteed1MinMsgBurstOk() (*int32, bool)`

GetQueueGuaranteed1MinMsgBurstOk returns a tuple with the QueueGuaranteed1MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueGuaranteed1MinMsgBurst

`func (o *ClientProfileRequest) SetQueueGuaranteed1MinMsgBurst(v int32)`

SetQueueGuaranteed1MinMsgBurst sets QueueGuaranteed1MinMsgBurst field to given value.

### HasQueueGuaranteed1MinMsgBurst

`func (o *ClientProfileRequest) HasQueueGuaranteed1MinMsgBurst() bool`

HasQueueGuaranteed1MinMsgBurst returns a boolean if a field has been set.

### GetQueueDirect1MaxDepth

`func (o *ClientProfileRequest) GetQueueDirect1MaxDepth() int32`

GetQueueDirect1MaxDepth returns the QueueDirect1MaxDepth field if non-nil, zero value otherwise.

### GetQueueDirect1MaxDepthOk

`func (o *ClientProfileRequest) GetQueueDirect1MaxDepthOk() (*int32, bool)`

GetQueueDirect1MaxDepthOk returns a tuple with the QueueDirect1MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect1MaxDepth

`func (o *ClientProfileRequest) SetQueueDirect1MaxDepth(v int32)`

SetQueueDirect1MaxDepth sets QueueDirect1MaxDepth field to given value.

### HasQueueDirect1MaxDepth

`func (o *ClientProfileRequest) HasQueueDirect1MaxDepth() bool`

HasQueueDirect1MaxDepth returns a boolean if a field has been set.

### GetQueueDirect1MinMsgBurst

`func (o *ClientProfileRequest) GetQueueDirect1MinMsgBurst() int32`

GetQueueDirect1MinMsgBurst returns the QueueDirect1MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueDirect1MinMsgBurstOk

`func (o *ClientProfileRequest) GetQueueDirect1MinMsgBurstOk() (*int32, bool)`

GetQueueDirect1MinMsgBurstOk returns a tuple with the QueueDirect1MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect1MinMsgBurst

`func (o *ClientProfileRequest) SetQueueDirect1MinMsgBurst(v int32)`

SetQueueDirect1MinMsgBurst sets QueueDirect1MinMsgBurst field to given value.

### HasQueueDirect1MinMsgBurst

`func (o *ClientProfileRequest) HasQueueDirect1MinMsgBurst() bool`

HasQueueDirect1MinMsgBurst returns a boolean if a field has been set.

### GetQueueDirect2MaxDepth

`func (o *ClientProfileRequest) GetQueueDirect2MaxDepth() int32`

GetQueueDirect2MaxDepth returns the QueueDirect2MaxDepth field if non-nil, zero value otherwise.

### GetQueueDirect2MaxDepthOk

`func (o *ClientProfileRequest) GetQueueDirect2MaxDepthOk() (*int32, bool)`

GetQueueDirect2MaxDepthOk returns a tuple with the QueueDirect2MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect2MaxDepth

`func (o *ClientProfileRequest) SetQueueDirect2MaxDepth(v int32)`

SetQueueDirect2MaxDepth sets QueueDirect2MaxDepth field to given value.

### HasQueueDirect2MaxDepth

`func (o *ClientProfileRequest) HasQueueDirect2MaxDepth() bool`

HasQueueDirect2MaxDepth returns a boolean if a field has been set.

### GetQueueDirect2MinMsgBurst

`func (o *ClientProfileRequest) GetQueueDirect2MinMsgBurst() int32`

GetQueueDirect2MinMsgBurst returns the QueueDirect2MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueDirect2MinMsgBurstOk

`func (o *ClientProfileRequest) GetQueueDirect2MinMsgBurstOk() (*int32, bool)`

GetQueueDirect2MinMsgBurstOk returns a tuple with the QueueDirect2MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect2MinMsgBurst

`func (o *ClientProfileRequest) SetQueueDirect2MinMsgBurst(v int32)`

SetQueueDirect2MinMsgBurst sets QueueDirect2MinMsgBurst field to given value.

### HasQueueDirect2MinMsgBurst

`func (o *ClientProfileRequest) HasQueueDirect2MinMsgBurst() bool`

HasQueueDirect2MinMsgBurst returns a boolean if a field has been set.

### GetQueueDirect3MaxDepth

`func (o *ClientProfileRequest) GetQueueDirect3MaxDepth() int32`

GetQueueDirect3MaxDepth returns the QueueDirect3MaxDepth field if non-nil, zero value otherwise.

### GetQueueDirect3MaxDepthOk

`func (o *ClientProfileRequest) GetQueueDirect3MaxDepthOk() (*int32, bool)`

GetQueueDirect3MaxDepthOk returns a tuple with the QueueDirect3MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect3MaxDepth

`func (o *ClientProfileRequest) SetQueueDirect3MaxDepth(v int32)`

SetQueueDirect3MaxDepth sets QueueDirect3MaxDepth field to given value.

### HasQueueDirect3MaxDepth

`func (o *ClientProfileRequest) HasQueueDirect3MaxDepth() bool`

HasQueueDirect3MaxDepth returns a boolean if a field has been set.

### GetQueueDirect3MinMsgBurst

`func (o *ClientProfileRequest) GetQueueDirect3MinMsgBurst() int32`

GetQueueDirect3MinMsgBurst returns the QueueDirect3MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueDirect3MinMsgBurstOk

`func (o *ClientProfileRequest) GetQueueDirect3MinMsgBurstOk() (*int32, bool)`

GetQueueDirect3MinMsgBurstOk returns a tuple with the QueueDirect3MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect3MinMsgBurst

`func (o *ClientProfileRequest) SetQueueDirect3MinMsgBurst(v int32)`

SetQueueDirect3MinMsgBurst sets QueueDirect3MinMsgBurst field to given value.

### HasQueueDirect3MinMsgBurst

`func (o *ClientProfileRequest) HasQueueDirect3MinMsgBurst() bool`

HasQueueDirect3MinMsgBurst returns a boolean if a field has been set.

### GetQueueControl1MaxDepth

`func (o *ClientProfileRequest) GetQueueControl1MaxDepth() int32`

GetQueueControl1MaxDepth returns the QueueControl1MaxDepth field if non-nil, zero value otherwise.

### GetQueueControl1MaxDepthOk

`func (o *ClientProfileRequest) GetQueueControl1MaxDepthOk() (*int32, bool)`

GetQueueControl1MaxDepthOk returns a tuple with the QueueControl1MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueControl1MaxDepth

`func (o *ClientProfileRequest) SetQueueControl1MaxDepth(v int32)`

SetQueueControl1MaxDepth sets QueueControl1MaxDepth field to given value.

### HasQueueControl1MaxDepth

`func (o *ClientProfileRequest) HasQueueControl1MaxDepth() bool`

HasQueueControl1MaxDepth returns a boolean if a field has been set.

### GetQueueControl1MinMsgBurst

`func (o *ClientProfileRequest) GetQueueControl1MinMsgBurst() int32`

GetQueueControl1MinMsgBurst returns the QueueControl1MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueControl1MinMsgBurstOk

`func (o *ClientProfileRequest) GetQueueControl1MinMsgBurstOk() (*int32, bool)`

GetQueueControl1MinMsgBurstOk returns a tuple with the QueueControl1MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueControl1MinMsgBurst

`func (o *ClientProfileRequest) SetQueueControl1MinMsgBurst(v int32)`

SetQueueControl1MinMsgBurst sets QueueControl1MinMsgBurst field to given value.

### HasQueueControl1MinMsgBurst

`func (o *ClientProfileRequest) HasQueueControl1MinMsgBurst() bool`

HasQueueControl1MinMsgBurst returns a boolean if a field has been set.

### GetTcpCongestionWindowSize

`func (o *ClientProfileRequest) GetTcpCongestionWindowSize() int32`

GetTcpCongestionWindowSize returns the TcpCongestionWindowSize field if non-nil, zero value otherwise.

### GetTcpCongestionWindowSizeOk

`func (o *ClientProfileRequest) GetTcpCongestionWindowSizeOk() (*int32, bool)`

GetTcpCongestionWindowSizeOk returns a tuple with the TcpCongestionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpCongestionWindowSize

`func (o *ClientProfileRequest) SetTcpCongestionWindowSize(v int32)`

SetTcpCongestionWindowSize sets TcpCongestionWindowSize field to given value.

### HasTcpCongestionWindowSize

`func (o *ClientProfileRequest) HasTcpCongestionWindowSize() bool`

HasTcpCongestionWindowSize returns a boolean if a field has been set.

### GetTcpKeepaliveCount

`func (o *ClientProfileRequest) GetTcpKeepaliveCount() int32`

GetTcpKeepaliveCount returns the TcpKeepaliveCount field if non-nil, zero value otherwise.

### GetTcpKeepaliveCountOk

`func (o *ClientProfileRequest) GetTcpKeepaliveCountOk() (*int32, bool)`

GetTcpKeepaliveCountOk returns a tuple with the TcpKeepaliveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpKeepaliveCount

`func (o *ClientProfileRequest) SetTcpKeepaliveCount(v int32)`

SetTcpKeepaliveCount sets TcpKeepaliveCount field to given value.

### HasTcpKeepaliveCount

`func (o *ClientProfileRequest) HasTcpKeepaliveCount() bool`

HasTcpKeepaliveCount returns a boolean if a field has been set.

### GetTcpKeepaliveIdleTime

`func (o *ClientProfileRequest) GetTcpKeepaliveIdleTime() int32`

GetTcpKeepaliveIdleTime returns the TcpKeepaliveIdleTime field if non-nil, zero value otherwise.

### GetTcpKeepaliveIdleTimeOk

`func (o *ClientProfileRequest) GetTcpKeepaliveIdleTimeOk() (*int32, bool)`

GetTcpKeepaliveIdleTimeOk returns a tuple with the TcpKeepaliveIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpKeepaliveIdleTime

`func (o *ClientProfileRequest) SetTcpKeepaliveIdleTime(v int32)`

SetTcpKeepaliveIdleTime sets TcpKeepaliveIdleTime field to given value.

### HasTcpKeepaliveIdleTime

`func (o *ClientProfileRequest) HasTcpKeepaliveIdleTime() bool`

HasTcpKeepaliveIdleTime returns a boolean if a field has been set.

### GetTcpKeepaliveInterval

`func (o *ClientProfileRequest) GetTcpKeepaliveInterval() int32`

GetTcpKeepaliveInterval returns the TcpKeepaliveInterval field if non-nil, zero value otherwise.

### GetTcpKeepaliveIntervalOk

`func (o *ClientProfileRequest) GetTcpKeepaliveIntervalOk() (*int32, bool)`

GetTcpKeepaliveIntervalOk returns a tuple with the TcpKeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpKeepaliveInterval

`func (o *ClientProfileRequest) SetTcpKeepaliveInterval(v int32)`

SetTcpKeepaliveInterval sets TcpKeepaliveInterval field to given value.

### HasTcpKeepaliveInterval

`func (o *ClientProfileRequest) HasTcpKeepaliveInterval() bool`

HasTcpKeepaliveInterval returns a boolean if a field has been set.

### GetTcpMaxSegmentSize

`func (o *ClientProfileRequest) GetTcpMaxSegmentSize() int32`

GetTcpMaxSegmentSize returns the TcpMaxSegmentSize field if non-nil, zero value otherwise.

### GetTcpMaxSegmentSizeOk

`func (o *ClientProfileRequest) GetTcpMaxSegmentSizeOk() (*int32, bool)`

GetTcpMaxSegmentSizeOk returns a tuple with the TcpMaxSegmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMaxSegmentSize

`func (o *ClientProfileRequest) SetTcpMaxSegmentSize(v int32)`

SetTcpMaxSegmentSize sets TcpMaxSegmentSize field to given value.

### HasTcpMaxSegmentSize

`func (o *ClientProfileRequest) HasTcpMaxSegmentSize() bool`

HasTcpMaxSegmentSize returns a boolean if a field has been set.

### GetTcpMaxWindowSize

`func (o *ClientProfileRequest) GetTcpMaxWindowSize() int32`

GetTcpMaxWindowSize returns the TcpMaxWindowSize field if non-nil, zero value otherwise.

### GetTcpMaxWindowSizeOk

`func (o *ClientProfileRequest) GetTcpMaxWindowSizeOk() (*int32, bool)`

GetTcpMaxWindowSizeOk returns a tuple with the TcpMaxWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMaxWindowSize

`func (o *ClientProfileRequest) SetTcpMaxWindowSize(v int32)`

SetTcpMaxWindowSize sets TcpMaxWindowSize field to given value.

### HasTcpMaxWindowSize

`func (o *ClientProfileRequest) HasTcpMaxWindowSize() bool`

HasTcpMaxWindowSize returns a boolean if a field has been set.

### GetElidingDelay

`func (o *ClientProfileRequest) GetElidingDelay() int32`

GetElidingDelay returns the ElidingDelay field if non-nil, zero value otherwise.

### GetElidingDelayOk

`func (o *ClientProfileRequest) GetElidingDelayOk() (*int32, bool)`

GetElidingDelayOk returns a tuple with the ElidingDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElidingDelay

`func (o *ClientProfileRequest) SetElidingDelay(v int32)`

SetElidingDelay sets ElidingDelay field to given value.

### HasElidingDelay

`func (o *ClientProfileRequest) HasElidingDelay() bool`

HasElidingDelay returns a boolean if a field has been set.

### GetElidingEnabled

`func (o *ClientProfileRequest) GetElidingEnabled() bool`

GetElidingEnabled returns the ElidingEnabled field if non-nil, zero value otherwise.

### GetElidingEnabledOk

`func (o *ClientProfileRequest) GetElidingEnabledOk() (*bool, bool)`

GetElidingEnabledOk returns a tuple with the ElidingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElidingEnabled

`func (o *ClientProfileRequest) SetElidingEnabled(v bool)`

SetElidingEnabled sets ElidingEnabled field to given value.

### HasElidingEnabled

`func (o *ClientProfileRequest) HasElidingEnabled() bool`

HasElidingEnabled returns a boolean if a field has been set.

### GetElidingMaxTopicCount

`func (o *ClientProfileRequest) GetElidingMaxTopicCount() int32`

GetElidingMaxTopicCount returns the ElidingMaxTopicCount field if non-nil, zero value otherwise.

### GetElidingMaxTopicCountOk

`func (o *ClientProfileRequest) GetElidingMaxTopicCountOk() (*int32, bool)`

GetElidingMaxTopicCountOk returns a tuple with the ElidingMaxTopicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElidingMaxTopicCount

`func (o *ClientProfileRequest) SetElidingMaxTopicCount(v int32)`

SetElidingMaxTopicCount sets ElidingMaxTopicCount field to given value.

### HasElidingMaxTopicCount

`func (o *ClientProfileRequest) HasElidingMaxTopicCount() bool`

HasElidingMaxTopicCount returns a boolean if a field has been set.

### GetRejectMsgToSenderOnNoSubscriptionMatchEnabled

`func (o *ClientProfileRequest) GetRejectMsgToSenderOnNoSubscriptionMatchEnabled() bool`

GetRejectMsgToSenderOnNoSubscriptionMatchEnabled returns the RejectMsgToSenderOnNoSubscriptionMatchEnabled field if non-nil, zero value otherwise.

### GetRejectMsgToSenderOnNoSubscriptionMatchEnabledOk

`func (o *ClientProfileRequest) GetRejectMsgToSenderOnNoSubscriptionMatchEnabledOk() (*bool, bool)`

GetRejectMsgToSenderOnNoSubscriptionMatchEnabledOk returns a tuple with the RejectMsgToSenderOnNoSubscriptionMatchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMsgToSenderOnNoSubscriptionMatchEnabled

`func (o *ClientProfileRequest) SetRejectMsgToSenderOnNoSubscriptionMatchEnabled(v bool)`

SetRejectMsgToSenderOnNoSubscriptionMatchEnabled sets RejectMsgToSenderOnNoSubscriptionMatchEnabled field to given value.

### HasRejectMsgToSenderOnNoSubscriptionMatchEnabled

`func (o *ClientProfileRequest) HasRejectMsgToSenderOnNoSubscriptionMatchEnabled() bool`

HasRejectMsgToSenderOnNoSubscriptionMatchEnabled returns a boolean if a field has been set.

### GetTlsAllowDowngradeToPlainTextEnabled

`func (o *ClientProfileRequest) GetTlsAllowDowngradeToPlainTextEnabled() bool`

GetTlsAllowDowngradeToPlainTextEnabled returns the TlsAllowDowngradeToPlainTextEnabled field if non-nil, zero value otherwise.

### GetTlsAllowDowngradeToPlainTextEnabledOk

`func (o *ClientProfileRequest) GetTlsAllowDowngradeToPlainTextEnabledOk() (*bool, bool)`

GetTlsAllowDowngradeToPlainTextEnabledOk returns a tuple with the TlsAllowDowngradeToPlainTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAllowDowngradeToPlainTextEnabled

`func (o *ClientProfileRequest) SetTlsAllowDowngradeToPlainTextEnabled(v bool)`

SetTlsAllowDowngradeToPlainTextEnabled sets TlsAllowDowngradeToPlainTextEnabled field to given value.

### HasTlsAllowDowngradeToPlainTextEnabled

`func (o *ClientProfileRequest) HasTlsAllowDowngradeToPlainTextEnabled() bool`

HasTlsAllowDowngradeToPlainTextEnabled returns a boolean if a field has been set.

### GetEventClientProvisionedEndpointSpoolUsageThreshold

`func (o *ClientProfileRequest) GetEventClientProvisionedEndpointSpoolUsageThreshold() ProvisionedEndpointSpoolUsageAlertThresholds`

GetEventClientProvisionedEndpointSpoolUsageThreshold returns the EventClientProvisionedEndpointSpoolUsageThreshold field if non-nil, zero value otherwise.

### GetEventClientProvisionedEndpointSpoolUsageThresholdOk

`func (o *ClientProfileRequest) GetEventClientProvisionedEndpointSpoolUsageThresholdOk() (*ProvisionedEndpointSpoolUsageAlertThresholds, bool)`

GetEventClientProvisionedEndpointSpoolUsageThresholdOk returns a tuple with the EventClientProvisionedEndpointSpoolUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClientProvisionedEndpointSpoolUsageThreshold

`func (o *ClientProfileRequest) SetEventClientProvisionedEndpointSpoolUsageThreshold(v ProvisionedEndpointSpoolUsageAlertThresholds)`

SetEventClientProvisionedEndpointSpoolUsageThreshold sets EventClientProvisionedEndpointSpoolUsageThreshold field to given value.

### HasEventClientProvisionedEndpointSpoolUsageThreshold

`func (o *ClientProfileRequest) HasEventClientProvisionedEndpointSpoolUsageThreshold() bool`

HasEventClientProvisionedEndpointSpoolUsageThreshold returns a boolean if a field has been set.

### GetId

`func (o *ClientProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientProfileRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientProfileRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ClientProfileRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientProfileRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientProfileRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientProfileRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


