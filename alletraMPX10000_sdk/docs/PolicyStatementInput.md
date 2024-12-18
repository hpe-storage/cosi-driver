# PolicyStatementInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **[]string** |  | 
**Condition** | Pointer to [**map[string]map[string]PolicyConditionValue**](map.md) |  | [optional] 
**Effect** | **string** |  | 
**Resource** | [**PolicyStatementInputResource**](PolicyStatementInputResource.md) |  | 
**Sid** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyStatementInput

`func NewPolicyStatementInput(action []string, effect string, resource PolicyStatementInputResource, ) *PolicyStatementInput`

NewPolicyStatementInput instantiates a new PolicyStatementInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyStatementInputWithDefaults

`func NewPolicyStatementInputWithDefaults() *PolicyStatementInput`

NewPolicyStatementInputWithDefaults instantiates a new PolicyStatementInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PolicyStatementInput) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyStatementInput) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyStatementInput) SetAction(v []string)`

SetAction sets Action field to given value.


### GetCondition

`func (o *PolicyStatementInput) GetCondition() map[string]map[string]PolicyConditionValue`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *PolicyStatementInput) GetConditionOk() (*map[string]map[string]PolicyConditionValue, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *PolicyStatementInput) SetCondition(v map[string]map[string]PolicyConditionValue)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *PolicyStatementInput) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEffect

`func (o *PolicyStatementInput) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *PolicyStatementInput) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *PolicyStatementInput) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetResource

`func (o *PolicyStatementInput) GetResource() PolicyStatementInputResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PolicyStatementInput) GetResourceOk() (*PolicyStatementInputResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PolicyStatementInput) SetResource(v PolicyStatementInputResource)`

SetResource sets Resource field to given value.


### GetSid

`func (o *PolicyStatementInput) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *PolicyStatementInput) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *PolicyStatementInput) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *PolicyStatementInput) HasSid() bool`

HasSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


