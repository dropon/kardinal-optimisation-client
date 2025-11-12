# \RegionAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClientRegion**](RegionAPI.md#DeleteClientRegion) | **Delete** /clients/{clientId}/countries/{countryId}/regions/{regionId} | Deletes an existing region
[**DeleteRegion**](RegionAPI.md#DeleteRegion) | **Delete** /regions/{regionId} | Deletes an existing region
[**GetClientRegions**](RegionAPI.md#GetClientRegions) | **Get** /clients/{clientId}/regions | Retrieves the regions of a single client (including their agencies)
[**GetRegions**](RegionAPI.md#GetRegions) | **Get** /regions | Retrieves the regions which are visible by the current user
[**PostClientRegion**](RegionAPI.md#PostClientRegion) | **Post** /clients/{clientId}/countries/{countryId}/regions | Creates a new region
[**PutClientRegion**](RegionAPI.md#PutClientRegion) | **Put** /clients/{clientId}/countries/{countryId}/regions/{regionId} | Updates an existing region or creates a new one if it doesn&#39;t exist



## DeleteClientRegion

> DeleteClientRegion(ctx, clientId, countryId, regionId).Execute()

Deletes an existing region

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
	clientId := TODO // ClientId | The client id.
	countryId := TODO // CountryId | The country id.
	regionId := TODO // RegionId | The region id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegionAPI.DeleteClientRegion(context.Background(), clientId, countryId, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.DeleteClientRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 
**countryId** | [**CountryId**](.md) | The country id. | 
**regionId** | [**RegionId**](.md) | The region id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRegionRequest struct via the builder pattern


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


## DeleteRegion

> DeleteRegion(ctx, regionId).Execute()

Deletes an existing region

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
	regionId := TODO // RegionId | The region id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegionAPI.DeleteRegion(context.Background(), regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.DeleteRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | [**RegionId**](.md) | The region id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegionRequest struct via the builder pattern


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


## GetClientRegions

> EnvelopedRegions GetClientRegions(ctx, clientId).Execute()

Retrieves the regions of a single client (including their agencies)

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
	clientId := TODO // ClientId | The client id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionAPI.GetClientRegions(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.GetClientRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRegions`: EnvelopedRegions
	fmt.Fprintf(os.Stdout, "Response from `RegionAPI.GetClientRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvelopedRegions**](EnvelopedRegions.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegions

> EnvelopedRegions GetRegions(ctx).UntilLevel(untilLevel).Execute()

Retrieves the regions which are visible by the current user



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
	untilLevel := "untilLevel_example" // string | The depth level. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionAPI.GetRegions(context.Background()).UntilLevel(untilLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.GetRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegions`: EnvelopedRegions
	fmt.Fprintf(os.Stdout, "Response from `RegionAPI.GetRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **untilLevel** | **string** | The depth level. | 

### Return type

[**EnvelopedRegions**](EnvelopedRegions.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientRegion

> EnvelopedPlainRegion PostClientRegion(ctx, clientId, countryId).RegionForCreation(regionForCreation).Execute()

Creates a new region

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
	clientId := TODO // ClientId | The client id.
	countryId := TODO // CountryId | The country id.
	regionForCreation := *openapiclient.NewRegionForCreation("TODO") // RegionForCreation | A region entity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionAPI.PostClientRegion(context.Background(), clientId, countryId).RegionForCreation(regionForCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.PostClientRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientRegion`: EnvelopedPlainRegion
	fmt.Fprintf(os.Stdout, "Response from `RegionAPI.PostClientRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 
**countryId** | [**CountryId**](.md) | The country id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **regionForCreation** | [**RegionForCreation**](RegionForCreation.md) | A region entity. | 

### Return type

[**EnvelopedPlainRegion**](EnvelopedPlainRegion.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClientRegion

> EnvelopedPlainRegion PutClientRegion(ctx, clientId, countryId, regionId).PutClientRegionRequest(putClientRegionRequest).Execute()

Updates an existing region or creates a new one if it doesn't exist

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
	clientId := TODO // ClientId | The client id.
	countryId := TODO // CountryId | The country id.
	regionId := TODO // RegionId | The region id.
	putClientRegionRequest := openapiclient.putClientRegion_request{RegionForCreation: openapiclient.NewRegionForCreation("TODO")} // PutClientRegionRequest | A region entity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionAPI.PutClientRegion(context.Background(), clientId, countryId, regionId).PutClientRegionRequest(putClientRegionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionAPI.PutClientRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientRegion`: EnvelopedPlainRegion
	fmt.Fprintf(os.Stdout, "Response from `RegionAPI.PutClientRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 
**countryId** | [**CountryId**](.md) | The country id. | 
**regionId** | [**RegionId**](.md) | The region id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putClientRegionRequest** | [**PutClientRegionRequest**](PutClientRegionRequest.md) | A region entity. | 

### Return type

[**EnvelopedPlainRegion**](EnvelopedPlainRegion.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

