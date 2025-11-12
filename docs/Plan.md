# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**interface{}**](AnyType.md) |  | 
**AgencyId** | [**interface{}**](AnyType.md) |  | 
**Version** | Pointer to **int32** | The plan version. | [optional] [readonly] 
**Running** | Pointer to **bool** | To know if the plan is running. | [optional] [readonly] 
**Status** | Pointer to [**PlanStatus**](PlanStatus.md) |  | [optional] 
**State** | Pointer to [**PlanState**](PlanState.md) |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Resources** | [**[]Resource**](Resource.md) |  | 
**Orders** | Pointer to [**[]Order**](Order.md) |  | [optional] 
**AdditionalOperationDurations** | Pointer to [**[]AdditionalOperationDuration**](AdditionalOperationDuration.md) | Additional operation time for a resource and a stop, according to tags (pairs of tags must be unique). | [optional] 
**OperationDurationPoliciesByResourceTag** | Pointer to [**map[string][]OperationDurationPolicy**](array.md) | Policies to remove the operation durations, by resource tag. | [optional] 
**ForbiddenAssignments** | Pointer to [**[]ForbiddenAssignment**](ForbiddenAssignment.md) |  | [optional] 
**IncompatibleStopTags** | Pointer to **[][]string** | Incompatibilities between stop tags (all elements must be different regardless of order). | [optional] 
**AdditionalConstraints** | Pointer to [**[]PlanAdditionalConstraintsInner**](PlanAdditionalConstraintsInner.md) |  | [optional] 
**GlobalConstraints** | Pointer to [**[]PlanGlobalConstraintsInner**](PlanGlobalConstraintsInner.md) | List of global constraints to be satisfied by the returned solution. | [optional] 
**AccessDurationsByStopTag** | Pointer to **map[string]string** | Access durations is an additional duration before the beginning of a group of stops with the same stop tag. | [optional] 
**OverlappingCapacitiesByStopTag** | Pointer to **map[string]int32** | This field allows users to define a limit in the number of resources that are simultaneously present at stops sharing the same stop tag. | [optional] 
**SetupDurations** | Pointer to [**[]SetupDuration**](SetupDuration.md) |  | [optional] 
**Objectives** | Pointer to [**[]PlanObjectivesInner**](PlanObjectivesInner.md) |  | [optional] 
**MaxOptimizationDuration** | Pointer to **string** | A period of time, expressed in the ISO8601 **duration** format. | [optional] 
**Tz** | Pointer to **string** | The time zone is a string code which identifies a region of the world in the \&quot;time zone database\&quot;, also called \&quot;tz database\&quot;. The tz database is a partition of the world into regions where local clocks all show the same time. This database gives the rules for time offset and daylight saving time in each region.  How do we use it?  In order to work with time events accurately, we usually use datetimes in the iso-8601 format, without explicit time zone. This format is quite well suported by many programming languages, and it is well suited for technical data exchange. But it is not easy to use for humans.  For instance, here are three datetimes in iso-8601 format, which give the same exact moment in time: - \&quot;2025-05-22T05:43:00Z\&quot; - \&quot;2025-05-22T06:43:00+01:00\&quot; - \&quot;2025-05-22T07:43:00+02:00\&quot;  For a non-technical user, it is difficult to know how to relate this to the time displayed on a watch or a clock.  We improve the user experience by adding the support of local datetimes, thanks to the use of the time zone, which allows to transform a local datetime into an iso-8601 datetime: - local datetime + timezone (tz) &#x3D; iso-8601 datetime  For instance, here are five datetimes which all give the same exact moment in time: - \&quot;2025-05-22T05:43:00Z\&quot; - \&quot;2025-05-22T06:43:00+01:00\&quot; - \&quot;2025-05-22T07:43:00+02:00\&quot; - \&quot;2025-05-22 07:43:00\&quot;       + timezone \&quot;tz\&quot;: \&quot;Europe/Paris\&quot; - \&quot;2025-05-22 07:43\&quot;          + timezone \&quot;tz\&quot;: \&quot;Europe/Paris\&quot;  Note: the last example (\&quot;2025-05-22 07:43\&quot;) illustrates the support of local datetimes without seconds, which can be very practical for users.  In order for local datetimes to be supported, some JSON input objects contain a \&quot;tz\&quot; time zone property. This \&quot;tz\&quot; property is used to pre-process the JSON input payload, like this: - We check if a valid timezone can be extracted from the \&quot;tz\&quot; property, - If so, we perform the following actions:   - Walk through the whole JSON content to look for local datetimes,   - Use the timezone to transform each local datetime into an iso-8601 datetime.  Important: some objects contain a \&quot;properties\&quot; sub-object, which is a map of custom client data; the content of the \&quot;properties\&quot; sub-objects is always excluded from the time zone pre-processing.  | [optional] 
**CreatedAt** | Pointer to **string** | The plan&#39;s creation datetime. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableString** | The plan&#39;s last update datetime. | [optional] [readonly] 
**ArchivedAt** | Pointer to **NullableString** | The plan&#39;s archiving datetime: if not null, the plan is archived. | [optional] [readonly] 
**LateDeparture** | Pointer to **bool** | True if lateDeparture is requested for Resources, false otherwise. | [optional] [default to false]
**SharedCapacities** | Pointer to **bool** | True if sharedCapacities is requested for Resources, false otherwise. It enables resources to share capacities between stops of different orders. | [optional] [default to false]

