# MaxStopTagGroupsViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** | The name of the MaxStopTagGroups constraint the violation is linked to. | [optional] 
**ExceededNbGroupsByStopTag** | [**map[string]ExceededNbGroups**](ExceededNbGroups.md) | The exceeded number of groups by stop tag. | 

## Methods

### NewMaxStopTagGroupsViolation

`func NewMaxStopTagGroupsViolation(type_ string, exceededNbGroupsByStopTag map[string]ExceededNbGroups, ) *MaxStopTagGroupsViolation`

NewMaxStopTagGroupsViolation instantiates a new MaxStopTagGroupsViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxStopTagGroupsViolationWithDefaults

`func NewMaxStopTagGroupsViolationWithDefaults() *MaxStopTagGroupsViolation`

NewMaxStopTagGroupsViolationWithDefaults instantiates a new MaxStopTagGroupsViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MaxStopTagGroupsViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MaxStopTagGroupsViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MaxStopTagGroupsViolation) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *MaxStopTagGroupsViolation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaxStopTagGroupsViolation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaxStopTagGroupsViolation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MaxStopTagGroupsViolation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExceededNbGroupsByStopTag

`func (o *MaxStopTagGroupsViolation) GetExceededNbGroupsByStopTag() map[string]ExceededNbGroups`

GetExceededNbGroupsByStopTag returns the ExceededNbGroupsByStopTag field if non-nil, zero value otherwise.

### GetExceededNbGroupsByStopTagOk

`func (o *MaxStopTagGroupsViolation) GetExceededNbGroupsByStopTagOk() (*map[string]ExceededNbGroups, bool)`

GetExceededNbGroupsByStopTagOk returns a tuple with the ExceededNbGroupsByStopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceededNbGroupsByStopTag

`func (o *MaxStopTagGroupsViolation) SetExceededNbGroupsByStopTag(v map[string]ExceededNbGroups)`

SetExceededNbGroupsByStopTag sets ExceededNbGroupsByStopTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


