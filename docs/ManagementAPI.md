# \ManagementAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAgencyConfGlobalTravelTimeCoefficient**](ManagementAPI.md#DeleteAgencyConfGlobalTravelTimeCoefficient) | **Delete** /agencies/{agencyId}/conf/globalTravelTimeCoefficient | Resets the agency&#39;s configured global travel time coefficient
[**FetchLastNPlanStates**](ManagementAPI.md#FetchLastNPlanStates) | **Get** /agencies/{agencyId}/plans/{planId}/states | Fetch the latest states of a plan
[**FetchLastPlanState**](ManagementAPI.md#FetchLastPlanState) | **Get** /agencies/{agencyId}/plans/{planId}/state | Fetch the latest state of a plan
[**GetPlanStatus**](ManagementAPI.md#GetPlanStatus) | **Get** /agencies/{agencyId}/plans/{planId}/status | Retrieve a plan status
[**PostLoginRefresh**](ManagementAPI.md#PostLoginRefresh) | **Post** /login/refresh | Refresh the access token
[**PutAgencyConfGlobalTravelTimeCoefficient**](ManagementAPI.md#PutAgencyConfGlobalTravelTimeCoefficient) | **Put** /agencies/{agencyId}/conf/globalTravelTimeCoefficient | Updates the agency&#39;s configured global travel time coefficient configuration
[**PutPlanRunning**](ManagementAPI.md#PutPlanRunning) | **Put** /agencies/{agencyId}/plans/{planId}/running | Stop or restart the optimization of a plan



## DeleteAgencyConfGlobalTravelTimeCoefficient

> DeleteAgencyConfGlobalTravelTimeCoefficient(ctx, agencyId).Execute()

Resets the agency's configured global travel time coefficient

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteAgencyConfGlobalTravelTimeCoefficient(context.Background(), agencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteAgencyConfGlobalTravelTimeCoefficient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgencyConfGlobalTravelTimeCoefficientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## FetchLastNPlanStates

> EnvelopedTimedPlanStates FetchLastNPlanStates(ctx, agencyId, planId).Limit(limit).Execute()

Fetch the latest states of a plan

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
	limit := int32(5) // int32 | The number of items to fetch. The value should be a positive integer and if the limit is set to 0, all items will be fetched.  (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.FetchLastNPlanStates(context.Background(), agencyId, planId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.FetchLastNPlanStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchLastNPlanStates`: EnvelopedTimedPlanStates
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.FetchLastNPlanStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchLastNPlanStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The number of items to fetch. The value should be a positive integer and if the limit is set to 0, all items will be fetched.  | [default to 5]

### Return type

[**EnvelopedTimedPlanStates**](EnvelopedTimedPlanStates.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchLastPlanState

> EnvelopedTimedPlanState FetchLastPlanState(ctx, agencyId, planId).Execute()

Fetch the latest state of a plan

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.FetchLastPlanState(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.FetchLastPlanState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchLastPlanState`: EnvelopedTimedPlanState
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.FetchLastPlanState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchLastPlanStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvelopedTimedPlanState**](EnvelopedTimedPlanState.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanStatus

> EnvelopedPlanStatus GetPlanStatus(ctx, agencyId, planId).Execute()

Retrieve a plan status

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetPlanStatus(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetPlanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanStatus`: EnvelopedPlanStatus
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetPlanStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvelopedPlanStatus**](EnvelopedPlanStatus.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLoginRefresh

> PostLoginRefresh200Response PostLoginRefresh(ctx).PostLoginRefreshRequest(postLoginRefreshRequest).Execute()

Refresh the access token

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
	postLoginRefreshRequest := *openapiclient.NewPostLoginRefreshRequest() // PostLoginRefreshRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.PostLoginRefresh(context.Background()).PostLoginRefreshRequest(postLoginRefreshRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.PostLoginRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLoginRefresh`: PostLoginRefresh200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.PostLoginRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLoginRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postLoginRefreshRequest** | [**PostLoginRefreshRequest**](PostLoginRefreshRequest.md) |  | 

### Return type

[**PostLoginRefresh200Response**](PostLoginRefresh200Response.md)

### Authorization

[refresh_token](../README.md#refresh_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAgencyConfGlobalTravelTimeCoefficient

> EnvelopedGlobalTravelTimeCoefficient PutAgencyConfGlobalTravelTimeCoefficient(ctx, agencyId).Body(body).Execute()

Updates the agency's configured global travel time coefficient configuration

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
	body := float32(8.14) // float32 | The floating value for the coefficient. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.PutAgencyConfGlobalTravelTimeCoefficient(context.Background(), agencyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.PutAgencyConfGlobalTravelTimeCoefficient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAgencyConfGlobalTravelTimeCoefficient`: EnvelopedGlobalTravelTimeCoefficient
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.PutAgencyConfGlobalTravelTimeCoefficient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAgencyConfGlobalTravelTimeCoefficientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **float32** | The floating value for the coefficient. | 

### Return type

[**EnvelopedGlobalTravelTimeCoefficient**](EnvelopedGlobalTravelTimeCoefficient.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlanRunning

> PutPlanRunning(ctx, agencyId, planId).Force(force).Body(body).Execute()

Stop or restart the optimization of a plan

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
	force := true // bool | If true, on an archived item, the requested action will be forced and the item will be unarchived. (optional) (default to false)
	body := true // bool | A boolean value: false will stop the optimization, true will restart the optimization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.PutPlanRunning(context.Background(), agencyId, planId).Force(force).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.PutPlanRunning``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlanRunningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | If true, on an archived item, the requested action will be forced and the item will be unarchived. | [default to false]
 **body** | **bool** | A boolean value: false will stop the optimization, true will restart the optimization. | 

### Return type

 (empty response body)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

