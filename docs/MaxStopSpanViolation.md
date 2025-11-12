# MaxStopSpanViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**OrderId** | **string** | The order id with the MaxStopSpan constraint violation. | 
**Delay** | **string** | The delay that exceeds the MaxStopSpan. | 

## Methods

### NewMaxStopSpanViolation

`func NewMaxStopSpanViolation(type_ string, orderId string, delay string, ) *MaxStopSpanViolation`

NewMaxStopSpanViolation instantiates a new MaxStopSpanViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxStopSpanViolationWithDefaults

`func NewMaxStopSpanViolationWithDefaults() *MaxStopSpanViolation`

NewMaxStopSpanViolationWithDefaults instantiates a new MaxStopSpanViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MaxStopSpanViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MaxStopSpanViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MaxStopSpanViolation) SetType(v string)`

SetType sets Type field to given value.


### GetOrderId

`func (o *MaxStopSpanViolation) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MaxStopSpanViolation) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MaxStopSpanViolation) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetDelay

`func (o *MaxStopSpanViolation) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *MaxStopSpanViolation) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *MaxStopSpanViolation) SetDelay(v string)`

SetDelay sets Delay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


