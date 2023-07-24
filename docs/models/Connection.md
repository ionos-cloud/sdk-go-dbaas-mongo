# Connection

The network connection  details for your cluster.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The datacenter to which your cluster will be connected. | |
|**LanId** | **string** | The numeric LAN ID with which you connect your cluster. | |
|**CidrList** | **[]string** | The list of IPs for your cluster. All IPs must be in a /24 network. Note the following unavailable IP ranges: 10.233.114.0/24  | |
|**Whitelist** | Pointer to **[]string** | List of whitelisted CIDRs.  | [optional] |

## Methods


### GetDatacenterId

`func (o *Connection) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *Connection) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *Connection) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLanId

`func (o *Connection) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *Connection) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *Connection) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetCidrList

`func (o *Connection) GetCidrList() []string`

GetCidrList returns the CidrList field if non-nil, zero value otherwise.

### GetCidrListOk

`func (o *Connection) GetCidrListOk() (*[]string, bool)`

GetCidrListOk returns a tuple with the CidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrList

`func (o *Connection) SetCidrList(v []string)`

SetCidrList sets CidrList field to given value.


### GetWhitelist

`func (o *Connection) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *Connection) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *Connection) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *Connection) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.



