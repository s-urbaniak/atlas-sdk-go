// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the SMSNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMSNotification{}

// SMSNotification SMS notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type SMSNotification struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int `json:"delayMin,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int `json:"intervalMin,omitempty"`
	// Mobile phone number to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.[n].typeName\" : \"SMS\"`.
	MobileNumber *string `json:"mobileNumber,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewSMSNotification instantiates a new SMSNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMSNotification(typeName string) *SMSNotification {
	this := SMSNotification{}
	this.TypeName = typeName
	return &this
}

// NewSMSNotificationWithDefaults instantiates a new SMSNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMSNotificationWithDefaults() *SMSNotification {
	this := SMSNotification{}
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *SMSNotification) GetDelayMin() int {
	if o == nil || IsNil(o.DelayMin) {
		var ret int
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSNotification) GetDelayMinOk() (*int, bool) {
	if o == nil || IsNil(o.DelayMin) {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *SMSNotification) HasDelayMin() bool {
	if o != nil && !IsNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int and assigns it to the DelayMin field.
func (o *SMSNotification) SetDelayMin(v int) {
	o.DelayMin = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *SMSNotification) GetIntervalMin() int {
	if o == nil || IsNil(o.IntervalMin) {
		var ret int
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSNotification) GetIntervalMinOk() (*int, bool) {
	if o == nil || IsNil(o.IntervalMin) {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *SMSNotification) HasIntervalMin() bool {
	if o != nil && !IsNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int and assigns it to the IntervalMin field.
func (o *SMSNotification) SetIntervalMin(v int) {
	o.IntervalMin = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *SMSNotification) GetMobileNumber() string {
	if o == nil || IsNil(o.MobileNumber) {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSNotification) GetMobileNumberOk() (*string, bool) {
	if o == nil || IsNil(o.MobileNumber) {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *SMSNotification) HasMobileNumber() bool {
	if o != nil && !IsNil(o.MobileNumber) {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *SMSNotification) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetTypeName returns the TypeName field value
func (o *SMSNotification) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *SMSNotification) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *SMSNotification) SetTypeName(v string) {
	o.TypeName = v
}

func (o SMSNotification) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o SMSNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !IsNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if !IsNil(o.MobileNumber) {
		toSerialize["mobileNumber"] = o.MobileNumber
	}
	toSerialize["typeName"] = o.TypeName
	return toSerialize, nil
}

type NullableSMSNotification struct {
	value *SMSNotification
	isSet bool
}

func (v NullableSMSNotification) Get() *SMSNotification {
	return v.value
}

func (v *NullableSMSNotification) Set(val *SMSNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSNotification(val *SMSNotification) *NullableSMSNotification {
	return &NullableSMSNotification{value: val, isSet: true}
}

func (v NullableSMSNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
