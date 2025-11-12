# SimpleResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Resource ids must be unique within a plan. | 
**Capacities** | Pointer to **map[string]float32** |  | [optional] 
**Departure** | Pointer to [**Position**](Position.md) |  | [optional] 
**Arrival** | Pointer to [**ResourceArrival**](ResourceArrival.md) |  | [optional] 
**WorkingTimeWindow** | [**TimeWindow**](TimeWindow.md) |  | 
**MaxWorkingDuration** | Pointer to **string** | A period of time, expressed in the ISO8601 **duration** format. | [optional] 
**MaxDistanceInKm** | Pointer to **float32** |  | [optional] 
**Breaks** | Pointer to [**[]Break**](Break.md) |  | [optional] 

## Methods

### NewSimpleResource

`func NewSimpleResource(id string, workingTimeWindow TimeWindow, ) *SimpleResource`

NewSimpleResource instantiates a new SimpleResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleResourceWithDefaults

`func NewSimpleResourceWithDefaults() *SimpleResource`

NewSimpleResourceWithDefaults instantiates a new SimpleResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleResource) SetId(v string)`

SetId sets Id field to given value.


### GetCapacities

`func (o *SimpleResource) GetCapacities() map[string]float32`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *SimpleResource) GetCapacitiesOk() (*map[string]float32, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *SimpleResource) SetCapacities(v map[string]float32)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *SimpleResource) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetDeparture

`func (o *SimpleResource) GetDeparture() Position`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *SimpleResource) GetDepartureOk() (*Position, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *SimpleResource) SetDeparture(v Position)`

SetDeparture sets Departure field to given value.

### HasDeparture

`func (o *SimpleResource) HasDeparture() bool`

HasDeparture returns a boolean if a field has been set.

### GetArrival

`func (o *SimpleResource) GetArrival() ResourceArrival`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *SimpleResource) GetArrivalOk() (*ResourceArrival, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *SimpleResource) SetArrival(v ResourceArrival)`

SetArrival sets Arrival field to given value.

### HasArrival

`func (o *SimpleResource) HasArrival() bool`

HasArrival returns a boolean if a field has been set.

### GetWorkingTimeWindow

`func (o *SimpleResource) GetWorkingTimeWindow() TimeWindow`

GetWorkingTimeWindow returns the WorkingTimeWindow field if non-nil, zero value otherwise.

### GetWorkingTimeWindowOk

`func (o *SimpleResource) GetWorkingTimeWindowOk() (*TimeWindow, bool)`

GetWorkingTimeWindowOk returns a tuple with the WorkingTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeWindow

`func (o *SimpleResource) SetWorkingTimeWindow(v TimeWindow)`

SetWorkingTimeWindow sets WorkingTimeWindow field to given value.


### GetMaxWorkingDuration

`func (o *SimpleResource) GetMaxWorkingDuration() string`

GetMaxWorkingDuration returns the MaxWorkingDuration field if non-nil, zero value otherwise.

### GetMaxWorkingDurationOk

`func (o *SimpleResource) GetMaxWorkingDurationOk() (*string, bool)`

GetMaxWorkingDurationOk returns a tuple with the MaxWorkingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkingDuration

`func (o *SimpleResource) SetMaxWorkingDuration(v string)`

SetMaxWorkingDuration sets MaxWorkingDuration field to given value.

### HasMaxWorkingDuration

`func (o *SimpleResource) HasMaxWorkingDuration() bool`

HasMaxWorkingDuration returns a boolean if a field has been set.

### GetMaxDistanceInKm

`func (o *SimpleResource) GetMaxDistanceInKm() float32`

GetMaxDistanceInKm returns the MaxDistanceInKm field if non-nil, zero value otherwise.

### GetMaxDistanceInKmOk

`func (o *SimpleResource) GetMaxDistanceInKmOk() (*float32, bool)`

GetMaxDistanceInKmOk returns a tuple with the MaxDistanceInKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistanceInKm

`func (o *SimpleResource) SetMaxDistanceInKm(v float32)`

SetMaxDistanceInKm sets MaxDistanceInKm field to given value.

### HasMaxDistanceInKm

`func (o *SimpleResource) HasMaxDistanceInKm() bool`

HasMaxDistanceInKm returns a boolean if a field has been set.

### GetBreaks

`func (o *SimpleResource) GetBreaks() []Break`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *SimpleResource) GetBreaksOk() (*[]Break, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *SimpleResource) SetBreaks(v []Break)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *SimpleResource) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


