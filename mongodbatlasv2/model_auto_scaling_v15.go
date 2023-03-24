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

// checks if the AutoScalingV15 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoScalingV15{}

// AutoScalingV15 Options that determine how this cluster handles resource scaling.
type AutoScalingV15 struct {
	Compute *ComputeAutoScalingV15 `json:"compute,omitempty"`
	DiskGB *DiskGBAutoScaling `json:"diskGB,omitempty"`
}

// NewAutoScalingV15 instantiates a new AutoScalingV15 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoScalingV15() *AutoScalingV15 {
	this := AutoScalingV15{}
	return &this
}

// NewAutoScalingV15WithDefaults instantiates a new AutoScalingV15 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoScalingV15WithDefaults() *AutoScalingV15 {
	this := AutoScalingV15{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *AutoScalingV15) GetCompute() ComputeAutoScalingV15 {
	if o == nil || IsNil(o.Compute) {
		var ret ComputeAutoScalingV15
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingV15) GetComputeOk() (*ComputeAutoScalingV15, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *AutoScalingV15) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given ComputeAutoScalingV15 and assigns it to the Compute field.
func (o *AutoScalingV15) SetCompute(v ComputeAutoScalingV15) {
	o.Compute = &v
}

// GetDiskGB returns the DiskGB field value if set, zero value otherwise.
func (o *AutoScalingV15) GetDiskGB() DiskGBAutoScaling {
	if o == nil || IsNil(o.DiskGB) {
		var ret DiskGBAutoScaling
		return ret
	}
	return *o.DiskGB
}

// GetDiskGBOk returns a tuple with the DiskGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScalingV15) GetDiskGBOk() (*DiskGBAutoScaling, bool) {
	if o == nil || IsNil(o.DiskGB) {
		return nil, false
	}
	return o.DiskGB, true
}

// HasDiskGB returns a boolean if a field has been set.
func (o *AutoScalingV15) HasDiskGB() bool {
	if o != nil && !IsNil(o.DiskGB) {
		return true
	}

	return false
}

// SetDiskGB gets a reference to the given DiskGBAutoScaling and assigns it to the DiskGB field.
func (o *AutoScalingV15) SetDiskGB(v DiskGBAutoScaling) {
	o.DiskGB = &v
}

func (o AutoScalingV15) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AutoScalingV15) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	if !IsNil(o.DiskGB) {
		toSerialize["diskGB"] = o.DiskGB
	}
	return toSerialize, nil
}

type NullableAutoScalingV15 struct {
	value *AutoScalingV15
	isSet bool
}

func (v NullableAutoScalingV15) Get() *AutoScalingV15 {
	return v.value
}

func (v *NullableAutoScalingV15) Set(val *AutoScalingV15) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoScalingV15) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoScalingV15) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoScalingV15(val *AutoScalingV15) *NullableAutoScalingV15 {
	return &NullableAutoScalingV15{value: val, isSet: true}
}

func (v NullableAutoScalingV15) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoScalingV15) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


