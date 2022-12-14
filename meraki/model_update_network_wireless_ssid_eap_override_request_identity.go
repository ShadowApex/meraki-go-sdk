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

// UpdateNetworkWirelessSsidEapOverrideRequestIdentity EAP settings for identity requests.
type UpdateNetworkWirelessSsidEapOverrideRequestIdentity struct {
	// Maximum number of EAP retries.
	Retries *int32 `json:"retries,omitempty"`
	// EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewUpdateNetworkWirelessSsidEapOverrideRequestIdentity instantiates a new UpdateNetworkWirelessSsidEapOverrideRequestIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidEapOverrideRequestIdentity() *UpdateNetworkWirelessSsidEapOverrideRequestIdentity {
	this := UpdateNetworkWirelessSsidEapOverrideRequestIdentity{}
	return &this
}

// NewUpdateNetworkWirelessSsidEapOverrideRequestIdentityWithDefaults instantiates a new UpdateNetworkWirelessSsidEapOverrideRequestIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidEapOverrideRequestIdentityWithDefaults() *UpdateNetworkWirelessSsidEapOverrideRequestIdentity {
	this := UpdateNetworkWirelessSsidEapOverrideRequestIdentity{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) GetRetries() int32 {
	if o == nil || o.Retries == nil {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) GetRetriesOk() (*int32, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) HasRetries() bool {
	if o != nil && o.Retries != nil {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) SetRetries(v int32) {
	o.Retries = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o UpdateNetworkWirelessSsidEapOverrideRequestIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity struct {
	value *UpdateNetworkWirelessSsidEapOverrideRequestIdentity
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity) Get() *UpdateNetworkWirelessSsidEapOverrideRequestIdentity {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity) Set(val *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity(val *UpdateNetworkWirelessSsidEapOverrideRequestIdentity) *NullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity {
	return &NullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidEapOverrideRequestIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


