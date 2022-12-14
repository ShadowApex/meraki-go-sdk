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

// CreateOrganizationBrandingPolicyRequest struct for CreateOrganizationBrandingPolicyRequest
type CreateOrganizationBrandingPolicyRequest struct {
	// Name of the Dashboard branding policy.
	Name string `json:"name"`
	// Boolean indicating whether this policy is enabled.
	Enabled bool `json:"enabled"`
	AdminSettings CreateOrganizationBrandingPolicyRequestAdminSettings `json:"adminSettings"`
	HelpSettings *CreateOrganizationBrandingPolicyRequestHelpSettings `json:"helpSettings,omitempty"`
	CustomLogo *CreateOrganizationBrandingPolicyRequestCustomLogo `json:"customLogo,omitempty"`
}

// NewCreateOrganizationBrandingPolicyRequest instantiates a new CreateOrganizationBrandingPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationBrandingPolicyRequest(name string, enabled bool, adminSettings CreateOrganizationBrandingPolicyRequestAdminSettings) *CreateOrganizationBrandingPolicyRequest {
	this := CreateOrganizationBrandingPolicyRequest{}
	this.Name = name
	this.Enabled = enabled
	this.AdminSettings = adminSettings
	return &this
}

// NewCreateOrganizationBrandingPolicyRequestWithDefaults instantiates a new CreateOrganizationBrandingPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationBrandingPolicyRequestWithDefaults() *CreateOrganizationBrandingPolicyRequest {
	this := CreateOrganizationBrandingPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrganizationBrandingPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationBrandingPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *CreateOrganizationBrandingPolicyRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateOrganizationBrandingPolicyRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAdminSettings returns the AdminSettings field value
func (o *CreateOrganizationBrandingPolicyRequest) GetAdminSettings() CreateOrganizationBrandingPolicyRequestAdminSettings {
	if o == nil {
		var ret CreateOrganizationBrandingPolicyRequestAdminSettings
		return ret
	}

	return o.AdminSettings
}

// GetAdminSettingsOk returns a tuple with the AdminSettings field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetAdminSettingsOk() (*CreateOrganizationBrandingPolicyRequestAdminSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminSettings, true
}

// SetAdminSettings sets field value
func (o *CreateOrganizationBrandingPolicyRequest) SetAdminSettings(v CreateOrganizationBrandingPolicyRequestAdminSettings) {
	o.AdminSettings = v
}

// GetHelpSettings returns the HelpSettings field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequest) GetHelpSettings() CreateOrganizationBrandingPolicyRequestHelpSettings {
	if o == nil || o.HelpSettings == nil {
		var ret CreateOrganizationBrandingPolicyRequestHelpSettings
		return ret
	}
	return *o.HelpSettings
}

// GetHelpSettingsOk returns a tuple with the HelpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetHelpSettingsOk() (*CreateOrganizationBrandingPolicyRequestHelpSettings, bool) {
	if o == nil || o.HelpSettings == nil {
		return nil, false
	}
	return o.HelpSettings, true
}

// HasHelpSettings returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequest) HasHelpSettings() bool {
	if o != nil && o.HelpSettings != nil {
		return true
	}

	return false
}

// SetHelpSettings gets a reference to the given CreateOrganizationBrandingPolicyRequestHelpSettings and assigns it to the HelpSettings field.
func (o *CreateOrganizationBrandingPolicyRequest) SetHelpSettings(v CreateOrganizationBrandingPolicyRequestHelpSettings) {
	o.HelpSettings = &v
}

// GetCustomLogo returns the CustomLogo field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequest) GetCustomLogo() CreateOrganizationBrandingPolicyRequestCustomLogo {
	if o == nil || o.CustomLogo == nil {
		var ret CreateOrganizationBrandingPolicyRequestCustomLogo
		return ret
	}
	return *o.CustomLogo
}

// GetCustomLogoOk returns a tuple with the CustomLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequest) GetCustomLogoOk() (*CreateOrganizationBrandingPolicyRequestCustomLogo, bool) {
	if o == nil || o.CustomLogo == nil {
		return nil, false
	}
	return o.CustomLogo, true
}

// HasCustomLogo returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequest) HasCustomLogo() bool {
	if o != nil && o.CustomLogo != nil {
		return true
	}

	return false
}

// SetCustomLogo gets a reference to the given CreateOrganizationBrandingPolicyRequestCustomLogo and assigns it to the CustomLogo field.
func (o *CreateOrganizationBrandingPolicyRequest) SetCustomLogo(v CreateOrganizationBrandingPolicyRequestCustomLogo) {
	o.CustomLogo = &v
}

func (o CreateOrganizationBrandingPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["adminSettings"] = o.AdminSettings
	}
	if o.HelpSettings != nil {
		toSerialize["helpSettings"] = o.HelpSettings
	}
	if o.CustomLogo != nil {
		toSerialize["customLogo"] = o.CustomLogo
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationBrandingPolicyRequest struct {
	value *CreateOrganizationBrandingPolicyRequest
	isSet bool
}

func (v NullableCreateOrganizationBrandingPolicyRequest) Get() *CreateOrganizationBrandingPolicyRequest {
	return v.value
}

func (v *NullableCreateOrganizationBrandingPolicyRequest) Set(val *CreateOrganizationBrandingPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationBrandingPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationBrandingPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationBrandingPolicyRequest(val *CreateOrganizationBrandingPolicyRequest) *NullableCreateOrganizationBrandingPolicyRequest {
	return &NullableCreateOrganizationBrandingPolicyRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationBrandingPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationBrandingPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


