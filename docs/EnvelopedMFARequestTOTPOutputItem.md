# EnvelopedMFARequestTOTPOutputItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpToken** | Pointer to **string** | A JSON Web Token with scope &#39;otp&#39;, to be used with an OTP (One-Time Password) value, valid for 10 minutes. | [optional] 
**TotpConfig** | Pointer to [**NullableTOTPConfig**](TOTPConfig.md) |  | [optional] 

## Methods

### NewEnvelopedMFARequestTOTPOutputItem

`func NewEnvelopedMFARequestTOTPOutputItem() *EnvelopedMFARequestTOTPOutputItem`

NewEnvelopedMFARequestTOTPOutputItem instantiates a new EnvelopedMFARequestTOTPOutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedMFARequestTOTPOutputItemWithDefaults

`func NewEnvelopedMFARequestTOTPOutputItemWithDefaults() *EnvelopedMFARequestTOTPOutputItem`

NewEnvelopedMFARequestTOTPOutputItemWithDefaults instantiates a new EnvelopedMFARequestTOTPOutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpToken

`func (o *EnvelopedMFARequestTOTPOutputItem) GetOtpToken() string`

GetOtpToken returns the OtpToken field if non-nil, zero value otherwise.

### GetOtpTokenOk

`func (o *EnvelopedMFARequestTOTPOutputItem) GetOtpTokenOk() (*string, bool)`

GetOtpTokenOk returns a tuple with the OtpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpToken

`func (o *EnvelopedMFARequestTOTPOutputItem) SetOtpToken(v string)`

SetOtpToken sets OtpToken field to given value.

### HasOtpToken

`func (o *EnvelopedMFARequestTOTPOutputItem) HasOtpToken() bool`

HasOtpToken returns a boolean if a field has been set.

### GetTotpConfig

`func (o *EnvelopedMFARequestTOTPOutputItem) GetTotpConfig() TOTPConfig`

GetTotpConfig returns the TotpConfig field if non-nil, zero value otherwise.

### GetTotpConfigOk

`func (o *EnvelopedMFARequestTOTPOutputItem) GetTotpConfigOk() (*TOTPConfig, bool)`

GetTotpConfigOk returns a tuple with the TotpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpConfig

`func (o *EnvelopedMFARequestTOTPOutputItem) SetTotpConfig(v TOTPConfig)`

SetTotpConfig sets TotpConfig field to given value.

### HasTotpConfig

`func (o *EnvelopedMFARequestTOTPOutputItem) HasTotpConfig() bool`

HasTotpConfig returns a boolean if a field has been set.

### SetTotpConfigNil

`func (o *EnvelopedMFARequestTOTPOutputItem) SetTotpConfigNil(b bool)`

 SetTotpConfigNil sets the value for TotpConfig to be an explicit nil

### UnsetTotpConfig
`func (o *EnvelopedMFARequestTOTPOutputItem) UnsetTotpConfig()`

UnsetTotpConfig ensures that no value is present for TotpConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


