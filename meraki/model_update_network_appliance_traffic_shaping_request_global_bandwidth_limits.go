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

// UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits Global per-client bandwidth limit
type UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits struct {
	// The upload bandwidth limit in Kbps. (0 represents no limit.)
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The download bandwidth limit in Kbps. (0 represents no limit.)
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits instantiates a new UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits() *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits {
	this := UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimitsWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimitsWithDefaults() *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits {
	this := UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) GetLimitUp() int32 {
	if o == nil || o.LimitUp == nil {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) GetLimitUpOk() (*int32, bool) {
	if o == nil || o.LimitUp == nil {
		return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) HasLimitUp() bool {
	if o != nil && o.LimitUp != nil {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) GetLimitDown() int32 {
	if o == nil || o.LimitDown == nil {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) GetLimitDownOk() (*int32, bool) {
	if o == nil || o.LimitDown == nil {
		return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) HasLimitDown() bool {
	if o != nil && o.LimitDown != nil {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LimitUp != nil {
		toSerialize["limitUp"] = o.LimitUp
	}
	if o.LimitDown != nil {
		toSerialize["limitDown"] = o.LimitDown
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits struct {
	value *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) Get() *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) Set(val *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits(val *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) *NullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits {
	return &NullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


