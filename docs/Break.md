# Break

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Duration** | **string** | A period of time, expressed in the ISO8601 **duration** format. | 
**TimeWindow** | [**TimeWindow**](TimeWindow.md) |  | 
**MinBreakDuration** | **string** | A period of time, expressed in the ISO8601 **duration** format. | 
**MaxInterBreakDuration** | **string** | A period of time, expressed in the ISO8601 **duration** format. | 

## Methods

### NewBreak

`func NewBreak(type_ string, duration string, timeWindow TimeWindow, minBreakDuration string, maxInterBreakDuration string, ) *Break`

NewBreak instantiates a new Break object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakWithDefaults

`func NewBreakWithDefaults() *Break`

NewBreakWithDefaults instantiates a new Break object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Break) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Break) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Break) SetType(v string)`

SetType sets Type field to given value.


### GetDuration

`func (o *Break) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Break) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Break) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetTimeWindow

`func (o *Break) GetTimeWindow() TimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *Break) GetTimeWindowOk() (*TimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *Break) SetTimeWindow(v TimeWindow)`

SetTimeWindow sets TimeWindow field to given value.


### GetMinBreakDuration

`func (o *Break) GetMinBreakDuration() string`

GetMinBreakDuration returns the MinBreakDuration field if non-nil, zero value otherwise.

### GetMinBreakDurationOk

`func (o *Break) GetMinBreakDurationOk() (*string, bool)`

GetMinBreakDurationOk returns a tuple with the MinBreakDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBreakDuration

`func (o *Break) SetMinBreakDuration(v string)`

SetMinBreakDuration sets MinBreakDuration field to given value.


### GetMaxInterBreakDuration

`func (o *Break) GetMaxInterBreakDuration() string`

GetMaxInterBreakDuration returns the MaxInterBreakDuration field if non-nil, zero value otherwise.

### GetMaxInterBreakDurationOk

`func (o *Break) GetMaxInterBreakDurationOk() (*string, bool)`

GetMaxInterBreakDurationOk returns a tuple with the MaxInterBreakDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInterBreakDuration

`func (o *Break) SetMaxInterBreakDuration(v string)`

SetMaxInterBreakDuration sets MaxInterBreakDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


