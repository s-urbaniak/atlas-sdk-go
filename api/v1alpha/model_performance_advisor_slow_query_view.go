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

// checks if the PerformanceAdvisorSlowQueryView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceAdvisorSlowQueryView{}

// PerformanceAdvisorSlowQueryView Details of one slow query that the Performance Advisor detected.
type PerformanceAdvisorSlowQueryView struct {
	// Text of the MongoDB log related to this slow query.
	Line *string `json:"line,omitempty"`
	// Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as `<database>.<collection>`.
	Namespace *string `json:"namespace,omitempty"`
}

// NewPerformanceAdvisorSlowQueryView instantiates a new PerformanceAdvisorSlowQueryView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceAdvisorSlowQueryView() *PerformanceAdvisorSlowQueryView {
	this := PerformanceAdvisorSlowQueryView{}
	return &this
}

// NewPerformanceAdvisorSlowQueryViewWithDefaults instantiates a new PerformanceAdvisorSlowQueryView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceAdvisorSlowQueryViewWithDefaults() *PerformanceAdvisorSlowQueryView {
	this := PerformanceAdvisorSlowQueryView{}
	return &this
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *PerformanceAdvisorSlowQueryView) GetLine() string {
	if o == nil || IsNil(o.Line) {
		var ret string
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQueryView) GetLineOk() (*string, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQueryView) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *PerformanceAdvisorSlowQueryView) SetLine(v string) {
	o.Line = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *PerformanceAdvisorSlowQueryView) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorSlowQueryView) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *PerformanceAdvisorSlowQueryView) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *PerformanceAdvisorSlowQueryView) SetNamespace(v string) {
	o.Namespace = &v
}

func (o PerformanceAdvisorSlowQueryView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerformanceAdvisorSlowQueryView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: line is readOnly
	// skip: namespace is readOnly
	return toSerialize, nil
}

type NullablePerformanceAdvisorSlowQueryView struct {
	value *PerformanceAdvisorSlowQueryView
	isSet bool
}

func (v NullablePerformanceAdvisorSlowQueryView) Get() *PerformanceAdvisorSlowQueryView {
	return v.value
}

func (v *NullablePerformanceAdvisorSlowQueryView) Set(val *PerformanceAdvisorSlowQueryView) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceAdvisorSlowQueryView) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceAdvisorSlowQueryView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceAdvisorSlowQueryView(val *PerformanceAdvisorSlowQueryView) *NullablePerformanceAdvisorSlowQueryView {
	return &NullablePerformanceAdvisorSlowQueryView{value: val, isSet: true}
}

func (v NullablePerformanceAdvisorSlowQueryView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceAdvisorSlowQueryView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


