# TaskResponseUi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Task Message task message. | [optional] 
**Status** | Pointer to **string** | Task Status Status of the task. | [optional] 
**TaskUri** | **string** | Task URI which can be used to monitor the status of the operation. | 

## Methods

### NewTaskResponseUi

`func NewTaskResponseUi(taskUri string, ) *TaskResponseUi`

NewTaskResponseUi instantiates a new TaskResponseUi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResponseUiWithDefaults

`func NewTaskResponseUiWithDefaults() *TaskResponseUi`

NewTaskResponseUiWithDefaults instantiates a new TaskResponseUi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TaskResponseUi) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskResponseUi) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskResponseUi) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TaskResponseUi) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *TaskResponseUi) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskResponseUi) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskResponseUi) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskResponseUi) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskUri

`func (o *TaskResponseUi) GetTaskUri() string`

GetTaskUri returns the TaskUri field if non-nil, zero value otherwise.

### GetTaskUriOk

`func (o *TaskResponseUi) GetTaskUriOk() (*string, bool)`

GetTaskUriOk returns a tuple with the TaskUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUri

`func (o *TaskResponseUi) SetTaskUri(v string)`

SetTaskUri sets TaskUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


