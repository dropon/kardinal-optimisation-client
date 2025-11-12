# TimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Begin** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 
**End** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 

## Methods

### NewTimeWindow

`func NewTimeWindow(begin string, end string, ) *TimeWindow`

NewTimeWindow instantiates a new TimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeWindowWithDefaults

`func NewTimeWindowWithDefaults() *TimeWindow`

NewTimeWindowWithDefaults instantiates a new TimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBegin

`func (o *TimeWindow) GetBegin() string`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *TimeWindow) GetBeginOk() (*string, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *TimeWindow) SetBegin(v string)`

SetBegin sets Begin field to given value.


### GetEnd

`func (o *TimeWindow) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TimeWindow) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TimeWindow) SetEnd(v string)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


