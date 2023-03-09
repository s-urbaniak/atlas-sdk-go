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

// checks if the DiskBackupRestoreJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskBackupRestoreJob{}

// DiskBackupRestoreJob struct for DiskBackupRestoreJob
type DiskBackupRestoreJob struct {
	// Flag that indicates whether someone canceled this restore job.
	Cancelled *bool `json:"cancelled,omitempty"`
	// Information on the restore job for each replica set in the sharded cluster.
	Components []DiskBackupBaseRestoreMember `json:"components,omitempty"`
	// Human-readable label that categorizes the restore job to create.
	DeliveryType string `json:"deliveryType"`
	// One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when `\"deliveryType\" : \"download\"`.
	DeliveryUrl []string `json:"deliveryUrl,omitempty"`
	DesiredTimestamp *BSONTimestamp `json:"desiredTimestamp,omitempty"`
	// Flag that indicates whether the restore job expired.
	Expired *bool `json:"expired,omitempty"`
	// Date and time when the restore job expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Flag that indicates whether the restore job failed.
	Failed *bool `json:"failed,omitempty"`
	// Date and time when the restore job completed. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// Unique 24-hexadecimal character string that identifies the restore job.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Oplog operation number from which you want to restore this snapshot. This number represents the second part of an Oplog timestamp. The resource returns this parameter when `\"deliveryType\" : \"pointInTime\"` and **oplogTs** exceeds `0`.
	OplogInc *int32 `json:"oplogInc,omitempty"`
	// Date and time from which you want to restore this snapshot. This parameter expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. This number represents the first part of an Oplog timestamp. The resource returns this parameter when `\"deliveryType\" : \"pointInTime\"` and **oplogTs** exceeds `0`.
	OplogTs *int32 `json:"oplogTs,omitempty"`
	// Date and time from which MongoDB Cloud restored this snapshot. This parameter expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. The resource returns this parameter when `\"deliveryType\" : \"pointInTime\"` and **pointInTimeUTCSeconds** exceeds `0`.
	PointInTimeUTCSeconds *int32 `json:"pointInTimeUTCSeconds,omitempty"`
	// Unique 24-hexadecimal character string that identifies the snapshot.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// Human-readable label that identifies the target cluster to which the restore job restores the snapshot. The resource returns this parameter when `\"deliveryType\":` `\"automated\"`.
	TargetClusterName string `json:"targetClusterName"`
	// Unique 24-hexadecimal digit string that identifies the target project for the specified **targetClusterName**.
	TargetGroupId string `json:"targetGroupId"`
	// Date and time when MongoDB Cloud took the snapshot associated with **snapshotId**. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewDiskBackupRestoreJob instantiates a new DiskBackupRestoreJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupRestoreJob(deliveryType string, targetClusterName string, targetGroupId string) *DiskBackupRestoreJob {
	this := DiskBackupRestoreJob{}
	this.DeliveryType = deliveryType
	this.TargetClusterName = targetClusterName
	this.TargetGroupId = targetGroupId
	return &this
}

// NewDiskBackupRestoreJobWithDefaults instantiates a new DiskBackupRestoreJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupRestoreJobWithDefaults() *DiskBackupRestoreJob {
	this := DiskBackupRestoreJob{}
	return &this
}

// GetCancelled returns the Cancelled field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetCancelled() bool {
	if o == nil || IsNil(o.Cancelled) {
		var ret bool
		return ret
	}
	return *o.Cancelled
}

// GetCancelledOk returns a tuple with the Cancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.Cancelled) {
		return nil, false
	}
	return o.Cancelled, true
}

// HasCancelled returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasCancelled() bool {
	if o != nil && !IsNil(o.Cancelled) {
		return true
	}

	return false
}

