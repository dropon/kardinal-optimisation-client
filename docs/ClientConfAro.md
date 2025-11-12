# ClientConfAro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** |  | [optional] 
**CustomProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**Archiving** | Pointer to **map[string]string** | The delay before data is archived, for each concerned module. | [optional] 
**Retention** | Pointer to **map[string]string** | The delay before data is cleaned up, for each concerned module. | [optional] 

## Methods

### NewClientConfAro

`func NewClientConfAro() *ClientConfAro`

NewClientConfAro instantiates a new ClientConfAro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfAroWithDefaults

`func NewClientConfAroWithDefaults() *ClientConfAro`

NewClientConfAroWithDefaults instantiates a new ClientConfAro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *ClientConfAro) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ClientConfAro) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ClientConfAro) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *ClientConfAro) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ClientConfAro) GetCustomProperties() map[string]interface{}`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ClientConfAro) GetCustomPropertiesOk() (*map[string]interface{}, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ClientConfAro) SetCustomProperties(v map[string]interface{})`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ClientConfAro) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetArchiving

`func (o *ClientConfAro) GetArchiving() map[string]string`

GetArchiving returns the Archiving field if non-nil, zero value otherwise.

### GetArchivingOk

`func (o *ClientConfAro) GetArchivingOk() (*map[string]string, bool)`

GetArchivingOk returns a tuple with the Archiving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiving

`func (o *ClientConfAro) SetArchiving(v map[string]string)`

SetArchiving sets Archiving field to given value.

### HasArchiving

`func (o *ClientConfAro) HasArchiving() bool`

HasArchiving returns a boolean if a field has been set.

### GetRetention

`func (o *ClientConfAro) GetRetention() map[string]string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *ClientConfAro) GetRetentionOk() (*map[string]string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *ClientConfAro) SetRetention(v map[string]string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *ClientConfAro) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


