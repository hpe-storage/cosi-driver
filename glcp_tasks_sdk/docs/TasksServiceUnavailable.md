# TasksServiceUnavailable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **interface{}** | Possible values: SERVICE_UNAVAILABLE | [optional] 
**Error** | **string** | A user friendly error message | 
**TraceId** | **string** | A unique identifier for the request | 

## Methods

### NewTasksServiceUnavailable

`func NewTasksServiceUnavailable(error_ string, traceId string, ) *TasksServiceUnavailable`

NewTasksServiceUnavailable instantiates a new TasksServiceUnavailable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksServiceUnavailableWithDefaults

`func NewTasksServiceUnavailableWithDefaults() *TasksServiceUnavailable`

NewTasksServiceUnavailableWithDefaults instantiates a new TasksServiceUnavailable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *TasksServiceUnavailable) GetErrorCode() interface{}`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TasksServiceUnavailable) GetErrorCodeOk() (*interface{}, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TasksServiceUnavailable) SetErrorCode(v interface{})`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TasksServiceUnavailable) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *TasksServiceUnavailable) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *TasksServiceUnavailable) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetError

`func (o *TasksServiceUnavailable) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TasksServiceUnavailable) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TasksServiceUnavailable) SetError(v string)`

SetError sets Error field to given value.


### GetTraceId

`func (o *TasksServiceUnavailable) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TasksServiceUnavailable) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TasksServiceUnavailable) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


