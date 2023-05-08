# \ClientProfilesApi

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientProfile**](ClientProfilesApi.md#CreateClientProfile) | **Post** /api/v2/missionControl/eventBrokerServices/{serviceId}/clientProfiles | (Beta) Create a client profile
[**DeleteClientProfile**](ClientProfilesApi.md#DeleteClientProfile) | **Delete** /api/v2/missionControl/eventBrokerServices/{serviceId}/clientProfiles/{name} | (Beta) Delete a client profile
[**GetClientProfile**](ClientProfilesApi.md#GetClientProfile) | **Get** /api/v2/missionControl/eventBrokerServices/{serviceId}/clientProfiles/{name} | (Beta) Get a specific client profile for event broker service by name
[**GetClientProfiles**](ClientProfilesApi.md#GetClientProfiles) | **Get** /api/v2/missionControl/eventBrokerServices/{serviceId}/clientProfiles | (Beta) Get a list of client profiles
[**ReplaceClientProfile**](ClientProfilesApi.md#ReplaceClientProfile) | **Put** /api/v2/missionControl/eventBrokerServices/{serviceId}/clientProfiles/{name} | (Beta) Replace a client profile
[**UpdateClientProfile**](ClientProfilesApi.md#UpdateClientProfile) | **Patch** /api/v2/missionControl/eventBrokerServices/{serviceId}/clientProfiles/{name} | (Beta) Update a client profile



## CreateClientProfile

> OperationResponse CreateClientProfile(ctx, serviceId).ClientProfileRequest(clientProfileRequest).Execute()

(Beta) Create a client profile



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
    serviceId := "serviceId_example" // string | The identifier of the event broker service.
    clientProfileRequest := *openapiclient.NewClientProfileRequest() // ClientProfileRequest | The new client profile specification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientProfilesApi.CreateClientProfile(context.Background(), serviceId).ClientProfileRequest(clientProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfilesApi.CreateClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClientProfile`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfilesApi.CreateClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The identifier of the event broker service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientProfileRequest** | [**ClientProfileRequest**](ClientProfileRequest.md) | The new client profile specification. | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientProfile

> OperationResponse DeleteClientProfile(ctx, serviceId, name).Execute()

(Beta) Delete a client profile



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
    serviceId := "serviceId_example" // string | The unique identifier of the event broker service.
    name := "name_example" // string | The name of the client profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientProfilesApi.DeleteClientProfile(context.Background(), serviceId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfilesApi.DeleteClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClientProfile`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfilesApi.DeleteClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the event broker service. | 
**name** | **string** | The name of the client profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientProfile

> ClientProfileResponse GetClientProfile(ctx, serviceId, name).Execute()

(Beta) Get a specific client profile for event broker service by name



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
    serviceId := "serviceId_example" // string | The identifier of the event broker service.
    name := "name_example" // string | The name of the client profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientProfilesApi.GetClientProfile(context.Background(), serviceId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfilesApi.GetClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientProfile`: ClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfilesApi.GetClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The identifier of the event broker service. | 
**name** | **string** | The name of the client profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientProfileResponse**](ClientProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientProfiles

> ClientProfilesResponse GetClientProfiles(ctx, serviceId).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()

(Beta) Get a list of client profiles



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
    serviceId := "serviceId_example" // string | The identifier of the event broker service.
    pageNumber := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of client profiles to return per page. (optional) (default to 100)
    sort := "name:asc" // string | The sorting criteria for the list of client profiles. (optional) (default to "name:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientProfilesApi.GetClientProfiles(context.Background(), serviceId).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfilesApi.GetClientProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientProfiles`: ClientProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfilesApi.GetClientProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The identifier of the event broker service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | The number of client profiles to return per page. | [default to 100]
 **sort** | **string** | The sorting criteria for the list of client profiles. | [default to &quot;name:asc&quot;]

### Return type

[**ClientProfilesResponse**](ClientProfilesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceClientProfile

> OperationResponse ReplaceClientProfile(ctx, serviceId, name).ClientProfileRequest(clientProfileRequest).Execute()

(Beta) Replace a client profile



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
    serviceId := "serviceId_example" // string | The unique identifier of the event broker service.
    name := "name_example" // string | 
    clientProfileRequest := *openapiclient.NewClientProfileRequest() // ClientProfileRequest | The fields to update the configuration for the client profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientProfilesApi.ReplaceClientProfile(context.Background(), serviceId, name).ClientProfileRequest(clientProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfilesApi.ReplaceClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceClientProfile`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfilesApi.ReplaceClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the event broker service. | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientProfileRequest** | [**ClientProfileRequest**](ClientProfileRequest.md) | The fields to update the configuration for the client profile. | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientProfile

> OperationResponse UpdateClientProfile(ctx, serviceId, name).ClientProfileRequest(clientProfileRequest).Execute()

(Beta) Update a client profile



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
    serviceId := "serviceId_example" // string | The unique identifier of the event broker service.
    name := "name_example" // string | The name of the client profile.
    clientProfileRequest := *openapiclient.NewClientProfileRequest() // ClientProfileRequest | The configuration to update for the client profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientProfilesApi.UpdateClientProfile(context.Background(), serviceId, name).ClientProfileRequest(clientProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfilesApi.UpdateClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientProfile`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfilesApi.UpdateClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the event broker service. | 
**name** | **string** | The name of the client profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientProfileRequest** | [**ClientProfileRequest**](ClientProfileRequest.md) | The configuration to update for the client profile. | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

