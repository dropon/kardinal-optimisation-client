# MFAConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **string** | The config&#39;s last update date. | [optional] [readonly] 
**PreferredType** | Pointer to [**OTPType**](OTPType.md) |  | [optional] [default to OTPTYPE_NONE]
**AvailableTypes** | Pointer to [**[]OTPType**](OTPType.md) | The list of OTP types which have been configured by the actor in his MFA configuration. | [optional] [readonly] 
**Email** | Pointer to **NullableString** | An obfuscated representation of the email used for MFA. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | An obfuscated representation of the phone number used for MFA | [optional] 
**TotpConfig** | Pointer to [**NullableTOTPConfig**](TOTPConfig.md) |  | [optional] [readonly] 
**BackupCodes** | Pointer to **[]string** | A list of 10 unique 16-characters backup codes. | [optional] [readonly] 

## Methods

### NewMFAConfig

`func NewMFAConfig() *MFAConfig`

NewMFAConfig instantiates a new MFAConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAConfigWithDefaults

`func NewMFAConfigWithDefaults() *MFAConfig`

NewMFAConfigWithDefaults instantiates a new MFAConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *MFAConfig) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MFAConfig) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MFAConfig) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MFAConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPreferredType

`func (o *MFAConfig) GetPreferredType() OTPType`

GetPreferredType returns the PreferredType field if non-nil, zero value otherwise.

### GetPreferredTypeOk

`func (o *MFAConfig) GetPreferredTypeOk() (*OTPType, bool)`

GetPreferredTypeOk returns a tuple with the PreferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredType

`func (o *MFAConfig) SetPreferredType(v OTPType)`

SetPreferredType sets PreferredType field to given value.

### HasPreferredType

`func (o *MFAConfig) HasPreferredType() bool`

HasPreferredType returns a boolean if a field has been set.

### GetAvailableTypes

`func (o *MFAConfig) GetAvailableTypes() []OTPType`

GetAvailableTypes returns the AvailableTypes field if non-nil, zero value otherwise.

### GetAvailableTypesOk

`func (o *MFAConfig) GetAvailableTypesOk() (*[]OTPType, bool)`

GetAvailableTypesOk returns a tuple with the AvailableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTypes

`func (o *MFAConfig) SetAvailableTypes(v []OTPType)`

SetAvailableTypes sets AvailableTypes field to given value.

### HasAvailableTypes

`func (o *MFAConfig) HasAvailableTypes() bool`

HasAvailableTypes returns a boolean if a field has been set.

### GetEmail

`func (o *MFAConfig) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MFAConfig) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MFAConfig) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MFAConfig) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MFAConfig) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MFAConfig) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *MFAConfig) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MFAConfig) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MFAConfig) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MFAConfig) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *MFAConfig) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *MFAConfig) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetTotpConfig

`func (o *MFAConfig) GetTotpConfig() TOTPConfig`

GetTotpConfig returns the TotpConfig field if non-nil, zero value otherwise.

### GetTotpConfigOk

`func (o *MFAConfig) GetTotpConfigOk() (*TOTPConfig, bool)`

GetTotpConfigOk returns a tuple with the TotpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpConfig

`func (o *MFAConfig) SetTotpConfig(v TOTPConfig)`

SetTotpConfig sets TotpConfig field to given value.

### HasTotpConfig

`func (o *MFAConfig) HasTotpConfig() bool`

HasTotpConfig returns a boolean if a field has been set.

### SetTotpConfigNil

`func (o *MFAConfig) SetTotpConfigNil(b bool)`

 SetTotpConfigNil sets the value for TotpConfig to be an explicit nil

### UnsetTotpConfig
`func (o *MFAConfig) UnsetTotpConfig()`

UnsetTotpConfig ensures that no value is present for TotpConfig, not even an explicit nil
### GetBackupCodes

`func (o *MFAConfig) GetBackupCodes() []string`

GetBackupCodes returns the BackupCodes field if non-nil, zero value otherwise.

### GetBackupCodesOk

`func (o *MFAConfig) GetBackupCodesOk() (*[]string, bool)`

GetBackupCodesOk returns a tuple with the BackupCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCodes

`func (o *MFAConfig) SetBackupCodes(v []string)`

SetBackupCodes sets BackupCodes field to given value.

### HasBackupCodes

`func (o *MFAConfig) HasBackupCodes() bool`

HasBackupCodes returns a boolean if a field has been set.

### SetBackupCodesNil

`func (o *MFAConfig) SetBackupCodesNil(b bool)`

 SetBackupCodesNil sets the value for BackupCodes to be an explicit nil

### UnsetBackupCodes
`func (o *MFAConfig) UnsetBackupCodes()`

UnsetBackupCodes ensures that no value is present for BackupCodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


