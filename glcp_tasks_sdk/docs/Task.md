# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | The customer application identifier | [optional] [readonly] 
**Generation** | Pointer to **int64** | A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date. | [optional] [readonly] 
**Id** | Pointer to **string** | An identifier for the resource, usually a UUID. | [optional] [readonly] 
**Name** | Pointer to **string** | A system specified name for the resource. | [optional] 
**ResourceUri** | Pointer to **string** | The &#39;self&#39; reference for this resource. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of resource. | [optional] [readonly] 
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

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *Task) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Task) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Task) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Task) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetGeneration

`func (o *Task) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *Task) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *Task) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *Task) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Task) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceUri

`func (o *Task) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *Task) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *Task) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.

### HasResourceUri

`func (o *Task) HasResourceUri() bool`

HasResourceUri returns a boolean if a field has been set.

### GetType

`func (o *Task) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Task) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Task) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Task) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdditionalDetails

`func (o *Task) GetAdditionalDetails() TaskConsoleReference`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *Task) GetAdditionalDetailsOk() (*TaskConsoleReference, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *Task) SetAdditionalDetails(v TaskConsoleReference)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *Task) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetAssociatedResources

`func (o *Task) GetAssociatedResources() []ResourceReference`

GetAssociatedResources returns the AssociatedResources field if non-nil, zero value otherwise.

### GetAssociatedResourcesOk

`func (o *Task) GetAssociatedResourcesOk() (*[]ResourceReference, bool)`

GetAssociatedResourcesOk returns a tuple with the AssociatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedResources

`func (o *Task) SetAssociatedResources(v []ResourceReference)`

SetAssociatedResources sets AssociatedResources field to given value.

### HasAssociatedResources

`func (o *Task) HasAssociatedResources() bool`

HasAssociatedResources returns a boolean if a field has been set.

### GetChildTasks

`func (o *Task) GetChildTasks() []ResourceReference`

GetChildTasks returns the ChildTasks field if non-nil, zero value otherwise.

### GetChildTasksOk

`func (o *Task) GetChildTasksOk() (*[]ResourceReference, bool)`

GetChildTasksOk returns a tuple with the ChildTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTasks

`func (o *Task) SetChildTasks(v []ResourceReference)`

SetChildTasks sets ChildTasks field to given value.

### HasChildTasks

`func (o *Task) HasChildTasks() bool`

HasChildTasks returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Task) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Task) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *Task) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Task) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Task) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Task) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndedAt

`func (o *Task) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Task) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Task) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Task) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *Task) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *Task) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetError

`func (o *Task) GetError() ErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Task) GetErrorOk() (*ErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Task) SetError(v ErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *Task) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *Task) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Task) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetEstimatedRunningDurationMinutes

`func (o *Task) GetEstimatedRunningDurationMinutes() int32`

GetEstimatedRunningDurationMinutes returns the EstimatedRunningDurationMinutes field if non-nil, zero value otherwise.

### GetEstimatedRunningDurationMinutesOk

`func (o *Task) GetEstimatedRunningDurationMinutesOk() (*int32, bool)`

GetEstimatedRunningDurationMinutesOk returns a tuple with the EstimatedRunningDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRunningDurationMinutes

`func (o *Task) SetEstimatedRunningDurationMinutes(v int32)`

SetEstimatedRunningDurationMinutes sets EstimatedRunningDurationMinutes field to given value.

### HasEstimatedRunningDurationMinutes

`func (o *Task) HasEstimatedRunningDurationMinutes() bool`

HasEstimatedRunningDurationMinutes returns a boolean if a field has been set.

### GetGroups

`func (o *Task) GetGroups() []GroupIDNamePair`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Task) GetGroupsOk() (*[]GroupIDNamePair, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Task) SetGroups(v []GroupIDNamePair)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Task) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHealthStatus

`func (o *Task) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *Task) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *Task) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *Task) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetLogMessages

`func (o *Task) GetLogMessages() []TaskLogMessage`

GetLogMessages returns the LogMessages field if non-nil, zero value otherwise.

### GetLogMessagesOk

`func (o *Task) GetLogMessagesOk() (*[]TaskLogMessage, bool)`

GetLogMessagesOk returns a tuple with the LogMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessages

`func (o *Task) SetLogMessages(v []TaskLogMessage)`

SetLogMessages sets LogMessages field to given value.

### HasLogMessages

`func (o *Task) HasLogMessages() bool`

HasLogMessages returns a boolean if a field has been set.

### GetParentTask

`func (o *Task) GetParentTask() ResourceReferenceWithId`

GetParentTask returns the ParentTask field if non-nil, zero value otherwise.

### GetParentTaskOk

`func (o *Task) GetParentTaskOk() (*ResourceReferenceWithId, bool)`

GetParentTaskOk returns a tuple with the ParentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTask

`func (o *Task) SetParentTask(v ResourceReferenceWithId)`

SetParentTask sets ParentTask field to given value.

### HasParentTask

