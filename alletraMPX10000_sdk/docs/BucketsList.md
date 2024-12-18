# BucketsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Buckets**](Buckets.md) |  | [optional] 
**PageLimit** | Pointer to **NullableInt32** | page limit | [optional] 
**PageOffset** | Pointer to **NullableInt32** | page offset | [optional] 
**Total** | Pointer to **NullableInt32** | Number of recommendation lists | [optional] 

## Methods

### NewBucketsList

`func NewBucketsList() *BucketsList`

NewBucketsList instantiates a new BucketsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketsListWithDefaults

`func NewBucketsListWithDefaults() *BucketsList`

NewBucketsListWithDefaults instantiates a new BucketsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BucketsList) GetItems() []Buckets`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BucketsList) GetItemsOk() (*[]Buckets, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BucketsList) SetItems(v []Buckets)`

SetItems sets Items field to given value.

### HasItems

`func (o *BucketsList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *BucketsList) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *BucketsList) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetPageLimit

`func (o *BucketsList) GetPageLimit() int32`

GetPageLimit returns the PageLimit field if non-nil, zero value otherwise.

### GetPageLimitOk

`func (o *BucketsList) GetPageLimitOk() (*int32, bool)`

GetPageLimitOk returns a tuple with the PageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLimit

`func (o *BucketsList) SetPageLimit(v int32)`

SetPageLimit sets PageLimit field to given value.

### HasPageLimit

`func (o *BucketsList) HasPageLimit() bool`

HasPageLimit returns a boolean if a field has been set.

### SetPageLimitNil

`func (o *BucketsList) SetPageLimitNil(b bool)`

 SetPageLimitNil sets the value for PageLimit to be an explicit nil

### UnsetPageLimit
`func (o *BucketsList) UnsetPageLimit()`

UnsetPageLimit ensures that no value is present for PageLimit, not even an explicit nil
### GetPageOffset

`func (o *BucketsList) GetPageOffset() int32`

GetPageOffset returns the PageOffset field if non-nil, zero value otherwise.

### GetPageOffsetOk

`func (o *BucketsList) GetPageOffsetOk() (*int32, bool)`

GetPageOffsetOk returns a tuple with the PageOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageOffset

`func (o *BucketsList) SetPageOffset(v int32)`

SetPageOffset sets PageOffset field to given value.

### HasPageOffset

`func (o *BucketsList) HasPageOffset() bool`

HasPageOffset returns a boolean if a field has been set.

### SetPageOffsetNil

`func (o *BucketsList) SetPageOffsetNil(b bool)`

 SetPageOffsetNil sets the value for PageOffset to be an explicit nil

### UnsetPageOffset
`func (o *BucketsList) UnsetPageOffset()`

UnsetPageOffset ensures that no value is present for PageOffset, not even an explicit nil
### GetTotal

`func (o *BucketsList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BucketsList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BucketsList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BucketsList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *BucketsList) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *BucketsList) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


