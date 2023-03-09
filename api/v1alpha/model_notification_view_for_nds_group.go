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

// NotificationViewForNdsGroup - One target that MongoDB Cloud sends notifications when an alert triggers.
type NotificationViewForNdsGroup struct {
	DatadogNotification *DatadogNotification
	EmailNotification *EmailNotification
	GroupNotification *GroupNotification
	HipChatNotification *HipChatNotification
	MicrosoftTeamsNotification *MicrosoftTeamsNotification
	OpsGenieNotification *OpsGenieNotification
	OrgNotification *OrgNotification
	PagerDutyNotification *PagerDutyNotification
	SMSNotification *SMSNotification
	SlackNotification *SlackNotification
	TeamNotification *TeamNotification
	UserNotification *UserNotification
	VictorOpsNotification *VictorOpsNotification
	WebhookNotification *WebhookNotification
}

// DatadogNotificationAsNotificationViewForNdsGroup is a convenience function that returns DatadogNotification wrapped in NotificationViewForNdsGroup
func DatadogNotificationAsNotificationViewForNdsGroup(v *DatadogNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		DatadogNotification: v,
	}
}

// EmailNotificationAsNotificationViewForNdsGroup is a convenience function that returns EmailNotification wrapped in NotificationViewForNdsGroup
func EmailNotificationAsNotificationViewForNdsGroup(v *EmailNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		EmailNotification: v,
	}
}

// GroupNotificationAsNotificationViewForNdsGroup is a convenience function that returns GroupNotification wrapped in NotificationViewForNdsGroup
func GroupNotificationAsNotificationViewForNdsGroup(v *GroupNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		GroupNotification: v,
	}
}

// HipChatNotificationAsNotificationViewForNdsGroup is a convenience function that returns HipChatNotification wrapped in NotificationViewForNdsGroup
func HipChatNotificationAsNotificationViewForNdsGroup(v *HipChatNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		HipChatNotification: v,
	}
}

// MicrosoftTeamsNotificationAsNotificationViewForNdsGroup is a convenience function that returns MicrosoftTeamsNotification wrapped in NotificationViewForNdsGroup
func MicrosoftTeamsNotificationAsNotificationViewForNdsGroup(v *MicrosoftTeamsNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		MicrosoftTeamsNotification: v,
	}
}

// OpsGenieNotificationAsNotificationViewForNdsGroup is a convenience function that returns OpsGenieNotification wrapped in NotificationViewForNdsGroup
func OpsGenieNotificationAsNotificationViewForNdsGroup(v *OpsGenieNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		OpsGenieNotification: v,
	}
}

// OrgNotificationAsNotificationViewForNdsGroup is a convenience function that returns OrgNotification wrapped in NotificationViewForNdsGroup
func OrgNotificationAsNotificationViewForNdsGroup(v *OrgNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		OrgNotification: v,
	}
}

// PagerDutyNotificationAsNotificationViewForNdsGroup is a convenience function that returns PagerDutyNotification wrapped in NotificationViewForNdsGroup
func PagerDutyNotificationAsNotificationViewForNdsGroup(v *PagerDutyNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		PagerDutyNotification: v,
	}
}

// SMSNotificationAsNotificationViewForNdsGroup is a convenience function that returns SMSNotification wrapped in NotificationViewForNdsGroup
func SMSNotificationAsNotificationViewForNdsGroup(v *SMSNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		SMSNotification: v,
	}
}

// SlackNotificationAsNotificationViewForNdsGroup is a convenience function that returns SlackNotification wrapped in NotificationViewForNdsGroup
func SlackNotificationAsNotificationViewForNdsGroup(v *SlackNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		SlackNotification: v,
	}
}

// TeamNotificationAsNotificationViewForNdsGroup is a convenience function that returns TeamNotification wrapped in NotificationViewForNdsGroup
func TeamNotificationAsNotificationViewForNdsGroup(v *TeamNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		TeamNotification: v,
	}
}

// UserNotificationAsNotificationViewForNdsGroup is a convenience function that returns UserNotification wrapped in NotificationViewForNdsGroup
func UserNotificationAsNotificationViewForNdsGroup(v *UserNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		UserNotification: v,
	}
}

// VictorOpsNotificationAsNotificationViewForNdsGroup is a convenience function that returns VictorOpsNotification wrapped in NotificationViewForNdsGroup
func VictorOpsNotificationAsNotificationViewForNdsGroup(v *VictorOpsNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		VictorOpsNotification: v,
	}
}

