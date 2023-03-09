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

// checks if the PerformanceAdvisorResponseView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceAdvisorResponseView{}

// PerformanceAdvisorResponseView struct for PerformanceAdvisorResponseView
type PerformanceAdvisorResponseView struct {
	// List of query predicates, sorts, and projections that the Performance Advisor suggests.
	Shapes []PerformanceAdvisorShapeView `json:"shapes,omitempty"`
	// List that contains the documents with information about the indexes that the Performance Advisor suggests.
	SuggestedIndexes []PerformanceAdvisorIndexView `json:"suggestedIndexes,omitempty"`
}

// NewPerformanceAdvisorResponseView instantiates a new PerformanceAdvisorResponseView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceAdvisorResponseView() *PerformanceAdvisorResponseView {
	this := PerformanceAdvisorResponseView{}
	return &this
}

// NewPerformanceAdvisorResponseViewWithDefaults instantiates a new PerformanceAdvisorResponseView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceAdvisorResponseViewWithDefaults() *PerformanceAdvisorResponseView {
	this := PerformanceAdvisorResponseView{}
	return &this
}

// GetShapes returns the Shapes field value if set, zero value otherwise.
func (o *PerformanceAdvisorResponseView) GetShapes() []PerformanceAdvisorShapeView {
	if o == nil || IsNil(o.Shapes) {
		var ret []PerformanceAdvisorShapeView
		return ret
	}
	return o.Shapes
}

// GetShapesOk returns a tuple with the Shapes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorResponseView) GetShapesOk() ([]PerformanceAdvisorShapeView, bool) {
	if o == nil || IsNil(o.Shapes) {
		return nil, false
	}
	return o.Shapes, true
}

// HasShapes returns a boolean if a field has been set.
func (o *PerformanceAdvisorResponseView) HasShapes() bool {
	if o != nil && !IsNil(o.Shapes) {
		return true
	}

	return false
}

// SetShapes gets a reference to the given []PerformanceAdvisorShapeView and assigns it to the Shapes field.
func (o *PerformanceAdvisorResponseView) SetShapes(v []PerformanceAdvisorShapeView) {
	o.Shapes = v
}

// GetSuggestedIndexes returns the SuggestedIndexes field value if set, zero value otherwise.
func (o *PerformanceAdvisorResponseView) GetSuggestedIndexes() []PerformanceAdvisorIndexView {
	if o == nil || IsNil(o.SuggestedIndexes) {
		var ret []PerformanceAdvisorIndexView
		return ret
	}
	return o.SuggestedIndexes
}

// GetSuggestedIndexesOk returns a tuple with the SuggestedIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorResponseView) GetSuggestedIndexesOk() ([]PerformanceAdvisorIndexView, bool) {
	if o == nil || IsNil(o.SuggestedIndexes) {
		return nil, false
	}
	return o.SuggestedIndexes, true
}

// HasSuggestedIndexes returns a boolean if a field has been set.
func (o *PerformanceAdvisorResponseView) HasSuggestedIndexes() bool {
	if o != nil && !IsNil(o.SuggestedIndexes) {
		return true
	}

	return false
}

// SetSuggestedIndexes gets a reference to the given []PerformanceAdvisorIndexView and assigns it to the SuggestedIndexes field.
func (o *PerformanceAdvisorResponseView) SetSuggestedIndexes(v []PerformanceAdvisorIndexView) {
	o.SuggestedIndexes = v
}

func (o PerformanceAdvisorResponseView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerformanceAdvisorResponseView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: shapes is readOnly
	// skip: suggestedIndexes is readOnly
	return toSerialize, nil
}

type NullablePerformanceAdvisorResponseView struct {
	value *PerformanceAdvisorResponseView
	isSet bool
}

func (v NullablePerformanceAdvisorResponseView) Get() *PerformanceAdvisorResponseView {
	return v.value
}

func (v *NullablePerformanceAdvisorResponseView) Set(val *PerformanceAdvisorResponseView) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceAdvisorResponseView) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceAdvisorResponseView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceAdvisorResponseView(val *PerformanceAdvisorResponseView) *NullablePerformanceAdvisorResponseView {
	return &NullablePerformanceAdvisorResponseView{value: val, isSet: true}
}

func (v NullablePerformanceAdvisorResponseView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceAdvisorResponseView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


