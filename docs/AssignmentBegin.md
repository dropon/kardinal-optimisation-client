# AssignmentBegin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | Pointer to [**AssignmentStatus**](AssignmentStatus.md) |  | [optional] [default to ASSIGNMENTSTATUS_FIXED]
**DepartureTime** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 

## Methods

### NewAssignmentBegin

`func NewAssignmentBegin(type_ string, departureTime string, ) *AssignmentBegin`

NewAssignmentBegin instantiates a new AssignmentBegin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentBeginWithDefaults

`func NewAssignmentBeginWithDefaults() *AssignmentBegin`

NewAssignmentBeginWithDefaults instantiates a new AssignmentBegin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AssignmentBegin) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignmentBegin) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignmentBegin) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *AssignmentBegin) GetStatus() AssignmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssignmentBegin) GetStatusOk() (*AssignmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssignmentBegin) SetStatus(v AssignmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssignmentBegin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDepartureTime

`func (o *AssignmentBegin) GetDepartureTime() string`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *AssignmentBegin) GetDepartureTimeOk() (*string, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *AssignmentBegin) SetDepartureTime(v string)`

SetDepartureTime sets DepartureTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


