# UserAPIOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User**](User.md) |  | [optional] 
**AccessTokens** | Pointer to [**UserAPIOutputAccessTokens**](UserAPIOutputAccessTokens.md) |  | [optional] 

## Methods

### NewUserAPIOutput

`func NewUserAPIOutput() *UserAPIOutput`

NewUserAPIOutput instantiates a new UserAPIOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAPIOutputWithDefaults

`func NewUserAPIOutputWithDefaults() *UserAPIOutput`

NewUserAPIOutputWithDefaults instantiates a new UserAPIOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserAPIOutput) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserAPIOutput) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserAPIOutput) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *UserAPIOutput) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAccessTokens

`func (o *UserAPIOutput) GetAccessTokens() UserAPIOutputAccessTokens`

GetAccessTokens returns the AccessTokens field if non-nil, zero value otherwise.

### GetAccessTokensOk

`func (o *UserAPIOutput) GetAccessTokensOk() (*UserAPIOutputAccessTokens, bool)`

GetAccessTokensOk returns a tuple with the AccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokens

`func (o *UserAPIOutput) SetAccessTokens(v UserAPIOutputAccessTokens)`

SetAccessTokens sets AccessTokens field to given value.

### HasAccessTokens

`func (o *UserAPIOutput) HasAccessTokens() bool`

HasAccessTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


