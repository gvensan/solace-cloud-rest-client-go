# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Longitude** | **string** | The longitude coordinate of the location for the datacenter. | 
**Latitude** | **string** | The latitude coordinate of the location for the datacenter. | 
**Continent** | **string** | The name of the continent where the datacenter is located. | 

## Methods

### NewLocation

`func NewLocation(longitude string, latitude string, continent string, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLongitude

`func (o *Location) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Location) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Location) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.


### GetLatitude

`func (o *Location) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Location) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Location) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.


### GetContinent

`func (o *Location) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *Location) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *Location) SetContinent(v string)`

SetContinent sets Continent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


