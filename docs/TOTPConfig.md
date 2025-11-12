# TOTPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | A 32-characters secret code. | [optional] 
**Url** | Pointer to **string** | The TOTP auth protocol url. | [optional] 

## Methods

### NewTOTPConfig

`func NewTOTPConfig() *TOTPConfig`

NewTOTPConfig instantiates a new TOTPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTOTPConfigWithDefaults

`func NewTOTPConfigWithDefaults() *TOTPConfig`

NewTOTPConfigWithDefaults instantiates a new TOTPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *TOTPConfig) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TOTPConfig) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TOTPConfig) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TOTPConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUrl

`func (o *TOTPConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TOTPConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TOTPConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TOTPConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


