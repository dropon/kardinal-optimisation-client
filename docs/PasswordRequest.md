# PasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**Password**](Password.md) | A new password. | [optional] 
**ConfirmPassword** | Pointer to [**Password**](Password.md) | The confirmation of the new password. | [optional] 

## Methods

### NewPasswordRequest

`func NewPasswordRequest() *PasswordRequest`

NewPasswordRequest instantiates a new PasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordRequestWithDefaults

`func NewPasswordRequestWithDefaults() *PasswordRequest`

NewPasswordRequestWithDefaults instantiates a new PasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PasswordRequest) GetPassword() Password`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordRequest) GetPasswordOk() (*Password, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordRequest) SetPassword(v Password)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PasswordRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetConfirmPassword

`func (o *PasswordRequest) GetConfirmPassword() Password`

GetConfirmPassword returns the ConfirmPassword field if non-nil, zero value otherwise.

### GetConfirmPasswordOk

`func (o *PasswordRequest) GetConfirmPasswordOk() (*Password, bool)`

GetConfirmPasswordOk returns a tuple with the ConfirmPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmPassword

`func (o *PasswordRequest) SetConfirmPassword(v Password)`

SetConfirmPassword sets ConfirmPassword field to given value.

### HasConfirmPassword

`func (o *PasswordRequest) HasConfirmPassword() bool`

HasConfirmPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


