# ManagementLoginCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username to log into the event broker service. | [optional] 
**Password** | Pointer to **string** | The password to log into the event broker service. | [optional] 
**Token** | Pointer to **string** | The token for management access. | [optional] 

## Methods

### NewManagementLoginCredential

`func NewManagementLoginCredential() *ManagementLoginCredential`

NewManagementLoginCredential instantiates a new ManagementLoginCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementLoginCredentialWithDefaults

`func NewManagementLoginCredentialWithDefaults() *ManagementLoginCredential`

NewManagementLoginCredentialWithDefaults instantiates a new ManagementLoginCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ManagementLoginCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ManagementLoginCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ManagementLoginCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ManagementLoginCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ManagementLoginCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ManagementLoginCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ManagementLoginCredential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ManagementLoginCredential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *ManagementLoginCredential) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ManagementLoginCredential) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ManagementLoginCredential) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ManagementLoginCredential) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


