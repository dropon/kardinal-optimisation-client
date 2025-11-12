# StopIncompatibilityViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** | The name of the StopIncompatibility constraint the violation is linked to. | [optional] 
**Incompatibilities** | **[][]string** | Lists of stop ids for each incompatible tag. | 
**Tags** | **[]string** | A pair of incompatible stop tags. | 

## Methods

### NewStopIncompatibilityViolation

`func NewStopIncompatibilityViolation(type_ string, incompatibilities [][]string, tags []string, ) *StopIncompatibilityViolation`

NewStopIncompatibilityViolation instantiates a new StopIncompatibilityViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopIncompatibilityViolationWithDefaults

`func NewStopIncompatibilityViolationWithDefaults() *StopIncompatibilityViolation`

NewStopIncompatibilityViolationWithDefaults instantiates a new StopIncompatibilityViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StopIncompatibilityViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StopIncompatibilityViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StopIncompatibilityViolation) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *StopIncompatibilityViolation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StopIncompatibilityViolation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StopIncompatibilityViolation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StopIncompatibilityViolation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIncompatibilities

`func (o *StopIncompatibilityViolation) GetIncompatibilities() [][]string`

GetIncompatibilities returns the Incompatibilities field if non-nil, zero value otherwise.

### GetIncompatibilitiesOk

`func (o *StopIncompatibilityViolation) GetIncompatibilitiesOk() (*[][]string, bool)`

GetIncompatibilitiesOk returns a tuple with the Incompatibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibilities

`func (o *StopIncompatibilityViolation) SetIncompatibilities(v [][]string)`

SetIncompatibilities sets Incompatibilities field to given value.


### GetTags

`func (o *StopIncompatibilityViolation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StopIncompatibilityViolation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StopIncompatibilityViolation) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


