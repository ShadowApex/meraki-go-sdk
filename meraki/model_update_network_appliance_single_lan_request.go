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

// UpdateNetworkApplianceSingleLanRequest struct for UpdateNetworkApplianceSingleLanRequest
type UpdateNetworkApplianceSingleLanRequest struct {
	// The subnet of the single LAN configuration
	Subnet *string `json:"subnet,omitempty"`
	// The appliance IP address of the single LAN
	ApplianceIp *string `json:"applianceIp,omitempty"`
}

// NewUpdateNetworkApplianceSingleLanRequest instantiates a new UpdateNetworkApplianceSingleLanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSingleLanRequest() *UpdateNetworkApplianceSingleLanRequest {
	this := UpdateNetworkApplianceSingleLanRequest{}
	return &this
}

// NewUpdateNetworkApplianceSingleLanRequestWithDefaults instantiates a new UpdateNetworkApplianceSingleLanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSingleLanRequestWithDefaults() *UpdateNetworkApplianceSingleLanRequest {
	this := UpdateNetworkApplianceSingleLanRequest{}
	return &this
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequest) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequest) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequest) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *UpdateNetworkApplianceSingleLanRequest) SetSubnet(v string) {
	o.Subnet = &v
}

// GetApplianceIp returns the ApplianceIp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequest) GetApplianceIp() string {
	if o == nil || o.ApplianceIp == nil {
		var ret string
		return ret
	}
	return *o.ApplianceIp
}

// GetApplianceIpOk returns a tuple with the ApplianceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequest) GetApplianceIpOk() (*string, bool) {
	if o == nil || o.ApplianceIp == nil {
		return nil, false
	}
	return o.ApplianceIp, true
}

// HasApplianceIp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequest) HasApplianceIp() bool {
	if o != nil && o.ApplianceIp != nil {
		return true
	}

	return false
}

// SetApplianceIp gets a reference to the given string and assigns it to the ApplianceIp field.
func (o *UpdateNetworkApplianceSingleLanRequest) SetApplianceIp(v string) {
	o.ApplianceIp = &v
}

func (o UpdateNetworkApplianceSingleLanRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.ApplianceIp != nil {
		toSerialize["applianceIp"] = o.ApplianceIp
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceSingleLanRequest struct {
	value *UpdateNetworkApplianceSingleLanRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceSingleLanRequest) Get() *UpdateNetworkApplianceSingleLanRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSingleLanRequest) Set(val *UpdateNetworkApplianceSingleLanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSingleLanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSingleLanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSingleLanRequest(val *UpdateNetworkApplianceSingleLanRequest) *NullableUpdateNetworkApplianceSingleLanRequest {
	return &NullableUpdateNetworkApplianceSingleLanRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSingleLanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSingleLanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


