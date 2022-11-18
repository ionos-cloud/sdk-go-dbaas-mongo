# PatchUserProperties

MongoDB database user patch request properties.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Password** | Pointer to **string** |  | [optional] |
|**Roles** | Pointer to [**[]UserRoles**](UserRoles.md) |  | [optional] |

## Methods


### GetPassword

`func (o *PatchUserProperties) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchUserProperties) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchUserProperties) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchUserProperties) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoles

`func (o *PatchUserProperties) GetRoles() []UserRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PatchUserProperties) GetRolesOk() (*[]UserRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PatchUserProperties) SetRoles(v []UserRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PatchUserProperties) HasRoles() bool`

HasRoles returns a boolean if a field has been set.



