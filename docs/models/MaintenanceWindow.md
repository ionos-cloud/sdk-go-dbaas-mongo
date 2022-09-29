# MaintenanceWindow

A weekly window of 4 hours during which maintenance work can be performed. 


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Time** | **string** |  | |
|**DayOfTheWeek** | [**DayOfTheWeek**](DayOfTheWeek.md) |  | |

## Methods


### GetTime

`func (o *MaintenanceWindow) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MaintenanceWindow) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MaintenanceWindow) SetTime(v string)`

SetTime sets Time field to given value.


### GetDayOfTheWeek

`func (o *MaintenanceWindow) GetDayOfTheWeek() DayOfTheWeek`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *MaintenanceWindow) GetDayOfTheWeekOk() (*DayOfTheWeek, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *MaintenanceWindow) SetDayOfTheWeek(v DayOfTheWeek)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.




