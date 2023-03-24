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

// checks if the RoleMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMapping{}

// RoleMapping Mapping settings that link one IdP and MongoDB Cloud.
type RoleMapping struct {
	// Unique human-readable label that identifies the identity provider group to whichthis role mapping applies.
	ExternalGroupName string `json:"externalGroupName"`
	// Unique 24-hexadecimal digit string that identifies this role mapping.
	Id *string `json:"id,omitempty"`
	// Atlas roles and the unique identifiers of the groups and organizations associated with each role.
	RoleAssignments []RoleAssignment `json:"roleAssignments,omitempty"`
}

// NewRoleMapping instantiates a new RoleMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMapping(externalGroupName string) *RoleMapping {
	this := RoleMapping{}
	this.ExternalGroupName = externalGroupName
	return &this
}

// NewRoleMappingWithDefaults instantiates a new RoleMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMappingWithDefaults() *RoleMapping {
	this := RoleMapping{}
	return &this
}

// GetExternalGroupName returns the ExternalGroupName field value
func (o *RoleMapping) GetExternalGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalGroupName
}

// GetExternalGroupNameOk returns a tuple with the ExternalGroupName field value
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetExternalGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalGroupName, true
}

// SetExternalGroupName sets field value
func (o *RoleMapping) SetExternalGroupName(v string) {
	o.ExternalGroupName = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleMapping) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleMapping) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleMapping) SetId(v string) {
	o.Id = &v
}

// GetRoleAssignments returns the RoleAssignments field value if set, zero value otherwise.
func (o *RoleMapping) GetRoleAssignments() []RoleAssignment {
	if o == nil || IsNil(o.RoleAssignments) {
		var ret []RoleAssignment
		return ret
	}
	return o.RoleAssignments
}

// GetRoleAssignmentsOk returns a tuple with the RoleAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetRoleAssignmentsOk() ([]RoleAssignment, bool) {
	if o == nil || IsNil(o.RoleAssignments) {
		return nil, false
	}
	return o.RoleAssignments, true
}

// HasRoleAssignments returns a boolean if a field has been set.
func (o *RoleMapping) HasRoleAssignments() bool {
	if o != nil && !IsNil(o.RoleAssignments) {
		return true
	}

	return false
}

// SetRoleAssignments gets a reference to the given []RoleAssignment and assigns it to the RoleAssignments field.
func (o *RoleMapping) SetRoleAssignments(v []RoleAssignment) {
	o.RoleAssignments = v
}

func (o RoleMapping) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o RoleMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalGroupName"] = o.ExternalGroupName
	if !IsNil(o.RoleAssignments) {
		toSerialize["roleAssignments"] = o.RoleAssignments
	}
	return toSerialize, nil
}

type NullableRoleMapping struct {
	value *RoleMapping
	isSet bool
}

func (v NullableRoleMapping) Get() *RoleMapping {
	return v.value
}

func (v *NullableRoleMapping) Set(val *RoleMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMapping(val *RoleMapping) *NullableRoleMapping {
	return &NullableRoleMapping{value: val, isSet: true}
}

func (v NullableRoleMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


