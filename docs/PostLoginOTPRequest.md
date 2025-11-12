# PostLoginOTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Otp** | **string** | A 6-digits OTP code. | 
**BackupCode** | **string** | A unique 16-characters backup code. | 

## Methods

### NewPostLoginOTPRequest

`func NewPostLoginOTPRequest(otp string, backupCode string, ) *PostLoginOTPRequest`

NewPostLoginOTPRequest instantiates a new PostLoginOTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLoginOTPRequestWithDefaults

`func NewPostLoginOTPRequestWithDefaults() *PostLoginOTPRequest`

NewPostLoginOTPRequestWithDefaults instantiates a new PostLoginOTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtp

`func (o *PostLoginOTPRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *PostLoginOTPRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *PostLoginOTPRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.


### GetBackupCode

`func (o *PostLoginOTPRequest) GetBackupCode() string`

GetBackupCode returns the BackupCode field if non-nil, zero value otherwise.

### GetBackupCodeOk

`func (o *PostLoginOTPRequest) GetBackupCodeOk() (*string, bool)`

GetBackupCodeOk returns a tuple with the BackupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCode

`func (o *PostLoginOTPRequest) SetBackupCode(v string)`

SetBackupCode sets BackupCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


