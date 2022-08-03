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

// GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients Clients info
type GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients struct {
	Counts *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts `json:"counts,omitempty"`
}

// NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients instantiates a new GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients {
	this := GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients{}
	return &this
}

// NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsWithDefaults instantiates a new GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsWithDefaults() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients {
	this := GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) GetCounts() GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts {
	if o == nil || o.Counts == nil {
		var ret GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) GetCountsOk() (*GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts, bool) {
	if o == nil || o.Counts == nil {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) HasCounts() bool {
	if o != nil && o.Counts != nil {
		return true
	}

	return false
}

// SetCounts gets a reference to the given GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts and assigns it to the Counts field.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) SetCounts(v GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) {
	o.Counts = &v
}

func (o GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Counts != nil {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients struct {
	value *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients
	isSet bool
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) Get() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) Set(val *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients(val *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients {
	return &NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


