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

// UpdateNetworkWirelessSsidRequestSpeedBurst The SpeedBurst setting for this SSID'
type UpdateNetworkWirelessSsidRequestSpeedBurst struct {
	// Boolean indicating whether or not to allow users to temporarily exceed the bandwidth limit for short periods while still keeping them under the bandwidth limit over time.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestSpeedBurst instantiates a new UpdateNetworkWirelessSsidRequestSpeedBurst object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestSpeedBurst() *UpdateNetworkWirelessSsidRequestSpeedBurst {
	this := UpdateNetworkWirelessSsidRequestSpeedBurst{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestSpeedBurstWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestSpeedBurst object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestSpeedBurstWithDefaults() *UpdateNetworkWirelessSsidRequestSpeedBurst {
	this := UpdateNetworkWirelessSsidRequestSpeedBurst{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestSpeedBurst) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestSpeedBurst) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestSpeedBurst) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidRequestSpeedBurst) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkWirelessSsidRequestSpeedBurst) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidRequestSpeedBurst struct {
	value *UpdateNetworkWirelessSsidRequestSpeedBurst
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestSpeedBurst) Get() *UpdateNetworkWirelessSsidRequestSpeedBurst {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestSpeedBurst) Set(val *UpdateNetworkWirelessSsidRequestSpeedBurst) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestSpeedBurst) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestSpeedBurst) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestSpeedBurst(val *UpdateNetworkWirelessSsidRequestSpeedBurst) *NullableUpdateNetworkWirelessSsidRequestSpeedBurst {
	return &NullableUpdateNetworkWirelessSsidRequestSpeedBurst{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestSpeedBurst) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestSpeedBurst) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


