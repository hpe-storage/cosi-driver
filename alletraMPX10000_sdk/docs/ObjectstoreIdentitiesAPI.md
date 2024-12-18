# \ObjectstoreIdentitiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyPolicy**](ObjectstoreIdentitiesAPI.md#ApplyPolicy) | **Put** /api/v1/storage-systems/device-type7/{systemId}/apply-policy | Apply policy to existing user or group in HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7CreateAccessPolicy**](ObjectstoreIdentitiesAPI.md#DeviceType7CreateAccessPolicy) | **Post** /api/v1/storage-systems/device-type7/{systemId}/access-policies | Create new access policy for HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7CreateGroup**](ObjectstoreIdentitiesAPI.md#DeviceType7CreateGroup) | **Post** /api/v1/storage-systems/device-type7/{systemId}/groups | Create new group in HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7CreateUser**](ObjectstoreIdentitiesAPI.md#DeviceType7CreateUser) | **Post** /api/v1/storage-systems/device-type7/{systemId}/users | Create new user in HPE Alletra Storage MP X10000 device
[**DeviceType7DeleteAccessPolicyById**](ObjectstoreIdentitiesAPI.md#DeviceType7DeleteAccessPolicyById) | **Delete** /api/v1/storage-systems/device-type7/{systemId}/access-policies/{policyId} | Delete HPE Alletra Storage MP X10000 ObjectStore access policy
[**DeviceType7DeleteGroupById**](ObjectstoreIdentitiesAPI.md#DeviceType7DeleteGroupById) | **Delete** /api/v1/storage-systems/device-type7/{systemId}/groups/{groupId} | Delete group from HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7DeleteUserById**](ObjectstoreIdentitiesAPI.md#DeviceType7DeleteUserById) | **Delete** /api/v1/storage-systems/device-type7/{systemId}/users/{userId} | Delete user from HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7EditUser**](ObjectstoreIdentitiesAPI.md#DeviceType7EditUser) | **Put** /api/v1/storage-systems/device-type7/{systemId}/users/{userId} | Edit the properties of an existing user in HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7GetAccessPolicyById**](ObjectstoreIdentitiesAPI.md#DeviceType7GetAccessPolicyById) | **Get** /api/v1/storage-systems/device-type7/{systemId}/access-policies/{policyId} | Get single HPE Alletra Storage MP X10000 ObjectStore access policy
[**DeviceType7GetGroupById**](ObjectstoreIdentitiesAPI.md#DeviceType7GetGroupById) | **Get** /api/v1/storage-systems/device-type7/{systemId}/groups/{groupId} | Get single group details for HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7GetSingleUser**](ObjectstoreIdentitiesAPI.md#DeviceType7GetSingleUser) | **Get** /api/v1/storage-systems/device-type7/{systemId}/users/{userId} | Get the properties of an existing user in ObjectStore
[**DeviceType7GetUser**](ObjectstoreIdentitiesAPI.md#DeviceType7GetUser) | **Get** /api/v1/storage-systems/device-type7/{systemId}/users | Get the properties of existing users in ObjectStore
[**DeviceType7ListAccessPolicies**](ObjectstoreIdentitiesAPI.md#DeviceType7ListAccessPolicies) | **Get** /api/v1/storage-systems/device-type7/{systemId}/access-policies | Get all access policies for HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7ListGroups**](ObjectstoreIdentitiesAPI.md#DeviceType7ListGroups) | **Get** /api/v1/storage-systems/device-type7/{systemId}/groups | Get all groups for HPE Alletra Storage MP X10000 ObjectStore
[**DeviceType7UpdateAccessPolicyById**](ObjectstoreIdentitiesAPI.md#DeviceType7UpdateAccessPolicyById) | **Put** /api/v1/storage-systems/device-type7/{systemId}/access-policies/{policyId} | Update HPE Alletra Storage MP X10000 ObjectStore access policy
[**DeviceType7UpdateGroupById**](ObjectstoreIdentitiesAPI.md#DeviceType7UpdateGroupById) | **Put** /api/v1/storage-systems/device-type7/{systemId}/groups/{groupId} | Update group details in HPE Alletra Storage MP X10000 ObjectStore



## ApplyPolicy

> TaskResponseUi ApplyPolicy(ctx, systemId).ApplyPolicy(applyPolicy).Execute()

Apply policy to existing user or group in HPE Alletra Storage MP X10000 ObjectStore



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
	applyPolicy := *openapiclient.NewApplyPolicy("NewAccessPolicy1", []string{"Policies_example"}, "Type_example") // ApplyPolicy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.ApplyPolicy(context.Background(), systemId).ApplyPolicy(applyPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.ApplyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyPolicy`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.ApplyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applyPolicy** | [**ApplyPolicy**](ApplyPolicy.md) |  | 

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


## DeviceType7CreateAccessPolicy

> TaskResponseUi DeviceType7CreateAccessPolicy(ctx, systemId).CreateAccessPolicyBody(createAccessPolicyBody).Execute()

Create new access policy for HPE Alletra Storage MP X10000 ObjectStore



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
	createAccessPolicyBody := *openapiclient.NewCreateAccessPolicyBody("NewAccessPolicy1", []openapiclient.PolicyStatementInput{*openapiclient.NewPolicyStatementInput([]string{"Action_example"}, "ALLOW", openapiclient.PolicyStatementInput_resource{ArrayOfString: new([]string)})}, "2012-10-17") // CreateAccessPolicyBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7CreateAccessPolicy(context.Background(), systemId).CreateAccessPolicyBody(createAccessPolicyBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7CreateAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7CreateAccessPolicy`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7CreateAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7CreateAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAccessPolicyBody** | [**CreateAccessPolicyBody**](CreateAccessPolicyBody.md) |  | 

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


## DeviceType7CreateGroup

> TaskResponseUi DeviceType7CreateGroup(ctx, systemId).ObjectCreateGroupBody(objectCreateGroupBody).Execute()

Create new group in HPE Alletra Storage MP X10000 ObjectStore



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
	objectCreateGroupBody := *openapiclient.NewObjectCreateGroupBody("NewGroup1") // ObjectCreateGroupBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7CreateGroup(context.Background(), systemId).ObjectCreateGroupBody(objectCreateGroupBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7CreateGroup`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7CreateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7CreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectCreateGroupBody** | [**ObjectCreateGroupBody**](ObjectCreateGroupBody.md) |  | 

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


## DeviceType7CreateUser

> UserResponseDetails DeviceType7CreateUser(ctx, systemId).UserDetails(userDetails).Execute()

Create new user in HPE Alletra Storage MP X10000 device



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
	userDetails := *openapiclient.NewUserDetails("Name_example") // UserDetails |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7CreateUser(context.Background(), systemId).UserDetails(userDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7CreateUser`: UserResponseDetails
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7CreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userDetails** | [**UserDetails**](UserDetails.md) |  | 

### Return type

[**UserResponseDetails**](UserResponseDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7DeleteAccessPolicyById

> TaskResponseUi DeviceType7DeleteAccessPolicyById(ctx, systemId, policyId).Execute()

Delete HPE Alletra Storage MP X10000 ObjectStore access policy



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
	policyId := "policy1" // string | ID unique to every access-policy created in objectstore

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7DeleteAccessPolicyById(context.Background(), systemId, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7DeleteAccessPolicyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7DeleteAccessPolicyById`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7DeleteAccessPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**policyId** | **string** | ID unique to every access-policy created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7DeleteAccessPolicyByIdRequest struct via the builder pattern


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


## DeviceType7DeleteGroupById

> TaskResponseUi DeviceType7DeleteGroupById(ctx, systemId, groupId).Execute()

Delete group from HPE Alletra Storage MP X10000 ObjectStore



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
	groupId := "testGroup" // string | ID unique to every group created in objectstore

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7DeleteGroupById(context.Background(), systemId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7DeleteGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7DeleteGroupById`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7DeleteGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**groupId** | **string** | ID unique to every group created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7DeleteGroupByIdRequest struct via the builder pattern


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


## DeviceType7DeleteUserById

> TaskResponseUi DeviceType7DeleteUserById(ctx, systemId, userId).Execute()

Delete user from HPE Alletra Storage MP X10000 ObjectStore



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
	userId := "testUser" // string | ID unique to every user created in objectstore

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7DeleteUserById(context.Background(), systemId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7DeleteUserById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7DeleteUserById`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7DeleteUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**userId** | **string** | ID unique to every user created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7DeleteUserByIdRequest struct via the builder pattern


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


## DeviceType7EditUser

> UserResponseDetails DeviceType7EditUser(ctx, systemId, userId).UserDetailsEdit(userDetailsEdit).Execute()

Edit the properties of an existing user in HPE Alletra Storage MP X10000 ObjectStore



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
	userId := "testUser" // string | ID unique to every user created in objectstore
	userDetailsEdit := *openapiclient.NewUserDetailsEdit(false) // UserDetailsEdit |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7EditUser(context.Background(), systemId, userId).UserDetailsEdit(userDetailsEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7EditUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7EditUser`: UserResponseDetails
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7EditUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**userId** | **string** | ID unique to every user created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7EditUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userDetailsEdit** | [**UserDetailsEdit**](UserDetailsEdit.md) |  | 

### Return type

[**UserResponseDetails**](UserResponseDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7GetAccessPolicyById

> AccessPolicy DeviceType7GetAccessPolicyById(ctx, systemId, policyId).Select_(select_).Execute()

Get single HPE Alletra Storage MP X10000 ObjectStore access policy



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
	policyId := "policy1" // string | ID unique to every access-policy created in objectstore
	select_ := "uid" // string | Query to select only the required parameters, separated by . if nested (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7GetAccessPolicyById(context.Background(), systemId, policyId).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7GetAccessPolicyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7GetAccessPolicyById`: AccessPolicy
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7GetAccessPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**policyId** | **string** | ID unique to every access-policy created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7GetAccessPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **string** | Query to select only the required parameters, separated by . if nested | 

### Return type

[**AccessPolicy**](AccessPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7GetGroupById

> GroupItem DeviceType7GetGroupById(ctx, groupId, systemId).Select_(select_).Execute()

Get single group details for HPE Alletra Storage MP X10000 ObjectStore



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
	groupId := "testGroup" // string | ID unique to every group created in objectstore
	systemId := "8UW0002928" // string | ID unique to every object service device
	select_ := "uid" // string | Query to select only the required parameters, separated by . if nested (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7GetGroupById(context.Background(), groupId, systemId).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7GetGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7GetGroupById`: GroupItem
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7GetGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID unique to every group created in objectstore | 
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7GetGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **string** | Query to select only the required parameters, separated by . if nested | 

### Return type

[**GroupItem**](GroupItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7GetSingleUser

> UserListDetails DeviceType7GetSingleUser(ctx, systemId, userId).Select_(select_).Execute()

Get the properties of an existing user in ObjectStore



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
	userId := "testUser" // string | ID unique to every user created in objectstore
	select_ := "uid" // string | Query to select only the required parameters, separated by . if nested (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7GetSingleUser(context.Background(), systemId, userId).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7GetSingleUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7GetSingleUser`: UserListDetails
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7GetSingleUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**userId** | **string** | ID unique to every user created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7GetSingleUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **string** | Query to select only the required parameters, separated by . if nested | 

### Return type

[**UserListDetails**](UserListDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7GetUser

> HomeUsersList DeviceType7GetUser(ctx, systemId).Filter(filter).Limit(limit).Offset(offset).Select_(select_).Sort(sort).Execute()

Get the properties of existing users in ObjectStore



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
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7GetUser(context.Background(), systemId).Filter(filter).Limit(limit).Offset(offset).Select_(select_).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7GetUser`: HomeUsersList
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | oData query to filter bucket by Key. | 
 **limit** | **int32** | Number of items to return at a time | 
 **offset** | **int32** | The offset of the first item in the collection to return | 
 **select_** | **string** | Query to select only the required parameters, separated by . if nested | 
 **sort** | **string** | A list of properties defining the sort order. This takes a single property name followed by the direction to sort in, separated by a space. The supported properties are &#x60;systemUid&#x60;, &#x60;tenantId&#x60;, &#x60;name&#x60;, &#x60;id&#x60; and &#x60;generation&#x60;. If not specified, the default behaviour is to sort by &#x60;generation&#x60;. The supported directions are &#x60;asc&#x60; and &#x60;desc&#x60; for ascending and descending respectively.  | [default to &quot;generation desc&quot;]

### Return type

[**HomeUsersList**](HomeUsersList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7ListAccessPolicies

> AccessPolicies DeviceType7ListAccessPolicies(ctx, systemId).Filter(filter).Limit(limit).Offset(offset).Select_(select_).Sort(sort).Execute()

Get all access policies for HPE Alletra Storage MP X10000 ObjectStore



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
	sort := "name desc" // string | A list of properties defining the sort order. This takes a single property name followed by the direction to sort in, separated by a space. The supported properties are `id`, `systemUid`, `tenantId`, `name` and `generation`. If not specified, the default behaviour is to sort by `generation`. The supported directions are `asc` and `desc` for ascending and descending respectively.  (optional) (default to "generation desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7ListAccessPolicies(context.Background(), systemId).Filter(filter).Limit(limit).Offset(offset).Select_(select_).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7ListAccessPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7ListAccessPolicies`: AccessPolicies
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7ListAccessPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7ListAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | oData query to filter bucket by Key. | 
 **limit** | **int32** | Number of items to return at a time | 
 **offset** | **int32** | The offset of the first item in the collection to return | 
 **select_** | **string** | Query to select only the required parameters, separated by . if nested | 
 **sort** | **string** | A list of properties defining the sort order. This takes a single property name followed by the direction to sort in, separated by a space. The supported properties are &#x60;id&#x60;, &#x60;systemUid&#x60;, &#x60;tenantId&#x60;, &#x60;name&#x60; and &#x60;generation&#x60;. If not specified, the default behaviour is to sort by &#x60;generation&#x60;. The supported directions are &#x60;asc&#x60; and &#x60;desc&#x60; for ascending and descending respectively.  | [default to &quot;generation desc&quot;]

### Return type

[**AccessPolicies**](AccessPolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7ListGroups

> GetGroupDetails DeviceType7ListGroups(ctx, systemId).Limit(limit).Offset(offset).Select_(select_).Filter(filter).Sort(sort).Execute()

Get all groups for HPE Alletra Storage MP X10000 ObjectStore



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
	limit := int32(10) // int32 | Number of items to return at a time (optional)
	offset := int32(5) // int32 | The offset of the first item in the collection to return (optional)
	select_ := "uid" // string | Query to select only the required parameters, separated by . if nested (optional)
	filter := "id eq "abc"" // string | oData query to filter bucket by Key. (optional)
	sort := "generation desc" // string | A list of properties defining the sort order. This takes a single property name followed by the direction to sort in, separated by a space. The supported properties are `name`, `systemUid`, `tenantId` and `generation`. If not specified, the default behaviour is to sort by `generation`. The supported directions are `asc` and `desc` for ascending and descending respectively.  (optional) (default to "generation desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7ListGroups(context.Background(), systemId).Limit(limit).Offset(offset).Select_(select_).Filter(filter).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7ListGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7ListGroups`: GetGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7ListGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7ListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of items to return at a time | 
 **offset** | **int32** | The offset of the first item in the collection to return | 
 **select_** | **string** | Query to select only the required parameters, separated by . if nested | 
 **filter** | **string** | oData query to filter bucket by Key. | 
 **sort** | **string** | A list of properties defining the sort order. This takes a single property name followed by the direction to sort in, separated by a space. The supported properties are &#x60;name&#x60;, &#x60;systemUid&#x60;, &#x60;tenantId&#x60; and &#x60;generation&#x60;. If not specified, the default behaviour is to sort by &#x60;generation&#x60;. The supported directions are &#x60;asc&#x60; and &#x60;desc&#x60; for ascending and descending respectively.  | [default to &quot;generation desc&quot;]

### Return type

[**GetGroupDetails**](GetGroupDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceType7UpdateAccessPolicyById

> TaskResponseUi DeviceType7UpdateAccessPolicyById(ctx, systemId, policyId).UpdateAccessPolicyBody(updateAccessPolicyBody).Execute()

Update HPE Alletra Storage MP X10000 ObjectStore access policy



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
	policyId := "policy1" // string | ID unique to every access-policy created in objectstore
	updateAccessPolicyBody := *openapiclient.NewUpdateAccessPolicyBody([]openapiclient.PolicyStatementInput{*openapiclient.NewPolicyStatementInput([]string{"Action_example"}, "ALLOW", openapiclient.PolicyStatementInput_resource{ArrayOfString: new([]string)})}, "2012-10-17") // UpdateAccessPolicyBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7UpdateAccessPolicyById(context.Background(), systemId, policyId).UpdateAccessPolicyBody(updateAccessPolicyBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7UpdateAccessPolicyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7UpdateAccessPolicyById`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7UpdateAccessPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**policyId** | **string** | ID unique to every access-policy created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7UpdateAccessPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAccessPolicyBody** | [**UpdateAccessPolicyBody**](UpdateAccessPolicyBody.md) |  | 

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


## DeviceType7UpdateGroupById

> TaskResponseUi DeviceType7UpdateGroupById(ctx, systemId, groupId).ObjectUpdateGroupBody(objectUpdateGroupBody).Execute()

Update group details in HPE Alletra Storage MP X10000 ObjectStore



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
	groupId := "testGroup" // string | ID unique to every group created in objectstore
	objectUpdateGroupBody := *openapiclient.NewObjectUpdateGroupBody() // ObjectUpdateGroupBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectstoreIdentitiesAPI.DeviceType7UpdateGroupById(context.Background(), systemId, groupId).ObjectUpdateGroupBody(objectUpdateGroupBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectstoreIdentitiesAPI.DeviceType7UpdateGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceType7UpdateGroupById`: TaskResponseUi
	fmt.Fprintf(os.Stdout, "Response from `ObjectstoreIdentitiesAPI.DeviceType7UpdateGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID unique to every object service device | 
**groupId** | **string** | ID unique to every group created in objectstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceType7UpdateGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **objectUpdateGroupBody** | [**ObjectUpdateGroupBody**](ObjectUpdateGroupBody.md) |  | 

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

