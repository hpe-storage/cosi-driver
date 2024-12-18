# UpdateAccessPolicyBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | [**[]PolicyStatementInput**](PolicyStatementInput.md) |  | 
**Version** | **string** | access policy version | 

## Methods

### NewUpdateAccessPolicyBody

`func NewUpdateAccessPolicyBody(statement []PolicyStatementInput, version string, ) *UpdateAccessPolicyBody`

NewUpdateAccessPolicyBody instantiates a new UpdateAccessPolicyBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccessPolicyBodyWithDefaults

`func NewUpdateAccessPolicyBodyWithDefaults() *UpdateAccessPolicyBody`

NewUpdateAccessPolicyBodyWithDefaults instantiates a new UpdateAccessPolicyBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *UpdateAccessPolicyBody) GetStatement() []PolicyStatementInput`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *UpdateAccessPolicyBody) GetStatementOk() (*[]PolicyStatementInput, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *UpdateAccessPolicyBody) SetStatement(v []PolicyStatementInput)`

SetStatement sets Statement field to given value.


### GetVersion

`func (o *UpdateAccessPolicyBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateAccessPolicyBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateAccessPolicyBody) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


