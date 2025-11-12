# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Order ids must be unique within a plan. | 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** | 0 by default, can be negative. | [optional] [default to 0]
**Optional** | Pointer to **bool** |  | [optional] 
**RequiredSkills** | Pointer to **[]string** |  | [optional] 
**Stops** | [**[]Stop**](Stop.md) |  | 
**SuccessiveStops** | Pointer to **bool** |  | [optional] 
**MaxStopSpan** | Pointer to **string** | A period of time, expressed in the ISO8601 **duration** format. | [optional] 

## Methods

### NewOrder

`func NewOrder(id string, stops []Stop, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *Order) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Order) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Order) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Order) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPriority

`func (o *Order) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Order) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Order) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Order) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOptional

`func (o *Order) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *Order) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *Order) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *Order) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetRequiredSkills

`func (o *Order) GetRequiredSkills() []string`

GetRequiredSkills returns the RequiredSkills field if non-nil, zero value otherwise.

### GetRequiredSkillsOk

`func (o *Order) GetRequiredSkillsOk() (*[]string, bool)`

GetRequiredSkillsOk returns a tuple with the RequiredSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSkills

`func (o *Order) SetRequiredSkills(v []string)`

SetRequiredSkills sets RequiredSkills field to given value.

### HasRequiredSkills

`func (o *Order) HasRequiredSkills() bool`

HasRequiredSkills returns a boolean if a field has been set.

### GetStops

`func (o *Order) GetStops() []Stop`

GetStops returns the Stops field if non-nil, zero value otherwise.

### GetStopsOk

`func (o *Order) GetStopsOk() (*[]Stop, bool)`

GetStopsOk returns a tuple with the Stops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStops

`func (o *Order) SetStops(v []Stop)`

SetStops sets Stops field to given value.


### GetSuccessiveStops

`func (o *Order) GetSuccessiveStops() bool`

GetSuccessiveStops returns the SuccessiveStops field if non-nil, zero value otherwise.

### GetSuccessiveStopsOk

`func (o *Order) GetSuccessiveStopsOk() (*bool, bool)`

GetSuccessiveStopsOk returns a tuple with the SuccessiveStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessiveStops

`func (o *Order) SetSuccessiveStops(v bool)`

SetSuccessiveStops sets SuccessiveStops field to given value.

### HasSuccessiveStops

`func (o *Order) HasSuccessiveStops() bool`

HasSuccessiveStops returns a boolean if a field has been set.

### GetMaxStopSpan

`func (o *Order) GetMaxStopSpan() string`

GetMaxStopSpan returns the MaxStopSpan field if non-nil, zero value otherwise.

### GetMaxStopSpanOk

`func (o *Order) GetMaxStopSpanOk() (*string, bool)`

GetMaxStopSpanOk returns a tuple with the MaxStopSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStopSpan

`func (o *Order) SetMaxStopSpan(v string)`

SetMaxStopSpan sets MaxStopSpan field to given value.

### HasMaxStopSpan

`func (o *Order) HasMaxStopSpan() bool`

HasMaxStopSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


