/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 July, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.23.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"encoding/json"
)

// CreateOrganizationAlertsProfileRequest struct for CreateOrganizationAlertsProfileRequest
type CreateOrganizationAlertsProfileRequest struct {
	// The alert type
	Type string `json:"type"`
	AlertCondition CreateOrganizationAlertsProfileRequestAlertCondition `json:"alertCondition"`
	Recipients CreateOrganizationAlertsProfileRequestRecipients `json:"recipients"`
	// Networks with these tags will be monitored for the alert
	NetworkTags []string `json:"networkTags"`
	// User supplied description of the alert
	Description *string `json:"description,omitempty"`
}

// NewCreateOrganizationAlertsProfileRequest instantiates a new CreateOrganizationAlertsProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationAlertsProfileRequest(type_ string, alertCondition CreateOrganizationAlertsProfileRequestAlertCondition, recipients CreateOrganizationAlertsProfileRequestRecipients, networkTags []string) *CreateOrganizationAlertsProfileRequest {
	this := CreateOrganizationAlertsProfileRequest{}
	this.Type = type_
	this.AlertCondition = alertCondition
	this.Recipients = recipients
	this.NetworkTags = networkTags
	return &this
}

// NewCreateOrganizationAlertsProfileRequestWithDefaults instantiates a new CreateOrganizationAlertsProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationAlertsProfileRequestWithDefaults() *CreateOrganizationAlertsProfileRequest {
	this := CreateOrganizationAlertsProfileRequest{}
	return &this
}

// GetType returns the Type field value
func (o *CreateOrganizationAlertsProfileRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateOrganizationAlertsProfileRequest) SetType(v string) {
	o.Type = v
}

// GetAlertCondition returns the AlertCondition field value
func (o *CreateOrganizationAlertsProfileRequest) GetAlertCondition() CreateOrganizationAlertsProfileRequestAlertCondition {
	if o == nil {
		var ret CreateOrganizationAlertsProfileRequestAlertCondition
		return ret
	}

	return o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequest) GetAlertConditionOk() (*CreateOrganizationAlertsProfileRequestAlertCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertCondition, true
}

// SetAlertCondition sets field value
func (o *CreateOrganizationAlertsProfileRequest) SetAlertCondition(v CreateOrganizationAlertsProfileRequestAlertCondition) {
	o.AlertCondition = v
}

// GetRecipients returns the Recipients field value
func (o *CreateOrganizationAlertsProfileRequest) GetRecipients() CreateOrganizationAlertsProfileRequestRecipients {
	if o == nil {
		var ret CreateOrganizationAlertsProfileRequestRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequest) GetRecipientsOk() (*CreateOrganizationAlertsProfileRequestRecipients, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *CreateOrganizationAlertsProfileRequest) SetRecipients(v CreateOrganizationAlertsProfileRequestRecipients) {
	o.Recipients = v
}

// GetNetworkTags returns the NetworkTags field value
func (o *CreateOrganizationAlertsProfileRequest) GetNetworkTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequest) GetNetworkTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkTags, true
}

// SetNetworkTags sets field value
func (o *CreateOrganizationAlertsProfileRequest) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrganizationAlertsProfileRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrganizationAlertsProfileRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrganizationAlertsProfileRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CreateOrganizationAlertsProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["alertCondition"] = o.AlertCondition
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["networkTags"] = o.NetworkTags
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationAlertsProfileRequest struct {
	value *CreateOrganizationAlertsProfileRequest
	isSet bool
}

func (v NullableCreateOrganizationAlertsProfileRequest) Get() *CreateOrganizationAlertsProfileRequest {
	return v.value
}

func (v *NullableCreateOrganizationAlertsProfileRequest) Set(val *CreateOrganizationAlertsProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationAlertsProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationAlertsProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationAlertsProfileRequest(val *CreateOrganizationAlertsProfileRequest) *NullableCreateOrganizationAlertsProfileRequest {
	return &NullableCreateOrganizationAlertsProfileRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationAlertsProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationAlertsProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


