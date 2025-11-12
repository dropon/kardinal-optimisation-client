# PlanStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanVersionInSolution** | Pointer to **int32** | The plan version taken into account in the current solution. | [optional] [readonly] 
**WaitingRoom** | Pointer to [**PlanStatusVersion**](PlanStatusVersion.md) | If the maximum number of simultaneous running plans has already been reached, the plan waits in the waiting room for one of the running plans to finish. | [optional] 
**WaitingTraffic** | Pointer to [**PlanStatusVersion**](PlanStatusVersion.md) | The plan is waiting for its traffic coefficients to be computed. | [optional] 
**Creation** | Pointer to [**PlanStatusVersion**](PlanStatusVersion.md) | The plan is being created in order to be optimized. | [optional] 
**Optimization** | Pointer to [**PlanStatusVersion**](PlanStatusVersion.md) | The plan is being optimized. | [optional] 

## Methods

### NewPlanStatus

`func NewPlanStatus() *PlanStatus`

NewPlanStatus instantiates a new PlanStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanStatusWithDefaults

`func NewPlanStatusWithDefaults() *PlanStatus`

NewPlanStatusWithDefaults instantiates a new PlanStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanVersionInSolution

`func (o *PlanStatus) GetPlanVersionInSolution() int32`

GetPlanVersionInSolution returns the PlanVersionInSolution field if non-nil, zero value otherwise.

### GetPlanVersionInSolutionOk

`func (o *PlanStatus) GetPlanVersionInSolutionOk() (*int32, bool)`

GetPlanVersionInSolutionOk returns a tuple with the PlanVersionInSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersionInSolution

`func (o *PlanStatus) SetPlanVersionInSolution(v int32)`

SetPlanVersionInSolution sets PlanVersionInSolution field to given value.

### HasPlanVersionInSolution

`func (o *PlanStatus) HasPlanVersionInSolution() bool`

HasPlanVersionInSolution returns a boolean if a field has been set.

### GetWaitingRoom

`func (o *PlanStatus) GetWaitingRoom() PlanStatusVersion`

GetWaitingRoom returns the WaitingRoom field if non-nil, zero value otherwise.

### GetWaitingRoomOk

`func (o *PlanStatus) GetWaitingRoomOk() (*PlanStatusVersion, bool)`

GetWaitingRoomOk returns a tuple with the WaitingRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingRoom

`func (o *PlanStatus) SetWaitingRoom(v PlanStatusVersion)`

SetWaitingRoom sets WaitingRoom field to given value.

### HasWaitingRoom

`func (o *PlanStatus) HasWaitingRoom() bool`

HasWaitingRoom returns a boolean if a field has been set.

### GetWaitingTraffic

`func (o *PlanStatus) GetWaitingTraffic() PlanStatusVersion`

GetWaitingTraffic returns the WaitingTraffic field if non-nil, zero value otherwise.

### GetWaitingTrafficOk

`func (o *PlanStatus) GetWaitingTrafficOk() (*PlanStatusVersion, bool)`

GetWaitingTrafficOk returns a tuple with the WaitingTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTraffic

`func (o *PlanStatus) SetWaitingTraffic(v PlanStatusVersion)`

SetWaitingTraffic sets WaitingTraffic field to given value.

### HasWaitingTraffic

`func (o *PlanStatus) HasWaitingTraffic() bool`

HasWaitingTraffic returns a boolean if a field has been set.

### GetCreation

`func (o *PlanStatus) GetCreation() PlanStatusVersion`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *PlanStatus) GetCreationOk() (*PlanStatusVersion, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *PlanStatus) SetCreation(v PlanStatusVersion)`

SetCreation sets Creation field to given value.

### HasCreation

`func (o *PlanStatus) HasCreation() bool`

HasCreation returns a boolean if a field has been set.

### GetOptimization

`func (o *PlanStatus) GetOptimization() PlanStatusVersion`

GetOptimization returns the Optimization field if non-nil, zero value otherwise.

### GetOptimizationOk

`func (o *PlanStatus) GetOptimizationOk() (*PlanStatusVersion, bool)`

GetOptimizationOk returns a tuple with the Optimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimization

`func (o *PlanStatus) SetOptimization(v PlanStatusVersion)`

SetOptimization sets Optimization field to given value.

### HasOptimization

`func (o *PlanStatus) HasOptimization() bool`

HasOptimization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


