# \SmartMeetingAppointmentAPI

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSmaMultiPlanScoring**](SmartMeetingAppointmentAPI.md#PostSmaMultiPlanScoring) | **Post** /agencies/{agencyId}/timeWindowScoring | Post an order to obtain time windows for a smart meeting appointment in multiple plans of an agency
[**PostSmaScoring**](SmartMeetingAppointmentAPI.md#PostSmaScoring) | **Post** /agencies/{agencyId}/plans/{planId}/timeWindowScoring | Post an order to obtain time windows for a smart meeting appointment



## PostSmaMultiPlanScoring

> EnvelopedSMAOutput PostSmaMultiPlanScoring(ctx, agencyId).SMAMultiPlanInput(sMAMultiPlanInput).Execute()

Post an order to obtain time windows for a smart meeting appointment in multiple plans of an agency

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
	sMAMultiPlanInput := *openapiclient.NewSMAMultiPlanInput(*openapiclient.NewSMAOrder("Id_example", []openapiclient.Stop{openapiclient.Stop{AlternativesStop: openapiclient.NewAlternativesStop("Type_example", []openapiclient.SingleStop{*openapiclient.NewSingleStop("Id_example", *openapiclient.NewPosition(float32(123), float32(123)))})}}), []openapiclient.TimeWindow{*openapiclient.NewTimeWindow("2019-11-15T12:34:56Z", "2019-11-15T12:34:56Z")}, []string{"PlanIds_example"}) // SMAMultiPlanInput | The smart meeting appointment multi plans input. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartMeetingAppointmentAPI.PostSmaMultiPlanScoring(context.Background(), agencyId).SMAMultiPlanInput(sMAMultiPlanInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartMeetingAppointmentAPI.PostSmaMultiPlanScoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSmaMultiPlanScoring`: EnvelopedSMAOutput
	fmt.Fprintf(os.Stdout, "Response from `SmartMeetingAppointmentAPI.PostSmaMultiPlanScoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSmaMultiPlanScoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sMAMultiPlanInput** | [**SMAMultiPlanInput**](SMAMultiPlanInput.md) | The smart meeting appointment multi plans input. | 

### Return type

[**EnvelopedSMAOutput**](EnvelopedSMAOutput.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSmaScoring

> EnvelopedSMAOutput PostSmaScoring(ctx, agencyId, planId).SMAInput(sMAInput).Execute()

Post an order to obtain time windows for a smart meeting appointment

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
	sMAInput := *openapiclient.NewSMAInput(*openapiclient.NewSMAOrder("Id_example", []openapiclient.Stop{openapiclient.Stop{AlternativesStop: openapiclient.NewAlternativesStop("Type_example", []openapiclient.SingleStop{*openapiclient.NewSingleStop("Id_example", *openapiclient.NewPosition(float32(123), float32(123)))})}}), []openapiclient.TimeWindow{*openapiclient.NewTimeWindow("2019-11-15T12:34:56Z", "2019-11-15T12:34:56Z")}) // SMAInput | The smart meeting appointment input. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartMeetingAppointmentAPI.PostSmaScoring(context.Background(), agencyId, planId).SMAInput(sMAInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartMeetingAppointmentAPI.PostSmaScoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSmaScoring`: EnvelopedSMAOutput
	fmt.Fprintf(os.Stdout, "Response from `SmartMeetingAppointmentAPI.PostSmaScoring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agencyId** | [**AgencyId**](.md) | The agency id. | 
**planId** | **string** | The plan id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSmaScoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sMAInput** | [**SMAInput**](SMAInput.md) | The smart meeting appointment input. | 

### Return type

[**EnvelopedSMAOutput**](EnvelopedSMAOutput.md)

### Authorization

[access_token](../README.md#access_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

