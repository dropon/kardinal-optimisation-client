# ClientForUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The client name. | [optional] 
**Conf** | Pointer to [**NullableClientConf**](ClientConf.md) |  | [optional] 

## Methods

### NewClientForUpdate

`func NewClientForUpdate() *ClientForUpdate`

NewClientForUpdate instantiates a new ClientForUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientForUpdateWithDefaults

`func NewClientForUpdateWithDefaults() *ClientForUpdate`

NewClientForUpdateWithDefaults instantiates a new ClientForUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientForUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientForUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientForUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientForUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConf

`func (o *ClientForUpdate) GetConf() ClientConf`

GetConf returns the Conf field if non-nil, zero value otherwise.

### GetConfOk

`func (o *ClientForUpdate) GetConfOk() (*ClientConf, bool)`

GetConfOk returns a tuple with the Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConf

`func (o *ClientForUpdate) SetConf(v ClientConf)`

SetConf sets Conf field to given value.

### HasConf

`func (o *ClientForUpdate) HasConf() bool`

HasConf returns a boolean if a field has been set.

### SetConfNil

`func (o *ClientForUpdate) SetConfNil(b bool)`

 SetConfNil sets the value for Conf to be an explicit nil

### UnsetConf
`func (o *ClientForUpdate) UnsetConf()`

UnsetConf ensures that no value is present for Conf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


