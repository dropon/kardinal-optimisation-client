# PlanAdditionalConstraintsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Capacities** | **[]string** |  | 
**ResourceTags** | Pointer to **[]string** |  | [optional] 
**ResourceTag** | **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | 
**StopTag** | **string** | Stop tag. | 
**StopTags** | **[]string** | A pair of incompatible stop tags. | 
**Constraints** | [**[]PlanAdditionalConstraintsInner**](PlanAdditionalConstraintsInner.md) |  | 
**MaxGroupsByStopTag** | **map[string]int32** |  | 
**RemovalStrategy** | Pointer to [**RemovalStrategyType**](RemovalStrategyType.md) |  | [optional] 

## Methods

### NewPlanAdditionalConstraintsInner

`func NewPlanAdditionalConstraintsInner(type_ string, capacities []string, resourceTag string, stopTag string, stopTags []string, constraints []PlanAdditionalConstraintsInner, maxGroupsByStopTag map[string]int32, ) *PlanAdditionalConstraintsInner`

NewPlanAdditionalConstraintsInner instantiates a new PlanAdditionalConstraintsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanAdditionalConstraintsInnerWithDefaults

`func NewPlanAdditionalConstraintsInnerWithDefaults() *PlanAdditionalConstraintsInner`

NewPlanAdditionalConstraintsInnerWithDefaults instantiates a new PlanAdditionalConstraintsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PlanAdditionalConstraintsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlanAdditionalConstraintsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlanAdditionalConstraintsInner) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PlanAdditionalConstraintsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanAdditionalConstraintsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanAdditionalConstraintsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanAdditionalConstraintsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCapacities

`func (o *PlanAdditionalConstraintsInner) GetCapacities() []string`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *PlanAdditionalConstraintsInner) GetCapacitiesOk() (*[]string, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *PlanAdditionalConstraintsInner) SetCapacities(v []string)`

SetCapacities sets Capacities field to given value.


### GetResourceTags

`func (o *PlanAdditionalConstraintsInner) GetResourceTags() []string`

GetResourceTags returns the ResourceTags field if non-nil, zero value otherwise.

### GetResourceTagsOk

`func (o *PlanAdditionalConstraintsInner) GetResourceTagsOk() (*[]string, bool)`

GetResourceTagsOk returns a tuple with the ResourceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTags

`func (o *PlanAdditionalConstraintsInner) SetResourceTags(v []string)`

SetResourceTags sets ResourceTags field to given value.

### HasResourceTags

`func (o *PlanAdditionalConstraintsInner) HasResourceTags() bool`

HasResourceTags returns a boolean if a field has been set.

### GetResourceTag

`func (o *PlanAdditionalConstraintsInner) GetResourceTag() string`

GetResourceTag returns the ResourceTag field if non-nil, zero value otherwise.

### GetResourceTagOk

`func (o *PlanAdditionalConstraintsInner) GetResourceTagOk() (*string, bool)`

GetResourceTagOk returns a tuple with the ResourceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTag

`func (o *PlanAdditionalConstraintsInner) SetResourceTag(v string)`

SetResourceTag sets ResourceTag field to given value.


### GetStopTag

`func (o *PlanAdditionalConstraintsInner) GetStopTag() string`

GetStopTag returns the StopTag field if non-nil, zero value otherwise.

### GetStopTagOk

`func (o *PlanAdditionalConstraintsInner) GetStopTagOk() (*string, bool)`

GetStopTagOk returns a tuple with the StopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTag

`func (o *PlanAdditionalConstraintsInner) SetStopTag(v string)`

SetStopTag sets StopTag field to given value.


### GetStopTags

`func (o *PlanAdditionalConstraintsInner) GetStopTags() []string`

GetStopTags returns the StopTags field if non-nil, zero value otherwise.

### GetStopTagsOk

`func (o *PlanAdditionalConstraintsInner) GetStopTagsOk() (*[]string, bool)`

GetStopTagsOk returns a tuple with the StopTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTags

`func (o *PlanAdditionalConstraintsInner) SetStopTags(v []string)`

SetStopTags sets StopTags field to given value.


### GetConstraints

`func (o *PlanAdditionalConstraintsInner) GetConstraints() []PlanAdditionalConstraintsInner`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *PlanAdditionalConstraintsInner) GetConstraintsOk() (*[]PlanAdditionalConstraintsInner, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *PlanAdditionalConstraintsInner) SetConstraints(v []PlanAdditionalConstraintsInner)`

SetConstraints sets Constraints field to given value.


### GetMaxGroupsByStopTag

`func (o *PlanAdditionalConstraintsInner) GetMaxGroupsByStopTag() map[string]int32`

GetMaxGroupsByStopTag returns the MaxGroupsByStopTag field if non-nil, zero value otherwise.

### GetMaxGroupsByStopTagOk

`func (o *PlanAdditionalConstraintsInner) GetMaxGroupsByStopTagOk() (*map[string]int32, bool)`

GetMaxGroupsByStopTagOk returns a tuple with the MaxGroupsByStopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGroupsByStopTag

`func (o *PlanAdditionalConstraintsInner) SetMaxGroupsByStopTag(v map[string]int32)`

SetMaxGroupsByStopTag sets MaxGroupsByStopTag field to given value.


### GetRemovalStrategy

`func (o *PlanAdditionalConstraintsInner) GetRemovalStrategy() RemovalStrategyType`

GetRemovalStrategy returns the RemovalStrategy field if non-nil, zero value otherwise.

### GetRemovalStrategyOk

`func (o *PlanAdditionalConstraintsInner) GetRemovalStrategyOk() (*RemovalStrategyType, bool)`

GetRemovalStrategyOk returns a tuple with the RemovalStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalStrategy

`func (o *PlanAdditionalConstraintsInner) SetRemovalStrategy(v RemovalStrategyType)`

SetRemovalStrategy sets RemovalStrategy field to given value.

### HasRemovalStrategy

`func (o *PlanAdditionalConstraintsInner) HasRemovalStrategy() bool`

HasRemovalStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


