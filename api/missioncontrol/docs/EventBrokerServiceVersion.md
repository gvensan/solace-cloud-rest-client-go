# EventBrokerServiceVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Event broker service version. | 
**SupportedServiceClasses** | **[]string** | Supported service classes. | 
**Capabilities** | Pointer to **map[string]interface{}** | Event broker service capabilities. | [optional] 

## Methods

### NewEventBrokerServiceVersion

`func NewEventBrokerServiceVersion(version string, supportedServiceClasses []string, ) *EventBrokerServiceVersion`

NewEventBrokerServiceVersion instantiates a new EventBrokerServiceVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBrokerServiceVersionWithDefaults

`func NewEventBrokerServiceVersionWithDefaults() *EventBrokerServiceVersion`

NewEventBrokerServiceVersionWithDefaults instantiates a new EventBrokerServiceVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *EventBrokerServiceVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EventBrokerServiceVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EventBrokerServiceVersion) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSupportedServiceClasses

`func (o *EventBrokerServiceVersion) GetSupportedServiceClasses() []string`

GetSupportedServiceClasses returns the SupportedServiceClasses field if non-nil, zero value otherwise.

### GetSupportedServiceClassesOk

`func (o *EventBrokerServiceVersion) GetSupportedServiceClassesOk() (*[]string, bool)`

GetSupportedServiceClassesOk returns a tuple with the SupportedServiceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedServiceClasses

`func (o *EventBrokerServiceVersion) SetSupportedServiceClasses(v []string)`

SetSupportedServiceClasses sets SupportedServiceClasses field to given value.


### GetCapabilities

`func (o *EventBrokerServiceVersion) GetCapabilities() map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *EventBrokerServiceVersion) GetCapabilitiesOk() (*map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *EventBrokerServiceVersion) SetCapabilities(v map[string]interface{})`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *EventBrokerServiceVersion) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


