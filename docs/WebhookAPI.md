# \WebhookAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgencyWebhook**](WebhookAPI.md#CreateAgencyWebhook) | **Post** /agencies/{agencyId}/webhooks | Create a new agency webhook
[**CreatePlanWebhook**](WebhookAPI.md#CreatePlanWebhook) | **Post** /agencies/{agencyId}/plans/{planId}/webhooks | Create a new plan webhook
[**DeleteAgencyWebhook**](WebhookAPI.md#DeleteAgencyWebhook) | **Delete** /agencies/{agencyId}/webhooks/{webhookId} | Deletes a single agency webhook
[**DeletePlanWebhook**](WebhookAPI.md#DeletePlanWebhook) | **Delete** /agencies/{agencyId}/plans/{planId}/webhooks/{webhookId} | Deletes a single plan webhook
[**FetchAgencyWebhook**](WebhookAPI.md#FetchAgencyWebhook) | **Get** /agencies/{agencyId}/webhooks/{webhookId} | Fetch a single agency webhook
[**FetchAgencyWebhookResults**](WebhookAPI.md#FetchAgencyWebhookResults) | **Get** /agencies/{agencyId}/webhooks/{webhookId}/results | Fetch an agency webhook&#39;s list of results
[**FetchAgencyWebhooks**](WebhookAPI.md#FetchAgencyWebhooks) | **Get** /agencies/{agencyId}/webhooks | Fetch all agency webhooks
[**FetchPlanWebhook**](WebhookAPI.md#FetchPlanWebhook) | **Get** /agencies/{agencyId}/plans/{planId}/webhooks/{webhookId} | Fetch a single plan webhook
[**FetchPlanWebhookResults**](WebhookAPI.md#FetchPlanWebhookResults) | **Get** /agencies/{agencyId}/plans/{planId}/webhooks/{webhookId}/results | Fetch a plan webhook&#39;s list of results
[**FetchPlanWebhooks**](WebhookAPI.md#FetchPlanWebhooks) | **Get** /agencies/{agencyId}/plans/{planId}/webhooks | Fetch all plan webhooks
[**UpdateAgencyWebhook**](WebhookAPI.md#UpdateAgencyWebhook) | **Put** /agencies/{agencyId}/webhooks/{webhookId} | Updates an existing agency webhook
[**UpdatePlanWebhook**](WebhookAPI.md#UpdatePlanWebhook) | **Put** /agencies/{agencyId}/plans/{planId}/webhooks/{webhookId} | Updates an existing plan webhook



## CreateAgencyWebhook

> EnvelopedAgencyWebhook CreateAgencyWebhook(ctx, agencyId).WebhookCreate(webhookCreate).Execute()

Create a new agency webhook

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
	webhookCreate := *openapiclient.NewWebhookCreate("https://mydomain.com/webhooks") // WebhookCreate | The webhook for creation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.CreateAgencyWebhook(context.Background(), agencyId).WebhookCreate(webhookCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.CreateAgencyWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgencyWebhook`: EnvelopedAgencyWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.CreateAgencyWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgencyWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookCreate** | [**WebhookCreate**](WebhookCreate.md) | The webhook for creation. | 

### Return type

[**EnvelopedAgencyWebhook**](EnvelopedAgencyWebhook.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlanWebhook

> EnvelopedPlanWebhook CreatePlanWebhook(ctx, agencyId, planId).WebhookCreate(webhookCreate).Execute()

Create a new plan webhook

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
	webhookCreate := *openapiclient.NewWebhookCreate("https://mydomain.com/webhooks") // WebhookCreate | The webhook for creation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.CreatePlanWebhook(context.Background(), agencyId, planId).WebhookCreate(webhookCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.CreatePlanWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlanWebhook`: EnvelopedPlanWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.CreatePlanWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **webhookCreate** | [**WebhookCreate**](WebhookCreate.md) | The webhook for creation. | 

### Return type

[**EnvelopedPlanWebhook**](EnvelopedPlanWebhook.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgencyWebhook

> DeleteAgencyWebhook(ctx, agencyId, webhookId).Execute()

Deletes a single agency webhook

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The webhook UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookAPI.DeleteAgencyWebhook(context.Background(), agencyId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.DeleteAgencyWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**webhookId** | **string** | The webhook UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgencyWebhookRequest struct via the builder pattern


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


## DeletePlanWebhook

> DeletePlanWebhook(ctx, agencyId, planId, webhookId).Execute()

Deletes a single plan webhook

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The webhook UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookAPI.DeletePlanWebhook(context.Background(), agencyId, planId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.DeletePlanWebhook``: %v\n", err)
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
**webhookId** | **string** | The webhook UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanWebhookRequest struct via the builder pattern


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


## FetchAgencyWebhook

> EnvelopedAgencyWebhook FetchAgencyWebhook(ctx, agencyId, webhookId).Execute()

Fetch a single agency webhook

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The webhook UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.FetchAgencyWebhook(context.Background(), agencyId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.FetchAgencyWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchAgencyWebhook`: EnvelopedAgencyWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.FetchAgencyWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**webhookId** | **string** | The webhook UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAgencyWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvelopedAgencyWebhook**](EnvelopedAgencyWebhook.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAgencyWebhookResults

> EnvelopedAgencyWebhookResults FetchAgencyWebhookResults(ctx, agencyId, webhookId).Execute()

Fetch an agency webhook's list of results

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The webhook UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.FetchAgencyWebhookResults(context.Background(), agencyId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.FetchAgencyWebhookResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchAgencyWebhookResults`: EnvelopedAgencyWebhookResults
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.FetchAgencyWebhookResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**webhookId** | **string** | The webhook UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAgencyWebhookResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvelopedAgencyWebhookResults**](EnvelopedAgencyWebhookResults.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAgencyWebhooks

> EnvelopedAgencyWebhooks FetchAgencyWebhooks(ctx, agencyId).Execute()

Fetch all agency webhooks

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
	resp, r, err := apiClient.WebhookAPI.FetchAgencyWebhooks(context.Background(), agencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.FetchAgencyWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchAgencyWebhooks`: EnvelopedAgencyWebhooks
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.FetchAgencyWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAgencyWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvelopedAgencyWebhooks**](EnvelopedAgencyWebhooks.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPlanWebhook

> EnvelopedPlanWebhook FetchPlanWebhook(ctx, agencyId, planId, webhookId).Execute()

Fetch a single plan webhook

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The webhook UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.FetchPlanWebhook(context.Background(), agencyId, planId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.FetchPlanWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPlanWebhook`: EnvelopedPlanWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.FetchPlanWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**webhookId** | **string** | The webhook UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPlanWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnvelopedPlanWebhook**](EnvelopedPlanWebhook.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPlanWebhookResults

> EnvelopedPlanWebhookResults FetchPlanWebhookResults(ctx, agencyId, planId, webhookId).Execute()

Fetch a plan webhook's list of results

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The webhook UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.FetchPlanWebhookResults(context.Background(), agencyId, planId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.FetchPlanWebhookResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPlanWebhookResults`: EnvelopedPlanWebhookResults
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.FetchPlanWebhookResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**webhookId** | **string** | The webhook UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPlanWebhookResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnvelopedPlanWebhookResults**](EnvelopedPlanWebhookResults.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPlanWebhooks

> EnvelopedPlanWebhooks FetchPlanWebhooks(ctx, agencyId, planId).Execute()

Fetch all plan webhooks

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
	resp, r, err := apiClient.WebhookAPI.FetchPlanWebhooks(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.FetchPlanWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPlanWebhooks`: EnvelopedPlanWebhooks
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.FetchPlanWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPlanWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvelopedPlanWebhooks**](EnvelopedPlanWebhooks.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgencyWebhook

> EnvelopedAgencyWebhook UpdateAgencyWebhook(ctx, agencyId, webhookId).WebhookUpdate(webhookUpdate).Execute()

Updates an existing agency webhook

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The webhook UUID.
	webhookUpdate := *openapiclient.NewWebhookUpdate("https://mydomain.com/webhooks") // WebhookUpdate | The webhook for update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.UpdateAgencyWebhook(context.Background(), agencyId, webhookId).WebhookUpdate(webhookUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.UpdateAgencyWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgencyWebhook`: EnvelopedAgencyWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.UpdateAgencyWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**webhookId** | **string** | The webhook UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgencyWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **webhookUpdate** | [**WebhookUpdate**](WebhookUpdate.md) | The webhook for update. | 

### Return type

[**EnvelopedAgencyWebhook**](EnvelopedAgencyWebhook.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlanWebhook

> EnvelopedPlanWebhook UpdatePlanWebhook(ctx, agencyId, planId, webhookId).WebhookUpdate(webhookUpdate).Execute()

Updates an existing plan webhook

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The webhook UUID.
	webhookUpdate := *openapiclient.NewWebhookUpdate("https://mydomain.com/webhooks") // WebhookUpdate | The webhook for update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.UpdatePlanWebhook(context.Background(), agencyId, planId, webhookId).WebhookUpdate(webhookUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.UpdatePlanWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlanWebhook`: EnvelopedPlanWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.UpdatePlanWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 
**webhookId** | **string** | The webhook UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **webhookUpdate** | [**WebhookUpdate**](WebhookUpdate.md) | The webhook for update. | 

### Return type

[**EnvelopedPlanWebhook**](EnvelopedPlanWebhook.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

