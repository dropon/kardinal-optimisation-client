# MFAConfigValidationEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpType** | **string** |  | 
**Email** | **string** | A valid email address. | 
**Otp** | **string** | A 6-digits OTP code. | 

## Methods

### NewMFAConfigValidationEmail

`func NewMFAConfigValidationEmail(otpType string, email string, otp string, ) *MFAConfigValidationEmail`

NewMFAConfigValidationEmail instantiates a new MFAConfigValidationEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAConfigValidationEmailWithDefaults

`func NewMFAConfigValidationEmailWithDefaults() *MFAConfigValidationEmail`

NewMFAConfigValidationEmailWithDefaults instantiates a new MFAConfigValidationEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpType

`func (o *MFAConfigValidationEmail) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *MFAConfigValidationEmail) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *MFAConfigValidationEmail) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.


### GetEmail

`func (o *MFAConfigValidationEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MFAConfigValidationEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MFAConfigValidationEmail) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOtp

`func (o *MFAConfigValidationEmail) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *MFAConfigValidationEmail) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *MFAConfigValidationEmail) SetOtp(v string)`

SetOtp sets Otp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


