# VehicleProfileTruck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**GrossWeight** | Pointer to **float32** |  | [optional] 
**WithTraffic** | Pointer to **bool** |  | [optional] 
**AvoidTollRoad** | Pointer to **bool** |  | [optional] 
**ShippedHazardousGoods** | Pointer to [**[]ShippedHazardousGood**](ShippedHazardousGood.md) |  | [optional] 
**ExcludedCountries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVehicleProfileTruck

`func NewVehicleProfileTruck(type_ string, ) *VehicleProfileTruck`

NewVehicleProfileTruck instantiates a new VehicleProfileTruck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleProfileTruckWithDefaults

`func NewVehicleProfileTruckWithDefaults() *VehicleProfileTruck`

NewVehicleProfileTruckWithDefaults instantiates a new VehicleProfileTruck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VehicleProfileTruck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VehicleProfileTruck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VehicleProfileTruck) SetType(v string)`

SetType sets Type field to given value.


### GetGrossWeight

`func (o *VehicleProfileTruck) GetGrossWeight() float32`

GetGrossWeight returns the GrossWeight field if non-nil, zero value otherwise.

### GetGrossWeightOk

`func (o *VehicleProfileTruck) GetGrossWeightOk() (*float32, bool)`

GetGrossWeightOk returns a tuple with the GrossWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossWeight

`func (o *VehicleProfileTruck) SetGrossWeight(v float32)`

SetGrossWeight sets GrossWeight field to given value.

### HasGrossWeight

`func (o *VehicleProfileTruck) HasGrossWeight() bool`

HasGrossWeight returns a boolean if a field has been set.

### GetWithTraffic

`func (o *VehicleProfileTruck) GetWithTraffic() bool`

GetWithTraffic returns the WithTraffic field if non-nil, zero value otherwise.

### GetWithTrafficOk

`func (o *VehicleProfileTruck) GetWithTrafficOk() (*bool, bool)`

GetWithTrafficOk returns a tuple with the WithTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithTraffic

`func (o *VehicleProfileTruck) SetWithTraffic(v bool)`

SetWithTraffic sets WithTraffic field to given value.

### HasWithTraffic

`func (o *VehicleProfileTruck) HasWithTraffic() bool`

HasWithTraffic returns a boolean if a field has been set.

### GetAvoidTollRoad

`func (o *VehicleProfileTruck) GetAvoidTollRoad() bool`

GetAvoidTollRoad returns the AvoidTollRoad field if non-nil, zero value otherwise.

### GetAvoidTollRoadOk

`func (o *VehicleProfileTruck) GetAvoidTollRoadOk() (*bool, bool)`

GetAvoidTollRoadOk returns a tuple with the AvoidTollRoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvoidTollRoad

`func (o *VehicleProfileTruck) SetAvoidTollRoad(v bool)`

SetAvoidTollRoad sets AvoidTollRoad field to given value.

### HasAvoidTollRoad

`func (o *VehicleProfileTruck) HasAvoidTollRoad() bool`

HasAvoidTollRoad returns a boolean if a field has been set.

### GetShippedHazardousGoods

`func (o *VehicleProfileTruck) GetShippedHazardousGoods() []ShippedHazardousGood`

GetShippedHazardousGoods returns the ShippedHazardousGoods field if non-nil, zero value otherwise.

### GetShippedHazardousGoodsOk

`func (o *VehicleProfileTruck) GetShippedHazardousGoodsOk() (*[]ShippedHazardousGood, bool)`

GetShippedHazardousGoodsOk returns a tuple with the ShippedHazardousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedHazardousGoods

`func (o *VehicleProfileTruck) SetShippedHazardousGoods(v []ShippedHazardousGood)`

SetShippedHazardousGoods sets ShippedHazardousGoods field to given value.

### HasShippedHazardousGoods

`func (o *VehicleProfileTruck) HasShippedHazardousGoods() bool`

HasShippedHazardousGoods returns a boolean if a field has been set.

### GetExcludedCountries

`func (o *VehicleProfileTruck) GetExcludedCountries() []string`

GetExcludedCountries returns the ExcludedCountries field if non-nil, zero value otherwise.

### GetExcludedCountriesOk

`func (o *VehicleProfileTruck) GetExcludedCountriesOk() (*[]string, bool)`

GetExcludedCountriesOk returns a tuple with the ExcludedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCountries

`func (o *VehicleProfileTruck) SetExcludedCountries(v []string)`

SetExcludedCountries sets ExcludedCountries field to given value.

### HasExcludedCountries

`func (o *VehicleProfileTruck) HasExcludedCountries() bool`

HasExcludedCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


