# EndViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**WayPointIndex** | Pointer to **int32** | The idx of the waypoint the violation applies to. | [optional] 

## Methods

### NewEndViolation

`func NewEndViolation(type_ string, ) *EndViolation`

NewEndViolation instantiates a new EndViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndViolationWithDefaults

`func NewEndViolationWithDefaults() *EndViolation`

NewEndViolationWithDefaults instantiates a new EndViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EndViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EndViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EndViolation) SetType(v string)`

SetType sets Type field to given value.


### GetWayPointIndex

`func (o *EndViolation) GetWayPointIndex() int32`

GetWayPointIndex returns the WayPointIndex field if non-nil, zero value otherwise.

### GetWayPointIndexOk

`func (o *EndViolation) GetWayPointIndexOk() (*int32, bool)`

GetWayPointIndexOk returns a tuple with the WayPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayPointIndex

`func (o *EndViolation) SetWayPointIndex(v int32)`

SetWayPointIndex sets WayPointIndex field to given value.

### HasWayPointIndex

`func (o *EndViolation) HasWayPointIndex() bool`

HasWayPointIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


