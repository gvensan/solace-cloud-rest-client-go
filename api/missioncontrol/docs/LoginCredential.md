# LoginCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username to log into the event broker service. | [optional] 
**Password** | Pointer to **string** | The password to log into the event broker service. | [optional] 

## Methods

### NewLoginCredential

`func NewLoginCredential() *LoginCredential`

NewLoginCredential instantiates a new LoginCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginCredentialWithDefaults

`func NewLoginCredentialWithDefaults() *LoginCredential`

NewLoginCredentialWithDefaults instantiates a new LoginCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *LoginCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LoginCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LoginCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LoginCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *LoginCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginCredential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoginCredential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


