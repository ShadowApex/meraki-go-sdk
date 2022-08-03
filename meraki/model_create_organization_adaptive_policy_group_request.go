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

// CreateOrganizationAdaptivePolicyGroupRequest struct for CreateOrganizationAdaptivePolicyGroupRequest
type CreateOrganizationAdaptivePolicyGroupRequest struct {
	// Name of the group
	Name string `json:"name"`
	// SGT value of the group
	Sgt int32 `json:"sgt"`
	// Description of the group (default: \"\")
	Description *string `json:"description,omitempty"`
	// The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group's SGT value if no other tagging scheme is being used (each requires one unique attribute) (default: [])
	PolicyObjects []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner `json:"policyObjects,omitempty"`
}

// NewCreateOrganizationAdaptivePolicyGroupRequest instantiates a new CreateOrganizationAdaptivePolicyGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationAdaptivePolicyGroupRequest(name string, sgt int32) *CreateOrganizationAdaptivePolicyGroupRequest {
	this := CreateOrganizationAdaptivePolicyGroupRequest{}
	this.Name = name
	this.Sgt = sgt
	return &this
}

// NewCreateOrganizationAdaptivePolicyGroupRequestWithDefaults instantiates a new CreateOrganizationAdaptivePolicyGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationAdaptivePolicyGroupRequestWithDefaults() *CreateOrganizationAdaptivePolicyGroupRequest {
	this := CreateOrganizationAdaptivePolicyGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationAdaptivePolicyGroupRequest) SetName(v string) {
	o.Name = v
}

// GetSgt returns the Sgt field value
func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetSgt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sgt
}

// GetSgtOk returns a tuple with the Sgt field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetSgtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sgt, true
}

// SetSgt sets field value
func (o *CreateOrganizationAdaptivePolicyGroupRequest) SetSgt(v int32) {
	o.Sgt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetPolicyObjects() []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner {
	if o == nil || o.PolicyObjects == nil {
		var ret []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner
		return ret
	}
	return o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) GetPolicyObjectsOk() ([]CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner, bool) {
	if o == nil || o.PolicyObjects == nil {
		return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) HasPolicyObjects() bool {
	if o != nil && o.PolicyObjects != nil {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner and assigns it to the PolicyObjects field.
func (o *CreateOrganizationAdaptivePolicyGroupRequest) SetPolicyObjects(v []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) {
	o.PolicyObjects = v
}

func (o CreateOrganizationAdaptivePolicyGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["sgt"] = o.Sgt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PolicyObjects != nil {
		toSerialize["policyObjects"] = o.PolicyObjects
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationAdaptivePolicyGroupRequest struct {
	value *CreateOrganizationAdaptivePolicyGroupRequest
	isSet bool
}

func (v NullableCreateOrganizationAdaptivePolicyGroupRequest) Get() *CreateOrganizationAdaptivePolicyGroupRequest {
	return v.value
}

func (v *NullableCreateOrganizationAdaptivePolicyGroupRequest) Set(val *CreateOrganizationAdaptivePolicyGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationAdaptivePolicyGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationAdaptivePolicyGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationAdaptivePolicyGroupRequest(val *CreateOrganizationAdaptivePolicyGroupRequest) *NullableCreateOrganizationAdaptivePolicyGroupRequest {
	return &NullableCreateOrganizationAdaptivePolicyGroupRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationAdaptivePolicyGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationAdaptivePolicyGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


