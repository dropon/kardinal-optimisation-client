# AdditionalConstraintMaxStopTagGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**ResourceTags** | Pointer to **[]string** |  | [optional] 
**MaxGroupsByStopTag** | **map[string]int32** |  | 

## Methods

### NewAdditionalConstraintMaxStopTagGroups

`func NewAdditionalConstraintMaxStopTagGroups(type_ string, maxGroupsByStopTag map[string]int32, ) *AdditionalConstraintMaxStopTagGroups`

NewAdditionalConstraintMaxStopTagGroups instantiates a new AdditionalConstraintMaxStopTagGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalConstraintMaxStopTagGroupsWithDefaults

`func NewAdditionalConstraintMaxStopTagGroupsWithDefaults() *AdditionalConstraintMaxStopTagGroups`

NewAdditionalConstraintMaxStopTagGroupsWithDefaults instantiates a new AdditionalConstraintMaxStopTagGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdditionalConstraintMaxStopTagGroups) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdditionalConstraintMaxStopTagGroups) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdditionalConstraintMaxStopTagGroups) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AdditionalConstraintMaxStopTagGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdditionalConstraintMaxStopTagGroups) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdditionalConstraintMaxStopTagGroups) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdditionalConstraintMaxStopTagGroups) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceTags

`func (o *AdditionalConstraintMaxStopTagGroups) GetResourceTags() []string`

GetResourceTags returns the ResourceTags field if non-nil, zero value otherwise.

### GetResourceTagsOk

`func (o *AdditionalConstraintMaxStopTagGroups) GetResourceTagsOk() (*[]string, bool)`

GetResourceTagsOk returns a tuple with the ResourceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTags

`func (o *AdditionalConstraintMaxStopTagGroups) SetResourceTags(v []string)`

SetResourceTags sets ResourceTags field to given value.

### HasResourceTags

`func (o *AdditionalConstraintMaxStopTagGroups) HasResourceTags() bool`

HasResourceTags returns a boolean if a field has been set.

### GetMaxGroupsByStopTag

`func (o *AdditionalConstraintMaxStopTagGroups) GetMaxGroupsByStopTag() map[string]int32`

GetMaxGroupsByStopTag returns the MaxGroupsByStopTag field if non-nil, zero value otherwise.

### GetMaxGroupsByStopTagOk

`func (o *AdditionalConstraintMaxStopTagGroups) GetMaxGroupsByStopTagOk() (*map[string]int32, bool)`

GetMaxGroupsByStopTagOk returns a tuple with the MaxGroupsByStopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGroupsByStopTag

`func (o *AdditionalConstraintMaxStopTagGroups) SetMaxGroupsByStopTag(v map[string]int32)`

SetMaxGroupsByStopTag sets MaxGroupsByStopTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


