# TourViolation

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
**Name** | **string** | The name of the RemovalStrategy constraint the violation is linked to. | 
**OverCapacities** | **map[string]float32** |  | 
**ExpectedSuccessiveStopIds** | **[]string** | The stop ids in the expected order. | 
**InterleavedStopIds** | Pointer to **[]string** | The stop ids which are not expected and interleaved between the order stops. | [optional] 
**ExpectedOrderedStopIds** | **[]string** | The stop ids in the expected order. | 
**MissedStopIds** | Pointer to **[]string** | The stop ids missing from the tour. | [optional] 
**BadlyOrderedStopIds** | Pointer to **[]string** | The stop ids badly ordered in the tour. | [optional] 
**OrderId** | **string** | The order id with the MaxStopSpan constraint violation. | 
**Delay** | **string** | The delay that exceeds the resource&#39;s TimeWindow. | 
**MaxWorkingDuration** | **string** | The expected maximum working duration. | 
**ExceededDuration** | **string** | The duration that exceeds the maximum working duration. | 
**MaxDistanceInKm** | **float32** | The expected maximum distance, in km. | 
**ExceededDistanceInKm** | **float32** | The distance that exceeds the maximum distance, in km. | 
**Incompatibilities** | **[][]string** | Lists of stop ids for each incompatible tag. | 
**Tags** | **[]string** | A pair of incompatible stop tags. | 
**ViolatedConstraints** | [**[]TourViolation**](TourViolation.md) | Lists of sub-constraints violations. | 
**ExceededNbGroupsByStopTag** | [**map[string]ExceededNbGroups**](ExceededNbGroups.md) | The exceeded number of groups by stop tag. | 
**RemovalStrategyType** | [**RemovalStrategyType**](RemovalStrategyType.md) |  | 

## Methods

### NewTourViolation

`func NewTourViolation(type_ string, forbiddenWindow TimeWindow, actualTime string, capacity string, overCapacity float32, name string, overCapacities map[string]float32, expectedSuccessiveStopIds []string, expectedOrderedStopIds []string, orderId string, delay string, maxWorkingDuration string, exceededDuration string, maxDistanceInKm float32, exceededDistanceInKm float32, incompatibilities [][]string, tags []string, violatedConstraints []TourViolation, exceededNbGroupsByStopTag map[string]ExceededNbGroups, removalStrategyType RemovalStrategyType, ) *TourViolation`

NewTourViolation instantiates a new TourViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTourViolationWithDefaults

`func NewTourViolationWithDefaults() *TourViolation`

