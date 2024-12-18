# EditBucketBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compression** | Pointer to **string** | Field to enable or disable compression for the bucket | [optional] 
**Tags** | Pointer to [**[]BucketTagsInner**](BucketTagsInner.md) | Tags for the bucket | [optional] 

## Methods

### NewEditBucketBody

`func NewEditBucketBody() *EditBucketBody`

NewEditBucketBody instantiates a new EditBucketBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditBucketBodyWithDefaults

`func NewEditBucketBodyWithDefaults() *EditBucketBody`

NewEditBucketBodyWithDefaults instantiates a new EditBucketBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompression

`func (o *EditBucketBody) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *EditBucketBody) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *EditBucketBody) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *EditBucketBody) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetTags

`func (o *EditBucketBody) GetTags() []BucketTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EditBucketBody) GetTagsOk() (*[]BucketTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EditBucketBody) SetTags(v []BucketTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EditBucketBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *EditBucketBody) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *EditBucketBody) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


