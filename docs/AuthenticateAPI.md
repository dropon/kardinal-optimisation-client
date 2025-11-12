# \AuthenticateAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableMFAConfig**](AuthenticateAPI.md#DisableMFAConfig) | **Post** /auth/mfa/disable | Disable an OTP type in the actor&#39;s MFA config
[**GetActorMFAConfig**](AuthenticateAPI.md#GetActorMFAConfig) | **Get** /auth/mfa/config | Fetch the actor&#39;s MFA config
[**GetLoginMethods**](AuthenticateAPI.md#GetLoginMethods) | **Get** /login/methods | Retrieve the available login methods for a given username
[**GetPublicKey**](AuthenticateAPI.md#GetPublicKey) | **Get** /public_key | Get the public key which can be used to check the tokens
[**PostLogin**](AuthenticateAPI.md#PostLogin) | **Post** /login | Login (returns an OTP token if MFA is configured for the user)
[**PostLoginOTP**](AuthenticateAPI.md#PostLoginOTP) | **Post** /login/otp | Confirm login with OTP
[**PostLoginRefresh**](AuthenticateAPI.md#PostLoginRefresh) | **Post** /login/refresh | Refresh the access token
[**PostLoginWithAzureSSO**](AuthenticateAPI.md#PostLoginWithAzureSSO) | **Post** /login/sso/azure | Login with Azure SSO
[**PostLoginWithGoogleSSO**](AuthenticateAPI.md#PostLoginWithGoogleSSO) | **Post** /login/sso/google | Login with Google SSO
[**PostRequestPasswordToken**](AuthenticateAPI.md#PostRequestPasswordToken) | **Post** /auth/password/requestToken | Request a password token
[**PutMFAConfigPreferredType**](AuthenticateAPI.md#PutMFAConfigPreferredType) | **Put** /auth/mfa/preferredType | Sets the actor&#39;s preferred OTP type in his MFA config
[**RegenerateMFABackupCodes**](AuthenticateAPI.md#RegenerateMFABackupCodes) | **Put** /auth/mfa/regenerateBackupCodes | Regenerate the actor&#39;s backup codes
[**RequestMFAConfigUpdate**](AuthenticateAPI.md#RequestMFAConfigUpdate) | **Post** /auth/mfa/request | Request the update of the actor&#39;s MFA config
[**RequestNewLoginOTPCode**](AuthenticateAPI.md#RequestNewLoginOTPCode) | **Post** /login/resendOTP | Request a new OTP code for login
[**RequestNewMFAOTPCode**](AuthenticateAPI.md#RequestNewMFAOTPCode) | **Post** /auth/mfa/resendOTP | Request a new OTP code for the MFA config update validation
[**ResetPassword**](AuthenticateAPI.md#ResetPassword) | **Post** /auth/password/reset | Reset a password for a user
[**ValidateMFAConfigUpdate**](AuthenticateAPI.md#ValidateMFAConfigUpdate) | **Post** /auth/mfa/validate | Validate the update of the actor&#39;s MFA config
[**ValidatePasswordToken**](AuthenticateAPI.md#ValidatePasswordToken) | **Post** /auth/password/validateToken | Check if a password token is valid



## DisableMFAConfig

> EnvelopedMFAConfig DisableMFAConfig(ctx).DisableMFAConfigRequest(disableMFAConfigRequest).Execute()

Disable an OTP type in the actor's MFA config

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
	disableMFAConfigRequest := openapiclient.disableMFAConfig_request{MFAConfigRequestEmail: openapiclient.NewMFAConfigRequestEmail("OtpType_example", "martin.dupont@kardinal.ai", *openapiclient.NewPassword())} // DisableMFAConfigRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticateAPI.DisableMFAConfig(context.Background()).DisableMFAConfigRequest(disableMFAConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.DisableMFAConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableMFAConfig`: EnvelopedMFAConfig
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.DisableMFAConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableMFAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableMFAConfigRequest** | [**DisableMFAConfigRequest**](DisableMFAConfigRequest.md) |  | 

### Return type

[**EnvelopedMFAConfig**](EnvelopedMFAConfig.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActorMFAConfig

> EnvelopedMFAConfig GetActorMFAConfig(ctx).Execute()

Fetch the actor's MFA config

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
	resp, r, err := apiClient.AuthenticateAPI.GetActorMFAConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.GetActorMFAConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActorMFAConfig`: EnvelopedMFAConfig
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.GetActorMFAConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActorMFAConfigRequest struct via the builder pattern


### Return type

[**EnvelopedMFAConfig**](EnvelopedMFAConfig.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginMethods

> EnvelopedAuthenticationMethods GetLoginMethods(ctx).Username(username).Execute()

Retrieve the available login methods for a given username

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
	username := "username_example" // Username | The username. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticateAPI.GetLoginMethods(context.Background()).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.GetLoginMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoginMethods`: EnvelopedAuthenticationMethods
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.GetLoginMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **Username** | The username. | 

### Return type

[**EnvelopedAuthenticationMethods**](EnvelopedAuthenticationMethods.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicKey

> map[string]interface{} GetPublicKey(ctx).Execute()

Get the public key which can be used to check the tokens

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
	resp, r, err := apiClient.AuthenticateAPI.GetPublicKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.GetPublicKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.GetPublicKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicKeyRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

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
	resp, r, err := apiClient.AuthenticateAPI.PostLogin(context.Background()).Login(login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.PostLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLogin`: PostLogin200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.PostLogin`: %v\n", resp)
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
	resp, r, err := apiClient.AuthenticateAPI.PostLoginOTP(context.Background()).PostLoginOTPRequest(postLoginOTPRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.PostLoginOTP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLoginOTP`: PostLoginOTP200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.PostLoginOTP`: %v\n", resp)
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


## PostLoginRefresh

> PostLoginRefresh200Response PostLoginRefresh(ctx).PostLoginRefreshRequest(postLoginRefreshRequest).Execute()

Refresh the access token

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
	postLoginRefreshRequest := *openapiclient.NewPostLoginRefreshRequest() // PostLoginRefreshRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticateAPI.PostLoginRefresh(context.Background()).PostLoginRefreshRequest(postLoginRefreshRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.PostLoginRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLoginRefresh`: PostLoginRefresh200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.PostLoginRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLoginRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postLoginRefreshRequest** | [**PostLoginRefreshRequest**](PostLoginRefreshRequest.md) |  | 

### Return type

[**PostLoginRefresh200Response**](PostLoginRefresh200Response.md)

### Authorization

[refresh_token](../README.md#refresh_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLoginWithAzureSSO

> PostLoginOTP200Response PostLoginWithAzureSSO(ctx).AzureSSOLogin(azureSSOLogin).Execute()

Login with Azure SSO

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
	azureSSOLogin := *openapiclient.NewAzureSSOLogin("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Imk2bEdrM0ZaenhSY1ViMkMzbkVRN3N5SEpsWSJ9.eyJhdWQiOiI2ZTc0MTcyYi1iZTU2LTQ4NDMtOWZmNC1lNjZhMzliYjEyZTMiLCJpc3MiOiJodHRwczovL2xvZ2luLm1pY3Jvc29mdG9ubGluZS5jb20vNzJmOTg4YmYtODZmMS00MWFmLTkxYWItMmQ3Y2QwMTFkYjQ3L3YyLjAiLCJpYXQiOjE1MzcyMzEwNDgsIm5iZiI6MTUzNzIzMTA0OCwiZXhwIjoxNTM3MjM0OTQ4LCJhaW8iOiJBWFFBaS84SUFBQUF0QWFaTG8zQ2hNaWY2S09udHRSQjdlQnE0L0RjY1F6amNKR3hQWXkvQzNqRGFOR3hYZDZ3TklJVkdSZ2hOUm53SjFsT2NBbk5aY2p2a295ckZ4Q3R0djMzMTQwUmlvT0ZKNGJDQ0dWdW9DYWcxdU9UVDIyMjIyZ0h3TFBZUS91Zjc5UVgrMEtJaWpkcm1wNjlSY3R6bVE9PSIsImF6cCI6IjZlNzQxNzJiLWJlNTYtNDg0My05ZmY0LWU2NmEzOWJiMTJlMyIsImF6cGFjciI6IjAiLCJuYW1lIjoiQWJlIExpbmNvbG4iLCJvaWQiOiI2OTAyMjJiZS1mZjFhLTRkNTYtYWJkMS03ZTRmN2QzOGU0NzQiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJhYmVsaUBtaWNyb3NvZnQuY29tIiwicmgiOiJJIiwic2NwIjoiYWNjZXNzX2FzX3VzZXIiLCJzdWIiOiJIS1pwZmFIeVdhZGVPb3VZbGl0anJJLUtmZlRtMjIyWDVyclYzeERxZktRIiwidGlkIjoiNzJmOTg4YmYtODZmMS00MWFmLTkxYWItMmQ3Y2QwMTFkYjQ3IiwidXRpIjoiZnFpQnFYTFBqMGVRYTgyUy1JWUZBQSIsInZlciI6IjIuMCJ9.pj4N-w_3Us9DrBLfpCt", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjFMVE16YWtpaGlSbGFfOHoyQkVKVlhlV01xbyJ9.eyJ2ZXIiOiIyLjAiLCJpc3MiOiJodHRwczovL2xvZ2luLm1pY3Jvc29mdG9ubGluZS5jb20vOTEyMjA0MGQtNmM2Ny00YzViLWIxMTItMzZhMzA0YjY2ZGFkL3YyLjAiLCJzdWIiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFJa3pxRlZyU2FTYUZIeTc4MmJidGFRIiwiYXVkIjoiNmNiMDQwMTgtYTNmNS00NmE3LWI5OTUtOTQwYzc4ZjVhZWYzIiwiZXhwIjoxNTM2MzYxNDExLCJpYXQiOjE1MzYyNzQ3MTEsIm5iZiI6MTUzNjI3NDcxMSwibmFtZSI6IkFiZSBMaW5jb2xuIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiQWJlTGlAbWljcm9zb2Z0LmNvbSIsIm9pZCI6IjAwMDAwMDAwLTAwMDAtMDAwMC02NmYzLTMzMzJlY2E3ZWE4MSIsInRpZCI6IjkxMjIwNDBkLTZjNjctNGM1Yi1iMTEyLTM2YTMwNGI2NmRhZCIsIm5vbmNlIjoiMTIzNTIzIiwiYWlvIjoiRGYyVVZYTDFpeCFsTUNXTVNPSkJjRmF0emNHZnZGR2hqS3Y4cTVnMHg3MzJkUjVNQjVCaXN2R1FPN1lXQnlqZDhpUURMcSFlR2JJRGFreXA1bW5PcmNkcUhlWVNubHRlcFFtUnA2QUlaOGpZIn0.1AFWW-Ck5nROwSlltm7GzZvDwUkqvhSQpm55TQsmVo9Y59cLhRXpvB8n-55HCr9Z6G_31_UbeUkoz612I2j_Sm9FFShSDDjoaLQr54CreGIJvjtmS3EkK9a7SJBbcpL1MpUtlfygow39tFjY7EVNW9plWUvRrTgVk7lYLprvfzw-CIqw3gHC-T7IK_m_xkr08INERBtaecwhTeN4chPC4W3jdmw_lIxzC48YoQ0dB1L9-ImX98Egypfrlbm0IBL5spFzL6JDZIRRJOu8vecJvj1mq-IUhGt0MacxX8jdxYLP-KUu2d9MbNKpCKJuZ7p8gwTL5B7NlUdh_dmSviPWrw") // AzureSSOLogin |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticateAPI.PostLoginWithAzureSSO(context.Background()).AzureSSOLogin(azureSSOLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.PostLoginWithAzureSSO``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLoginWithAzureSSO`: PostLoginOTP200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.PostLoginWithAzureSSO`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLoginWithAzureSSORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureSSOLogin** | [**AzureSSOLogin**](AzureSSOLogin.md) |  | 

### Return type

[**PostLoginOTP200Response**](PostLoginOTP200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLoginWithGoogleSSO

> PostLoginOTP200Response PostLoginWithGoogleSSO(ctx).GoogleSSOLogin(googleSSOLogin).Execute()

Login with Google SSO

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
	googleSSOLogin := *openapiclient.NewGoogleSSOLogin("ya28.a0AfB_byCbEgO5rOeu73KBdAv0ZeiEFz7xzgxtdCzailCy9ew6BCgKUaiheRxDJ1Do8HbiXVxMcm5jyo_5fV_yAtYKhSLKayMvGH0n00E-NtOgtH5myIhpc7heq5zZ788XSzoHxJ9aSWR_y1Vr7-zWiRWe_YANqnmEo6FaaCgYKAckERASSFQHGX2MilpI4NsOBgWYkOo1zKgJPmg0171") // GoogleSSOLogin |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticateAPI.PostLoginWithGoogleSSO(context.Background()).GoogleSSOLogin(googleSSOLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.PostLoginWithGoogleSSO``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLoginWithGoogleSSO`: PostLoginOTP200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.PostLoginWithGoogleSSO`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLoginWithGoogleSSORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **googleSSOLogin** | [**GoogleSSOLogin**](GoogleSSOLogin.md) |  | 

### Return type

[**PostLoginOTP200Response**](PostLoginOTP200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRequestPasswordToken

> PostRequestPasswordToken(ctx).OriginHeader(originHeader).UsernameRequest(usernameRequest).Execute()

Request a password token

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
	originHeader := "originHeader_example" // string | The origin of the request. When an email needs to be sent, the \"Origin\" header value is used (after validation) to correctly generate the hyperlinks inside the email body.  (optional)
	usernameRequest := *openapiclient.NewUsernameRequest() // UsernameRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticateAPI.PostRequestPasswordToken(context.Background()).OriginHeader(originHeader).UsernameRequest(usernameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.PostRequestPasswordToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRequestPasswordTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **originHeader** | **string** | The origin of the request. When an email needs to be sent, the \&quot;Origin\&quot; header value is used (after validation) to correctly generate the hyperlinks inside the email body.  | 
 **usernameRequest** | [**UsernameRequest**](UsernameRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMFAConfigPreferredType

> EnvelopedOTPType PutMFAConfigPreferredType(ctx).Body(body).Execute()

Sets the actor's preferred OTP type in his MFA config

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
	body := string(987) // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticateAPI.PutMFAConfigPreferredType(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.PutMFAConfigPreferredType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMFAConfigPreferredType`: EnvelopedOTPType
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.PutMFAConfigPreferredType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutMFAConfigPreferredTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**EnvelopedOTPType**](EnvelopedOTPType.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateMFABackupCodes

> EnvelopedBackupCodes RegenerateMFABackupCodes(ctx).MFAConfigRegenerateBackupCodes(mFAConfigRegenerateBackupCodes).Execute()

Regenerate the actor's backup codes

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
	mFAConfigRegenerateBackupCodes := *openapiclient.NewMFAConfigRegenerateBackupCodes(*openapiclient.NewPassword()) // MFAConfigRegenerateBackupCodes |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticateAPI.RegenerateMFABackupCodes(context.Background()).MFAConfigRegenerateBackupCodes(mFAConfigRegenerateBackupCodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.RegenerateMFABackupCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateMFABackupCodes`: EnvelopedBackupCodes
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.RegenerateMFABackupCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateMFABackupCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mFAConfigRegenerateBackupCodes** | [**MFAConfigRegenerateBackupCodes**](MFAConfigRegenerateBackupCodes.md) |  | 

### Return type

[**EnvelopedBackupCodes**](EnvelopedBackupCodes.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestMFAConfigUpdate

> RequestMFAConfigUpdate200Response RequestMFAConfigUpdate(ctx).DisableMFAConfigRequest(disableMFAConfigRequest).Execute()

Request the update of the actor's MFA config

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
	disableMFAConfigRequest := openapiclient.disableMFAConfig_request{MFAConfigRequestEmail: openapiclient.NewMFAConfigRequestEmail("OtpType_example", "martin.dupont@kardinal.ai", *openapiclient.NewPassword())} // DisableMFAConfigRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticateAPI.RequestMFAConfigUpdate(context.Background()).DisableMFAConfigRequest(disableMFAConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.RequestMFAConfigUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestMFAConfigUpdate`: RequestMFAConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.RequestMFAConfigUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestMFAConfigUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableMFAConfigRequest** | [**DisableMFAConfigRequest**](DisableMFAConfigRequest.md) |  | 

### Return type

[**RequestMFAConfigUpdate200Response**](RequestMFAConfigUpdate200Response.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestNewLoginOTPCode

> RequestNewLoginOTPCode(ctx).Execute()

Request a new OTP code for login

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
	r, err := apiClient.AuthenticateAPI.RequestNewLoginOTPCode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.RequestNewLoginOTPCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRequestNewLoginOTPCodeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[otp_token](../README.md#otp_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestNewMFAOTPCode

> RequestNewMFAOTPCode(ctx).RequestNewMFAOTPCodeRequest(requestNewMFAOTPCodeRequest).Execute()

Request a new OTP code for the MFA config update validation

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
	requestNewMFAOTPCodeRequest := openapiclient.requestNewMFAOTPCode_request{MFAConfigResendEmail: openapiclient.NewMFAConfigResendEmail("OtpType_example", "martin.dupont@kardinal.ai")} // RequestNewMFAOTPCodeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticateAPI.RequestNewMFAOTPCode(context.Background()).RequestNewMFAOTPCodeRequest(requestNewMFAOTPCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.RequestNewMFAOTPCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestNewMFAOTPCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestNewMFAOTPCodeRequest** | [**RequestNewMFAOTPCodeRequest**](RequestNewMFAOTPCodeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[otp_token](../README.md#otp_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> ResetPassword(ctx).PasswordRequest(passwordRequest).Execute()

Reset a password for a user

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
	passwordRequest := *openapiclient.NewPasswordRequest() // PasswordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticateAPI.ResetPassword(context.Background()).PasswordRequest(passwordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.ResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordRequest** | [**PasswordRequest**](PasswordRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[password_token](../README.md#password_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMFAConfigUpdate

> EnvelopedMFAConfig ValidateMFAConfigUpdate(ctx).ValidateMFAConfigUpdateRequest(validateMFAConfigUpdateRequest).Execute()

Validate the update of the actor's MFA config

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
	validateMFAConfigUpdateRequest := openapiclient.validateMFAConfigUpdate_request{MFAConfigValidationEmail: openapiclient.NewMFAConfigValidationEmail("OtpType_example", "martin.dupont@kardinal.ai", "403852")} // ValidateMFAConfigUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticateAPI.ValidateMFAConfigUpdate(context.Background()).ValidateMFAConfigUpdateRequest(validateMFAConfigUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.ValidateMFAConfigUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateMFAConfigUpdate`: EnvelopedMFAConfig
	fmt.Fprintf(os.Stdout, "Response from `AuthenticateAPI.ValidateMFAConfigUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateMFAConfigUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateMFAConfigUpdateRequest** | [**ValidateMFAConfigUpdateRequest**](ValidateMFAConfigUpdateRequest.md) |  | 

### Return type

[**EnvelopedMFAConfig**](EnvelopedMFAConfig.md)

### Authorization

[otp_token](../README.md#otp_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePasswordToken

> ValidatePasswordToken(ctx).TokenRequest(tokenRequest).Execute()

Check if a password token is valid

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
	tokenRequest := *openapiclient.NewTokenRequest() // TokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticateAPI.ValidatePasswordToken(context.Background()).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateAPI.ValidatePasswordToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePasswordTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

