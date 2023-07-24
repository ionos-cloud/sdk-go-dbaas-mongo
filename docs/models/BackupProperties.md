# BackupProperties

Backup related properties. 


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SnapshotIntervalHours** | Pointer to **int32** | Number of hours between snapshots.  | [optional] [default to 24]|
|**PointInTimeWindowHours** | Pointer to **int32** | Number of hours in the past for which a point-in-time snapshot can be created.  | [optional] |
|**BackupRetention** | Pointer to [**BackupRetentionProperties**](BackupRetentionProperties.md) |  | [optional] |
|**Location** | Pointer to **string** | The location where the cluster backups will be stored. If not set, the backup is stored in the nearest location of the cluster.  | [optional] |

## Methods


### GetSnapshotIntervalHours

`func (o *BackupProperties) GetSnapshotIntervalHours() int32`

GetSnapshotIntervalHours returns the SnapshotIntervalHours field if non-nil, zero value otherwise.

### GetSnapshotIntervalHoursOk

`func (o *BackupProperties) GetSnapshotIntervalHoursOk() (*int32, bool)`

GetSnapshotIntervalHoursOk returns a tuple with the SnapshotIntervalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIntervalHours

`func (o *BackupProperties) SetSnapshotIntervalHours(v int32)`

SetSnapshotIntervalHours sets SnapshotIntervalHours field to given value.

### HasSnapshotIntervalHours

`func (o *BackupProperties) HasSnapshotIntervalHours() bool`

HasSnapshotIntervalHours returns a boolean if a field has been set.

### GetPointInTimeWindowHours

`func (o *BackupProperties) GetPointInTimeWindowHours() int32`

GetPointInTimeWindowHours returns the PointInTimeWindowHours field if non-nil, zero value otherwise.

### GetPointInTimeWindowHoursOk

`func (o *BackupProperties) GetPointInTimeWindowHoursOk() (*int32, bool)`

GetPointInTimeWindowHoursOk returns a tuple with the PointInTimeWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeWindowHours

`func (o *BackupProperties) SetPointInTimeWindowHours(v int32)`

SetPointInTimeWindowHours sets PointInTimeWindowHours field to given value.

### HasPointInTimeWindowHours

`func (o *BackupProperties) HasPointInTimeWindowHours() bool`

HasPointInTimeWindowHours returns a boolean if a field has been set.

### GetBackupRetention

`func (o *BackupProperties) GetBackupRetention() BackupRetentionProperties`

GetBackupRetention returns the BackupRetention field if non-nil, zero value otherwise.

### GetBackupRetentionOk

`func (o *BackupProperties) GetBackupRetentionOk() (*BackupRetentionProperties, bool)`

GetBackupRetentionOk returns a tuple with the BackupRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetention

`func (o *BackupProperties) SetBackupRetention(v BackupRetentionProperties)`

SetBackupRetention sets BackupRetention field to given value.

### HasBackupRetention

`func (o *BackupProperties) HasBackupRetention() bool`

HasBackupRetention returns a boolean if a field has been set.

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



