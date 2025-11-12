# \ResourceAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlanResource**](ResourceAPI.md#DeletePlanResource) | **Delete** /agencies/{agencyId}/plans/{planId}/resources/{resourceId} | Delete a plan&#39;s resource
[**GetPlanResource**](ResourceAPI.md#GetPlanResource) | **Get** /agencies/{agencyId}/plans/{planId}/resources/{resourceId} | Retrieve a plan&#39;s resource
[**GetPlanResourceState**](ResourceAPI.md#GetPlanResourceState) | **Get** /agencies/{agencyId}/plans/{planId}/resources/{resourceId}/state | Retrieve a plan&#39;s resource state
[**PutForbidResourceStop**](ResourceAPI.md#PutForbidResourceStop) | **Put** /agencies/{agencyId}/plans/{planId}/resources/{resourceId}/forbid/{stopId} | Forbid a resource from doing a stop, and return the updated plan
[**PutPlanResource**](ResourceAPI.md#PutPlanResource) | **Put** /agencies/{agencyId}/plans/{planId}/resources/{resourceId} | Create or update a plan&#39;s resource
[**PutPlanResourceState**](ResourceAPI.md#PutPlanResourceState) | **Put** /agencies/{agencyId}/plans/{planId}/resources/{resourceId}/state | Update a plan&#39;s resource state



## DeletePlanResource

> DeletePlanResource(ctx, agencyId, planId, resourceId).Force(force).Execute()

Delete a plan's resource

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
	agencyId := TODO // AgencyId | The agency id.
	planId := "planId_example" // string | The plan id.
	resourceId := "resourceId_example" // string | The resource id.
	force := true // bool | If true, on an archived item, the requested action will be forced and the item will be unarchived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceAPI.DeletePlanResource(context.Background(), agencyId, planId, resourceId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceAPI.DeletePlanResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**resourceId** | **string** | The resource id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | If true, on an archived item, the requested action will be forced and the item will be unarchived. | [default to false]

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanResource

> EnvelopedResource GetPlanResource(ctx, agencyId, planId, resourceId).Execute()

Retrieve a plan's resource

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
	agencyId := TODO // AgencyId | The agency id.
	planId := "planId_example" // string | The plan id.
	resourceId := "resourceId_example" // string | The resource id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceAPI.GetPlanResource(context.Background(), agencyId, planId, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceAPI.GetPlanResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanResource`: EnvelopedResource
	fmt.Fprintf(os.Stdout, "Response from `ResourceAPI.GetPlanResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**resourceId** | **string** | The resource id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnvelopedResource**](EnvelopedResource.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanResourceState

> EnvelopedState GetPlanResourceState(ctx, agencyId, planId, resourceId).Execute()

Retrieve a plan's resource state

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
	agencyId := TODO // AgencyId | The agency id.
	planId := "planId_example" // string | The plan id.
	resourceId := "resourceId_example" // string | The resource id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceAPI.GetPlanResourceState(context.Background(), agencyId, planId, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceAPI.GetPlanResourceState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanResourceState`: EnvelopedState
	fmt.Fprintf(os.Stdout, "Response from `ResourceAPI.GetPlanResourceState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**resourceId** | **string** | The resource id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanResourceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnvelopedState**](EnvelopedState.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutForbidResourceStop

> EnvelopedPlan PutForbidResourceStop(ctx, agencyId, planId, resourceId, stopId).Force(force).Execute()

Forbid a resource from doing a stop, and return the updated plan



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
	agencyId := TODO // AgencyId | The agency id.
	planId := "planId_example" // string | The plan id.
	resourceId := "resourceId_example" // string | The resource id.
	stopId := "stopId_example" // string | The stop id.
	force := true // bool | If true, on an archived item, the requested action will be forced and the item will be unarchived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceAPI.PutForbidResourceStop(context.Background(), agencyId, planId, resourceId, stopId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceAPI.PutForbidResourceStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutForbidResourceStop`: EnvelopedPlan
	fmt.Fprintf(os.Stdout, "Response from `ResourceAPI.PutForbidResourceStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**resourceId** | **string** | The resource id. | 
**stopId** | **string** | The stop id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutForbidResourceStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **force** | **bool** | If true, on an archived item, the requested action will be forced and the item will be unarchived. | [default to false]

### Return type

[**EnvelopedPlan**](EnvelopedPlan.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlanResource

> EnvelopedResource PutPlanResource(ctx, agencyId, planId, resourceId).Force(force).Resource(resource).Execute()

Create or update a plan's resource

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
	agencyId := TODO // AgencyId | The agency id.
	planId := "planId_example" // string | The plan id.
	resourceId := "resourceId_example" // string | The resource id.
	force := true // bool | If true, on an archived item, the requested action will be forced and the item will be unarchived. (optional) (default to false)
	resource := *openapiclient.NewResource("Id_example", openapiclient.Resource_vehicleProfile{VehicleProfileBicycle: openapiclient.NewVehicleProfileBicycle("Type_example")}, *openapiclient.NewTimeWindow("2019-11-15T12:34:56Z", "2019-11-15T12:34:56Z")) // Resource | The Resource to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceAPI.PutPlanResource(context.Background(), agencyId, planId, resourceId).Force(force).Resource(resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceAPI.PutPlanResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlanResource`: EnvelopedResource
	fmt.Fprintf(os.Stdout, "Response from `ResourceAPI.PutPlanResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**resourceId** | **string** | The resource id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlanResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | If true, on an archived item, the requested action will be forced and the item will be unarchived. | [default to false]
 **resource** | [**Resource**](Resource.md) | The Resource to update. | 

### Return type

[**EnvelopedResource**](EnvelopedResource.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlanResourceState

> EnvelopedState PutPlanResourceState(ctx, agencyId, planId, resourceId).Force(force).State(state).Execute()

Update a plan's resource state

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
	agencyId := TODO // AgencyId | The agency id.
	planId := "planId_example" // string | The plan id.
	resourceId := "resourceId_example" // string | The resource id.
	force := true // bool | If true, on an archived item, the requested action will be forced and the item will be unarchived. (optional) (default to false)
	state := *openapiclient.NewState() // State | The State to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceAPI.PutPlanResourceState(context.Background(), agencyId, planId, resourceId).Force(force).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceAPI.PutPlanResourceState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlanResourceState`: EnvelopedState
	fmt.Fprintf(os.Stdout, "Response from `ResourceAPI.PutPlanResourceState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**resourceId** | **string** | The resource id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlanResourceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | If true, on an archived item, the requested action will be forced and the item will be unarchived. | [default to false]
 **state** | [**State**](State.md) | The State to update. | 

### Return type

[**EnvelopedState**](EnvelopedState.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

