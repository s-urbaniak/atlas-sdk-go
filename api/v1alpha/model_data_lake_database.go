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

// checks if the DataLakeDatabase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeDatabase{}

// DataLakeDatabase Database associated with this data lake. Databases contain collections and views.
type DataLakeDatabase struct {
	// Array of collections and data sources that map to a ``stores`` data store.
	Collections []DataLakeDatabaseCollection `json:"collections,omitempty"`
	// Maximum number of wildcard collections in the database. This only applies to S3 data sources.
	MaxWildcardCollections *int32 `json:"maxWildcardCollections,omitempty"`
	// Human-readable label that identifies the database to which the data lake maps data.
	Name *string `json:"name,omitempty"`
	// Array of aggregation pipelines that apply to the collection. This only applies to S3 data sources.
	Views []DataLakeView `json:"views,omitempty"`
}

// NewDataLakeDatabase instantiates a new DataLakeDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeDatabase() *DataLakeDatabase {
	this := DataLakeDatabase{}
	var maxWildcardCollections int32 = 100
	this.MaxWildcardCollections = &maxWildcardCollections
	return &this
}

// NewDataLakeDatabaseWithDefaults instantiates a new DataLakeDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeDatabaseWithDefaults() *DataLakeDatabase {
	this := DataLakeDatabase{}
	var maxWildcardCollections int32 = 100
	this.MaxWildcardCollections = &maxWildcardCollections
	return &this
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *DataLakeDatabase) GetCollections() []DataLakeDatabaseCollection {
	if o == nil || IsNil(o.Collections) {
		var ret []DataLakeDatabaseCollection
		return ret
	}
	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeDatabase) GetCollectionsOk() ([]DataLakeDatabaseCollection, bool) {
	if o == nil || IsNil(o.Collections) {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *DataLakeDatabase) HasCollections() bool {
	if o != nil && !IsNil(o.Collections) {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []DataLakeDatabaseCollection and assigns it to the Collections field.
func (o *DataLakeDatabase) SetCollections(v []DataLakeDatabaseCollection) {
	o.Collections = v
}

// GetMaxWildcardCollections returns the MaxWildcardCollections field value if set, zero value otherwise.
func (o *DataLakeDatabase) GetMaxWildcardCollections() int32 {
	if o == nil || IsNil(o.MaxWildcardCollections) {
		var ret int32
		return ret
	}
	return *o.MaxWildcardCollections
}

// GetMaxWildcardCollectionsOk returns a tuple with the MaxWildcardCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeDatabase) GetMaxWildcardCollectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxWildcardCollections) {
		return nil, false
	}
	return o.MaxWildcardCollections, true
}

// HasMaxWildcardCollections returns a boolean if a field has been set.
func (o *DataLakeDatabase) HasMaxWildcardCollections() bool {
	if o != nil && !IsNil(o.MaxWildcardCollections) {
		return true
	}

	return false
}

// SetMaxWildcardCollections gets a reference to the given int32 and assigns it to the MaxWildcardCollections field.
func (o *DataLakeDatabase) SetMaxWildcardCollections(v int32) {
	o.MaxWildcardCollections = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataLakeDatabase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeDatabase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataLakeDatabase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataLakeDatabase) SetName(v string) {
	o.Name = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *DataLakeDatabase) GetViews() []DataLakeView {
	if o == nil || IsNil(o.Views) {
		var ret []DataLakeView
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeDatabase) GetViewsOk() ([]DataLakeView, bool) {
	if o == nil || IsNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *DataLakeDatabase) HasViews() bool {
	if o != nil && !IsNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given []DataLakeView and assigns it to the Views field.
func (o *DataLakeDatabase) SetViews(v []DataLakeView) {
	o.Views = v
}

func (o DataLakeDatabase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataLakeDatabase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collections) {
		toSerialize["collections"] = o.Collections
	}
	if !IsNil(o.MaxWildcardCollections) {
		toSerialize["maxWildcardCollections"] = o.MaxWildcardCollections
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Views) {
		toSerialize["views"] = o.Views
	}
	return toSerialize, nil
}

type NullableDataLakeDatabase struct {
	value *DataLakeDatabase
	isSet bool
}

func (v NullableDataLakeDatabase) Get() *DataLakeDatabase {
	return v.value
}

func (v *NullableDataLakeDatabase) Set(val *DataLakeDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeDatabase(val *DataLakeDatabase) *NullableDataLakeDatabase {
	return &NullableDataLakeDatabase{value: val, isSet: true}
}

func (v NullableDataLakeDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


