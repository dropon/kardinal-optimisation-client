# \ClientAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClient**](ClientAPI.md#DeleteClient) | **Delete** /clients/{clientId} | Deletes an existing client
[**DeleteClientAgency**](ClientAPI.md#DeleteClientAgency) | **Delete** /clients/{clientId}/countries/{countryId}/regions/{regionId}/agencies/{agencyId} | Deletes an existing agency
[**DeleteClientCountry**](ClientAPI.md#DeleteClientCountry) | **Delete** /clients/{clientId}/countries/{countryId} | Deletes an existing country
[**DeleteClientRegion**](ClientAPI.md#DeleteClientRegion) | **Delete** /clients/{clientId}/countries/{countryId}/regions/{regionId} | Deletes an existing region
[**GetClient**](ClientAPI.md#GetClient) | **Get** /clients/{clientId} | Retrieves a single client
[**GetClientAgencies**](ClientAPI.md#GetClientAgencies) | **Get** /clients/{clientId}/agencies | Retrieves the agencies of a single client
[**GetClientConf**](ClientAPI.md#GetClientConf) | **Get** /clients/{clientId}/conf | Retrieves the conf of a single client
[**GetClientCountries**](ClientAPI.md#GetClientCountries) | **Get** /clients/{clientId}/countries | Retrieves the countries of a single client (including their regions and agencies)
[**GetClientRegions**](ClientAPI.md#GetClientRegions) | **Get** /clients/{clientId}/regions | Retrieves the regions of a single client (including their agencies)
[**GetClients**](ClientAPI.md#GetClients) | **Get** /clients | Retrieves the clients which are visible by the current user
[**PostClient**](ClientAPI.md#PostClient) | **Post** /clients | Creates a single client
[**PostClientAgency**](ClientAPI.md#PostClientAgency) | **Post** /clients/{clientId}/countries/{countryId}/regions/{regionId}/agencies | Creates a new agency
[**PostClientCountry**](ClientAPI.md#PostClientCountry) | **Post** /clients/{clientId}/countries | Creates a new country
[**PostClientRegion**](ClientAPI.md#PostClientRegion) | **Post** /clients/{clientId}/countries/{countryId}/regions | Creates a new region
[**PutClient**](ClientAPI.md#PutClient) | **Put** /clients/{clientId} | Updates or creates a single client
[**PutClientAgency**](ClientAPI.md#PutClientAgency) | **Put** /clients/{clientId}/countries/{countryId}/regions/{regionId}/agencies/{agencyId} | Updates an existing agency or creates a new one if it doesn&#39;t exist
[**PutClientConf**](ClientAPI.md#PutClientConf) | **Put** /clients/{clientId}/conf | Updates the conf of a single client
[**PutClientCountry**](ClientAPI.md#PutClientCountry) | **Put** /clients/{clientId}/countries/{countryId} | Updates an existing country or creates a new one if it doesn&#39;t exist
[**PutClientRegion**](ClientAPI.md#PutClientRegion) | **Put** /clients/{clientId}/countries/{countryId}/regions/{regionId} | Updates an existing region or creates a new one if it doesn&#39;t exist
[**PutIsAvailableClientName**](ClientAPI.md#PutIsAvailableClientName) | **Put** /clients/isAvailable/name | Checks the availability of a client name
[**PutIsAvailableClientPrefix**](ClientAPI.md#PutIsAvailableClientPrefix) | **Put** /clients/isAvailable/prefix | Checks the availability of a client prefix



## DeleteClient

> DeleteClient(ctx, clientId).Execute()

Deletes an existing client

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
	r, err := apiClient.ClientAPI.DeleteClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRequest struct via the builder pattern


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
	r, err := apiClient.ClientAPI.DeleteClientAgency(context.Background(), clientId, countryId, regionId, agencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteClientAgency``: %v\n", err)
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
	r, err := apiClient.ClientAPI.DeleteClientCountry(context.Background(), clientId, countryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteClientCountry``: %v\n", err)
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
	r, err := apiClient.ClientAPI.DeleteClientRegion(context.Background(), clientId, countryId, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.DeleteClientRegion``: %v\n", err)
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


## GetClient

> EnvelopedClient GetClient(ctx, clientId).UntilLevel(untilLevel).Execute()

Retrieves a single client



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
	untilLevel := "untilLevel_example" // string | The depth level. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.GetClient(context.Background(), clientId).UntilLevel(untilLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClient`: EnvelopedClient
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **untilLevel** | **string** | The depth level. | 

### Return type

[**EnvelopedClient**](EnvelopedClient.md)

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
	resp, r, err := apiClient.ClientAPI.GetClientAgencies(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientAgencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientAgencies`: EnvelopedAgencies
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientAgencies`: %v\n", resp)
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


## GetClientConf

> EnvelopedClientConf GetClientConf(ctx, clientId).Execute()

Retrieves the conf of a single client

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
	resp, r, err := apiClient.ClientAPI.GetClientConf(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientConf`: EnvelopedClientConf
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientConf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientConfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvelopedClientConf**](EnvelopedClientConf.md)

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
	resp, r, err := apiClient.ClientAPI.GetClientCountries(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientCountries`: EnvelopedCountries
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientCountries`: %v\n", resp)
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
	resp, r, err := apiClient.ClientAPI.GetClientRegions(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientRegions`: EnvelopedRegions
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientRegions`: %v\n", resp)
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


## GetClients

> EnvelopedClients GetClients(ctx).UntilLevel(untilLevel).Execute()

Retrieves the clients which are visible by the current user



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
	resp, r, err := apiClient.ClientAPI.GetClients(context.Background()).UntilLevel(untilLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: EnvelopedClients
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **untilLevel** | **string** | The depth level. | 

### Return type

[**EnvelopedClients**](EnvelopedClients.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClient

> EnvelopedClient PostClient(ctx).ClientForCreation(clientForCreation).Execute()

Creates a single client



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
	clientForCreation := *openapiclient.NewClientForCreation("TODO", "TODO") // ClientForCreation | A client entity contains the following informations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.PostClient(context.Background()).ClientForCreation(clientForCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PostClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClient`: EnvelopedClient
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PostClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientForCreation** | [**ClientForCreation**](ClientForCreation.md) | A client entity contains the following informations. | 

### Return type

[**EnvelopedClient**](EnvelopedClient.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.ClientAPI.PostClientAgency(context.Background(), clientId, countryId, regionId).AgencyForCreation(agencyForCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PostClientAgency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientAgency`: EnvelopedAgency
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PostClientAgency`: %v\n", resp)
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
	resp, r, err := apiClient.ClientAPI.PostClientCountry(context.Background(), clientId).CountryForCreation(countryForCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PostClientCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientCountry`: EnvelopedPlainCountry
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PostClientCountry`: %v\n", resp)
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
	resp, r, err := apiClient.ClientAPI.PostClientRegion(context.Background(), clientId, countryId).RegionForCreation(regionForCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PostClientRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClientRegion`: EnvelopedPlainRegion
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PostClientRegion`: %v\n", resp)
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


## PutClient

> EnvelopedClient PutClient(ctx, clientId).PutClientRequest(putClientRequest).Execute()

Updates or creates a single client



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
	putClientRequest := openapiclient.putClient_request{ClientForCreation: openapiclient.NewClientForCreation("TODO", "TODO")} // PutClientRequest | The client to update / create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.PutClient(context.Background(), clientId).PutClientRequest(putClientRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PutClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClient`: EnvelopedClient
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PutClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putClientRequest** | [**PutClientRequest**](PutClientRequest.md) | The client to update / create. | 

### Return type

[**EnvelopedClient**](EnvelopedClient.md)

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
	resp, r, err := apiClient.ClientAPI.PutClientAgency(context.Background(), clientId, countryId, regionId, agencyId).PutClientAgencyRequest(putClientAgencyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PutClientAgency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientAgency`: EnvelopedAgency
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PutClientAgency`: %v\n", resp)
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


## PutClientConf

> EnvelopedClientConf PutClientConf(ctx, clientId).ClientConf(clientConf).Execute()

Updates the conf of a single client

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
	clientConf := *openapiclient.NewClientConf() // ClientConf | The client conf to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.PutClientConf(context.Background(), clientId).ClientConf(clientConf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PutClientConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientConf`: EnvelopedClientConf
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PutClientConf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | [**ClientId**](.md) | The client id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClientConfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientConf** | [**ClientConf**](ClientConf.md) | The client conf to update. | 

### Return type

[**EnvelopedClientConf**](EnvelopedClientConf.md)

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
	resp, r, err := apiClient.ClientAPI.PutClientCountry(context.Background(), clientId, countryId).PutClientCountryRequest(putClientCountryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PutClientCountry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientCountry`: EnvelopedPlainCountry
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PutClientCountry`: %v\n", resp)
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
	resp, r, err := apiClient.ClientAPI.PutClientRegion(context.Background(), clientId, countryId, regionId).PutClientRegionRequest(putClientRegionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PutClientRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClientRegion`: EnvelopedPlainRegion
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PutClientRegion`: %v\n", resp)
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


## PutIsAvailableClientName

> EnvelopedString PutIsAvailableClientName(ctx).Body(body).Execute()

Checks the availability of a client name

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
	body := "body_example" // string | The client name to check. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.PutIsAvailableClientName(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PutIsAvailableClientName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIsAvailableClientName`: EnvelopedString
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PutIsAvailableClientName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIsAvailableClientNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | The client name to check. | 

### Return type

[**EnvelopedString**](EnvelopedString.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIsAvailableClientPrefix

> EnvelopedString PutIsAvailableClientPrefix(ctx).ClientPrefix(clientPrefix).Execute()

Checks the availability of a client prefix

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
	clientPrefix := TODO // ClientPrefix | The client prefix to check. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAPI.PutIsAvailableClientPrefix(context.Background()).ClientPrefix(clientPrefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.PutIsAvailableClientPrefix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIsAvailableClientPrefix`: EnvelopedString
	fmt.Fprintf(os.Stdout, "Response from `ClientAPI.PutIsAvailableClientPrefix`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIsAvailableClientPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientPrefix** | [**ClientPrefix**](ClientPrefix.md) | The client prefix to check. | 

### Return type

[**EnvelopedString**](EnvelopedString.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

