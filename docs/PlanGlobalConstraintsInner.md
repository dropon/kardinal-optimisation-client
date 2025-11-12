# PlanGlobalConstraintsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Maximum** | Pointer to **float32** |  | [optional] 
**CostsByResourceTag** | [**map[string]Cost**](Cost.md) |  | 

## Methods

### NewPlanGlobalConstraintsInner

`func NewPlanGlobalConstraintsInner(type_ string, costsByResourceTag map[string]Cost, ) *PlanGlobalConstraintsInner`

NewPlanGlobalConstraintsInner instantiates a new PlanGlobalConstraintsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanGlobalConstraintsInnerWithDefaults

`func NewPlanGlobalConstraintsInnerWithDefaults() *PlanGlobalConstraintsInner`

NewPlanGlobalConstraintsInnerWithDefaults instantiates a new PlanGlobalConstraintsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PlanGlobalConstraintsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlanGlobalConstraintsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlanGlobalConstraintsInner) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PlanGlobalConstraintsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanGlobalConstraintsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanGlobalConstraintsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanGlobalConstraintsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaximum

`func (o *PlanGlobalConstraintsInner) GetMaximum() float32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *PlanGlobalConstraintsInner) GetMaximumOk() (*float32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *PlanGlobalConstraintsInner) SetMaximum(v float32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *PlanGlobalConstraintsInner) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetCostsByResourceTag

`func (o *PlanGlobalConstraintsInner) GetCostsByResourceTag() map[string]Cost`

GetCostsByResourceTag returns the CostsByResourceTag field if non-nil, zero value otherwise.

### GetCostsByResourceTagOk

`func (o *PlanGlobalConstraintsInner) GetCostsByResourceTagOk() (*map[string]Cost, bool)`

GetCostsByResourceTagOk returns a tuple with the CostsByResourceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostsByResourceTag

`func (o *PlanGlobalConstraintsInner) SetCostsByResourceTag(v map[string]Cost)`

SetCostsByResourceTag sets CostsByResourceTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


