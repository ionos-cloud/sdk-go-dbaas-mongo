# PatchUserRequest

MongoDB database user patch request.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**UserMetadata**](UserMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**PatchUserProperties**](PatchUserProperties.md) |  | [optional] |

## Methods


### GetMetadata

`func (o *PatchUserRequest) GetMetadata() UserMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchUserRequest) GetMetadataOk() (*UserMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchUserRequest) SetMetadata(v UserMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchUserRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *PatchUserRequest) GetProperties() PatchUserProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PatchUserRequest) GetPropertiesOk() (*PatchUserProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PatchUserRequest) SetProperties(v PatchUserProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PatchUserRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



