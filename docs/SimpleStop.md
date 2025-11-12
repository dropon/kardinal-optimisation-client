# SimpleStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Single stop ids must be unique within a plan. | 
**Position** | [**Position**](Position.md) |  | 
**Kind** | Pointer to [**StopKind**](StopKind.md) |  | [optional] [default to STOPKIND_DELIVERY]
**OperationDuration** | Pointer to **string** | A period of time, expressed in the ISO8601 **duration** format. | [optional] 
**Capacities** | Pointer to **map[string]float32** |  | [optional] 
**TimeWindow** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 

## Methods

### NewSimpleStop

`func NewSimpleStop(id string, position Position, ) *SimpleStop`

NewSimpleStop instantiates a new SimpleStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleStopWithDefaults

`func NewSimpleStopWithDefaults() *SimpleStop`

NewSimpleStopWithDefaults instantiates a new SimpleStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleStop) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleStop) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleStop) SetId(v string)`

SetId sets Id field to given value.


### GetPosition

`func (o *SimpleStop) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SimpleStop) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SimpleStop) SetPosition(v Position)`

SetPosition sets Position field to given value.


### GetKind

`func (o *SimpleStop) GetKind() StopKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SimpleStop) GetKindOk() (*StopKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SimpleStop) SetKind(v StopKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SimpleStop) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOperationDuration

`func (o *SimpleStop) GetOperationDuration() string`

GetOperationDuration returns the OperationDuration field if non-nil, zero value otherwise.

### GetOperationDurationOk

`func (o *SimpleStop) GetOperationDurationOk() (*string, bool)`

GetOperationDurationOk returns a tuple with the OperationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationDuration

`func (o *SimpleStop) SetOperationDuration(v string)`

SetOperationDuration sets OperationDuration field to given value.

### HasOperationDuration

`func (o *SimpleStop) HasOperationDuration() bool`

HasOperationDuration returns a boolean if a field has been set.

### GetCapacities

`func (o *SimpleStop) GetCapacities() map[string]float32`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *SimpleStop) GetCapacitiesOk() (*map[string]float32, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *SimpleStop) SetCapacities(v map[string]float32)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *SimpleStop) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetTimeWindow

`func (o *SimpleStop) GetTimeWindow() TimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *SimpleStop) GetTimeWindowOk() (*TimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *SimpleStop) SetTimeWindow(v TimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *SimpleStop) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


