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

// WeeklyScheduleView struct for WeeklyScheduleView
type WeeklyScheduleView struct {
	// Day of the week when the scheduled archive starts. The week starts with Monday (`1`) and ends with Sunday (`7`).
	DayOfWeek *int32 `json:"dayOfWeek,omitempty"`
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

// NewWeeklyScheduleView instantiates a new WeeklyScheduleView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeeklyScheduleView() *WeeklyScheduleView {
	this := WeeklyScheduleView{}
	return &this
}

// NewWeeklyScheduleViewWithDefaults instantiates a new WeeklyScheduleView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeeklyScheduleViewWithDefaults() *WeeklyScheduleView {
	this := WeeklyScheduleView{}
	return &this
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *WeeklyScheduleView) GetDayOfWeek() int32 {
	if o == nil || o.DayOfWeek == nil {
		var ret int32
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleView) GetDayOfWeekOk() (*int32, bool) {
	if o == nil || o.DayOfWeek == nil {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *WeeklyScheduleView) HasDayOfWeek() bool {
	if o != nil && o.DayOfWeek != nil {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given int32 and assigns it to the DayOfWeek field.
func (o *WeeklyScheduleView) SetDayOfWeek(v int32) {
	o.DayOfWeek = &v
}

// GetEndHour returns the EndHour field value if set, zero value otherwise.
func (o *WeeklyScheduleView) GetEndHour() int32 {
	if o == nil || o.EndHour == nil {
		var ret int32
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleView) GetEndHourOk() (*int32, bool) {
	if o == nil || o.EndHour == nil {
		return nil, false
	}
	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *WeeklyScheduleView) HasEndHour() bool {
	if o != nil && o.EndHour != nil {
		return true
	}

	return false
}

// SetEndHour gets a reference to the given int32 and assigns it to the EndHour field.
func (o *WeeklyScheduleView) SetEndHour(v int32) {
	o.EndHour = &v
}

// GetEndMinute returns the EndMinute field value if set, zero value otherwise.
func (o *WeeklyScheduleView) GetEndMinute() int32 {
	if o == nil || o.EndMinute == nil {
		var ret int32
		return ret
	}
	return *o.EndMinute
}

// GetEndMinuteOk returns a tuple with the EndMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleView) GetEndMinuteOk() (*int32, bool) {
	if o == nil || o.EndMinute == nil {
		return nil, false
	}
	return o.EndMinute, true
}

// HasEndMinute returns a boolean if a field has been set.
func (o *WeeklyScheduleView) HasEndMinute() bool {
	if o != nil && o.EndMinute != nil {
		return true
	}

	return false
}

// SetEndMinute gets a reference to the given int32 and assigns it to the EndMinute field.
func (o *WeeklyScheduleView) SetEndMinute(v int32) {
	o.EndMinute = &v
}

// GetStartHour returns the StartHour field value if set, zero value otherwise.
func (o *WeeklyScheduleView) GetStartHour() int32 {
	if o == nil || o.StartHour == nil {
		var ret int32
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleView) GetStartHourOk() (*int32, bool) {
	if o == nil || o.StartHour == nil {
		return nil, false
	}
	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *WeeklyScheduleView) HasStartHour() bool {
	if o != nil && o.StartHour != nil {
		return true
	}

	return false
}

// SetStartHour gets a reference to the given int32 and assigns it to the StartHour field.
func (o *WeeklyScheduleView) SetStartHour(v int32) {
	o.StartHour = &v
}

// GetStartMinute returns the StartMinute field value if set, zero value otherwise.
func (o *WeeklyScheduleView) GetStartMinute() int32 {
	if o == nil || o.StartMinute == nil {
		var ret int32
		return ret
	}
	return *o.StartMinute
}

// GetStartMinuteOk returns a tuple with the StartMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleView) GetStartMinuteOk() (*int32, bool) {
	if o == nil || o.StartMinute == nil {
		return nil, false
	}
	return o.StartMinute, true
}

// HasStartMinute returns a boolean if a field has been set.
func (o *WeeklyScheduleView) HasStartMinute() bool {
	if o != nil && o.StartMinute != nil {
		return true
	}

	return false
}

// SetStartMinute gets a reference to the given int32 and assigns it to the StartMinute field.
func (o *WeeklyScheduleView) SetStartMinute(v int32) {
	o.StartMinute = &v
}

// GetType returns the Type field value
func (o *WeeklyScheduleView) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WeeklyScheduleView) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WeeklyScheduleView) SetType(v string) {
	o.Type = v
}

func (o WeeklyScheduleView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DayOfWeek != nil {
		toSerialize["dayOfWeek"] = o.DayOfWeek
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

type NullableWeeklyScheduleView struct {
	value *WeeklyScheduleView
	isSet bool
}

func (v NullableWeeklyScheduleView) Get() *WeeklyScheduleView {
	return v.value
}

func (v *NullableWeeklyScheduleView) Set(val *WeeklyScheduleView) {
	v.value = val
	v.isSet = true
}

func (v NullableWeeklyScheduleView) IsSet() bool {
	return v.isSet
}

func (v *NullableWeeklyScheduleView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeeklyScheduleView(val *WeeklyScheduleView) *NullableWeeklyScheduleView {
	return &NullableWeeklyScheduleView{value: val, isSet: true}
}

func (v NullableWeeklyScheduleView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeeklyScheduleView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


