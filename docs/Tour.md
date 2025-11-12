# Tour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | [optional] 
**Mode** | Pointer to [**ResourceMode**](ResourceMode.md) |  | [optional] [default to RESOURCEMODE_FREE]
**DistanceInKm** | Pointer to **float32** |  | [optional] 
**TourCost** | Pointer to **float32** |  | [optional] 
**TotalDelay** | Pointer to **string** | Total delay of the tour related to the preferred time windows. | [optional] 
**TravelDuration** | Pointer to **string** | A period of time, expressed in the ISO8601 **duration** format. | [optional] 
**WorkingDuration** | Pointer to **string** | A period of time, expressed in the ISO8601 **duration** format. | [optional] 
**WaitingDuration** | Pointer to **string** | A period of time, expressed in the ISO8601 **duration** format. | [optional] 
**ResourceCapacities** | Pointer to **map[string]float32** | List of capacities and associated quantity for which the resource has a restriction on the quantity to be carried at each stop during the tour. | [optional] 
**MaxFilledCapacities** | Pointer to **map[string]float32** | List of each capacity transported by the resource, as well as the maximum quantity reached for this capacity at a stop during the tour. | [optional] 
**FilledCapacitiesAtBegin** | Pointer to **map[string]float32** |  | [optional] 
**FilledCapacitiesAtEnd** | Pointer to **map[string]float32** |  | [optional] 
**IsValid** | Pointer to **bool** | Indicates whether the tour satisfies all resource constraints. | [optional] 
**Violations** | Pointer to [**[]TourViolation**](TourViolation.md) | List of the Tour&#39;s violations. | [optional] 
**WayPoints** | Pointer to [**[]TourWayPointsInner**](TourWayPointsInner.md) |  | [optional] 
**VehicleProfile** | Pointer to [**ResourceVehicleProfile**](ResourceVehicleProfile.md) |  | [optional] 

## Methods

### NewTour

`func NewTour() *Tour`

NewTour instantiates a new Tour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTourWithDefaults

`func NewTourWithDefaults() *Tour`

NewTourWithDefaults instantiates a new Tour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *Tour) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Tour) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Tour) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Tour) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetMode

