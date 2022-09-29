# UserRoles

a list of mongodb user role.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Role** | Pointer to **string** |  | [optional] |
|**Database** | Pointer to **string** |  | [optional] |

## Methods


### GetRole

`func (o *UserRoles) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserRoles) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserRoles) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserRoles) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetDatabase

`func (o *UserRoles) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *UserRoles) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *UserRoles) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *UserRoles) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.



