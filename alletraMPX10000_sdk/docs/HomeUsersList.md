# HomeUsersList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]UserListDetails**](UserListDetails.md) |  | [optional] 
**PageLimit** | Pointer to **NullableInt32** | page limit | [optional] 
**PageOffset** | Pointer to **NullableInt32** | page offset | [optional] 
**Total** | Pointer to **NullableInt32** | Number of users lists | [optional] 

## Methods

### NewHomeUsersList

`func NewHomeUsersList() *HomeUsersList`

NewHomeUsersList instantiates a new HomeUsersList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHomeUsersListWithDefaults

`func NewHomeUsersListWithDefaults() *HomeUsersList`

NewHomeUsersListWithDefaults instantiates a new HomeUsersList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *HomeUsersList) GetItems() []UserListDetails`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HomeUsersList) GetItemsOk() (*[]UserListDetails, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HomeUsersList) SetItems(v []UserListDetails)`

SetItems sets Items field to given value.

### HasItems

`func (o *HomeUsersList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *HomeUsersList) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *HomeUsersList) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetPageLimit

`func (o *HomeUsersList) GetPageLimit() int32`

GetPageLimit returns the PageLimit field if non-nil, zero value otherwise.

### GetPageLimitOk

`func (o *HomeUsersList) GetPageLimitOk() (*int32, bool)`

GetPageLimitOk returns a tuple with the PageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLimit

`func (o *HomeUsersList) SetPageLimit(v int32)`

SetPageLimit sets PageLimit field to given value.

### HasPageLimit

`func (o *HomeUsersList) HasPageLimit() bool`

HasPageLimit returns a boolean if a field has been set.

### SetPageLimitNil

`func (o *HomeUsersList) SetPageLimitNil(b bool)`

 SetPageLimitNil sets the value for PageLimit to be an explicit nil

### UnsetPageLimit
`func (o *HomeUsersList) UnsetPageLimit()`

UnsetPageLimit ensures that no value is present for PageLimit, not even an explicit nil
### GetPageOffset

`func (o *HomeUsersList) GetPageOffset() int32`

GetPageOffset returns the PageOffset field if non-nil, zero value otherwise.

### GetPageOffsetOk

`func (o *HomeUsersList) GetPageOffsetOk() (*int32, bool)`

GetPageOffsetOk returns a tuple with the PageOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageOffset

`func (o *HomeUsersList) SetPageOffset(v int32)`

SetPageOffset sets PageOffset field to given value.

### HasPageOffset

`func (o *HomeUsersList) HasPageOffset() bool`

HasPageOffset returns a boolean if a field has been set.

### SetPageOffsetNil

`func (o *HomeUsersList) SetPageOffsetNil(b bool)`

 SetPageOffsetNil sets the value for PageOffset to be an explicit nil

### UnsetPageOffset
`func (o *HomeUsersList) UnsetPageOffset()`

UnsetPageOffset ensures that no value is present for PageOffset, not even an explicit nil
### GetTotal

`func (o *HomeUsersList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HomeUsersList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HomeUsersList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HomeUsersList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *HomeUsersList) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *HomeUsersList) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