NewTourViolationWithDefaults instantiates a new TourViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TourViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TourViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TourViolation) SetType(v string)`

SetType sets Type field to given value.


### GetWayPointIndex

`func (o *TourViolation) GetWayPointIndex() int32`

GetWayPointIndex returns the WayPointIndex field if non-nil, zero value otherwise.

### GetWayPointIndexOk

`func (o *TourViolation) GetWayPointIndexOk() (*int32, bool)`

GetWayPointIndexOk returns a tuple with the WayPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayPointIndex

`func (o *TourViolation) SetWayPointIndex(v int32)`

SetWayPointIndex sets WayPointIndex field to given value.

### HasWayPointIndex

`func (o *TourViolation) HasWayPointIndex() bool`

HasWayPointIndex returns a boolean if a field has been set.

### GetLackingSkills

`func (o *TourViolation) GetLackingSkills() []string`

GetLackingSkills returns the LackingSkills field if non-nil, zero value otherwise.

### GetLackingSkillsOk

`func (o *TourViolation) GetLackingSkillsOk() (*[]string, bool)`

GetLackingSkillsOk returns a tuple with the LackingSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLackingSkills

`func (o *TourViolation) SetLackingSkills(v []string)`

SetLackingSkills sets LackingSkills field to given value.

### HasLackingSkills

`func (o *TourViolation) HasLackingSkills() bool`

HasLackingSkills returns a boolean if a field has been set.

### GetForbiddenAssignments

`func (o *TourViolation) GetForbiddenAssignments() []ForbiddenAssignment`

GetForbiddenAssignments returns the ForbiddenAssignments field if non-nil, zero value otherwise.

### GetForbiddenAssignmentsOk

`func (o *TourViolation) GetForbiddenAssignmentsOk() (*[]ForbiddenAssignment, bool)`

GetForbiddenAssignmentsOk returns a tuple with the ForbiddenAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAssignments

`func (o *TourViolation) SetForbiddenAssignments(v []ForbiddenAssignment)`

SetForbiddenAssignments sets ForbiddenAssignments field to given value.

### HasForbiddenAssignments

`func (o *TourViolation) HasForbiddenAssignments() bool`

HasForbiddenAssignments returns a boolean if a field has been set.

### GetForbiddenWindow

`func (o *TourViolation) GetForbiddenWindow() TimeWindow`

GetForbiddenWindow returns the ForbiddenWindow field if non-nil, zero value otherwise.

### GetForbiddenWindowOk

`func (o *TourViolation) GetForbiddenWindowOk() (*TimeWindow, bool)`

GetForbiddenWindowOk returns a tuple with the ForbiddenWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenWindow

`func (o *TourViolation) SetForbiddenWindow(v TimeWindow)`

SetForbiddenWindow sets ForbiddenWindow field to given value.


### GetActualTime

`func (o *TourViolation) GetActualTime() string`

GetActualTime returns the ActualTime field if non-nil, zero value otherwise.

### GetActualTimeOk

`func (o *TourViolation) GetActualTimeOk() (*string, bool)`

GetActualTimeOk returns a tuple with the ActualTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTime

`func (o *TourViolation) SetActualTime(v string)`

SetActualTime sets ActualTime field to given value.


### GetCapacity

`func (o *TourViolation) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *TourViolation) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *TourViolation) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.


### GetOverCapacity

`func (o *TourViolation) GetOverCapacity() float32`

GetOverCapacity returns the OverCapacity field if non-nil, zero value otherwise.

### GetOverCapacityOk

`func (o *TourViolation) GetOverCapacityOk() (*float32, bool)`

GetOverCapacityOk returns a tuple with the OverCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCapacity

`func (o *TourViolation) SetOverCapacity(v float32)`

SetOverCapacity sets OverCapacity field to given value.


### GetName

`func (o *TourViolation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TourViolation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TourViolation) SetName(v string)`

SetName sets Name field to given value.


### GetOverCapacities

`func (o *TourViolation) GetOverCapacities() map[string]float32`

GetOverCapacities returns the OverCapacities field if non-nil, zero value otherwise.

### GetOverCapacitiesOk

`func (o *TourViolation) GetOverCapacitiesOk() (*map[string]float32, bool)`

GetOverCapacitiesOk returns a tuple with the OverCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCapacities

`func (o *TourViolation) SetOverCapacities(v map[string]float32)`

SetOverCapacities sets OverCapacities field to given value.


### GetExpectedSuccessiveStopIds

`func (o *TourViolation) GetExpectedSuccessiveStopIds() []string`

GetExpectedSuccessiveStopIds returns the ExpectedSuccessiveStopIds field if non-nil, zero value otherwise.

### GetExpectedSuccessiveStopIdsOk

`func (o *TourViolation) GetExpectedSuccessiveStopIdsOk() (*[]string, bool)`

GetExpectedSuccessiveStopIdsOk returns a tuple with the ExpectedSuccessiveStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSuccessiveStopIds

`func (o *TourViolation) SetExpectedSuccessiveStopIds(v []string)`

SetExpectedSuccessiveStopIds sets ExpectedSuccessiveStopIds field to given value.


### GetInterleavedStopIds

`func (o *TourViolation) GetInterleavedStopIds() []string`

