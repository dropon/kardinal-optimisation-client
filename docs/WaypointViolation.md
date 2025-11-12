# WaypointViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The violation&#39;s type. | 
**WayPointIndex** | Pointer to **int32** | The idx of the waypoint the violation applies to. | [optional] 
**LackingSkills** | Pointer to **[]string** | The list of skills needed to do the waypoint that the resource lacks. | [optional] 
**ForbiddenAssignments** | Pointer to [**[]ForbiddenAssignment**](ForbiddenAssignment.md) | The list of forbidden assignments that are violated with the waypoint assigned to the resource. | [optional] 
**ForbiddenWindow** | [**TimeWindow**](TimeWindow.md) | The violated forbidden window. | 
**ActualTime** | **string** | The time at which the waypoint is actually executed. | 
**Capacity** | **string** | The capacity&#39;s name the constraint is linked to. | 
**OverCapacity** | **float32** | The exceeding capacity. | 
**Name** | **string** | The name of the AtLeastOneValidCapacity constraint the violation is linked to. | 
**OverCapacities** | **map[string]float32** |  | 

## Methods

### NewWaypointViolation

`func NewWaypointViolation(type_ string, forbiddenWindow TimeWindow, actualTime string, capacity string, overCapacity float32, name string, overCapacities map[string]float32, ) *WaypointViolation`

NewWaypointViolation instantiates a new WaypointViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaypointViolationWithDefaults

`func NewWaypointViolationWithDefaults() *WaypointViolation`

NewWaypointViolationWithDefaults instantiates a new WaypointViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WaypointViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WaypointViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WaypointViolation) SetType(v string)`

SetType sets Type field to given value.


### GetWayPointIndex

`func (o *WaypointViolation) GetWayPointIndex() int32`

GetWayPointIndex returns the WayPointIndex field if non-nil, zero value otherwise.

### GetWayPointIndexOk

`func (o *WaypointViolation) GetWayPointIndexOk() (*int32, bool)`

GetWayPointIndexOk returns a tuple with the WayPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayPointIndex

`func (o *WaypointViolation) SetWayPointIndex(v int32)`

SetWayPointIndex sets WayPointIndex field to given value.

### HasWayPointIndex

`func (o *WaypointViolation) HasWayPointIndex() bool`

HasWayPointIndex returns a boolean if a field has been set.

### GetLackingSkills

`func (o *WaypointViolation) GetLackingSkills() []string`

GetLackingSkills returns the LackingSkills field if non-nil, zero value otherwise.

### GetLackingSkillsOk

`func (o *WaypointViolation) GetLackingSkillsOk() (*[]string, bool)`

GetLackingSkillsOk returns a tuple with the LackingSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLackingSkills

`func (o *WaypointViolation) SetLackingSkills(v []string)`

SetLackingSkills sets LackingSkills field to given value.

### HasLackingSkills

`func (o *WaypointViolation) HasLackingSkills() bool`

HasLackingSkills returns a boolean if a field has been set.

### GetForbiddenAssignments

`func (o *WaypointViolation) GetForbiddenAssignments() []ForbiddenAssignment`

GetForbiddenAssignments returns the ForbiddenAssignments field if non-nil, zero value otherwise.

### GetForbiddenAssignmentsOk

`func (o *WaypointViolation) GetForbiddenAssignmentsOk() (*[]ForbiddenAssignment, bool)`

GetForbiddenAssignmentsOk returns a tuple with the ForbiddenAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAssignments

`func (o *WaypointViolation) SetForbiddenAssignments(v []ForbiddenAssignment)`

SetForbiddenAssignments sets ForbiddenAssignments field to given value.

### HasForbiddenAssignments

`func (o *WaypointViolation) HasForbiddenAssignments() bool`

HasForbiddenAssignments returns a boolean if a field has been set.

### GetForbiddenWindow

`func (o *WaypointViolation) GetForbiddenWindow() TimeWindow`

GetForbiddenWindow returns the ForbiddenWindow field if non-nil, zero value otherwise.

### GetForbiddenWindowOk

`func (o *WaypointViolation) GetForbiddenWindowOk() (*TimeWindow, bool)`

GetForbiddenWindowOk returns a tuple with the ForbiddenWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenWindow

`func (o *WaypointViolation) SetForbiddenWindow(v TimeWindow)`

SetForbiddenWindow sets ForbiddenWindow field to given value.


### GetActualTime

`func (o *WaypointViolation) GetActualTime() string`

GetActualTime returns the ActualTime field if non-nil, zero value otherwise.

### GetActualTimeOk

`func (o *WaypointViolation) GetActualTimeOk() (*string, bool)`

GetActualTimeOk returns a tuple with the ActualTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTime

`func (o *WaypointViolation) SetActualTime(v string)`

SetActualTime sets ActualTime field to given value.


### GetCapacity

`func (o *WaypointViolation) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *WaypointViolation) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *WaypointViolation) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.


### GetOverCapacity

`func (o *WaypointViolation) GetOverCapacity() float32`

GetOverCapacity returns the OverCapacity field if non-nil, zero value otherwise.

### GetOverCapacityOk

`func (o *WaypointViolation) GetOverCapacityOk() (*float32, bool)`

GetOverCapacityOk returns a tuple with the OverCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCapacity

`func (o *WaypointViolation) SetOverCapacity(v float32)`

SetOverCapacity sets OverCapacity field to given value.


### GetName

`func (o *WaypointViolation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WaypointViolation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WaypointViolation) SetName(v string)`

SetName sets Name field to given value.


### GetOverCapacities

`func (o *WaypointViolation) GetOverCapacities() map[string]float32`

GetOverCapacities returns the OverCapacities field if non-nil, zero value otherwise.

### GetOverCapacitiesOk

`func (o *WaypointViolation) GetOverCapacitiesOk() (*map[string]float32, bool)`

GetOverCapacitiesOk returns a tuple with the OverCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCapacities

`func (o *WaypointViolation) SetOverCapacities(v map[string]float32)`

SetOverCapacities sets OverCapacities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


