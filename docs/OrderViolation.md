# OrderViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ExpectedOrderedStopIds** | **[]string** | The stop ids in the expected order. | 
**MissedStopIds** | Pointer to **[]string** | The stop ids missing from the tour. | [optional] 
**BadlyOrderedStopIds** | Pointer to **[]string** | The stop ids badly ordered in the tour. | [optional] 

## Methods

### NewOrderViolation

`func NewOrderViolation(type_ string, expectedOrderedStopIds []string, ) *OrderViolation`

NewOrderViolation instantiates a new OrderViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderViolationWithDefaults

`func NewOrderViolationWithDefaults() *OrderViolation`

NewOrderViolationWithDefaults instantiates a new OrderViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderViolation) SetType(v string)`

SetType sets Type field to given value.


### GetExpectedOrderedStopIds

`func (o *OrderViolation) GetExpectedOrderedStopIds() []string`

GetExpectedOrderedStopIds returns the ExpectedOrderedStopIds field if non-nil, zero value otherwise.

### GetExpectedOrderedStopIdsOk

`func (o *OrderViolation) GetExpectedOrderedStopIdsOk() (*[]string, bool)`

GetExpectedOrderedStopIdsOk returns a tuple with the ExpectedOrderedStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOrderedStopIds

`func (o *OrderViolation) SetExpectedOrderedStopIds(v []string)`

SetExpectedOrderedStopIds sets ExpectedOrderedStopIds field to given value.


### GetMissedStopIds

`func (o *OrderViolation) GetMissedStopIds() []string`

GetMissedStopIds returns the MissedStopIds field if non-nil, zero value otherwise.

### GetMissedStopIdsOk

`func (o *OrderViolation) GetMissedStopIdsOk() (*[]string, bool)`

GetMissedStopIdsOk returns a tuple with the MissedStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissedStopIds

`func (o *OrderViolation) SetMissedStopIds(v []string)`

SetMissedStopIds sets MissedStopIds field to given value.

### HasMissedStopIds

`func (o *OrderViolation) HasMissedStopIds() bool`

HasMissedStopIds returns a boolean if a field has been set.

### GetBadlyOrderedStopIds

`func (o *OrderViolation) GetBadlyOrderedStopIds() []string`

GetBadlyOrderedStopIds returns the BadlyOrderedStopIds field if non-nil, zero value otherwise.

### GetBadlyOrderedStopIdsOk

`func (o *OrderViolation) GetBadlyOrderedStopIdsOk() (*[]string, bool)`

GetBadlyOrderedStopIdsOk returns a tuple with the BadlyOrderedStopIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadlyOrderedStopIds

`func (o *OrderViolation) SetBadlyOrderedStopIds(v []string)`

SetBadlyOrderedStopIds sets BadlyOrderedStopIds field to given value.

### HasBadlyOrderedStopIds

`func (o *OrderViolation) HasBadlyOrderedStopIds() bool`

HasBadlyOrderedStopIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


