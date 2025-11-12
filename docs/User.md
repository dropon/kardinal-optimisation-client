# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Kardinal&#39;s unique identifier. | [optional] [readonly] 
**Username** | [**Username**](Username.md) |  | 
**Type** | Pointer to [**UserType**](UserType.md) |  | [optional] [default to USERTYPE_STANDARD]
**Password** | Pointer to [**Password**](Password.md) | The user&#39;s password. | [optional] 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] [default to ROLE_ANALYST]
**Firstname** | Pointer to **NullableString** | The user&#39;s first name. | [optional] 
**Lastname** | Pointer to **NullableString** | The user&#39;s last name. | [optional] 
**Level** | Pointer to [**Level**](Level.md) |  | [optional] [default to LEVEL_AGENCY]
**Countries** | Pointer to [**[]interface{}**](interface{}.md) | The countries whose data the user is allowed to see. | [optional] 
**Regions** | Pointer to [**[]interface{}**](interface{}.md) | The regions whose data the user is allowed to see. | [optional] 
**Agencies** | Pointer to [**[]interface{}**](interface{}.md) | The agencies whose data the user is allowed to see. | [optional] 
**Active** | Pointer to **bool** | Represents whether the user is enabled or disabled internally. | [optional] 
**Status** | Pointer to [**UserStatus**](UserStatus.md) |  | [optional] 
**LastLoginAt** | Pointer to **NullableString** | Represents the last login date of the user. | [optional] 
**CreatedAt** | Pointer to **string** | The user&#39;s creation date. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableString** | The user&#39;s last update date. | [optional] [readonly] 
**ClientId** | Pointer to [**interface{}**](AnyType.md) | The id of the client whom the user is linked to. It can be empty for a user which level is greater than &#39;owner&#39;.  When creating a new user, this id can be left empty if: - the caller has the level &#39;region&#39; or &#39;country&#39; or &#39;owner&#39;: in this case, the clientId of the new user is inherited from the clientId of the caller, - the new user has a level greater than &#39;owner&#39;.  | [optional] 
**GdprInfo** | Pointer to [**UserGdprInfo**](UserGdprInfo.md) |  | [optional] 
**ManagedClientsIds** | Pointer to [**[]interface{}**](interface{}.md) | For an &#39;admin&#39; user only: the ids of the clients which are managed by this user. | [optional] 

## Methods

### NewUser

`func NewUser(username Username, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() Username`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*Username, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v Username)`

SetUsername sets Username field to given value.


### GetType

`func (o *User) GetType() UserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User) GetTypeOk() (*UserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User) SetType(v UserType)`

SetType sets Type field to given value.

### HasType

`func (o *User) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPassword

`func (o *User) GetPassword() Password`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*Password, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v Password)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRole

`func (o *User) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *User) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *User) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *User) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetFirstname

`func (o *User) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *User) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *User) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *User) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### SetFirstnameNil

`func (o *User) SetFirstnameNil(b bool)`

 SetFirstnameNil sets the value for Firstname to be an explicit nil

### UnsetFirstname
`func (o *User) UnsetFirstname()`

UnsetFirstname ensures that no value is present for Firstname, not even an explicit nil
### GetLastname

`func (o *User) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *User) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *User) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *User) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### SetLastnameNil

`func (o *User) SetLastnameNil(b bool)`

 SetLastnameNil sets the value for Lastname to be an explicit nil

### UnsetLastname
`func (o *User) UnsetLastname()`

UnsetLastname ensures that no value is present for Lastname, not even an explicit nil
### GetLevel

`func (o *User) GetLevel() Level`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *User) GetLevelOk() (*Level, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *User) SetLevel(v Level)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *User) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetCountries

`func (o *User) GetCountries() []interface{}`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *User) GetCountriesOk() (*[]interface{}, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *User) SetCountries(v []interface{})`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *User) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetRegions

`func (o *User) GetRegions() []interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *User) GetRegionsOk() (*[]interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *User) SetRegions(v []interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *User) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetAgencies

`func (o *User) GetAgencies() []interface{}`

GetAgencies returns the Agencies field if non-nil, zero value otherwise.

### GetAgenciesOk

`func (o *User) GetAgenciesOk() (*[]interface{}, bool)`

GetAgenciesOk returns a tuple with the Agencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencies

`func (o *User) SetAgencies(v []interface{})`

SetAgencies sets Agencies field to given value.

### HasAgencies

`func (o *User) HasAgencies() bool`

HasAgencies returns a boolean if a field has been set.

### GetActive

`func (o *User) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *User) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *User) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *User) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStatus

`func (o *User) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *User) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *User) GetLastLoginAt() string`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *User) GetLastLoginAtOk() (*string, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *User) SetLastLoginAt(v string)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *User) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### SetLastLoginAtNil

`func (o *User) SetLastLoginAtNil(b bool)`

 SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil

### UnsetLastLoginAt
`func (o *User) UnsetLastLoginAt()`

UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
### GetCreatedAt

`func (o *User) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *User) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *User) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetClientId

`func (o *User) GetClientId() interface{}`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *User) GetClientIdOk() (*interface{}, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *User) SetClientId(v interface{})`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *User) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetGdprInfo

`func (o *User) GetGdprInfo() UserGdprInfo`

GetGdprInfo returns the GdprInfo field if non-nil, zero value otherwise.

### GetGdprInfoOk

`func (o *User) GetGdprInfoOk() (*UserGdprInfo, bool)`

GetGdprInfoOk returns a tuple with the GdprInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdprInfo

`func (o *User) SetGdprInfo(v UserGdprInfo)`

SetGdprInfo sets GdprInfo field to given value.

### HasGdprInfo

`func (o *User) HasGdprInfo() bool`

HasGdprInfo returns a boolean if a field has been set.

### GetManagedClientsIds

`func (o *User) GetManagedClientsIds() []interface{}`

GetManagedClientsIds returns the ManagedClientsIds field if non-nil, zero value otherwise.

### GetManagedClientsIdsOk

`func (o *User) GetManagedClientsIdsOk() (*[]interface{}, bool)`

GetManagedClientsIdsOk returns a tuple with the ManagedClientsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedClientsIds

`func (o *User) SetManagedClientsIds(v []interface{})`

SetManagedClientsIds sets ManagedClientsIds field to given value.

### HasManagedClientsIds

`func (o *User) HasManagedClientsIds() bool`

HasManagedClientsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


