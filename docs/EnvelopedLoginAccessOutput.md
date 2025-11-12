# EnvelopedLoginAccessOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**EnvelopedLoginAccessOutputItem**](EnvelopedLoginAccessOutputItem.md) |  | [optional] 
**AccessToken** | Pointer to **string** | A JSON Web Token with scope &#39;access&#39;, to be used to access protected data, valid for 1 hour. | [optional] 
**RefreshToken** | Pointer to **string** | A JSON Web Token with scope &#39;refresh&#39;, to be used to refresh the access token, valid for 30 days. | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewEnvelopedLoginAccessOutput

`func NewEnvelopedLoginAccessOutput() *EnvelopedLoginAccessOutput`

NewEnvelopedLoginAccessOutput instantiates a new EnvelopedLoginAccessOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedLoginAccessOutputWithDefaults

`func NewEnvelopedLoginAccessOutputWithDefaults() *EnvelopedLoginAccessOutput`

NewEnvelopedLoginAccessOutputWithDefaults instantiates a new EnvelopedLoginAccessOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EnvelopedLoginAccessOutput) GetItem() EnvelopedLoginAccessOutputItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EnvelopedLoginAccessOutput) GetItemOk() (*EnvelopedLoginAccessOutputItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EnvelopedLoginAccessOutput) SetItem(v EnvelopedLoginAccessOutputItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *EnvelopedLoginAccessOutput) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetAccessToken

`func (o *EnvelopedLoginAccessOutput) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *EnvelopedLoginAccessOutput) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *EnvelopedLoginAccessOutput) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *EnvelopedLoginAccessOutput) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *EnvelopedLoginAccessOutput) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *EnvelopedLoginAccessOutput) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *EnvelopedLoginAccessOutput) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *EnvelopedLoginAccessOutput) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetUser

`func (o *EnvelopedLoginAccessOutput) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EnvelopedLoginAccessOutput) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EnvelopedLoginAccessOutput) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *EnvelopedLoginAccessOutput) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


