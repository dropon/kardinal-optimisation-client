# WayPointEnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Position** | [**Position**](Position.md) |  | 
**Status** | Pointer to [**WayPointStatus**](WayPointStatus.md) |  | [optional] 
**ArrivalTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**Violations** | Pointer to [**[]WaypointViolation**](WaypointViolation.md) | List of the Waypoint&#39;s violations. | [optional] 
**FromPrevious** | Pointer to [**FromPrevious**](FromPrevious.md) |  | [optional] 

## Methods

### NewWayPointEnd

`func NewWayPointEnd(type_ string, position Position, ) *WayPointEnd`

NewWayPointEnd instantiates a new WayPointEnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWayPointEndWithDefaults

`func NewWayPointEndWithDefaults() *WayPointEnd`

NewWayPointEndWithDefaults instantiates a new WayPointEnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WayPointEnd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WayPointEnd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WayPointEnd) SetType(v string)`

SetType sets Type field to given value.


### GetPosition

`func (o *WayPointEnd) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WayPointEnd) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WayPointEnd) SetPosition(v Position)`

SetPosition sets Position field to given value.


### GetStatus

`func (o *WayPointEnd) GetStatus() WayPointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WayPointEnd) GetStatusOk() (*WayPointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WayPointEnd) SetStatus(v WayPointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WayPointEnd) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetArrivalTime

`func (o *WayPointEnd) GetArrivalTime() string`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *WayPointEnd) GetArrivalTimeOk() (*string, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *WayPointEnd) SetArrivalTime(v string)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *WayPointEnd) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.

### GetViolations

`func (o *WayPointEnd) GetViolations() []WaypointViolation`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *WayPointEnd) GetViolationsOk() (*[]WaypointViolation, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *WayPointEnd) SetViolations(v []WaypointViolation)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *WayPointEnd) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### GetFromPrevious

`func (o *WayPointEnd) GetFromPrevious() FromPrevious`

GetFromPrevious returns the FromPrevious field if non-nil, zero value otherwise.

### GetFromPreviousOk

`func (o *WayPointEnd) GetFromPreviousOk() (*FromPrevious, bool)`

GetFromPreviousOk returns a tuple with the FromPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPrevious

`func (o *WayPointEnd) SetFromPrevious(v FromPrevious)`

SetFromPrevious sets FromPrevious field to given value.

### HasFromPrevious

`func (o *WayPointEnd) HasFromPrevious() bool`

HasFromPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


