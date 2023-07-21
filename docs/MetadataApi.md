# \MetadataApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfosVersionGet**](MetadataApi.md#InfosVersionGet) | **Get** /infos/version | Get API Version
[**InfosVersionsGet**](MetadataApi.md#InfosVersionsGet) | **Get** /infos/versions | Get All API Versions



## InfosVersionGet

> APIVersion InfosVersionGet(ctx).Execute()

Get API Version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.InfosVersionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.InfosVersionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InfosVersionGet`: APIVersion
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.InfosVersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInfosVersionGetRequest struct via the builder pattern


### Return type

[**APIVersion**](APIVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfosVersionsGet

> []APIVersion InfosVersionsGet(ctx).Execute()

Get All API Versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.InfosVersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.InfosVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InfosVersionsGet`: []APIVersion
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.InfosVersionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInfosVersionsGetRequest struct via the builder pattern


### Return type

[**[]APIVersion**](APIVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

