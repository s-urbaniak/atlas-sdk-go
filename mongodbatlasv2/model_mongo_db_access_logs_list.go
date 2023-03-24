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

// checks if the MongoDBAccessLogsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MongoDBAccessLogsList{}

// MongoDBAccessLogsList struct for MongoDBAccessLogsList
type MongoDBAccessLogsList struct {
	// Authentication attempt, one per object, made against the cluster.
	AccessLogs []MongoDBAccessLogs `json:"accessLogs,omitempty"`
}

// NewMongoDBAccessLogsList instantiates a new MongoDBAccessLogsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongoDBAccessLogsList() *MongoDBAccessLogsList {
	this := MongoDBAccessLogsList{}
	return &this
}

// NewMongoDBAccessLogsListWithDefaults instantiates a new MongoDBAccessLogsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongoDBAccessLogsListWithDefaults() *MongoDBAccessLogsList {
	this := MongoDBAccessLogsList{}
	return &this
}

// GetAccessLogs returns the AccessLogs field value if set, zero value otherwise.
func (o *MongoDBAccessLogsList) GetAccessLogs() []MongoDBAccessLogs {
	if o == nil || IsNil(o.AccessLogs) {
		var ret []MongoDBAccessLogs
		return ret
	}
	return o.AccessLogs
}

// GetAccessLogsOk returns a tuple with the AccessLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogsList) GetAccessLogsOk() ([]MongoDBAccessLogs, bool) {
	if o == nil || IsNil(o.AccessLogs) {
		return nil, false
	}
	return o.AccessLogs, true
}

// HasAccessLogs returns a boolean if a field has been set.
func (o *MongoDBAccessLogsList) HasAccessLogs() bool {
	if o != nil && !IsNil(o.AccessLogs) {
		return true
	}

	return false
}

// SetAccessLogs gets a reference to the given []MongoDBAccessLogs and assigns it to the AccessLogs field.
func (o *MongoDBAccessLogsList) SetAccessLogs(v []MongoDBAccessLogs) {
	o.AccessLogs = v
}

func (o MongoDBAccessLogsList) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o MongoDBAccessLogsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableMongoDBAccessLogsList struct {
	value *MongoDBAccessLogsList
	isSet bool
}

func (v NullableMongoDBAccessLogsList) Get() *MongoDBAccessLogsList {
	return v.value
}

func (v *NullableMongoDBAccessLogsList) Set(val *MongoDBAccessLogsList) {
	v.value = val
	v.isSet = true
}

func (v NullableMongoDBAccessLogsList) IsSet() bool {
	return v.isSet
}

func (v *NullableMongoDBAccessLogsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMongoDBAccessLogsList(val *MongoDBAccessLogsList) *NullableMongoDBAccessLogsList {
	return &NullableMongoDBAccessLogsList{value: val, isSet: true}
}

func (v NullableMongoDBAccessLogsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMongoDBAccessLogsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


