# BackupProperties

Backup related properties. 


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Location** | Pointer to **string** | The location where the cluster backups will be stored. If not set, the backup is stored in the nearest location of the cluster.  | [optional] |

## Methods


### GetLocation

`func (o *BackupProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BackupProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BackupProperties) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BackupProperties) HasLocation() bool`

HasLocation returns a boolean if a field has been set.



