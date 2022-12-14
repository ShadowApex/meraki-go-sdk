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

// UpdateOrganizationRequest struct for UpdateOrganizationRequest
type UpdateOrganizationRequest struct {
	// The name of the organization
	Name *string `json:"name,omitempty"`
	Api *UpdateOrganizationRequestApi `json:"api,omitempty"`
}

// NewUpdateOrganizationRequest instantiates a new UpdateOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationRequest() *UpdateOrganizationRequest {
	this := UpdateOrganizationRequest{}
	return &this
}

// NewUpdateOrganizationRequestWithDefaults instantiates a new UpdateOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationRequestWithDefaults() *UpdateOrganizationRequest {
	this := UpdateOrganizationRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrganizationRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrganizationRequest) SetName(v string) {
	o.Name = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *UpdateOrganizationRequest) GetApi() UpdateOrganizationRequestApi {
	if o == nil || o.Api == nil {
		var ret UpdateOrganizationRequestApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationRequest) GetApiOk() (*UpdateOrganizationRequestApi, bool) {
	if o == nil || o.Api == nil {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *UpdateOrganizationRequest) HasApi() bool {
	if o != nil && o.Api != nil {
		return true
	}

	return false
}

// SetApi gets a reference to the given UpdateOrganizationRequestApi and assigns it to the Api field.
func (o *UpdateOrganizationRequest) SetApi(v UpdateOrganizationRequestApi) {
	o.Api = &v
}

func (o UpdateOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Api != nil {
		toSerialize["api"] = o.Api
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationRequest struct {
	value *UpdateOrganizationRequest
	isSet bool
}

func (v NullableUpdateOrganizationRequest) Get() *UpdateOrganizationRequest {
	return v.value
}

func (v *NullableUpdateOrganizationRequest) Set(val *UpdateOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationRequest(val *UpdateOrganizationRequest) *NullableUpdateOrganizationRequest {
	return &NullableUpdateOrganizationRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


