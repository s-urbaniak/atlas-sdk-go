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

// checks if the AvailableProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableProject{}

// AvailableProject struct for AvailableProject
type AvailableProject struct {
	// List of clusters that can be migrated to MongoDB Cloud.
	Deployments []AvailableDeployment `json:"deployments,omitempty"`
	// Hostname of MongoDB Agent list that you configured to perform a migration.
	MigrationHosts []string `json:"migrationHosts,omitempty"`
	// Human-readable label that identifies this project.
	Name string `json:"name"`
	// Unique 24-hexadecimal digit string that identifies the project to be migrated.
	ProjectId string `json:"projectId"`
}

// NewAvailableProject instantiates a new AvailableProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableProject(name string, projectId string) *AvailableProject {
	this := AvailableProject{}
	this.Name = name
	this.ProjectId = projectId
	return &this
}

// NewAvailableProjectWithDefaults instantiates a new AvailableProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableProjectWithDefaults() *AvailableProject {
	this := AvailableProject{}
	return &this
}

// GetDeployments returns the Deployments field value if set, zero value otherwise.
func (o *AvailableProject) GetDeployments() []AvailableDeployment {
	if o == nil || IsNil(o.Deployments) {
		var ret []AvailableDeployment
		return ret
	}
	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableProject) GetDeploymentsOk() ([]AvailableDeployment, bool) {
	if o == nil || IsNil(o.Deployments) {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *AvailableProject) HasDeployments() bool {
	if o != nil && !IsNil(o.Deployments) {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given []AvailableDeployment and assigns it to the Deployments field.
func (o *AvailableProject) SetDeployments(v []AvailableDeployment) {
	o.Deployments = v
}

// GetMigrationHosts returns the MigrationHosts field value if set, zero value otherwise.
func (o *AvailableProject) GetMigrationHosts() []string {
	if o == nil || IsNil(o.MigrationHosts) {
		var ret []string
		return ret
	}
	return o.MigrationHosts
}

// GetMigrationHostsOk returns a tuple with the MigrationHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableProject) GetMigrationHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.MigrationHosts) {
		return nil, false
	}
	return o.MigrationHosts, true
}

// HasMigrationHosts returns a boolean if a field has been set.
func (o *AvailableProject) HasMigrationHosts() bool {
	if o != nil && !IsNil(o.MigrationHosts) {
		return true
	}

	return false
}

// SetMigrationHosts gets a reference to the given []string and assigns it to the MigrationHosts field.
func (o *AvailableProject) SetMigrationHosts(v []string) {
	o.MigrationHosts = v
}

// GetName returns the Name field value
func (o *AvailableProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AvailableProject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AvailableProject) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *AvailableProject) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AvailableProject) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AvailableProject) SetProjectId(v string) {
	o.ProjectId = v
}

func (o AvailableProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailableProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deployments) {
		toSerialize["deployments"] = o.Deployments
	}
	if !IsNil(o.MigrationHosts) {
		toSerialize["migrationHosts"] = o.MigrationHosts
	}
	// skip: name is readOnly
	// skip: projectId is readOnly
	return toSerialize, nil
}

type NullableAvailableProject struct {
	value *AvailableProject
	isSet bool
}

func (v NullableAvailableProject) Get() *AvailableProject {
	return v.value
}

func (v *NullableAvailableProject) Set(val *AvailableProject) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableProject(val *AvailableProject) *NullableAvailableProject {
	return &NullableAvailableProject{value: val, isSet: true}
}

func (v NullableAvailableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


