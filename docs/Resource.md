# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Resource ids must be unique within a plan. | 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Priority** | Pointer to **int32** | 0 by default, can be negative. | [optional] [default to 0]
**Skills** | Pointer to **[]string** |  | [optional] 
**PreferredStopTags** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**VehicleProfile** | [**ResourceVehicleProfile**](ResourceVehicleProfile.md) |  | 
**Capacities** | Pointer to **map[string]float32** |  | [optional] 
**Departure** | Pointer to [**Position**](Position.md) |  | [optional] 
**Arrival** | Pointer to [**ResourceArrival**](ResourceArrival.md) |  | [optional] 
**WorkingTimeWindow** | [**TimeWindow**](TimeWindow.md) |  | 
**MaxWorkingDuration** | Pointer to **string** | A period of time, expressed in the ISO8601 **duration** format. | [optional] 
**MaxDistanceInKm** | Pointer to **float32** |  | [optional] 
**Breaks** | Pointer to [**[]Break**](Break.md) |  | [optional] 
**OperationDurationPolicies** | Pointer to [**[]OperationDurationPolicy**](OperationDurationPolicy.md) |  | [optional] 
**TravelTimeCoefficient** | Pointer to **float32** |  | [optional] 

## Methods

### NewResource

`func NewResource(id string, vehicleProfile ResourceVehicleProfile, workingTimeWindow TimeWindow, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *Resource) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Resource) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Resource) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Resource) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetState

`func (o *Resource) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Resource) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Resource) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *Resource) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCost

`func (o *Resource) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Resource) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Resource) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Resource) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPriority

`func (o *Resource) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Resource) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Resource) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Resource) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSkills

`func (o *Resource) GetSkills() []string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *Resource) GetSkillsOk() (*[]string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *Resource) SetSkills(v []string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *Resource) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetPreferredStopTags

`func (o *Resource) GetPreferredStopTags() []string`

GetPreferredStopTags returns the PreferredStopTags field if non-nil, zero value otherwise.

### GetPreferredStopTagsOk

`func (o *Resource) GetPreferredStopTagsOk() (*[]string, bool)`

GetPreferredStopTagsOk returns a tuple with the PreferredStopTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredStopTags

`func (o *Resource) SetPreferredStopTags(v []string)`

SetPreferredStopTags sets PreferredStopTags field to given value.

### HasPreferredStopTags

`func (o *Resource) HasPreferredStopTags() bool`

HasPreferredStopTags returns a boolean if a field has been set.

### GetTags

`func (o *Resource) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Resource) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Resource) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Resource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVehicleProfile

`func (o *Resource) GetVehicleProfile() ResourceVehicleProfile`

GetVehicleProfile returns the VehicleProfile field if non-nil, zero value otherwise.

### GetVehicleProfileOk

`func (o *Resource) GetVehicleProfileOk() (*ResourceVehicleProfile, bool)`

GetVehicleProfileOk returns a tuple with the VehicleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleProfile

`func (o *Resource) SetVehicleProfile(v ResourceVehicleProfile)`

SetVehicleProfile sets VehicleProfile field to given value.


### GetCapacities

`func (o *Resource) GetCapacities() map[string]float32`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *Resource) GetCapacitiesOk() (*map[string]float32, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *Resource) SetCapacities(v map[string]float32)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *Resource) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetDeparture

`func (o *Resource) GetDeparture() Position`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *Resource) GetDepartureOk() (*Position, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *Resource) SetDeparture(v Position)`

SetDeparture sets Departure field to given value.

### HasDeparture

`func (o *Resource) HasDeparture() bool`

HasDeparture returns a boolean if a field has been set.

### GetArrival

`func (o *Resource) GetArrival() ResourceArrival`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *Resource) GetArrivalOk() (*ResourceArrival, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *Resource) SetArrival(v ResourceArrival)`

SetArrival sets Arrival field to given value.

### HasArrival

