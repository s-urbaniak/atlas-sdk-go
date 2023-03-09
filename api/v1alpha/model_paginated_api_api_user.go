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

// checks if the PaginatedApiApiUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedApiApiUser{}

// PaginatedApiApiUser struct for PaginatedApiApiUser
type PaginatedApiApiUser struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ApiUser `json:"results,omitempty"`
	// Number of documents returned in this response.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewPaginatedApiApiUser instantiates a new PaginatedApiApiUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiApiUser() *PaginatedApiApiUser {
	this := PaginatedApiApiUser{}
	return &this
}

// NewPaginatedApiApiUserWithDefaults instantiates a new PaginatedApiApiUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiApiUserWithDefaults() *PaginatedApiApiUser {
	this := PaginatedApiApiUser{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginatedApiApiUser) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiApiUser) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedApiApiUser) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedApiApiUser) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedApiApiUser) GetResults() []ApiUser {
	if o == nil || IsNil(o.Results) {
		var ret []ApiUser
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiApiUser) GetResultsOk() ([]ApiUser, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedApiApiUser) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ApiUser and assigns it to the Results field.
func (o *PaginatedApiApiUser) SetResults(v []ApiUser) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedApiApiUser) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiApiUser) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedApiApiUser) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *PaginatedApiApiUser) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o PaginatedApiApiUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedApiApiUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: links is readOnly
	// skip: results is readOnly
	// skip: totalCount is readOnly
	return toSerialize, nil
}

type NullablePaginatedApiApiUser struct {
	value *PaginatedApiApiUser
	isSet bool
}

func (v NullablePaginatedApiApiUser) Get() *PaginatedApiApiUser {
	return v.value
}

func (v *NullablePaginatedApiApiUser) Set(val *PaginatedApiApiUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedApiApiUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedApiApiUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedApiApiUser(val *PaginatedApiApiUser) *NullablePaginatedApiApiUser {
	return &NullablePaginatedApiApiUser{value: val, isSet: true}
}

func (v NullablePaginatedApiApiUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedApiApiUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


