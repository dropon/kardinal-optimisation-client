# AgencyWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Universally Unique Identifier. | [optional] [readonly] 
**AgencyId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**Url** | Pointer to **string** | The secure webhook URL (must be https). | [optional] 
**Headers** | Pointer to **map[string]string** | A map of header key-value pairs. | [optional] 
**WithPlan** | Pointer to **bool** | Specifies whether we should also provide the plan in the webhook&#39;s payload. | [optional] [default to false]
**Active** | Pointer to **bool** | Specifies whether the webhook is active or not. | [optional] [default to true]
**CreatedAt** | Pointer to **string** | The webhook&#39;s creation date. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The webhook&#39;s last update date. | [optional] [readonly] 
**ExpiresAt** | Pointer to **string** | The webhook&#39;s expiration date. | [optional] [readonly] 

## Methods

### NewAgencyWebhook

`func NewAgencyWebhook() *AgencyWebhook`

NewAgencyWebhook instantiates a new AgencyWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgencyWebhookWithDefaults

`func NewAgencyWebhookWithDefaults() *AgencyWebhook`

NewAgencyWebhookWithDefaults instantiates a new AgencyWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgencyWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgencyWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgencyWebhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgencyWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgencyId

`func (o *AgencyWebhook) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *AgencyWebhook) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *AgencyWebhook) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.

### HasAgencyId

`func (o *AgencyWebhook) HasAgencyId() bool`

HasAgencyId returns a boolean if a field has been set.

### GetUrl

`func (o *AgencyWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AgencyWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AgencyWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AgencyWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeaders

`func (o *AgencyWebhook) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AgencyWebhook) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AgencyWebhook) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AgencyWebhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetWithPlan

`func (o *AgencyWebhook) GetWithPlan() bool`

GetWithPlan returns the WithPlan field if non-nil, zero value otherwise.

### GetWithPlanOk

`func (o *AgencyWebhook) GetWithPlanOk() (*bool, bool)`

GetWithPlanOk returns a tuple with the WithPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPlan

`func (o *AgencyWebhook) SetWithPlan(v bool)`

SetWithPlan sets WithPlan field to given value.

### HasWithPlan

`func (o *AgencyWebhook) HasWithPlan() bool`

HasWithPlan returns a boolean if a field has been set.

### GetActive

`func (o *AgencyWebhook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AgencyWebhook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AgencyWebhook) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AgencyWebhook) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgencyWebhook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgencyWebhook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgencyWebhook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgencyWebhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgencyWebhook) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgencyWebhook) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgencyWebhook) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgencyWebhook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AgencyWebhook) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AgencyWebhook) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AgencyWebhook) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AgencyWebhook) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


