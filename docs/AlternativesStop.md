# AlternativesStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Alternatives** | [**[]SingleStop**](SingleStop.md) |  | 

## Methods

### NewAlternativesStop

`func NewAlternativesStop(type_ string, alternatives []SingleStop, ) *AlternativesStop`

NewAlternativesStop instantiates a new AlternativesStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativesStopWithDefaults

`func NewAlternativesStopWithDefaults() *AlternativesStop`

NewAlternativesStopWithDefaults instantiates a new AlternativesStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AlternativesStop) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlternativesStop) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlternativesStop) SetType(v string)`

SetType sets Type field to given value.


### GetAlternatives

`func (o *AlternativesStop) GetAlternatives() []SingleStop`

GetAlternatives returns the Alternatives field if non-nil, zero value otherwise.

### GetAlternativesOk

`func (o *AlternativesStop) GetAlternativesOk() (*[]SingleStop, bool)`

GetAlternativesOk returns a tuple with the Alternatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatives

`func (o *AlternativesStop) SetAlternatives(v []SingleStop)`

SetAlternatives sets Alternatives field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


