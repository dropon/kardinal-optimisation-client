# ForbiddenAssignmentViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**WayPointIndex** | Pointer to **int32** | The idx of the waypoint the violation applies to. | [optional] 
**ForbiddenAssignments** | Pointer to [**[]ForbiddenAssignment**](ForbiddenAssignment.md) | The list of forbidden assignments that are violated with the waypoint assigned to the resource. | [optional] 

## Methods

### NewForbiddenAssignmentViolation

`func NewForbiddenAssignmentViolation(type_ string, ) *ForbiddenAssignmentViolation`

NewForbiddenAssignmentViolation instantiates a new ForbiddenAssignmentViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForbiddenAssignmentViolationWithDefaults

`func NewForbiddenAssignmentViolationWithDefaults() *ForbiddenAssignmentViolation`

NewForbiddenAssignmentViolationWithDefaults instantiates a new ForbiddenAssignmentViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ForbiddenAssignmentViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ForbiddenAssignmentViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ForbiddenAssignmentViolation) SetType(v string)`

SetType sets Type field to given value.


### GetWayPointIndex

`func (o *ForbiddenAssignmentViolation) GetWayPointIndex() int32`

GetWayPointIndex returns the WayPointIndex field if non-nil, zero value otherwise.

### GetWayPointIndexOk

`func (o *ForbiddenAssignmentViolation) GetWayPointIndexOk() (*int32, bool)`

GetWayPointIndexOk returns a tuple with the WayPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayPointIndex

`func (o *ForbiddenAssignmentViolation) SetWayPointIndex(v int32)`

SetWayPointIndex sets WayPointIndex field to given value.

### HasWayPointIndex

`func (o *ForbiddenAssignmentViolation) HasWayPointIndex() bool`

HasWayPointIndex returns a boolean if a field has been set.

### GetForbiddenAssignments

`func (o *ForbiddenAssignmentViolation) GetForbiddenAssignments() []ForbiddenAssignment`

GetForbiddenAssignments returns the ForbiddenAssignments field if non-nil, zero value otherwise.

### GetForbiddenAssignmentsOk

`func (o *ForbiddenAssignmentViolation) GetForbiddenAssignmentsOk() (*[]ForbiddenAssignment, bool)`

GetForbiddenAssignmentsOk returns a tuple with the ForbiddenAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAssignments

`func (o *ForbiddenAssignmentViolation) SetForbiddenAssignments(v []ForbiddenAssignment)`

SetForbiddenAssignments sets ForbiddenAssignments field to given value.

### HasForbiddenAssignments

`func (o *ForbiddenAssignmentViolation) HasForbiddenAssignments() bool`

HasForbiddenAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


