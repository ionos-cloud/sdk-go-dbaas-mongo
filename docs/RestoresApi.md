# \RestoresApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersRestorePost**](RestoresApi.md#ClustersRestorePost) | **Post** /clusters/{clusterId}/restore | In-place restore of a cluster



## ClustersRestorePost

> ClustersRestorePost(ctx, clusterId).CreateRestoreRequest(createRestoreRequest).Execute()

In-place restore of a cluster



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
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    createRestoreRequest := *openapiclient.NewCreateRestoreRequest("dcd31531-3ac8-11eb-9feb-046c59cc737e") // CreateRestoreRequest | The restore request to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RestoresApi.ClustersRestorePost(context.Background(), clusterId).CreateRestoreRequest(createRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoresApi.ClustersRestorePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The unique ID of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersRestorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRestoreRequest** | [**CreateRestoreRequest**](CreateRestoreRequest.md) | The restore request to create. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

