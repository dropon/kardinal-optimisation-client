# WebhookUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The secure webhook URL (must be https). | 
**Headers** | Pointer to **map[string]string** | A map of header key-value pairs. | [optional] 
**WithPlan** | Pointer to **bool** | Specifies whether we should also provide the plan in the webhook&#39;s payload. | [optional] [default to false]
**Active** | Pointer to **bool** | Specifies whether the webhook is active or not. | [optional] [default to true]

## Methods

### NewWebhookUpdate

`func NewWebhookUpdate(url string, ) *WebhookUpdate`

NewWebhookUpdate instantiates a new WebhookUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdateWithDefaults

`func NewWebhookUpdateWithDefaults() *WebhookUpdate`

NewWebhookUpdateWithDefaults instantiates a new WebhookUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookUpdate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookUpdate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookUpdate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *WebhookUpdate) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookUpdate) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookUpdate) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookUpdate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetWithPlan

`func (o *WebhookUpdate) GetWithPlan() bool`

GetWithPlan returns the WithPlan field if non-nil, zero value otherwise.

### GetWithPlanOk

`func (o *WebhookUpdate) GetWithPlanOk() (*bool, bool)`

GetWithPlanOk returns a tuple with the WithPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPlan

`func (o *WebhookUpdate) SetWithPlan(v bool)`

SetWithPlan sets WithPlan field to given value.

### HasWithPlan

`func (o *WebhookUpdate) HasWithPlan() bool`

HasWithPlan returns a boolean if a field has been set.

### GetActive

`func (o *WebhookUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


