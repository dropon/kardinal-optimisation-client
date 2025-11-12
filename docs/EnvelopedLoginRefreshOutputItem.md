# EnvelopedLoginRefreshOutputItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | A JSON Web Token with scope &#39;access&#39;, to be used to access protected data, valid for 1 hour. | [optional] 

## Methods

### NewEnvelopedLoginRefreshOutputItem

`func NewEnvelopedLoginRefreshOutputItem() *EnvelopedLoginRefreshOutputItem`

NewEnvelopedLoginRefreshOutputItem instantiates a new EnvelopedLoginRefreshOutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedLoginRefreshOutputItemWithDefaults

`func NewEnvelopedLoginRefreshOutputItemWithDefaults() *EnvelopedLoginRefreshOutputItem`

NewEnvelopedLoginRefreshOutputItemWithDefaults instantiates a new EnvelopedLoginRefreshOutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *EnvelopedLoginRefreshOutputItem) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *EnvelopedLoginRefreshOutputItem) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *EnvelopedLoginRefreshOutputItem) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *EnvelopedLoginRefreshOutputItem) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


