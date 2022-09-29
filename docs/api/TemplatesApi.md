# \TemplatesApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**TemplatesGet**](TemplatesApi.md#TemplatesGet) | **Get** /templates | Get Templates|



## TemplatesGet

```go
var result TemplateList = TemplatesGet(ctx)
                      .Execute()
```

Get Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-mongo"
)

func main() {

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.TemplatesApi.TemplatesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TemplatesGet`: TemplateList
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiTemplatesGetRequest struct via the builder pattern


### Return type

[**TemplateList**](TemplateList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


