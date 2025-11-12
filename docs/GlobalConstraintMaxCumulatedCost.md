# GlobalConstraintMaxCumulatedCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Maximum** | Pointer to **float32** |  | [optional] 
**CostsByResourceTag** | [**map[string]Cost**](Cost.md) |  | 

## Methods

### NewGlobalConstraintMaxCumulatedCost

`func NewGlobalConstraintMaxCumulatedCost(type_ string, costsByResourceTag map[string]Cost, ) *GlobalConstraintMaxCumulatedCost`

NewGlobalConstraintMaxCumulatedCost instantiates a new GlobalConstraintMaxCumulatedCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalConstraintMaxCumulatedCostWithDefaults

`func NewGlobalConstraintMaxCumulatedCostWithDefaults() *GlobalConstraintMaxCumulatedCost`

NewGlobalConstraintMaxCumulatedCostWithDefaults instantiates a new GlobalConstraintMaxCumulatedCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GlobalConstraintMaxCumulatedCost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalConstraintMaxCumulatedCost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalConstraintMaxCumulatedCost) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *GlobalConstraintMaxCumulatedCost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalConstraintMaxCumulatedCost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalConstraintMaxCumulatedCost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalConstraintMaxCumulatedCost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaximum

`func (o *GlobalConstraintMaxCumulatedCost) GetMaximum() float32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *GlobalConstraintMaxCumulatedCost) GetMaximumOk() (*float32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *GlobalConstraintMaxCumulatedCost) SetMaximum(v float32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *GlobalConstraintMaxCumulatedCost) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetCostsByResourceTag

`func (o *GlobalConstraintMaxCumulatedCost) GetCostsByResourceTag() map[string]Cost`

GetCostsByResourceTag returns the CostsByResourceTag field if non-nil, zero value otherwise.

### GetCostsByResourceTagOk

`func (o *GlobalConstraintMaxCumulatedCost) GetCostsByResourceTagOk() (*map[string]Cost, bool)`

GetCostsByResourceTagOk returns a tuple with the CostsByResourceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostsByResourceTag

`func (o *GlobalConstraintMaxCumulatedCost) SetCostsByResourceTag(v map[string]Cost)`

SetCostsByResourceTag sets CostsByResourceTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


