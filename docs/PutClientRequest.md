# PutClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The client name. | [optional] 
**Conf** | Pointer to [**NullableClientConf**](ClientConf.md) |  | [optional] 
**Id** | [**interface{}**](AnyType.md) |  | 
**Prefix** | [**interface{}**](AnyType.md) |  | 

## Methods

### NewPutClientRequest

`func NewPutClientRequest(id interface{}, prefix interface{}, ) *PutClientRequest`

NewPutClientRequest instantiates a new PutClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutClientRequestWithDefaults

`func NewPutClientRequestWithDefaults() *PutClientRequest`

NewPutClientRequestWithDefaults instantiates a new PutClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutClientRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutClientRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConf

`func (o *PutClientRequest) GetConf() ClientConf`

GetConf returns the Conf field if non-nil, zero value otherwise.

### GetConfOk

`func (o *PutClientRequest) GetConfOk() (*ClientConf, bool)`

GetConfOk returns a tuple with the Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConf

`func (o *PutClientRequest) SetConf(v ClientConf)`

SetConf sets Conf field to given value.

### HasConf

`func (o *PutClientRequest) HasConf() bool`

HasConf returns a boolean if a field has been set.

### SetConfNil

`func (o *PutClientRequest) SetConfNil(b bool)`

 SetConfNil sets the value for Conf to be an explicit nil

### UnsetConf
`func (o *PutClientRequest) UnsetConf()`

UnsetConf ensures that no value is present for Conf, not even an explicit nil
### GetId

`func (o *PutClientRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutClientRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutClientRequest) SetId(v interface{})`

SetId sets Id field to given value.


### GetPrefix

`func (o *PutClientRequest) GetPrefix() interface{}`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PutClientRequest) GetPrefixOk() (*interface{}, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PutClientRequest) SetPrefix(v interface{})`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


