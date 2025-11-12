# SolutionObjective

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Direction** | Pointer to [**OptimizationDirection**](OptimizationDirection.md) |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewSolutionObjective

`func NewSolutionObjective() *SolutionObjective`

NewSolutionObjective instantiates a new SolutionObjective object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolutionObjectiveWithDefaults

`func NewSolutionObjectiveWithDefaults() *SolutionObjective`

NewSolutionObjectiveWithDefaults instantiates a new SolutionObjective object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SolutionObjective) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SolutionObjective) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SolutionObjective) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SolutionObjective) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *SolutionObjective) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SolutionObjective) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SolutionObjective) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SolutionObjective) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDirection

`func (o *SolutionObjective) GetDirection() OptimizationDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SolutionObjective) GetDirectionOk() (*OptimizationDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SolutionObjective) SetDirection(v OptimizationDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SolutionObjective) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetValue

`func (o *SolutionObjective) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SolutionObjective) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SolutionObjective) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *SolutionObjective) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


