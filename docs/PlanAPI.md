# \PlanAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckPlanManualMode**](PlanAPI.md#CheckPlanManualMode) | **Get** /agencies/{agencyId}/plans/{planId}/manual/check | Check if a plan is in manual mode
[**DeletePlan**](PlanAPI.md#DeletePlan) | **Delete** /agencies/{agencyId}/plans/{planId} | Delete a plan
[**FetchLastNPlanStates**](PlanAPI.md#FetchLastNPlanStates) | **Get** /agencies/{agencyId}/plans/{planId}/states | Fetch the latest states of a plan
[**FetchLastPlanState**](PlanAPI.md#FetchLastPlanState) | **Get** /agencies/{agencyId}/plans/{planId}/state | Fetch the latest state of a plan
[**GetPlan**](PlanAPI.md#GetPlan) | **Get** /agencies/{agencyId}/plans/{planId} | Retrieve a plan
[**GetPlanSolution**](PlanAPI.md#GetPlanSolution) | **Get** /agencies/{agencyId}/plans/{planId}/solution | Retrieve a plan solution
[**GetPlanSolutionObjectives**](PlanAPI.md#GetPlanSolutionObjectives) | **Get** /agencies/{agencyId}/plans/{planId}/solution/objectives | Retrieve the objectives of a plan solution
[**GetPlanStatus**](PlanAPI.md#GetPlanStatus) | **Get** /agencies/{agencyId}/plans/{planId}/status | Retrieve a plan status
[**GetPlans**](PlanAPI.md#GetPlans) | **Get** /agencies/{agencyId}/plans | Retrieves a collection of plans
[**PutPlan**](PlanAPI.md#PutPlan) | **Put** /agencies/{agencyId}/plans/{planId} | Create or update a plan
[**PutPlanMode**](PlanAPI.md#PutPlanMode) | **Put** /agencies/{agencyId}/plans/{planId}/mode | Changes a plan&#39;s mode
[**PutPlanRunning**](PlanAPI.md#PutPlanRunning) | **Put** /agencies/{agencyId}/plans/{planId}/running | Stop or restart the optimization of a plan



## CheckPlanManualMode

> EnvelopedBoolResponse CheckPlanManualMode(ctx, agencyId, planId).Execute()

Check if a plan is in manual mode

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
	resp, r, err := apiClient.PlanAPI.CheckPlanManualMode(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.CheckPlanManualMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckPlanManualMode`: EnvelopedBoolResponse
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.CheckPlanManualMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckPlanManualModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvelopedBoolResponse**](EnvelopedBoolResponse.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlan

> DeletePlan(ctx, agencyId, planId).Execute()

Delete a plan

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
	r, err := apiClient.PlanAPI.DeletePlan(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.DeletePlan``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePlanRequest struct via the builder pattern


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
	resp, r, err := apiClient.PlanAPI.FetchLastNPlanStates(context.Background(), agencyId, planId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.FetchLastNPlanStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchLastNPlanStates`: EnvelopedTimedPlanStates
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.FetchLastNPlanStates`: %v\n", resp)
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
	resp, r, err := apiClient.PlanAPI.FetchLastPlanState(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.FetchLastPlanState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchLastPlanState`: EnvelopedTimedPlanState
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.FetchLastPlanState`: %v\n", resp)
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


## GetPlan

> EnvelopedPlan GetPlan(ctx, agencyId, planId).Mode(mode).Execute()

Retrieve a plan

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
	mode := openapiclient.PlanMode("standard") // PlanMode | The plan mode. (optional) (default to "standard")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlan(context.Background(), agencyId, planId).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlan`: EnvelopedPlan
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | [**PlanMode**](PlanMode.md) | The plan mode. | [default to &quot;standard&quot;]

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


## GetPlanSolution

> EnvelopedSolution GetPlanSolution(ctx, agencyId, planId).Execute()

Retrieve a plan solution

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
	resp, r, err := apiClient.PlanAPI.GetPlanSolution(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanSolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanSolution`: EnvelopedSolution
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanSolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanSolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvelopedSolution**](EnvelopedSolution.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanSolutionObjectives

> EnvelopedSolutionObjectives GetPlanSolutionObjectives(ctx, agencyId, planId).Execute()

Retrieve the objectives of a plan solution

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
	resp, r, err := apiClient.PlanAPI.GetPlanSolutionObjectives(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanSolutionObjectives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanSolutionObjectives`: EnvelopedSolutionObjectives
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanSolutionObjectives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanSolutionObjectivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvelopedSolutionObjectives**](EnvelopedSolutionObjectives.md)

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
	resp, r, err := apiClient.PlanAPI.GetPlanStatus(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanStatus`: EnvelopedPlanStatus
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanStatus`: %v\n", resp)
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


## GetPlans

> EnvelopedPlansLight GetPlans(ctx, agencyId).Page(page).ItemsPerPage(itemsPerPage).Archived(archived).Execute()

Retrieves a collection of plans



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
	page := int32(1) // int32 | The page number of a paginated list of records. (optional) (default to 1)
	itemsPerPage := int32(30) // int32 | The number of items per page in a paginated list of records. (optional) (default to 20)
	archived := "archived_example" // string | Indicates if and how archived items should be part of the result. (optional) (default to "excluded")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlans(context.Background(), agencyId).Page(page).ItemsPerPage(itemsPerPage).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlans`: EnvelopedPlansLight
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number of a paginated list of records. | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page in a paginated list of records. | [default to 20]
 **archived** | **string** | Indicates if and how archived items should be part of the result. | [default to &quot;excluded&quot;]

### Return type

[**EnvelopedPlansLight**](EnvelopedPlansLight.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlan

> EnvelopedPlan PutPlan(ctx, agencyId, planId).Force(force).Plan(plan).Execute()

Create or update a plan

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
	plan := *openapiclient.NewPlan("TODO", "TODO", []openapiclient.Resource{*openapiclient.NewResource("Id_example", openapiclient.Resource_vehicleProfile{VehicleProfileBicycle: openapiclient.NewVehicleProfileBicycle("Type_example")}, *openapiclient.NewTimeWindow("2019-11-15T12:34:56Z", "2019-11-15T12:34:56Z"))}) // Plan | The Plan to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.PutPlan(context.Background(), agencyId, planId).Force(force).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.PutPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlan`: EnvelopedPlan
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.PutPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | If true, on an archived item, the requested action will be forced and the item will be unarchived. | [default to false]
 **plan** | [**Plan**](Plan.md) | The Plan to update. | 

### Return type

[**EnvelopedPlan**](EnvelopedPlan.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlanMode

> EnvelopedPlanMode PutPlanMode(ctx, agencyId, planId).Force(force).Body(body).Execute()

Changes a plan's mode



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
	body := string(987) // string | The mode to set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.PutPlanMode(context.Background(), agencyId, planId).Force(force).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.PutPlanMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlanMode`: EnvelopedPlanMode
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.PutPlanMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlanModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | If true, on an archived item, the requested action will be forced and the item will be unarchived. | [default to false]
 **body** | **string** | The mode to set. | 

### Return type

[**EnvelopedPlanMode**](EnvelopedPlanMode.md)

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
	r, err := apiClient.PlanAPI.PutPlanRunning(context.Background(), agencyId, planId).Force(force).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.PutPlanRunning``: %v\n", err)
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

