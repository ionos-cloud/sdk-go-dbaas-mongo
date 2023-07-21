# \SnapshotsApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersSnapshotsGet**](SnapshotsApi.md#ClustersSnapshotsGet) | **Get** /clusters/{clusterId}/snapshots | Get the snapshots of your cluster



## ClustersSnapshotsGet

> SnapshotList ClustersSnapshotsGet(ctx, clusterId).Limit(limit).Offset(offset).Execute()

Get the snapshots of your cluster



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
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with 'offset' for pagination. (optional) (default to 100)
    offset := int32(200) // int32 | The first element to return. Use together with 'limit' for pagination. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.ClustersSnapshotsGet(context.Background(), clusterId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.ClustersSnapshotsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersSnapshotsGet`: SnapshotList
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.ClustersSnapshotsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The unique ID of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersSnapshotsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]
 **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]

### Return type

[**SnapshotList**](SnapshotList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

