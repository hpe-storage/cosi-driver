# BucketsCapacityHistoricalMetricData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketId** | Pointer to **NullableString** | bucketId | [optional] 
**CustomerId** | Pointer to **NullableString** | customerId | [optional] 
**Items** | Pointer to [**[]BucketsCapacityMetricSeriesData**](BucketsCapacityMetricSeriesData.md) |  | [optional] 
**SystemId** | Pointer to **NullableString** | systemId | [optional] 
**Total** | Pointer to **NullableInt32** | count of series data | [optional] 

## Methods

### NewBucketsCapacityHistoricalMetricData

`func NewBucketsCapacityHistoricalMetricData() *BucketsCapacityHistoricalMetricData`

NewBucketsCapacityHistoricalMetricData instantiates a new BucketsCapacityHistoricalMetricData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketsCapacityHistoricalMetricDataWithDefaults

`func NewBucketsCapacityHistoricalMetricDataWithDefaults() *BucketsCapacityHistoricalMetricData`

NewBucketsCapacityHistoricalMetricDataWithDefaults instantiates a new BucketsCapacityHistoricalMetricData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketId

`func (o *BucketsCapacityHistoricalMetricData) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *BucketsCapacityHistoricalMetricData) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *BucketsCapacityHistoricalMetricData) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *BucketsCapacityHistoricalMetricData) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### SetBucketIdNil

`func (o *BucketsCapacityHistoricalMetricData) SetBucketIdNil(b bool)`

 SetBucketIdNil sets the value for BucketId to be an explicit nil

### UnsetBucketId
`func (o *BucketsCapacityHistoricalMetricData) UnsetBucketId()`

UnsetBucketId ensures that no value is present for BucketId, not even an explicit nil
### GetCustomerId

`func (o *BucketsCapacityHistoricalMetricData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *BucketsCapacityHistoricalMetricData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *BucketsCapacityHistoricalMetricData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *BucketsCapacityHistoricalMetricData) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *BucketsCapacityHistoricalMetricData) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *BucketsCapacityHistoricalMetricData) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetItems

`func (o *BucketsCapacityHistoricalMetricData) GetItems() []BucketsCapacityMetricSeriesData`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BucketsCapacityHistoricalMetricData) GetItemsOk() (*[]BucketsCapacityMetricSeriesData, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BucketsCapacityHistoricalMetricData) SetItems(v []BucketsCapacityMetricSeriesData)`

SetItems sets Items field to given value.

### HasItems

`func (o *BucketsCapacityHistoricalMetricData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *BucketsCapacityHistoricalMetricData) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *BucketsCapacityHistoricalMetricData) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetSystemId

`func (o *BucketsCapacityHistoricalMetricData) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *BucketsCapacityHistoricalMetricData) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *BucketsCapacityHistoricalMetricData) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *BucketsCapacityHistoricalMetricData) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### SetSystemIdNil

`func (o *BucketsCapacityHistoricalMetricData) SetSystemIdNil(b bool)`

 SetSystemIdNil sets the value for SystemId to be an explicit nil

### UnsetSystemId
`func (o *BucketsCapacityHistoricalMetricData) UnsetSystemId()`

UnsetSystemId ensures that no value is present for SystemId, not even an explicit nil
### GetTotal

`func (o *BucketsCapacityHistoricalMetricData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BucketsCapacityHistoricalMetricData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BucketsCapacityHistoricalMetricData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BucketsCapacityHistoricalMetricData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *BucketsCapacityHistoricalMetricData) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *BucketsCapacityHistoricalMetricData) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


