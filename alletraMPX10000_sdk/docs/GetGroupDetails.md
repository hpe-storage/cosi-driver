# GetGroupDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]GroupItem**](GroupItem.md) |  | [optional] 
**PageLimit** | Pointer to **NullableInt64** |  | [optional] 
**PageOffset** | Pointer to **NullableInt64** |  | [optional] 
**Total** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewGetGroupDetails

`func NewGetGroupDetails() *GetGroupDetails`

NewGetGroupDetails instantiates a new GetGroupDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupDetailsWithDefaults

`func NewGetGroupDetailsWithDefaults() *GetGroupDetails`

NewGetGroupDetailsWithDefaults instantiates a new GetGroupDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetGroupDetails) GetItems() []GroupItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetGroupDetails) GetItemsOk() (*[]GroupItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetGroupDetails) SetItems(v []GroupItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetGroupDetails) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *GetGroupDetails) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *GetGroupDetails) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetPageLimit

`func (o *GetGroupDetails) GetPageLimit() int64`

GetPageLimit returns the PageLimit field if non-nil, zero value otherwise.

### GetPageLimitOk

`func (o *GetGroupDetails) GetPageLimitOk() (*int64, bool)`

GetPageLimitOk returns a tuple with the PageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLimit

`func (o *GetGroupDetails) SetPageLimit(v int64)`

SetPageLimit sets PageLimit field to given value.

### HasPageLimit

`func (o *GetGroupDetails) HasPageLimit() bool`

HasPageLimit returns a boolean if a field has been set.

### SetPageLimitNil

`func (o *GetGroupDetails) SetPageLimitNil(b bool)`

 SetPageLimitNil sets the value for PageLimit to be an explicit nil

### UnsetPageLimit
`func (o *GetGroupDetails) UnsetPageLimit()`

UnsetPageLimit ensures that no value is present for PageLimit, not even an explicit nil
### GetPageOffset

`func (o *GetGroupDetails) GetPageOffset() int64`

GetPageOffset returns the PageOffset field if non-nil, zero value otherwise.

### GetPageOffsetOk

`func (o *GetGroupDetails) GetPageOffsetOk() (*int64, bool)`

GetPageOffsetOk returns a tuple with the PageOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageOffset

`func (o *GetGroupDetails) SetPageOffset(v int64)`

SetPageOffset sets PageOffset field to given value.

### HasPageOffset

`func (o *GetGroupDetails) HasPageOffset() bool`

HasPageOffset returns a boolean if a field has been set.

### SetPageOffsetNil

`func (o *GetGroupDetails) SetPageOffsetNil(b bool)`

 SetPageOffsetNil sets the value for PageOffset to be an explicit nil

### UnsetPageOffset
`func (o *GetGroupDetails) UnsetPageOffset()`

UnsetPageOffset ensures that no value is present for PageOffset, not even an explicit nil
### GetTotal

`func (o *GetGroupDetails) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetGroupDetails) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetGroupDetails) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetGroupDetails) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GetGroupDetails) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GetGroupDetails) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


