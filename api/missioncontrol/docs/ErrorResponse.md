# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | A user-friendly message that describes the error. | [optional] 
**ErrorId** | Pointer to **string** | The UUID (Universally Unique Identifier) that is logged with an appropriate stack trace for a WARN or ERROR log that allows support to better determine what cause the error. | [optional] 
**Meta** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ValidationDetails** | Pointer to **map[string][]string** | When applicable, these are the details of issues with the fields provided for the REST call. | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorResponse

`func NewErrorResponse() *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorId

`func (o *ErrorResponse) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *ErrorResponse) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *ErrorResponse) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *ErrorResponse) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetMeta

`func (o *ErrorResponse) GetMeta() map[string]map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorResponse) GetMetaOk() (*map[string]map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorResponse) SetMeta(v map[string]map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValidationDetails

`func (o *ErrorResponse) GetValidationDetails() map[string][]string`

GetValidationDetails returns the ValidationDetails field if non-nil, zero value otherwise.

### GetValidationDetailsOk

`func (o *ErrorResponse) GetValidationDetailsOk() (*map[string][]string, bool)`

GetValidationDetailsOk returns a tuple with the ValidationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationDetails

`func (o *ErrorResponse) SetValidationDetails(v map[string][]string)`

SetValidationDetails sets ValidationDetails field to given value.

### HasValidationDetails

`func (o *ErrorResponse) HasValidationDetails() bool`

HasValidationDetails returns a boolean if a field has been set.

### GetErrorType

`func (o *ErrorResponse) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ErrorResponse) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ErrorResponse) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *ErrorResponse) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


