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

// MonthlyScheduleView struct for MonthlyScheduleView
type MonthlyScheduleView struct {
	// Day of the month when the scheduled archive starts.
	DayOfMonth *int32 `json:"dayOfMonth,omitempty"`
	// Hour of the day when the scheduled window to run one online archive ends.
	EndHour *int32 `json:"endHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive ends.
	EndMinute *int32 `json:"endMinute,omitempty"`
	// Hour of the day when the when the scheduled window to run one online archive starts.
	StartHour *int32 `json:"startHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive starts.
	StartMinute *int32 `json:"startMinute,omitempty"`
	Type string `json:"type"`
}

// NewMonthlyScheduleView instantiates a new MonthlyScheduleView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonthlyScheduleView() *MonthlyScheduleView {
	this := MonthlyScheduleView{}
	return &this
}

// NewMonthlyScheduleViewWithDefaults instantiates a new MonthlyScheduleView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonthlyScheduleViewWithDefaults() *MonthlyScheduleView {
	this := MonthlyScheduleView{}
	return &this
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise.
func (o *MonthlyScheduleView) GetDayOfMonth() int32 {
	if o == nil || o.DayOfMonth == nil {
		var ret int32
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleView) GetDayOfMonthOk() (*int32, bool) {
	if o == nil || o.DayOfMonth == nil {
		return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *MonthlyScheduleView) HasDayOfMonth() bool {
	if o != nil && o.DayOfMonth != nil {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given int32 and assigns it to the DayOfMonth field.
func (o *MonthlyScheduleView) SetDayOfMonth(v int32) {
	o.DayOfMonth = &v
}

// GetEndHour returns the EndHour field value if set, zero value otherwise.
func (o *MonthlyScheduleView) GetEndHour() int32 {
	if o == nil || o.EndHour == nil {
		var ret int32
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleView) GetEndHourOk() (*int32, bool) {
	if o == nil || o.EndHour == nil {
		return nil, false
	}
	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *MonthlyScheduleView) HasEndHour() bool {
	if o != nil && o.EndHour != nil {
		return true
	}

	return false
}

// SetEndHour gets a reference to the given int32 and assigns it to the EndHour field.
func (o *MonthlyScheduleView) SetEndHour(v int32) {
	o.EndHour = &v
}

// GetEndMinute returns the EndMinute field value if set, zero value otherwise.
func (o *MonthlyScheduleView) GetEndMinute() int32 {
	if o == nil || o.EndMinute == nil {
		var ret int32
		return ret
	}
	return *o.EndMinute
}

// GetEndMinuteOk returns a tuple with the EndMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleView) GetEndMinuteOk() (*int32, bool) {
	if o == nil || o.EndMinute == nil {
		return nil, false
	}
	return o.EndMinute, true
}

// HasEndMinute returns a boolean if a field has been set.
func (o *MonthlyScheduleView) HasEndMinute() bool {
	if o != nil && o.EndMinute != nil {
		return true
	}

	return false
}

// SetEndMinute gets a reference to the given int32 and assigns it to the EndMinute field.
func (o *MonthlyScheduleView) SetEndMinute(v int32) {
	o.EndMinute = &v
}

// GetStartHour returns the StartHour field value if set, zero value otherwise.
func (o *MonthlyScheduleView) GetStartHour() int32 {
	if o == nil || o.StartHour == nil {
		var ret int32
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleView) GetStartHourOk() (*int32, bool) {
	if o == nil || o.StartHour == nil {
		return nil, false
	}
	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *MonthlyScheduleView) HasStartHour() bool {
	if o != nil && o.StartHour != nil {
		return true
	}

	return false
}

// SetStartHour gets a reference to the given int32 and assigns it to the StartHour field.
func (o *MonthlyScheduleView) SetStartHour(v int32) {
	o.StartHour = &v
}

// GetStartMinute returns the StartMinute field value if set, zero value otherwise.
func (o *MonthlyScheduleView) GetStartMinute() int32 {
	if o == nil || o.StartMinute == nil {
		var ret int32
		return ret
	}
	return *o.StartMinute
}

// GetStartMinuteOk returns a tuple with the StartMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleView) GetStartMinuteOk() (*int32, bool) {
	if o == nil || o.StartMinute == nil {
		return nil, false
	}
	return o.StartMinute, true
}

// HasStartMinute returns a boolean if a field has been set.
func (o *MonthlyScheduleView) HasStartMinute() bool {
	if o != nil && o.StartMinute != nil {
		return true
	}

	return false
}

// SetStartMinute gets a reference to the given int32 and assigns it to the StartMinute field.
func (o *MonthlyScheduleView) SetStartMinute(v int32) {
	o.StartMinute = &v
}

// GetType returns the Type field value
func (o *MonthlyScheduleView) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleView) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MonthlyScheduleView) SetType(v string) {
	o.Type = v
}

func (o MonthlyScheduleView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DayOfMonth != nil {
		toSerialize["dayOfMonth"] = o.DayOfMonth
	}
	if o.EndHour != nil {
		toSerialize["endHour"] = o.EndHour
	}
	if o.EndMinute != nil {
		toSerialize["endMinute"] = o.EndMinute
	}
	if o.StartHour != nil {
		toSerialize["startHour"] = o.StartHour
	}
	if o.StartMinute != nil {
		toSerialize["startMinute"] = o.StartMinute
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMonthlyScheduleView struct {
	value *MonthlyScheduleView
	isSet bool
}

func (v NullableMonthlyScheduleView) Get() *MonthlyScheduleView {
	return v.value
}

func (v *NullableMonthlyScheduleView) Set(val *MonthlyScheduleView) {
	v.value = val
	v.isSet = true
}

func (v NullableMonthlyScheduleView) IsSet() bool {
	return v.isSet
}

func (v *NullableMonthlyScheduleView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonthlyScheduleView(val *MonthlyScheduleView) *NullableMonthlyScheduleView {
	return &NullableMonthlyScheduleView{value: val, isSet: true}
}

func (v NullableMonthlyScheduleView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonthlyScheduleView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


