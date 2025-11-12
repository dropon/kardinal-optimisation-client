# \CountryAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClientCountry**](CountryAPI.md#DeleteClientCountry) | **Delete** /clients/{clientId}/countries/{countryId} | Deletes an existing country
[**DeleteCountry**](CountryAPI.md#DeleteCountry) | **Delete** /countries/{countryId} | Deletes an existing country
[**GetClientCountries**](CountryAPI.md#GetClientCountries) | **Get** /clients/{clientId}/countries | Retrieves the countries of a single client (including their regions and agencies)
[**GetCountries**](CountryAPI.md#GetCountries) | **Get** /countries | Retrieves the countries which are visible by the current user
[**PostClientCountry**](CountryAPI.md#PostClientCountry) | **Post** /clients/{clientId}/countries | Creates a new country
[**PutClientCountry**](CountryAPI.md#PutClientCountry) | **Put** /clients/{clientId}/countries/{countryId} | Updates an existing country or creates a new one if it doesn&#39;t exist



## DeleteClientCountry

> DeleteClientCountry(ctx, clientId, countryId).Execute()

Deletes an existing country

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CountryAPI.DeleteClientCountry(context.Background(), clientId, countryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.DeleteClientCountry``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientCountryRequest struct via the builder pattern


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


## DeleteCountry

> DeleteCountry(ctx, countryId).Execute()

Deletes an existing country

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
	countryId := TODO // CountryId | The country id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CountryAPI.DeleteCountry(context.Background(), countryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.DeleteCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | [**CountryId**](.md) | The country id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCountryRequest struct via the builder pattern


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


## GetClientCountries

> EnvelopedCountries GetClientCountries(ctx, clientId).Execute()

Retrieves the countries of a single client (including their regions and agencies)

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
	resp, r, err := apiClient.CountryAPI.GetClientCountries(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.GetClientCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientCountries`: EnvelopedCountries
	fmt.Fprintf(os.Stdout, "Response from `CountryAPI.GetClientCountries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvelopedCountries**](EnvelopedCountries.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountries

> EnvelopedCountries GetCountries(ctx).UntilLevel(untilLevel).Execute()

Retrieves the countries which are visible by the current user



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
	resp, r, err := apiClient.CountryAPI.GetCountries(context.Background()).UntilLevel(untilLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountries`: EnvelopedCountries
	fmt.Fprintf(os.Stdout, "Response from `CountryAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **untilLevel** | **string** | The depth level. | 

### Return type

[**EnvelopedCountries**](EnvelopedCountries.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClientCountry

> EnvelopedPlainCountry PostClientCountry(ctx, clientId).CountryForCreation(countryForCreation).Execute()

Creates a new country

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
	countryForCreation := *openapiclient.NewCountryForCreation("TODO") // CountryForCreation | A country entity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountryAPI.PostClientCountry(context.Background(), clientId).CountryForCreation(countryForCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.PostClientCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientCountry`: EnvelopedPlainCountry
	fmt.Fprintf(os.Stdout, "Response from `CountryAPI.PostClientCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClientCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countryForCreation** | [**CountryForCreation**](CountryForCreation.md) | A country entity. | 

### Return type

[**EnvelopedPlainCountry**](EnvelopedPlainCountry.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClientCountry

> EnvelopedPlainCountry PutClientCountry(ctx, clientId, countryId).PutClientCountryRequest(putClientCountryRequest).Execute()

Updates an existing country or creates a new one if it doesn't exist

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
	putClientCountryRequest := openapiclient.putClientCountry_request{CountryForCreation: openapiclient.NewCountryForCreation("TODO")} // PutClientCountryRequest | A country entity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountryAPI.PutClientCountry(context.Background(), clientId, countryId).PutClientCountryRequest(putClientCountryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.PutClientCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientCountry`: EnvelopedPlainCountry
	fmt.Fprintf(os.Stdout, "Response from `CountryAPI.PutClientCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 
**countryId** | [**CountryId**](.md) | The country id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putClientCountryRequest** | [**PutClientCountryRequest**](PutClientCountryRequest.md) | A country entity. | 

### Return type

[**EnvelopedPlainCountry**](EnvelopedPlainCountry.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

