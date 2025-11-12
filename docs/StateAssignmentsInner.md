# StateAssignmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | Pointer to [**AssignmentStatus**](AssignmentStatus.md) |  | [optional] [default to ASSIGNMENTSTATUS_FIXED]
**StopId** | **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | 
**ArrivalTime** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 
**BeginTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**DepartureTime** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 

## Methods

### NewStateAssignmentsInner

`func NewStateAssignmentsInner(type_ string, stopId string, arrivalTime string, departureTime string, ) *StateAssignmentsInner`

NewStateAssignmentsInner instantiates a new StateAssignmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateAssignmentsInnerWithDefaults

`func NewStateAssignmentsInnerWithDefaults() *StateAssignmentsInner`

NewStateAssignmentsInnerWithDefaults instantiates a new StateAssignmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StateAssignmentsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StateAssignmentsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StateAssignmentsInner) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *StateAssignmentsInner) GetStatus() AssignmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StateAssignmentsInner) GetStatusOk() (*AssignmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StateAssignmentsInner) SetStatus(v AssignmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StateAssignmentsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopId

`func (o *StateAssignmentsInner) GetStopId() string`

GetStopId returns the StopId field if non-nil, zero value otherwise.

### GetStopIdOk

`func (o *StateAssignmentsInner) GetStopIdOk() (*string, bool)`

GetStopIdOk returns a tuple with the StopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopId

`func (o *StateAssignmentsInner) SetStopId(v string)`

SetStopId sets StopId field to given value.


### GetArrivalTime

`func (o *StateAssignmentsInner) GetArrivalTime() string`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *StateAssignmentsInner) GetArrivalTimeOk() (*string, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *StateAssignmentsInner) SetArrivalTime(v string)`

SetArrivalTime sets ArrivalTime field to given value.


### GetBeginTime

`func (o *StateAssignmentsInner) GetBeginTime() string`

GetBeginTime returns the BeginTime field if non-nil, zero value otherwise.

### GetBeginTimeOk

`func (o *StateAssignmentsInner) GetBeginTimeOk() (*string, bool)`

GetBeginTimeOk returns a tuple with the BeginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginTime

`func (o *StateAssignmentsInner) SetBeginTime(v string)`

SetBeginTime sets BeginTime field to given value.

### HasBeginTime

`func (o *StateAssignmentsInner) HasBeginTime() bool`

HasBeginTime returns a boolean if a field has been set.

### GetDepartureTime

`func (o *StateAssignmentsInner) GetDepartureTime() string`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *StateAssignmentsInner) GetDepartureTimeOk() (*string, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *StateAssignmentsInner) SetDepartureTime(v string)`

SetDepartureTime sets DepartureTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


