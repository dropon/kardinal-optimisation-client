# RemovalStrategyViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** | The name of the RemovalStrategy constraint the violation is linked to. | [optional] 
**RemovalStrategyType** | [**RemovalStrategyType**](RemovalStrategyType.md) |  | 

## Methods

### NewRemovalStrategyViolation

`func NewRemovalStrategyViolation(type_ string, removalStrategyType RemovalStrategyType, ) *RemovalStrategyViolation`

NewRemovalStrategyViolation instantiates a new RemovalStrategyViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemovalStrategyViolationWithDefaults

`func NewRemovalStrategyViolationWithDefaults() *RemovalStrategyViolation`

NewRemovalStrategyViolationWithDefaults instantiates a new RemovalStrategyViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RemovalStrategyViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RemovalStrategyViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RemovalStrategyViolation) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RemovalStrategyViolation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemovalStrategyViolation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemovalStrategyViolation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemovalStrategyViolation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemovalStrategyType

`func (o *RemovalStrategyViolation) GetRemovalStrategyType() RemovalStrategyType`

GetRemovalStrategyType returns the RemovalStrategyType field if non-nil, zero value otherwise.

### GetRemovalStrategyTypeOk

`func (o *RemovalStrategyViolation) GetRemovalStrategyTypeOk() (*RemovalStrategyType, bool)`

GetRemovalStrategyTypeOk returns a tuple with the RemovalStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalStrategyType

`func (o *RemovalStrategyViolation) SetRemovalStrategyType(v RemovalStrategyType)`

SetRemovalStrategyType sets RemovalStrategyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


