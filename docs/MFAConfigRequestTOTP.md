# MFAConfigRequestTOTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpType** | **string** |  | 
**Password** | [**Password**](Password.md) |  | 

## Methods

### NewMFAConfigRequestTOTP

`func NewMFAConfigRequestTOTP(otpType string, password Password, ) *MFAConfigRequestTOTP`

NewMFAConfigRequestTOTP instantiates a new MFAConfigRequestTOTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAConfigRequestTOTPWithDefaults

`func NewMFAConfigRequestTOTPWithDefaults() *MFAConfigRequestTOTP`

NewMFAConfigRequestTOTPWithDefaults instantiates a new MFAConfigRequestTOTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpType

`func (o *MFAConfigRequestTOTP) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *MFAConfigRequestTOTP) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *MFAConfigRequestTOTP) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.


### GetPassword

`func (o *MFAConfigRequestTOTP) GetPassword() Password`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MFAConfigRequestTOTP) GetPasswordOk() (*Password, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MFAConfigRequestTOTP) SetPassword(v Password)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


