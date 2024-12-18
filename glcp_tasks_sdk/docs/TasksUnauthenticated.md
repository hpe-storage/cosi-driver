# TasksUnauthenticated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **interface{}** | Possible values: UNAUTHENTICATED | [optional] 
**Error** | **string** | A user friendly error message | 
**TraceId** | **string** | A unique identifier for the request | 

## Methods

### NewTasksUnauthenticated

`func NewTasksUnauthenticated(error_ string, traceId string, ) *TasksUnauthenticated`

NewTasksUnauthenticated instantiates a new TasksUnauthenticated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksUnauthenticatedWithDefaults

`func NewTasksUnauthenticatedWithDefaults() *TasksUnauthenticated`

NewTasksUnauthenticatedWithDefaults instantiates a new TasksUnauthenticated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *TasksUnauthenticated) GetErrorCode() interface{}`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TasksUnauthenticated) GetErrorCodeOk() (*interface{}, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TasksUnauthenticated) SetErrorCode(v interface{})`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TasksUnauthenticated) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *TasksUnauthenticated) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *TasksUnauthenticated) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetError

`func (o *TasksUnauthenticated) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TasksUnauthenticated) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TasksUnauthenticated) SetError(v string)`

SetError sets Error field to given value.


### GetTraceId

`func (o *TasksUnauthenticated) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TasksUnauthenticated) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TasksUnauthenticated) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


