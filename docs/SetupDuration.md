# SetupDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromStopTag** | **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | 
**ToStopTag** | **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | 
**SetupDuration** | **string** | A period of time, expressed in the ISO8601 **duration** format. | 

## Methods

### NewSetupDuration

`func NewSetupDuration(fromStopTag string, toStopTag string, setupDuration string, ) *SetupDuration`

NewSetupDuration instantiates a new SetupDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupDurationWithDefaults

`func NewSetupDurationWithDefaults() *SetupDuration`

NewSetupDurationWithDefaults instantiates a new SetupDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromStopTag

`func (o *SetupDuration) GetFromStopTag() string`

GetFromStopTag returns the FromStopTag field if non-nil, zero value otherwise.

### GetFromStopTagOk

`func (o *SetupDuration) GetFromStopTagOk() (*string, bool)`

GetFromStopTagOk returns a tuple with the FromStopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromStopTag

`func (o *SetupDuration) SetFromStopTag(v string)`

SetFromStopTag sets FromStopTag field to given value.


### GetToStopTag

`func (o *SetupDuration) GetToStopTag() string`

GetToStopTag returns the ToStopTag field if non-nil, zero value otherwise.

### GetToStopTagOk

`func (o *SetupDuration) GetToStopTagOk() (*string, bool)`

GetToStopTagOk returns a tuple with the ToStopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToStopTag

`func (o *SetupDuration) SetToStopTag(v string)`

SetToStopTag sets ToStopTag field to given value.


### GetSetupDuration

`func (o *SetupDuration) GetSetupDuration() string`

GetSetupDuration returns the SetupDuration field if non-nil, zero value otherwise.

### GetSetupDurationOk

`func (o *SetupDuration) GetSetupDurationOk() (*string, bool)`

GetSetupDurationOk returns a tuple with the SetupDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupDuration

`func (o *SetupDuration) SetSetupDuration(v string)`

SetSetupDuration sets SetupDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


