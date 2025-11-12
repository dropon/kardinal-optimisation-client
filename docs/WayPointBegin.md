# WayPointBegin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Position** | [**Position**](Position.md) |  | 
**Status** | Pointer to [**WayPointStatus**](WayPointStatus.md) |  | [optional] 
**DepartureTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**Violations** | Pointer to [**[]WaypointViolation**](WaypointViolation.md) | List of the Waypoint&#39;s violations. | [optional] 
**FromPrevious** | Pointer to [**FromPrevious**](FromPrevious.md) |  | [optional] 

## Methods

### NewWayPointBegin

`func NewWayPointBegin(type_ string, position Position, ) *WayPointBegin`

NewWayPointBegin instantiates a new WayPointBegin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWayPointBeginWithDefaults

`func NewWayPointBeginWithDefaults() *WayPointBegin`

NewWayPointBeginWithDefaults instantiates a new WayPointBegin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WayPointBegin) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WayPointBegin) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WayPointBegin) SetType(v string)`

SetType sets Type field to given value.


### GetPosition

`func (o *WayPointBegin) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WayPointBegin) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WayPointBegin) SetPosition(v Position)`

SetPosition sets Position field to given value.


### GetStatus

`func (o *WayPointBegin) GetStatus() WayPointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WayPointBegin) GetStatusOk() (*WayPointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WayPointBegin) SetStatus(v WayPointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WayPointBegin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDepartureTime

`func (o *WayPointBegin) GetDepartureTime() string`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *WayPointBegin) GetDepartureTimeOk() (*string, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *WayPointBegin) SetDepartureTime(v string)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *WayPointBegin) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### GetViolations

`func (o *WayPointBegin) GetViolations() []WaypointViolation`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *WayPointBegin) GetViolationsOk() (*[]WaypointViolation, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *WayPointBegin) SetViolations(v []WaypointViolation)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *WayPointBegin) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### GetFromPrevious

`func (o *WayPointBegin) GetFromPrevious() FromPrevious`

GetFromPrevious returns the FromPrevious field if non-nil, zero value otherwise.

### GetFromPreviousOk

`func (o *WayPointBegin) GetFromPreviousOk() (*FromPrevious, bool)`

GetFromPreviousOk returns a tuple with the FromPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPrevious

`func (o *WayPointBegin) SetFromPrevious(v FromPrevious)`

SetFromPrevious sets FromPrevious field to given value.

### HasFromPrevious

`func (o *WayPointBegin) HasFromPrevious() bool`

HasFromPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