`func (o *Tour) GetMode() ResourceMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Tour) GetModeOk() (*ResourceMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Tour) SetMode(v ResourceMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Tour) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDistanceInKm

`func (o *Tour) GetDistanceInKm() float32`

GetDistanceInKm returns the DistanceInKm field if non-nil, zero value otherwise.

### GetDistanceInKmOk

`func (o *Tour) GetDistanceInKmOk() (*float32, bool)`

GetDistanceInKmOk returns a tuple with the DistanceInKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceInKm

`func (o *Tour) SetDistanceInKm(v float32)`

SetDistanceInKm sets DistanceInKm field to given value.

### HasDistanceInKm

`func (o *Tour) HasDistanceInKm() bool`

HasDistanceInKm returns a boolean if a field has been set.

### GetTourCost

`func (o *Tour) GetTourCost() float32`

GetTourCost returns the TourCost field if non-nil, zero value otherwise.

### GetTourCostOk

`func (o *Tour) GetTourCostOk() (*float32, bool)`

GetTourCostOk returns a tuple with the TourCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTourCost

`func (o *Tour) SetTourCost(v float32)`

SetTourCost sets TourCost field to given value.

### HasTourCost

`func (o *Tour) HasTourCost() bool`

HasTourCost returns a boolean if a field has been set.

### GetTotalDelay

`func (o *Tour) GetTotalDelay() string`

GetTotalDelay returns the TotalDelay field if non-nil, zero value otherwise.

### GetTotalDelayOk

`func (o *Tour) GetTotalDelayOk() (*string, bool)`

GetTotalDelayOk returns a tuple with the TotalDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDelay

`func (o *Tour) SetTotalDelay(v string)`

SetTotalDelay sets TotalDelay field to given value.

### HasTotalDelay

`func (o *Tour) HasTotalDelay() bool`

HasTotalDelay returns a boolean if a field has been set.

### GetTravelDuration

`func (o *Tour) GetTravelDuration() string`

GetTravelDuration returns the TravelDuration field if non-nil, zero value otherwise.

### GetTravelDurationOk

`func (o *Tour) GetTravelDurationOk() (*string, bool)`

GetTravelDurationOk returns a tuple with the TravelDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelDuration

`func (o *Tour) SetTravelDuration(v string)`

SetTravelDuration sets TravelDuration field to given value.

### HasTravelDuration

`func (o *Tour) HasTravelDuration() bool`

HasTravelDuration returns a boolean if a field has been set.

### GetWorkingDuration

`func (o *Tour) GetWorkingDuration() string`

GetWorkingDuration returns the WorkingDuration field if non-nil, zero value otherwise.

### GetWorkingDurationOk

`func (o *Tour) GetWorkingDurationOk() (*string, bool)`

GetWorkingDurationOk returns a tuple with the WorkingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDuration

`func (o *Tour) SetWorkingDuration(v string)`

SetWorkingDuration sets WorkingDuration field to given value.

### HasWorkingDuration

`func (o *Tour) HasWorkingDuration() bool`

HasWorkingDuration returns a boolean if a field has been set.

### GetWaitingDuration

`func (o *Tour) GetWaitingDuration() string`

GetWaitingDuration returns the WaitingDuration field if non-nil, zero value otherwise.

### GetWaitingDurationOk

`func (o *Tour) GetWaitingDurationOk() (*string, bool)`

GetWaitingDurationOk returns a tuple with the WaitingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingDuration

`func (o *Tour) SetWaitingDuration(v string)`

SetWaitingDuration sets WaitingDuration field to given value.

### HasWaitingDuration

`func (o *Tour) HasWaitingDuration() bool`

HasWaitingDuration returns a boolean if a field has been set.

### GetResourceCapacities

`func (o *Tour) GetResourceCapacities() map[string]float32`

GetResourceCapacities returns the ResourceCapacities field if non-nil, zero value otherwise.

### GetResourceCapacitiesOk

`func (o *Tour) GetResourceCapacitiesOk() (*map[string]float32, bool)`

GetResourceCapacitiesOk returns a tuple with the ResourceCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCapacities

`func (o *Tour) SetResourceCapacities(v map[string]float32)`

SetResourceCapacities sets ResourceCapacities field to given value.

### HasResourceCapacities

`func (o *Tour) HasResourceCapacities() bool`

HasResourceCapacities returns a boolean if a field has been set.

### GetMaxFilledCapacities

`func (o *Tour) GetMaxFilledCapacities() map[string]float32`

GetMaxFilledCapacities returns the MaxFilledCapacities field if non-nil, zero value otherwise.

### GetMaxFilledCapacitiesOk

`func (o *Tour) GetMaxFilledCapacitiesOk() (*map[string]float32, bool)`

GetMaxFilledCapacitiesOk returns a tuple with the MaxFilledCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFilledCapacities

`func (o *Tour) SetMaxFilledCapacities(v map[string]float32)`

SetMaxFilledCapacities sets MaxFilledCapacities field to given value.

### HasMaxFilledCapacities

`func (o *Tour) HasMaxFilledCapacities() bool`

HasMaxFilledCapacities returns a boolean if a field has been set.

### GetFilledCapacitiesAtBegin

`func (o *Tour) GetFilledCapacitiesAtBegin() map[string]float32`

GetFilledCapacitiesAtBegin returns the FilledCapacitiesAtBegin field if non-nil, zero value otherwise.

### GetFilledCapacitiesAtBeginOk

`func (o *Tour) GetFilledCapacitiesAtBeginOk() (*map[string]float32, bool)`

GetFilledCapacitiesAtBeginOk returns a tuple with the FilledCapacitiesAtBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledCapacitiesAtBegin

`func (o *Tour) SetFilledCapacitiesAtBegin(v map[string]float32)`

SetFilledCapacitiesAtBegin sets FilledCapacitiesAtBegin field to given value.

### HasFilledCapacitiesAtBegin

`func (o *Tour) HasFilledCapacitiesAtBegin() bool`

HasFilledCapacitiesAtBegin returns a boolean if a field has been set.

### GetFilledCapacitiesAtEnd

`func (o *Tour) GetFilledCapacitiesAtEnd() map[string]float32`

GetFilledCapacitiesAtEnd returns the FilledCapacitiesAtEnd field if non-nil, zero value otherwise.

### GetFilledCapacitiesAtEndOk

`func (o *Tour) GetFilledCapacitiesAtEndOk() (*map[string]float32, bool)`

GetFilledCapacitiesAtEndOk returns a tuple with the FilledCapacitiesAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledCapacitiesAtEnd

`func (o *Tour) SetFilledCapacitiesAtEnd(v map[string]float32)`

SetFilledCapacitiesAtEnd sets FilledCapacitiesAtEnd field to given value.

### HasFilledCapacitiesAtEnd

`func (o *Tour) HasFilledCapacitiesAtEnd() bool`

HasFilledCapacitiesAtEnd returns a boolean if a field has been set.

### GetIsValid

`func (o *Tour) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *Tour) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *Tour) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *Tour) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetViolations

`func (o *Tour) GetViolations() []TourViolation`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *Tour) GetViolationsOk() (*[]TourViolation, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *Tour) SetViolations(v []TourViolation)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *Tour) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### GetWayPoints

`func (o *Tour) GetWayPoints() []TourWayPointsInner`

GetWayPoints returns the WayPoints field if non-nil, zero value otherwise.

### GetWayPointsOk

`func (o *Tour) GetWayPointsOk() (*[]TourWayPointsInner, bool)`

GetWayPointsOk returns a tuple with the WayPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayPoints

`func (o *Tour) SetWayPoints(v []TourWayPointsInner)`

SetWayPoints sets WayPoints field to given value.

### HasWayPoints

`func (o *Tour) HasWayPoints() bool`

HasWayPoints returns a boolean if a field has been set.

### GetVehicleProfile

`func (o *Tour) GetVehicleProfile() ResourceVehicleProfile`

GetVehicleProfile returns the VehicleProfile field if non-nil, zero value otherwise.

### GetVehicleProfileOk

`func (o *Tour) GetVehicleProfileOk() (*ResourceVehicleProfile, bool)`

GetVehicleProfileOk returns a tuple with the VehicleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleProfile

`func (o *Tour) SetVehicleProfile(v ResourceVehicleProfile)`

SetVehicleProfile sets VehicleProfile field to given value.

### HasVehicleProfile

`func (o *Tour) HasVehicleProfile() bool`

HasVehicleProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


