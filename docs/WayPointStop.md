# WayPointStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StopId** | Pointer to **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | [optional] 
**Position** | [**Position**](Position.md) |  | 
**Status** | Pointer to [**WayPointStatus**](WayPointStatus.md) |  | [optional] 
**StopProperties** | Pointer to **map[string]string** |  | [optional] 
**OrderProperties** | Pointer to **map[string]string** |  | [optional] 
**FilledCapacitiesAfterStop** | Pointer to **map[string]float32** |  | [optional] 
**Delay** | Pointer to **string** | Delay related to the stop&#39;s preferred time windows. | [optional] 
**ArrivalTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**BeginTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**DepartureTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**Violations** | Pointer to [**[]WaypointViolation**](WaypointViolation.md) | List of the Waypoint&#39;s violations. | [optional] 
**FromPrevious** | Pointer to [**FromPrevious**](FromPrevious.md) |  | [optional] 
**StopKind** | Pointer to [**StopKind**](StopKind.md) |  | [optional] [default to STOPKIND_DELIVERY]

## Methods

### NewWayPointStop

`func NewWayPointStop(type_ string, position Position, ) *WayPointStop`

NewWayPointStop instantiates a new WayPointStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWayPointStopWithDefaults

`func NewWayPointStopWithDefaults() *WayPointStop`

NewWayPointStopWithDefaults instantiates a new WayPointStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WayPointStop) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WayPointStop) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WayPointStop) SetType(v string)`

SetType sets Type field to given value.


### GetStopId

`func (o *WayPointStop) GetStopId() string`

GetStopId returns the StopId field if non-nil, zero value otherwise.

### GetStopIdOk

`func (o *WayPointStop) GetStopIdOk() (*string, bool)`

GetStopIdOk returns a tuple with the StopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopId

`func (o *WayPointStop) SetStopId(v string)`

SetStopId sets StopId field to given value.

### HasStopId

`func (o *WayPointStop) HasStopId() bool`

HasStopId returns a boolean if a field has been set.

### GetPosition

`func (o *WayPointStop) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WayPointStop) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WayPointStop) SetPosition(v Position)`

SetPosition sets Position field to given value.


### GetStatus

`func (o *WayPointStop) GetStatus() WayPointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WayPointStop) GetStatusOk() (*WayPointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WayPointStop) SetStatus(v WayPointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WayPointStop) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopProperties

`func (o *WayPointStop) GetStopProperties() map[string]string`

GetStopProperties returns the StopProperties field if non-nil, zero value otherwise.

### GetStopPropertiesOk

`func (o *WayPointStop) GetStopPropertiesOk() (*map[string]string, bool)`

GetStopPropertiesOk returns a tuple with the StopProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProperties

`func (o *WayPointStop) SetStopProperties(v map[string]string)`

SetStopProperties sets StopProperties field to given value.

### HasStopProperties

`func (o *WayPointStop) HasStopProperties() bool`

HasStopProperties returns a boolean if a field has been set.

### GetOrderProperties

`func (o *WayPointStop) GetOrderProperties() map[string]string`

GetOrderProperties returns the OrderProperties field if non-nil, zero value otherwise.

### GetOrderPropertiesOk

`func (o *WayPointStop) GetOrderPropertiesOk() (*map[string]string, bool)`

GetOrderPropertiesOk returns a tuple with the OrderProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProperties

`func (o *WayPointStop) SetOrderProperties(v map[string]string)`

SetOrderProperties sets OrderProperties field to given value.

### HasOrderProperties

`func (o *WayPointStop) HasOrderProperties() bool`

HasOrderProperties returns a boolean if a field has been set.

### GetFilledCapacitiesAfterStop

`func (o *WayPointStop) GetFilledCapacitiesAfterStop() map[string]float32`

GetFilledCapacitiesAfterStop returns the FilledCapacitiesAfterStop field if non-nil, zero value otherwise.

### GetFilledCapacitiesAfterStopOk

`func (o *WayPointStop) GetFilledCapacitiesAfterStopOk() (*map[string]float32, bool)`

GetFilledCapacitiesAfterStopOk returns a tuple with the FilledCapacitiesAfterStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledCapacitiesAfterStop

`func (o *WayPointStop) SetFilledCapacitiesAfterStop(v map[string]float32)`

SetFilledCapacitiesAfterStop sets FilledCapacitiesAfterStop field to given value.

### HasFilledCapacitiesAfterStop

`func (o *WayPointStop) HasFilledCapacitiesAfterStop() bool`

HasFilledCapacitiesAfterStop returns a boolean if a field has been set.

### GetDelay

`func (o *WayPointStop) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *WayPointStop) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *WayPointStop) SetDelay(v string)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *WayPointStop) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetArrivalTime

`func (o *WayPointStop) GetArrivalTime() string`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *WayPointStop) GetArrivalTimeOk() (*string, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *WayPointStop) SetArrivalTime(v string)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *WayPointStop) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.

### GetBeginTime

`func (o *WayPointStop) GetBeginTime() string`

GetBeginTime returns the BeginTime field if non-nil, zero value otherwise.

### GetBeginTimeOk

`func (o *WayPointStop) GetBeginTimeOk() (*string, bool)`

GetBeginTimeOk returns a tuple with the BeginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginTime

`func (o *WayPointStop) SetBeginTime(v string)`

SetBeginTime sets BeginTime field to given value.

### HasBeginTime

`func (o *WayPointStop) HasBeginTime() bool`

HasBeginTime returns a boolean if a field has been set.

### GetDepartureTime

`func (o *WayPointStop) GetDepartureTime() string`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *WayPointStop) GetDepartureTimeOk() (*string, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *WayPointStop) SetDepartureTime(v string)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *WayPointStop) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### GetViolations

`func (o *WayPointStop) GetViolations() []WaypointViolation`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *WayPointStop) GetViolationsOk() (*[]WaypointViolation, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *WayPointStop) SetViolations(v []WaypointViolation)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *WayPointStop) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### GetFromPrevious

`func (o *WayPointStop) GetFromPrevious() FromPrevious`

GetFromPrevious returns the FromPrevious field if non-nil, zero value otherwise.

### GetFromPreviousOk

`func (o *WayPointStop) GetFromPreviousOk() (*FromPrevious, bool)`

GetFromPreviousOk returns a tuple with the FromPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPrevious

`func (o *WayPointStop) SetFromPrevious(v FromPrevious)`

SetFromPrevious sets FromPrevious field to given value.

### HasFromPrevious

`func (o *WayPointStop) HasFromPrevious() bool`

HasFromPrevious returns a boolean if a field has been set.

### GetStopKind

`func (o *WayPointStop) GetStopKind() StopKind`

GetStopKind returns the StopKind field if non-nil, zero value otherwise.

### GetStopKindOk

`func (o *WayPointStop) GetStopKindOk() (*StopKind, bool)`

GetStopKindOk returns a tuple with the StopKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopKind

`func (o *WayPointStop) SetStopKind(v StopKind)`

SetStopKind sets StopKind field to given value.

### HasStopKind

`func (o *WayPointStop) HasStopKind() bool`

HasStopKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


