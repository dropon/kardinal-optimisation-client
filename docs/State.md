# State

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**ResourceMode**](ResourceMode.md) |  | [optional] [default to RESOURCEMODE_FREE]
**Assignments** | Pointer to [**[]StateAssignmentsInner**](StateAssignmentsInner.md) |  | [optional] 

## Methods

### NewState

`func NewState() *State`

NewState instantiates a new State object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateWithDefaults

`func NewStateWithDefaults() *State`

NewStateWithDefaults instantiates a new State object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *State) GetMode() ResourceMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *State) GetModeOk() (*ResourceMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *State) SetMode(v ResourceMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *State) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAssignments

`func (o *State) GetAssignments() []StateAssignmentsInner`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *State) GetAssignmentsOk() (*[]StateAssignmentsInner, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *State) SetAssignments(v []StateAssignmentsInner)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *State) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


