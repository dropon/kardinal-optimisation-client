# MFAConfigRequestEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpType** | **string** |  | 
**Email** | **string** | A valid email address. | 
**Password** | [**Password**](Password.md) |  | 

## Methods

### NewMFAConfigRequestEmail

`func NewMFAConfigRequestEmail(otpType string, email string, password Password, ) *MFAConfigRequestEmail`

NewMFAConfigRequestEmail instantiates a new MFAConfigRequestEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAConfigRequestEmailWithDefaults

`func NewMFAConfigRequestEmailWithDefaults() *MFAConfigRequestEmail`

NewMFAConfigRequestEmailWithDefaults instantiates a new MFAConfigRequestEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpType

`func (o *MFAConfigRequestEmail) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *MFAConfigRequestEmail) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *MFAConfigRequestEmail) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.


### GetEmail

`func (o *MFAConfigRequestEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MFAConfigRequestEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MFAConfigRequestEmail) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *MFAConfigRequestEmail) GetPassword() Password`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MFAConfigRequestEmail) GetPasswordOk() (*Password, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MFAConfigRequestEmail) SetPassword(v Password)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


