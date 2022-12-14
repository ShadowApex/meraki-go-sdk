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

// UpdateNetworkSettingsRequestLocalStatusPage A hash of Local Status page(s) options applied to the Network.
type UpdateNetworkSettingsRequestLocalStatusPage struct {
	Authentication *UpdateNetworkSettingsRequestLocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// NewUpdateNetworkSettingsRequestLocalStatusPage instantiates a new UpdateNetworkSettingsRequestLocalStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSettingsRequestLocalStatusPage() *UpdateNetworkSettingsRequestLocalStatusPage {
	this := UpdateNetworkSettingsRequestLocalStatusPage{}
	return &this
}

// NewUpdateNetworkSettingsRequestLocalStatusPageWithDefaults instantiates a new UpdateNetworkSettingsRequestLocalStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSettingsRequestLocalStatusPageWithDefaults() *UpdateNetworkSettingsRequestLocalStatusPage {
	this := UpdateNetworkSettingsRequestLocalStatusPage{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *UpdateNetworkSettingsRequestLocalStatusPage) GetAuthentication() UpdateNetworkSettingsRequestLocalStatusPageAuthentication {
	if o == nil || o.Authentication == nil {
		var ret UpdateNetworkSettingsRequestLocalStatusPageAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSettingsRequestLocalStatusPage) GetAuthenticationOk() (*UpdateNetworkSettingsRequestLocalStatusPageAuthentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *UpdateNetworkSettingsRequestLocalStatusPage) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given UpdateNetworkSettingsRequestLocalStatusPageAuthentication and assigns it to the Authentication field.
func (o *UpdateNetworkSettingsRequestLocalStatusPage) SetAuthentication(v UpdateNetworkSettingsRequestLocalStatusPageAuthentication) {
	o.Authentication = &v
}

func (o UpdateNetworkSettingsRequestLocalStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSettingsRequestLocalStatusPage struct {
	value *UpdateNetworkSettingsRequestLocalStatusPage
	isSet bool
}

func (v NullableUpdateNetworkSettingsRequestLocalStatusPage) Get() *UpdateNetworkSettingsRequestLocalStatusPage {
	return v.value
}

func (v *NullableUpdateNetworkSettingsRequestLocalStatusPage) Set(val *UpdateNetworkSettingsRequestLocalStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSettingsRequestLocalStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSettingsRequestLocalStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSettingsRequestLocalStatusPage(val *UpdateNetworkSettingsRequestLocalStatusPage) *NullableUpdateNetworkSettingsRequestLocalStatusPage {
	return &NullableUpdateNetworkSettingsRequestLocalStatusPage{value: val, isSet: true}
}

func (v NullableUpdateNetworkSettingsRequestLocalStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSettingsRequestLocalStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