GetInterleavedStopIds returns the InterleavedStopIds field if non-nil, zero value otherwise.

### GetInterleavedStopIdsOk

`func (o *TourViolation) GetInterleavedStopIdsOk() (*[]string, bool)`

GetInterleavedStopIdsOk returns a tuple with the InterleavedStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterleavedStopIds

`func (o *TourViolation) SetInterleavedStopIds(v []string)`

SetInterleavedStopIds sets InterleavedStopIds field to given value.

### HasInterleavedStopIds

`func (o *TourViolation) HasInterleavedStopIds() bool`

HasInterleavedStopIds returns a boolean if a field has been set.

### GetExpectedOrderedStopIds

`func (o *TourViolation) GetExpectedOrderedStopIds() []string`

GetExpectedOrderedStopIds returns the ExpectedOrderedStopIds field if non-nil, zero value otherwise.

### GetExpectedOrderedStopIdsOk

`func (o *TourViolation) GetExpectedOrderedStopIdsOk() (*[]string, bool)`

GetExpectedOrderedStopIdsOk returns a tuple with the ExpectedOrderedStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOrderedStopIds

`func (o *TourViolation) SetExpectedOrderedStopIds(v []string)`

SetExpectedOrderedStopIds sets ExpectedOrderedStopIds field to given value.


### GetMissedStopIds

`func (o *TourViolation) GetMissedStopIds() []string`

GetMissedStopIds returns the MissedStopIds field if non-nil, zero value otherwise.

### GetMissedStopIdsOk

`func (o *TourViolation) GetMissedStopIdsOk() (*[]string, bool)`

GetMissedStopIdsOk returns a tuple with the MissedStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissedStopIds

`func (o *TourViolation) SetMissedStopIds(v []string)`

SetMissedStopIds sets MissedStopIds field to given value.

### HasMissedStopIds

`func (o *TourViolation) HasMissedStopIds() bool`

HasMissedStopIds returns a boolean if a field has been set.

### GetBadlyOrderedStopIds

`func (o *TourViolation) GetBadlyOrderedStopIds() []string`

GetBadlyOrderedStopIds returns the BadlyOrderedStopIds field if non-nil, zero value otherwise.

### GetBadlyOrderedStopIdsOk

`func (o *TourViolation) GetBadlyOrderedStopIdsOk() (*[]string, bool)`

GetBadlyOrderedStopIdsOk returns a tuple with the BadlyOrderedStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadlyOrderedStopIds

`func (o *TourViolation) SetBadlyOrderedStopIds(v []string)`

SetBadlyOrderedStopIds sets BadlyOrderedStopIds field to given value.

### HasBadlyOrderedStopIds

`func (o *TourViolation) HasBadlyOrderedStopIds() bool`

HasBadlyOrderedStopIds returns a boolean if a field has been set.

### GetOrderId

`func (o *TourViolation) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *TourViolation) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *TourViolation) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetDelay

`func (o *TourViolation) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *TourViolation) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *TourViolation) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetMaxWorkingDuration

`func (o *TourViolation) GetMaxWorkingDuration() string`

GetMaxWorkingDuration returns the MaxWorkingDuration field if non-nil, zero value otherwise.

### GetMaxWorkingDurationOk

`func (o *TourViolation) GetMaxWorkingDurationOk() (*string, bool)`

GetMaxWorkingDurationOk returns a tuple with the MaxWorkingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkingDuration

`func (o *TourViolation) SetMaxWorkingDuration(v string)`

SetMaxWorkingDuration sets MaxWorkingDuration field to given value.


### GetExceededDuration

`func (o *TourViolation) GetExceededDuration() string`

GetExceededDuration returns the ExceededDuration field if non-nil, zero value otherwise.

### GetExceededDurationOk

`func (o *TourViolation) GetExceededDurationOk() (*string, bool)`

GetExceededDurationOk returns a tuple with the ExceededDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceededDuration

