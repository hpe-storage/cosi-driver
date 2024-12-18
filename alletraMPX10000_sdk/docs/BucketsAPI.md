# \BucketsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceType7CreateBucket**](BucketsAPI.md#DeviceType7CreateBucket) | **Post** /api/v1/storage-systems/device-type7/{systemId}/buckets | Create new bucket in HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7DeleteBucketById**](BucketsAPI.md#DeviceType7DeleteBucketById) | **Delete** /api/v1/storage-systems/device-type7/{systemId}/buckets/{bucketId} | Delete bucket from HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7EditBucket**](BucketsAPI.md#DeviceType7EditBucket) | **Put** /api/v1/storage-systems/device-type7/{systemId}/buckets/{bucketId} | Edit the properties of an existing bucket in HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7ListBuckets**](BucketsAPI.md#DeviceType7ListBuckets) | **Get** /api/v1/storage-systems/device-type7/{systemId}/buckets | Get all buckets for HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7SingleBuckets**](BucketsAPI.md#DeviceType7SingleBuckets) | **Get** /api/v1/storage-systems/device-type7/{systemId}/buckets/{bucketId} | Get single HPE Alletra Storage MP X10000 ObjectStore bucket



## DeviceType7CreateBucket

> TaskResponseUi DeviceType7CreateBucket(ctx, systemId).CreateBucketBody(createBucketBody).Execute()

Create new bucket in HPE Alletra Storage MP X10000 ObjectStore



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
	createBucketBody := *openapiclient.NewCreateBucketBody("myBucketName") // CreateBucketBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketsAPI.DeviceType7CreateBucket(context.Background(), systemId).CreateBucketBody(createBucketBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.DeviceType7CreateBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7CreateBucket`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.DeviceType7CreateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7CreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBucketBody** | [**CreateBucketBody**](CreateBucketBody.md) |  | 

### Return type

[**TaskResponseUi**](TaskResponseUi.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7DeleteBucketById

> TaskResponseUi DeviceType7DeleteBucketById(ctx, systemId, bucketId).Execute()

Delete bucket from HPE Alletra Storage MP X10000 ObjectStore



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketsAPI.DeviceType7DeleteBucketById(context.Background(), systemId, bucketId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.DeviceType7DeleteBucketById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7DeleteBucketById`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.DeviceType7DeleteBucketById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**bucketId** | **string** | ID unique to every bucket created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7DeleteBucketByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaskResponseUi**](TaskResponseUi.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7EditBucket

> TaskResponseUi DeviceType7EditBucket(ctx, systemId, bucketId).EditBucketBody(editBucketBody).Execute()

Edit the properties of an existing bucket in HPE Alletra Storage MP X10000 ObjectStore



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
	editBucketBody := *openapiclient.NewEditBucketBody() // EditBucketBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketsAPI.DeviceType7EditBucket(context.Background(), systemId, bucketId).EditBucketBody(editBucketBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.DeviceType7EditBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7EditBucket`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.DeviceType7EditBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**bucketId** | **string** | ID unique to every bucket created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7EditBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editBucketBody** | [**EditBucketBody**](EditBucketBody.md) |  | 

### Return type

[**TaskResponseUi**](TaskResponseUi.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7ListBuckets

> BucketsList DeviceType7ListBuckets(ctx, systemId).Filter(filter).Limit(limit).Offset(offset).Select_(select_).Sort(sort).Execute()

Get all buckets for HPE Alletra Storage MP X10000 ObjectStore



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
	filter := "id eq "abc"" // string | oData query to filter bucket by Key. (optional)
	limit := int32(10) // int32 | Number of items to return at a time (optional)
	offset := int32(5) // int32 | The offset of the first item in the collection to return (optional)
	select_ := "uid" // string | Query to select only the required parameters, separated by . if nested (optional)
	sort := "name desc" // string | A list of properties defining the sort order. This takes a single property name followed by the direction to sort in, separated by a space. The supported properties are `systemUid`, `tenantId`, `name`, `id` and `generation`. If not specified, the default behaviour is to sort by `generation`. The supported directions are `asc` and `desc` for ascending and descending respectively.  (optional) (default to "generation desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketsAPI.DeviceType7ListBuckets(context.Background(), systemId).Filter(filter).Limit(limit).Offset(offset).Select_(select_).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.DeviceType7ListBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7ListBuckets`: BucketsList
	fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.DeviceType7ListBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7ListBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | oData query to filter bucket by Key. | 
 **limit** | **int32** | Number of items to return at a time | 
 **offset** | **int32** | The offset of the first item in the collection to return | 
 **select_** | **string** | Query to select only the required parameters, separated by . if nested | 
 **sort** | **string** | A list of properties defining the sort order. This takes a single property name followed by the direction to sort in, separated by a space. The supported properties are &#x60;systemUid&#x60;, &#x60;tenantId&#x60;, &#x60;name&#x60;, &#x60;id&#x60; and &#x60;generation&#x60;. If not specified, the default behaviour is to sort by &#x60;generation&#x60;. The supported directions are &#x60;asc&#x60; and &#x60;desc&#x60; for ascending and descending respectively.  | [default to &quot;generation desc&quot;]

### Return type

[**BucketsList**](BucketsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7SingleBuckets

> Buckets DeviceType7SingleBuckets(ctx, systemId, bucketId).Select_(select_).Execute()

Get single HPE Alletra Storage MP X10000 ObjectStore bucket



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
	select_ := "uid" // string | Query to select only the required parameters, separated by . if nested (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketsAPI.DeviceType7SingleBuckets(context.Background(), systemId, bucketId).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.DeviceType7SingleBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7SingleBuckets`: Buckets
	fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.DeviceType7SingleBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**bucketId** | **string** | ID unique to every bucket created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7SingleBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **string** | Query to select only the required parameters, separated by . if nested | 

### Return type

[**Buckets**](Buckets.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

