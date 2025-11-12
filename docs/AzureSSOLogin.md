# AzureSSOLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**IdToken** | **string** |  | 

## Methods

### NewAzureSSOLogin

`func NewAzureSSOLogin(accessToken string, idToken string, ) *AzureSSOLogin`

NewAzureSSOLogin instantiates a new AzureSSOLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSSOLoginWithDefaults

`func NewAzureSSOLoginWithDefaults() *AzureSSOLogin`

NewAzureSSOLoginWithDefaults instantiates a new AzureSSOLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AzureSSOLogin) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AzureSSOLogin) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AzureSSOLogin) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetIdToken

`func (o *AzureSSOLogin) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *AzureSSOLogin) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *AzureSSOLogin) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


