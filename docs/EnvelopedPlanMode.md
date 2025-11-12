# EnvelopedPlanMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**PlanMode**](PlanMode.md) |  | [optional] [default to PLANMODE_STANDARD]
**AgencyId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanVersion** | Pointer to **int32** | The plan version. | [optional] [readonly] 

## Methods

### NewEnvelopedPlanMode

`func NewEnvelopedPlanMode() *EnvelopedPlanMode`

NewEnvelopedPlanMode instantiates a new EnvelopedPlanMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedPlanModeWithDefaults

`func NewEnvelopedPlanModeWithDefaults() *EnvelopedPlanMode`

NewEnvelopedPlanModeWithDefaults instantiates a new EnvelopedPlanMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EnvelopedPlanMode) GetItem() PlanMode`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EnvelopedPlanMode) GetItemOk() (*PlanMode, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EnvelopedPlanMode) SetItem(v PlanMode)`

SetItem sets Item field to given value.

### HasItem

`func (o *EnvelopedPlanMode) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetAgencyId

`func (o *EnvelopedPlanMode) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *EnvelopedPlanMode) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *EnvelopedPlanMode) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.

### HasAgencyId

`func (o *EnvelopedPlanMode) HasAgencyId() bool`

HasAgencyId returns a boolean if a field has been set.

### GetPlanId

`func (o *EnvelopedPlanMode) GetPlanId() interface{}`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *EnvelopedPlanMode) GetPlanIdOk() (*interface{}, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *EnvelopedPlanMode) SetPlanId(v interface{})`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *EnvelopedPlanMode) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanVersion

`func (o *EnvelopedPlanMode) GetPlanVersion() int32`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *EnvelopedPlanMode) GetPlanVersionOk() (*int32, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *EnvelopedPlanMode) SetPlanVersion(v int32)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *EnvelopedPlanMode) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


