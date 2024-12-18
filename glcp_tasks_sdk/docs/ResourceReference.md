# ResourceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [readonly] 
**ResourceUri** | **string** |  | [readonly] 
**Type** | **string** |  | [readonly] 

## Methods

### NewResourceReference

`func NewResourceReference(resourceUri string, type_ string, ) *ResourceReference`

NewResourceReference instantiates a new ResourceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceReferenceWithDefaults

`func NewResourceReferenceWithDefaults() *ResourceReference`

NewResourceReferenceWithDefaults instantiates a new ResourceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceUri

`func (o *ResourceReference) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *ResourceReference) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *ResourceReference) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.


### GetType

`func (o *ResourceReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceReference) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


