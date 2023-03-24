/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the ExportStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportStatus{}

// ExportStatus State of the export job for the collections on the replica set only.
type ExportStatus struct {
	// Number of collections on the replica set that MongoDB Cloud exported.
	ExportedCollections *int32 `json:"exportedCollections,omitempty"`
	// Total number of collections on the replica set to export.
	TotalCollections *int32 `json:"totalCollections,omitempty"`
}

// NewExportStatus instantiates a new ExportStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportStatus() *ExportStatus {
	this := ExportStatus{}
	return &this
}

// NewExportStatusWithDefaults instantiates a new ExportStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportStatusWithDefaults() *ExportStatus {
	this := ExportStatus{}
	return &this
}

// GetExportedCollections returns the ExportedCollections field value if set, zero value otherwise.
func (o *ExportStatus) GetExportedCollections() int32 {
	if o == nil || IsNil(o.ExportedCollections) {
		var ret int32
		return ret
	}
	return *o.ExportedCollections
}

// GetExportedCollectionsOk returns a tuple with the ExportedCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportStatus) GetExportedCollectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.ExportedCollections) {
		return nil, false
	}
	return o.ExportedCollections, true
}

// HasExportedCollections returns a boolean if a field has been set.
func (o *ExportStatus) HasExportedCollections() bool {
	if o != nil && !IsNil(o.ExportedCollections) {
		return true
	}

	return false
}

// SetExportedCollections gets a reference to the given int32 and assigns it to the ExportedCollections field.
func (o *ExportStatus) SetExportedCollections(v int32) {
	o.ExportedCollections = &v
}

// GetTotalCollections returns the TotalCollections field value if set, zero value otherwise.
func (o *ExportStatus) GetTotalCollections() int32 {
	if o == nil || IsNil(o.TotalCollections) {
		var ret int32
		return ret
	}
	return *o.TotalCollections
}

// GetTotalCollectionsOk returns a tuple with the TotalCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportStatus) GetTotalCollectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCollections) {
		return nil, false
	}
	return o.TotalCollections, true
}

// HasTotalCollections returns a boolean if a field has been set.
func (o *ExportStatus) HasTotalCollections() bool {
	if o != nil && !IsNil(o.TotalCollections) {
		return true
	}

	return false
}

// SetTotalCollections gets a reference to the given int32 and assigns it to the TotalCollections field.
func (o *ExportStatus) SetTotalCollections(v int32) {
	o.TotalCollections = &v
}

func (o ExportStatus) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ExportStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableExportStatus struct {
	value *ExportStatus
	isSet bool
}

func (v NullableExportStatus) Get() *ExportStatus {
	return v.value
}

func (v *NullableExportStatus) Set(val *ExportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportStatus(val *ExportStatus) *NullableExportStatus {
	return &NullableExportStatus{value: val, isSet: true}
}

func (v NullableExportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


