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

// checks if the TenantHardwareSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantHardwareSpec{}

// TenantHardwareSpec struct for TenantHardwareSpec
type TenantHardwareSpec struct {
	// Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size.
	InstanceSize *string `json:"instanceSize,omitempty"`
}

// NewTenantHardwareSpec instantiates a new TenantHardwareSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantHardwareSpec() *TenantHardwareSpec {
	this := TenantHardwareSpec{}
	return &this
}

// NewTenantHardwareSpecWithDefaults instantiates a new TenantHardwareSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantHardwareSpecWithDefaults() *TenantHardwareSpec {
	this := TenantHardwareSpec{}
	return &this
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise.
func (o *TenantHardwareSpec) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantHardwareSpec) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}
	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *TenantHardwareSpec) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *TenantHardwareSpec) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

func (o TenantHardwareSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TenantHardwareSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceSize) {
		toSerialize["instanceSize"] = o.InstanceSize
	}
	return toSerialize, nil
}

type NullableTenantHardwareSpec struct {
	value *TenantHardwareSpec
	isSet bool
}

func (v NullableTenantHardwareSpec) Get() *TenantHardwareSpec {
	return v.value
}

func (v *NullableTenantHardwareSpec) Set(val *TenantHardwareSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantHardwareSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantHardwareSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantHardwareSpec(val *TenantHardwareSpec) *NullableTenantHardwareSpec {
	return &NullableTenantHardwareSpec{value: val, isSet: true}
}

func (v NullableTenantHardwareSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantHardwareSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


