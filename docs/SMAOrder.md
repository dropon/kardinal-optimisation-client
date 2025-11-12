# SMAOrder

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

### NewSMAOrder

`func NewSMAOrder(id string, stops []Stop, ) *SMAOrder`

NewSMAOrder instantiates a new SMAOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMAOrderWithDefaults

`func NewSMAOrderWithDefaults() *SMAOrder`

NewSMAOrderWithDefaults instantiates a new SMAOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SMAOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SMAOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SMAOrder) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *SMAOrder) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SMAOrder) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SMAOrder) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SMAOrder) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPriority

`func (o *SMAOrder) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SMAOrder) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SMAOrder) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SMAOrder) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetOptional

`func (o *SMAOrder) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *SMAOrder) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *SMAOrder) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *SMAOrder) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetRequiredSkills

`func (o *SMAOrder) GetRequiredSkills() []string`

GetRequiredSkills returns the RequiredSkills field if non-nil, zero value otherwise.

### GetRequiredSkillsOk

`func (o *SMAOrder) GetRequiredSkillsOk() (*[]string, bool)`

GetRequiredSkillsOk returns a tuple with the RequiredSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSkills

`func (o *SMAOrder) SetRequiredSkills(v []string)`

SetRequiredSkills sets RequiredSkills field to given value.

### HasRequiredSkills

`func (o *SMAOrder) HasRequiredSkills() bool`

HasRequiredSkills returns a boolean if a field has been set.

### GetStops

`func (o *SMAOrder) GetStops() []Stop`

GetStops returns the Stops field if non-nil, zero value otherwise.

### GetStopsOk

`func (o *SMAOrder) GetStopsOk() (*[]Stop, bool)`

GetStopsOk returns a tuple with the Stops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStops

`func (o *SMAOrder) SetStops(v []Stop)`

SetStops sets Stops field to given value.


### GetSuccessiveStops

`func (o *SMAOrder) GetSuccessiveStops() bool`

GetSuccessiveStops returns the SuccessiveStops field if non-nil, zero value otherwise.

### GetSuccessiveStopsOk

`func (o *SMAOrder) GetSuccessiveStopsOk() (*bool, bool)`

GetSuccessiveStopsOk returns a tuple with the SuccessiveStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessiveStops

`func (o *SMAOrder) SetSuccessiveStops(v bool)`

SetSuccessiveStops sets SuccessiveStops field to given value.

### HasSuccessiveStops

`func (o *SMAOrder) HasSuccessiveStops() bool`

HasSuccessiveStops returns a boolean if a field has been set.

### GetMaxStopSpan

`func (o *SMAOrder) GetMaxStopSpan() string`

GetMaxStopSpan returns the MaxStopSpan field if non-nil, zero value otherwise.

### GetMaxStopSpanOk

`func (o *SMAOrder) GetMaxStopSpanOk() (*string, bool)`

GetMaxStopSpanOk returns a tuple with the MaxStopSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStopSpan

`func (o *SMAOrder) SetMaxStopSpan(v string)`

SetMaxStopSpan sets MaxStopSpan field to given value.

### HasMaxStopSpan

`func (o *SMAOrder) HasMaxStopSpan() bool`

HasMaxStopSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


