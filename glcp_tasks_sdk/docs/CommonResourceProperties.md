# CommonResourceProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | The customer application identifier | [optional] [readonly] 
**Generation** | Pointer to **int64** | A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date. | [optional] [readonly] 
**Id** | Pointer to **string** | An identifier for the resource, usually a UUID. | [optional] [readonly] 
**Name** | Pointer to **string** | A system specified name for the resource. | [optional] 
**ResourceUri** | Pointer to **string** | The &#39;self&#39; reference for this resource. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of resource. | [optional] [readonly] 

## Methods

### NewCommonResourceProperties

`func NewCommonResourceProperties() *CommonResourceProperties`

NewCommonResourceProperties instantiates a new CommonResourceProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResourcePropertiesWithDefaults

`func NewCommonResourcePropertiesWithDefaults() *CommonResourceProperties`

NewCommonResourcePropertiesWithDefaults instantiates a new CommonResourceProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CommonResourceProperties) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CommonResourceProperties) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CommonResourceProperties) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CommonResourceProperties) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetGeneration

`func (o *CommonResourceProperties) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *CommonResourceProperties) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *CommonResourceProperties) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *CommonResourceProperties) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetId

`func (o *CommonResourceProperties) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonResourceProperties) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonResourceProperties) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommonResourceProperties) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CommonResourceProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonResourceProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonResourceProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonResourceProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceUri

`func (o *CommonResourceProperties) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *CommonResourceProperties) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *CommonResourceProperties) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.

### HasResourceUri

`func (o *CommonResourceProperties) HasResourceUri() bool`

HasResourceUri returns a boolean if a field has been set.

### GetType

`func (o *CommonResourceProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommonResourceProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommonResourceProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommonResourceProperties) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


