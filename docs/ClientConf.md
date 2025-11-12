# ClientConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to [**ClientConfApi**](ClientConfApi.md) |  | [optional] 
**Aro** | Pointer to [**ClientConfAro**](ClientConfAro.md) |  | [optional] 
**Tao** | Pointer to [**ClientConfTao**](ClientConfTao.md) |  | [optional] 
**Pudo** | Pointer to [**ClientConfPudo**](ClientConfPudo.md) |  | [optional] 
**CustomProperties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewClientConf

`func NewClientConf() *ClientConf`

NewClientConf instantiates a new ClientConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfWithDefaults

`func NewClientConfWithDefaults() *ClientConf`

NewClientConfWithDefaults instantiates a new ClientConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *ClientConf) GetApi() ClientConfApi`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *ClientConf) GetApiOk() (*ClientConfApi, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *ClientConf) SetApi(v ClientConfApi)`

SetApi sets Api field to given value.

### HasApi

`func (o *ClientConf) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetAro

`func (o *ClientConf) GetAro() ClientConfAro`

GetAro returns the Aro field if non-nil, zero value otherwise.

### GetAroOk

`func (o *ClientConf) GetAroOk() (*ClientConfAro, bool)`

GetAroOk returns a tuple with the Aro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAro

`func (o *ClientConf) SetAro(v ClientConfAro)`

SetAro sets Aro field to given value.

### HasAro

`func (o *ClientConf) HasAro() bool`

HasAro returns a boolean if a field has been set.

### GetTao

`func (o *ClientConf) GetTao() ClientConfTao`

GetTao returns the Tao field if non-nil, zero value otherwise.

### GetTaoOk

`func (o *ClientConf) GetTaoOk() (*ClientConfTao, bool)`

GetTaoOk returns a tuple with the Tao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTao

`func (o *ClientConf) SetTao(v ClientConfTao)`

SetTao sets Tao field to given value.

### HasTao

`func (o *ClientConf) HasTao() bool`

HasTao returns a boolean if a field has been set.

### GetPudo

`func (o *ClientConf) GetPudo() ClientConfPudo`

GetPudo returns the Pudo field if non-nil, zero value otherwise.

### GetPudoOk

`func (o *ClientConf) GetPudoOk() (*ClientConfPudo, bool)`

GetPudoOk returns a tuple with the Pudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPudo

`func (o *ClientConf) SetPudo(v ClientConfPudo)`

SetPudo sets Pudo field to given value.

### HasPudo

`func (o *ClientConf) HasPudo() bool`

HasPudo returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ClientConf) GetCustomProperties() map[string]interface{}`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ClientConf) GetCustomPropertiesOk() (*map[string]interface{}, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ClientConf) SetCustomProperties(v map[string]interface{})`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ClientConf) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


