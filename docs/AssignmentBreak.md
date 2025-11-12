# AssignmentBreak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | Pointer to [**AssignmentStatus**](AssignmentStatus.md) |  | [optional] [default to ASSIGNMENTSTATUS_FIXED]
**ArrivalTime** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 
**DepartureTime** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 

## Methods

### NewAssignmentBreak

`func NewAssignmentBreak(type_ string, arrivalTime string, departureTime string, ) *AssignmentBreak`

NewAssignmentBreak instantiates a new AssignmentBreak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentBreakWithDefaults

`func NewAssignmentBreakWithDefaults() *AssignmentBreak`

NewAssignmentBreakWithDefaults instantiates a new AssignmentBreak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AssignmentBreak) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignmentBreak) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignmentBreak) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *AssignmentBreak) GetStatus() AssignmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssignmentBreak) GetStatusOk() (*AssignmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssignmentBreak) SetStatus(v AssignmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssignmentBreak) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetArrivalTime

`func (o *AssignmentBreak) GetArrivalTime() string`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *AssignmentBreak) GetArrivalTimeOk() (*string, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *AssignmentBreak) SetArrivalTime(v string)`

SetArrivalTime sets ArrivalTime field to given value.


### GetDepartureTime

`func (o *AssignmentBreak) GetDepartureTime() string`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *AssignmentBreak) GetDepartureTimeOk() (*string, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *AssignmentBreak) SetDepartureTime(v string)`

SetDepartureTime sets DepartureTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


