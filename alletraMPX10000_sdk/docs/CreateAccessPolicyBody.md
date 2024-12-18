# CreateAccessPolicyBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Access Policy to be created | 
**Statement** | [**[]PolicyStatementInput**](PolicyStatementInput.md) |  | 
**Version** | **string** | access policy version | 

## Methods

### NewCreateAccessPolicyBody

`func NewCreateAccessPolicyBody(name string, statement []PolicyStatementInput, version string, ) *CreateAccessPolicyBody`

NewCreateAccessPolicyBody instantiates a new CreateAccessPolicyBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessPolicyBodyWithDefaults

`func NewCreateAccessPolicyBodyWithDefaults() *CreateAccessPolicyBody`

NewCreateAccessPolicyBodyWithDefaults instantiates a new CreateAccessPolicyBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAccessPolicyBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccessPolicyBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccessPolicyBody) SetName(v string)`

SetName sets Name field to given value.


### GetStatement

`func (o *CreateAccessPolicyBody) GetStatement() []PolicyStatementInput`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *CreateAccessPolicyBody) GetStatementOk() (*[]PolicyStatementInput, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *CreateAccessPolicyBody) SetStatement(v []PolicyStatementInput)`

SetStatement sets Statement field to given value.


### GetVersion

`func (o *CreateAccessPolicyBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateAccessPolicyBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateAccessPolicyBody) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


