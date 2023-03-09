/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the GreaterThanTimeThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GreaterThanTimeThreshold{}

// GreaterThanTimeThreshold A Limit that triggers an alert when greater than a time period.
type GreaterThanTimeThreshold struct {
	// Comparison operator to apply when checking the current metric value.
	Operator *string `json:"operator,omitempty"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *int32 `json:"threshold,omitempty"`
	Units *TimeMetricUnits `json:"units,omitempty"`
}

// NewGreaterThanTimeThreshold instantiates a new GreaterThanTimeThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreaterThanTimeThreshold() *GreaterThanTimeThreshold {
	this := GreaterThanTimeThreshold{}
	var units TimeMetricUnits = TIMEMETRICUNITS_HOURS
	this.Units = &units
	return &this
}

// NewGreaterThanTimeThresholdWithDefaults instantiates a new GreaterThanTimeThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreaterThanTimeThresholdWithDefaults() *GreaterThanTimeThreshold {
	this := GreaterThanTimeThreshold{}
	var units TimeMetricUnits = TIMEMETRICUNITS_HOURS
	this.Units = &units
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *GreaterThanTimeThreshold) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanTimeThreshold) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *GreaterThanTimeThreshold) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *GreaterThanTimeThreshold) SetOperator(v string) {
	o.Operator = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *GreaterThanTimeThreshold) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanTimeThreshold) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *GreaterThanTimeThreshold) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *GreaterThanTimeThreshold) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *GreaterThanTimeThreshold) GetUnits() TimeMetricUnits {
	if o == nil || IsNil(o.Units) {
		var ret TimeMetricUnits
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanTimeThreshold) GetUnitsOk() (*TimeMetricUnits, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GreaterThanTimeThreshold) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given TimeMetricUnits and assigns it to the Units field.
func (o *GreaterThanTimeThreshold) SetUnits(v TimeMetricUnits) {
	o.Units = &v
}

func (o GreaterThanTimeThreshold) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GreaterThanTimeThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}

type NullableGreaterThanTimeThreshold struct {
	value *GreaterThanTimeThreshold
	isSet bool
}

func (v NullableGreaterThanTimeThreshold) Get() *GreaterThanTimeThreshold {
	return v.value
}

func (v *NullableGreaterThanTimeThreshold) Set(val *GreaterThanTimeThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableGreaterThanTimeThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableGreaterThanTimeThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreaterThanTimeThreshold(val *GreaterThanTimeThreshold) *NullableGreaterThanTimeThreshold {
	return &NullableGreaterThanTimeThreshold{value: val, isSet: true}
}

func (v NullableGreaterThanTimeThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreaterThanTimeThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


