# ScoredTimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeWindow** | [**TimeWindow**](TimeWindow.md) |  | 
**PlanId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanVersion** | Pointer to **int32** | The plan version. | [optional] [readonly] 
**Objectives** | [**[]SolutionObjective**](SolutionObjective.md) |  | 
**InitialObjectives** | Pointer to [**[]SolutionObjective**](SolutionObjective.md) |  | [optional] 
**SolutionTimestamp** | Pointer to **NullableString** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**ResourceId** | Pointer to **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | [optional] 
**BeginTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 

## Methods

### NewScoredTimeWindow

`func NewScoredTimeWindow(timeWindow TimeWindow, objectives []SolutionObjective, ) *ScoredTimeWindow`

NewScoredTimeWindow instantiates a new ScoredTimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoredTimeWindowWithDefaults

`func NewScoredTimeWindowWithDefaults() *ScoredTimeWindow`

NewScoredTimeWindowWithDefaults instantiates a new ScoredTimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeWindow

`func (o *ScoredTimeWindow) GetTimeWindow() TimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *ScoredTimeWindow) GetTimeWindowOk() (*TimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *ScoredTimeWindow) SetTimeWindow(v TimeWindow)`

SetTimeWindow sets TimeWindow field to given value.


### GetPlanId

`func (o *ScoredTimeWindow) GetPlanId() interface{}`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ScoredTimeWindow) GetPlanIdOk() (*interface{}, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ScoredTimeWindow) SetPlanId(v interface{})`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ScoredTimeWindow) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanVersion

`func (o *ScoredTimeWindow) GetPlanVersion() int32`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *ScoredTimeWindow) GetPlanVersionOk() (*int32, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *ScoredTimeWindow) SetPlanVersion(v int32)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *ScoredTimeWindow) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.

### GetObjectives

`func (o *ScoredTimeWindow) GetObjectives() []SolutionObjective`

GetObjectives returns the Objectives field if non-nil, zero value otherwise.

### GetObjectivesOk

`func (o *ScoredTimeWindow) GetObjectivesOk() (*[]SolutionObjective, bool)`

GetObjectivesOk returns a tuple with the Objectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectives

`func (o *ScoredTimeWindow) SetObjectives(v []SolutionObjective)`

SetObjectives sets Objectives field to given value.


### GetInitialObjectives

`func (o *ScoredTimeWindow) GetInitialObjectives() []SolutionObjective`

GetInitialObjectives returns the InitialObjectives field if non-nil, zero value otherwise.

### GetInitialObjectivesOk

`func (o *ScoredTimeWindow) GetInitialObjectivesOk() (*[]SolutionObjective, bool)`

GetInitialObjectivesOk returns a tuple with the InitialObjectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialObjectives

`func (o *ScoredTimeWindow) SetInitialObjectives(v []SolutionObjective)`

SetInitialObjectives sets InitialObjectives field to given value.

### HasInitialObjectives

`func (o *ScoredTimeWindow) HasInitialObjectives() bool`

HasInitialObjectives returns a boolean if a field has been set.

### GetSolutionTimestamp

`func (o *ScoredTimeWindow) GetSolutionTimestamp() string`

GetSolutionTimestamp returns the SolutionTimestamp field if non-nil, zero value otherwise.

### GetSolutionTimestampOk

`func (o *ScoredTimeWindow) GetSolutionTimestampOk() (*string, bool)`

GetSolutionTimestampOk returns a tuple with the SolutionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionTimestamp

`func (o *ScoredTimeWindow) SetSolutionTimestamp(v string)`

SetSolutionTimestamp sets SolutionTimestamp field to given value.

### HasSolutionTimestamp

`func (o *ScoredTimeWindow) HasSolutionTimestamp() bool`

HasSolutionTimestamp returns a boolean if a field has been set.

### SetSolutionTimestampNil

`func (o *ScoredTimeWindow) SetSolutionTimestampNil(b bool)`

 SetSolutionTimestampNil sets the value for SolutionTimestamp to be an explicit nil

### UnsetSolutionTimestamp
`func (o *ScoredTimeWindow) UnsetSolutionTimestamp()`

UnsetSolutionTimestamp ensures that no value is present for SolutionTimestamp, not even an explicit nil
### GetResourceId

`func (o *ScoredTimeWindow) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ScoredTimeWindow) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ScoredTimeWindow) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ScoredTimeWindow) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetBeginTime

`func (o *ScoredTimeWindow) GetBeginTime() string`

GetBeginTime returns the BeginTime field if non-nil, zero value otherwise.

### GetBeginTimeOk

`func (o *ScoredTimeWindow) GetBeginTimeOk() (*string, bool)`

GetBeginTimeOk returns a tuple with the BeginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginTime

`func (o *ScoredTimeWindow) SetBeginTime(v string)`

SetBeginTime sets BeginTime field to given value.

### HasBeginTime

`func (o *ScoredTimeWindow) HasBeginTime() bool`

HasBeginTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


