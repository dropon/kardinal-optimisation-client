# CapacityViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**WayPointIndex** | Pointer to **int32** | The idx of the waypoint the violation applies to. | [optional] 
**Capacity** | **string** | The capacity&#39;s name the constraint is linked to. | 
**OverCapacity** | **float32** | The exceeding capacity. | 

## Methods

### NewCapacityViolation

`func NewCapacityViolation(type_ string, capacity string, overCapacity float32, ) *CapacityViolation`

NewCapacityViolation instantiates a new CapacityViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityViolationWithDefaults

`func NewCapacityViolationWithDefaults() *CapacityViolation`

NewCapacityViolationWithDefaults instantiates a new CapacityViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CapacityViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CapacityViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CapacityViolation) SetType(v string)`

SetType sets Type field to given value.


### GetWayPointIndex

`func (o *CapacityViolation) GetWayPointIndex() int32`

GetWayPointIndex returns the WayPointIndex field if non-nil, zero value otherwise.

### GetWayPointIndexOk

`func (o *CapacityViolation) GetWayPointIndexOk() (*int32, bool)`

GetWayPointIndexOk returns a tuple with the WayPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayPointIndex

`func (o *CapacityViolation) SetWayPointIndex(v int32)`

SetWayPointIndex sets WayPointIndex field to given value.

### HasWayPointIndex

`func (o *CapacityViolation) HasWayPointIndex() bool`

HasWayPointIndex returns a boolean if a field has been set.

### GetCapacity

`func (o *CapacityViolation) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CapacityViolation) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CapacityViolation) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.


### GetOverCapacity

`func (o *CapacityViolation) GetOverCapacity() float32`

GetOverCapacity returns the OverCapacity field if non-nil, zero value otherwise.

### GetOverCapacityOk

`func (o *CapacityViolation) GetOverCapacityOk() (*float32, bool)`

GetOverCapacityOk returns a tuple with the OverCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCapacity

`func (o *CapacityViolation) SetOverCapacity(v float32)`

SetOverCapacity sets OverCapacity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


