# RequestNewMFAOTPCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpType** | **string** |  | 
**Email** | **string** | A valid email address. | 
**PhoneNumber** | **string** | Phone number in E.164 international format (e.g., +33612345678). | 

## Methods

### NewRequestNewMFAOTPCodeRequest

`func NewRequestNewMFAOTPCodeRequest(otpType string, email string, phoneNumber string, ) *RequestNewMFAOTPCodeRequest`

NewRequestNewMFAOTPCodeRequest instantiates a new RequestNewMFAOTPCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestNewMFAOTPCodeRequestWithDefaults

`func NewRequestNewMFAOTPCodeRequestWithDefaults() *RequestNewMFAOTPCodeRequest`

NewRequestNewMFAOTPCodeRequestWithDefaults instantiates a new RequestNewMFAOTPCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpType

`func (o *RequestNewMFAOTPCodeRequest) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *RequestNewMFAOTPCodeRequest) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *RequestNewMFAOTPCodeRequest) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.


### GetEmail

`func (o *RequestNewMFAOTPCodeRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RequestNewMFAOTPCodeRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RequestNewMFAOTPCodeRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhoneNumber

`func (o *RequestNewMFAOTPCodeRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RequestNewMFAOTPCodeRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RequestNewMFAOTPCodeRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


