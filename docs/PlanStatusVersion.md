# PlanStatusVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaitingVersion** | Pointer to **int32** | The version currently waiting. | [optional] [readonly] 
**RunningVersion** | Pointer to **int32** | The version currently running. | [optional] [readonly] 

## Methods

### NewPlanStatusVersion

`func NewPlanStatusVersion() *PlanStatusVersion`

NewPlanStatusVersion instantiates a new PlanStatusVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanStatusVersionWithDefaults

`func NewPlanStatusVersionWithDefaults() *PlanStatusVersion`

NewPlanStatusVersionWithDefaults instantiates a new PlanStatusVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaitingVersion

`func (o *PlanStatusVersion) GetWaitingVersion() int32`

GetWaitingVersion returns the WaitingVersion field if non-nil, zero value otherwise.

### GetWaitingVersionOk

`func (o *PlanStatusVersion) GetWaitingVersionOk() (*int32, bool)`

GetWaitingVersionOk returns a tuple with the WaitingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingVersion

`func (o *PlanStatusVersion) SetWaitingVersion(v int32)`

SetWaitingVersion sets WaitingVersion field to given value.

### HasWaitingVersion

`func (o *PlanStatusVersion) HasWaitingVersion() bool`

HasWaitingVersion returns a boolean if a field has been set.

### GetRunningVersion

`func (o *PlanStatusVersion) GetRunningVersion() int32`

GetRunningVersion returns the RunningVersion field if non-nil, zero value otherwise.

### GetRunningVersionOk

`func (o *PlanStatusVersion) GetRunningVersionOk() (*int32, bool)`

GetRunningVersionOk returns a tuple with the RunningVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningVersion

`func (o *PlanStatusVersion) SetRunningVersion(v int32)`

SetRunningVersion sets RunningVersion field to given value.

### HasRunningVersion

`func (o *PlanStatusVersion) HasRunningVersion() bool`

HasRunningVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


