# AssignmentEnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | Pointer to [**AssignmentStatus**](AssignmentStatus.md) |  | [optional] [default to ASSIGNMENTSTATUS_FIXED]
**ArrivalTime** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 

## Methods

### NewAssignmentEnd

`func NewAssignmentEnd(type_ string, arrivalTime string, ) *AssignmentEnd`

NewAssignmentEnd instantiates a new AssignmentEnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentEndWithDefaults

`func NewAssignmentEndWithDefaults() *AssignmentEnd`

NewAssignmentEndWithDefaults instantiates a new AssignmentEnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AssignmentEnd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignmentEnd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignmentEnd) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *AssignmentEnd) GetStatus() AssignmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssignmentEnd) GetStatusOk() (*AssignmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssignmentEnd) SetStatus(v AssignmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssignmentEnd) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetArrivalTime

`func (o *AssignmentEnd) GetArrivalTime() string`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *AssignmentEnd) GetArrivalTimeOk() (*string, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *AssignmentEnd) SetArrivalTime(v string)`

SetArrivalTime sets ArrivalTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


