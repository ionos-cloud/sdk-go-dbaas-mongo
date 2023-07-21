# \TemplatesApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplatesGet**](TemplatesApi.md#TemplatesGet) | **Get** /templates | Get Templates



## TemplatesGet

> TemplateList TemplatesGet(ctx).Limit(limit).Offset(offset).Execute()

Get Templates



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesGet(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGet`: TemplateList
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]
 **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]

### Return type

[**TemplateList**](TemplateList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

