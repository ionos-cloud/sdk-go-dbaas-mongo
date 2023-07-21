# \ClustersApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersDelete**](ClustersApi.md#ClustersDelete) | **Delete** /clusters/{clusterId} | Delete a Cluster
[**ClustersFindById**](ClustersApi.md#ClustersFindById) | **Get** /clusters/{clusterId} | Get a cluster by id
[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | Get Clusters
[**ClustersPatch**](ClustersApi.md#ClustersPatch) | **Patch** /clusters/{clusterId} | Patch a cluster
[**ClustersPost**](ClustersApi.md#ClustersPost) | **Post** /clusters | Create a Cluster



## ClustersDelete

> ClusterResponse ClustersDelete(ctx, clusterId).Execute()

Delete a Cluster



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ClustersDelete(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersDelete`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The unique ID of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersFindById

> ClusterResponse ClustersFindById(ctx, clusterId).Execute()

Get a cluster by id



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ClustersFindById(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersFindById`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersFindById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The unique ID of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersFindByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersGet

> ClusterList ClustersGet(ctx).Limit(limit).Offset(offset).FilterName(filterName).Execute()

Get Clusters



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
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with 'offset' for pagination. (optional) (default to 100)
    offset := int32(200) // int32 | The first element to return. Use together with 'limit' for pagination. (optional) (default to 0)
    filterName := "filterName_example" // string | Response filter to list only the MongoDB clusters that contain the specified name. The value is case insensitive and matched on the 'displayName' field.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ClustersGet(context.Background()).Limit(limit).Offset(offset).FilterName(filterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersGet`: ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClustersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]
 **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]
 **filterName** | **string** | Response filter to list only the MongoDB clusters that contain the specified name. The value is case insensitive and matched on the &#39;displayName&#39; field.  | 

### Return type

[**ClusterList**](ClusterList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersPatch

> ClusterResponse ClustersPatch(ctx, clusterId).PatchClusterRequest(patchClusterRequest).Execute()

Patch a cluster



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
    patchClusterRequest := *openapiclient.NewPatchClusterRequest() // PatchClusterRequest | Part of the cluster which should be modified.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ClustersPatch(context.Background(), clusterId).PatchClusterRequest(patchClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersPatch`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The unique ID of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchClusterRequest** | [**PatchClusterRequest**](PatchClusterRequest.md) | Part of the cluster which should be modified. | 

### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersPost

> ClusterResponse ClustersPost(ctx).CreateClusterRequest(createClusterRequest).Execute()

Create a Cluster



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
    createClusterRequest := *openapiclient.NewCreateClusterRequest() // CreateClusterRequest | The cluster to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ClustersPost(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersPost`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClustersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClusterRequest** | [**CreateClusterRequest**](CreateClusterRequest.md) | The cluster to be created. | 

### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

