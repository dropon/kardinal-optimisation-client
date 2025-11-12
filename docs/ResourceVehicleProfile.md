# ResourceVehicleProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Kmph** | Pointer to **float32** |  | [optional] 
**ExcludedCountries** | Pointer to **[]string** |  | [optional] 
**WithTraffic** | Pointer to **bool** |  | [optional] 
**GrossWeight** | Pointer to **float32** |  | [optional] 
**AvoidTollRoad** | Pointer to **bool** |  | [optional] 
**ShippedHazardousGoods** | Pointer to [**[]ShippedHazardousGood**](ShippedHazardousGood.md) |  | [optional] 

## Methods

### NewResourceVehicleProfile

`func NewResourceVehicleProfile(type_ string, ) *ResourceVehicleProfile`

NewResourceVehicleProfile instantiates a new ResourceVehicleProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceVehicleProfileWithDefaults

`func NewResourceVehicleProfileWithDefaults() *ResourceVehicleProfile`

NewResourceVehicleProfileWithDefaults instantiates a new ResourceVehicleProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceVehicleProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceVehicleProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceVehicleProfile) SetType(v string)`

SetType sets Type field to given value.


### GetKmph

`func (o *ResourceVehicleProfile) GetKmph() float32`

GetKmph returns the Kmph field if non-nil, zero value otherwise.

### GetKmphOk

`func (o *ResourceVehicleProfile) GetKmphOk() (*float32, bool)`

GetKmphOk returns a tuple with the Kmph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmph

`func (o *ResourceVehicleProfile) SetKmph(v float32)`

SetKmph sets Kmph field to given value.

### HasKmph

`func (o *ResourceVehicleProfile) HasKmph() bool`

HasKmph returns a boolean if a field has been set.

### GetExcludedCountries

`func (o *ResourceVehicleProfile) GetExcludedCountries() []string`

GetExcludedCountries returns the ExcludedCountries field if non-nil, zero value otherwise.

### GetExcludedCountriesOk

`func (o *ResourceVehicleProfile) GetExcludedCountriesOk() (*[]string, bool)`

GetExcludedCountriesOk returns a tuple with the ExcludedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCountries

`func (o *ResourceVehicleProfile) SetExcludedCountries(v []string)`

SetExcludedCountries sets ExcludedCountries field to given value.

### HasExcludedCountries

`func (o *ResourceVehicleProfile) HasExcludedCountries() bool`

HasExcludedCountries returns a boolean if a field has been set.

### GetWithTraffic

`func (o *ResourceVehicleProfile) GetWithTraffic() bool`

GetWithTraffic returns the WithTraffic field if non-nil, zero value otherwise.

### GetWithTrafficOk

`func (o *ResourceVehicleProfile) GetWithTrafficOk() (*bool, bool)`

GetWithTrafficOk returns a tuple with the WithTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithTraffic

`func (o *ResourceVehicleProfile) SetWithTraffic(v bool)`

SetWithTraffic sets WithTraffic field to given value.

### HasWithTraffic

`func (o *ResourceVehicleProfile) HasWithTraffic() bool`

HasWithTraffic returns a boolean if a field has been set.

### GetGrossWeight

`func (o *ResourceVehicleProfile) GetGrossWeight() float32`

GetGrossWeight returns the GrossWeight field if non-nil, zero value otherwise.

### GetGrossWeightOk

`func (o *ResourceVehicleProfile) GetGrossWeightOk() (*float32, bool)`

GetGrossWeightOk returns a tuple with the GrossWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossWeight

`func (o *ResourceVehicleProfile) SetGrossWeight(v float32)`

SetGrossWeight sets GrossWeight field to given value.

### HasGrossWeight

`func (o *ResourceVehicleProfile) HasGrossWeight() bool`

HasGrossWeight returns a boolean if a field has been set.

### GetAvoidTollRoad

`func (o *ResourceVehicleProfile) GetAvoidTollRoad() bool`

GetAvoidTollRoad returns the AvoidTollRoad field if non-nil, zero value otherwise.

### GetAvoidTollRoadOk

`func (o *ResourceVehicleProfile) GetAvoidTollRoadOk() (*bool, bool)`

GetAvoidTollRoadOk returns a tuple with the AvoidTollRoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvoidTollRoad

`func (o *ResourceVehicleProfile) SetAvoidTollRoad(v bool)`

SetAvoidTollRoad sets AvoidTollRoad field to given value.

### HasAvoidTollRoad

`func (o *ResourceVehicleProfile) HasAvoidTollRoad() bool`

HasAvoidTollRoad returns a boolean if a field has been set.

### GetShippedHazardousGoods

`func (o *ResourceVehicleProfile) GetShippedHazardousGoods() []ShippedHazardousGood`

GetShippedHazardousGoods returns the ShippedHazardousGoods field if non-nil, zero value otherwise.

### GetShippedHazardousGoodsOk

`func (o *ResourceVehicleProfile) GetShippedHazardousGoodsOk() (*[]ShippedHazardousGood, bool)`

GetShippedHazardousGoodsOk returns a tuple with the ShippedHazardousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedHazardousGoods

`func (o *ResourceVehicleProfile) SetShippedHazardousGoods(v []ShippedHazardousGood)`

SetShippedHazardousGoods sets ShippedHazardousGoods field to given value.

### HasShippedHazardousGoods

`func (o *ResourceVehicleProfile) HasShippedHazardousGoods() bool`

HasShippedHazardousGoods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


