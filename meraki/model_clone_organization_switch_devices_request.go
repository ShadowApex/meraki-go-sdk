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

// CloneOrganizationSwitchDevicesRequest struct for CloneOrganizationSwitchDevicesRequest
type CloneOrganizationSwitchDevicesRequest struct {
	// Serial number of the source switch (must be on a network not bound to a template)
	SourceSerial string `json:"sourceSerial"`
	// Array of serial numbers of one or more target switches (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials"`
}

// NewCloneOrganizationSwitchDevicesRequest instantiates a new CloneOrganizationSwitchDevicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneOrganizationSwitchDevicesRequest(sourceSerial string, targetSerials []string) *CloneOrganizationSwitchDevicesRequest {
	this := CloneOrganizationSwitchDevicesRequest{}
	this.SourceSerial = sourceSerial
	this.TargetSerials = targetSerials
	return &this
}

// NewCloneOrganizationSwitchDevicesRequestWithDefaults instantiates a new CloneOrganizationSwitchDevicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneOrganizationSwitchDevicesRequestWithDefaults() *CloneOrganizationSwitchDevicesRequest {
	this := CloneOrganizationSwitchDevicesRequest{}
	return &this
}

// GetSourceSerial returns the SourceSerial field value
func (o *CloneOrganizationSwitchDevicesRequest) GetSourceSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceSerial
}

// GetSourceSerialOk returns a tuple with the SourceSerial field value
// and a boolean to check if the value has been set.
func (o *CloneOrganizationSwitchDevicesRequest) GetSourceSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceSerial, true
}

// SetSourceSerial sets field value
func (o *CloneOrganizationSwitchDevicesRequest) SetSourceSerial(v string) {
	o.SourceSerial = v
}

// GetTargetSerials returns the TargetSerials field value
func (o *CloneOrganizationSwitchDevicesRequest) GetTargetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetSerials
}

// GetTargetSerialsOk returns a tuple with the TargetSerials field value
// and a boolean to check if the value has been set.
func (o *CloneOrganizationSwitchDevicesRequest) GetTargetSerialsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetSerials, true
}

// SetTargetSerials sets field value
func (o *CloneOrganizationSwitchDevicesRequest) SetTargetSerials(v []string) {
	o.TargetSerials = v
}

func (o CloneOrganizationSwitchDevicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceSerial"] = o.SourceSerial
	}
	if true {
		toSerialize["targetSerials"] = o.TargetSerials
	}
	return json.Marshal(toSerialize)
}

type NullableCloneOrganizationSwitchDevicesRequest struct {
	value *CloneOrganizationSwitchDevicesRequest
	isSet bool
}

func (v NullableCloneOrganizationSwitchDevicesRequest) Get() *CloneOrganizationSwitchDevicesRequest {
	return v.value
}

func (v *NullableCloneOrganizationSwitchDevicesRequest) Set(val *CloneOrganizationSwitchDevicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneOrganizationSwitchDevicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneOrganizationSwitchDevicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneOrganizationSwitchDevicesRequest(val *CloneOrganizationSwitchDevicesRequest) *NullableCloneOrganizationSwitchDevicesRequest {
	return &NullableCloneOrganizationSwitchDevicesRequest{value: val, isSet: true}
}

func (v NullableCloneOrganizationSwitchDevicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneOrganizationSwitchDevicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


