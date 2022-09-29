# SnapshotListAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] |
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Items** | Pointer to [**[]SnapshotResponse**](SnapshotResponse.md) |  | [optional] |

## Methods


### GetType

`func (o *SnapshotListAllOf) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotListAllOf) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotListAllOf) SetType(v ResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *SnapshotListAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *SnapshotListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotListAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotListAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *SnapshotListAllOf) GetItems() []SnapshotResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SnapshotListAllOf) GetItemsOk() (*[]SnapshotResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SnapshotListAllOf) SetItems(v []SnapshotResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *SnapshotListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.



