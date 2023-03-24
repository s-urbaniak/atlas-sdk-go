/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the ClusterDescriptionConnectionStringsPrivateEndpointEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterDescriptionConnectionStringsPrivateEndpointEndpoint{}

// ClusterDescriptionConnectionStringsPrivateEndpointEndpoint Details of a private endpoint deployed for this cluster.
type ClusterDescriptionConnectionStringsPrivateEndpointEndpoint struct {
	// Unique string that the cloud provider uses to identify the private endpoint.
	EndpointId *string `json:"endpointId,omitempty"`
	// Cloud provider in which MongoDB Cloud deploys the private endpoint.
	ProviderName *string `json:"providerName,omitempty"`
	// Region where the private endpoint is deployed.
	Region *string `json:"region,omitempty"`
}

// NewClusterDescriptionConnectionStringsPrivateEndpointEndpoint instantiates a new ClusterDescriptionConnectionStringsPrivateEndpointEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDescriptionConnectionStringsPrivateEndpointEndpoint() *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint {
	this := ClusterDescriptionConnectionStringsPrivateEndpointEndpoint{}
	return &this
}

// NewClusterDescriptionConnectionStringsPrivateEndpointEndpointWithDefaults instantiates a new ClusterDescriptionConnectionStringsPrivateEndpointEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDescriptionConnectionStringsPrivateEndpointEndpointWithDefaults() *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint {
	this := ClusterDescriptionConnectionStringsPrivateEndpointEndpoint{}
	return &this
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetEndpointId() string {
	if o == nil || IsNil(o.EndpointId) {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointId) {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) HasEndpointId() bool {
	if o != nil && !IsNil(o.EndpointId) {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) SetEndpointId(v string) {
	o.EndpointId = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) SetRegion(v string) {
	o.Region = &v
}

func (o ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint struct {
	value *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint
	isSet bool
}

func (v NullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint) Get() *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint {
	return v.value
}

func (v *NullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint) Set(val *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint(val *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) *NullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint {
	return &NullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint{value: val, isSet: true}
}

func (v NullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDescriptionConnectionStringsPrivateEndpointEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


