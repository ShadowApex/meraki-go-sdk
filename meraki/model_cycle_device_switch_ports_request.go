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

// CycleDeviceSwitchPortsRequest struct for CycleDeviceSwitchPortsRequest
type CycleDeviceSwitchPortsRequest struct {
	// List of switch ports. Example: [1, 2-5, 1_MA-MOD-8X10G_1, 1_MA-MOD-8X10G_2-1_MA-MOD-8X10G_8]
	Ports []string `json:"ports"`
}

// NewCycleDeviceSwitchPortsRequest instantiates a new CycleDeviceSwitchPortsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCycleDeviceSwitchPortsRequest(ports []string) *CycleDeviceSwitchPortsRequest {
	this := CycleDeviceSwitchPortsRequest{}
	this.Ports = ports
	return &this
}

// NewCycleDeviceSwitchPortsRequestWithDefaults instantiates a new CycleDeviceSwitchPortsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCycleDeviceSwitchPortsRequestWithDefaults() *CycleDeviceSwitchPortsRequest {
	this := CycleDeviceSwitchPortsRequest{}
	return &this
}

// GetPorts returns the Ports field value
func (o *CycleDeviceSwitchPortsRequest) GetPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *CycleDeviceSwitchPortsRequest) GetPortsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *CycleDeviceSwitchPortsRequest) SetPorts(v []string) {
	o.Ports = v
}

func (o CycleDeviceSwitchPortsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableCycleDeviceSwitchPortsRequest struct {
	value *CycleDeviceSwitchPortsRequest
	isSet bool
}

func (v NullableCycleDeviceSwitchPortsRequest) Get() *CycleDeviceSwitchPortsRequest {
	return v.value
}

func (v *NullableCycleDeviceSwitchPortsRequest) Set(val *CycleDeviceSwitchPortsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCycleDeviceSwitchPortsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCycleDeviceSwitchPortsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCycleDeviceSwitchPortsRequest(val *CycleDeviceSwitchPortsRequest) *NullableCycleDeviceSwitchPortsRequest {
	return &NullableCycleDeviceSwitchPortsRequest{value: val, isSet: true}
}

func (v NullableCycleDeviceSwitchPortsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCycleDeviceSwitchPortsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


