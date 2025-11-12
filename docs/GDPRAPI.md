# \GDPRAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchGDPRApprove**](GDPRAPI.md#PatchGDPRApprove) | **Patch** /gdpr/approve | Approve a GDPR policy



## PatchGDPRApprove

> EnvelopedLoginAccessOutput PatchGDPRApprove(ctx).GDPRApproval(gDPRApproval).Execute()

Approve a GDPR policy

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
	gDPRApproval := *openapiclient.NewGDPRApproval("Version_example") // GDPRApproval |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GDPRAPI.PatchGDPRApprove(context.Background()).GDPRApproval(gDPRApproval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GDPRAPI.PatchGDPRApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchGDPRApprove`: EnvelopedLoginAccessOutput
	fmt.Fprintf(os.Stdout, "Response from `GDPRAPI.PatchGDPRApprove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchGDPRApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gDPRApproval** | [**GDPRApproval**](GDPRApproval.md) |  | 

### Return type

[**EnvelopedLoginAccessOutput**](EnvelopedLoginAccessOutput.md)

### Authorization

[gdpr_token](../README.md#gdpr_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

