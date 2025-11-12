# WayPointBreak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | Pointer to [**WayPointStatus**](WayPointStatus.md) |  | [optional] 
**ArrivalTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**DepartureTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**Violations** | Pointer to [**[]WaypointViolation**](WaypointViolation.md) | List of the Waypoint&#39;s violations. | [optional] 
**FromPrevious** | Pointer to [**FromPrevious**](FromPrevious.md) |  | [optional] 

## Methods

### NewWayPointBreak

`func NewWayPointBreak(type_ string, ) *WayPointBreak`

NewWayPointBreak instantiates a new WayPointBreak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWayPointBreakWithDefaults

`func NewWayPointBreakWithDefaults() *WayPointBreak`

NewWayPointBreakWithDefaults instantiates a new WayPointBreak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WayPointBreak) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WayPointBreak) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WayPointBreak) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *WayPointBreak) GetStatus() WayPointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WayPointBreak) GetStatusOk() (*WayPointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WayPointBreak) SetStatus(v WayPointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WayPointBreak) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetArrivalTime

`func (o *WayPointBreak) GetArrivalTime() string`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *WayPointBreak) GetArrivalTimeOk() (*string, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *WayPointBreak) SetArrivalTime(v string)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *WayPointBreak) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.

### GetDepartureTime

`func (o *WayPointBreak) GetDepartureTime() string`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *WayPointBreak) GetDepartureTimeOk() (*string, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *WayPointBreak) SetDepartureTime(v string)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *WayPointBreak) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### GetViolations

`func (o *WayPointBreak) GetViolations() []WaypointViolation`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *WayPointBreak) GetViolationsOk() (*[]WaypointViolation, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *WayPointBreak) SetViolations(v []WaypointViolation)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *WayPointBreak) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### GetFromPrevious

`func (o *WayPointBreak) GetFromPrevious() FromPrevious`

GetFromPrevious returns the FromPrevious field if non-nil, zero value otherwise.

### GetFromPreviousOk

`func (o *WayPointBreak) GetFromPreviousOk() (*FromPrevious, bool)`

GetFromPreviousOk returns a tuple with the FromPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPrevious

`func (o *WayPointBreak) SetFromPrevious(v FromPrevious)`

SetFromPrevious sets FromPrevious field to given value.

### HasFromPrevious

`func (o *WayPointBreak) HasFromPrevious() bool`

HasFromPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


