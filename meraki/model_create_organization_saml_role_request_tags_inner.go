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

// CreateOrganizationSamlRoleRequestTagsInner struct for CreateOrganizationSamlRoleRequestTagsInner
type CreateOrganizationSamlRoleRequestTagsInner struct {
	// The name of the tag
	Tag string `json:"tag"`
	// The privilege of the SAML administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Access string `json:"access"`
}

// NewCreateOrganizationSamlRoleRequestTagsInner instantiates a new CreateOrganizationSamlRoleRequestTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationSamlRoleRequestTagsInner(tag string, access string) *CreateOrganizationSamlRoleRequestTagsInner {
	this := CreateOrganizationSamlRoleRequestTagsInner{}
	this.Tag = tag
	this.Access = access
	return &this
}

// NewCreateOrganizationSamlRoleRequestTagsInnerWithDefaults instantiates a new CreateOrganizationSamlRoleRequestTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationSamlRoleRequestTagsInnerWithDefaults() *CreateOrganizationSamlRoleRequestTagsInner {
	this := CreateOrganizationSamlRoleRequestTagsInner{}
	return &this
}

// GetTag returns the Tag field value
func (o *CreateOrganizationSamlRoleRequestTagsInner) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequestTagsInner) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *CreateOrganizationSamlRoleRequestTagsInner) SetTag(v string) {
	o.Tag = v
}

// GetAccess returns the Access field value
func (o *CreateOrganizationSamlRoleRequestTagsInner) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequestTagsInner) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *CreateOrganizationSamlRoleRequestTagsInner) SetAccess(v string) {
	o.Access = v
}

func (o CreateOrganizationSamlRoleRequestTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationSamlRoleRequestTagsInner struct {
	value *CreateOrganizationSamlRoleRequestTagsInner
	isSet bool
}

func (v NullableCreateOrganizationSamlRoleRequestTagsInner) Get() *CreateOrganizationSamlRoleRequestTagsInner {
	return v.value
}

func (v *NullableCreateOrganizationSamlRoleRequestTagsInner) Set(val *CreateOrganizationSamlRoleRequestTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationSamlRoleRequestTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationSamlRoleRequestTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationSamlRoleRequestTagsInner(val *CreateOrganizationSamlRoleRequestTagsInner) *NullableCreateOrganizationSamlRoleRequestTagsInner {
	return &NullableCreateOrganizationSamlRoleRequestTagsInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationSamlRoleRequestTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationSamlRoleRequestTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


