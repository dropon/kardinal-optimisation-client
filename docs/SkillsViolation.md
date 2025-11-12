# SkillsViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**WayPointIndex** | Pointer to **int32** | The idx of the waypoint the violation applies to. | [optional] 
**LackingSkills** | Pointer to **[]string** | The list of skills needed to do the waypoint that the resource lacks. | [optional] 

## Methods

### NewSkillsViolation

`func NewSkillsViolation(type_ string, ) *SkillsViolation`

NewSkillsViolation instantiates a new SkillsViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillsViolationWithDefaults

`func NewSkillsViolationWithDefaults() *SkillsViolation`

NewSkillsViolationWithDefaults instantiates a new SkillsViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkillsViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkillsViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkillsViolation) SetType(v string)`

SetType sets Type field to given value.


### GetWayPointIndex

`func (o *SkillsViolation) GetWayPointIndex() int32`

GetWayPointIndex returns the WayPointIndex field if non-nil, zero value otherwise.

### GetWayPointIndexOk

`func (o *SkillsViolation) GetWayPointIndexOk() (*int32, bool)`

GetWayPointIndexOk returns a tuple with the WayPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayPointIndex

`func (o *SkillsViolation) SetWayPointIndex(v int32)`

SetWayPointIndex sets WayPointIndex field to given value.

### HasWayPointIndex

`func (o *SkillsViolation) HasWayPointIndex() bool`

HasWayPointIndex returns a boolean if a field has been set.

### GetLackingSkills

`func (o *SkillsViolation) GetLackingSkills() []string`

GetLackingSkills returns the LackingSkills field if non-nil, zero value otherwise.

### GetLackingSkillsOk

`func (o *SkillsViolation) GetLackingSkillsOk() (*[]string, bool)`

GetLackingSkillsOk returns a tuple with the LackingSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLackingSkills

`func (o *SkillsViolation) SetLackingSkills(v []string)`

SetLackingSkills sets LackingSkills field to given value.

### HasLackingSkills

`func (o *SkillsViolation) HasLackingSkills() bool`

HasLackingSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


