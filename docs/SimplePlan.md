# SimplePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Universally Unique Identifier. | [readonly] 
**AgencyId** | [**interface{}**](AnyType.md) |  | 
**Version** | Pointer to **int32** | The plan version. | [optional] [readonly] 
**State** | Pointer to [**PlanState**](PlanState.md) |  | [optional] 
**Resources** | [**[]SimpleResource**](SimpleResource.md) |  | 
**Stops** | Pointer to [**[]SimpleStop**](SimpleStop.md) |  | [optional] 
**Objective** | Pointer to [**SimpleObjective**](SimpleObjective.md) |  | [optional] [default to SIMPLEOBJECTIVE_MINIMIZE_RESOURCES]
**Tz** | Pointer to **string** | The time zone is a string code which identifies a region of the world in the \&quot;time zone database\&quot;, also called \&quot;tz database\&quot;. The tz database is a partition of the world into regions where local clocks all show the same time. This database gives the rules for time offset and daylight saving time in each region.  How do we use it?  In order to work with time events accurately, we usually use datetimes in the iso-8601 format, without explicit time zone. This format is quite well suported by many programming languages, and it is well suited for technical data exchange. But it is not easy to use for humans.  For instance, here are three datetimes in iso-8601 format, which give the same exact moment in time: - \&quot;2025-05-22T05:43:00Z\&quot; - \&quot;2025-05-22T06:43:00+01:00\&quot; - \&quot;2025-05-22T07:43:00+02:00\&quot;  For a non-technical user, it is difficult to know how to relate this to the time displayed on a watch or a clock.  We improve the user experience by adding the support of local datetimes, thanks to the use of the time zone, which allows to transform a local datetime into an iso-8601 datetime: - local datetime + timezone (tz) &#x3D; iso-8601 datetime  For instance, here are five datetimes which all give the same exact moment in time: - \&quot;2025-05-22T05:43:00Z\&quot; - \&quot;2025-05-22T06:43:00+01:00\&quot; - \&quot;2025-05-22T07:43:00+02:00\&quot; - \&quot;2025-05-22 07:43:00\&quot;       + timezone \&quot;tz\&quot;: \&quot;Europe/Paris\&quot; - \&quot;2025-05-22 07:43\&quot;          + timezone \&quot;tz\&quot;: \&quot;Europe/Paris\&quot;  Note: the last example (\&quot;2025-05-22 07:43\&quot;) illustrates the support of local datetimes without seconds, which can be very practical for users.  In order for local datetimes to be supported, some JSON input objects contain a \&quot;tz\&quot; time zone property. This \&quot;tz\&quot; property is used to pre-process the JSON input payload, like this: - We check if a valid timezone can be extracted from the \&quot;tz\&quot; property, - If so, we perform the following actions:   - Walk through the whole JSON content to look for local datetimes,   - Use the timezone to transform each local datetime into an iso-8601 datetime.  Important: some objects contain a \&quot;properties\&quot; sub-object, which is a map of custom client data; the content of the \&quot;properties\&quot; sub-objects is always excluded from the time zone pre-processing.  | [optional] 
**LateDeparture** | Pointer to **bool** | True if lateDeparture is requested for Resources, false otherwise. | [optional] [default to false]
**CreatedAt** | Pointer to **string** | The plan&#39;s creation datetime. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableString** | The plan&#39;s last update datetime. | [optional] [readonly] 
**ArchivedAt** | Pointer to **NullableString** | The plan&#39;s archiving datetime: if not null, the plan is archived. | [optional] [readonly] 

## Methods

### NewSimplePlan

`func NewSimplePlan(id string, agencyId interface{}, resources []SimpleResource, ) *SimplePlan`

NewSimplePlan instantiates a new SimplePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePlanWithDefaults

`func NewSimplePlanWithDefaults() *SimplePlan`

NewSimplePlanWithDefaults instantiates a new SimplePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimplePlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplePlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplePlan) SetId(v string)`

SetId sets Id field to given value.


### GetAgencyId

`func (o *SimplePlan) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *SimplePlan) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *SimplePlan) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.


### GetVersion

`func (o *SimplePlan) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SimplePlan) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SimplePlan) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SimplePlan) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetState

`func (o *SimplePlan) GetState() PlanState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SimplePlan) GetStateOk() (*PlanState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SimplePlan) SetState(v PlanState)`

SetState sets State field to given value.

### HasState

`func (o *SimplePlan) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResources

`func (o *SimplePlan) GetResources() []SimpleResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SimplePlan) GetResourcesOk() (*[]SimpleResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SimplePlan) SetResources(v []SimpleResource)`

SetResources sets Resources field to given value.


### GetStops

`func (o *SimplePlan) GetStops() []SimpleStop`

GetStops returns the Stops field if non-nil, zero value otherwise.

### GetStopsOk

`func (o *SimplePlan) GetStopsOk() (*[]SimpleStop, bool)`

GetStopsOk returns a tuple with the Stops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStops

`func (o *SimplePlan) SetStops(v []SimpleStop)`

SetStops sets Stops field to given value.

### HasStops

`func (o *SimplePlan) HasStops() bool`

HasStops returns a boolean if a field has been set.

### GetObjective

`func (o *SimplePlan) GetObjective() SimpleObjective`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *SimplePlan) GetObjectiveOk() (*SimpleObjective, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *SimplePlan) SetObjective(v SimpleObjective)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *SimplePlan) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetTz

`func (o *SimplePlan) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *SimplePlan) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *SimplePlan) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *SimplePlan) HasTz() bool`

HasTz returns a boolean if a field has been set.

### GetLateDeparture

`func (o *SimplePlan) GetLateDeparture() bool`

GetLateDeparture returns the LateDeparture field if non-nil, zero value otherwise.

### GetLateDepartureOk

`func (o *SimplePlan) GetLateDepartureOk() (*bool, bool)`

GetLateDepartureOk returns a tuple with the LateDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateDeparture

`func (o *SimplePlan) SetLateDeparture(v bool)`

SetLateDeparture sets LateDeparture field to given value.

### HasLateDeparture

`func (o *SimplePlan) HasLateDeparture() bool`

HasLateDeparture returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SimplePlan) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimplePlan) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimplePlan) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SimplePlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SimplePlan) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimplePlan) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimplePlan) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SimplePlan) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SimplePlan) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SimplePlan) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetArchivedAt

`func (o *SimplePlan) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *SimplePlan) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *SimplePlan) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *SimplePlan) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *SimplePlan) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *SimplePlan) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


