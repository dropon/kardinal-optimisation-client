# EnvelopedSMAOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ScoredTimeWindow**](ScoredTimeWindow.md) |  | [optional] 
**AgencyId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanVersion** | Pointer to **int32** | The plan version. | [optional] [readonly] 

## Methods

### NewEnvelopedSMAOutput

`func NewEnvelopedSMAOutput() *EnvelopedSMAOutput`

NewEnvelopedSMAOutput instantiates a new EnvelopedSMAOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedSMAOutputWithDefaults

`func NewEnvelopedSMAOutputWithDefaults() *EnvelopedSMAOutput`

NewEnvelopedSMAOutputWithDefaults instantiates a new EnvelopedSMAOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EnvelopedSMAOutput) GetItems() []ScoredTimeWindow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnvelopedSMAOutput) GetItemsOk() (*[]ScoredTimeWindow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnvelopedSMAOutput) SetItems(v []ScoredTimeWindow)`

SetItems sets Items field to given value.

### HasItems

`func (o *EnvelopedSMAOutput) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAgencyId

`func (o *EnvelopedSMAOutput) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *EnvelopedSMAOutput) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *EnvelopedSMAOutput) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.

### HasAgencyId

`func (o *EnvelopedSMAOutput) HasAgencyId() bool`

HasAgencyId returns a boolean if a field has been set.

### GetPlanId

`func (o *EnvelopedSMAOutput) GetPlanId() interface{}`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *EnvelopedSMAOutput) GetPlanIdOk() (*interface{}, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *EnvelopedSMAOutput) SetPlanId(v interface{})`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *EnvelopedSMAOutput) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanVersion

`func (o *EnvelopedSMAOutput) GetPlanVersion() int32`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *EnvelopedSMAOutput) GetPlanVersionOk() (*int32, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *EnvelopedSMAOutput) SetPlanVersion(v int32)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *EnvelopedSMAOutput) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


