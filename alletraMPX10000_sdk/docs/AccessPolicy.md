# AccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generation** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** | policy ID | [optional] 
**Name** | Pointer to **string** | Name of the policy | [optional] 
**Statement** | Pointer to [**[]PolicyStatement**](PolicyStatement.md) |  | [optional] 
**SystemUid** | Pointer to **string** | Cluster ID | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | Version of the policy | [optional] 

## Methods

### NewAccessPolicy

`func NewAccessPolicy() *AccessPolicy`

NewAccessPolicy instantiates a new AccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyWithDefaults

`func NewAccessPolicyWithDefaults() *AccessPolicy`

NewAccessPolicyWithDefaults instantiates a new AccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneration

`func (o *AccessPolicy) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *AccessPolicy) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *AccessPolicy) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *AccessPolicy) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetId

`func (o *AccessPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccessPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatement

`func (o *AccessPolicy) GetStatement() []PolicyStatement`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *AccessPolicy) GetStatementOk() (*[]PolicyStatement, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *AccessPolicy) SetStatement(v []PolicyStatement)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *AccessPolicy) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetSystemUid

`func (o *AccessPolicy) GetSystemUid() string`

GetSystemUid returns the SystemUid field if non-nil, zero value otherwise.

### GetSystemUidOk

`func (o *AccessPolicy) GetSystemUidOk() (*string, bool)`

GetSystemUidOk returns a tuple with the SystemUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUid

`func (o *AccessPolicy) SetSystemUid(v string)`

SetSystemUid sets SystemUid field to given value.

### HasSystemUid

`func (o *AccessPolicy) HasSystemUid() bool`

HasSystemUid returns a boolean if a field has been set.

### GetTenantId

`func (o *AccessPolicy) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AccessPolicy) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AccessPolicy) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AccessPolicy) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetVersion

`func (o *AccessPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AccessPolicy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AccessPolicy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AccessPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


