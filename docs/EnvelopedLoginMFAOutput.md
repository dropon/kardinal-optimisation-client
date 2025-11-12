# EnvelopedLoginMFAOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**EnvelopedLoginMFAOutputItem**](EnvelopedLoginMFAOutputItem.md) |  | [optional] 
**AccessToken** | Pointer to **string** | A JSON Web Token with scope &#39;mfa&#39;, to be used to setup MFA authentication, valid for 1 hour. | [optional] 

## Methods

### NewEnvelopedLoginMFAOutput

`func NewEnvelopedLoginMFAOutput() *EnvelopedLoginMFAOutput`

NewEnvelopedLoginMFAOutput instantiates a new EnvelopedLoginMFAOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedLoginMFAOutputWithDefaults

`func NewEnvelopedLoginMFAOutputWithDefaults() *EnvelopedLoginMFAOutput`

NewEnvelopedLoginMFAOutputWithDefaults instantiates a new EnvelopedLoginMFAOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EnvelopedLoginMFAOutput) GetItem() EnvelopedLoginMFAOutputItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EnvelopedLoginMFAOutput) GetItemOk() (*EnvelopedLoginMFAOutputItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EnvelopedLoginMFAOutput) SetItem(v EnvelopedLoginMFAOutputItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *EnvelopedLoginMFAOutput) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetAccessToken

`func (o *EnvelopedLoginMFAOutput) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *EnvelopedLoginMFAOutput) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *EnvelopedLoginMFAOutput) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *EnvelopedLoginMFAOutput) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


