# TaskLogMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**TimestampAt** | **time.Time** |  | 

## Methods

### NewTaskLogMessage

`func NewTaskLogMessage(message string, timestampAt time.Time, ) *TaskLogMessage`

NewTaskLogMessage instantiates a new TaskLogMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskLogMessageWithDefaults

`func NewTaskLogMessageWithDefaults() *TaskLogMessage`

NewTaskLogMessageWithDefaults instantiates a new TaskLogMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TaskLogMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskLogMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskLogMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTimestampAt

`func (o *TaskLogMessage) GetTimestampAt() time.Time`

GetTimestampAt returns the TimestampAt field if non-nil, zero value otherwise.

### GetTimestampAtOk

`func (o *TaskLogMessage) GetTimestampAtOk() (*time.Time, bool)`

GetTimestampAtOk returns a tuple with the TimestampAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampAt

`func (o *TaskLogMessage) SetTimestampAt(v time.Time)`

SetTimestampAt sets TimestampAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


