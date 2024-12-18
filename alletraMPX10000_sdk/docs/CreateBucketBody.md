# CreateBucketBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compression** | Pointer to **string** | Field to enable or disable compression for the bucket | [optional] 
**Name** | **string** | Name of the bucket to be created | 
**Tags** | Pointer to [**[]BucketTagsInner**](BucketTagsInner.md) | Tags for the bucket | [optional] 

## Methods

### NewCreateBucketBody

`func NewCreateBucketBody(name string, ) *CreateBucketBody`

NewCreateBucketBody instantiates a new CreateBucketBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBucketBodyWithDefaults

`func NewCreateBucketBodyWithDefaults() *CreateBucketBody`

NewCreateBucketBodyWithDefaults instantiates a new CreateBucketBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompression

`func (o *CreateBucketBody) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *CreateBucketBody) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *CreateBucketBody) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *CreateBucketBody) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetName

`func (o *CreateBucketBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBucketBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBucketBody) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *CreateBucketBody) GetTags() []BucketTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateBucketBody) GetTagsOk() (*[]BucketTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateBucketBody) SetTags(v []BucketTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateBucketBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateBucketBody) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateBucketBody) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


