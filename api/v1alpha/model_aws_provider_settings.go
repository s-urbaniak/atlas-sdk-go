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

// AWSProviderSettings struct for AWSProviderSettings
type AWSProviderSettings struct {
	AutoScaling *AWSAutoScaling `json:"autoScaling,omitempty"`
	// Maximum Disk Input/Output Operations per Second (IOPS) that the database host can perform.
	DiskIOPS *int32 `json:"diskIOPS,omitempty"`
	// Flag that indicates whether the Amazon Elastic Block Store (EBS) encryption feature encrypts the host's root volume for both data at rest within the volume and for data moving between the volume and the cluster. Clusters always have this setting enabled.
	// Deprecated
	EncryptEBSVolume *bool `json:"encryptEBSVolume,omitempty"`
	// Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster.
	InstanceSizeName *string `json:"instanceSizeName,omitempty"`
	//  Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts.
	RegionName *string `json:"regionName,omitempty"`
	// Disk Input/Output Operations per Second (IOPS) setting for Amazon Web Services (AWS) storage that you configure only for abbr title=\"Amazon Web Services\">AWS</abbr>. Specify whether Disk Input/Output Operations per Second (IOPS) must not exceed the default Input/Output Operations per Second (IOPS) rate for the selected volume size (`STANDARD`), or must fall within the allowable Input/Output Operations per Second (IOPS) range for the selected volume size (`PROVISIONED`).
	VolumeType *string `json:"volumeType,omitempty"`
	ProviderName string `json:"providerName"`
}

// NewAWSProviderSettings instantiates a new AWSProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSProviderSettings() *AWSProviderSettings {
	this := AWSProviderSettings{}
	var encryptEBSVolume bool = true
	this.EncryptEBSVolume = &encryptEBSVolume
	return &this
}

// NewAWSProviderSettingsWithDefaults instantiates a new AWSProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSProviderSettingsWithDefaults() *AWSProviderSettings {
	this := AWSProviderSettings{}
	var encryptEBSVolume bool = true
	this.EncryptEBSVolume = &encryptEBSVolume
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetAutoScaling() AWSAutoScaling {
	if o == nil || o.AutoScaling == nil {
		var ret AWSAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetAutoScalingOk() (*AWSAutoScaling, bool) {
	if o == nil || o.AutoScaling == nil {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasAutoScaling() bool {
	if o != nil && o.AutoScaling != nil {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given AWSAutoScaling and assigns it to the AutoScaling field.
func (o *AWSProviderSettings) SetAutoScaling(v AWSAutoScaling) {
	o.AutoScaling = &v
}

// GetDiskIOPS returns the DiskIOPS field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetDiskIOPS() int32 {
	if o == nil || o.DiskIOPS == nil {
		var ret int32
		return ret
	}
	return *o.DiskIOPS
}

// GetDiskIOPSOk returns a tuple with the DiskIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetDiskIOPSOk() (*int32, bool) {
	if o == nil || o.DiskIOPS == nil {
		return nil, false
	}
	return o.DiskIOPS, true
}

// HasDiskIOPS returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasDiskIOPS() bool {
	if o != nil && o.DiskIOPS != nil {
		return true
	}

	return false
}

// SetDiskIOPS gets a reference to the given int32 and assigns it to the DiskIOPS field.
func (o *AWSProviderSettings) SetDiskIOPS(v int32) {
	o.DiskIOPS = &v
}

// GetEncryptEBSVolume returns the EncryptEBSVolume field value if set, zero value otherwise.
// Deprecated
func (o *AWSProviderSettings) GetEncryptEBSVolume() bool {
	if o == nil || o.EncryptEBSVolume == nil {
		var ret bool
		return ret
	}
	return *o.EncryptEBSVolume
}

// GetEncryptEBSVolumeOk returns a tuple with the EncryptEBSVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AWSProviderSettings) GetEncryptEBSVolumeOk() (*bool, bool) {
	if o == nil || o.EncryptEBSVolume == nil {
		return nil, false
	}
	return o.EncryptEBSVolume, true
}

// HasEncryptEBSVolume returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasEncryptEBSVolume() bool {
	if o != nil && o.EncryptEBSVolume != nil {
		return true
	}

	return false
}

// SetEncryptEBSVolume gets a reference to the given bool and assigns it to the EncryptEBSVolume field.
// Deprecated
func (o *AWSProviderSettings) SetEncryptEBSVolume(v bool) {
	o.EncryptEBSVolume = &v
}

// GetInstanceSizeName returns the InstanceSizeName field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetInstanceSizeName() string {
	if o == nil || o.InstanceSizeName == nil {
		var ret string
		return ret
	}
	return *o.InstanceSizeName
}

// GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetInstanceSizeNameOk() (*string, bool) {
	if o == nil || o.InstanceSizeName == nil {
		return nil, false
	}
	return o.InstanceSizeName, true
}

// HasInstanceSizeName returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasInstanceSizeName() bool {
	if o != nil && o.InstanceSizeName != nil {
		return true
	}

	return false
}

// SetInstanceSizeName gets a reference to the given string and assigns it to the InstanceSizeName field.
func (o *AWSProviderSettings) SetInstanceSizeName(v string) {
	o.InstanceSizeName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AWSProviderSettings) SetRegionName(v string) {
	o.RegionName = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *AWSProviderSettings) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetProviderName returns the ProviderName field value
func (o *AWSProviderSettings) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AWSProviderSettings) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AWSProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoScaling != nil {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if o.DiskIOPS != nil {
		toSerialize["diskIOPS"] = o.DiskIOPS
	}
	if o.EncryptEBSVolume != nil {
		toSerialize["encryptEBSVolume"] = o.EncryptEBSVolume
	}
	if o.InstanceSizeName != nil {
		toSerialize["instanceSizeName"] = o.InstanceSizeName
	}
	if o.RegionName != nil {
		toSerialize["regionName"] = o.RegionName
	}
	if o.VolumeType != nil {
		toSerialize["volumeType"] = o.VolumeType
	}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	return json.Marshal(toSerialize)
}

type NullableAWSProviderSettings struct {
	value *AWSProviderSettings
	isSet bool
}

func (v NullableAWSProviderSettings) Get() *AWSProviderSettings {
	return v.value
}

func (v *NullableAWSProviderSettings) Set(val *AWSProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSProviderSettings(val *AWSProviderSettings) *NullableAWSProviderSettings {
	return &NullableAWSProviderSettings{value: val, isSet: true}
}

func (v NullableAWSProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


