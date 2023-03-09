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

// checks if the ApiUserView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUserView{}

// ApiUserView struct for ApiUserView
type ApiUserView struct {
	// Purpose or explanation provided when someone created this organization API key.
	Desc *string `json:"desc,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this organization API key assigned to this project.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Redacted private key returned for this organization API key. This key displays unredacted when first created.
	PrivateKey *string `json:"privateKey,omitempty"`
	// Public API key value set for the specified organization API key.
	PublicKey *string `json:"publicKey,omitempty"`
	// List that contains the roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the API key.
	Roles []RoleAssignmentView `json:"roles,omitempty"`
}

// NewApiUserView instantiates a new ApiUserView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserView() *ApiUserView {
	this := ApiUserView{}
	return &this
}

// NewApiUserViewWithDefaults instantiates a new ApiUserView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserViewWithDefaults() *ApiUserView {
	this := ApiUserView{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *ApiUserView) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserView) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *ApiUserView) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *ApiUserView) SetDesc(v string) {
	o.Desc = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiUserView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiUserView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiUserView) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiUserView) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserView) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiUserView) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ApiUserView) SetLinks(v []Link) {
	o.Links = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *ApiUserView) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserView) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *ApiUserView) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *ApiUserView) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *ApiUserView) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserView) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *ApiUserView) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *ApiUserView) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiUserView) GetRoles() []RoleAssignmentView {
	if o == nil || IsNil(o.Roles) {
		var ret []RoleAssignmentView
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserView) GetRolesOk() ([]RoleAssignmentView, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiUserView) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []RoleAssignmentView and assigns it to the Roles field.
func (o *ApiUserView) SetRoles(v []RoleAssignmentView) {
	o.Roles = v
}

func (o ApiUserView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	// skip: id is readOnly
	// skip: links is readOnly
	// skip: privateKey is readOnly
	// skip: publicKey is readOnly
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableApiUserView struct {
	value *ApiUserView
	isSet bool
}

func (v NullableApiUserView) Get() *ApiUserView {
	return v.value
}

func (v *NullableApiUserView) Set(val *ApiUserView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserView(val *ApiUserView) *NullableApiUserView {
	return &NullableApiUserView{value: val, isSet: true}
}

func (v NullableApiUserView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


