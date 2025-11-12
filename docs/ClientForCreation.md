# ClientForCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**interface{}**](AnyType.md) |  | 
**Name** | Pointer to **string** | The client name. | [optional] 
**Prefix** | [**interface{}**](AnyType.md) |  | 
**Conf** | Pointer to [**NullableClientConf**](ClientConf.md) |  | [optional] 

## Methods

### NewClientForCreation

`func NewClientForCreation(id interface{}, prefix interface{}, ) *ClientForCreation`

NewClientForCreation instantiates a new ClientForCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientForCreationWithDefaults

`func NewClientForCreationWithDefaults() *ClientForCreation`

NewClientForCreationWithDefaults instantiates a new ClientForCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientForCreation) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientForCreation) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientForCreation) SetId(v interface{})`

SetId sets Id field to given value.


### GetName

`func (o *ClientForCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientForCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientForCreation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientForCreation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *ClientForCreation) GetPrefix() interface{}`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ClientForCreation) GetPrefixOk() (*interface{}, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ClientForCreation) SetPrefix(v interface{})`

SetPrefix sets Prefix field to given value.


### GetConf

`func (o *ClientForCreation) GetConf() ClientConf`

GetConf returns the Conf field if non-nil, zero value otherwise.

### GetConfOk

`func (o *ClientForCreation) GetConfOk() (*ClientConf, bool)`

GetConfOk returns a tuple with the Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConf

`func (o *ClientForCreation) SetConf(v ClientConf)`

SetConf sets Conf field to given value.

### HasConf

`func (o *ClientForCreation) HasConf() bool`

HasConf returns a boolean if a field has been set.

### SetConfNil

`func (o *ClientForCreation) SetConfNil(b bool)`

 SetConfNil sets the value for Conf to be an explicit nil

### UnsetConf
`func (o *ClientForCreation) UnsetConf()`

UnsetConf ensures that no value is present for Conf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


