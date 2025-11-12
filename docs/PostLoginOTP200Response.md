# PostLoginOTP200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**EnvelopedLoginAccessOutputItem**](EnvelopedLoginAccessOutputItem.md) |  | [optional] 
**AccessToken** | Pointer to **string** | A JSON Web Token with scope &#39;access&#39;, to be used to access protected data, valid for 1 hour. | [optional] 
**RefreshToken** | Pointer to **string** | A JSON Web Token with scope &#39;refresh&#39;, to be used to refresh the access token, valid for 30 days. | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewPostLoginOTP200Response

`func NewPostLoginOTP200Response() *PostLoginOTP200Response`

NewPostLoginOTP200Response instantiates a new PostLoginOTP200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLoginOTP200ResponseWithDefaults

`func NewPostLoginOTP200ResponseWithDefaults() *PostLoginOTP200Response`

NewPostLoginOTP200ResponseWithDefaults instantiates a new PostLoginOTP200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *PostLoginOTP200Response) GetItem() EnvelopedLoginAccessOutputItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *PostLoginOTP200Response) GetItemOk() (*EnvelopedLoginAccessOutputItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *PostLoginOTP200Response) SetItem(v EnvelopedLoginAccessOutputItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *PostLoginOTP200Response) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetAccessToken

`func (o *PostLoginOTP200Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *PostLoginOTP200Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *PostLoginOTP200Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *PostLoginOTP200Response) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *PostLoginOTP200Response) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *PostLoginOTP200Response) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *PostLoginOTP200Response) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *PostLoginOTP200Response) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetUser

`func (o *PostLoginOTP200Response) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PostLoginOTP200Response) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PostLoginOTP200Response) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *PostLoginOTP200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


