# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**interface{}**](AnyType.md) |  | 
**CreatedAt** | Pointer to **string** | The client&#39;s creation date. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableString** | The client&#39;s last update date. | [optional] [readonly] 
**Name** | Pointer to **string** | The client name. | [optional] 
**Prefix** | [**interface{}**](AnyType.md) |  | 
**OwnerIds** | Pointer to **[]string** | The identifier(s) of the user(s) who can manage the client. | [optional] 
**Countries** | Pointer to [**[]Country**](Country.md) |  | [optional] 
**Conf** | Pointer to [**NullableClientConf**](ClientConf.md) |  | [optional] 

## Methods

### NewClient

`func NewClient(id interface{}, prefix interface{}, ) *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Client) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Client) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Client) SetId(v interface{})`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Client) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Client) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Client) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Client) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Client) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Client) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Client) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Client) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Client) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Client) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetName

`func (o *Client) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Client) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Client) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Client) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *Client) GetPrefix() interface{}`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *Client) GetPrefixOk() (*interface{}, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *Client) SetPrefix(v interface{})`

SetPrefix sets Prefix field to given value.


### GetOwnerIds

`func (o *Client) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *Client) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *Client) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *Client) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetCountries

`func (o *Client) GetCountries() []Country`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *Client) GetCountriesOk() (*[]Country, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *Client) SetCountries(v []Country)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *Client) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetConf

`func (o *Client) GetConf() ClientConf`

GetConf returns the Conf field if non-nil, zero value otherwise.

### GetConfOk

`func (o *Client) GetConfOk() (*ClientConf, bool)`

GetConfOk returns a tuple with the Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConf

`func (o *Client) SetConf(v ClientConf)`

SetConf sets Conf field to given value.

### HasConf

`func (o *Client) HasConf() bool`

HasConf returns a boolean if a field has been set.

### SetConfNil

`func (o *Client) SetConfNil(b bool)`

 SetConfNil sets the value for Conf to be an explicit nil

### UnsetConf
`func (o *Client) UnsetConf()`

UnsetConf ensures that no value is present for Conf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


