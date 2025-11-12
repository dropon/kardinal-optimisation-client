# RequestMFAConfigUpdate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**EnvelopedMFARequestTOTPOutputItem**](EnvelopedMFARequestTOTPOutputItem.md) |  | [optional] 

## Methods

### NewRequestMFAConfigUpdate200Response

`func NewRequestMFAConfigUpdate200Response() *RequestMFAConfigUpdate200Response`

NewRequestMFAConfigUpdate200Response instantiates a new RequestMFAConfigUpdate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMFAConfigUpdate200ResponseWithDefaults

`func NewRequestMFAConfigUpdate200ResponseWithDefaults() *RequestMFAConfigUpdate200Response`

NewRequestMFAConfigUpdate200ResponseWithDefaults instantiates a new RequestMFAConfigUpdate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *RequestMFAConfigUpdate200Response) GetItem() EnvelopedMFARequestTOTPOutputItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *RequestMFAConfigUpdate200Response) GetItemOk() (*EnvelopedMFARequestTOTPOutputItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *RequestMFAConfigUpdate200Response) SetItem(v EnvelopedMFARequestTOTPOutputItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *RequestMFAConfigUpdate200Response) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


