# ErrorResponse



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | Pointer to **int32** | The HTTP status code of the operation. | [optional] |
|**Messages** | Pointer to [**[]ErrorMessage**](ErrorMessage.md) |  | [optional] |

## Methods


### GetHttpStatus

`func (o *ErrorResponse) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ErrorResponse) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ErrorResponse) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *ErrorResponse) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessages

`func (o *ErrorResponse) GetMessages() []ErrorMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ErrorResponse) GetMessagesOk() (*[]ErrorMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ErrorResponse) SetMessages(v []ErrorMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ErrorResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.



