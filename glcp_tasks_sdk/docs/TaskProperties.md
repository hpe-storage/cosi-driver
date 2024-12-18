# TaskProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to [**TaskConsoleReference**](TaskConsoleReference.md) | A link to be displayed in the Tasks UI. This can be used when a task is paused to take the user to the console UI page with information on how to unpause the task, or for more general information when the task is in other states. | [optional] 
**AssociatedResources** | Pointer to [**[]ResourceReference**](ResourceReference.md) | Resources that are associated with the task. These may be created by the task or other resources that are involved in the task. | [optional] 
**ChildTasks** | Pointer to [**[]ResourceReference**](ResourceReference.md) | A list of sub-tasks that were initiated by this task. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time this task was created. | [optional] 
**DisplayName** | Pointer to **string** | The displayed name for the task. | [optional] 
**EndedAt** | Pointer to **NullableTime** | The time this task completed. | [optional] 
**Error** | Pointer to [**NullableErrorResponse**](ErrorResponse.md) | The error response status of the operation. | [optional] 
**EstimatedRunningDurationMinutes** | Pointer to **int32** | An estimate of how long the task will run before completing. | [optional] 
**Groups** | Pointer to [**[]GroupIDNamePair**](GroupIDNamePair.md) | A list of groups associated with this task. | [optional] 
**HealthStatus** | Pointer to **string** | The health status indicates if any errors or problems have been encountered during the processing of the task.  Expected values are OK, ERROR, WARNING, UNKNOWN, and UNSPECIFIED.  | [optional] 
**LogMessages** | Pointer to [**[]TaskLogMessage**](TaskLogMessage.md) | Time stamped messages that record the progress of the task. | [optional] 
**ParentTask** | Pointer to [**NullableResourceReferenceWithId**](ResourceReferenceWithId.md) | The parent is the task that initiated this sub-task. If this task is not a sub-task this will be null. | [optional] 
**ProgressPercent** | Pointer to **int32** | A percentage representation of progress to completion. | [optional] 
**Recommendations** | Pointer to [**[]TaskRecommendations**](TaskRecommendations.md) | Recommendations on how to fix failing tasks. | [optional] 
**RootTask** | Pointer to [**NullableResourceReferenceWithId**](ResourceReferenceWithId.md) | The root of the tree of tasks. If this task is not part of a tree this will be null. | [optional] 
**Services** | Pointer to **[]string** | List of services this task belongs to, can be used to filter to specific services in the UI | [optional] 
**SourceResource** | Pointer to [**NullableResourceReference**](ResourceReference.md) | The resource that was used to initiate the task. | [optional] 
**StartedAt** | Pointer to **NullableTime** | The time this task was started. | [optional] 
**State** | Pointer to **string** | A message to indicate the current state of the task, for example the current step in a workflow. Expected values are INITIALIZED, RUNNING, FAILED, SUCCEEDED, TIMEDOUT, PAUSED, and UNSPECIFIED.  | [optional] 
**SubtreeTaskCount** | Pointer to **int32** | The count of the number of child Tasks below this one, recursively. | [optional] 
**SuggestedPollingIntervalSeconds** | Pointer to **int32** | This attribute suggests a suitable interval to use when polling for progress. Where specified this will be based on the frequency with which the task is likely to be updated. | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The time this task was last updated. | [optional] 
**UserId** | Pointer to **string** | The ID or email address of the user that initiated the task. | [optional] 

## Methods

### NewTaskProperties

`func NewTaskProperties() *TaskProperties`

NewTaskProperties instantiates a new TaskProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPropertiesWithDefaults

`func NewTaskPropertiesWithDefaults() *TaskProperties`

