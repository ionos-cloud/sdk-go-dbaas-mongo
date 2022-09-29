# ClusterLogs

The logs of the MongoDB cluster.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Instances** | Pointer to [**[]ClusterLogsInstances**](ClusterLogsInstances.md) |  | [optional] |

## Methods


### GetInstances

`func (o *ClusterLogs) GetInstances() []ClusterLogsInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ClusterLogs) GetInstancesOk() (*[]ClusterLogsInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ClusterLogs) SetInstances(v []ClusterLogsInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ClusterLogs) HasInstances() bool`

HasInstances returns a boolean if a field has been set.



