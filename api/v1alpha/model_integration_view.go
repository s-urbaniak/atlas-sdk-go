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

// checks if the IntegrationView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationView{}

// IntegrationView Collection of settings that describe third-party integrations.
type IntegrationView struct {
	Type *string `json:"type,omitempty"`
}

// NewIntegrationView instantiates a new IntegrationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationView() *IntegrationView {
	this := IntegrationView{}
	return &this
}

// NewIntegrationViewWithDefaults instantiates a new IntegrationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationViewWithDefaults() *IntegrationView {
	this := IntegrationView{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IntegrationView) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationView) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IntegrationView) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IntegrationView) SetType(v string) {
	o.Type = &v
}

func (o IntegrationView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIntegrationView struct {
	value *IntegrationView
	isSet bool
}

func (v NullableIntegrationView) Get() *IntegrationView {
	return v.value
}

func (v *NullableIntegrationView) Set(val *IntegrationView) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationView) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationView(val *IntegrationView) *NullableIntegrationView {
	return &NullableIntegrationView{value: val, isSet: true}
}

func (v NullableIntegrationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