`func (o *TourViolation) SetExceededDuration(v string)`

SetExceededDuration sets ExceededDuration field to given value.


### GetMaxDistanceInKm

`func (o *TourViolation) GetMaxDistanceInKm() float32`

GetMaxDistanceInKm returns the MaxDistanceInKm field if non-nil, zero value otherwise.

### GetMaxDistanceInKmOk

`func (o *TourViolation) GetMaxDistanceInKmOk() (*float32, bool)`

GetMaxDistanceInKmOk returns a tuple with the MaxDistanceInKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistanceInKm

`func (o *TourViolation) SetMaxDistanceInKm(v float32)`

SetMaxDistanceInKm sets MaxDistanceInKm field to given value.


### GetExceededDistanceInKm

`func (o *TourViolation) GetExceededDistanceInKm() float32`

GetExceededDistanceInKm returns the ExceededDistanceInKm field if non-nil, zero value otherwise.

### GetExceededDistanceInKmOk

`func (o *TourViolation) GetExceededDistanceInKmOk() (*float32, bool)`

GetExceededDistanceInKmOk returns a tuple with the ExceededDistanceInKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceededDistanceInKm

`func (o *TourViolation) SetExceededDistanceInKm(v float32)`

SetExceededDistanceInKm sets ExceededDistanceInKm field to given value.


### GetIncompatibilities

`func (o *TourViolation) GetIncompatibilities() [][]string`

GetIncompatibilities returns the Incompatibilities field if non-nil, zero value otherwise.

### GetIncompatibilitiesOk

`func (o *TourViolation) GetIncompatibilitiesOk() (*[][]string, bool)`

GetIncompatibilitiesOk returns a tuple with the Incompatibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibilities

`func (o *TourViolation) SetIncompatibilities(v [][]string)`

SetIncompatibilities sets Incompatibilities field to given value.


### GetTags

`func (o *TourViolation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TourViolation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TourViolation) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetViolatedConstraints

`func (o *TourViolation) GetViolatedConstraints() []TourViolation`

GetViolatedConstraints returns the ViolatedConstraints field if non-nil, zero value otherwise.

### GetViolatedConstraintsOk

`func (o *TourViolation) GetViolatedConstraintsOk() (*[]TourViolation, bool)`

GetViolatedConstraintsOk returns a tuple with the ViolatedConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatedConstraints

`func (o *TourViolation) SetViolatedConstraints(v []TourViolation)`

SetViolatedConstraints sets ViolatedConstraints field to given value.


### GetExceededNbGroupsByStopTag

`func (o *TourViolation) GetExceededNbGroupsByStopTag() map[string]ExceededNbGroups`

GetExceededNbGroupsByStopTag returns the ExceededNbGroupsByStopTag field if non-nil, zero value otherwise.

### GetExceededNbGroupsByStopTagOk

`func (o *TourViolation) GetExceededNbGroupsByStopTagOk() (*map[string]ExceededNbGroups, bool)`

GetExceededNbGroupsByStopTagOk returns a tuple with the ExceededNbGroupsByStopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceededNbGroupsByStopTag

`func (o *TourViolation) SetExceededNbGroupsByStopTag(v map[string]ExceededNbGroups)`

SetExceededNbGroupsByStopTag sets ExceededNbGroupsByStopTag field to given value.


### GetRemovalStrategyType

`func (o *TourViolation) GetRemovalStrategyType() RemovalStrategyType`

GetRemovalStrategyType returns the RemovalStrategyType field if non-nil, zero value otherwise.

### GetRemovalStrategyTypeOk

`func (o *TourViolation) GetRemovalStrategyTypeOk() (*RemovalStrategyType, bool)`

GetRemovalStrategyTypeOk returns a tuple with the RemovalStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalStrategyType

`func (o *TourViolation) SetRemovalStrategyType(v RemovalStrategyType)`

SetRemovalStrategyType sets RemovalStrategyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


