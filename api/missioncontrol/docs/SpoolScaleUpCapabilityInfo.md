# SpoolScaleUpCapabilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpoolScaleUpCapabilityState** | Pointer to **string** | The current state of  the Scalable Message Spool Size feature. The values can be &#39;INPROGRESS&#39;, &#39;SUPPORTED&#39;, &#39;NOT SUPPORTED&#39;, or &#39;UNKNOWN&#39;. | [optional] 
**SpoolScaleUpTestTimestamp** | Pointer to **time.Time** | The time of the last test was performed, in ISO 8601 date/time format. | [optional] 
**SpoolScaleUpTestMessage** | Pointer to **string** | Message resulting from a self-test of support for the Scalable Message Spool Size feature. | [optional] 

## Methods

### NewSpoolScaleUpCapabilityInfo

`func NewSpoolScaleUpCapabilityInfo() *SpoolScaleUpCapabilityInfo`

NewSpoolScaleUpCapabilityInfo instantiates a new SpoolScaleUpCapabilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpoolScaleUpCapabilityInfoWithDefaults

`func NewSpoolScaleUpCapabilityInfoWithDefaults() *SpoolScaleUpCapabilityInfo`

NewSpoolScaleUpCapabilityInfoWithDefaults instantiates a new SpoolScaleUpCapabilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpoolScaleUpCapabilityState

`func (o *SpoolScaleUpCapabilityInfo) GetSpoolScaleUpCapabilityState() string`

GetSpoolScaleUpCapabilityState returns the SpoolScaleUpCapabilityState field if non-nil, zero value otherwise.

### GetSpoolScaleUpCapabilityStateOk

`func (o *SpoolScaleUpCapabilityInfo) GetSpoolScaleUpCapabilityStateOk() (*string, bool)`

GetSpoolScaleUpCapabilityStateOk returns a tuple with the SpoolScaleUpCapabilityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoolScaleUpCapabilityState

`func (o *SpoolScaleUpCapabilityInfo) SetSpoolScaleUpCapabilityState(v string)`

SetSpoolScaleUpCapabilityState sets SpoolScaleUpCapabilityState field to given value.

### HasSpoolScaleUpCapabilityState

`func (o *SpoolScaleUpCapabilityInfo) HasSpoolScaleUpCapabilityState() bool`

HasSpoolScaleUpCapabilityState returns a boolean if a field has been set.

### GetSpoolScaleUpTestTimestamp

`func (o *SpoolScaleUpCapabilityInfo) GetSpoolScaleUpTestTimestamp() time.Time`

GetSpoolScaleUpTestTimestamp returns the SpoolScaleUpTestTimestamp field if non-nil, zero value otherwise.

### GetSpoolScaleUpTestTimestampOk

`func (o *SpoolScaleUpCapabilityInfo) GetSpoolScaleUpTestTimestampOk() (*time.Time, bool)`

GetSpoolScaleUpTestTimestampOk returns a tuple with the SpoolScaleUpTestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoolScaleUpTestTimestamp

`func (o *SpoolScaleUpCapabilityInfo) SetSpoolScaleUpTestTimestamp(v time.Time)`

SetSpoolScaleUpTestTimestamp sets SpoolScaleUpTestTimestamp field to given value.

### HasSpoolScaleUpTestTimestamp

`func (o *SpoolScaleUpCapabilityInfo) HasSpoolScaleUpTestTimestamp() bool`

HasSpoolScaleUpTestTimestamp returns a boolean if a field has been set.

### GetSpoolScaleUpTestMessage

`func (o *SpoolScaleUpCapabilityInfo) GetSpoolScaleUpTestMessage() string`

GetSpoolScaleUpTestMessage returns the SpoolScaleUpTestMessage field if non-nil, zero value otherwise.

### GetSpoolScaleUpTestMessageOk

`func (o *SpoolScaleUpCapabilityInfo) GetSpoolScaleUpTestMessageOk() (*string, bool)`

GetSpoolScaleUpTestMessageOk returns a tuple with the SpoolScaleUpTestMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoolScaleUpTestMessage

`func (o *SpoolScaleUpCapabilityInfo) SetSpoolScaleUpTestMessage(v string)`

SetSpoolScaleUpTestMessage sets SpoolScaleUpTestMessage field to given value.

### HasSpoolScaleUpTestMessage

`func (o *SpoolScaleUpCapabilityInfo) HasSpoolScaleUpTestMessage() bool`

HasSpoolScaleUpTestMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


