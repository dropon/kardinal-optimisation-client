# AuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | The name of the authentication provider. | 
**Properties** | Pointer to **map[string]interface{}** | Additional properties needed by the authentication method. | [optional] 
**AuthUrl** | **string** | The absolute path of the login endpoint to use. | 

## Methods

### NewAuthenticationMethod

`func NewAuthenticationMethod(provider string, authUrl string, ) *AuthenticationMethod`

NewAuthenticationMethod instantiates a new AuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodWithDefaults

`func NewAuthenticationMethodWithDefaults() *AuthenticationMethod`

NewAuthenticationMethodWithDefaults instantiates a new AuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AuthenticationMethod) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthenticationMethod) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthenticationMethod) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProperties

`func (o *AuthenticationMethod) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AuthenticationMethod) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AuthenticationMethod) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AuthenticationMethod) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAuthUrl

`func (o *AuthenticationMethod) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *AuthenticationMethod) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *AuthenticationMethod) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


