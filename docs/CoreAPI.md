# \CoreAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlanSolution**](CoreAPI.md#GetPlanSolution) | **Get** /agencies/{agencyId}/plans/{planId}/solution | Retrieve a plan solution
[**PostLogin**](CoreAPI.md#PostLogin) | **Post** /login | Login (returns an OTP token if MFA is configured for the user)
[**PostLoginOTP**](CoreAPI.md#PostLoginOTP) | **Post** /login/otp | Confirm login with OTP
[**PutPlan**](CoreAPI.md#PutPlan) | **Put** /agencies/{agencyId}/plans/{planId} | Create or update a plan



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
	resp, r, err := apiClient.CoreAPI.GetPlanSolution(context.Background(), agencyId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.GetPlanSolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanSolution`: EnvelopedSolution
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.GetPlanSolution`: %v\n", resp)
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


## PostLogin

> PostLogin200Response PostLogin(ctx).Login(login).Execute()

Login (returns an OTP token if MFA is configured for the user)

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
	login := *openapiclient.NewLogin(*openapiclient.NewUsername(), *openapiclient.NewPassword()) // Login |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.PostLogin(context.Background()).Login(login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.PostLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLogin`: PostLogin200Response
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.PostLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | [**Login**](Login.md) |  | 

### Return type

[**PostLogin200Response**](PostLogin200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLoginOTP

> PostLoginOTP200Response PostLoginOTP(ctx).PostLoginOTPRequest(postLoginOTPRequest).Execute()

Confirm login with OTP

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
	postLoginOTPRequest := openapiclient.postLoginOTP_request{LoginBackupCodeInput: openapiclient.NewLoginBackupCodeInput("E0ZE97NY1Z6WAW1L")} // PostLoginOTPRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.PostLoginOTP(context.Background()).PostLoginOTPRequest(postLoginOTPRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.PostLoginOTP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLoginOTP`: PostLoginOTP200Response
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.PostLoginOTP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLoginOTPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postLoginOTPRequest** | [**PostLoginOTPRequest**](PostLoginOTPRequest.md) |  | 

### Return type

[**PostLoginOTP200Response**](PostLoginOTP200Response.md)

### Authorization

[otp_token](../README.md#otp_token)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.CoreAPI.PutPlan(context.Background(), agencyId, planId).Force(force).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.PutPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlan`: EnvelopedPlan
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.PutPlan`: %v\n", resp)
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

