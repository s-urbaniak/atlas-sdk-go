// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the GreaterThanDaysThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GreaterThanDaysThreshold{}

// GreaterThanDaysThreshold Threshold value that triggers an alert.
type GreaterThanDaysThreshold struct {
	// Comparison operator to apply when checking the current metric value.
	Operator *string `json:"operator,omitempty"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *int `json:"threshold,omitempty"`
	// Element used to express the quantity. This can be an element of time, storage capacity, and the like.
	Units *string `json:"units,omitempty"`
}

// NewGreaterThanDaysThreshold instantiates a new GreaterThanDaysThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreaterThanDaysThreshold() *GreaterThanDaysThreshold {
	this := GreaterThanDaysThreshold{}
	return &this
}

// NewGreaterThanDaysThresholdWithDefaults instantiates a new GreaterThanDaysThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreaterThanDaysThresholdWithDefaults() *GreaterThanDaysThreshold {
	this := GreaterThanDaysThreshold{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *GreaterThanDaysThreshold) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanDaysThreshold) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *GreaterThanDaysThreshold) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *GreaterThanDaysThreshold) SetOperator(v string) {
	o.Operator = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *GreaterThanDaysThreshold) GetThreshold() int {
	if o == nil || IsNil(o.Threshold) {
		var ret int
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanDaysThreshold) GetThresholdOk() (*int, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *GreaterThanDaysThreshold) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int and assigns it to the Threshold field.
func (o *GreaterThanDaysThreshold) SetThreshold(v int) {
	o.Threshold = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *GreaterThanDaysThreshold) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanDaysThreshold) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GreaterThanDaysThreshold) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *GreaterThanDaysThreshold) SetUnits(v string) {
	o.Units = &v
}

func (o GreaterThanDaysThreshold) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GreaterThanDaysThreshold) ToMap() (map[string]interface{}, error) {
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

type NullableGreaterThanDaysThreshold struct {
	value *GreaterThanDaysThreshold
	isSet bool
}

func (v NullableGreaterThanDaysThreshold) Get() *GreaterThanDaysThreshold {
	return v.value
}

func (v *NullableGreaterThanDaysThreshold) Set(val *GreaterThanDaysThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableGreaterThanDaysThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableGreaterThanDaysThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreaterThanDaysThreshold(val *GreaterThanDaysThreshold) *NullableGreaterThanDaysThreshold {
	return &NullableGreaterThanDaysThreshold{value: val, isSet: true}
}

func (v NullableGreaterThanDaysThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreaterThanDaysThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
