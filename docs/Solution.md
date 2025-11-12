# Solution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgencyId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**PlanVersion** | Pointer to **int32** | The plan version. | [optional] [readonly] 
**Timestamp** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**UnaffectedStopIds** | Pointer to **[]string** |  | [optional] 
**UnaffectedAlternativeIds** | Pointer to **[]string** |  | [optional] 
**UnusedResourceIds** | Pointer to **[]string** |  | [optional] 
**Objectives** | Pointer to [**[]SolutionObjective**](SolutionObjective.md) |  | [optional] 
**Tours** | Pointer to [**[]Tour**](Tour.md) |  | [optional] 
**TotalDelay** | Pointer to **string** | Total delay of the solution related to the preferred time windows. | [optional] 
**GlobalViolations** | Pointer to [**[]SolutionGlobalViolationsInner**](SolutionGlobalViolationsInner.md) |  | [optional] 

## Methods

### NewSolution

`func NewSolution() *Solution`

NewSolution instantiates a new Solution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolutionWithDefaults

`func NewSolutionWithDefaults() *Solution`

NewSolutionWithDefaults instantiates a new Solution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgencyId

`func (o *Solution) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *Solution) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *Solution) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.

### HasAgencyId

`func (o *Solution) HasAgencyId() bool`

HasAgencyId returns a boolean if a field has been set.

### GetPlanId

`func (o *Solution) GetPlanId() interface{}`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *Solution) GetPlanIdOk() (*interface{}, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *Solution) SetPlanId(v interface{})`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *Solution) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanVersion

`func (o *Solution) GetPlanVersion() int32`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *Solution) GetPlanVersionOk() (*int32, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *Solution) SetPlanVersion(v int32)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *Solution) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.

### GetTimestamp

`func (o *Solution) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Solution) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Solution) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Solution) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUnaffectedStopIds

`func (o *Solution) GetUnaffectedStopIds() []string`

GetUnaffectedStopIds returns the UnaffectedStopIds field if non-nil, zero value otherwise.

### GetUnaffectedStopIdsOk

`func (o *Solution) GetUnaffectedStopIdsOk() (*[]string, bool)`

GetUnaffectedStopIdsOk returns a tuple with the UnaffectedStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnaffectedStopIds

`func (o *Solution) SetUnaffectedStopIds(v []string)`

SetUnaffectedStopIds sets UnaffectedStopIds field to given value.

### HasUnaffectedStopIds

`func (o *Solution) HasUnaffectedStopIds() bool`

HasUnaffectedStopIds returns a boolean if a field has been set.

### GetUnaffectedAlternativeIds

`func (o *Solution) GetUnaffectedAlternativeIds() []string`

GetUnaffectedAlternativeIds returns the UnaffectedAlternativeIds field if non-nil, zero value otherwise.

### GetUnaffectedAlternativeIdsOk

`func (o *Solution) GetUnaffectedAlternativeIdsOk() (*[]string, bool)`

GetUnaffectedAlternativeIdsOk returns a tuple with the UnaffectedAlternativeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnaffectedAlternativeIds

`func (o *Solution) SetUnaffectedAlternativeIds(v []string)`

SetUnaffectedAlternativeIds sets UnaffectedAlternativeIds field to given value.

### HasUnaffectedAlternativeIds

`func (o *Solution) HasUnaffectedAlternativeIds() bool`

HasUnaffectedAlternativeIds returns a boolean if a field has been set.

### GetUnusedResourceIds

`func (o *Solution) GetUnusedResourceIds() []string`

GetUnusedResourceIds returns the UnusedResourceIds field if non-nil, zero value otherwise.

### GetUnusedResourceIdsOk

`func (o *Solution) GetUnusedResourceIdsOk() (*[]string, bool)`

GetUnusedResourceIdsOk returns a tuple with the UnusedResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedResourceIds

`func (o *Solution) SetUnusedResourceIds(v []string)`

SetUnusedResourceIds sets UnusedResourceIds field to given value.

### HasUnusedResourceIds

`func (o *Solution) HasUnusedResourceIds() bool`

HasUnusedResourceIds returns a boolean if a field has been set.

### GetObjectives

`func (o *Solution) GetObjectives() []SolutionObjective`

GetObjectives returns the Objectives field if non-nil, zero value otherwise.

### GetObjectivesOk

`func (o *Solution) GetObjectivesOk() (*[]SolutionObjective, bool)`

GetObjectivesOk returns a tuple with the Objectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectives

`func (o *Solution) SetObjectives(v []SolutionObjective)`

SetObjectives sets Objectives field to given value.

### HasObjectives

`func (o *Solution) HasObjectives() bool`

HasObjectives returns a boolean if a field has been set.

### GetTours

`func (o *Solution) GetTours() []Tour`

GetTours returns the Tours field if non-nil, zero value otherwise.

### GetToursOk

`func (o *Solution) GetToursOk() (*[]Tour, bool)`

GetToursOk returns a tuple with the Tours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTours

`func (o *Solution) SetTours(v []Tour)`

SetTours sets Tours field to given value.

### HasTours

`func (o *Solution) HasTours() bool`

HasTours returns a boolean if a field has been set.

### GetTotalDelay

`func (o *Solution) GetTotalDelay() string`

GetTotalDelay returns the TotalDelay field if non-nil, zero value otherwise.

### GetTotalDelayOk

`func (o *Solution) GetTotalDelayOk() (*string, bool)`

GetTotalDelayOk returns a tuple with the TotalDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDelay

`func (o *Solution) SetTotalDelay(v string)`

SetTotalDelay sets TotalDelay field to given value.

### HasTotalDelay

`func (o *Solution) HasTotalDelay() bool`

HasTotalDelay returns a boolean if a field has been set.

### GetGlobalViolations

`func (o *Solution) GetGlobalViolations() []SolutionGlobalViolationsInner`

GetGlobalViolations returns the GlobalViolations field if non-nil, zero value otherwise.

### GetGlobalViolationsOk

`func (o *Solution) GetGlobalViolationsOk() (*[]SolutionGlobalViolationsInner, bool)`

GetGlobalViolationsOk returns a tuple with the GlobalViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalViolations

`func (o *Solution) SetGlobalViolations(v []SolutionGlobalViolationsInner)`

SetGlobalViolations sets GlobalViolations field to given value.

### HasGlobalViolations

`func (o *Solution) HasGlobalViolations() bool`

HasGlobalViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


