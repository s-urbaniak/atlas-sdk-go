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

// NumberMetricUnits Element used to express the quantity. This can be an element of time, storage capacity, and the like.
type NumberMetricUnits string

// List of NumberMetricUnits
const (
	NUMBERMETRICUNITS_COUNT NumberMetricUnits = "COUNT"
	NUMBERMETRICUNITS_THOUSAND NumberMetricUnits = "THOUSAND"
	NUMBERMETRICUNITS_MILLION NumberMetricUnits = "MILLION"
	NUMBERMETRICUNITS_BILLION NumberMetricUnits = "BILLION"
)

// All allowed values of NumberMetricUnits enum
var AllowedNumberMetricUnitsEnumValues = []NumberMetricUnits{
	"COUNT",
	"THOUSAND",
	"MILLION",
	"BILLION",
}

func (v *NumberMetricUnits) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NumberMetricUnits(value)
	for _, existing := range AllowedNumberMetricUnitsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NumberMetricUnits", value)
}

// NewNumberMetricUnitsFromValue returns a pointer to a valid NumberMetricUnits
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNumberMetricUnitsFromValue(v string) (*NumberMetricUnits, error) {
	ev := NumberMetricUnits(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NumberMetricUnits: valid values are %v", v, AllowedNumberMetricUnitsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NumberMetricUnits) IsValid() bool {
	for _, existing := range AllowedNumberMetricUnitsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NumberMetricUnits value
func (v NumberMetricUnits) Ptr() *NumberMetricUnits {
	return &v
}


