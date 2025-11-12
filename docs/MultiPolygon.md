# MultiPolygon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MultiPolygonType**](MultiPolygonType.md) |  | 
**Coordinates** | [**[][][][]float32**]([][][]float32.md) | A multi polygon is an array of polygons. | 

## Methods

### NewMultiPolygon

`func NewMultiPolygon(type_ MultiPolygonType, coordinates [][][][]float32, ) *MultiPolygon`

NewMultiPolygon instantiates a new MultiPolygon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiPolygonWithDefaults

`func NewMultiPolygonWithDefaults() *MultiPolygon`

NewMultiPolygonWithDefaults instantiates a new MultiPolygon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MultiPolygon) GetType() MultiPolygonType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultiPolygon) GetTypeOk() (*MultiPolygonType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultiPolygon) SetType(v MultiPolygonType)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *MultiPolygon) GetCoordinates() [][][][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *MultiPolygon) GetCoordinatesOk() (*[][][][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *MultiPolygon) SetCoordinates(v [][][][]float32)`

SetCoordinates sets Coordinates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


