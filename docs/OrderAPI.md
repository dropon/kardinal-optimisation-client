# \OrderAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlanOrder**](OrderAPI.md#DeletePlanOrder) | **Delete** /agencies/{agencyId}/plans/{planId}/orders/{orderId} | Delete a plan&#39;s order
[**GetPlanOrder**](OrderAPI.md#GetPlanOrder) | **Get** /agencies/{agencyId}/plans/{planId}/orders/{orderId} | Retrieve a plan&#39;s order
[**PutPlanOrder**](OrderAPI.md#PutPlanOrder) | **Put** /agencies/{agencyId}/plans/{planId}/orders/{orderId} | Create or update a plan&#39;s order



## DeletePlanOrder

> DeletePlanOrder(ctx, agencyId, planId, orderId).Force(force).Execute()

Delete a plan's order

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
	orderId := "orderId_example" // string | The order id.
	force := true // bool | If true, on an archived item, the requested action will be forced and the item will be unarchived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrderAPI.DeletePlanOrder(context.Background(), agencyId, planId, orderId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.DeletePlanOrder``: %v\n", err)
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
**orderId** | **string** | The order id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanOrderRequest struct via the builder pattern


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


## GetPlanOrder

> EnvelopedOrder GetPlanOrder(ctx, agencyId, planId, orderId).Execute()

Retrieve a plan's order

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
	orderId := "orderId_example" // string | The order id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetPlanOrder(context.Background(), agencyId, planId, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetPlanOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanOrder`: EnvelopedOrder
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetPlanOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**orderId** | **string** | The order id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnvelopedOrder**](EnvelopedOrder.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlanOrder

> EnvelopedOrder PutPlanOrder(ctx, agencyId, planId, orderId).Force(force).Order(order).Execute()

Create or update a plan's order

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
	orderId := "orderId_example" // string | The order id.
	force := true // bool | If true, on an archived item, the requested action will be forced and the item will be unarchived. (optional) (default to false)
	order := *openapiclient.NewOrder("Id_example", []openapiclient.Stop{openapiclient.Stop{AlternativesStop: openapiclient.NewAlternativesStop("Type_example", []openapiclient.SingleStop{*openapiclient.NewSingleStop("Id_example", *openapiclient.NewPosition(float32(123), float32(123)))})}}) // Order | The Order to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.PutPlanOrder(context.Background(), agencyId, planId, orderId).Force(force).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.PutPlanOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlanOrder`: EnvelopedOrder
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.PutPlanOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**orderId** | **string** | The order id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlanOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | If true, on an archived item, the requested action will be forced and the item will be unarchived. | [default to false]
 **order** | [**Order**](Order.md) | The Order to update. | 

### Return type

[**EnvelopedOrder**](EnvelopedOrder.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

