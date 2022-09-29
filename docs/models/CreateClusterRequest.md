# CreateClusterRequest

The request payload with all data needed to create a new MongoDB cluster. 


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |
|**Properties** | Pointer to [**CreateClusterProperties**](CreateClusterProperties.md) |  | [optional] |

## Methods


### GetMetadata

`func (o *CreateClusterRequest) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateClusterRequest) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateClusterRequest) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateClusterRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *CreateClusterRequest) GetProperties() CreateClusterProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateClusterRequest) GetPropertiesOk() (*CreateClusterProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateClusterRequest) SetProperties(v CreateClusterProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateClusterRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