## Methods

### NewPlan

`func NewPlan(id interface{}, agencyId interface{}, resources []Resource, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plan) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v interface{})`

SetId sets Id field to given value.


### GetAgencyId

`func (o *Plan) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *Plan) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *Plan) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.


### GetVersion

`func (o *Plan) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Plan) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Plan) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Plan) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRunning

`func (o *Plan) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *Plan) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *Plan) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *Plan) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetStatus

`func (o *Plan) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Plan) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Plan) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Plan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetState

`func (o *Plan) GetState() PlanState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Plan) GetStateOk() (*PlanState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Plan) SetState(v PlanState)`

SetState sets State field to given value.

### HasState

`func (o *Plan) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProperties

`func (o *Plan) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Plan) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Plan) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Plan) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetResources

`func (o *Plan) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Plan) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Plan) SetResources(v []Resource)`

SetResources sets Resources field to given value.


### GetOrders

`func (o *Plan) GetOrders() []Order`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *Plan) GetOrdersOk() (*[]Order, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *Plan) SetOrders(v []Order)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *Plan) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetAdditionalOperationDurations

`func (o *Plan) GetAdditionalOperationDurations() []AdditionalOperationDuration`

GetAdditionalOperationDurations returns the AdditionalOperationDurations field if non-nil, zero value otherwise.

### GetAdditionalOperationDurationsOk

`func (o *Plan) GetAdditionalOperationDurationsOk() (*[]AdditionalOperationDuration, bool)`

GetAdditionalOperationDurationsOk returns a tuple with the AdditionalOperationDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalOperationDurations

`func (o *Plan) SetAdditionalOperationDurations(v []AdditionalOperationDuration)`

SetAdditionalOperationDurations sets AdditionalOperationDurations field to given value.

### HasAdditionalOperationDurations

`func (o *Plan) HasAdditionalOperationDurations() bool`

HasAdditionalOperationDurations returns a boolean if a field has been set.

### GetOperationDurationPoliciesByResourceTag

`func (o *Plan) GetOperationDurationPoliciesByResourceTag() map[string][]OperationDurationPolicy`

GetOperationDurationPoliciesByResourceTag returns the OperationDurationPoliciesByResourceTag field if non-nil, zero value otherwise.

### GetOperationDurationPoliciesByResourceTagOk

`func (o *Plan) GetOperationDurationPoliciesByResourceTagOk() (*map[string][]OperationDurationPolicy, bool)`

GetOperationDurationPoliciesByResourceTagOk returns a tuple with the OperationDurationPoliciesByResourceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationDurationPoliciesByResourceTag

`func (o *Plan) SetOperationDurationPoliciesByResourceTag(v map[string][]OperationDurationPolicy)`

SetOperationDurationPoliciesByResourceTag sets OperationDurationPoliciesByResourceTag field to given value.

### HasOperationDurationPoliciesByResourceTag

`func (o *Plan) HasOperationDurationPoliciesByResourceTag() bool`

HasOperationDurationPoliciesByResourceTag returns a boolean if a field has been set.

### GetForbiddenAssignments

`func (o *Plan) GetForbiddenAssignments() []ForbiddenAssignment`

GetForbiddenAssignments returns the ForbiddenAssignments field if non-nil, zero value otherwise.

