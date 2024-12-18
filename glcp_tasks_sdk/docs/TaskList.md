# TaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Task**](Task.md) |  | 
**PageLimit** | Pointer to **int32** | The limit query parameter specified in the request. | [optional] 
**PageOffset** | Pointer to **int32** | The offset query parameter specified in the request. | [optional] 
**Total** | Pointer to **int32** | Total number of items matching the filter query parameter in the request. | [optional] 

## Methods

### NewTaskList

`func NewTaskList(items []Task, ) *TaskList`

NewTaskList instantiates a new TaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskListWithDefaults

`func NewTaskListWithDefaults() *TaskList`

NewTaskListWithDefaults instantiates a new TaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TaskList) GetItems() []Task`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TaskList) GetItemsOk() (*[]Task, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TaskList) SetItems(v []Task)`

SetItems sets Items field to given value.


### GetPageLimit

`func (o *TaskList) GetPageLimit() int32`

GetPageLimit returns the PageLimit field if non-nil, zero value otherwise.

### GetPageLimitOk

`func (o *TaskList) GetPageLimitOk() (*int32, bool)`

GetPageLimitOk returns a tuple with the PageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLimit

`func (o *TaskList) SetPageLimit(v int32)`

SetPageLimit sets PageLimit field to given value.

### HasPageLimit

`func (o *TaskList) HasPageLimit() bool`

HasPageLimit returns a boolean if a field has been set.

### GetPageOffset

`func (o *TaskList) GetPageOffset() int32`

GetPageOffset returns the PageOffset field if non-nil, zero value otherwise.

### GetPageOffsetOk

`func (o *TaskList) GetPageOffsetOk() (*int32, bool)`

GetPageOffsetOk returns a tuple with the PageOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageOffset

`func (o *TaskList) SetPageOffset(v int32)`

SetPageOffset sets PageOffset field to given value.

### HasPageOffset

`func (o *TaskList) HasPageOffset() bool`

HasPageOffset returns a boolean if a field has been set.

### GetTotal

`func (o *TaskList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TaskList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TaskList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TaskList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


