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

// checks if the TimeMetricValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeMetricValue{}

// TimeMetricValue Measurement of the **metricName** recorded at the time of the event.
type TimeMetricValue struct {
	// Amount of the **metricName** recorded at the time of the event. This value triggered the alert.
	Number *float64 `json:"number,omitempty"`
	Units *TimeMetricUnits `json:"units,omitempty"`
}

// NewTimeMetricValue instantiates a new TimeMetricValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeMetricValue() *TimeMetricValue {
	this := TimeMetricValue{}
	var units TimeMetricUnits = TIMEMETRICUNITS_HOURS
	this.Units = &units
	return &this
}

// NewTimeMetricValueWithDefaults instantiates a new TimeMetricValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeMetricValueWithDefaults() *TimeMetricValue {
	this := TimeMetricValue{}
	var units TimeMetricUnits = TIMEMETRICUNITS_HOURS
	this.Units = &units
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *TimeMetricValue) GetNumber() float64 {
	if o == nil || IsNil(o.Number) {
		var ret float64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricValue) GetNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *TimeMetricValue) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given float64 and assigns it to the Number field.
func (o *TimeMetricValue) SetNumber(v float64) {
	o.Number = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *TimeMetricValue) GetUnits() TimeMetricUnits {
	if o == nil || IsNil(o.Units) {
		var ret TimeMetricUnits
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricValue) GetUnitsOk() (*TimeMetricUnits, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *TimeMetricValue) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given TimeMetricUnits and assigns it to the Units field.
func (o *TimeMetricValue) SetUnits(v TimeMetricUnits) {
	o.Units = &v
}

func (o TimeMetricValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeMetricValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: number is readOnly
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}

type NullableTimeMetricValue struct {
	value *TimeMetricValue
	isSet bool
}

func (v NullableTimeMetricValue) Get() *TimeMetricValue {
	return v.value
}

func (v *NullableTimeMetricValue) Set(val *TimeMetricValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeMetricValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeMetricValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeMetricValue(val *TimeMetricValue) *NullableTimeMetricValue {
	return &NullableTimeMetricValue{value: val, isSet: true}
}

func (v NullableTimeMetricValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeMetricValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


