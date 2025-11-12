# \SolutionAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlanSolution**](SolutionAPI.md#GetPlanSolution) | **Get** /agencies/{agencyId}/plans/{planId}/solution | Retrieve a plan solution
[**GetPlanSolutionObjectives**](SolutionAPI.md#GetPlanSolutionObjectives) | **Get** /agencies/{agencyId}/plans/{planId}/solution/objectives | Retrieve the objectives of a plan solution



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
	resp, r, err := apiClient.SolutionAPI.GetPlanSolution(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SolutionAPI.GetPlanSolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanSolution`: EnvelopedSolution
	fmt.Fprintf(os.Stdout, "Response from `SolutionAPI.GetPlanSolution`: %v\n", resp)
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
	resp, r, err := apiClient.SolutionAPI.GetPlanSolutionObjectives(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SolutionAPI.GetPlanSolutionObjectives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanSolutionObjectives`: EnvelopedSolutionObjectives
	fmt.Fprintf(os.Stdout, "Response from `SolutionAPI.GetPlanSolutionObjectives`: %v\n", resp)
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