`func (o *Resource) HasArrival() bool`

HasArrival returns a boolean if a field has been set.

### GetWorkingTimeWindow

`func (o *Resource) GetWorkingTimeWindow() TimeWindow`

GetWorkingTimeWindow returns the WorkingTimeWindow field if non-nil, zero value otherwise.

### GetWorkingTimeWindowOk

`func (o *Resource) GetWorkingTimeWindowOk() (*TimeWindow, bool)`

GetWorkingTimeWindowOk returns a tuple with the WorkingTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeWindow

`func (o *Resource) SetWorkingTimeWindow(v TimeWindow)`

SetWorkingTimeWindow sets WorkingTimeWindow field to given value.


### GetMaxWorkingDuration

`func (o *Resource) GetMaxWorkingDuration() string`

GetMaxWorkingDuration returns the MaxWorkingDuration field if non-nil, zero value otherwise.

### GetMaxWorkingDurationOk

`func (o *Resource) GetMaxWorkingDurationOk() (*string, bool)`

GetMaxWorkingDurationOk returns a tuple with the MaxWorkingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkingDuration

`func (o *Resource) SetMaxWorkingDuration(v string)`

SetMaxWorkingDuration sets MaxWorkingDuration field to given value.

### HasMaxWorkingDuration

`func (o *Resource) HasMaxWorkingDuration() bool`

HasMaxWorkingDuration returns a boolean if a field has been set.

### GetMaxDistanceInKm

`func (o *Resource) GetMaxDistanceInKm() float32`

GetMaxDistanceInKm returns the MaxDistanceInKm field if non-nil, zero value otherwise.

### GetMaxDistanceInKmOk

`func (o *Resource) GetMaxDistanceInKmOk() (*float32, bool)`

GetMaxDistanceInKmOk returns a tuple with the MaxDistanceInKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistanceInKm

`func (o *Resource) SetMaxDistanceInKm(v float32)`

SetMaxDistanceInKm sets MaxDistanceInKm field to given value.

### HasMaxDistanceInKm

`func (o *Resource) HasMaxDistanceInKm() bool`

HasMaxDistanceInKm returns a boolean if a field has been set.

### GetBreaks

`func (o *Resource) GetBreaks() []Break`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *Resource) GetBreaksOk() (*[]Break, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *Resource) SetBreaks(v []Break)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *Resource) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.

### GetOperationDurationPolicies

`func (o *Resource) GetOperationDurationPolicies() []OperationDurationPolicy`

GetOperationDurationPolicies returns the OperationDurationPolicies field if non-nil, zero value otherwise.

### GetOperationDurationPoliciesOk

`func (o *Resource) GetOperationDurationPoliciesOk() (*[]OperationDurationPolicy, bool)`

GetOperationDurationPoliciesOk returns a tuple with the OperationDurationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationDurationPolicies

`func (o *Resource) SetOperationDurationPolicies(v []OperationDurationPolicy)`

SetOperationDurationPolicies sets OperationDurationPolicies field to given value.

### HasOperationDurationPolicies

`func (o *Resource) HasOperationDurationPolicies() bool`

HasOperationDurationPolicies returns a boolean if a field has been set.

### GetTravelTimeCoefficient

`func (o *Resource) GetTravelTimeCoefficient() float32`

GetTravelTimeCoefficient returns the TravelTimeCoefficient field if non-nil, zero value otherwise.

### GetTravelTimeCoefficientOk

`func (o *Resource) GetTravelTimeCoefficientOk() (*float32, bool)`

GetTravelTimeCoefficientOk returns a tuple with the TravelTimeCoefficient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTimeCoefficient

`func (o *Resource) SetTravelTimeCoefficient(v float32)`

SetTravelTimeCoefficient sets TravelTimeCoefficient field to given value.

### HasTravelTimeCoefficient

`func (o *Resource) HasTravelTimeCoefficient() bool`

HasTravelTimeCoefficient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


