# \FirmwareUpgradesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkFirmwareUpgradesRollback**](FirmwareUpgradesApi.md#CreateNetworkFirmwareUpgradesRollback) | **Post** /networks/{networkId}/firmwareUpgrades/rollbacks | Rollback a Firmware Upgrade For A Network
[**GetNetworkFirmwareUpgrades**](FirmwareUpgradesApi.md#GetNetworkFirmwareUpgrades) | **Get** /networks/{networkId}/firmwareUpgrades | Get firmware upgrade information for a network
[**UpdateNetworkFirmwareUpgrades**](FirmwareUpgradesApi.md#UpdateNetworkFirmwareUpgrades) | **Put** /networks/{networkId}/firmwareUpgrades | Update firmware upgrade information for a network



## CreateNetworkFirmwareUpgradesRollback

> CreateNetworkFirmwareUpgradesRollback200Response CreateNetworkFirmwareUpgradesRollback(ctx, networkId).CreateNetworkFirmwareUpgradesRollback(createNetworkFirmwareUpgradesRollback).Execute()

Rollback a Firmware Upgrade For A Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 
    createNetworkFirmwareUpgradesRollback := *openapiclient.NewCreateNetworkFirmwareUpgradesRollbackRequest([]openapiclient.CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner{*openapiclient.NewCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner("Category_example", "Comment_example")}) // CreateNetworkFirmwareUpgradesRollbackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesRollback(context.Background(), networkId).CreateNetworkFirmwareUpgradesRollback(createNetworkFirmwareUpgradesRollback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesRollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesRollback`: CreateNetworkFirmwareUpgradesRollback200Response
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesRollback** | [**CreateNetworkFirmwareUpgradesRollbackRequest**](CreateNetworkFirmwareUpgradesRollbackRequest.md) |  | 

### Return type

[**CreateNetworkFirmwareUpgradesRollback200Response**](CreateNetworkFirmwareUpgradesRollback200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgrades

> GetNetworkFirmwareUpgrades200Response GetNetworkFirmwareUpgrades(ctx, networkId).Execute()

Get firmware upgrade information for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.GetNetworkFirmwareUpgrades(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.GetNetworkFirmwareUpgrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgrades`: GetNetworkFirmwareUpgrades200Response
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.GetNetworkFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFirmwareUpgrades200Response**](GetNetworkFirmwareUpgrades200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgrades

> GetNetworkFirmwareUpgrades200Response UpdateNetworkFirmwareUpgrades(ctx, networkId).UpdateNetworkFirmwareUpgrades(updateNetworkFirmwareUpgrades).Execute()

Update firmware upgrade information for a network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkId := "networkId_example" // string | 
    updateNetworkFirmwareUpgrades := *openapiclient.NewUpdateNetworkFirmwareUpgradesRequest() // UpdateNetworkFirmwareUpgradesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.UpdateNetworkFirmwareUpgrades(context.Background(), networkId).UpdateNetworkFirmwareUpgrades(updateNetworkFirmwareUpgrades).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgrades`: GetNetworkFirmwareUpgrades200Response
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgrades** | [**UpdateNetworkFirmwareUpgradesRequest**](UpdateNetworkFirmwareUpgradesRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgrades200Response**](GetNetworkFirmwareUpgrades200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

