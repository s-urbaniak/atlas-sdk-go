/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"fmt"
)

// NDSServerlessInstanceAuditType Unique identifier of event type.
type NDSServerlessInstanceAuditType string

// List of NDSServerlessInstanceAuditType
const (
	NDSSERVERLESSINSTANCEAUDITTYPE_CREATED NDSServerlessInstanceAuditType = "SERVERLESS_INSTANCE_CREATED"
	NDSSERVERLESSINSTANCEAUDITTYPE_READY NDSServerlessInstanceAuditType = "SERVERLESS_INSTANCE_READY"
	NDSSERVERLESSINSTANCEAUDITTYPE_UPDATE_SUBMITTED NDSServerlessInstanceAuditType = "SERVERLESS_INSTANCE_UPDATE_SUBMITTED"
	NDSSERVERLESSINSTANCEAUDITTYPE_UPDATE_STARTED NDSServerlessInstanceAuditType = "SERVERLESS_INSTANCE_UPDATE_STARTED"
	NDSSERVERLESSINSTANCEAUDITTYPE_UPDATE_COMPLETED NDSServerlessInstanceAuditType = "SERVERLESS_INSTANCE_UPDATE_COMPLETED"
	NDSSERVERLESSINSTANCEAUDITTYPE_DELETE_SUBMITTED NDSServerlessInstanceAuditType = "SERVERLESS_INSTANCE_DELETE_SUBMITTED"
	NDSSERVERLESSINSTANCEAUDITTYPE_DELETED NDSServerlessInstanceAuditType = "SERVERLESS_INSTANCE_DELETED"
)

// All allowed values of NDSServerlessInstanceAuditType enum
var AllowedNDSServerlessInstanceAuditTypeEnumValues = []NDSServerlessInstanceAuditType{
	"SERVERLESS_INSTANCE_CREATED",
	"SERVERLESS_INSTANCE_READY",
	"SERVERLESS_INSTANCE_UPDATE_SUBMITTED",
	"SERVERLESS_INSTANCE_UPDATE_STARTED",
	"SERVERLESS_INSTANCE_UPDATE_COMPLETED",
	"SERVERLESS_INSTANCE_DELETE_SUBMITTED",
	"SERVERLESS_INSTANCE_DELETED",
}

func (v *NDSServerlessInstanceAuditType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NDSServerlessInstanceAuditType(value)
	for _, existing := range AllowedNDSServerlessInstanceAuditTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NDSServerlessInstanceAuditType", value)
}

// NewNDSServerlessInstanceAuditTypeFromValue returns a pointer to a valid NDSServerlessInstanceAuditType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNDSServerlessInstanceAuditTypeFromValue(v string) (*NDSServerlessInstanceAuditType, error) {
	ev := NDSServerlessInstanceAuditType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NDSServerlessInstanceAuditType: valid values are %v", v, AllowedNDSServerlessInstanceAuditTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NDSServerlessInstanceAuditType) IsValid() bool {
	for _, existing := range AllowedNDSServerlessInstanceAuditTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NDSServerlessInstanceAuditType value
func (v NDSServerlessInstanceAuditType) Ptr() *NDSServerlessInstanceAuditType {
	return &v
}

type NullableNDSServerlessInstanceAuditType struct {
	value *NDSServerlessInstanceAuditType
	isSet bool
}

func (v NullableNDSServerlessInstanceAuditType) Get() *NDSServerlessInstanceAuditType {
	return v.value
}

func (v *NullableNDSServerlessInstanceAuditType) Set(val *NDSServerlessInstanceAuditType) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSServerlessInstanceAuditType) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSServerlessInstanceAuditType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSServerlessInstanceAuditType(val *NDSServerlessInstanceAuditType) *NullableNDSServerlessInstanceAuditType {
	return &NullableNDSServerlessInstanceAuditType{value: val, isSet: true}
}

func (v NullableNDSServerlessInstanceAuditType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSServerlessInstanceAuditType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

