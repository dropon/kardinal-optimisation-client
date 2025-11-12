# EnvelopedLoginRefreshOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**EnvelopedLoginRefreshOutputItem**](EnvelopedLoginRefreshOutputItem.md) |  | [optional] 
**AccessToken** | Pointer to **string** | A JSON Web Token with scope &#39;access&#39;, to be used to access protected data, valid for 1 hour. | [optional] 

## Methods

### NewEnvelopedLoginRefreshOutput

`func NewEnvelopedLoginRefreshOutput() *EnvelopedLoginRefreshOutput`

NewEnvelopedLoginRefreshOutput instantiates a new EnvelopedLoginRefreshOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedLoginRefreshOutputWithDefaults

`func NewEnvelopedLoginRefreshOutputWithDefaults() *EnvelopedLoginRefreshOutput`

NewEnvelopedLoginRefreshOutputWithDefaults instantiates a new EnvelopedLoginRefreshOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EnvelopedLoginRefreshOutput) GetItem() EnvelopedLoginRefreshOutputItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EnvelopedLoginRefreshOutput) GetItemOk() (*EnvelopedLoginRefreshOutputItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EnvelopedLoginRefreshOutput) SetItem(v EnvelopedLoginRefreshOutputItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *EnvelopedLoginRefreshOutput) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetAccessToken

`func (o *EnvelopedLoginRefreshOutput) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *EnvelopedLoginRefreshOutput) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *EnvelopedLoginRefreshOutput) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *EnvelopedLoginRefreshOutput) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


