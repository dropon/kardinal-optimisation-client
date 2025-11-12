# TourWayPointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Position** | [**Position**](Position.md) |  | 
**Status** | Pointer to [**WayPointStatus**](WayPointStatus.md) |  | [optional] 
**DepartureTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**Violations** | Pointer to [**[]WaypointViolation**](WaypointViolation.md) | List of the Waypoint&#39;s violations. | [optional] 
**FromPrevious** | Pointer to [**FromPrevious**](FromPrevious.md) |  | [optional] 
**StopId** | Pointer to **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | [optional] 
**StopProperties** | Pointer to **map[string]string** |  | [optional] 
**OrderProperties** | Pointer to **map[string]string** |  | [optional] 
**FilledCapacitiesAfterStop** | Pointer to **map[string]float32** |  | [optional] 
**Delay** | Pointer to **string** | Delay related to the stop&#39;s preferred time windows. | [optional] 
**ArrivalTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**BeginTime** | Pointer to **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | [optional] 
**StopKind** | Pointer to [**StopKind**](StopKind.md) |  | [optional] [default to STOPKIND_DELIVERY]

## Methods

### NewTourWayPointsInner

`func NewTourWayPointsInner(type_ string, position Position, ) *TourWayPointsInner`

NewTourWayPointsInner instantiates a new TourWayPointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTourWayPointsInnerWithDefaults

`func NewTourWayPointsInnerWithDefaults() *TourWayPointsInner`

NewTourWayPointsInnerWithDefaults instantiates a new TourWayPointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TourWayPointsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TourWayPointsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TourWayPointsInner) SetType(v string)`

SetType sets Type field to given value.


### GetPosition

`func (o *TourWayPointsInner) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TourWayPointsInner) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TourWayPointsInner) SetPosition(v Position)`

SetPosition sets Position field to given value.


### GetStatus

`func (o *TourWayPointsInner) GetStatus() WayPointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TourWayPointsInner) GetStatusOk() (*WayPointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TourWayPointsInner) SetStatus(v WayPointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TourWayPointsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDepartureTime

`func (o *TourWayPointsInner) GetDepartureTime() string`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *TourWayPointsInner) GetDepartureTimeOk() (*string, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *TourWayPointsInner) SetDepartureTime(v string)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *TourWayPointsInner) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### GetViolations

`func (o *TourWayPointsInner) GetViolations() []WaypointViolation`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *TourWayPointsInner) GetViolationsOk() (*[]WaypointViolation, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *TourWayPointsInner) SetViolations(v []WaypointViolation)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *TourWayPointsInner) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### GetFromPrevious

`func (o *TourWayPointsInner) GetFromPrevious() FromPrevious`

GetFromPrevious returns the FromPrevious field if non-nil, zero value otherwise.

### GetFromPreviousOk

`func (o *TourWayPointsInner) GetFromPreviousOk() (*FromPrevious, bool)`

GetFromPreviousOk returns a tuple with the FromPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPrevious

`func (o *TourWayPointsInner) SetFromPrevious(v FromPrevious)`

SetFromPrevious sets FromPrevious field to given value.

### HasFromPrevious

`func (o *TourWayPointsInner) HasFromPrevious() bool`

HasFromPrevious returns a boolean if a field has been set.

### GetStopId

`func (o *TourWayPointsInner) GetStopId() string`

GetStopId returns the StopId field if non-nil, zero value otherwise.

### GetStopIdOk

`func (o *TourWayPointsInner) GetStopIdOk() (*string, bool)`

GetStopIdOk returns a tuple with the StopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopId

`func (o *TourWayPointsInner) SetStopId(v string)`

SetStopId sets StopId field to given value.

### HasStopId

`func (o *TourWayPointsInner) HasStopId() bool`

HasStopId returns a boolean if a field has been set.

### GetStopProperties

`func (o *TourWayPointsInner) GetStopProperties() map[string]string`

GetStopProperties returns the StopProperties field if non-nil, zero value otherwise.

### GetStopPropertiesOk

`func (o *TourWayPointsInner) GetStopPropertiesOk() (*map[string]string, bool)`

GetStopPropertiesOk returns a tuple with the StopProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProperties

`func (o *TourWayPointsInner) SetStopProperties(v map[string]string)`

SetStopProperties sets StopProperties field to given value.

### HasStopProperties

`func (o *TourWayPointsInner) HasStopProperties() bool`

HasStopProperties returns a boolean if a field has been set.

### GetOrderProperties

`func (o *TourWayPointsInner) GetOrderProperties() map[string]string`

GetOrderProperties returns the OrderProperties field if non-nil, zero value otherwise.

### GetOrderPropertiesOk

`func (o *TourWayPointsInner) GetOrderPropertiesOk() (*map[string]string, bool)`

GetOrderPropertiesOk returns a tuple with the OrderProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProperties

`func (o *TourWayPointsInner) SetOrderProperties(v map[string]string)`

SetOrderProperties sets OrderProperties field to given value.

### HasOrderProperties

`func (o *TourWayPointsInner) HasOrderProperties() bool`

HasOrderProperties returns a boolean if a field has been set.

### GetFilledCapacitiesAfterStop

`func (o *TourWayPointsInner) GetFilledCapacitiesAfterStop() map[string]float32`

GetFilledCapacitiesAfterStop returns the FilledCapacitiesAfterStop field if non-nil, zero value otherwise.

### GetFilledCapacitiesAfterStopOk

`func (o *TourWayPointsInner) GetFilledCapacitiesAfterStopOk() (*map[string]float32, bool)`

GetFilledCapacitiesAfterStopOk returns a tuple with the FilledCapacitiesAfterStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledCapacitiesAfterStop

`func (o *TourWayPointsInner) SetFilledCapacitiesAfterStop(v map[string]float32)`

SetFilledCapacitiesAfterStop sets FilledCapacitiesAfterStop field to given value.

### HasFilledCapacitiesAfterStop

`func (o *TourWayPointsInner) HasFilledCapacitiesAfterStop() bool`

HasFilledCapacitiesAfterStop returns a boolean if a field has been set.

### GetDelay

`func (o *TourWayPointsInner) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *TourWayPointsInner) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *TourWayPointsInner) SetDelay(v string)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *TourWayPointsInner) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetArrivalTime

`func (o *TourWayPointsInner) GetArrivalTime() string`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *TourWayPointsInner) GetArrivalTimeOk() (*string, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *TourWayPointsInner) SetArrivalTime(v string)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *TourWayPointsInner) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.

### GetBeginTime

`func (o *TourWayPointsInner) GetBeginTime() string`

GetBeginTime returns the BeginTime field if non-nil, zero value otherwise.

### GetBeginTimeOk

`func (o *TourWayPointsInner) GetBeginTimeOk() (*string, bool)`

GetBeginTimeOk returns a tuple with the BeginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginTime

`func (o *TourWayPointsInner) SetBeginTime(v string)`

SetBeginTime sets BeginTime field to given value.

### HasBeginTime

`func (o *TourWayPointsInner) HasBeginTime() bool`

HasBeginTime returns a boolean if a field has been set.

### GetStopKind

`func (o *TourWayPointsInner) GetStopKind() StopKind`

GetStopKind returns the StopKind field if non-nil, zero value otherwise.

### GetStopKindOk

`func (o *TourWayPointsInner) GetStopKindOk() (*StopKind, bool)`

GetStopKindOk returns a tuple with the StopKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopKind

`func (o *TourWayPointsInner) SetStopKind(v StopKind)`

SetStopKind sets StopKind field to given value.

### HasStopKind

`func (o *TourWayPointsInner) HasStopKind() bool`

HasStopKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


