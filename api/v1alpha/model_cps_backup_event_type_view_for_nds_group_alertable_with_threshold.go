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

// CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold Event type that triggers an alert.
type CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold string

// List of CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold
const (
	CPSBACKUPEVENTTYPEVIEWFORNDSGROUPALERTABLEWITHTHRESHOLD_CPS_SNAPSHOT_BEHIND CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold = "CPS_SNAPSHOT_BEHIND"
)

// All allowed values of CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold enum
var AllowedCpsBackupEventTypeViewForNdsGroupAlertableWithThresholdEnumValues = []CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold{
	"CPS_SNAPSHOT_BEHIND",
}

func (v *CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold(value)
	for _, existing := range AllowedCpsBackupEventTypeViewForNdsGroupAlertableWithThresholdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold", value)
}

// NewCpsBackupEventTypeViewForNdsGroupAlertableWithThresholdFromValue returns a pointer to a valid CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCpsBackupEventTypeViewForNdsGroupAlertableWithThresholdFromValue(v string) (*CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold, error) {
	ev := CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold: valid values are %v", v, AllowedCpsBackupEventTypeViewForNdsGroupAlertableWithThresholdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold) IsValid() bool {
	for _, existing := range AllowedCpsBackupEventTypeViewForNdsGroupAlertableWithThresholdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold value
func (v CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold) Ptr() *CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold {
	return &v
}


