# MFAConfigValidationSMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpType** | **string** |  | 
**PhoneNumber** | **string** | Phone number in E.164 international format (e.g., +33612345678). | 
**Otp** | **string** | A 6-digits OTP code. | 

## Methods

### NewMFAConfigValidationSMS

`func NewMFAConfigValidationSMS(otpType string, phoneNumber string, otp string, ) *MFAConfigValidationSMS`

NewMFAConfigValidationSMS instantiates a new MFAConfigValidationSMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAConfigValidationSMSWithDefaults

`func NewMFAConfigValidationSMSWithDefaults() *MFAConfigValidationSMS`

NewMFAConfigValidationSMSWithDefaults instantiates a new MFAConfigValidationSMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpType

`func (o *MFAConfigValidationSMS) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *MFAConfigValidationSMS) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *MFAConfigValidationSMS) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.


### GetPhoneNumber

`func (o *MFAConfigValidationSMS) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MFAConfigValidationSMS) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MFAConfigValidationSMS) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetOtp

`func (o *MFAConfigValidationSMS) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *MFAConfigValidationSMS) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *MFAConfigValidationSMS) SetOtp(v string)`

SetOtp sets Otp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


