# PutClientAgencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The agency name. | [optional] 
**Id** | [**interface{}**](AnyType.md) |  | 

## Methods

### NewPutClientAgencyRequest

`func NewPutClientAgencyRequest(id interface{}, ) *PutClientAgencyRequest`

NewPutClientAgencyRequest instantiates a new PutClientAgencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutClientAgencyRequestWithDefaults

`func NewPutClientAgencyRequestWithDefaults() *PutClientAgencyRequest`

NewPutClientAgencyRequestWithDefaults instantiates a new PutClientAgencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutClientAgencyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutClientAgencyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutClientAgencyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutClientAgencyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *PutClientAgencyRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutClientAgencyRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutClientAgencyRequest) SetId(v interface{})`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


