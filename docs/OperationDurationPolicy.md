# OperationDurationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** |  | 
**StopTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOperationDurationPolicy

`func NewOperationDurationPolicy(policy string, ) *OperationDurationPolicy`

NewOperationDurationPolicy instantiates a new OperationDurationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationDurationPolicyWithDefaults

`func NewOperationDurationPolicyWithDefaults() *OperationDurationPolicy`

NewOperationDurationPolicyWithDefaults instantiates a new OperationDurationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *OperationDurationPolicy) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *OperationDurationPolicy) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *OperationDurationPolicy) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetStopTags

`func (o *OperationDurationPolicy) GetStopTags() []string`

GetStopTags returns the StopTags field if non-nil, zero value otherwise.

### GetStopTagsOk

`func (o *OperationDurationPolicy) GetStopTagsOk() (*[]string, bool)`

GetStopTagsOk returns a tuple with the StopTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTags

`func (o *OperationDurationPolicy) SetStopTags(v []string)`

SetStopTags sets StopTags field to given value.

### HasStopTags

`func (o *OperationDurationPolicy) HasStopTags() bool`

HasStopTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