### GetForbiddenAssignmentsOk

`func (o *Plan) GetForbiddenAssignmentsOk() (*[]ForbiddenAssignment, bool)`

GetForbiddenAssignmentsOk returns a tuple with the ForbiddenAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAssignments

`func (o *Plan) SetForbiddenAssignments(v []ForbiddenAssignment)`

SetForbiddenAssignments sets ForbiddenAssignments field to given value.

### HasForbiddenAssignments

`func (o *Plan) HasForbiddenAssignments() bool`

HasForbiddenAssignments returns a boolean if a field has been set.

### GetIncompatibleStopTags

`func (o *Plan) GetIncompatibleStopTags() [][]string`

GetIncompatibleStopTags returns the IncompatibleStopTags field if non-nil, zero value otherwise.

### GetIncompatibleStopTagsOk

`func (o *Plan) GetIncompatibleStopTagsOk() (*[][]string, bool)`

GetIncompatibleStopTagsOk returns a tuple with the IncompatibleStopTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleStopTags

`func (o *Plan) SetIncompatibleStopTags(v [][]string)`

SetIncompatibleStopTags sets IncompatibleStopTags field to given value.

### HasIncompatibleStopTags

`func (o *Plan) HasIncompatibleStopTags() bool`

HasIncompatibleStopTags returns a boolean if a field has been set.

### GetAdditionalConstraints

`func (o *Plan) GetAdditionalConstraints() []PlanAdditionalConstraintsInner`

GetAdditionalConstraints returns the AdditionalConstraints field if non-nil, zero value otherwise.

### GetAdditionalConstraintsOk

`func (o *Plan) GetAdditionalConstraintsOk() (*[]PlanAdditionalConstraintsInner, bool)`

GetAdditionalConstraintsOk returns a tuple with the AdditionalConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConstraints

`func (o *Plan) SetAdditionalConstraints(v []PlanAdditionalConstraintsInner)`

SetAdditionalConstraints sets AdditionalConstraints field to given value.

### HasAdditionalConstraints

`func (o *Plan) HasAdditionalConstraints() bool`

HasAdditionalConstraints returns a boolean if a field has been set.

### GetGlobalConstraints

`func (o *Plan) GetGlobalConstraints() []PlanGlobalConstraintsInner`

GetGlobalConstraints returns the GlobalConstraints field if non-nil, zero value otherwise.

### GetGlobalConstraintsOk

`func (o *Plan) GetGlobalConstraintsOk() (*[]PlanGlobalConstraintsInner, bool)`

GetGlobalConstraintsOk returns a tuple with the GlobalConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalConstraints

`func (o *Plan) SetGlobalConstraints(v []PlanGlobalConstraintsInner)`

SetGlobalConstraints sets GlobalConstraints field to given value.

### HasGlobalConstraints

`func (o *Plan) HasGlobalConstraints() bool`

HasGlobalConstraints returns a boolean if a field has been set.

### GetAccessDurationsByStopTag

`func (o *Plan) GetAccessDurationsByStopTag() map[string]string`

GetAccessDurationsByStopTag returns the AccessDurationsByStopTag field if non-nil, zero value otherwise.

### GetAccessDurationsByStopTagOk

`func (o *Plan) GetAccessDurationsByStopTagOk() (*map[string]string, bool)`

GetAccessDurationsByStopTagOk returns a tuple with the AccessDurationsByStopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationsByStopTag

`func (o *Plan) SetAccessDurationsByStopTag(v map[string]string)`

SetAccessDurationsByStopTag sets AccessDurationsByStopTag field to given value.

### HasAccessDurationsByStopTag

`func (o *Plan) HasAccessDurationsByStopTag() bool`

HasAccessDurationsByStopTag returns a boolean if a field has been set.

### GetOverlappingCapacitiesByStopTag

`func (o *Plan) GetOverlappingCapacitiesByStopTag() map[string]int32`

GetOverlappingCapacitiesByStopTag returns the OverlappingCapacitiesByStopTag field if non-nil, zero value otherwise.

### GetOverlappingCapacitiesByStopTagOk

`func (o *Plan) GetOverlappingCapacitiesByStopTagOk() (*map[string]int32, bool)`

GetOverlappingCapacitiesByStopTagOk returns a tuple with the OverlappingCapacitiesByStopTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlappingCapacitiesByStopTag

