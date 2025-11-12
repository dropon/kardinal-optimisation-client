# AtLeastOneConstraintViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** | The name of the AtLeastOneConstraint constraint the violation is linked to. | 
**ViolatedConstraints** | [**[]TourViolation**](TourViolation.md) | Lists of sub-constraints violations. | 

## Methods

### NewAtLeastOneConstraintViolation

`func NewAtLeastOneConstraintViolation(type_ string, name string, violatedConstraints []TourViolation, ) *AtLeastOneConstraintViolation`

NewAtLeastOneConstraintViolation instantiates a new AtLeastOneConstraintViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtLeastOneConstraintViolationWithDefaults

`func NewAtLeastOneConstraintViolationWithDefaults() *AtLeastOneConstraintViolation`

NewAtLeastOneConstraintViolationWithDefaults instantiates a new AtLeastOneConstraintViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AtLeastOneConstraintViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AtLeastOneConstraintViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AtLeastOneConstraintViolation) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AtLeastOneConstraintViolation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AtLeastOneConstraintViolation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AtLeastOneConstraintViolation) SetName(v string)`

SetName sets Name field to given value.


### GetViolatedConstraints

`func (o *AtLeastOneConstraintViolation) GetViolatedConstraints() []TourViolation`

GetViolatedConstraints returns the ViolatedConstraints field if non-nil, zero value otherwise.

### GetViolatedConstraintsOk

`func (o *AtLeastOneConstraintViolation) GetViolatedConstraintsOk() (*[]TourViolation, bool)`

GetViolatedConstraintsOk returns a tuple with the ViolatedConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatedConstraints

`func (o *AtLeastOneConstraintViolation) SetViolatedConstraints(v []TourViolation)`

SetViolatedConstraints sets ViolatedConstraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


