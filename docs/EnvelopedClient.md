# EnvelopedClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**Client**](Client.md) |  | [optional] 

## Methods

### NewEnvelopedClient

`func NewEnvelopedClient() *EnvelopedClient`

NewEnvelopedClient instantiates a new EnvelopedClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedClientWithDefaults

`func NewEnvelopedClientWithDefaults() *EnvelopedClient`

NewEnvelopedClientWithDefaults instantiates a new EnvelopedClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EnvelopedClient) GetItem() Client`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EnvelopedClient) GetItemOk() (*Client, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EnvelopedClient) SetItem(v Client)`

SetItem sets Item field to given value.

### HasItem

`func (o *EnvelopedClient) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


