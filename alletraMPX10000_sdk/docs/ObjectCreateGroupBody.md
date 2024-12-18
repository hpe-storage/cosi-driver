# ObjectCreateGroupBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the group to be created | 
**Users** | Pointer to **[]string** | Users name list to add in the group | [optional] 

## Methods

### NewObjectCreateGroupBody

`func NewObjectCreateGroupBody(name string, ) *ObjectCreateGroupBody`

NewObjectCreateGroupBody instantiates a new ObjectCreateGroupBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectCreateGroupBodyWithDefaults

`func NewObjectCreateGroupBodyWithDefaults() *ObjectCreateGroupBody`

NewObjectCreateGroupBodyWithDefaults instantiates a new ObjectCreateGroupBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ObjectCreateGroupBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectCreateGroupBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectCreateGroupBody) SetName(v string)`

SetName sets Name field to given value.


### GetUsers

`func (o *ObjectCreateGroupBody) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ObjectCreateGroupBody) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ObjectCreateGroupBody) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ObjectCreateGroupBody) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


