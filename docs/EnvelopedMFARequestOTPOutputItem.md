# EnvelopedMFARequestOTPOutputItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpToken** | Pointer to **string** | A JSON Web Token with scope &#39;otp&#39;, to be used with an OTP (One-Time Password) value, valid for 10 minutes. | [optional] 

## Methods

### NewEnvelopedMFARequestOTPOutputItem

`func NewEnvelopedMFARequestOTPOutputItem() *EnvelopedMFARequestOTPOutputItem`

NewEnvelopedMFARequestOTPOutputItem instantiates a new EnvelopedMFARequestOTPOutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedMFARequestOTPOutputItemWithDefaults

`func NewEnvelopedMFARequestOTPOutputItemWithDefaults() *EnvelopedMFARequestOTPOutputItem`

NewEnvelopedMFARequestOTPOutputItemWithDefaults instantiates a new EnvelopedMFARequestOTPOutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpToken

`func (o *EnvelopedMFARequestOTPOutputItem) GetOtpToken() string`

GetOtpToken returns the OtpToken field if non-nil, zero value otherwise.

### GetOtpTokenOk

`func (o *EnvelopedMFARequestOTPOutputItem) GetOtpTokenOk() (*string, bool)`

GetOtpTokenOk returns a tuple with the OtpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpToken

`func (o *EnvelopedMFARequestOTPOutputItem) SetOtpToken(v string)`

SetOtpToken sets OtpToken field to given value.

### HasOtpToken

`func (o *EnvelopedMFARequestOTPOutputItem) HasOtpToken() bool`

HasOtpToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


