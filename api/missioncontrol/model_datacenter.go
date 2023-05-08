/*
MISSION CONTROL

Documentation for Mission Control API

API version: V2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Datacenter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Datacenter{}

// Datacenter struct for Datacenter
type Datacenter struct {
	// The identifier of the datacenter.
	Id *string `json:"id,omitempty"`
	// The type of object for informational purposes.
	Type *string `json:"type,omitempty"`
	// The name of the datacenter.
	Name string `json:"name"`
	// The type of the datacenter, in terms of ownership. The valus can be a Public Region( 'SolacePublic'), a Dedicated Region('SolaceDedicated'), or a Customer-Controlled Region ('CustomerCloud' or 'CustomerOnPrem').
	DatacenterType string `json:"datacenterType"`
	// The name of the cloud provider for the datacenter.
	Provider string `json:"provider"`
	// The operational state of the datacenter. The values can be 'up' or 'down'.
	OperState string `json:"operState"`
	Location *Location `json:"location,omitempty"`
	// The unique identifier representing the user who created the datacenter.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The time the datacenter was created, in ISO 8601 date/time format.
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	// The unique identifier representing the user who last updated the datacenter.
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// The time of the last update was performed on the datacenter, in ISO 8601 date/time format.
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
	// Whether the datacenter is available.
	Available bool `json:"available"`
	// The list of supported service classes in the datacenter.
	SupportedServiceClasses []string `json:"supportedServiceClasses,omitempty"`
	// The version of the Mission Control Agent.
	CloudAgentVersion *string `json:"cloudAgentVersion,omitempty"`
	// The type of the Kubernetes service. The values can be 'LoadBalancer' or 'NodePort'.
	K8sServiceType *string `json:"k8sServiceType,omitempty"`
	// The number of supported private endpoints.
	NumSupportedPrivateEndpoints *int32 `json:"numSupportedPrivateEndpoints,omitempty"`
	// The number of supported public endpoints.
	NumSupportedPublicEndpoints *int32 `json:"numSupportedPublicEndpoints,omitempty"`
	// The identifier of the datacenter's organization.
	OrganizationId *string `json:"organizationId,omitempty"`
	SpoolScaleUpCapabilityInfo *SpoolScaleUpCapabilityInfo `json:"spoolScaleUpCapabilityInfo,omitempty"`
}

// NewDatacenter instantiates a new Datacenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatacenter(name string, datacenterType string, provider string, operState string, available bool) *Datacenter {
	this := Datacenter{}
	this.Name = name
	this.DatacenterType = datacenterType
	this.Provider = provider
	this.OperState = operState
	this.Available = available
	return &this
}

// NewDatacenterWithDefaults instantiates a new Datacenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatacenterWithDefaults() *Datacenter {
	this := Datacenter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Datacenter) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Datacenter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Datacenter) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Datacenter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Datacenter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Datacenter) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *Datacenter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Datacenter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Datacenter) SetName(v string) {
	o.Name = v
}

// GetDatacenterType returns the DatacenterType field value
func (o *Datacenter) GetDatacenterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatacenterType
}

// GetDatacenterTypeOk returns a tuple with the DatacenterType field value
// and a boolean to check if the value has been set.
func (o *Datacenter) GetDatacenterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatacenterType, true
}

// SetDatacenterType sets field value
func (o *Datacenter) SetDatacenterType(v string) {
	o.DatacenterType = v
}

// GetProvider returns the Provider field value
func (o *Datacenter) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Datacenter) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *Datacenter) SetProvider(v string) {
	o.Provider = v
}

// GetOperState returns the OperState field value
func (o *Datacenter) GetOperState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value
// and a boolean to check if the value has been set.
func (o *Datacenter) GetOperStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperState, true
}

// SetOperState sets field value
func (o *Datacenter) SetOperState(v string) {
	o.OperState = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Datacenter) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Datacenter) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *Datacenter) SetLocation(v Location) {
	o.Location = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Datacenter) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Datacenter) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Datacenter) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Datacenter) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Datacenter) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *Datacenter) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Datacenter) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Datacenter) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *Datacenter) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *Datacenter) GetUpdatedTime() time.Time {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *Datacenter) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given time.Time and assigns it to the UpdatedTime field.
func (o *Datacenter) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = &v
}

// GetAvailable returns the Available field value
func (o *Datacenter) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *Datacenter) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *Datacenter) SetAvailable(v bool) {
	o.Available = v
}

// GetSupportedServiceClasses returns the SupportedServiceClasses field value if set, zero value otherwise.
func (o *Datacenter) GetSupportedServiceClasses() []string {
	if o == nil || IsNil(o.SupportedServiceClasses) {
		var ret []string
		return ret
	}
	return o.SupportedServiceClasses
}

// GetSupportedServiceClassesOk returns a tuple with the SupportedServiceClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetSupportedServiceClassesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedServiceClasses) {
		return nil, false
	}
	return o.SupportedServiceClasses, true
}

// HasSupportedServiceClasses returns a boolean if a field has been set.
func (o *Datacenter) HasSupportedServiceClasses() bool {
	if o != nil && !IsNil(o.SupportedServiceClasses) {
		return true
	}

	return false
}

// SetSupportedServiceClasses gets a reference to the given []string and assigns it to the SupportedServiceClasses field.
func (o *Datacenter) SetSupportedServiceClasses(v []string) {
	o.SupportedServiceClasses = v
}

// GetCloudAgentVersion returns the CloudAgentVersion field value if set, zero value otherwise.
func (o *Datacenter) GetCloudAgentVersion() string {
	if o == nil || IsNil(o.CloudAgentVersion) {
		var ret string
		return ret
	}
	return *o.CloudAgentVersion
}

// GetCloudAgentVersionOk returns a tuple with the CloudAgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetCloudAgentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CloudAgentVersion) {
		return nil, false
	}
	return o.CloudAgentVersion, true
}

// HasCloudAgentVersion returns a boolean if a field has been set.
func (o *Datacenter) HasCloudAgentVersion() bool {
	if o != nil && !IsNil(o.CloudAgentVersion) {
		return true
	}

	return false
}

// SetCloudAgentVersion gets a reference to the given string and assigns it to the CloudAgentVersion field.
func (o *Datacenter) SetCloudAgentVersion(v string) {
	o.CloudAgentVersion = &v
}

// GetK8sServiceType returns the K8sServiceType field value if set, zero value otherwise.
func (o *Datacenter) GetK8sServiceType() string {
	if o == nil || IsNil(o.K8sServiceType) {
		var ret string
		return ret
	}
	return *o.K8sServiceType
}

// GetK8sServiceTypeOk returns a tuple with the K8sServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetK8sServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.K8sServiceType) {
		return nil, false
	}
	return o.K8sServiceType, true
}

// HasK8sServiceType returns a boolean if a field has been set.
func (o *Datacenter) HasK8sServiceType() bool {
	if o != nil && !IsNil(o.K8sServiceType) {
		return true
	}

	return false
}

// SetK8sServiceType gets a reference to the given string and assigns it to the K8sServiceType field.
func (o *Datacenter) SetK8sServiceType(v string) {
	o.K8sServiceType = &v
}

// GetNumSupportedPrivateEndpoints returns the NumSupportedPrivateEndpoints field value if set, zero value otherwise.
func (o *Datacenter) GetNumSupportedPrivateEndpoints() int32 {
	if o == nil || IsNil(o.NumSupportedPrivateEndpoints) {
		var ret int32
		return ret
	}
	return *o.NumSupportedPrivateEndpoints
}

// GetNumSupportedPrivateEndpointsOk returns a tuple with the NumSupportedPrivateEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetNumSupportedPrivateEndpointsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumSupportedPrivateEndpoints) {
		return nil, false
	}
	return o.NumSupportedPrivateEndpoints, true
}

// HasNumSupportedPrivateEndpoints returns a boolean if a field has been set.
func (o *Datacenter) HasNumSupportedPrivateEndpoints() bool {
	if o != nil && !IsNil(o.NumSupportedPrivateEndpoints) {
		return true
	}

	return false
}

// SetNumSupportedPrivateEndpoints gets a reference to the given int32 and assigns it to the NumSupportedPrivateEndpoints field.
func (o *Datacenter) SetNumSupportedPrivateEndpoints(v int32) {
	o.NumSupportedPrivateEndpoints = &v
}

// GetNumSupportedPublicEndpoints returns the NumSupportedPublicEndpoints field value if set, zero value otherwise.
func (o *Datacenter) GetNumSupportedPublicEndpoints() int32 {
	if o == nil || IsNil(o.NumSupportedPublicEndpoints) {
		var ret int32
		return ret
	}
	return *o.NumSupportedPublicEndpoints
}

// GetNumSupportedPublicEndpointsOk returns a tuple with the NumSupportedPublicEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetNumSupportedPublicEndpointsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumSupportedPublicEndpoints) {
		return nil, false
	}
	return o.NumSupportedPublicEndpoints, true
}

// HasNumSupportedPublicEndpoints returns a boolean if a field has been set.
func (o *Datacenter) HasNumSupportedPublicEndpoints() bool {
	if o != nil && !IsNil(o.NumSupportedPublicEndpoints) {
		return true
	}

	return false
}

// SetNumSupportedPublicEndpoints gets a reference to the given int32 and assigns it to the NumSupportedPublicEndpoints field.
func (o *Datacenter) SetNumSupportedPublicEndpoints(v int32) {
	o.NumSupportedPublicEndpoints = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *Datacenter) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Datacenter) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *Datacenter) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetSpoolScaleUpCapabilityInfo returns the SpoolScaleUpCapabilityInfo field value if set, zero value otherwise.
func (o *Datacenter) GetSpoolScaleUpCapabilityInfo() SpoolScaleUpCapabilityInfo {
	if o == nil || IsNil(o.SpoolScaleUpCapabilityInfo) {
		var ret SpoolScaleUpCapabilityInfo
		return ret
	}
	return *o.SpoolScaleUpCapabilityInfo
}

// GetSpoolScaleUpCapabilityInfoOk returns a tuple with the SpoolScaleUpCapabilityInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datacenter) GetSpoolScaleUpCapabilityInfoOk() (*SpoolScaleUpCapabilityInfo, bool) {
	if o == nil || IsNil(o.SpoolScaleUpCapabilityInfo) {
		return nil, false
	}
	return o.SpoolScaleUpCapabilityInfo, true
}

// HasSpoolScaleUpCapabilityInfo returns a boolean if a field has been set.
func (o *Datacenter) HasSpoolScaleUpCapabilityInfo() bool {
	if o != nil && !IsNil(o.SpoolScaleUpCapabilityInfo) {
		return true
	}

	return false
}

// SetSpoolScaleUpCapabilityInfo gets a reference to the given SpoolScaleUpCapabilityInfo and assigns it to the SpoolScaleUpCapabilityInfo field.
func (o *Datacenter) SetSpoolScaleUpCapabilityInfo(v SpoolScaleUpCapabilityInfo) {
	o.SpoolScaleUpCapabilityInfo = &v
}

func (o Datacenter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Datacenter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	toSerialize["datacenterType"] = o.DatacenterType
	toSerialize["provider"] = o.Provider
	toSerialize["operState"] = o.OperState
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updatedTime"] = o.UpdatedTime
	}
	toSerialize["available"] = o.Available
	if !IsNil(o.SupportedServiceClasses) {
		toSerialize["supportedServiceClasses"] = o.SupportedServiceClasses
	}
	if !IsNil(o.CloudAgentVersion) {
		toSerialize["cloudAgentVersion"] = o.CloudAgentVersion
	}
	if !IsNil(o.K8sServiceType) {
		toSerialize["k8sServiceType"] = o.K8sServiceType
	}
	if !IsNil(o.NumSupportedPrivateEndpoints) {
		toSerialize["numSupportedPrivateEndpoints"] = o.NumSupportedPrivateEndpoints
	}
	if !IsNil(o.NumSupportedPublicEndpoints) {
		toSerialize["numSupportedPublicEndpoints"] = o.NumSupportedPublicEndpoints
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.SpoolScaleUpCapabilityInfo) {
		toSerialize["spoolScaleUpCapabilityInfo"] = o.SpoolScaleUpCapabilityInfo
	}
	return toSerialize, nil
}

type NullableDatacenter struct {
	value *Datacenter
	isSet bool
}

func (v NullableDatacenter) Get() *Datacenter {
	return v.value
}

func (v *NullableDatacenter) Set(val *Datacenter) {
	v.value = val
	v.isSet = true
}

func (v NullableDatacenter) IsSet() bool {
	return v.isSet
}

func (v *NullableDatacenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatacenter(val *Datacenter) *NullableDatacenter {
	return &NullableDatacenter{value: val, isSet: true}
}

func (v NullableDatacenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatacenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


