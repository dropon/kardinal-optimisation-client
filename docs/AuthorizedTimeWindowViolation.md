# AuthorizedTimeWindowViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**WayPointIndex** | Pointer to **int32** | The idx of the waypoint the violation applies to. | [optional] 
**ForbiddenWindow** | [**TimeWindow**](TimeWindow.md) | The violated forbidden window. | 
**ActualTime** | **string** | The time at which the waypoint is actually executed. | 

## Methods

### NewAuthorizedTimeWindowViolation

`func NewAuthorizedTimeWindowViolation(type_ string, forbiddenWindow TimeWindow, actualTime string, ) *AuthorizedTimeWindowViolation`

NewAuthorizedTimeWindowViolation instantiates a new AuthorizedTimeWindowViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizedTimeWindowViolationWithDefaults

`func NewAuthorizedTimeWindowViolationWithDefaults() *AuthorizedTimeWindowViolation`

NewAuthorizedTimeWindowViolationWithDefaults instantiates a new AuthorizedTimeWindowViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizedTimeWindowViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizedTimeWindowViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizedTimeWindowViolation) SetType(v string)`

SetType sets Type field to given value.


### GetWayPointIndex

`func (o *AuthorizedTimeWindowViolation) GetWayPointIndex() int32`

GetWayPointIndex returns the WayPointIndex field if non-nil, zero value otherwise.

### GetWayPointIndexOk

`func (o *AuthorizedTimeWindowViolation) GetWayPointIndexOk() (*int32, bool)`

GetWayPointIndexOk returns a tuple with the WayPointIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWayPointIndex

`func (o *AuthorizedTimeWindowViolation) SetWayPointIndex(v int32)`

SetWayPointIndex sets WayPointIndex field to given value.

### HasWayPointIndex

`func (o *AuthorizedTimeWindowViolation) HasWayPointIndex() bool`

HasWayPointIndex returns a boolean if a field has been set.

### GetForbiddenWindow

`func (o *AuthorizedTimeWindowViolation) GetForbiddenWindow() TimeWindow`

GetForbiddenWindow returns the ForbiddenWindow field if non-nil, zero value otherwise.

### GetForbiddenWindowOk

`func (o *AuthorizedTimeWindowViolation) GetForbiddenWindowOk() (*TimeWindow, bool)`

GetForbiddenWindowOk returns a tuple with the ForbiddenWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenWindow

`func (o *AuthorizedTimeWindowViolation) SetForbiddenWindow(v TimeWindow)`

SetForbiddenWindow sets ForbiddenWindow field to given value.


### GetActualTime

`func (o *AuthorizedTimeWindowViolation) GetActualTime() string`

GetActualTime returns the ActualTime field if non-nil, zero value otherwise.

### GetActualTimeOk

`func (o *AuthorizedTimeWindowViolation) GetActualTimeOk() (*string, bool)`

GetActualTimeOk returns a tuple with the ActualTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTime

`func (o *AuthorizedTimeWindowViolation) SetActualTime(v string)`

SetActualTime sets ActualTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


