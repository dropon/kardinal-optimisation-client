# AdditionalConstraintCapacities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Capacities** | **map[string]float32** |  | 
**ResourceTag** | **string** | At least one character among those allowed: unaccented alpha-numeric characters, \&quot;-\&quot;, \&quot;.\&quot;, \&quot;_\&quot;, \&quot;~\&quot;, \&quot;:\&quot;, \&quot;@\&quot;, \&quot;!\&quot;, \&quot;$\&quot;, \&quot;,\&quot;. | 

## Methods

### NewAdditionalConstraintCapacities

`func NewAdditionalConstraintCapacities(type_ string, capacities map[string]float32, resourceTag string, ) *AdditionalConstraintCapacities`

NewAdditionalConstraintCapacities instantiates a new AdditionalConstraintCapacities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalConstraintCapacitiesWithDefaults

`func NewAdditionalConstraintCapacitiesWithDefaults() *AdditionalConstraintCapacities`

NewAdditionalConstraintCapacitiesWithDefaults instantiates a new AdditionalConstraintCapacities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdditionalConstraintCapacities) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdditionalConstraintCapacities) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdditionalConstraintCapacities) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AdditionalConstraintCapacities) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdditionalConstraintCapacities) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdditionalConstraintCapacities) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdditionalConstraintCapacities) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCapacities

`func (o *AdditionalConstraintCapacities) GetCapacities() map[string]float32`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *AdditionalConstraintCapacities) GetCapacitiesOk() (*map[string]float32, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *AdditionalConstraintCapacities) SetCapacities(v map[string]float32)`

SetCapacities sets Capacities field to given value.


### GetResourceTag

`func (o *AdditionalConstraintCapacities) GetResourceTag() string`

GetResourceTag returns the ResourceTag field if non-nil, zero value otherwise.

### GetResourceTagOk

`func (o *AdditionalConstraintCapacities) GetResourceTagOk() (*string, bool)`

GetResourceTagOk returns a tuple with the ResourceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTag

`func (o *AdditionalConstraintCapacities) SetResourceTag(v string)`

SetResourceTag sets ResourceTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


