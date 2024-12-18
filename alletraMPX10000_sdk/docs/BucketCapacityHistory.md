# BucketCapacityHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityData** | Pointer to [**NullableBucketsCapacityHistoricalMetricData**](BucketsCapacityHistoricalMetricData.md) |  | [optional] 
**EndTime** | Pointer to **NullableInt32** | end time of history data | [optional] 
**RequestUri** | Pointer to **NullableString** | requestUri for detailed storage object7 | [optional] 
**StartTime** | Pointer to **NullableInt32** | start time of history data | [optional] 

## Methods

### NewBucketCapacityHistory

`func NewBucketCapacityHistory() *BucketCapacityHistory`

NewBucketCapacityHistory instantiates a new BucketCapacityHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketCapacityHistoryWithDefaults

`func NewBucketCapacityHistoryWithDefaults() *BucketCapacityHistory`

NewBucketCapacityHistoryWithDefaults instantiates a new BucketCapacityHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityData

`func (o *BucketCapacityHistory) GetCapacityData() BucketsCapacityHistoricalMetricData`

GetCapacityData returns the CapacityData field if non-nil, zero value otherwise.

### GetCapacityDataOk

`func (o *BucketCapacityHistory) GetCapacityDataOk() (*BucketsCapacityHistoricalMetricData, bool)`

GetCapacityDataOk returns a tuple with the CapacityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityData

`func (o *BucketCapacityHistory) SetCapacityData(v BucketsCapacityHistoricalMetricData)`

SetCapacityData sets CapacityData field to given value.

### HasCapacityData

`func (o *BucketCapacityHistory) HasCapacityData() bool`

HasCapacityData returns a boolean if a field has been set.

### SetCapacityDataNil

`func (o *BucketCapacityHistory) SetCapacityDataNil(b bool)`

 SetCapacityDataNil sets the value for CapacityData to be an explicit nil

### UnsetCapacityData
`func (o *BucketCapacityHistory) UnsetCapacityData()`

UnsetCapacityData ensures that no value is present for CapacityData, not even an explicit nil
### GetEndTime

`func (o *BucketCapacityHistory) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BucketCapacityHistory) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BucketCapacityHistory) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *BucketCapacityHistory) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *BucketCapacityHistory) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *BucketCapacityHistory) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetRequestUri

`func (o *BucketCapacityHistory) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *BucketCapacityHistory) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *BucketCapacityHistory) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *BucketCapacityHistory) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### SetRequestUriNil

`func (o *BucketCapacityHistory) SetRequestUriNil(b bool)`

 SetRequestUriNil sets the value for RequestUri to be an explicit nil

### UnsetRequestUri
`func (o *BucketCapacityHistory) UnsetRequestUri()`

UnsetRequestUri ensures that no value is present for RequestUri, not even an explicit nil
### GetStartTime

`func (o *BucketCapacityHistory) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BucketCapacityHistory) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BucketCapacityHistory) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BucketCapacityHistory) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *BucketCapacityHistory) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BucketCapacityHistory) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


