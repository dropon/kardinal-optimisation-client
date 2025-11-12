# ProductConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** |  | [optional] 
**CustomProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**Archiving** | Pointer to **map[string]string** | The delay before data is archived, for each concerned module. | [optional] 
**Retention** | Pointer to **map[string]string** | The delay before data is cleaned up, for each concerned module. | [optional] 

## Methods

### NewProductConf

`func NewProductConf() *ProductConf`

NewProductConf instantiates a new ProductConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductConfWithDefaults

`func NewProductConfWithDefaults() *ProductConf`

NewProductConfWithDefaults instantiates a new ProductConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *ProductConf) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ProductConf) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ProductConf) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *ProductConf) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ProductConf) GetCustomProperties() map[string]interface{}`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ProductConf) GetCustomPropertiesOk() (*map[string]interface{}, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ProductConf) SetCustomProperties(v map[string]interface{})`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ProductConf) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetArchiving

`func (o *ProductConf) GetArchiving() map[string]string`

GetArchiving returns the Archiving field if non-nil, zero value otherwise.

### GetArchivingOk

`func (o *ProductConf) GetArchivingOk() (*map[string]string, bool)`

GetArchivingOk returns a tuple with the Archiving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiving

`func (o *ProductConf) SetArchiving(v map[string]string)`

SetArchiving sets Archiving field to given value.

### HasArchiving

`func (o *ProductConf) HasArchiving() bool`

HasArchiving returns a boolean if a field has been set.

### GetRetention

`func (o *ProductConf) GetRetention() map[string]string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *ProductConf) GetRetentionOk() (*map[string]string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *ProductConf) SetRetention(v map[string]string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *ProductConf) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