`func (o *Task) HasParentTask() bool`

HasParentTask returns a boolean if a field has been set.

### SetParentTaskNil

`func (o *Task) SetParentTaskNil(b bool)`

 SetParentTaskNil sets the value for ParentTask to be an explicit nil

### UnsetParentTask
`func (o *Task) UnsetParentTask()`

UnsetParentTask ensures that no value is present for ParentTask, not even an explicit nil
### GetProgressPercent

`func (o *Task) GetProgressPercent() int32`

GetProgressPercent returns the ProgressPercent field if non-nil, zero value otherwise.

### GetProgressPercentOk

`func (o *Task) GetProgressPercentOk() (*int32, bool)`

GetProgressPercentOk returns a tuple with the ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercent

`func (o *Task) SetProgressPercent(v int32)`

SetProgressPercent sets ProgressPercent field to given value.

### HasProgressPercent

`func (o *Task) HasProgressPercent() bool`

HasProgressPercent returns a boolean if a field has been set.

### GetRecommendations

`func (o *Task) GetRecommendations() []TaskRecommendations`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *Task) GetRecommendationsOk() (*[]TaskRecommendations, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *Task) SetRecommendations(v []TaskRecommendations)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *Task) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetRootTask

`func (o *Task) GetRootTask() ResourceReferenceWithId`

GetRootTask returns the RootTask field if non-nil, zero value otherwise.

### GetRootTaskOk

`func (o *Task) GetRootTaskOk() (*ResourceReferenceWithId, bool)`

GetRootTaskOk returns a tuple with the RootTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootTask

`func (o *Task) SetRootTask(v ResourceReferenceWithId)`

SetRootTask sets RootTask field to given value.

### HasRootTask

`func (o *Task) HasRootTask() bool`

HasRootTask returns a boolean if a field has been set.

### SetRootTaskNil

`func (o *Task) SetRootTaskNil(b bool)`

 SetRootTaskNil sets the value for RootTask to be an explicit nil

### UnsetRootTask
`func (o *Task) UnsetRootTask()`

UnsetRootTask ensures that no value is present for RootTask, not even an explicit nil
### GetServices

`func (o *Task) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Task) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Task) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *Task) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSourceResource

`func (o *Task) GetSourceResource() ResourceReference`

GetSourceResource returns the SourceResource field if non-nil, zero value otherwise.

### GetSourceResourceOk

`func (o *Task) GetSourceResourceOk() (*ResourceReference, bool)`

GetSourceResourceOk returns a tuple with the SourceResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceResource

`func (o *Task) SetSourceResource(v ResourceReference)`

SetSourceResource sets SourceResource field to given value.

### HasSourceResource

`func (o *Task) HasSourceResource() bool`

HasSourceResource returns a boolean if a field has been set.

### SetSourceResourceNil

`func (o *Task) SetSourceResourceNil(b bool)`

 SetSourceResourceNil sets the value for SourceResource to be an explicit nil

### UnsetSourceResource
`func (o *Task) UnsetSourceResource()`

UnsetSourceResource ensures that no value is present for SourceResource, not even an explicit nil
### GetStartedAt

`func (o *Task) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Task) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Task) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Task) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *Task) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Task) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetState

`func (o *Task) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Task) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Task) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Task) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubtreeTaskCount

`func (o *Task) GetSubtreeTaskCount() int32`

GetSubtreeTaskCount returns the SubtreeTaskCount field if non-nil, zero value otherwise.

### GetSubtreeTaskCountOk

`func (o *Task) GetSubtreeTaskCountOk() (*int32, bool)`

GetSubtreeTaskCountOk returns a tuple with the SubtreeTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeTaskCount

`func (o *Task) SetSubtreeTaskCount(v int32)`

SetSubtreeTaskCount sets SubtreeTaskCount field to given value.

### HasSubtreeTaskCount

`func (o *Task) HasSubtreeTaskCount() bool`

HasSubtreeTaskCount returns a boolean if a field has been set.

### GetSuggestedPollingIntervalSeconds

`func (o *Task) GetSuggestedPollingIntervalSeconds() int32`

GetSuggestedPollingIntervalSeconds returns the SuggestedPollingIntervalSeconds field if non-nil, zero value otherwise.

### GetSuggestedPollingIntervalSecondsOk

`func (o *Task) GetSuggestedPollingIntervalSecondsOk() (*int32, bool)`

GetSuggestedPollingIntervalSecondsOk returns a tuple with the SuggestedPollingIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPollingIntervalSeconds

`func (o *Task) SetSuggestedPollingIntervalSeconds(v int32)`

SetSuggestedPollingIntervalSeconds sets SuggestedPollingIntervalSeconds field to given value.

### HasSuggestedPollingIntervalSeconds

`func (o *Task) HasSuggestedPollingIntervalSeconds() bool`

HasSuggestedPollingIntervalSeconds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Task) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Task) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Task) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Task) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserId

`func (o *Task) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Task) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Task) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Task) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


