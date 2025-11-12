# ClientConfPudo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** |  | [optional] 
**CustomProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**Archiving** | Pointer to **map[string]string** | The delay before data is archived, for each concerned module. | [optional] 
**Retention** | Pointer to **map[string]string** | The delay before data is cleaned up, for each concerned module. | [optional] 

## Methods

### NewClientConfPudo

`func NewClientConfPudo() *ClientConfPudo`

NewClientConfPudo instantiates a new ClientConfPudo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfPudoWithDefaults

`func NewClientConfPudoWithDefaults() *ClientConfPudo`

NewClientConfPudoWithDefaults instantiates a new ClientConfPudo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *ClientConfPudo) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ClientConfPudo) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ClientConfPudo) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *ClientConfPudo) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ClientConfPudo) GetCustomProperties() map[string]interface{}`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ClientConfPudo) GetCustomPropertiesOk() (*map[string]interface{}, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ClientConfPudo) SetCustomProperties(v map[string]interface{})`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ClientConfPudo) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetArchiving

`func (o *ClientConfPudo) GetArchiving() map[string]string`

GetArchiving returns the Archiving field if non-nil, zero value otherwise.

### GetArchivingOk

`func (o *ClientConfPudo) GetArchivingOk() (*map[string]string, bool)`

GetArchivingOk returns a tuple with the Archiving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiving

`func (o *ClientConfPudo) SetArchiving(v map[string]string)`

SetArchiving sets Archiving field to given value.

### HasArchiving

`func (o *ClientConfPudo) HasArchiving() bool`

HasArchiving returns a boolean if a field has been set.

### GetRetention

`func (o *ClientConfPudo) GetRetention() map[string]string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *ClientConfPudo) GetRetentionOk() (*map[string]string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *ClientConfPudo) SetRetention(v map[string]string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *ClientConfPudo) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


