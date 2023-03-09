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

// checks if the VictorOpsView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VictorOpsView{}

// VictorOpsView Details to integrate one Splunk On-Call account with one MongoDB Cloud project.
type VictorOpsView struct {
	// Key that allows MongoDB Cloud to access your VictorOps account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ApiKey string `json:"apiKey"`
	// Routing key associated with your Splunk On-Call account.
	RoutingKey *string `json:"routingKey,omitempty"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
}

// NewVictorOpsView instantiates a new VictorOpsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVictorOpsView(apiKey string) *VictorOpsView {
	this := VictorOpsView{}
	this.ApiKey = apiKey
	return &this
}

// NewVictorOpsViewWithDefaults instantiates a new VictorOpsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVictorOpsViewWithDefaults() *VictorOpsView {
	this := VictorOpsView{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *VictorOpsView) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *VictorOpsView) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *VictorOpsView) SetApiKey(v string) {
	o.ApiKey = v
}

// GetRoutingKey returns the RoutingKey field value if set, zero value otherwise.
func (o *VictorOpsView) GetRoutingKey() string {
	if o == nil || IsNil(o.RoutingKey) {
		var ret string
		return ret
	}
	return *o.RoutingKey
}

// GetRoutingKeyOk returns a tuple with the RoutingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsView) GetRoutingKeyOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingKey) {
		return nil, false
	}
	return o.RoutingKey, true
}

// HasRoutingKey returns a boolean if a field has been set.
func (o *VictorOpsView) HasRoutingKey() bool {
	if o != nil && !IsNil(o.RoutingKey) {
		return true
	}

	return false
}

// SetRoutingKey gets a reference to the given string and assigns it to the RoutingKey field.
func (o *VictorOpsView) SetRoutingKey(v string) {
	o.RoutingKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VictorOpsView) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VictorOpsView) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VictorOpsView) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VictorOpsView) SetType(v string) {
	o.Type = &v
}

func (o VictorOpsView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VictorOpsView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiKey"] = o.ApiKey
	if !IsNil(o.RoutingKey) {
		toSerialize["routingKey"] = o.RoutingKey
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableVictorOpsView struct {
	value *VictorOpsView
	isSet bool
}

func (v NullableVictorOpsView) Get() *VictorOpsView {
	return v.value
}

func (v *NullableVictorOpsView) Set(val *VictorOpsView) {
	v.value = val
	v.isSet = true
}

func (v NullableVictorOpsView) IsSet() bool {
	return v.isSet
}

func (v *NullableVictorOpsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVictorOpsView(val *VictorOpsView) *NullableVictorOpsView {
	return &NullableVictorOpsView{value: val, isSet: true}
}

func (v NullableVictorOpsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVictorOpsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


