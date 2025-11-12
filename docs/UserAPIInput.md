# UserAPIInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | A valid email address. | 
**Company** | **string** | The client name to use for this simpleApi/complexApi user. | 
**Type** | Pointer to **string** | The input type describes the desired API usage: - simpleApi: will create a user with type \&quot;simpleApi\&quot; and access to a simplified ARO API, - complexApi: will create a user with type \&quot;standard\&quot; and access to the full ARO API.  | [optional] [default to "simpleApi"]
**Firstname** | Pointer to **string** | The user&#39;s first name | [optional] 
**Lastname** | Pointer to **string** | The user&#39;s last name | [optional] 

## Methods

### NewUserAPIInput

`func NewUserAPIInput(email string, company string, ) *UserAPIInput`

NewUserAPIInput instantiates a new UserAPIInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAPIInputWithDefaults

`func NewUserAPIInputWithDefaults() *UserAPIInput`

NewUserAPIInputWithDefaults instantiates a new UserAPIInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserAPIInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAPIInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAPIInput) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCompany

`func (o *UserAPIInput) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserAPIInput) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserAPIInput) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetType

`func (o *UserAPIInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAPIInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAPIInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserAPIInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFirstname

`func (o *UserAPIInput) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserAPIInput) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserAPIInput) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserAPIInput) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserAPIInput) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserAPIInput) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserAPIInput) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserAPIInput) HasLastname() bool`

HasLastname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


