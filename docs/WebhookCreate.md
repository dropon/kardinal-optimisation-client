# WebhookCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The secure webhook URL (must be https). | 
**Headers** | Pointer to **map[string]string** | A map of header key-value pairs. | [optional] 
**WithPlan** | Pointer to **bool** | Specifies whether we should also provide the plan in the webhook&#39;s payload. | [optional] [default to false]
**Ttl** | Pointer to **string** | Defines the webhook&#39;s time to live in iso 8601 duration format. The value should be at least 30 minutes and at most 14 days.  | [optional] [default to "PT30M"]

## Methods

### NewWebhookCreate

`func NewWebhookCreate(url string, ) *WebhookCreate`

NewWebhookCreate instantiates a new WebhookCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCreateWithDefaults

`func NewWebhookCreateWithDefaults() *WebhookCreate`

NewWebhookCreateWithDefaults instantiates a new WebhookCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookCreate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *WebhookCreate) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookCreate) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookCreate) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookCreate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetWithPlan

`func (o *WebhookCreate) GetWithPlan() bool`

GetWithPlan returns the WithPlan field if non-nil, zero value otherwise.

### GetWithPlanOk

`func (o *WebhookCreate) GetWithPlanOk() (*bool, bool)`

GetWithPlanOk returns a tuple with the WithPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPlan

`func (o *WebhookCreate) SetWithPlan(v bool)`

SetWithPlan sets WithPlan field to given value.

### HasWithPlan

`func (o *WebhookCreate) HasWithPlan() bool`

HasWithPlan returns a boolean if a field has been set.

### GetTtl

`func (o *WebhookCreate) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *WebhookCreate) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *WebhookCreate) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *WebhookCreate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


