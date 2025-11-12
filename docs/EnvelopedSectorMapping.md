# EnvelopedSectorMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**map[string]MultiPolygonFeature**](MultiPolygonFeature.md) | A sector mapping consists in GeoJSON MultiPolygon features, mapped by sector names (one multipolygon per sector name): each multipolygon defines a geographical sector associated with a name.  When a sector mapping exists in an agency, it is applied on every incoming ARO plan which fulfills the following condition: - Each plan order contains one and only one stop.  If the condition is fulfilled, the sector mapping is applied as follows: - For each stop in the plan, we check whether the stop position is contained inside a sector (a multipolygon) of the sector mapping, - If so, we check whether the sector name is a prefix for a resource id of the plan, - If so, we set the resource mode to \&quot;fixed\&quot;, and we push the stop inside the resource assignments, with the status \&quot;assigned\&quot;.  Once done, the plan follows the rest of the usual ARO process.  Note: if an error occurs during this process, we abort the application of the sector mapping, and the plan simply goes on with the usual ARO process.  | [optional] 
**AgencyId** | Pointer to [**interface{}**](AnyType.md) |  | [optional] 

## Methods

### NewEnvelopedSectorMapping

`func NewEnvelopedSectorMapping() *EnvelopedSectorMapping`

NewEnvelopedSectorMapping instantiates a new EnvelopedSectorMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedSectorMappingWithDefaults

`func NewEnvelopedSectorMappingWithDefaults() *EnvelopedSectorMapping`

NewEnvelopedSectorMappingWithDefaults instantiates a new EnvelopedSectorMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EnvelopedSectorMapping) GetItem() map[string]MultiPolygonFeature`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EnvelopedSectorMapping) GetItemOk() (*map[string]MultiPolygonFeature, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EnvelopedSectorMapping) SetItem(v map[string]MultiPolygonFeature)`

SetItem sets Item field to given value.

### HasItem

`func (o *EnvelopedSectorMapping) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetAgencyId

`func (o *EnvelopedSectorMapping) GetAgencyId() interface{}`

GetAgencyId returns the AgencyId field if non-nil, zero value otherwise.

### GetAgencyIdOk

`func (o *EnvelopedSectorMapping) GetAgencyIdOk() (*interface{}, bool)`

GetAgencyIdOk returns a tuple with the AgencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyId

`func (o *EnvelopedSectorMapping) SetAgencyId(v interface{})`

SetAgencyId sets AgencyId field to given value.

### HasAgencyId

`func (o *EnvelopedSectorMapping) HasAgencyId() bool`

HasAgencyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