// WebhookNotificationAsNotificationViewForNdsGroup is a convenience function that returns WebhookNotification wrapped in NotificationViewForNdsGroup
func WebhookNotificationAsNotificationViewForNdsGroup(v *WebhookNotification) NotificationViewForNdsGroup {
	return NotificationViewForNdsGroup{
		WebhookNotification: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotificationViewForNdsGroup) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatadogNotification
	err = json.Unmarshal(data, &dst.DatadogNotification)
	if err == nil {
		jsonDatadogNotification, _ := json.Marshal(dst.DatadogNotification)
		if string(jsonDatadogNotification) == "{}" { // empty struct
			dst.DatadogNotification = nil
		} else {
			match++
		}
	} else {
		dst.DatadogNotification = nil
	}

	// try to unmarshal data into EmailNotification
	err = json.Unmarshal(data, &dst.EmailNotification)
	if err == nil {
		jsonEmailNotification, _ := json.Marshal(dst.EmailNotification)
		if string(jsonEmailNotification) == "{}" { // empty struct
			dst.EmailNotification = nil
		} else {
			match++
		}
	} else {
		dst.EmailNotification = nil
	}

	// try to unmarshal data into GroupNotification
	err = json.Unmarshal(data, &dst.GroupNotification)
	if err == nil {
		jsonGroupNotification, _ := json.Marshal(dst.GroupNotification)
		if string(jsonGroupNotification) == "{}" { // empty struct
			dst.GroupNotification = nil
		} else {
			match++
		}
	} else {
		dst.GroupNotification = nil
	}

	// try to unmarshal data into HipChatNotification
	err = json.Unmarshal(data, &dst.HipChatNotification)
	if err == nil {
		jsonHipChatNotification, _ := json.Marshal(dst.HipChatNotification)
		if string(jsonHipChatNotification) == "{}" { // empty struct
			dst.HipChatNotification = nil
		} else {
			match++
		}
	} else {
		dst.HipChatNotification = nil
	}

	// try to unmarshal data into MicrosoftTeamsNotification
	err = json.Unmarshal(data, &dst.MicrosoftTeamsNotification)
	if err == nil {
		jsonMicrosoftTeamsNotification, _ := json.Marshal(dst.MicrosoftTeamsNotification)
		if string(jsonMicrosoftTeamsNotification) == "{}" { // empty struct
			dst.MicrosoftTeamsNotification = nil
		} else {
			match++
		}
	} else {
		dst.MicrosoftTeamsNotification = nil
	}

	// try to unmarshal data into OpsGenieNotification
	err = json.Unmarshal(data, &dst.OpsGenieNotification)
	if err == nil {
		jsonOpsGenieNotification, _ := json.Marshal(dst.OpsGenieNotification)
		if string(jsonOpsGenieNotification) == "{}" { // empty struct
			dst.OpsGenieNotification = nil
		} else {
			match++
		}
	} else {
		dst.OpsGenieNotification = nil
	}

	// try to unmarshal data into OrgNotification
	err = json.Unmarshal(data, &dst.OrgNotification)
	if err == nil {
		jsonOrgNotification, _ := json.Marshal(dst.OrgNotification)
		if string(jsonOrgNotification) == "{}" { // empty struct
			dst.OrgNotification = nil
		} else {
			match++
		}
	} else {
		dst.OrgNotification = nil
	}

	// try to unmarshal data into PagerDutyNotification
	err = json.Unmarshal(data, &dst.PagerDutyNotification)
	if err == nil {
		jsonPagerDutyNotification, _ := json.Marshal(dst.PagerDutyNotification)
		if string(jsonPagerDutyNotification) == "{}" { // empty struct
			dst.PagerDutyNotification = nil
		} else {
			match++
		}
	} else {
		dst.PagerDutyNotification = nil
	}

	// try to unmarshal data into SMSNotification
	err = json.Unmarshal(data, &dst.SMSNotification)
	if err == nil {
		jsonSMSNotification, _ := json.Marshal(dst.SMSNotification)
		if string(jsonSMSNotification) == "{}" { // empty struct
			dst.SMSNotification = nil
		} else {
			match++
		}
	} else {
		dst.SMSNotification = nil
	}

	// try to unmarshal data into SlackNotification
	err = json.Unmarshal(data, &dst.SlackNotification)
	if err == nil {
		jsonSlackNotification, _ := json.Marshal(dst.SlackNotification)
		if string(jsonSlackNotification) == "{}" { // empty struct
			dst.SlackNotification = nil
		} else {
			match++
		}
	} else {
		dst.SlackNotification = nil
	}

	// try to unmarshal data into TeamNotification
	err = json.Unmarshal(data, &dst.TeamNotification)
	if err == nil {
		jsonTeamNotification, _ := json.Marshal(dst.TeamNotification)
		if string(jsonTeamNotification) == "{}" { // empty struct
			dst.TeamNotification = nil
		} else {
			match++
		}
	} else {
		dst.TeamNotification = nil
	}

	// try to unmarshal data into UserNotification
	err = json.Unmarshal(data, &dst.UserNotification)
	if err == nil {
		jsonUserNotification, _ := json.Marshal(dst.UserNotification)
		if string(jsonUserNotification) == "{}" { // empty struct
			dst.UserNotification = nil
		} else {
			match++
		}
	} else {
		dst.UserNotification = nil
	}

	// try to unmarshal data into VictorOpsNotification
	err = json.Unmarshal(data, &dst.VictorOpsNotification)
	if err == nil {
		jsonVictorOpsNotification, _ := json.Marshal(dst.VictorOpsNotification)
		if string(jsonVictorOpsNotification) == "{}" { // empty struct
			dst.VictorOpsNotification = nil
		} else {
			match++
		}
	} else {
		dst.VictorOpsNotification = nil
	}

	// try to unmarshal data into WebhookNotification
	err = json.Unmarshal(data, &dst.WebhookNotification)
	if err == nil {
		jsonWebhookNotification, _ := json.Marshal(dst.WebhookNotification)
		if string(jsonWebhookNotification) == "{}" { // empty struct
			dst.WebhookNotification = nil
		} else {
			match++
		}
	} else {
		dst.WebhookNotification = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DatadogNotification = nil
		dst.EmailNotification = nil
		dst.GroupNotification = nil
		dst.HipChatNotification = nil
		dst.MicrosoftTeamsNotification = nil
		dst.OpsGenieNotification = nil
		dst.OrgNotification = nil
		dst.PagerDutyNotification = nil
		dst.SMSNotification = nil
		dst.SlackNotification = nil
		dst.TeamNotification = nil
		dst.UserNotification = nil
		dst.VictorOpsNotification = nil
		dst.WebhookNotification = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NotificationViewForNdsGroup)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NotificationViewForNdsGroup)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotificationViewForNdsGroup) MarshalJSON() ([]byte, error) {
	if src.DatadogNotification != nil {
		return json.Marshal(&src.DatadogNotification)
	}

	if src.EmailNotification != nil {
		return json.Marshal(&src.EmailNotification)
	}

	if src.GroupNotification != nil {
		return json.Marshal(&src.GroupNotification)
	}

	if src.HipChatNotification != nil {
		return json.Marshal(&src.HipChatNotification)
	}

	if src.MicrosoftTeamsNotification != nil {
		return json.Marshal(&src.MicrosoftTeamsNotification)
	}

	if src.OpsGenieNotification != nil {
		return json.Marshal(&src.OpsGenieNotification)
	}

	if src.OrgNotification != nil {
		return json.Marshal(&src.OrgNotification)
	}

	if src.PagerDutyNotification != nil {
		return json.Marshal(&src.PagerDutyNotification)
	}

	if src.SMSNotification != nil {
		return json.Marshal(&src.SMSNotification)
	}

	if src.SlackNotification != nil {
		return json.Marshal(&src.SlackNotification)
	}

	if src.TeamNotification != nil {
		return json.Marshal(&src.TeamNotification)
	}

	if src.UserNotification != nil {
		return json.Marshal(&src.UserNotification)
	}

	if src.VictorOpsNotification != nil {
		return json.Marshal(&src.VictorOpsNotification)
	}

	if src.WebhookNotification != nil {
		return json.Marshal(&src.WebhookNotification)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotificationViewForNdsGroup) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DatadogNotification != nil {
		return obj.DatadogNotification
	}

	if obj.EmailNotification != nil {
		return obj.EmailNotification
	}

	if obj.GroupNotification != nil {
		return obj.GroupNotification
	}

	if obj.HipChatNotification != nil {
		return obj.HipChatNotification
	}

	if obj.MicrosoftTeamsNotification != nil {
		return obj.MicrosoftTeamsNotification
	}

	if obj.OpsGenieNotification != nil {
		return obj.OpsGenieNotification
	}

	if obj.OrgNotification != nil {
		return obj.OrgNotification
	}

	if obj.PagerDutyNotification != nil {
		return obj.PagerDutyNotification
	}

	if obj.SMSNotification != nil {
		return obj.SMSNotification
	}

	if obj.SlackNotification != nil {
		return obj.SlackNotification
	}

	if obj.TeamNotification != nil {
		return obj.TeamNotification
	}

	if obj.UserNotification != nil {
		return obj.UserNotification
	}

	if obj.VictorOpsNotification != nil {
		return obj.VictorOpsNotification
	}

	if obj.WebhookNotification != nil {
		return obj.WebhookNotification
	}

	// all schemas are nil
	return nil
}

type NullableNotificationViewForNdsGroup struct {
	value *NotificationViewForNdsGroup
	isSet bool
}

func (v NullableNotificationViewForNdsGroup) Get() *NotificationViewForNdsGroup {
	return v.value
}

func (v *NullableNotificationViewForNdsGroup) Set(val *NotificationViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationViewForNdsGroup(val *NotificationViewForNdsGroup) *NullableNotificationViewForNdsGroup {
	return &NullableNotificationViewForNdsGroup{value: val, isSet: true}
}

func (v NullableNotificationViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


