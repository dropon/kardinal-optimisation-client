# EnvelopedOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**Order**](Order.md) |  | [optional] 
**AgencyId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanVersion** | Pointer to **int32** | The plan version. | [optional] [readonly] 

## Methods

### NewEnvelopedOrder

`func NewEnvelopedOrder() *EnvelopedOrder`

NewEnvelopedOrder instantiates a new EnvelopedOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedOrderWithDefaults

`func NewEnvelopedOrderWithDefaults() *EnvelopedOrder`

NewEnvelopedOrderWithDefaults instantiates a new EnvelopedOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EnvelopedOrder) GetItem() Order`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EnvelopedOrder) GetItemOk() (*Order, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EnvelopedOrder) SetItem(v Order)`

SetItem sets Item field to given value.

### HasItem

`func (o *EnvelopedOrder) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetAgencyId

`func (o *EnvelopedOrder) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *EnvelopedOrder) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *EnvelopedOrder) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.

### HasAgencyId

`func (o *EnvelopedOrder) HasAgencyId() bool`

HasAgencyId returns a boolean if a field has been set.

### GetPlanId

`func (o *EnvelopedOrder) GetPlanId() interface{}`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *EnvelopedOrder) GetPlanIdOk() (*interface{}, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *EnvelopedOrder) SetPlanId(v interface{})`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *EnvelopedOrder) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanVersion

`func (o *EnvelopedOrder) GetPlanVersion() int32`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *EnvelopedOrder) GetPlanVersionOk() (*int32, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *EnvelopedOrder) SetPlanVersion(v int32)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *EnvelopedOrder) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