// SetCancelled gets a reference to the given bool and assigns it to the Cancelled field.
func (o *DiskBackupRestoreJob) SetCancelled(v bool) {
	o.Cancelled = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetComponents() []DiskBackupBaseRestoreMember {
	if o == nil || IsNil(o.Components) {
		var ret []DiskBackupBaseRestoreMember
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetComponentsOk() ([]DiskBackupBaseRestoreMember, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []DiskBackupBaseRestoreMember and assigns it to the Components field.
func (o *DiskBackupRestoreJob) SetComponents(v []DiskBackupBaseRestoreMember) {
	o.Components = v
}

// GetDeliveryType returns the DeliveryType field value
func (o *DiskBackupRestoreJob) GetDeliveryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryType
}

// GetDeliveryTypeOk returns a tuple with the DeliveryType field value
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetDeliveryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryType, true
}

// SetDeliveryType sets field value
func (o *DiskBackupRestoreJob) SetDeliveryType(v string) {
	o.DeliveryType = v
}

// GetDeliveryUrl returns the DeliveryUrl field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetDeliveryUrl() []string {
	if o == nil || IsNil(o.DeliveryUrl) {
		var ret []string
		return ret
	}
	return o.DeliveryUrl
}

// GetDeliveryUrlOk returns a tuple with the DeliveryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetDeliveryUrlOk() ([]string, bool) {
	if o == nil || IsNil(o.DeliveryUrl) {
		return nil, false
	}
	return o.DeliveryUrl, true
}

// HasDeliveryUrl returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasDeliveryUrl() bool {
	if o != nil && !IsNil(o.DeliveryUrl) {
		return true
	}

	return false
}

// SetDeliveryUrl gets a reference to the given []string and assigns it to the DeliveryUrl field.
func (o *DiskBackupRestoreJob) SetDeliveryUrl(v []string) {
	o.DeliveryUrl = v
}

// GetDesiredTimestamp returns the DesiredTimestamp field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetDesiredTimestamp() BSONTimestamp {
	if o == nil || IsNil(o.DesiredTimestamp) {
		var ret BSONTimestamp
		return ret
	}
	return *o.DesiredTimestamp
}

// GetDesiredTimestampOk returns a tuple with the DesiredTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetDesiredTimestampOk() (*BSONTimestamp, bool) {
	if o == nil || IsNil(o.DesiredTimestamp) {
		return nil, false
	}
	return o.DesiredTimestamp, true
}

// HasDesiredTimestamp returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasDesiredTimestamp() bool {
	if o != nil && !IsNil(o.DesiredTimestamp) {
		return true
	}

	return false
}

// SetDesiredTimestamp gets a reference to the given BSONTimestamp and assigns it to the DesiredTimestamp field.
func (o *DiskBackupRestoreJob) SetDesiredTimestamp(v BSONTimestamp) {
	o.DesiredTimestamp = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetExpired() bool {
	if o == nil || IsNil(o.Expired) {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Expired) {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasExpired() bool {
	if o != nil && !IsNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *DiskBackupRestoreJob) SetExpired(v bool) {
	o.Expired = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *DiskBackupRestoreJob) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *DiskBackupRestoreJob) SetFailed(v bool) {
	o.Failed = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DiskBackupRestoreJob) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiskBackupRestoreJob) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupRestoreJob) SetLinks(v []Link) {
	o.Links = v
}

// GetOplogInc returns the OplogInc field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetOplogInc() int32 {
	if o == nil || IsNil(o.OplogInc) {
		var ret int32
		return ret
	}
	return *o.OplogInc
}

// GetOplogIncOk returns a tuple with the OplogInc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetOplogIncOk() (*int32, bool) {
	if o == nil || IsNil(o.OplogInc) {
		return nil, false
	}
	return o.OplogInc, true
}

// HasOplogInc returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasOplogInc() bool {
	if o != nil && !IsNil(o.OplogInc) {
		return true
	}

	return false
}

// SetOplogInc gets a reference to the given int32 and assigns it to the OplogInc field.
func (o *DiskBackupRestoreJob) SetOplogInc(v int32) {
	o.OplogInc = &v
}

// GetOplogTs returns the OplogTs field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetOplogTs() int32 {
	if o == nil || IsNil(o.OplogTs) {
		var ret int32
		return ret
	}
	return *o.OplogTs
}

// GetOplogTsOk returns a tuple with the OplogTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetOplogTsOk() (*int32, bool) {
	if o == nil || IsNil(o.OplogTs) {
		return nil, false
	}
	return o.OplogTs, true
}

