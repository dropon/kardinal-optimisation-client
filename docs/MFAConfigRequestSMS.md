# MFAConfigRequestSMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpType** | **string** |  | 
**PhoneNumber** | **string** | Phone number in E.164 international format (e.g., +33612345678). | 
**Password** | [**Password**](Password.md) |  | 

## Methods

### NewMFAConfigRequestSMS

`func NewMFAConfigRequestSMS(otpType string, phoneNumber string, password Password, ) *MFAConfigRequestSMS`

NewMFAConfigRequestSMS instantiates a new MFAConfigRequestSMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAConfigRequestSMSWithDefaults

`func NewMFAConfigRequestSMSWithDefaults() *MFAConfigRequestSMS`

NewMFAConfigRequestSMSWithDefaults instantiates a new MFAConfigRequestSMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpType

`func (o *MFAConfigRequestSMS) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *MFAConfigRequestSMS) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *MFAConfigRequestSMS) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.


### GetPhoneNumber

`func (o *MFAConfigRequestSMS) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MFAConfigRequestSMS) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MFAConfigRequestSMS) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetPassword

`func (o *MFAConfigRequestSMS) GetPassword() Password`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MFAConfigRequestSMS) GetPasswordOk() (*Password, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MFAConfigRequestSMS) SetPassword(v Password)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


