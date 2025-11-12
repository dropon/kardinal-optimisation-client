# \UserAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](UserAPI.md#DeleteUser) | **Delete** /users/{id} | Deletes a single user
[**GetMe**](UserAPI.md#GetMe) | **Get** /users/me | Retrieves the current user
[**GetUser**](UserAPI.md#GetUser) | **Get** /users/{id} | Retrieves a single user
[**GetUsers**](UserAPI.md#GetUsers) | **Get** /users | Retrieves a collection of users
[**PostUser**](UserAPI.md#PostUser) | **Post** /users | Creates a single user
[**PostUserApi**](UserAPI.md#PostUserApi) | **Post** /users/api | Creates a user with simpleApi or complexApi usage
[**PutIsAvailableUsername**](UserAPI.md#PutIsAvailableUsername) | **Put** /users/isAvailable/username | Checks the availability of a username
[**PutUser**](UserAPI.md#PutUser) | **Put** /users/{id} | Updates a single user



## DeleteUser

> DeleteUser(ctx, id).Execute()

Deletes a single user

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
	id := "id_example" // string | Internal identifier: the id which was returned by the POST request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Internal identifier: the id which was returned by the POST request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetMe

> EnvelopedUser GetMe(ctx).Execute()

Retrieves the current user

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
	resp, r, err := apiClient.UserAPI.GetMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMe`: EnvelopedUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeRequest struct via the builder pattern


### Return type

[**EnvelopedUser**](EnvelopedUser.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> EnvelopedUser GetUser(ctx, id).Execute()

Retrieves a single user

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
	id := "id_example" // string | Internal identifier: the id which was returned by the POST request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: EnvelopedUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Internal identifier: the id which was returned by the POST request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvelopedUser**](EnvelopedUser.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> EnvelopedUsers GetUsers(ctx).Execute()

Retrieves a collection of users

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
	resp, r, err := apiClient.UserAPI.GetUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: EnvelopedUsers
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


### Return type

[**EnvelopedUsers**](EnvelopedUsers.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUser

> EnvelopedUser PostUser(ctx).OriginHeader(originHeader).User(user).Execute()

Creates a single user

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
	user := *openapiclient.NewUser(*openapiclient.NewUsername()) // User | A user entity contains the following informations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PostUser(context.Background()).OriginHeader(originHeader).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PostUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUser`: EnvelopedUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PostUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **originHeader** | **string** | The origin of the request. When an email needs to be sent, the \&quot;Origin\&quot; header value is used (after validation) to correctly generate the hyperlinks inside the email body.  | 
 **user** | [**User**](User.md) | A user entity contains the following informations. | 

### Return type

[**EnvelopedUser**](EnvelopedUser.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserApi

> EnvelopedUserAPIOutput PostUserApi(ctx).OriginHeader(originHeader).UserAPIInput(userAPIInput).Execute()

Creates a user with simpleApi or complexApi usage

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
	userAPIInput := *openapiclient.NewUserAPIInput("martin.dupont@kardinal.ai", "Company_example") // UserAPIInput | An input which allows to create a user with simpleApi or complexApi usage, the associated client, and the client hierarchy.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PostUserApi(context.Background()).OriginHeader(originHeader).UserAPIInput(userAPIInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PostUserApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUserApi`: EnvelopedUserAPIOutput
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PostUserApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUserApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **originHeader** | **string** | The origin of the request. When an email needs to be sent, the \&quot;Origin\&quot; header value is used (after validation) to correctly generate the hyperlinks inside the email body.  | 
 **userAPIInput** | [**UserAPIInput**](UserAPIInput.md) | An input which allows to create a user with simpleApi or complexApi usage, the associated client, and the client hierarchy.  | 

### Return type

[**EnvelopedUserAPIOutput**](EnvelopedUserAPIOutput.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIsAvailableUsername

> EnvelopedString PutIsAvailableUsername(ctx).Body(body).Execute()

Checks the availability of a username

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
	body := "body_example" // string | The username to check. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PutIsAvailableUsername(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PutIsAvailableUsername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIsAvailableUsername`: EnvelopedString
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PutIsAvailableUsername`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIsAvailableUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | The username to check. | 

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


## PutUser

> EnvelopedUser PutUser(ctx, id).User(user).Execute()

Updates a single user

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
	id := "id_example" // string | Internal identifier: the id which was returned by the POST request.
	user := *openapiclient.NewUser(*openapiclient.NewUsername()) // User | The user to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PutUser(context.Background(), id).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PutUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUser`: EnvelopedUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PutUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Internal identifier: the id which was returned by the POST request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) | The user to update. | 

### Return type

[**EnvelopedUser**](EnvelopedUser.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

