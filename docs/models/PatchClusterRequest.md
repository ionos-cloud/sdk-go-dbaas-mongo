# PatchClusterRequest

Request payload to change a cluster.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |
|**Properties** | Pointer to [**PatchClusterProperties**](PatchClusterProperties.md) |  | [optional] |

## Methods


### GetMetadata

`func (o *PatchClusterRequest) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchClusterRequest) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchClusterRequest) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchClusterRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *PatchClusterRequest) GetProperties() PatchClusterProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PatchClusterRequest) GetPropertiesOk() (*PatchClusterProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PatchClusterRequest) SetProperties(v PatchClusterProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PatchClusterRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



