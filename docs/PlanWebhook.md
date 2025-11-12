# PlanWebhook

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
**PlanId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 

## Methods

### NewPlanWebhook

`func NewPlanWebhook() *PlanWebhook`

NewPlanWebhook instantiates a new PlanWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWebhookWithDefaults

`func NewPlanWebhookWithDefaults() *PlanWebhook`

NewPlanWebhookWithDefaults instantiates a new PlanWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanWebhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlanWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgencyId

`func (o *PlanWebhook) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *PlanWebhook) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *PlanWebhook) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.

### HasAgencyId

`func (o *PlanWebhook) HasAgencyId() bool`

HasAgencyId returns a boolean if a field has been set.

### GetUrl

`func (o *PlanWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PlanWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PlanWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PlanWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeaders

`func (o *PlanWebhook) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PlanWebhook) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PlanWebhook) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PlanWebhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetWithPlan

`func (o *PlanWebhook) GetWithPlan() bool`

GetWithPlan returns the WithPlan field if non-nil, zero value otherwise.

### GetWithPlanOk

`func (o *PlanWebhook) GetWithPlanOk() (*bool, bool)`

GetWithPlanOk returns a tuple with the WithPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPlan

`func (o *PlanWebhook) SetWithPlan(v bool)`

SetWithPlan sets WithPlan field to given value.

### HasWithPlan

`func (o *PlanWebhook) HasWithPlan() bool`

HasWithPlan returns a boolean if a field has been set.

### GetActive

`func (o *PlanWebhook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PlanWebhook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PlanWebhook) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PlanWebhook) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlanWebhook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanWebhook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanWebhook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlanWebhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlanWebhook) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlanWebhook) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlanWebhook) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlanWebhook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PlanWebhook) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PlanWebhook) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PlanWebhook) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PlanWebhook) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetPlanId

`func (o *PlanWebhook) GetPlanId() interface{}`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanWebhook) GetPlanIdOk() (*interface{}, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanWebhook) SetPlanId(v interface{})`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PlanWebhook) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


