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

// checks if the TimeMetricThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeMetricThreshold{}

// TimeMetricThreshold struct for TimeMetricThreshold
type TimeMetricThreshold struct {
	// Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**.
	MetricName *string `json:"metricName,omitempty"`
	// MongoDB Cloud computes the current metric value as an average.
	Mode *string `json:"mode,omitempty"`
	Operator *Operator `json:"operator,omitempty"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *float64 `json:"threshold,omitempty"`
	Units *TimeMetricUnits `json:"units,omitempty"`
}

// NewTimeMetricThreshold instantiates a new TimeMetricThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeMetricThreshold() *TimeMetricThreshold {
	this := TimeMetricThreshold{}
	var units TimeMetricUnits = TIMEMETRICUNITS_HOURS
	this.Units = &units
	return &this
}

// NewTimeMetricThresholdWithDefaults instantiates a new TimeMetricThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeMetricThresholdWithDefaults() *TimeMetricThreshold {
	this := TimeMetricThreshold{}
	var units TimeMetricUnits = TIMEMETRICUNITS_HOURS
	this.Units = &units
	return &this
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *TimeMetricThreshold) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricThreshold) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *TimeMetricThreshold) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *TimeMetricThreshold) SetMetricName(v string) {
	o.MetricName = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *TimeMetricThreshold) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricThreshold) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *TimeMetricThreshold) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *TimeMetricThreshold) SetMode(v string) {
	o.Mode = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *TimeMetricThreshold) GetOperator() Operator {
	if o == nil || IsNil(o.Operator) {
		var ret Operator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricThreshold) GetOperatorOk() (*Operator, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *TimeMetricThreshold) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given Operator and assigns it to the Operator field.
func (o *TimeMetricThreshold) SetOperator(v Operator) {
	o.Operator = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *TimeMetricThreshold) GetThreshold() float64 {
	if o == nil || IsNil(o.Threshold) {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricThreshold) GetThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *TimeMetricThreshold) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *TimeMetricThreshold) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *TimeMetricThreshold) GetUnits() TimeMetricUnits {
	if o == nil || IsNil(o.Units) {
		var ret TimeMetricUnits
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeMetricThreshold) GetUnitsOk() (*TimeMetricUnits, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *TimeMetricThreshold) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given TimeMetricUnits and assigns it to the Units field.
func (o *TimeMetricThreshold) SetUnits(v TimeMetricUnits) {
	o.Units = &v
}

func (o TimeMetricThreshold) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TimeMetricThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetricName) {
		toSerialize["metricName"] = o.MetricName
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
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

type NullableTimeMetricThreshold struct {
	value *TimeMetricThreshold
	isSet bool
}

func (v NullableTimeMetricThreshold) Get() *TimeMetricThreshold {
	return v.value
}

func (v *NullableTimeMetricThreshold) Set(val *TimeMetricThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeMetricThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeMetricThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeMetricThreshold(val *TimeMetricThreshold) *NullableTimeMetricThreshold {
	return &NullableTimeMetricThreshold{value: val, isSet: true}
}

func (v NullableTimeMetricThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeMetricThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


