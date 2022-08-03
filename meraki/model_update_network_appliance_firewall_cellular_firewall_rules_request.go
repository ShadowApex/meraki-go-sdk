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

// UpdateNetworkApplianceFirewallCellularFirewallRulesRequest struct for UpdateNetworkApplianceFirewallCellularFirewallRulesRequest
type UpdateNetworkApplianceFirewallCellularFirewallRulesRequest struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner `json:"rules,omitempty"`
}

// NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequest instantiates a new UpdateNetworkApplianceFirewallCellularFirewallRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequest() *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest {
	this := UpdateNetworkApplianceFirewallCellularFirewallRulesRequest{}
	return &this
}

// NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequestWithDefaults instantiates a new UpdateNetworkApplianceFirewallCellularFirewallRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequestWithDefaults() *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest {
	this := UpdateNetworkApplianceFirewallCellularFirewallRulesRequest{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest) GetRules() []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner {
	if o == nil || o.Rules == nil {
		var ret []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest) GetRulesOk() ([]UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner and assigns it to the Rules field.
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest) SetRules(v []UpdateNetworkApplianceFirewallCellularFirewallRulesRequestRulesInner) {
	o.Rules = v
}

func (o UpdateNetworkApplianceFirewallCellularFirewallRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest struct {
	value *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest) Get() *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest) Set(val *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest(val *UpdateNetworkApplianceFirewallCellularFirewallRulesRequest) *NullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest {
	return &NullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallCellularFirewallRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


