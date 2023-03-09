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

// checks if the ConnectedOrgConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectedOrgConfig{}

// ConnectedOrgConfig struct for ConnectedOrgConfig
type ConnectedOrgConfig struct {
	// Approved domains that restrict users who can join the organization based on their email address.
	DomainAllowList []string `json:"domainAllowList,omitempty"`
	// Value that indicates whether domain restriction is enabled for this connected org.
	DomainRestrictionEnabled bool `json:"domainRestrictionEnabled"`
	// Unique 20-hexadecimal digit string that identifies the identity provider that this connected org config is associated with.
	IdentityProviderId string `json:"identityProviderId"`
	// Unique 24-hexadecimal digit string that identifies the connected organization configuration.
	OrgId string `json:"orgId"`
	// Atlas roles that are granted to a user in this organization after authenticating.
	PostAuthRoleGrants []string `json:"postAuthRoleGrants,omitempty"`
	// Role mappings that are configured in this organization.
	RoleMappings []RoleMapping `json:"roleMappings,omitempty"`
	// List that contains the users who have an email address that doesn't match any domain on the allowed list.
	UserConflicts []FederatedUser `json:"userConflicts,omitempty"`
}

// NewConnectedOrgConfig instantiates a new ConnectedOrgConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedOrgConfig(domainRestrictionEnabled bool, identityProviderId string, orgId string) *ConnectedOrgConfig {
	this := ConnectedOrgConfig{}
	this.DomainRestrictionEnabled = domainRestrictionEnabled
	this.IdentityProviderId = identityProviderId
	this.OrgId = orgId
	return &this
}

// NewConnectedOrgConfigWithDefaults instantiates a new ConnectedOrgConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedOrgConfigWithDefaults() *ConnectedOrgConfig {
	this := ConnectedOrgConfig{}
	return &this
}

// GetDomainAllowList returns the DomainAllowList field value if set, zero value otherwise.
func (o *ConnectedOrgConfig) GetDomainAllowList() []string {
	if o == nil || IsNil(o.DomainAllowList) {
		var ret []string
		return ret
	}
	return o.DomainAllowList
}

// GetDomainAllowListOk returns a tuple with the DomainAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedOrgConfig) GetDomainAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainAllowList) {
		return nil, false
	}
	return o.DomainAllowList, true
}

// HasDomainAllowList returns a boolean if a field has been set.
func (o *ConnectedOrgConfig) HasDomainAllowList() bool {
	if o != nil && !IsNil(o.DomainAllowList) {
		return true
	}

	return false
}

// SetDomainAllowList gets a reference to the given []string and assigns it to the DomainAllowList field.
func (o *ConnectedOrgConfig) SetDomainAllowList(v []string) {
	o.DomainAllowList = v
}

// GetDomainRestrictionEnabled returns the DomainRestrictionEnabled field value
func (o *ConnectedOrgConfig) GetDomainRestrictionEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DomainRestrictionEnabled
}

// GetDomainRestrictionEnabledOk returns a tuple with the DomainRestrictionEnabled field value
// and a boolean to check if the value has been set.
func (o *ConnectedOrgConfig) GetDomainRestrictionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainRestrictionEnabled, true
}

// SetDomainRestrictionEnabled sets field value
func (o *ConnectedOrgConfig) SetDomainRestrictionEnabled(v bool) {
	o.DomainRestrictionEnabled = v
}

// GetIdentityProviderId returns the IdentityProviderId field value
func (o *ConnectedOrgConfig) GetIdentityProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderId
}

// GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field value
// and a boolean to check if the value has been set.
func (o *ConnectedOrgConfig) GetIdentityProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderId, true
}

// SetIdentityProviderId sets field value
func (o *ConnectedOrgConfig) SetIdentityProviderId(v string) {
	o.IdentityProviderId = v
}

// GetOrgId returns the OrgId field value
func (o *ConnectedOrgConfig) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ConnectedOrgConfig) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ConnectedOrgConfig) SetOrgId(v string) {
	o.OrgId = v
}

// GetPostAuthRoleGrants returns the PostAuthRoleGrants field value if set, zero value otherwise.
func (o *ConnectedOrgConfig) GetPostAuthRoleGrants() []string {
	if o == nil || IsNil(o.PostAuthRoleGrants) {
		var ret []string
		return ret
	}
	return o.PostAuthRoleGrants
}

