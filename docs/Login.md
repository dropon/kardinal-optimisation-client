# Login

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | [**Username**](Username.md) |  | 
**Password** | [**Password**](Password.md) | The user&#39;s password. | 

## Methods

### NewLogin

`func NewLogin(username Username, password Password, ) *Login`

NewLogin instantiates a new Login object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWithDefaults

`func NewLoginWithDefaults() *Login`

NewLoginWithDefaults instantiates a new Login object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *Login) GetUsername() Username`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Login) GetUsernameOk() (*Username, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Login) SetUsername(v Username)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *Login) GetPassword() Password`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Login) GetPasswordOk() (*Password, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Login) SetPassword(v Password)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


