# PlanObjectivesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Direction** | [**OptimizationDirection**](OptimizationDirection.md) |  | 
**CostsByResourceTag** | [**map[string]Cost**](Cost.md) |  | 

## Methods

### NewPlanObjectivesInner

`func NewPlanObjectivesInner(type_ string, name string, direction OptimizationDirection, costsByResourceTag map[string]Cost, ) *PlanObjectivesInner`

NewPlanObjectivesInner instantiates a new PlanObjectivesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanObjectivesInnerWithDefaults

`func NewPlanObjectivesInnerWithDefaults() *PlanObjectivesInner`

NewPlanObjectivesInnerWithDefaults instantiates a new PlanObjectivesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PlanObjectivesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlanObjectivesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlanObjectivesInner) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PlanObjectivesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanObjectivesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanObjectivesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDirection

`func (o *PlanObjectivesInner) GetDirection() OptimizationDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PlanObjectivesInner) GetDirectionOk() (*OptimizationDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PlanObjectivesInner) SetDirection(v OptimizationDirection)`

SetDirection sets Direction field to given value.


### GetCostsByResourceTag

`func (o *PlanObjectivesInner) GetCostsByResourceTag() map[string]Cost`

GetCostsByResourceTag returns the CostsByResourceTag field if non-nil, zero value otherwise.

### GetCostsByResourceTagOk

`func (o *PlanObjectivesInner) GetCostsByResourceTagOk() (*map[string]Cost, bool)`

GetCostsByResourceTagOk returns a tuple with the CostsByResourceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostsByResourceTag

`func (o *PlanObjectivesInner) SetCostsByResourceTag(v map[string]Cost)`

SetCostsByResourceTag sets CostsByResourceTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


