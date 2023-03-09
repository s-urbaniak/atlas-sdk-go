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

// checks if the PolicyView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyView{}

// PolicyView List that contains a document for each backup policy item in the desired backup policy.
type PolicyView struct {
	// Unique 24-hexadecimal digit string that identifies this backup policy.
	Id *string `json:"id,omitempty"`
	// List that contains the specifications for one policy.
	PolicyItems []PolicyItemView `json:"policyItems,omitempty"`
}

// NewPolicyView instantiates a new PolicyView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyView() *PolicyView {
	this := PolicyView{}
	return &this
}

// NewPolicyViewWithDefaults instantiates a new PolicyView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyViewWithDefaults() *PolicyView {
	this := PolicyView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PolicyView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PolicyView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PolicyView) SetId(v string) {
	o.Id = &v
}

// GetPolicyItems returns the PolicyItems field value if set, zero value otherwise.
func (o *PolicyView) GetPolicyItems() []PolicyItemView {
	if o == nil || IsNil(o.PolicyItems) {
		var ret []PolicyItemView
		return ret
	}
	return o.PolicyItems
}

// GetPolicyItemsOk returns a tuple with the PolicyItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyView) GetPolicyItemsOk() ([]PolicyItemView, bool) {
	if o == nil || IsNil(o.PolicyItems) {
		return nil, false
	}
	return o.PolicyItems, true
}

// HasPolicyItems returns a boolean if a field has been set.
func (o *PolicyView) HasPolicyItems() bool {
	if o != nil && !IsNil(o.PolicyItems) {
		return true
	}

	return false
}

// SetPolicyItems gets a reference to the given []PolicyItemView and assigns it to the PolicyItems field.
func (o *PolicyView) SetPolicyItems(v []PolicyItemView) {
	o.PolicyItems = v
}

func (o PolicyView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PolicyItems) {
		toSerialize["policyItems"] = o.PolicyItems
	}
	return toSerialize, nil
}

type NullablePolicyView struct {
	value *PolicyView
	isSet bool
}

func (v NullablePolicyView) Get() *PolicyView {
	return v.value
}

func (v *NullablePolicyView) Set(val *PolicyView) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyView) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyView(val *PolicyView) *NullablePolicyView {
	return &NullablePolicyView{value: val, isSet: true}
}

func (v NullablePolicyView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


