# TimeWindowBreak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Duration** | **string** | A period of time, expressed in the ISO8601 **duration** format. | 
**TimeWindow** | [**TimeWindow**](TimeWindow.md) |  | 

## Methods

### NewTimeWindowBreak

`func NewTimeWindowBreak(duration string, timeWindow TimeWindow, ) *TimeWindowBreak`

NewTimeWindowBreak instantiates a new TimeWindowBreak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeWindowBreakWithDefaults

`func NewTimeWindowBreakWithDefaults() *TimeWindowBreak`

NewTimeWindowBreakWithDefaults instantiates a new TimeWindowBreak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimeWindowBreak) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeWindowBreak) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeWindowBreak) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TimeWindowBreak) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDuration

`func (o *TimeWindowBreak) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TimeWindowBreak) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TimeWindowBreak) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetTimeWindow

`func (o *TimeWindowBreak) GetTimeWindow() TimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *TimeWindowBreak) GetTimeWindowOk() (*TimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *TimeWindowBreak) SetTimeWindow(v TimeWindow)`

SetTimeWindow sets TimeWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