// HasOplogTs returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasOplogTs() bool {
	if o != nil && !IsNil(o.OplogTs) {
		return true
	}

	return false
}

// SetOplogTs gets a reference to the given int32 and assigns it to the OplogTs field.
func (o *DiskBackupRestoreJob) SetOplogTs(v int32) {
	o.OplogTs = &v
}

// GetPointInTimeUTCSeconds returns the PointInTimeUTCSeconds field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetPointInTimeUTCSeconds() int32 {
	if o == nil || IsNil(o.PointInTimeUTCSeconds) {
		var ret int32
		return ret
	}
	return *o.PointInTimeUTCSeconds
}

// GetPointInTimeUTCSecondsOk returns a tuple with the PointInTimeUTCSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetPointInTimeUTCSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.PointInTimeUTCSeconds) {
		return nil, false
	}
	return o.PointInTimeUTCSeconds, true
}

// HasPointInTimeUTCSeconds returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasPointInTimeUTCSeconds() bool {
	if o != nil && !IsNil(o.PointInTimeUTCSeconds) {
		return true
	}

	return false
}

// SetPointInTimeUTCSeconds gets a reference to the given int32 and assigns it to the PointInTimeUTCSeconds field.
func (o *DiskBackupRestoreJob) SetPointInTimeUTCSeconds(v int32) {
	o.PointInTimeUTCSeconds = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *DiskBackupRestoreJob) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetTargetClusterName returns the TargetClusterName field value
func (o *DiskBackupRestoreJob) GetTargetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetClusterName
}

// GetTargetClusterNameOk returns a tuple with the TargetClusterName field value
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetTargetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetClusterName, true
}

// SetTargetClusterName sets field value
func (o *DiskBackupRestoreJob) SetTargetClusterName(v string) {
	o.TargetClusterName = v
}

// GetTargetGroupId returns the TargetGroupId field value
func (o *DiskBackupRestoreJob) GetTargetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetTargetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetGroupId, true
}

// SetTargetGroupId sets field value
func (o *DiskBackupRestoreJob) SetTargetGroupId(v string) {
	o.TargetGroupId = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DiskBackupRestoreJob) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupRestoreJob) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DiskBackupRestoreJob) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *DiskBackupRestoreJob) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o DiskBackupRestoreJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskBackupRestoreJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: cancelled is readOnly
	// skip: components is readOnly
	toSerialize["deliveryType"] = o.DeliveryType
	// skip: deliveryUrl is readOnly
	if !IsNil(o.DesiredTimestamp) {
		toSerialize["desiredTimestamp"] = o.DesiredTimestamp
	}
	// skip: expired is readOnly
	// skip: expiresAt is readOnly
	// skip: failed is readOnly
	// skip: finishedAt is readOnly
	// skip: id is readOnly
	// skip: links is readOnly
	if !IsNil(o.OplogInc) {
		toSerialize["oplogInc"] = o.OplogInc
	}
	if !IsNil(o.OplogTs) {
		toSerialize["oplogTs"] = o.OplogTs
	}
	if !IsNil(o.PointInTimeUTCSeconds) {
		toSerialize["pointInTimeUTCSeconds"] = o.PointInTimeUTCSeconds
	}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	toSerialize["targetClusterName"] = o.TargetClusterName
	toSerialize["targetGroupId"] = o.TargetGroupId
	// skip: timestamp is readOnly
	return toSerialize, nil
}

type NullableDiskBackupRestoreJob struct {
	value *DiskBackupRestoreJob
	isSet bool
}

func (v NullableDiskBackupRestoreJob) Get() *DiskBackupRestoreJob {
	return v.value
}

func (v *NullableDiskBackupRestoreJob) Set(val *DiskBackupRestoreJob) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskBackupRestoreJob) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskBackupRestoreJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskBackupRestoreJob(val *DiskBackupRestoreJob) *NullableDiskBackupRestoreJob {
	return &NullableDiskBackupRestoreJob{value: val, isSet: true}
}

func (v NullableDiskBackupRestoreJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskBackupRestoreJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


