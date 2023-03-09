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

// checks if the CreateOrganizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationResponse{}

// CreateOrganizationResponse struct for CreateOrganizationResponse
type CreateOrganizationResponse struct {
	ApiKey *User `json:"apiKey,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Atlas user that you want to assign the Organization Owner role.
	OrgOwnerId *string `json:"orgOwnerId,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
}

// NewCreateOrganizationResponse instantiates a new CreateOrganizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationResponse() *CreateOrganizationResponse {
	this := CreateOrganizationResponse{}
	return &this
}

// NewCreateOrganizationResponseWithDefaults instantiates a new CreateOrganizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationResponseWithDefaults() *CreateOrganizationResponse {
	this := CreateOrganizationResponse{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CreateOrganizationResponse) GetApiKey() User {
	if o == nil || IsNil(o.ApiKey) {
		var ret User
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationResponse) GetApiKeyOk() (*User, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateOrganizationResponse) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given User and assigns it to the ApiKey field.
func (o *CreateOrganizationResponse) SetApiKey(v User) {
	o.ApiKey = &v
}

// GetOrgOwnerId returns the OrgOwnerId field value if set, zero value otherwise.
func (o *CreateOrganizationResponse) GetOrgOwnerId() string {
	if o == nil || IsNil(o.OrgOwnerId) {
		var ret string
		return ret
	}
	return *o.OrgOwnerId
}

// GetOrgOwnerIdOk returns a tuple with the OrgOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationResponse) GetOrgOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgOwnerId) {
		return nil, false
	}
	return o.OrgOwnerId, true
}

// HasOrgOwnerId returns a boolean if a field has been set.
func (o *CreateOrganizationResponse) HasOrgOwnerId() bool {
	if o != nil && !IsNil(o.OrgOwnerId) {
		return true
	}

	return false
}

// SetOrgOwnerId gets a reference to the given string and assigns it to the OrgOwnerId field.
func (o *CreateOrganizationResponse) SetOrgOwnerId(v string) {
	o.OrgOwnerId = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CreateOrganizationResponse) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationResponse) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CreateOrganizationResponse) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *CreateOrganizationResponse) SetOrganization(v Organization) {
	o.Organization = &v
}

func (o CreateOrganizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	// skip: orgOwnerId is readOnly
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableCreateOrganizationResponse struct {
	value *CreateOrganizationResponse
	isSet bool
}

func (v NullableCreateOrganizationResponse) Get() *CreateOrganizationResponse {
	return v.value
}

func (v *NullableCreateOrganizationResponse) Set(val *CreateOrganizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationResponse(val *CreateOrganizationResponse) *NullableCreateOrganizationResponse {
	return &NullableCreateOrganizationResponse{value: val, isSet: true}
}

func (v NullableCreateOrganizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


