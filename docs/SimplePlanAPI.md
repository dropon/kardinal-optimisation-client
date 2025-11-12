# \SimplePlanAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSimplePlan**](SimplePlanAPI.md#PostSimplePlan) | **Post** /agencies/{agencyId}/simplePlans | Create a simple plan



## PostSimplePlan

> EnvelopedSimplePlan PostSimplePlan(ctx, agencyId).SimplePlan(simplePlan).Execute()

Create a simple plan

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
	simplePlan := *openapiclient.NewSimplePlan("cd4ce4e3-0208-4b10-b346-25f235214e4f", "TODO", []openapiclient.SimpleResource{*openapiclient.NewSimpleResource("Id_example", *openapiclient.NewTimeWindow("2019-11-15T12:34:56Z", "2019-11-15T12:34:56Z"))}) // SimplePlan | The SimplePlan to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimplePlanAPI.PostSimplePlan(context.Background(), agencyId).SimplePlan(simplePlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimplePlanAPI.PostSimplePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimplePlan`: EnvelopedSimplePlan
	fmt.Fprintf(os.Stdout, "Response from `SimplePlanAPI.PostSimplePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSimplePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simplePlan** | [**SimplePlan**](SimplePlan.md) | The SimplePlan to create. | 

### Return type

[**EnvelopedSimplePlan**](EnvelopedSimplePlan.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

