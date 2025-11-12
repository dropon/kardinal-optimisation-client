# PutClientCountryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The country name. | [optional] 
**Id** | [**interface{}**](AnyType.md) |  | 

## Methods

### NewPutClientCountryRequest

`func NewPutClientCountryRequest(id interface{}, ) *PutClientCountryRequest`

NewPutClientCountryRequest instantiates a new PutClientCountryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutClientCountryRequestWithDefaults

`func NewPutClientCountryRequestWithDefaults() *PutClientCountryRequest`

NewPutClientCountryRequestWithDefaults instantiates a new PutClientCountryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutClientCountryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutClientCountryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutClientCountryRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutClientCountryRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *PutClientCountryRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutClientCountryRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutClientCountryRequest) SetId(v interface{})`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


