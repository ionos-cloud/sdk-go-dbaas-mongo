/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.  MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// PatchUserProperties MongoDB database user patch request properties.
type PatchUserProperties struct {
	Password *string      `json:"password,omitempty"`
	Roles    *[]UserRoles `json:"roles,omitempty"`
}

// NewPatchUserProperties instantiates a new PatchUserProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchUserProperties() *PatchUserProperties {
	this := PatchUserProperties{}

	return &this
}

// NewPatchUserPropertiesWithDefaults instantiates a new PatchUserProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchUserPropertiesWithDefaults() *PatchUserProperties {
	this := PatchUserProperties{}
	return &this
}

// GetPassword returns the Password field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PatchUserProperties) GetPassword() *string {
	if o == nil {
		return nil
	}

	return o.Password

}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchUserProperties) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Password, true
}

// SetPassword sets field value
func (o *PatchUserProperties) SetPassword(v string) {

	o.Password = &v

}

// HasPassword returns a boolean if a field has been set.
func (o *PatchUserProperties) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// GetRoles returns the Roles field value
// If the value is explicit nil, the zero value for []UserRoles will be returned
func (o *PatchUserProperties) GetRoles() *[]UserRoles {
	if o == nil {
		return nil
	}

	return o.Roles

}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchUserProperties) GetRolesOk() (*[]UserRoles, bool) {
	if o == nil {
		return nil, false
	}

	return o.Roles, true
}

// SetRoles sets field value
func (o *PatchUserProperties) SetRoles(v []UserRoles) {

	o.Roles = &v

}

// HasRoles returns a boolean if a field has been set.
func (o *PatchUserProperties) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

func (o PatchUserProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}

	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}

	return json.Marshal(toSerialize)
}

type NullablePatchUserProperties struct {
	value *PatchUserProperties
	isSet bool
}

func (v NullablePatchUserProperties) Get() *PatchUserProperties {
	return v.value
}

func (v *NullablePatchUserProperties) Set(val *PatchUserProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchUserProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchUserProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchUserProperties(val *PatchUserProperties) *NullablePatchUserProperties {
	return &NullablePatchUserProperties{value: val, isSet: true}
}

func (v NullablePatchUserProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchUserProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
