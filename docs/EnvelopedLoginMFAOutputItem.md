# EnvelopedLoginMFAOutputItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaToken** | Pointer to **string** | A JSON Web Token with scope &#39;mfa&#39;, to be used to setup MFA authentication, valid for 1 hour. | [optional] 

## Methods

### NewEnvelopedLoginMFAOutputItem

`func NewEnvelopedLoginMFAOutputItem() *EnvelopedLoginMFAOutputItem`

NewEnvelopedLoginMFAOutputItem instantiates a new EnvelopedLoginMFAOutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedLoginMFAOutputItemWithDefaults

`func NewEnvelopedLoginMFAOutputItemWithDefaults() *EnvelopedLoginMFAOutputItem`

NewEnvelopedLoginMFAOutputItemWithDefaults instantiates a new EnvelopedLoginMFAOutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfaToken

`func (o *EnvelopedLoginMFAOutputItem) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *EnvelopedLoginMFAOutputItem) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *EnvelopedLoginMFAOutputItem) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.

### HasMfaToken

`func (o *EnvelopedLoginMFAOutputItem) HasMfaToken() bool`

HasMfaToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


