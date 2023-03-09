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

// checks if the RestoreJobFileHashView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreJobFileHashView{}

// RestoreJobFileHashView Key and value pair that map one restore file to one hashed checksum. This parameter applies after you download the corresponding **delivery.url**.
type RestoreJobFileHashView struct {
	// Human-readable label that identifies the hashed file.
	FileName *string `json:"fileName,omitempty"`
	// Hashed checksum that maps to the restore file.
	Hash *string `json:"hash,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Human-readable label that identifies the hashing algorithm used to compute the hash value.
	TypeName *string `json:"typeName,omitempty"`
}

// NewRestoreJobFileHashView instantiates a new RestoreJobFileHashView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreJobFileHashView() *RestoreJobFileHashView {
	this := RestoreJobFileHashView{}
	return &this
}

// NewRestoreJobFileHashViewWithDefaults instantiates a new RestoreJobFileHashView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreJobFileHashViewWithDefaults() *RestoreJobFileHashView {
	this := RestoreJobFileHashView{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *RestoreJobFileHashView) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreJobFileHashView) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *RestoreJobFileHashView) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *RestoreJobFileHashView) SetFileName(v string) {
	o.FileName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *RestoreJobFileHashView) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreJobFileHashView) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *RestoreJobFileHashView) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *RestoreJobFileHashView) SetHash(v string) {
	o.Hash = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RestoreJobFileHashView) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreJobFileHashView) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RestoreJobFileHashView) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *RestoreJobFileHashView) SetLinks(v []Link) {
	o.Links = v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *RestoreJobFileHashView) GetTypeName() string {
	if o == nil || IsNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreJobFileHashView) GetTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.TypeName) {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *RestoreJobFileHashView) HasTypeName() bool {
	if o != nil && !IsNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *RestoreJobFileHashView) SetTypeName(v string) {
	o.TypeName = &v
}

func (o RestoreJobFileHashView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreJobFileHashView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: fileName is readOnly
	// skip: hash is readOnly
	// skip: links is readOnly
	// skip: typeName is readOnly
	return toSerialize, nil
}

type NullableRestoreJobFileHashView struct {
	value *RestoreJobFileHashView
	isSet bool
}

func (v NullableRestoreJobFileHashView) Get() *RestoreJobFileHashView {
	return v.value
}

func (v *NullableRestoreJobFileHashView) Set(val *RestoreJobFileHashView) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreJobFileHashView) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreJobFileHashView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreJobFileHashView(val *RestoreJobFileHashView) *NullableRestoreJobFileHashView {
	return &NullableRestoreJobFileHashView{value: val, isSet: true}
}

func (v NullableRestoreJobFileHashView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreJobFileHashView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


