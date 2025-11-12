# EnvelopedLoginGDPROutputItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GdprToken** | Pointer to **string** | A JSON Web Token with scope &#39;gdpr&#39;, to be used to approve a GDPR policy, valid for 1 hour. | [optional] 

## Methods

### NewEnvelopedLoginGDPROutputItem

`func NewEnvelopedLoginGDPROutputItem() *EnvelopedLoginGDPROutputItem`

NewEnvelopedLoginGDPROutputItem instantiates a new EnvelopedLoginGDPROutputItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedLoginGDPROutputItemWithDefaults

`func NewEnvelopedLoginGDPROutputItemWithDefaults() *EnvelopedLoginGDPROutputItem`

NewEnvelopedLoginGDPROutputItemWithDefaults instantiates a new EnvelopedLoginGDPROutputItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGdprToken

`func (o *EnvelopedLoginGDPROutputItem) GetGdprToken() string`

GetGdprToken returns the GdprToken field if non-nil, zero value otherwise.

### GetGdprTokenOk

`func (o *EnvelopedLoginGDPROutputItem) GetGdprTokenOk() (*string, bool)`

GetGdprTokenOk returns a tuple with the GdprToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdprToken

`func (o *EnvelopedLoginGDPROutputItem) SetGdprToken(v string)`

SetGdprToken sets GdprToken field to given value.

### HasGdprToken

`func (o *EnvelopedLoginGDPROutputItem) HasGdprToken() bool`

HasGdprToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


