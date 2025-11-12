# Paging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total number of items in the requested collection. | [optional] [readonly] 
**TotalPages** | Pointer to **int32** | The total number of pages in the requested collection. | [optional] [readonly] 
**Page** | Pointer to **int32** | The current page of items. | [optional] [readonly] 
**NextPage** | Pointer to **NullableInt32** | The next available page of items. | [optional] [readonly] 
**PreviousPage** | Pointer to **NullableInt32** | The previous page of items. | [optional] [readonly] 
**ItemsPerPage** | Pointer to **int32** | The number of items per page, set internally or by the request. | [optional] [readonly] [default to 20]

## Methods

### NewPaging

`func NewPaging() *Paging`

NewPaging instantiates a new Paging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingWithDefaults

`func NewPagingWithDefaults() *Paging`

NewPagingWithDefaults instantiates a new Paging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *Paging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Paging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Paging) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Paging) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalPages

`func (o *Paging) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *Paging) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *Paging) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *Paging) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetPage

`func (o *Paging) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Paging) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Paging) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *Paging) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetNextPage

`func (o *Paging) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *Paging) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *Paging) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *Paging) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *Paging) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *Paging) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil
### GetPreviousPage

`func (o *Paging) GetPreviousPage() int32`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *Paging) GetPreviousPageOk() (*int32, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *Paging) SetPreviousPage(v int32)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *Paging) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.

### SetPreviousPageNil

`func (o *Paging) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *Paging) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetItemsPerPage

`func (o *Paging) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *Paging) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *Paging) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *Paging) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


