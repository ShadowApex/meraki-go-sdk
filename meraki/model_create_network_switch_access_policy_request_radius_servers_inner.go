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

// CreateNetworkSwitchAccessPolicyRequestRadiusServersInner struct for CreateNetworkSwitchAccessPolicyRequestRadiusServersInner
type CreateNetworkSwitchAccessPolicyRequestRadiusServersInner struct {
	// Public IP address of the RADIUS server
	Host string `json:"host"`
	// UDP port that the RADIUS server listens on for access requests
	Port int32 `json:"port"`
	// RADIUS client shared secret
	Secret string `json:"secret"`
}

// NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInner instantiates a new CreateNetworkSwitchAccessPolicyRequestRadiusServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInner(host string, port int32, secret string) *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner {
	this := CreateNetworkSwitchAccessPolicyRequestRadiusServersInner{}
	this.Host = host
	this.Port = port
	this.Secret = secret
	return &this
}

// NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInnerWithDefaults instantiates a new CreateNetworkSwitchAccessPolicyRequestRadiusServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInnerWithDefaults() *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner {
	this := CreateNetworkSwitchAccessPolicyRequestRadiusServersInner{}
	return &this
}

// GetHost returns the Host field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) SetPort(v int32) {
	o.Port = v
}

// GetSecret returns the Secret field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) SetSecret(v string) {
	o.Secret = v
}

func (o CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner struct {
	value *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner
	isSet bool
}

func (v NullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner) Get() *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner {
	return v.value
}

func (v *NullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner) Set(val *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner(val *CreateNetworkSwitchAccessPolicyRequestRadiusServersInner) *NullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner {
	return &NullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchAccessPolicyRequestRadiusServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


