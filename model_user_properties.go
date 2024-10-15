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

// UserProperties Mongodb user properties.
type UserProperties struct {
	Username *string      `json:"username"`
	Password *string      `json:"password"`
	Roles    *[]UserRoles `json:"roles,omitempty"`
}

// NewUserProperties instantiates a new UserProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProperties(username string, password string) *UserProperties {
	this := UserProperties{}

	this.Username = &username
	this.Password = &password

	return &this
}

// NewUserPropertiesWithDefaults instantiates a new UserProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPropertiesWithDefaults() *UserProperties {
	this := UserProperties{}
	return &this
}

// GetUsername returns the Username field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserProperties) GetUsername() *string {
	if o == nil {
		return nil
	}

	return o.Username

}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProperties) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Username, true
}

// SetUsername sets field value
func (o *UserProperties) SetUsername(v string) {

	o.Username = &v

}

// HasUsername returns a boolean if a field has been set.
func (o *UserProperties) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// GetPassword returns the Password field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserProperties) GetPassword() *string {
	if o == nil {
		return nil
	}

	return o.Password

}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProperties) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Password, true
}

// SetPassword sets field value
func (o *UserProperties) SetPassword(v string) {

	o.Password = &v

}

// HasPassword returns a boolean if a field has been set.
func (o *UserProperties) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// GetRoles returns the Roles field value
// If the value is explicit nil, the zero value for []UserRoles will be returned
func (o *UserProperties) GetRoles() *[]UserRoles {
	if o == nil {
		return nil
	}

	return o.Roles

}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProperties) GetRolesOk() (*[]UserRoles, bool) {
	if o == nil {
		return nil, false
	}

	return o.Roles, true
}

// SetRoles sets field value
func (o *UserProperties) SetRoles(v []UserRoles) {

	o.Roles = &v

}

// HasRoles returns a boolean if a field has been set.
func (o *UserProperties) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

func (o UserProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}

	if o.Password != nil {
		toSerialize["password"] = o.Password
	}

	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}

	return json.Marshal(toSerialize)
}

type NullableUserProperties struct {
	value *UserProperties
	isSet bool
}

func (v NullableUserProperties) Get() *UserProperties {
	return v.value
}

func (v *NullableUserProperties) Set(val *UserProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProperties(val *UserProperties) *NullableUserProperties {
	return &NullableUserProperties{value: val, isSet: true}
}

func (v NullableUserProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
