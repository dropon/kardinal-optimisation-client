# AdditionalConstraintRemovalStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**ResourceTags** | Pointer to **[]string** |  | [optional] 
**Capacities** | Pointer to **[]string** |  | [optional] 
**RemovalStrategy** | Pointer to [**RemovalStrategyType**](RemovalStrategyType.md) |  | [optional] 

## Methods

### NewAdditionalConstraintRemovalStrategy

`func NewAdditionalConstraintRemovalStrategy(type_ string, ) *AdditionalConstraintRemovalStrategy`

NewAdditionalConstraintRemovalStrategy instantiates a new AdditionalConstraintRemovalStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalConstraintRemovalStrategyWithDefaults

`func NewAdditionalConstraintRemovalStrategyWithDefaults() *AdditionalConstraintRemovalStrategy`

NewAdditionalConstraintRemovalStrategyWithDefaults instantiates a new AdditionalConstraintRemovalStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdditionalConstraintRemovalStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdditionalConstraintRemovalStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdditionalConstraintRemovalStrategy) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AdditionalConstraintRemovalStrategy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdditionalConstraintRemovalStrategy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdditionalConstraintRemovalStrategy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdditionalConstraintRemovalStrategy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceTags

`func (o *AdditionalConstraintRemovalStrategy) GetResourceTags() []string`

GetResourceTags returns the ResourceTags field if non-nil, zero value otherwise.

### GetResourceTagsOk

`func (o *AdditionalConstraintRemovalStrategy) GetResourceTagsOk() (*[]string, bool)`

GetResourceTagsOk returns a tuple with the ResourceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTags

`func (o *AdditionalConstraintRemovalStrategy) SetResourceTags(v []string)`

SetResourceTags sets ResourceTags field to given value.

### HasResourceTags

`func (o *AdditionalConstraintRemovalStrategy) HasResourceTags() bool`

HasResourceTags returns a boolean if a field has been set.

### GetCapacities

`func (o *AdditionalConstraintRemovalStrategy) GetCapacities() []string`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *AdditionalConstraintRemovalStrategy) GetCapacitiesOk() (*[]string, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *AdditionalConstraintRemovalStrategy) SetCapacities(v []string)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *AdditionalConstraintRemovalStrategy) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetRemovalStrategy

`func (o *AdditionalConstraintRemovalStrategy) GetRemovalStrategy() RemovalStrategyType`

GetRemovalStrategy returns the RemovalStrategy field if non-nil, zero value otherwise.

### GetRemovalStrategyOk

`func (o *AdditionalConstraintRemovalStrategy) GetRemovalStrategyOk() (*RemovalStrategyType, bool)`

GetRemovalStrategyOk returns a tuple with the RemovalStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalStrategy

`func (o *AdditionalConstraintRemovalStrategy) SetRemovalStrategy(v RemovalStrategyType)`

SetRemovalStrategy sets RemovalStrategy field to given value.

### HasRemovalStrategy

`func (o *AdditionalConstraintRemovalStrategy) HasRemovalStrategy() bool`

HasRemovalStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


