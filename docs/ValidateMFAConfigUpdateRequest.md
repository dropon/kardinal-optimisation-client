# ValidateMFAConfigUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpType** | **string** |  | 
**Email** | **string** | A valid email address. | 
**Otp** | **string** | A 6-digits OTP code. | 
**PhoneNumber** | **string** | Phone number in E.164 international format (e.g., +33612345678). | 

## Methods

### NewValidateMFAConfigUpdateRequest

`func NewValidateMFAConfigUpdateRequest(otpType string, email string, otp string, phoneNumber string, ) *ValidateMFAConfigUpdateRequest`

NewValidateMFAConfigUpdateRequest instantiates a new ValidateMFAConfigUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateMFAConfigUpdateRequestWithDefaults

`func NewValidateMFAConfigUpdateRequestWithDefaults() *ValidateMFAConfigUpdateRequest`

NewValidateMFAConfigUpdateRequestWithDefaults instantiates a new ValidateMFAConfigUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpType

`func (o *ValidateMFAConfigUpdateRequest) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *ValidateMFAConfigUpdateRequest) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *ValidateMFAConfigUpdateRequest) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.


### GetEmail

`func (o *ValidateMFAConfigUpdateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ValidateMFAConfigUpdateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ValidateMFAConfigUpdateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOtp

`func (o *ValidateMFAConfigUpdateRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *ValidateMFAConfigUpdateRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *ValidateMFAConfigUpdateRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.


### GetPhoneNumber

`func (o *ValidateMFAConfigUpdateRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ValidateMFAConfigUpdateRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ValidateMFAConfigUpdateRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


