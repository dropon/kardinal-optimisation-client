# SingleStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "single"]
**Id** | **string** | Single stop ids must be unique within a plan. | 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Tags** | Pointer to **[]string** | prefix:suffix best practice, not forced. | [optional] 
**Position** | [**Position**](Position.md) |  | 
**Kind** | Pointer to [**StopKind**](StopKind.md) |  | [optional] [default to STOPKIND_DELIVERY]
**OperationDuration** | Pointer to **string** | A period of time, expressed in the ISO8601 **duration** format. | [optional] 
**Capacities** | Pointer to **map[string]float32** |  | [optional] 
**AuthorizedTimeWindows** | Pointer to [**[]TaggedTimeWindow**](TaggedTimeWindow.md) |  | [optional] 
**PreferredTimeWindows** | Pointer to [**[]TaggedTimeWindow**](TaggedTimeWindow.md) |  | [optional] 

## Methods

### NewSingleStop

`func NewSingleStop(id string, position Position, ) *SingleStop`

NewSingleStop instantiates a new SingleStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleStopWithDefaults

`func NewSingleStopWithDefaults() *SingleStop`

NewSingleStopWithDefaults instantiates a new SingleStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SingleStop) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SingleStop) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SingleStop) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SingleStop) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *SingleStop) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SingleStop) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SingleStop) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *SingleStop) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SingleStop) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SingleStop) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SingleStop) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *SingleStop) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SingleStop) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SingleStop) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SingleStop) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPosition

`func (o *SingleStop) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SingleStop) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SingleStop) SetPosition(v Position)`

SetPosition sets Position field to given value.


### GetKind

`func (o *SingleStop) GetKind() StopKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SingleStop) GetKindOk() (*StopKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SingleStop) SetKind(v StopKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SingleStop) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOperationDuration

`func (o *SingleStop) GetOperationDuration() string`

GetOperationDuration returns the OperationDuration field if non-nil, zero value otherwise.

### GetOperationDurationOk

`func (o *SingleStop) GetOperationDurationOk() (*string, bool)`

GetOperationDurationOk returns a tuple with the OperationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationDuration

`func (o *SingleStop) SetOperationDuration(v string)`

SetOperationDuration sets OperationDuration field to given value.

### HasOperationDuration

`func (o *SingleStop) HasOperationDuration() bool`

HasOperationDuration returns a boolean if a field has been set.

### GetCapacities

`func (o *SingleStop) GetCapacities() map[string]float32`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *SingleStop) GetCapacitiesOk() (*map[string]float32, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *SingleStop) SetCapacities(v map[string]float32)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *SingleStop) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetAuthorizedTimeWindows

`func (o *SingleStop) GetAuthorizedTimeWindows() []TaggedTimeWindow`

GetAuthorizedTimeWindows returns the AuthorizedTimeWindows field if non-nil, zero value otherwise.

### GetAuthorizedTimeWindowsOk

`func (o *SingleStop) GetAuthorizedTimeWindowsOk() (*[]TaggedTimeWindow, bool)`

GetAuthorizedTimeWindowsOk returns a tuple with the AuthorizedTimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedTimeWindows

`func (o *SingleStop) SetAuthorizedTimeWindows(v []TaggedTimeWindow)`

SetAuthorizedTimeWindows sets AuthorizedTimeWindows field to given value.

### HasAuthorizedTimeWindows

`func (o *SingleStop) HasAuthorizedTimeWindows() bool`

HasAuthorizedTimeWindows returns a boolean if a field has been set.

### GetPreferredTimeWindows

`func (o *SingleStop) GetPreferredTimeWindows() []TaggedTimeWindow`

GetPreferredTimeWindows returns the PreferredTimeWindows field if non-nil, zero value otherwise.

### GetPreferredTimeWindowsOk

`func (o *SingleStop) GetPreferredTimeWindowsOk() (*[]TaggedTimeWindow, bool)`

GetPreferredTimeWindowsOk returns a tuple with the PreferredTimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTimeWindows

`func (o *SingleStop) SetPreferredTimeWindows(v []TaggedTimeWindow)`

SetPreferredTimeWindows sets PreferredTimeWindows field to given value.

### HasPreferredTimeWindows

`func (o *SingleStop) HasPreferredTimeWindows() bool`

HasPreferredTimeWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


