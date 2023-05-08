# \EventBrokerServicesApi

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](EventBrokerServicesApi.md#CreateService) | **Post** /api/v2/missionControl/eventBrokerServices | (Beta) Create an event broker service
[**DeleteService**](EventBrokerServicesApi.md#DeleteService) | **Delete** /api/v2/missionControl/eventBrokerServices/{id} | (Beta) Delete an event broker service
[**GetDatacenter**](EventBrokerServicesApi.md#GetDatacenter) | **Get** /api/v2/missionControl/datacenters/{id} | (Beta) Get a specific datacenter by identifier
[**GetDatacenters**](EventBrokerServicesApi.md#GetDatacenters) | **Get** /api/v2/missionControl/datacenters | (Beta) Get a list of datacenters
[**GetEventBrokerServiceVersions**](EventBrokerServicesApi.md#GetEventBrokerServiceVersions) | **Get** /api/v2/missionControl/datacenters/{id}/eventBrokerServiceVersions | (Beta) Get Event Broker Service Versions.
[**GetService**](EventBrokerServicesApi.md#GetService) | **Get** /api/v2/missionControl/eventBrokerServices/{id} | (Beta) Get an event broker service
[**GetServiceClass**](EventBrokerServicesApi.md#GetServiceClass) | **Get** /api/v2/missionControl/serviceClasses/{id} | (Beta) Get a specific service class by identifier
[**GetServiceClasses**](EventBrokerServicesApi.md#GetServiceClasses) | **Get** /api/v2/missionControl/serviceClasses | (Beta) Get a list of service classes
[**GetServiceOperation**](EventBrokerServicesApi.md#GetServiceOperation) | **Get** /api/v2/missionControl/eventBrokerServices/{serviceId}/operations/{operationId} | (Beta) Get the status of a service operation
[**GetServices**](EventBrokerServicesApi.md#GetServices) | **Get** /api/v2/missionControl/eventBrokerServices | (Beta) Get a list of event broker services
[**UpdateService**](EventBrokerServicesApi.md#UpdateService) | **Patch** /api/v2/missionControl/eventBrokerServices/{id} | (Beta) Update an event broker service



## CreateService

> OperationResponse CreateService(ctx).CreateServiceRequest(createServiceRequest).Execute()

(Beta) Create an event broker service



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
    createServiceRequest := *openapiclient.NewCreateServiceRequest("my-service-name", "DEVELOPER", "k8s-us-east-1") // CreateServiceRequest | The new service specification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.CreateService(context.Background()).CreateServiceRequest(createServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.CreateService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceRequest** | [**CreateServiceRequest**](CreateServiceRequest.md) | The new service specification. | 

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


## DeleteService

> OperationResponse DeleteService(ctx, id).Execute()

(Beta) Delete an event broker service



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
    id := "id_example" // string | The unique identifier of the service.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.DeleteService(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteService`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.DeleteService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


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


## GetDatacenter

> DatacenterResponse GetDatacenter(ctx, id).Execute()

(Beta) Get a specific datacenter by identifier



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
    id := "id_example" // string | The unique identifier of the datacenter.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.GetDatacenter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.GetDatacenter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatacenter`: DatacenterResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.GetDatacenter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the datacenter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatacenterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatacenterResponse**](DatacenterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatacenters

> DatacentersResponse GetDatacenters(ctx).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).DatacenterType(datacenterType).Provider(provider).Execute()

(Beta) Get a list of datacenters



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
    pageNumber := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of datacenters to return per page. (optional) (default to 100)
    sort := "name:asc,datacenterType:asc" // string | The sorting criteria for the list of datacenters. (optional) (default to "name:asc,datacenterType:asc")
    datacenterType := "datacenterType_example" // string | The datacenter type to filter datacenters list. (optional)
    provider := "provider_example" // string | The provider name to filter datacenters list. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.GetDatacenters(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).DatacenterType(datacenterType).Provider(provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.GetDatacenters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatacenters`: DatacentersResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.GetDatacenters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatacentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | The number of datacenters to return per page. | [default to 100]
 **sort** | **string** | The sorting criteria for the list of datacenters. | [default to &quot;name:asc,datacenterType:asc&quot;]
 **datacenterType** | **string** | The datacenter type to filter datacenters list. | 
 **provider** | **string** | The provider name to filter datacenters list. | 

### Return type

[**DatacentersResponse**](DatacentersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventBrokerServiceVersions

> EventBrokerServiceVersions GetEventBrokerServiceVersions(ctx, id).Execute()

(Beta) Get Event Broker Service Versions.



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.GetEventBrokerServiceVersions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.GetEventBrokerServiceVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventBrokerServiceVersions`: EventBrokerServiceVersions
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.GetEventBrokerServiceVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventBrokerServiceVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventBrokerServiceVersions**](EventBrokerServiceVersions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> GetServiceResponse GetService(ctx, id).Expand(expand).Execute()

(Beta) Get an event broker service



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
    id := "id_example" // string | The unique identifier of the service.
    expand := []string{"Expand_example"} // []string | Additional information to retrieve for event broker service, such as connection endpoint information or broker details. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.GetService(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.GetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetService`: GetServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Additional information to retrieve for event broker service, such as connection endpoint information or broker details. | 

### Return type

[**GetServiceResponse**](GetServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceClass

> ServiceClassResponse GetServiceClass(ctx, id).Execute()

(Beta) Get a specific service class by identifier



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
    id := "id_example" // string | The unique identifier of the service class.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.GetServiceClass(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.GetServiceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceClass`: ServiceClassResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.GetServiceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the service class. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceClassResponse**](ServiceClassResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceClasses

> ServiceClassesResponse GetServiceClasses(ctx).Execute()

(Beta) Get a list of service classes



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
    resp, r, err := apiClient.EventBrokerServicesApi.GetServiceClasses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.GetServiceClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceClasses`: ServiceClassesResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.GetServiceClasses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceClassesRequest struct via the builder pattern


### Return type

[**ServiceClassesResponse**](ServiceClassesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceOperation

> OperationResponse GetServiceOperation(ctx, serviceId, operationId).Execute()

(Beta) Get the status of a service operation



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
    serviceId := "serviceId_example" // string | The unique identifier of the service.
    operationId := "operationId_example" // string | The identifier of the operation being performed on the event broker service.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.GetServiceOperation(context.Background(), serviceId, operationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.GetServiceOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceOperation`: OperationResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.GetServiceOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The unique identifier of the service. | 
**operationId** | **string** | The identifier of the operation being performed on the event broker service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceOperationRequest struct via the builder pattern


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


## GetServices

> GetAllServicesResponse GetServices(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

(Beta) Get a list of event broker services



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
    pageNumber := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of event brokers to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.GetServices(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.GetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServices`: GetAllServicesResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | The number of event brokers to return per page. | [default to 100]

### Return type

[**GetAllServicesResponse**](GetAllServicesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> GetServiceResponse UpdateService(ctx, id).UpdateServiceRequest(updateServiceRequest).Execute()

(Beta) Update an event broker service



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
    id := "id_example" // string | The unique identifier of the event broker service.
    updateServiceRequest := *openapiclient.NewUpdateServiceRequest() // UpdateServiceRequest | The fields to update the configuration for the event broker service.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventBrokerServicesApi.UpdateService(context.Background(), id).UpdateServiceRequest(updateServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventBrokerServicesApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: GetServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `EventBrokerServicesApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the event broker service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServiceRequest** | [**UpdateServiceRequest**](UpdateServiceRequest.md) | The fields to update the configuration for the event broker service. | 

### Return type

[**GetServiceResponse**](GetServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