NewTaskPropertiesWithDefaults instantiates a new TaskProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *TaskProperties) GetAdditionalDetails() TaskConsoleReference`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *TaskProperties) GetAdditionalDetailsOk() (*TaskConsoleReference, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *TaskProperties) SetAdditionalDetails(v TaskConsoleReference)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *TaskProperties) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetAssociatedResources

`func (o *TaskProperties) GetAssociatedResources() []ResourceReference`

GetAssociatedResources returns the AssociatedResources field if non-nil, zero value otherwise.

### GetAssociatedResourcesOk

`func (o *TaskProperties) GetAssociatedResourcesOk() (*[]ResourceReference, bool)`

GetAssociatedResourcesOk returns a tuple with the AssociatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedResources

`func (o *TaskProperties) SetAssociatedResources(v []ResourceReference)`

SetAssociatedResources sets AssociatedResources field to given value.

### HasAssociatedResources

`func (o *TaskProperties) HasAssociatedResources() bool`

HasAssociatedResources returns a boolean if a field has been set.

### GetChildTasks

`func (o *TaskProperties) GetChildTasks() []ResourceReference`

GetChildTasks returns the ChildTasks field if non-nil, zero value otherwise.

### GetChildTasksOk

`func (o *TaskProperties) GetChildTasksOk() (*[]ResourceReference, bool)`

GetChildTasksOk returns a tuple with the ChildTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTasks

`func (o *TaskProperties) SetChildTasks(v []ResourceReference)`

SetChildTasks sets ChildTasks field to given value.

### HasChildTasks

`func (o *TaskProperties) HasChildTasks() bool`

HasChildTasks returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskProperties) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskProperties) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskProperties) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskProperties) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *TaskProperties) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TaskProperties) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TaskProperties) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TaskProperties) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndedAt

`func (o *TaskProperties) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *TaskProperties) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *TaskProperties) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *TaskProperties) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *TaskProperties) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *TaskProperties) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetError

`func (o *TaskProperties) GetError() ErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TaskProperties) GetErrorOk() (*ErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TaskProperties) SetError(v ErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *TaskProperties) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *TaskProperties) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *TaskProperties) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetEstimatedRunningDurationMinutes

`func (o *TaskProperties) GetEstimatedRunningDurationMinutes() int32`

GetEstimatedRunningDurationMinutes returns the EstimatedRunningDurationMinutes field if non-nil, zero value otherwise.

### GetEstimatedRunningDurationMinutesOk

`func (o *TaskProperties) GetEstimatedRunningDurationMinutesOk() (*int32, bool)`

GetEstimatedRunningDurationMinutesOk returns a tuple with the EstimatedRunningDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRunningDurationMinutes

`func (o *TaskProperties) SetEstimatedRunningDurationMinutes(v int32)`

SetEstimatedRunningDurationMinutes sets EstimatedRunningDurationMinutes field to given value.

### HasEstimatedRunningDurationMinutes

`func (o *TaskProperties) HasEstimatedRunningDurationMinutes() bool`

HasEstimatedRunningDurationMinutes returns a boolean if a field has been set.

### GetGroups

`func (o *TaskProperties) GetGroups() []GroupIDNamePair`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TaskProperties) GetGroupsOk() (*[]GroupIDNamePair, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TaskProperties) SetGroups(v []GroupIDNamePair)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *TaskProperties) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHealthStatus

`func (o *TaskProperties) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *TaskProperties) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *TaskProperties) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *TaskProperties) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetLogMessages

`func (o *TaskProperties) GetLogMessages() []TaskLogMessage`

GetLogMessages returns the LogMessages field if non-nil, zero value otherwise.

### GetLogMessagesOk

`func (o *TaskProperties) GetLogMessagesOk() (*[]TaskLogMessage, bool)`

GetLogMessagesOk returns a tuple with the LogMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessages

`func (o *TaskProperties) SetLogMessages(v []TaskLogMessage)`

SetLogMessages sets LogMessages field to given value.

### HasLogMessages

`func (o *TaskProperties) HasLogMessages() bool`

HasLogMessages returns a boolean if a field has been set.

### GetParentTask

`func (o *TaskProperties) GetParentTask() ResourceReferenceWithId`

GetParentTask returns the ParentTask field if non-nil, zero value otherwise.

### GetParentTaskOk

`func (o *TaskProperties) GetParentTaskOk() (*ResourceReferenceWithId, bool)`

GetParentTaskOk returns a tuple with the ParentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTask

`func (o *TaskProperties) SetParentTask(v ResourceReferenceWithId)`

SetParentTask sets ParentTask field to given value.

### HasParentTask

`func (o *TaskProperties) HasParentTask() bool`

HasParentTask returns a boolean if a field has been set.

### SetParentTaskNil

`func (o *TaskProperties) SetParentTaskNil(b bool)`

 SetParentTaskNil sets the value for ParentTask to be an explicit nil

### UnsetParentTask
`func (o *TaskProperties) UnsetParentTask()`

UnsetParentTask ensures that no value is present for ParentTask, not even an explicit nil
### GetProgressPercent

`func (o *TaskProperties) GetProgressPercent() int32`

GetProgressPercent returns the ProgressPercent field if non-nil, zero value otherwise.

### GetProgressPercentOk

`func (o *TaskProperties) GetProgressPercentOk() (*int32, bool)`

GetProgressPercentOk returns a tuple with the ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercent

`func (o *TaskProperties) SetProgressPercent(v int32)`

SetProgressPercent sets ProgressPercent field to given value.