`func (o *Plan) SetOverlappingCapacitiesByStopTag(v map[string]int32)`

SetOverlappingCapacitiesByStopTag sets OverlappingCapacitiesByStopTag field to given value.

### HasOverlappingCapacitiesByStopTag

`func (o *Plan) HasOverlappingCapacitiesByStopTag() bool`

HasOverlappingCapacitiesByStopTag returns a boolean if a field has been set.

### GetSetupDurations

`func (o *Plan) GetSetupDurations() []SetupDuration`

GetSetupDurations returns the SetupDurations field if non-nil, zero value otherwise.

### GetSetupDurationsOk

`func (o *Plan) GetSetupDurationsOk() (*[]SetupDuration, bool)`

GetSetupDurationsOk returns a tuple with the SetupDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupDurations

`func (o *Plan) SetSetupDurations(v []SetupDuration)`

SetSetupDurations sets SetupDurations field to given value.

### HasSetupDurations

`func (o *Plan) HasSetupDurations() bool`

HasSetupDurations returns a boolean if a field has been set.

### GetObjectives

`func (o *Plan) GetObjectives() []PlanObjectivesInner`

GetObjectives returns the Objectives field if non-nil, zero value otherwise.

### GetObjectivesOk

`func (o *Plan) GetObjectivesOk() (*[]PlanObjectivesInner, bool)`

GetObjectivesOk returns a tuple with the Objectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectives

`func (o *Plan) SetObjectives(v []PlanObjectivesInner)`

SetObjectives sets Objectives field to given value.

### HasObjectives

`func (o *Plan) HasObjectives() bool`

HasObjectives returns a boolean if a field has been set.

### GetMaxOptimizationDuration

`func (o *Plan) GetMaxOptimizationDuration() string`

GetMaxOptimizationDuration returns the MaxOptimizationDuration field if non-nil, zero value otherwise.

### GetMaxOptimizationDurationOk

`func (o *Plan) GetMaxOptimizationDurationOk() (*string, bool)`

GetMaxOptimizationDurationOk returns a tuple with the MaxOptimizationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOptimizationDuration

`func (o *Plan) SetMaxOptimizationDuration(v string)`

SetMaxOptimizationDuration sets MaxOptimizationDuration field to given value.

### HasMaxOptimizationDuration

`func (o *Plan) HasMaxOptimizationDuration() bool`

HasMaxOptimizationDuration returns a boolean if a field has been set.

### GetTz

`func (o *Plan) GetTz() string`

GetTz returns the Tz field if non-nil, zero value otherwise.

### GetTzOk

`func (o *Plan) GetTzOk() (*string, bool)`

GetTzOk returns a tuple with the Tz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTz

`func (o *Plan) SetTz(v string)`

SetTz sets Tz field to given value.

### HasTz

`func (o *Plan) HasTz() bool`

HasTz returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Plan) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Plan) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Plan) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Plan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Plan) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Plan) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Plan) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Plan) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Plan) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Plan) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetArchivedAt

`func (o *Plan) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Plan) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Plan) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Plan) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *Plan) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *Plan) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetLateDeparture

`func (o *Plan) GetLateDeparture() bool`

GetLateDeparture returns the LateDeparture field if non-nil, zero value otherwise.

### GetLateDepartureOk

`func (o *Plan) GetLateDepartureOk() (*bool, bool)`

GetLateDepartureOk returns a tuple with the LateDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateDeparture

`func (o *Plan) SetLateDeparture(v bool)`

SetLateDeparture sets LateDeparture field to given value.

### HasLateDeparture

`func (o *Plan) HasLateDeparture() bool`

HasLateDeparture returns a boolean if a field has been set.

### GetSharedCapacities

`func (o *Plan) GetSharedCapacities() bool`

GetSharedCapacities returns the SharedCapacities field if non-nil, zero value otherwise.

### GetSharedCapacitiesOk

`func (o *Plan) GetSharedCapacitiesOk() (*bool, bool)`

GetSharedCapacitiesOk returns a tuple with the SharedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedCapacities

`func (o *Plan) SetSharedCapacities(v bool)`

SetSharedCapacities sets SharedCapacities field to given value.

### HasSharedCapacities

`func (o *Plan) HasSharedCapacities() bool`

HasSharedCapacities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


