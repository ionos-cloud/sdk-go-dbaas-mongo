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

// TemplateProperties The properties of a MongoDB template.
type TemplateProperties struct {
	// The name of the template.
	Name *string `json:"name,omitempty"`
	// The edition of the template (e.g. enterprise)
	Edition *string `json:"edition,omitempty"`
	// The number of CPU cores.
	Cores *int32 `json:"cores,omitempty"`
	// The amount of memory in MB.
	Ram *int32 `json:"ram,omitempty"`
	// The amount of storage size in GB.
	StorageSize *int32 `json:"storageSize,omitempty"`
}

// NewTemplateProperties instantiates a new TemplateProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateProperties() *TemplateProperties {
	this := TemplateProperties{}

	return &this
}

// NewTemplatePropertiesWithDefaults instantiates a new TemplateProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatePropertiesWithDefaults() *TemplateProperties {
	this := TemplateProperties{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplateProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *TemplateProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *TemplateProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetEdition returns the Edition field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TemplateProperties) GetEdition() *string {
	if o == nil {
		return nil
	}

	return o.Edition

}

// GetEditionOk returns a tuple with the Edition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProperties) GetEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Edition, true
}

// SetEdition sets field value
func (o *TemplateProperties) SetEdition(v string) {

	o.Edition = &v

}

// HasEdition returns a boolean if a field has been set.
func (o *TemplateProperties) HasEdition() bool {
	if o != nil && o.Edition != nil {
		return true
	}

	return false
}

// GetCores returns the Cores field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TemplateProperties) GetCores() *int32 {
	if o == nil {
		return nil
	}

	return o.Cores

}

// GetCoresOk returns a tuple with the Cores field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProperties) GetCoresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Cores, true
}

// SetCores sets field value
func (o *TemplateProperties) SetCores(v int32) {

	o.Cores = &v

}

// HasCores returns a boolean if a field has been set.
func (o *TemplateProperties) HasCores() bool {
	if o != nil && o.Cores != nil {
		return true
	}

	return false
}

// GetRam returns the Ram field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TemplateProperties) GetRam() *int32 {
	if o == nil {
		return nil
	}

	return o.Ram

}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProperties) GetRamOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Ram, true
}

// SetRam sets field value
func (o *TemplateProperties) SetRam(v int32) {

	o.Ram = &v

}

// HasRam returns a boolean if a field has been set.
func (o *TemplateProperties) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// GetStorageSize returns the StorageSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TemplateProperties) GetStorageSize() *int32 {
	if o == nil {
		return nil
	}

	return o.StorageSize

}

// GetStorageSizeOk returns a tuple with the StorageSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProperties) GetStorageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.StorageSize, true
}

// SetStorageSize sets field value
func (o *TemplateProperties) SetStorageSize(v int32) {

	o.StorageSize = &v

}

// HasStorageSize returns a boolean if a field has been set.
func (o *TemplateProperties) HasStorageSize() bool {
	if o != nil && o.StorageSize != nil {
		return true
	}

	return false
}

func (o TemplateProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Edition != nil {
		toSerialize["edition"] = o.Edition
	}

	if o.Cores != nil {
		toSerialize["cores"] = o.Cores
	}

	if o.Ram != nil {
		toSerialize["ram"] = o.Ram
	}

	if o.StorageSize != nil {
		toSerialize["storageSize"] = o.StorageSize
	}

	return json.Marshal(toSerialize)
}

type NullableTemplateProperties struct {
	value *TemplateProperties
	isSet bool
}

func (v NullableTemplateProperties) Get() *TemplateProperties {
	return v.value
}

func (v *NullableTemplateProperties) Set(val *TemplateProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateProperties(val *TemplateProperties) *NullableTemplateProperties {
	return &NullableTemplateProperties{value: val, isSet: true}
}

func (v NullableTemplateProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
