# SMAInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**SMAOrder**](SMAOrder.md) |  | 
**MaxNbPropositions** | Pointer to **NullableInt32** |  | [optional] 
**TimeWindows** | [**[]TimeWindow**](TimeWindow.md) |  | 
**Tz** | Pointer to **string** | The time zone is a string code which identifies a region of the world in the \&quot;time zone database\&quot;, also called \&quot;tz database\&quot;. The tz database is a partition of the world into regions where local clocks all show the same time. This database gives the rules for time offset and daylight saving time in each region.  How do we use it?  In order to work with time events accurately, we usually use datetimes in the iso-8601 format, without explicit time zone. This format is quite well suported by many programming languages, and it is well suited for technical data exchange. But it is not easy to use for humans.  For instance, here are three datetimes in iso-8601 format, which give the same exact moment in time: - \&quot;2025-05-22T05:43:00Z\&quot; - \&quot;2025-05-22T06:43:00+01:00\&quot; - \&quot;2025-05-22T07:43:00+02:00\&quot;  For a non-technical user, it is difficult to know how to relate this to the time displayed on a watch or a clock.  We improve the user experience by adding the support of local datetimes, thanks to the use of the time zone, which allows to transform a local datetime into an iso-8601 datetime: - local datetime + timezone (tz) &#x3D; iso-8601 datetime  For instance, here are five datetimes which all give the same exact moment in time: - \&quot;2025-05-22T05:43:00Z\&quot; - \&quot;2025-05-22T06:43:00+01:00\&quot; - \&quot;2025-05-22T07:43:00+02:00\&quot; - \&quot;2025-05-22 07:43:00\&quot;       + timezone \&quot;tz\&quot;: \&quot;Europe/Paris\&quot; - \&quot;2025-05-22 07:43\&quot;          + timezone \&quot;tz\&quot;: \&quot;Europe/Paris\&quot;  Note: the last example (\&quot;2025-05-22 07:43\&quot;) illustrates the support of local datetimes without seconds, which can be very practical for users.  In order for local datetimes to be supported, some JSON input objects contain a \&quot;tz\&quot; time zone property. This \&quot;tz\&quot; property is used to pre-process the JSON input payload, like this: - We check if a valid timezone can be extracted from the \&quot;tz\&quot; property, - If so, we perform the following actions:   - Walk through the whole JSON content to look for local datetimes,   - Use the timezone to transform each local datetime into an iso-8601 datetime.  Important: some objects contain a \&quot;properties\&quot; sub-object, which is a map of custom client data; the content of the \&quot;properties\&quot; sub-objects is always excluded from the time zone pre-processing.  | [optional] 

## Methods

### NewSMAInput

`func NewSMAInput(order SMAOrder, timeWindows []TimeWindow, ) *SMAInput`

NewSMAInput instantiates a new SMAInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMAInputWithDefaults

`func NewSMAInputWithDefaults() *SMAInput`

NewSMAInputWithDefaults instantiates a new SMAInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *SMAInput) GetOrder() SMAOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SMAInput) GetOrderOk() (*SMAOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SMAInput) SetOrder(v SMAOrder)`

SetOrder sets Order field to given value.


### GetMaxNbPropositions

`func (o *SMAInput) GetMaxNbPropositions() int32`

GetMaxNbPropositions returns the MaxNbPropositions field if non-nil, zero value otherwise.

### GetMaxNbPropositionsOk

`func (o *SMAInput) GetMaxNbPropositionsOk() (*int32, bool)`

GetMaxNbPropositionsOk returns a tuple with the MaxNbPropositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNbPropositions

`func (o *SMAInput) SetMaxNbPropositions(v int32)`

SetMaxNbPropositions sets MaxNbPropositions field to given value.

### HasMaxNbPropositions

`func (o *SMAInput) HasMaxNbPropositions() bool`

HasMaxNbPropositions returns a boolean if a field has been set.

### SetMaxNbPropositionsNil

`func (o *SMAInput) SetMaxNbPropositionsNil(b bool)`

 SetMaxNbPropositionsNil sets the value for MaxNbPropositions to be an explicit nil

### UnsetMaxNbPropositions
`func (o *SMAInput) UnsetMaxNbPropositions()`

UnsetMaxNbPropositions ensures that no value is present for MaxNbPropositions, not even an explicit nil
### GetTimeWindows

`func (o *SMAInput) GetTimeWindows() []TimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *SMAInput) GetTimeWindowsOk() (*[]TimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *SMAInput) SetTimeWindows(v []TimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetTz

`func (o *SMAInput) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *SMAInput) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *SMAInput) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *SMAInput) HasTz() bool`

HasTz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


