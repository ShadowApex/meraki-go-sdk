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

// UpdateNetworkWirelessSsidEapOverrideRequestEapolKey EAPOL Key settings.
type UpdateNetworkWirelessSsidEapOverrideRequestEapolKey struct {
	// Maximum number of EAPOL key retries.
	Retries *int32 `json:"retries,omitempty"`
	// EAPOL Key timeout in milliseconds.
	TimeoutInMs *int32 `json:"timeoutInMs,omitempty"`
}

// NewUpdateNetworkWirelessSsidEapOverrideRequestEapolKey instantiates a new UpdateNetworkWirelessSsidEapOverrideRequestEapolKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidEapOverrideRequestEapolKey() *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey {
	this := UpdateNetworkWirelessSsidEapOverrideRequestEapolKey{}
	return &this
}

// NewUpdateNetworkWirelessSsidEapOverrideRequestEapolKeyWithDefaults instantiates a new UpdateNetworkWirelessSsidEapOverrideRequestEapolKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidEapOverrideRequestEapolKeyWithDefaults() *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey {
	this := UpdateNetworkWirelessSsidEapOverrideRequestEapolKey{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) GetRetries() int32 {
	if o == nil || o.Retries == nil {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) GetRetriesOk() (*int32, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) HasRetries() bool {
	if o != nil && o.Retries != nil {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) SetRetries(v int32) {
	o.Retries = &v
}

// GetTimeoutInMs returns the TimeoutInMs field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) GetTimeoutInMs() int32 {
	if o == nil || o.TimeoutInMs == nil {
		var ret int32
		return ret
	}
	return *o.TimeoutInMs
}

// GetTimeoutInMsOk returns a tuple with the TimeoutInMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) GetTimeoutInMsOk() (*int32, bool) {
	if o == nil || o.TimeoutInMs == nil {
		return nil, false
	}
	return o.TimeoutInMs, true
}

// HasTimeoutInMs returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) HasTimeoutInMs() bool {
	if o != nil && o.TimeoutInMs != nil {
		return true
	}

	return false
}

// SetTimeoutInMs gets a reference to the given int32 and assigns it to the TimeoutInMs field.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) SetTimeoutInMs(v int32) {
	o.TimeoutInMs = &v
}

func (o UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}
	if o.TimeoutInMs != nil {
		toSerialize["timeoutInMs"] = o.TimeoutInMs
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey struct {
	value *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey) Get() *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey) Set(val *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey(val *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) *NullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey {
	return &NullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidEapOverrideRequestEapolKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