// GetPostAuthRoleGrantsOk returns a tuple with the PostAuthRoleGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedOrgConfig) GetPostAuthRoleGrantsOk() ([]string, bool) {
	if o == nil || IsNil(o.PostAuthRoleGrants) {
		return nil, false
	}
	return o.PostAuthRoleGrants, true
}

// HasPostAuthRoleGrants returns a boolean if a field has been set.
func (o *ConnectedOrgConfig) HasPostAuthRoleGrants() bool {
	if o != nil && !IsNil(o.PostAuthRoleGrants) {
		return true
	}

	return false
}

// SetPostAuthRoleGrants gets a reference to the given []string and assigns it to the PostAuthRoleGrants field.
func (o *ConnectedOrgConfig) SetPostAuthRoleGrants(v []string) {
	o.PostAuthRoleGrants = v
}

// GetRoleMappings returns the RoleMappings field value if set, zero value otherwise.
func (o *ConnectedOrgConfig) GetRoleMappings() []RoleMapping {
	if o == nil || IsNil(o.RoleMappings) {
		var ret []RoleMapping
		return ret
	}
	return o.RoleMappings
}

// GetRoleMappingsOk returns a tuple with the RoleMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedOrgConfig) GetRoleMappingsOk() ([]RoleMapping, bool) {
	if o == nil || IsNil(o.RoleMappings) {
		return nil, false
	}
	return o.RoleMappings, true
}

// HasRoleMappings returns a boolean if a field has been set.
func (o *ConnectedOrgConfig) HasRoleMappings() bool {
	if o != nil && !IsNil(o.RoleMappings) {
		return true
	}

	return false
}

// SetRoleMappings gets a reference to the given []RoleMapping and assigns it to the RoleMappings field.
func (o *ConnectedOrgConfig) SetRoleMappings(v []RoleMapping) {
	o.RoleMappings = v
}

// GetUserConflicts returns the UserConflicts field value if set, zero value otherwise.
func (o *ConnectedOrgConfig) GetUserConflicts() []FederatedUser {
	if o == nil || IsNil(o.UserConflicts) {
		var ret []FederatedUser
		return ret
	}
	return o.UserConflicts
}

// GetUserConflictsOk returns a tuple with the UserConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedOrgConfig) GetUserConflictsOk() ([]FederatedUser, bool) {
	if o == nil || IsNil(o.UserConflicts) {
		return nil, false
	}
	return o.UserConflicts, true
}

// HasUserConflicts returns a boolean if a field has been set.
func (o *ConnectedOrgConfig) HasUserConflicts() bool {
	if o != nil && !IsNil(o.UserConflicts) {
		return true
	}

	return false
}

// SetUserConflicts gets a reference to the given []FederatedUser and assigns it to the UserConflicts field.
func (o *ConnectedOrgConfig) SetUserConflicts(v []FederatedUser) {
	o.UserConflicts = v
}

func (o ConnectedOrgConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectedOrgConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DomainAllowList) {
		toSerialize["domainAllowList"] = o.DomainAllowList
	}
	toSerialize["domainRestrictionEnabled"] = o.DomainRestrictionEnabled
	toSerialize["identityProviderId"] = o.IdentityProviderId
	// skip: orgId is readOnly
	if !IsNil(o.PostAuthRoleGrants) {
		toSerialize["postAuthRoleGrants"] = o.PostAuthRoleGrants
	}
	if !IsNil(o.RoleMappings) {
		toSerialize["roleMappings"] = o.RoleMappings
	}
	if !IsNil(o.UserConflicts) {
		toSerialize["userConflicts"] = o.UserConflicts
	}
	return toSerialize, nil
}

type NullableConnectedOrgConfig struct {
	value *ConnectedOrgConfig
	isSet bool
}

func (v NullableConnectedOrgConfig) Get() *ConnectedOrgConfig {
	return v.value
}

func (v *NullableConnectedOrgConfig) Set(val *ConnectedOrgConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedOrgConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedOrgConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedOrgConfig(val *ConnectedOrgConfig) *NullableConnectedOrgConfig {
	return &NullableConnectedOrgConfig{value: val, isSet: true}
}

func (v NullableConnectedOrgConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedOrgConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


