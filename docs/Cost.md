# Cost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkedHours** | Pointer to [**CostFloorsAndCoeffs**](CostFloorsAndCoeffs.md) |  | [optional] 
**Km** | Pointer to [**CostFloorsAndCoeffs**](CostFloorsAndCoeffs.md) |  | [optional] 
**Using** | Pointer to **float32** |  | [optional] 
**CostsByStopTag** | Pointer to [**map[string]CostFloorsAndCoeffs**](CostFloorsAndCoeffs.md) |  | [optional] 
**CostsByCapacity** | Pointer to [**map[string]TaggedCostFloorsAndCoeffs**](TaggedCostFloorsAndCoeffs.md) |  | [optional] 

## Methods

### NewCost

`func NewCost() *Cost`

NewCost instantiates a new Cost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostWithDefaults

`func NewCostWithDefaults() *Cost`

NewCostWithDefaults instantiates a new Cost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkedHours

`func (o *Cost) GetWorkedHours() CostFloorsAndCoeffs`

GetWorkedHours returns the WorkedHours field if non-nil, zero value otherwise.

### GetWorkedHoursOk

`func (o *Cost) GetWorkedHoursOk() (*CostFloorsAndCoeffs, bool)`

GetWorkedHoursOk returns a tuple with the WorkedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkedHours

`func (o *Cost) SetWorkedHours(v CostFloorsAndCoeffs)`

SetWorkedHours sets WorkedHours field to given value.

### HasWorkedHours

`func (o *Cost) HasWorkedHours() bool`

HasWorkedHours returns a boolean if a field has been set.

### GetKm

`func (o *Cost) GetKm() CostFloorsAndCoeffs`

GetKm returns the Km field if non-nil, zero value otherwise.

### GetKmOk

`func (o *Cost) GetKmOk() (*CostFloorsAndCoeffs, bool)`

GetKmOk returns a tuple with the Km field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKm

`func (o *Cost) SetKm(v CostFloorsAndCoeffs)`

SetKm sets Km field to given value.

### HasKm

`func (o *Cost) HasKm() bool`

HasKm returns a boolean if a field has been set.

### GetUsing

`func (o *Cost) GetUsing() float32`

GetUsing returns the Using field if non-nil, zero value otherwise.

### GetUsingOk

`func (o *Cost) GetUsingOk() (*float32, bool)`

GetUsingOk returns a tuple with the Using field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsing

`func (o *Cost) SetUsing(v float32)`

SetUsing sets Using field to given value.

### HasUsing

`func (o *Cost) HasUsing() bool`

HasUsing returns a boolean if a field has been set.

### GetCostsByStopTag

`func (o *Cost) GetCostsByStopTag() map[string]CostFloorsAndCoeffs`

GetCostsByStopTag returns the CostsByStopTag field if non-nil, zero value otherwise.

### GetCostsByStopTagOk

`func (o *Cost) GetCostsByStopTagOk() (*map[string]CostFloorsAndCoeffs, bool)`

GetCostsByStopTagOk returns a tuple with the CostsByStopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostsByStopTag

`func (o *Cost) SetCostsByStopTag(v map[string]CostFloorsAndCoeffs)`

SetCostsByStopTag sets CostsByStopTag field to given value.

### HasCostsByStopTag

`func (o *Cost) HasCostsByStopTag() bool`

HasCostsByStopTag returns a boolean if a field has been set.

### GetCostsByCapacity

`func (o *Cost) GetCostsByCapacity() map[string]TaggedCostFloorsAndCoeffs`

GetCostsByCapacity returns the CostsByCapacity field if non-nil, zero value otherwise.

### GetCostsByCapacityOk

`func (o *Cost) GetCostsByCapacityOk() (*map[string]TaggedCostFloorsAndCoeffs, bool)`

GetCostsByCapacityOk returns a tuple with the CostsByCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostsByCapacity

`func (o *Cost) SetCostsByCapacity(v map[string]TaggedCostFloorsAndCoeffs)`

SetCostsByCapacity sets CostsByCapacity field to given value.

### HasCostsByCapacity

`func (o *Cost) HasCostsByCapacity() bool`

HasCostsByCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


