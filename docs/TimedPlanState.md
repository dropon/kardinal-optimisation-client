# TimedPlanState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanVersion** | Pointer to **float32** | The corresponding plan&#39;s version. | [optional] 
**Timestamp** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**State** | Pointer to [**PlanState**](PlanState.md) |  | [optional] 

## Methods

### NewTimedPlanState

`func NewTimedPlanState() *TimedPlanState`

NewTimedPlanState instantiates a new TimedPlanState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimedPlanStateWithDefaults

`func NewTimedPlanStateWithDefaults() *TimedPlanState`

NewTimedPlanStateWithDefaults instantiates a new TimedPlanState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanVersion

`func (o *TimedPlanState) GetPlanVersion() float32`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *TimedPlanState) GetPlanVersionOk() (*float32, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *TimedPlanState) SetPlanVersion(v float32)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *TimedPlanState) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.

### GetTimestamp

`func (o *TimedPlanState) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimedPlanState) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimedPlanState) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TimedPlanState) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetState

`func (o *TimedPlanState) GetState() PlanState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TimedPlanState) GetStateOk() (*PlanState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TimedPlanState) SetState(v PlanState)`

SetState sets State field to given value.

### HasState

`func (o *TimedPlanState) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


