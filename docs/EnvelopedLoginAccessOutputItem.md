# EnvelopedLoginAccessOutputItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | A JSON Web Token with scope &#39;access&#39;, to be used to access protected data, valid for 1 hour. | [optional] 
**RefreshToken** | Pointer to **string** | A JSON Web Token with scope &#39;refresh&#39;, to be used to refresh the access token, valid for 30 days. | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewEnvelopedLoginAccessOutputItem

`func NewEnvelopedLoginAccessOutputItem() *EnvelopedLoginAccessOutputItem`

NewEnvelopedLoginAccessOutputItem instantiates a new EnvelopedLoginAccessOutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedLoginAccessOutputItemWithDefaults

`func NewEnvelopedLoginAccessOutputItemWithDefaults() *EnvelopedLoginAccessOutputItem`

NewEnvelopedLoginAccessOutputItemWithDefaults instantiates a new EnvelopedLoginAccessOutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *EnvelopedLoginAccessOutputItem) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *EnvelopedLoginAccessOutputItem) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *EnvelopedLoginAccessOutputItem) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *EnvelopedLoginAccessOutputItem) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *EnvelopedLoginAccessOutputItem) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *EnvelopedLoginAccessOutputItem) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *EnvelopedLoginAccessOutputItem) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *EnvelopedLoginAccessOutputItem) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetUser

`func (o *EnvelopedLoginAccessOutputItem) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EnvelopedLoginAccessOutputItem) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EnvelopedLoginAccessOutputItem) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *EnvelopedLoginAccessOutputItem) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


