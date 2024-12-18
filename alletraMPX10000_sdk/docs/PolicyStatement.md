# PolicyStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **[]string** |  | [optional] 
**Condition** | Pointer to [**map[string]map[string]PolicyConditionValue**](map.md) |  | [optional] 
**Effect** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **[]string** |  | [optional] 
**Sid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPolicyStatement

`func NewPolicyStatement() *PolicyStatement`

NewPolicyStatement instantiates a new PolicyStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyStatementWithDefaults

`func NewPolicyStatementWithDefaults() *PolicyStatement`

NewPolicyStatementWithDefaults instantiates a new PolicyStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PolicyStatement) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyStatement) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyStatement) SetAction(v []string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PolicyStatement) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCondition

`func (o *PolicyStatement) GetCondition() map[string]map[string]PolicyConditionValue`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *PolicyStatement) GetConditionOk() (*map[string]map[string]PolicyConditionValue, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *PolicyStatement) SetCondition(v map[string]map[string]PolicyConditionValue)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *PolicyStatement) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *PolicyStatement) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *PolicyStatement) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetEffect

`func (o *PolicyStatement) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *PolicyStatement) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *PolicyStatement) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *PolicyStatement) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetResource

`func (o *PolicyStatement) GetResource() []string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PolicyStatement) GetResourceOk() (*[]string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PolicyStatement) SetResource(v []string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PolicyStatement) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *PolicyStatement) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *PolicyStatement) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetSid

`func (o *PolicyStatement) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *PolicyStatement) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *PolicyStatement) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *PolicyStatement) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *PolicyStatement) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *PolicyStatement) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


