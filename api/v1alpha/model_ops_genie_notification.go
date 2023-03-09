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

// checks if the OpsGenieNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpsGenieNotification{}

// OpsGenieNotification OpsGenie notification configuration for MongoDB Cloud to send information when an event triggers an alert condition.
type OpsGenieNotification struct {
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// API Key that MongoDB Cloud needs to send this notification via Opsgenie. The resource requires this parameter when `\"notifications.[n].typeName\" : \"OPS_GENIE\"`. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	OpsGenieApiKey *string `json:"opsGenieApiKey,omitempty"`
	// Opsgenie region that indicates which API Uniform Resource Locator (URL) to use.
	OpsGenieRegion *string `json:"opsGenieRegion,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName string `json:"typeName"`
}

// NewOpsGenieNotification instantiates a new OpsGenieNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsGenieNotification(typeName string) *OpsGenieNotification {
	this := OpsGenieNotification{}
	var opsGenieRegion string = "US"
	this.OpsGenieRegion = &opsGenieRegion
	this.TypeName = typeName
	return &this
}

// NewOpsGenieNotificationWithDefaults instantiates a new OpsGenieNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsGenieNotificationWithDefaults() *OpsGenieNotification {
	this := OpsGenieNotification{}
	var opsGenieRegion string = "US"
	this.OpsGenieRegion = &opsGenieRegion
	return &this
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *OpsGenieNotification) GetDelayMin() int32 {
	if o == nil || IsNil(o.DelayMin) {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsGenieNotification) GetDelayMinOk() (*int32, bool) {
	if o == nil || IsNil(o.DelayMin) {
		return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *OpsGenieNotification) HasDelayMin() bool {
	if o != nil && !IsNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *OpsGenieNotification) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *OpsGenieNotification) GetIntervalMin() int32 {
	if o == nil || IsNil(o.IntervalMin) {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsGenieNotification) GetIntervalMinOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalMin) {
		return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *OpsGenieNotification) HasIntervalMin() bool {
	if o != nil && !IsNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *OpsGenieNotification) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetOpsGenieApiKey returns the OpsGenieApiKey field value if set, zero value otherwise.
func (o *OpsGenieNotification) GetOpsGenieApiKey() string {
	if o == nil || IsNil(o.OpsGenieApiKey) {
		var ret string
		return ret
	}
	return *o.OpsGenieApiKey
}

// GetOpsGenieApiKeyOk returns a tuple with the OpsGenieApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsGenieNotification) GetOpsGenieApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OpsGenieApiKey) {
		return nil, false
	}
	return o.OpsGenieApiKey, true
}

// HasOpsGenieApiKey returns a boolean if a field has been set.
func (o *OpsGenieNotification) HasOpsGenieApiKey() bool {
	if o != nil && !IsNil(o.OpsGenieApiKey) {
		return true
	}

	return false
}

// SetOpsGenieApiKey gets a reference to the given string and assigns it to the OpsGenieApiKey field.
func (o *OpsGenieNotification) SetOpsGenieApiKey(v string) {
	o.OpsGenieApiKey = &v
}

// GetOpsGenieRegion returns the OpsGenieRegion field value if set, zero value otherwise.
func (o *OpsGenieNotification) GetOpsGenieRegion() string {
	if o == nil || IsNil(o.OpsGenieRegion) {
		var ret string
		return ret
	}
	return *o.OpsGenieRegion
}

// GetOpsGenieRegionOk returns a tuple with the OpsGenieRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsGenieNotification) GetOpsGenieRegionOk() (*string, bool) {
	if o == nil || IsNil(o.OpsGenieRegion) {
		return nil, false
	}
	return o.OpsGenieRegion, true
}

// HasOpsGenieRegion returns a boolean if a field has been set.
func (o *OpsGenieNotification) HasOpsGenieRegion() bool {
	if o != nil && !IsNil(o.OpsGenieRegion) {
		return true
	}

	return false
}

// SetOpsGenieRegion gets a reference to the given string and assigns it to the OpsGenieRegion field.
func (o *OpsGenieNotification) SetOpsGenieRegion(v string) {
	o.OpsGenieRegion = &v
}

// GetTypeName returns the TypeName field value
func (o *OpsGenieNotification) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *OpsGenieNotification) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *OpsGenieNotification) SetTypeName(v string) {
	o.TypeName = v
}

func (o OpsGenieNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpsGenieNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !IsNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if !IsNil(o.OpsGenieApiKey) {
		toSerialize["opsGenieApiKey"] = o.OpsGenieApiKey
	}
	if !IsNil(o.OpsGenieRegion) {
		toSerialize["opsGenieRegion"] = o.OpsGenieRegion
	}
	toSerialize["typeName"] = o.TypeName
	return toSerialize, nil
}

type NullableOpsGenieNotification struct {
	value *OpsGenieNotification
	isSet bool
}

func (v NullableOpsGenieNotification) Get() *OpsGenieNotification {
	return v.value
}

func (v *NullableOpsGenieNotification) Set(val *OpsGenieNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsGenieNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsGenieNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsGenieNotification(val *OpsGenieNotification) *NullableOpsGenieNotification {
	return &NullableOpsGenieNotification{value: val, isSet: true}
}

func (v NullableOpsGenieNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsGenieNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


