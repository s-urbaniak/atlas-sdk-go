/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"time"
)

// checks if the SnapshotView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotView{}

// SnapshotView struct for SnapshotView
type SnapshotView struct {
	// Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return.
	ClusterId *string `json:"clusterId,omitempty"`
	// Flag that indicates whether the snapshot exists. This flag returns `false` while MongoDB Cloud creates the snapshot.
	Complete *bool `json:"complete,omitempty"`
	Created *BSONTimestampView `json:"created,omitempty"`
	// Flag that indicates whether someone can delete this snapshot. You can't set `\"doNotDelete\" : true` and set a timestamp for **expires** in the same request.
	DoNotDelete *bool `json:"doNotDelete,omitempty"`
	// Date and time when MongoDB Cloud deletes the snapshot. If `\"doNotDelete\" : true`, MongoDB Cloud removes any value set for this parameter.
	Expires *time.Time `json:"expires,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project that owns the snapshots.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the snapshot.
	Id *string `json:"id,omitempty"`
	LastOplogAppliedTimestamp *BSONTimestampView `json:"lastOplogAppliedTimestamp,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Metadata that describes the complete snapshot.  - For a replica set, this array contains a single document. - For a sharded cluster, this array contains one document for each shard plus one document for the config host.
	Parts []SnapshotPartView `json:"parts,omitempty"`
}

// NewSnapshotView instantiates a new SnapshotView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotView() *SnapshotView {
	this := SnapshotView{}
	return &this
}

// NewSnapshotViewWithDefaults instantiates a new SnapshotView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotViewWithDefaults() *SnapshotView {
	this := SnapshotView{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *SnapshotView) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *SnapshotView) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *SnapshotView) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *SnapshotView) GetComplete() bool {
	if o == nil || IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *SnapshotView) HasComplete() bool {
	if o != nil && !IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *SnapshotView) SetComplete(v bool) {
	o.Complete = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SnapshotView) GetCreated() BSONTimestampView {
	if o == nil || IsNil(o.Created) {
		var ret BSONTimestampView
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetCreatedOk() (*BSONTimestampView, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SnapshotView) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given BSONTimestampView and assigns it to the Created field.
func (o *SnapshotView) SetCreated(v BSONTimestampView) {
	o.Created = &v
}

// GetDoNotDelete returns the DoNotDelete field value if set, zero value otherwise.
func (o *SnapshotView) GetDoNotDelete() bool {
	if o == nil || IsNil(o.DoNotDelete) {
		var ret bool
		return ret
	}
	return *o.DoNotDelete
}

// GetDoNotDeleteOk returns a tuple with the DoNotDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetDoNotDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.DoNotDelete) {
		return nil, false
	}
	return o.DoNotDelete, true
}

// HasDoNotDelete returns a boolean if a field has been set.
func (o *SnapshotView) HasDoNotDelete() bool {
	if o != nil && !IsNil(o.DoNotDelete) {
		return true
	}

	return false
}

// SetDoNotDelete gets a reference to the given bool and assigns it to the DoNotDelete field.
func (o *SnapshotView) SetDoNotDelete(v bool) {
	o.DoNotDelete = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *SnapshotView) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *SnapshotView) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *SnapshotView) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SnapshotView) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SnapshotView) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SnapshotView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SnapshotView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SnapshotView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SnapshotView) SetId(v string) {
	o.Id = &v
}

// GetLastOplogAppliedTimestamp returns the LastOplogAppliedTimestamp field value if set, zero value otherwise.
func (o *SnapshotView) GetLastOplogAppliedTimestamp() BSONTimestampView {
	if o == nil || IsNil(o.LastOplogAppliedTimestamp) {
		var ret BSONTimestampView
		return ret
	}
	return *o.LastOplogAppliedTimestamp
}

// GetLastOplogAppliedTimestampOk returns a tuple with the LastOplogAppliedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetLastOplogAppliedTimestampOk() (*BSONTimestampView, bool) {
	if o == nil || IsNil(o.LastOplogAppliedTimestamp) {
		return nil, false
	}
	return o.LastOplogAppliedTimestamp, true
}

// HasLastOplogAppliedTimestamp returns a boolean if a field has been set.
func (o *SnapshotView) HasLastOplogAppliedTimestamp() bool {
	if o != nil && !IsNil(o.LastOplogAppliedTimestamp) {
		return true
	}

	return false
}

// SetLastOplogAppliedTimestamp gets a reference to the given BSONTimestampView and assigns it to the LastOplogAppliedTimestamp field.
func (o *SnapshotView) SetLastOplogAppliedTimestamp(v BSONTimestampView) {
	o.LastOplogAppliedTimestamp = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SnapshotView) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SnapshotView) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *SnapshotView) SetLinks(v []Link) {
	o.Links = v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *SnapshotView) GetParts() []SnapshotPartView {
	if o == nil || IsNil(o.Parts) {
		var ret []SnapshotPartView
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotView) GetPartsOk() ([]SnapshotPartView, bool) {
	if o == nil || IsNil(o.Parts) {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *SnapshotView) HasParts() bool {
	if o != nil && !IsNil(o.Parts) {
		return true
	}

	return false
}

// SetParts gets a reference to the given []SnapshotPartView and assigns it to the Parts field.
func (o *SnapshotView) SetParts(v []SnapshotPartView) {
	o.Parts = v
}

func (o SnapshotView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: clusterId is readOnly
	// skip: complete is readOnly
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.DoNotDelete) {
		toSerialize["doNotDelete"] = o.DoNotDelete
	}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	// skip: groupId is readOnly
	// skip: id is readOnly
	if !IsNil(o.LastOplogAppliedTimestamp) {
		toSerialize["lastOplogAppliedTimestamp"] = o.LastOplogAppliedTimestamp
	}
	// skip: links is readOnly
	// skip: parts is readOnly
	return toSerialize, nil
}

type NullableSnapshotView struct {
	value *SnapshotView
	isSet bool
}

func (v NullableSnapshotView) Get() *SnapshotView {
	return v.value
}

func (v *NullableSnapshotView) Set(val *SnapshotView) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotView) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotView(val *SnapshotView) *NullableSnapshotView {
	return &NullableSnapshotView{value: val, isSet: true}
}

func (v NullableSnapshotView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


