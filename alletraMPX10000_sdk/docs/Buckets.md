# Buckets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketPolicy** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **bool** | Field to check if compression is  enabled or disabled | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Generation** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** | id of the bucket | [optional] 
**Name** | Pointer to **string** | Name of the bucket | [optional] 
**QuotaLimitMB** | Pointer to **int32** | Field to check the quotaLimit set in MB | [optional] 
**QuotaType** | Pointer to **string** | Field to check the quotaType set on the bucket | [optional] 
**SystemUid** | Pointer to **string** | systemUid where the bucket is created | [optional] 
**Tags** | Pointer to [**[]BucketTagsInner**](BucketTagsInner.md) | Tags for the bucket | [optional] 
**TenantId** | Pointer to **string** | tenantId of the customer | [optional] 
**Versioning** | Pointer to **string** | Field to check if versioning is enabled or disabled | [optional] 

## Methods

### NewBuckets

`func NewBuckets() *Buckets`

NewBuckets instantiates a new Buckets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketsWithDefaults

`func NewBucketsWithDefaults() *Buckets`

NewBucketsWithDefaults instantiates a new Buckets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketPolicy

`func (o *Buckets) GetBucketPolicy() string`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *Buckets) GetBucketPolicyOk() (*string, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *Buckets) SetBucketPolicy(v string)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *Buckets) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.

### GetCompression

`func (o *Buckets) GetCompression() bool`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *Buckets) GetCompressionOk() (*bool, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *Buckets) SetCompression(v bool)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *Buckets) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Buckets) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Buckets) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Buckets) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Buckets) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGeneration

`func (o *Buckets) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *Buckets) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *Buckets) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *Buckets) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetId

`func (o *Buckets) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Buckets) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Buckets) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Buckets) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Buckets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Buckets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Buckets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Buckets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuotaLimitMB

`func (o *Buckets) GetQuotaLimitMB() int32`

GetQuotaLimitMB returns the QuotaLimitMB field if non-nil, zero value otherwise.

### GetQuotaLimitMBOk

`func (o *Buckets) GetQuotaLimitMBOk() (*int32, bool)`

GetQuotaLimitMBOk returns a tuple with the QuotaLimitMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimitMB

`func (o *Buckets) SetQuotaLimitMB(v int32)`

SetQuotaLimitMB sets QuotaLimitMB field to given value.

### HasQuotaLimitMB

`func (o *Buckets) HasQuotaLimitMB() bool`

HasQuotaLimitMB returns a boolean if a field has been set.

### GetQuotaType

`func (o *Buckets) GetQuotaType() string`

GetQuotaType returns the QuotaType field if non-nil, zero value otherwise.

### GetQuotaTypeOk

`func (o *Buckets) GetQuotaTypeOk() (*string, bool)`

GetQuotaTypeOk returns a tuple with the QuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaType

`func (o *Buckets) SetQuotaType(v string)`

SetQuotaType sets QuotaType field to given value.

### HasQuotaType

`func (o *Buckets) HasQuotaType() bool`

HasQuotaType returns a boolean if a field has been set.

### GetSystemUid

`func (o *Buckets) GetSystemUid() string`

GetSystemUid returns the SystemUid field if non-nil, zero value otherwise.

### GetSystemUidOk

`func (o *Buckets) GetSystemUidOk() (*string, bool)`

GetSystemUidOk returns a tuple with the SystemUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUid

`func (o *Buckets) SetSystemUid(v string)`

SetSystemUid sets SystemUid field to given value.

### HasSystemUid

`func (o *Buckets) HasSystemUid() bool`

HasSystemUid returns a boolean if a field has been set.

### GetTags

`func (o *Buckets) GetTags() []BucketTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Buckets) GetTagsOk() (*[]BucketTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Buckets) SetTags(v []BucketTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Buckets) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Buckets) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Buckets) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTenantId

`func (o *Buckets) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Buckets) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Buckets) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Buckets) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetVersioning

`func (o *Buckets) GetVersioning() string`

GetVersioning returns the Versioning field if non-nil, zero value otherwise.

### GetVersioningOk

`func (o *Buckets) GetVersioningOk() (*string, bool)`

GetVersioningOk returns a tuple with the Versioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioning

`func (o *Buckets) SetVersioning(v string)`

SetVersioning sets Versioning field to given value.

### HasVersioning

`func (o *Buckets) HasVersioning() bool`

HasVersioning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


