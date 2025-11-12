# PlanWebhookResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | Pointer to **string** | Universally Unique Identifier. | [optional] [readonly] 
**AgencyId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**StatusCode** | Pointer to **float32** | The HTTP status code from the webhook&#39;s response. | [optional] 
**Timestamp** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**InputSize** | Pointer to **float32** | The webhook&#39;s input payload size in KB. | [optional] 
**OutputSize** | Pointer to **float32** | The webhook&#39;s output payload size in KB. | [optional] 
**TryNumber** | Pointer to **float32** | The try number of calling the webhook:  - 1: This is the first time we are calling the webhook.  - greater than 1: Calling the webhook failed at least once and we are re-trying.  | [optional] 
**PlanId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanVersion** | Pointer to **int32** | The plan version. | [optional] [readonly] 

## Methods

### NewPlanWebhookResult

`func NewPlanWebhookResult() *PlanWebhookResult`

NewPlanWebhookResult instantiates a new PlanWebhookResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWebhookResultWithDefaults

`func NewPlanWebhookResultWithDefaults() *PlanWebhookResult`

NewPlanWebhookResultWithDefaults instantiates a new PlanWebhookResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *PlanWebhookResult) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *PlanWebhookResult) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *PlanWebhookResult) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *PlanWebhookResult) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetAgencyId

`func (o *PlanWebhookResult) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *PlanWebhookResult) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *PlanWebhookResult) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.

### HasAgencyId

`func (o *PlanWebhookResult) HasAgencyId() bool`

HasAgencyId returns a boolean if a field has been set.

### GetStatusCode

`func (o *PlanWebhookResult) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *PlanWebhookResult) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *PlanWebhookResult) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *PlanWebhookResult) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *PlanWebhookResult) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PlanWebhookResult) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PlanWebhookResult) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PlanWebhookResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetInputSize

`func (o *PlanWebhookResult) GetInputSize() float32`

GetInputSize returns the InputSize field if non-nil, zero value otherwise.

### GetInputSizeOk

`func (o *PlanWebhookResult) GetInputSizeOk() (*float32, bool)`

GetInputSizeOk returns a tuple with the InputSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSize

`func (o *PlanWebhookResult) SetInputSize(v float32)`

SetInputSize sets InputSize field to given value.

### HasInputSize

`func (o *PlanWebhookResult) HasInputSize() bool`

HasInputSize returns a boolean if a field has been set.

### GetOutputSize

`func (o *PlanWebhookResult) GetOutputSize() float32`

GetOutputSize returns the OutputSize field if non-nil, zero value otherwise.

### GetOutputSizeOk

`func (o *PlanWebhookResult) GetOutputSizeOk() (*float32, bool)`

GetOutputSizeOk returns a tuple with the OutputSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSize

`func (o *PlanWebhookResult) SetOutputSize(v float32)`

SetOutputSize sets OutputSize field to given value.

### HasOutputSize

`func (o *PlanWebhookResult) HasOutputSize() bool`

HasOutputSize returns a boolean if a field has been set.

### GetTryNumber

`func (o *PlanWebhookResult) GetTryNumber() float32`

GetTryNumber returns the TryNumber field if non-nil, zero value otherwise.

### GetTryNumberOk

`func (o *PlanWebhookResult) GetTryNumberOk() (*float32, bool)`

GetTryNumberOk returns a tuple with the TryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryNumber

`func (o *PlanWebhookResult) SetTryNumber(v float32)`

SetTryNumber sets TryNumber field to given value.

### HasTryNumber

`func (o *PlanWebhookResult) HasTryNumber() bool`

HasTryNumber returns a boolean if a field has been set.

### GetPlanId

`func (o *PlanWebhookResult) GetPlanId() interface{}`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanWebhookResult) GetPlanIdOk() (*interface{}, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanWebhookResult) SetPlanId(v interface{})`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PlanWebhookResult) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanVersion

`func (o *PlanWebhookResult) GetPlanVersion() int32`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *PlanWebhookResult) GetPlanVersionOk() (*int32, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *PlanWebhookResult) SetPlanVersion(v int32)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *PlanWebhookResult) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


