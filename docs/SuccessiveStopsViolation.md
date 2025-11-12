# SuccessiveStopsViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ExpectedSuccessiveStopIds** | **[]string** | The stop ids in the expected order. | 
**InterleavedStopIds** | Pointer to **[]string** | The stop ids which are not expected and interleaved between the order stops. | [optional] 

## Methods

### NewSuccessiveStopsViolation

`func NewSuccessiveStopsViolation(type_ string, expectedSuccessiveStopIds []string, ) *SuccessiveStopsViolation`

NewSuccessiveStopsViolation instantiates a new SuccessiveStopsViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessiveStopsViolationWithDefaults

`func NewSuccessiveStopsViolationWithDefaults() *SuccessiveStopsViolation`

NewSuccessiveStopsViolationWithDefaults instantiates a new SuccessiveStopsViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SuccessiveStopsViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SuccessiveStopsViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SuccessiveStopsViolation) SetType(v string)`

SetType sets Type field to given value.


### GetExpectedSuccessiveStopIds

`func (o *SuccessiveStopsViolation) GetExpectedSuccessiveStopIds() []string`

GetExpectedSuccessiveStopIds returns the ExpectedSuccessiveStopIds field if non-nil, zero value otherwise.

### GetExpectedSuccessiveStopIdsOk

`func (o *SuccessiveStopsViolation) GetExpectedSuccessiveStopIdsOk() (*[]string, bool)`

GetExpectedSuccessiveStopIdsOk returns a tuple with the ExpectedSuccessiveStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSuccessiveStopIds

`func (o *SuccessiveStopsViolation) SetExpectedSuccessiveStopIds(v []string)`

SetExpectedSuccessiveStopIds sets ExpectedSuccessiveStopIds field to given value.


### GetInterleavedStopIds

`func (o *SuccessiveStopsViolation) GetInterleavedStopIds() []string`

GetInterleavedStopIds returns the InterleavedStopIds field if non-nil, zero value otherwise.

### GetInterleavedStopIdsOk

`func (o *SuccessiveStopsViolation) GetInterleavedStopIdsOk() (*[]string, bool)`

GetInterleavedStopIdsOk returns a tuple with the InterleavedStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterleavedStopIds

`func (o *SuccessiveStopsViolation) SetInterleavedStopIds(v []string)`

SetInterleavedStopIds sets InterleavedStopIds field to given value.

### HasInterleavedStopIds

`func (o *SuccessiveStopsViolation) HasInterleavedStopIds() bool`

HasInterleavedStopIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


