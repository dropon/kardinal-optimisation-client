# MFAConfigResendEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpType** | **string** |  | 
**Email** | **string** | A valid email address. | 

## Methods

### NewMFAConfigResendEmail

`func NewMFAConfigResendEmail(otpType string, email string, ) *MFAConfigResendEmail`

NewMFAConfigResendEmail instantiates a new MFAConfigResendEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAConfigResendEmailWithDefaults

`func NewMFAConfigResendEmailWithDefaults() *MFAConfigResendEmail`

NewMFAConfigResendEmailWithDefaults instantiates a new MFAConfigResendEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpType

`func (o *MFAConfigResendEmail) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *MFAConfigResendEmail) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *MFAConfigResendEmail) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.


### GetEmail

`func (o *MFAConfigResendEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MFAConfigResendEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MFAConfigResendEmail) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


