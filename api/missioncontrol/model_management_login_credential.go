/*
MISSION CONTROL

Documentation for Mission Control API

API version: V2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ManagementLoginCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementLoginCredential{}

// ManagementLoginCredential The management admin login credentials
type ManagementLoginCredential struct {
	// The username to log into the event broker service.
	Username *string `json:"username,omitempty"`
	// The password to log into the event broker service.
	Password *string `json:"password,omitempty"`
	// The token for management access.
	Token *string `json:"token,omitempty"`
}

// NewManagementLoginCredential instantiates a new ManagementLoginCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementLoginCredential() *ManagementLoginCredential {
	this := ManagementLoginCredential{}
	return &this
}

// NewManagementLoginCredentialWithDefaults instantiates a new ManagementLoginCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementLoginCredentialWithDefaults() *ManagementLoginCredential {
	this := ManagementLoginCredential{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ManagementLoginCredential) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementLoginCredential) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ManagementLoginCredential) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ManagementLoginCredential) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ManagementLoginCredential) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementLoginCredential) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ManagementLoginCredential) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ManagementLoginCredential) SetPassword(v string) {
	o.Password = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ManagementLoginCredential) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementLoginCredential) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ManagementLoginCredential) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ManagementLoginCredential) SetToken(v string) {
	o.Token = &v
}

func (o ManagementLoginCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementLoginCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableManagementLoginCredential struct {
	value *ManagementLoginCredential
	isSet bool
}

func (v NullableManagementLoginCredential) Get() *ManagementLoginCredential {
	return v.value
}

func (v *NullableManagementLoginCredential) Set(val *ManagementLoginCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementLoginCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementLoginCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementLoginCredential(val *ManagementLoginCredential) *NullableManagementLoginCredential {
	return &NullableManagementLoginCredential{value: val, isSet: true}
}

func (v NullableManagementLoginCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementLoginCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