### HasProgressPercent

`func (o *TaskProperties) HasProgressPercent() bool`

HasProgressPercent returns a boolean if a field has been set.

### GetRecommendations

`func (o *TaskProperties) GetRecommendations() []TaskRecommendations`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *TaskProperties) GetRecommendationsOk() (*[]TaskRecommendations, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *TaskProperties) SetRecommendations(v []TaskRecommendations)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *TaskProperties) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetRootTask

`func (o *TaskProperties) GetRootTask() ResourceReferenceWithId`

GetRootTask returns the RootTask field if non-nil, zero value otherwise.

### GetRootTaskOk

`func (o *TaskProperties) GetRootTaskOk() (*ResourceReferenceWithId, bool)`

GetRootTaskOk returns a tuple with the RootTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootTask

`func (o *TaskProperties) SetRootTask(v ResourceReferenceWithId)`

SetRootTask sets RootTask field to given value.

### HasRootTask

`func (o *TaskProperties) HasRootTask() bool`

HasRootTask returns a boolean if a field has been set.

### SetRootTaskNil

`func (o *TaskProperties) SetRootTaskNil(b bool)`

 SetRootTaskNil sets the value for RootTask to be an explicit nil

### UnsetRootTask
`func (o *TaskProperties) UnsetRootTask()`

UnsetRootTask ensures that no value is present for RootTask, not even an explicit nil
### GetServices

`func (o *TaskProperties) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *TaskProperties) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *TaskProperties) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *TaskProperties) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSourceResource

`func (o *TaskProperties) GetSourceResource() ResourceReference`

GetSourceResource returns the SourceResource field if non-nil, zero value otherwise.

### GetSourceResourceOk

`func (o *TaskProperties) GetSourceResourceOk() (*ResourceReference, bool)`

GetSourceResourceOk returns a tuple with the SourceResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceResource

`func (o *TaskProperties) SetSourceResource(v ResourceReference)`

SetSourceResource sets SourceResource field to given value.

### HasSourceResource

`func (o *TaskProperties) HasSourceResource() bool`

HasSourceResource returns a boolean if a field has been set.

### SetSourceResourceNil

`func (o *TaskProperties) SetSourceResourceNil(b bool)`

 SetSourceResourceNil sets the value for SourceResource to be an explicit nil

### UnsetSourceResource
`func (o *TaskProperties) UnsetSourceResource()`

UnsetSourceResource ensures that no value is present for SourceResource, not even an explicit nil
### GetStartedAt

`func (o *TaskProperties) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TaskProperties) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TaskProperties) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TaskProperties) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *TaskProperties) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *TaskProperties) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetState

`func (o *TaskProperties) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskProperties) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskProperties) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TaskProperties) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubtreeTaskCount

`func (o *TaskProperties) GetSubtreeTaskCount() int32`

GetSubtreeTaskCount returns the SubtreeTaskCount field if non-nil, zero value otherwise.

### GetSubtreeTaskCountOk

`func (o *TaskProperties) GetSubtreeTaskCountOk() (*int32, bool)`

GetSubtreeTaskCountOk returns a tuple with the SubtreeTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeTaskCount

`func (o *TaskProperties) SetSubtreeTaskCount(v int32)`

SetSubtreeTaskCount sets SubtreeTaskCount field to given value.

### HasSubtreeTaskCount

`func (o *TaskProperties) HasSubtreeTaskCount() bool`

HasSubtreeTaskCount returns a boolean if a field has been set.

### GetSuggestedPollingIntervalSeconds

`func (o *TaskProperties) GetSuggestedPollingIntervalSeconds() int32`

GetSuggestedPollingIntervalSeconds returns the SuggestedPollingIntervalSeconds field if non-nil, zero value otherwise.

### GetSuggestedPollingIntervalSecondsOk

`func (o *TaskProperties) GetSuggestedPollingIntervalSecondsOk() (*int32, bool)`

GetSuggestedPollingIntervalSecondsOk returns a tuple with the SuggestedPollingIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPollingIntervalSeconds

`func (o *TaskProperties) SetSuggestedPollingIntervalSeconds(v int32)`

SetSuggestedPollingIntervalSeconds sets SuggestedPollingIntervalSeconds field to given value.

### HasSuggestedPollingIntervalSeconds

`func (o *TaskProperties) HasSuggestedPollingIntervalSeconds() bool`

HasSuggestedPollingIntervalSeconds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskProperties) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskProperties) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskProperties) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskProperties) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TaskProperties) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TaskProperties) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserId

`func (o *TaskProperties) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TaskProperties) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TaskProperties) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TaskProperties) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


