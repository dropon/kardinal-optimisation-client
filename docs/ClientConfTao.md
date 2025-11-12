# ClientConfTao

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** |  | [optional] 
**CustomProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**Archiving** | Pointer to **map[string]string** | The delay before data is archived, for each concerned module. | [optional] 
**Retention** | Pointer to **map[string]string** | The delay before data is cleaned up, for each concerned module. | [optional] 
**RequiredProperties** | Pointer to [**TAOConfRequiredProperties**](TAOConfRequiredProperties.md) |  | [optional] 
**MaxSectorizationsInParallel** | Pointer to **float32** |  | [optional] [default to 5]

## Methods

### NewClientConfTao

`func NewClientConfTao() *ClientConfTao`

NewClientConfTao instantiates a new ClientConfTao object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfTaoWithDefaults

`func NewClientConfTaoWithDefaults() *ClientConfTao`

NewClientConfTaoWithDefaults instantiates a new ClientConfTao object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *ClientConfTao) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ClientConfTao) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ClientConfTao) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *ClientConfTao) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ClientConfTao) GetCustomProperties() map[string]interface{}`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ClientConfTao) GetCustomPropertiesOk() (*map[string]interface{}, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ClientConfTao) SetCustomProperties(v map[string]interface{})`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ClientConfTao) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetArchiving

`func (o *ClientConfTao) GetArchiving() map[string]string`

GetArchiving returns the Archiving field if non-nil, zero value otherwise.

### GetArchivingOk

`func (o *ClientConfTao) GetArchivingOk() (*map[string]string, bool)`

GetArchivingOk returns a tuple with the Archiving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiving

`func (o *ClientConfTao) SetArchiving(v map[string]string)`

SetArchiving sets Archiving field to given value.

### HasArchiving

`func (o *ClientConfTao) HasArchiving() bool`

HasArchiving returns a boolean if a field has been set.

### GetRetention

`func (o *ClientConfTao) GetRetention() map[string]string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *ClientConfTao) GetRetentionOk() (*map[string]string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *ClientConfTao) SetRetention(v map[string]string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *ClientConfTao) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetRequiredProperties

`func (o *ClientConfTao) GetRequiredProperties() TAOConfRequiredProperties`

GetRequiredProperties returns the RequiredProperties field if non-nil, zero value otherwise.

### GetRequiredPropertiesOk

`func (o *ClientConfTao) GetRequiredPropertiesOk() (*TAOConfRequiredProperties, bool)`

GetRequiredPropertiesOk returns a tuple with the RequiredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredProperties

`func (o *ClientConfTao) SetRequiredProperties(v TAOConfRequiredProperties)`

SetRequiredProperties sets RequiredProperties field to given value.

### HasRequiredProperties

`func (o *ClientConfTao) HasRequiredProperties() bool`

HasRequiredProperties returns a boolean if a field has been set.

### GetMaxSectorizationsInParallel

`func (o *ClientConfTao) GetMaxSectorizationsInParallel() float32`

GetMaxSectorizationsInParallel returns the MaxSectorizationsInParallel field if non-nil, zero value otherwise.

### GetMaxSectorizationsInParallelOk

`func (o *ClientConfTao) GetMaxSectorizationsInParallelOk() (*float32, bool)`

GetMaxSectorizationsInParallelOk returns a tuple with the MaxSectorizationsInParallel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSectorizationsInParallel

`func (o *ClientConfTao) SetMaxSectorizationsInParallel(v float32)`

SetMaxSectorizationsInParallel sets MaxSectorizationsInParallel field to given value.

### HasMaxSectorizationsInParallel

`func (o *ClientConfTao) HasMaxSectorizationsInParallel() bool`

HasMaxSectorizationsInParallel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


