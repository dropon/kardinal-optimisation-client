# PlanLight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**AgencyId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 
**Version** | Pointer to **int32** | The plan version. | [optional] [readonly] 
**Running** | Pointer to **bool** | To know if the plan is running. | [optional] [readonly] 
**Status** | Pointer to [**PlanStatus**](PlanStatus.md) |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPlanLight

`func NewPlanLight() *PlanLight`

NewPlanLight instantiates a new PlanLight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanLightWithDefaults

`func NewPlanLightWithDefaults() *PlanLight`

NewPlanLightWithDefaults instantiates a new PlanLight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanLight) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanLight) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanLight) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *PlanLight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgencyId

`func (o *PlanLight) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *PlanLight) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *PlanLight) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.

### HasAgencyId

`func (o *PlanLight) HasAgencyId() bool`

HasAgencyId returns a boolean if a field has been set.

### GetVersion

`func (o *PlanLight) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlanLight) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlanLight) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PlanLight) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRunning

`func (o *PlanLight) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *PlanLight) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *PlanLight) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *PlanLight) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetStatus

`func (o *PlanLight) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlanLight) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlanLight) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlanLight) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProperties

`func (o *PlanLight) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PlanLight) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PlanLight) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PlanLight) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


