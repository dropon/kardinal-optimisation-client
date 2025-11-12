# AtLeastOneValidCapacityViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**WayPointIndex** | Pointer to **int32** | The idx of the waypoint the violation applies to. | [optional] 
**Name** | **string** | The name of the AtLeastOneValidCapacity constraint the violation is linked to. | 
**OverCapacities** | **map[string]float32** |  | 

## Methods

### NewAtLeastOneValidCapacityViolation

`func NewAtLeastOneValidCapacityViolation(type_ string, name string, overCapacities map[string]float32, ) *AtLeastOneValidCapacityViolation`

NewAtLeastOneValidCapacityViolation instantiates a new AtLeastOneValidCapacityViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtLeastOneValidCapacityViolationWithDefaults

`func NewAtLeastOneValidCapacityViolationWithDefaults() *AtLeastOneValidCapacityViolation`

NewAtLeastOneValidCapacityViolationWithDefaults instantiates a new AtLeastOneValidCapacityViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AtLeastOneValidCapacityViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AtLeastOneValidCapacityViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AtLeastOneValidCapacityViolation) SetType(v string)`

SetType sets Type field to given value.


### GetWayPointIndex

`func (o *AtLeastOneValidCapacityViolation) GetWayPointIndex() int32`

GetWayPointIndex returns the WayPointIndex field if non-nil, zero value otherwise.

### GetWayPointIndexOk

`func (o *AtLeastOneValidCapacityViolation) GetWayPointIndexOk() (*int32, bool)`

GetWayPointIndexOk returns a tuple with the WayPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayPointIndex

`func (o *AtLeastOneValidCapacityViolation) SetWayPointIndex(v int32)`

SetWayPointIndex sets WayPointIndex field to given value.

### HasWayPointIndex

`func (o *AtLeastOneValidCapacityViolation) HasWayPointIndex() bool`

HasWayPointIndex returns a boolean if a field has been set.

### GetName

`func (o *AtLeastOneValidCapacityViolation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtLeastOneValidCapacityViolation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtLeastOneValidCapacityViolation) SetName(v string)`

SetName sets Name field to given value.


### GetOverCapacities

`func (o *AtLeastOneValidCapacityViolation) GetOverCapacities() map[string]float32`

GetOverCapacities returns the OverCapacities field if non-nil, zero value otherwise.

### GetOverCapacitiesOk

`func (o *AtLeastOneValidCapacityViolation) GetOverCapacitiesOk() (*map[string]float32, bool)`

GetOverCapacitiesOk returns a tuple with the OverCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCapacities

`func (o *AtLeastOneValidCapacityViolation) SetOverCapacities(v map[string]float32)`

SetOverCapacities sets OverCapacities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


