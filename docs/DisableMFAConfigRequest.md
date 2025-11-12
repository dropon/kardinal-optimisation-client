# DisableMFAConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpType** | **string** |  | 
**Email** | **string** | A valid email address. | 
**Password** | [**Password**](Password.md) |  | 
**PhoneNumber** | **string** | Phone number in E.164 international format (e.g., +33612345678). | 

## Methods

### NewDisableMFAConfigRequest

`func NewDisableMFAConfigRequest(otpType string, email string, password Password, phoneNumber string, ) *DisableMFAConfigRequest`

NewDisableMFAConfigRequest instantiates a new DisableMFAConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableMFAConfigRequestWithDefaults

`func NewDisableMFAConfigRequestWithDefaults() *DisableMFAConfigRequest`

NewDisableMFAConfigRequestWithDefaults instantiates a new DisableMFAConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpType

`func (o *DisableMFAConfigRequest) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *DisableMFAConfigRequest) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *DisableMFAConfigRequest) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.


### GetEmail

`func (o *DisableMFAConfigRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DisableMFAConfigRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DisableMFAConfigRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *DisableMFAConfigRequest) GetPassword() Password`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DisableMFAConfigRequest) GetPasswordOk() (*Password, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DisableMFAConfigRequest) SetPassword(v Password)`

SetPassword sets Password field to given value.


### GetPhoneNumber

`func (o *DisableMFAConfigRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *DisableMFAConfigRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *DisableMFAConfigRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


