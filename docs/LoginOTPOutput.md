# LoginOTPOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpToken** | Pointer to **string** | A JWT token with scope &#39;otp&#39; valid for 10 minutes. | [optional] 
**PreferredType** | Pointer to [**OTPType**](OTPType.md) |  | [optional] [default to OTPTYPE_NONE]
**AvailableTypes** | Pointer to [**[]OTPType**](OTPType.md) | The list of OTP types which have been configured by the actor in his MFA configuration | [optional] [readonly] 
**Email** | Pointer to **NullableString** | An obfuscated representation of the email used for MFA. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | An obfuscated representation of the phone number used for MFA | [optional] 

## Methods

### NewLoginOTPOutput

`func NewLoginOTPOutput() *LoginOTPOutput`

NewLoginOTPOutput instantiates a new LoginOTPOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginOTPOutputWithDefaults

`func NewLoginOTPOutputWithDefaults() *LoginOTPOutput`

NewLoginOTPOutputWithDefaults instantiates a new LoginOTPOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpToken

`func (o *LoginOTPOutput) GetOtpToken() string`

GetOtpToken returns the OtpToken field if non-nil, zero value otherwise.

### GetOtpTokenOk

`func (o *LoginOTPOutput) GetOtpTokenOk() (*string, bool)`

GetOtpTokenOk returns a tuple with the OtpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpToken

`func (o *LoginOTPOutput) SetOtpToken(v string)`

SetOtpToken sets OtpToken field to given value.

### HasOtpToken

`func (o *LoginOTPOutput) HasOtpToken() bool`

HasOtpToken returns a boolean if a field has been set.

### GetPreferredType

`func (o *LoginOTPOutput) GetPreferredType() OTPType`

GetPreferredType returns the PreferredType field if non-nil, zero value otherwise.

### GetPreferredTypeOk

`func (o *LoginOTPOutput) GetPreferredTypeOk() (*OTPType, bool)`

GetPreferredTypeOk returns a tuple with the PreferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredType

`func (o *LoginOTPOutput) SetPreferredType(v OTPType)`

SetPreferredType sets PreferredType field to given value.

### HasPreferredType

`func (o *LoginOTPOutput) HasPreferredType() bool`

HasPreferredType returns a boolean if a field has been set.

### GetAvailableTypes

`func (o *LoginOTPOutput) GetAvailableTypes() []OTPType`

GetAvailableTypes returns the AvailableTypes field if non-nil, zero value otherwise.

### GetAvailableTypesOk

`func (o *LoginOTPOutput) GetAvailableTypesOk() (*[]OTPType, bool)`

GetAvailableTypesOk returns a tuple with the AvailableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTypes

`func (o *LoginOTPOutput) SetAvailableTypes(v []OTPType)`

SetAvailableTypes sets AvailableTypes field to given value.

### HasAvailableTypes

`func (o *LoginOTPOutput) HasAvailableTypes() bool`

HasAvailableTypes returns a boolean if a field has been set.

### GetEmail

`func (o *LoginOTPOutput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoginOTPOutput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoginOTPOutput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LoginOTPOutput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LoginOTPOutput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LoginOTPOutput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *LoginOTPOutput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *LoginOTPOutput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *LoginOTPOutput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *LoginOTPOutput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *LoginOTPOutput) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *LoginOTPOutput) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


