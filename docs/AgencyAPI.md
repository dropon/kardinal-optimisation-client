# \AgencyAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAgency**](AgencyAPI.md#DeleteAgency) | **Delete** /agencies/{agencyId} | Deletes an existing agency
[**DeleteAgencyConfGlobalTravelTimeCoefficient**](AgencyAPI.md#DeleteAgencyConfGlobalTravelTimeCoefficient) | **Delete** /agencies/{agencyId}/conf/globalTravelTimeCoefficient | Resets the agency&#39;s configured global travel time coefficient
[**DeleteClientAgency**](AgencyAPI.md#DeleteClientAgency) | **Delete** /clients/{clientId}/countries/{countryId}/regions/{regionId}/agencies/{agencyId} | Deletes an existing agency
[**GetAgencies**](AgencyAPI.md#GetAgencies) | **Get** /agencies | Retrieves the agencies which are visible by the current user
[**GetAgencySectorMapping**](AgencyAPI.md#GetAgencySectorMapping) | **Get** /agencies/{agencyId}/sectorMapping | Get the sector mapping which is currently applied in an agency, if it exists
[**GetClientAgencies**](AgencyAPI.md#GetClientAgencies) | **Get** /clients/{clientId}/agencies | Retrieves the agencies of a single client
[**PostClientAgency**](AgencyAPI.md#PostClientAgency) | **Post** /clients/{clientId}/countries/{countryId}/regions/{regionId}/agencies | Creates a new agency
[**PutAgencyConfGlobalTravelTimeCoefficient**](AgencyAPI.md#PutAgencyConfGlobalTravelTimeCoefficient) | **Put** /agencies/{agencyId}/conf/globalTravelTimeCoefficient | Updates the agency&#39;s configured global travel time coefficient configuration
[**PutAgencySectorMapping**](AgencyAPI.md#PutAgencySectorMapping) | **Put** /agencies/{agencyId}/sectorMapping | Put the sector mapping to be applied in an agency
[**PutClientAgency**](AgencyAPI.md#PutClientAgency) | **Put** /clients/{clientId}/countries/{countryId}/regions/{regionId}/agencies/{agencyId} | Updates an existing agency or creates a new one if it doesn&#39;t exist



## DeleteAgency

> DeleteAgency(ctx, agencyId).Execute()

Deletes an existing agency

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
	r, err := apiClient.AgencyAPI.DeleteAgency(context.Background(), agencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.DeleteAgency``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAgencyRequest struct via the builder pattern


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
	r, err := apiClient.AgencyAPI.DeleteAgencyConfGlobalTravelTimeCoefficient(context.Background(), agencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.DeleteAgencyConfGlobalTravelTimeCoefficient``: %v\n", err)
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


## DeleteClientAgency

> DeleteClientAgency(ctx, clientId, countryId, regionId, agencyId).Execute()

Deletes an existing agency

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
	agencyId := TODO // AgencyId | The agency id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgencyAPI.DeleteClientAgency(context.Background(), clientId, countryId, regionId, agencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.DeleteClientAgency``: %v\n", err)
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
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientAgencyRequest struct via the builder pattern


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


## GetAgencies

> EnvelopedAgencies GetAgencies(ctx).Execute()

Retrieves the agencies which are visible by the current user

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
	resp, r, err := apiClient.AgencyAPI.GetAgencies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.GetAgencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgencies`: EnvelopedAgencies
	fmt.Fprintf(os.Stdout, "Response from `AgencyAPI.GetAgencies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgenciesRequest struct via the builder pattern


### Return type

[**EnvelopedAgencies**](EnvelopedAgencies.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgencySectorMapping

> EnvelopedSectorMapping GetAgencySectorMapping(ctx, agencyId).Execute()

Get the sector mapping which is currently applied in an agency, if it exists

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
	resp, r, err := apiClient.AgencyAPI.GetAgencySectorMapping(context.Background(), agencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.GetAgencySectorMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgencySectorMapping`: EnvelopedSectorMapping
	fmt.Fprintf(os.Stdout, "Response from `AgencyAPI.GetAgencySectorMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgencySectorMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvelopedSectorMapping**](EnvelopedSectorMapping.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientAgencies

> EnvelopedAgencies GetClientAgencies(ctx, clientId).Execute()

Retrieves the agencies of a single client

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
	resp, r, err := apiClient.AgencyAPI.GetClientAgencies(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.GetClientAgencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientAgencies`: EnvelopedAgencies
	fmt.Fprintf(os.Stdout, "Response from `AgencyAPI.GetClientAgencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientAgenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvelopedAgencies**](EnvelopedAgencies.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientAgency

> EnvelopedAgency PostClientAgency(ctx, clientId, countryId, regionId).AgencyForCreation(agencyForCreation).Execute()

Creates a new agency

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
	agencyForCreation := *openapiclient.NewAgencyForCreation("TODO") // AgencyForCreation | An agency entity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgencyAPI.PostClientAgency(context.Background(), clientId, countryId, regionId).AgencyForCreation(agencyForCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.PostClientAgency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientAgency`: EnvelopedAgency
	fmt.Fprintf(os.Stdout, "Response from `AgencyAPI.PostClientAgency`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPostClientAgencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **agencyForCreation** | [**AgencyForCreation**](AgencyForCreation.md) | An agency entity. | 

### Return type

[**EnvelopedAgency**](EnvelopedAgency.md)

### Authorization

[access_token](../README.md#access_token)

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
	resp, r, err := apiClient.AgencyAPI.PutAgencyConfGlobalTravelTimeCoefficient(context.Background(), agencyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.PutAgencyConfGlobalTravelTimeCoefficient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAgencyConfGlobalTravelTimeCoefficient`: EnvelopedGlobalTravelTimeCoefficient
	fmt.Fprintf(os.Stdout, "Response from `AgencyAPI.PutAgencyConfGlobalTravelTimeCoefficient`: %v\n", resp)
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


## PutAgencySectorMapping

> EnvelopedSectorMapping PutAgencySectorMapping(ctx, agencyId).RequestBody(requestBody).Execute()

Put the sector mapping to be applied in an agency

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
	requestBody := map[string]MultiPolygonFeature{"key": *openapiclient.NewMultiPolygonFeature(openapiclient.FeatureType("Feature"), *openapiclient.NewMultiPolygon(openapiclient.MultiPolygonType("MultiPolygon"), [][][][]float32{[][][]float32{[][]float32{[]float32{float32(123)}}}}))} // map[string]MultiPolygonFeature | The sector mapping to be applied. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgencyAPI.PutAgencySectorMapping(context.Background(), agencyId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.PutAgencySectorMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAgencySectorMapping`: EnvelopedSectorMapping
	fmt.Fprintf(os.Stdout, "Response from `AgencyAPI.PutAgencySectorMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAgencySectorMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**map[string]MultiPolygonFeature**](MultiPolygonFeature.md) | The sector mapping to be applied. | 

### Return type

[**EnvelopedSectorMapping**](EnvelopedSectorMapping.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClientAgency

> EnvelopedAgency PutClientAgency(ctx, clientId, countryId, regionId, agencyId).PutClientAgencyRequest(putClientAgencyRequest).Execute()

Updates an existing agency or creates a new one if it doesn't exist

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
	agencyId := TODO // AgencyId | The agency id.
	putClientAgencyRequest := openapiclient.putClientAgency_request{AgencyForCreation: openapiclient.NewAgencyForCreation("TODO")} // PutClientAgencyRequest | An agency entity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgencyAPI.PutClientAgency(context.Background(), clientId, countryId, regionId, agencyId).PutClientAgencyRequest(putClientAgencyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgencyAPI.PutClientAgency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientAgency`: EnvelopedAgency
	fmt.Fprintf(os.Stdout, "Response from `AgencyAPI.PutClientAgency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 
**countryId** | [**CountryId**](.md) | The country id. | 
**regionId** | [**RegionId**](.md) | The region id. | 
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientAgencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **putClientAgencyRequest** | [**PutClientAgencyRequest**](PutClientAgencyRequest.md) | An agency entity. | 

### Return type

[**EnvelopedAgency**](EnvelopedAgency.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

