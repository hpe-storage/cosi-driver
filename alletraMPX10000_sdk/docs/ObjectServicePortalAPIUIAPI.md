# \ObjectServicePortalAPIUIAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BucketCapacityHistory**](ObjectServicePortalAPIUIAPI.md#BucketCapacityHistory) | **Get** /api/v1/storage-systems/device-type7/{systemId}/buckets/{bucketId}/capacity-history | Get capacity trend data of buckets



## BucketCapacityHistory

> BucketCapacityHistory BucketCapacityHistory(ctx, systemId, bucketId).Range_(range_).TimeIntervalMin(timeIntervalMin).Execute()

Get capacity trend data of buckets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "alletraMPX10000_sdk"
)

func main() {
	systemId := "8UW0002928" // string | ID unique to every object service device
	bucketId := "bucket1" // string | ID unique to every bucket created in objectstore
	range_ := "startTime eq 1605063600 and endTime eq 1605186000" // string | range will define start and end time in which query has to be made. (optional)
	timeIntervalMin := int32(56) // int32 | It defines granularity in minutes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectServicePortalAPIUIAPI.BucketCapacityHistory(context.Background(), systemId, bucketId).Range_(range_).TimeIntervalMin(timeIntervalMin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectServicePortalAPIUIAPI.BucketCapacityHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BucketCapacityHistory`: BucketCapacityHistory
	fmt.Fprintf(os.Stdout, "Response from `ObjectServicePortalAPIUIAPI.BucketCapacityHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**bucketId** | **string** | ID unique to every bucket created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiBucketCapacityHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **range_** | **string** | range will define start and end time in which query has to be made. | 
 **timeIntervalMin** | **int32** | It defines granularity in minutes. | 

### Return type

[**BucketCapacityHistory**](BucketCapacityHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

