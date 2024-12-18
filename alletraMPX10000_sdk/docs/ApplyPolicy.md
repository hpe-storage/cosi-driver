# ApplyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the user or group | 
**Policies** | **[]string** | list of policies to apply | 
**Type** | **string** | user or group | 

## Methods

### NewApplyPolicy

`func NewApplyPolicy(name string, policies []string, type_ string, ) *ApplyPolicy`

NewApplyPolicy instantiates a new ApplyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyPolicyWithDefaults

`func NewApplyPolicyWithDefaults() *ApplyPolicy`

NewApplyPolicyWithDefaults instantiates a new ApplyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplyPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplyPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplyPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetPolicies

`func (o *ApplyPolicy) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ApplyPolicy) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ApplyPolicy) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### GetType

`func (o *ApplyPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplyPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplyPolicy) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


