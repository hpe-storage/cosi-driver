# BucketsCapacityMetricSeriesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimestampMs** | Pointer to **NullableInt64** | epoch timestamp | [optional] 
**UsageMiB** | Pointer to **NullableFloat32** | average logical usage value at particular timestamp | [optional] 

## Methods

### NewBucketsCapacityMetricSeriesData

`func NewBucketsCapacityMetricSeriesData() *BucketsCapacityMetricSeriesData`

NewBucketsCapacityMetricSeriesData instantiates a new BucketsCapacityMetricSeriesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketsCapacityMetricSeriesDataWithDefaults

`func NewBucketsCapacityMetricSeriesDataWithDefaults() *BucketsCapacityMetricSeriesData`

NewBucketsCapacityMetricSeriesDataWithDefaults instantiates a new BucketsCapacityMetricSeriesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestampMs

`func (o *BucketsCapacityMetricSeriesData) GetTimestampMs() int64`

GetTimestampMs returns the TimestampMs field if non-nil, zero value otherwise.

### GetTimestampMsOk

`func (o *BucketsCapacityMetricSeriesData) GetTimestampMsOk() (*int64, bool)`

GetTimestampMsOk returns a tuple with the TimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMs

`func (o *BucketsCapacityMetricSeriesData) SetTimestampMs(v int64)`

SetTimestampMs sets TimestampMs field to given value.

### HasTimestampMs

`func (o *BucketsCapacityMetricSeriesData) HasTimestampMs() bool`

HasTimestampMs returns a boolean if a field has been set.

### SetTimestampMsNil

`func (o *BucketsCapacityMetricSeriesData) SetTimestampMsNil(b bool)`

 SetTimestampMsNil sets the value for TimestampMs to be an explicit nil

### UnsetTimestampMs
`func (o *BucketsCapacityMetricSeriesData) UnsetTimestampMs()`

UnsetTimestampMs ensures that no value is present for TimestampMs, not even an explicit nil
### GetUsageMiB

`func (o *BucketsCapacityMetricSeriesData) GetUsageMiB() float32`

GetUsageMiB returns the UsageMiB field if non-nil, zero value otherwise.

### GetUsageMiBOk

`func (o *BucketsCapacityMetricSeriesData) GetUsageMiBOk() (*float32, bool)`

GetUsageMiBOk returns a tuple with the UsageMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMiB

`func (o *BucketsCapacityMetricSeriesData) SetUsageMiB(v float32)`

SetUsageMiB sets UsageMiB field to given value.

### HasUsageMiB

`func (o *BucketsCapacityMetricSeriesData) HasUsageMiB() bool`

HasUsageMiB returns a boolean if a field has been set.

### SetUsageMiBNil

`func (o *BucketsCapacityMetricSeriesData) SetUsageMiBNil(b bool)`

 SetUsageMiBNil sets the value for UsageMiB to be an explicit nil

### UnsetUsageMiB
`func (o *BucketsCapacityMetricSeriesData) UnsetUsageMiB()`

UnsetUsageMiB ensures that no value is present for UsageMiB, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


