# \TasksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTask**](TasksAPI.md#GetTask) | **Get** /api/v1/tasks/{id} | Returns details of a specific task
[**ListTasks**](TasksAPI.md#ListTasks) | **Get** /api/v1/tasks | Returns a list of tasks accessible by the user



## GetTask

> Task GetTask(ctx, id).Select_(select_).Execute()

Returns details of a specific task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "glcp_tasks_sdk"
)

func main() {
	id := "c1a0eb78-41a0-4151-93b2-f057ffeca3f3" // string | The UUID of the object
	select_ := "id,name" // string | A list of properties to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.GetTask(context.Background(), id).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The UUID of the object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **string** | A list of properties to include in the response. | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> TaskList ListTasks(ctx).Offset(offset).Limit(limit).Filter(filter).Sort(sort).Select_(select_).Execute()

Returns a list of tasks accessible by the user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "glcp_tasks_sdk"
)

func main() {
	offset := int32(30) // int32 | The offset query parameter should be used in conjunction with limit for paging, e.g.: offset=30&&limit=10. The offset is the number of items from the beginning of the result set to the first item included in the response.  (optional)
	limit := int32(10) // int32 | The limit query parameter should be used in conjunction with offset for paging, e.g.: offset=30&&limit=10. The limit is the maximum number of items to include in the response.  (optional)
	filter := "owner.name eq fred@example.com" // string | The expression to filter responses. (optional)
	sort := "id desc,name asc" // string | A comma separated list of properties to sort by, followed by a direction  indicator (\"asc\" or \"desc\"). If no direction indicator is specified the  default order is ascending.  (optional)
	select_ := "id,name" // string | A list of properties in the items collection to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ListTasks(context.Background()).Offset(offset).Limit(limit).Filter(filter).Sort(sort).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasks`: TaskList
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The offset query parameter should be used in conjunction with limit for paging, e.g.: offset&#x3D;30&amp;&amp;limit&#x3D;10. The offset is the number of items from the beginning of the result set to the first item included in the response.  | 
 **limit** | **int32** | The limit query parameter should be used in conjunction with offset for paging, e.g.: offset&#x3D;30&amp;&amp;limit&#x3D;10. The limit is the maximum number of items to include in the response.  | 
 **filter** | **string** | The expression to filter responses. | 
 **sort** | **string** | A comma separated list of properties to sort by, followed by a direction  indicator (\&quot;asc\&quot; or \&quot;desc\&quot;). If no direction indicator is specified the  default order is ascending.  | 
 **select_** | **string** | A list of properties in the items collection to include in the response. | 

### Return type

[**TaskList**](TaskList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

