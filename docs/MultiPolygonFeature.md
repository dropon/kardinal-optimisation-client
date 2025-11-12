# MultiPolygonFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FeatureType**](FeatureType.md) |  | 
**Geometry** | [**MultiPolygon**](MultiPolygon.md) |  | 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMultiPolygonFeature

`func NewMultiPolygonFeature(type_ FeatureType, geometry MultiPolygon, ) *MultiPolygonFeature`

NewMultiPolygonFeature instantiates a new MultiPolygonFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiPolygonFeatureWithDefaults

`func NewMultiPolygonFeatureWithDefaults() *MultiPolygonFeature`

NewMultiPolygonFeatureWithDefaults instantiates a new MultiPolygonFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MultiPolygonFeature) GetType() FeatureType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultiPolygonFeature) GetTypeOk() (*FeatureType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultiPolygonFeature) SetType(v FeatureType)`

SetType sets Type field to given value.


### GetGeometry

`func (o *MultiPolygonFeature) GetGeometry() MultiPolygon`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *MultiPolygonFeature) GetGeometryOk() (*MultiPolygon, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *MultiPolygonFeature) SetGeometry(v MultiPolygon)`

SetGeometry sets Geometry field to given value.


### GetProperties

`func (o *MultiPolygonFeature) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MultiPolygonFeature) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MultiPolygonFeature) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MultiPolygonFeature) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


