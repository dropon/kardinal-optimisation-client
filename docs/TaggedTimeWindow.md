# TaggedTimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Begin** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 
**End** | **string** | A full calendar date time, expressed in the ISO8601 **date** format: YYYY-MM-DDThh:mm:ssZ. | 
**ResourceTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTaggedTimeWindow

`func NewTaggedTimeWindow(begin string, end string, ) *TaggedTimeWindow`

NewTaggedTimeWindow instantiates a new TaggedTimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggedTimeWindowWithDefaults

`func NewTaggedTimeWindowWithDefaults() *TaggedTimeWindow`

NewTaggedTimeWindowWithDefaults instantiates a new TaggedTimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBegin

`func (o *TaggedTimeWindow) GetBegin() string`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *TaggedTimeWindow) GetBeginOk() (*string, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *TaggedTimeWindow) SetBegin(v string)`

SetBegin sets Begin field to given value.


### GetEnd

`func (o *TaggedTimeWindow) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TaggedTimeWindow) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TaggedTimeWindow) SetEnd(v string)`

SetEnd sets End field to given value.


### GetResourceTags

`func (o *TaggedTimeWindow) GetResourceTags() []string`

GetResourceTags returns the ResourceTags field if non-nil, zero value otherwise.

### GetResourceTagsOk

`func (o *TaggedTimeWindow) GetResourceTagsOk() (*[]string, bool)`

GetResourceTagsOk returns a tuple with the ResourceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTags

`func (o *TaggedTimeWindow) SetResourceTags(v []string)`

SetResourceTags sets ResourceTags field to given value.

### HasResourceTags

`func (o *TaggedTimeWindow) HasResourceTags() bool`

HasResourceTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


