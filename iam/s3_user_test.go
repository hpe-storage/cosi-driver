// © Copyright 2024 Hewlett Packard Enterprise Development LP

package iam

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	sdk "hpe-cosi-osp/alletraMPX10000_sdk"
	"io"
	"net/http"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"gotest.tools/v3/assert"
)

var badRoutingResponse = map[string]string{
	"httpStatusCode": "404",
	"errorCode":      "HPE_GL_ERROR_NOT_FOUND",
	"message":        "Routing details for the customer not found.",
	"debugId":        "CsggFMdSMOR55P3k99--gd9plOITW8d3hxmw5b67dCsw6ww615TNzw==",
}

var resourceNotFoundResponse = map[string]string{
	"errorCode": string(ERR_RESOURCE_NOT_FOUND),
	"error":     "Invalid resource or Resource not found",
}

func Test_GetS3User(t *testing.T) {
	user := getNewS3User()
	name := user.name
	systemId := user.systemId
	request := sdk.ApiDeviceType7GetSingleUserRequest{}
	gId := int32(1725948764)
	t.Run("GetS3User successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7GetSingleUser", request)
		us := sdk.UserListDetails{
			Generation: *sdk.NewNullableInt32(&gId),
			Id:         &name,
			Name:       &name,
			SystemUid:  &systemId,
		}
		r := http.Response{StatusCode: http.StatusOK}
		patch = patch.ApplyMethodReturn(request, "Execute", &us, &r, nil)
		defer patch.Reset()
		_, err := user.GetS3User()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
	})
	t.Run("GetS3User failure", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7GetSingleUser", request)
		r := http.Response{StatusCode: http.StatusBadRequest}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, nil)
		defer patch.Reset()
		if _, err := user.GetS3User(); err == nil {
			t.Errorf("FAILED: expected error not found")
		}
	})
}

func Test_UserExists(t *testing.T) {
	user := getNewS3User()
	name := user.name
	systemId := user.systemId
	request := sdk.ApiDeviceType7GetSingleUserRequest{}
	gId := int32(1725948764)
	t.Run("UserExists successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7GetSingleUser", request)
		request := sdk.ApiDeviceType7GetSingleUserRequest{}
		ap := sdk.UserListDetails{
			Generation: *sdk.NewNullableInt32(&gId),
			Id:         &name,
			Name:       &name,
			SystemUid:  &systemId,
		}
		r := http.Response{StatusCode: http.StatusOK}
		patch = patch.ApplyMethodReturn(request, "Execute", &ap, &r, nil)
		defer patch.Reset()
		isExist, err := user.UserExists()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, isExist, true)
	})
	t.Run("UserExists false", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7GetSingleUser", request)
		body, _ := json.Marshal(resourceNotFoundResponse)
		r := http.Response{StatusCode: http.StatusNotFound, Body: io.NopCloser(bytes.NewBuffer(body))}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, errors.New("404 Not Found"))
		defer patch.Reset()
		isExist, err := user.UserExists()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, isExist, false)
	})

	t.Run("InCorrect DSCC URL", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7GetSingleUser", request)
		body, _ := json.Marshal(badRoutingResponse)
		r := http.Response{StatusCode: http.StatusNotFound, Body: io.NopCloser(bytes.NewBuffer(body))}
		e := errors.New("404 Not Found")
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, e)
		defer patch.Reset()
		isExist, err := user.UserExists()
		if err == nil {
			t.Errorf("FAILED: expected error not found")
		}
		assert.Equal(t, isExist, false)
		assert.Equal(t, err, e)
	})

	t.Run("Invalid URL with inline character", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7GetSingleUser", request)
		err := fmt.Errorf("parse \"https://%s\\n/api/v1/storage-systems/device-type7/HFDT1Y87BC/users/user_ba-4ea8396e-53eb-43aa-ac52-10224f0c4273\": net/url: invalid control character in URL", host)
		patch = patch.ApplyMethodReturn(request, "Execute", nil, nil, err)
		defer patch.Reset()
		isExist, e := user.UserExists()
		if e == nil {
			t.Errorf("FAILED: expected error not found")
		}
		assert.Equal(t, isExist, false)
		assert.Equal(t, err, e)
	})

}

func Test_CreateS3User(t *testing.T) {
	user := getNewS3User()
	request := sdk.ApiDeviceType7CreateUserRequest{}
	t.Run("CreateS3User successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7CreateUser", request)
		patch = patch.ApplyMethodReturn(request, "UserDetails", request)
		secretKey := "rd39awiMIao1FvwIg4U"
		resp := sdk.UserResponseDetails{
			SecretKey: &secretKey,
		}
		r := http.Response{StatusCode: http.StatusCreated}
		patch = patch.ApplyMethodReturn(request, "Execute", &resp, &r, nil)
		defer patch.Reset()
		secret, err := user.CreateS3User()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, secret, secretKey)
	})
	t.Run("CreateS3User failure", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7CreateUser", request)
		patch = patch.ApplyMethodReturn(request, "UserDetails", request)
		r := http.Response{StatusCode: http.StatusBadRequest}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, nil)
		defer patch.Reset()
		if _, err := user.CreateS3User(); err == nil {
			t.Errorf("FAILED: expected error not found")
		}
	})
}

func Test_ResetPassword(t *testing.T) {
	user := getNewS3User()
	request := sdk.ApiDeviceType7EditUserRequest{}
	t.Run("ResetPassword successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7EditUser", request)
		patch = patch.ApplyMethodReturn(request, "UserDetailsEdit", request)
		secretKey := "rd39awiMIao1FvwIg4U"
		resp := sdk.UserResponseDetails{
			SecretKey: &secretKey,
		}
		r := http.Response{StatusCode: http.StatusCreated}
		patch = patch.ApplyMethodReturn(request, "Execute", &resp, &r, nil)
		defer patch.Reset()
		secret, err := user.ResetPassword()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, secret, secretKey)
	})
	t.Run("ResetPassword failure", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7EditUser", request)
		patch = patch.ApplyMethodReturn(request, "UserDetailsEdit", request)
		r := http.Response{StatusCode: http.StatusBadRequest}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, nil)
		defer patch.Reset()
		if _, err := user.ResetPassword(); err == nil {
			t.Errorf("FAILED: expected error not found")
		}
	})
}

func Test_ApplyPolicy(t *testing.T) {
	user := getNewS3User()
	request := sdk.ApiApplyPolicyRequest{}
	policyName := "bucket1_policy"
	t.Run("ApplyPolicy successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "ApplyPolicy", request)
		patch = patch.ApplyMethodReturn(request, "ApplyPolicy", request)
		taskUi := GetMockTaskResponseUi()

		r := http.Response{StatusCode: http.StatusAccepted}
		patch = patch.ApplyMethodReturn(request, "Execute", taskUi, &r, nil)
		defer patch.Reset()
		taskResp, err := user.ApplyPolicy([]string{policyName})
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, taskResp, taskUi)
	})
	t.Run("ApplyPolicy failure", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "ApplyPolicy", request)
		patch = patch.ApplyMethodReturn(request, "ApplyPolicy", request)
		r := http.Response{StatusCode: http.StatusBadRequest}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, nil)
		defer patch.Reset()
		if _, err := user.ApplyPolicy([]string{policyName}); err == nil {
			t.Errorf("FAILED: expected error not found")
		}
	})
}

func Test_DeleteS3User(t *testing.T) {
	user := getNewS3User()
	request := sdk.ApiDeviceType7DeleteUserByIdRequest{}
	t.Run("DeleteS3User successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7DeleteUserById", request)
		taskUi := GetMockTaskResponseUi()
		r := http.Response{StatusCode: http.StatusAccepted}
		patch = patch.ApplyMethodReturn(request, "Execute", taskUi, &r, nil)
		defer patch.Reset()
		taskResp, err := user.DeleteS3User()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.DeepEqual(t, taskResp, taskUi)
	})
	t.Run("DeleteS3User failure", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(user.client.ObjectstoreIdentitiesAPI, "DeviceType7DeleteUserById", request)
		r := http.Response{StatusCode: http.StatusBadRequest}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, nil)
		defer patch.Reset()
		if _, err := user.DeleteS3User(); err == nil {
			t.Errorf("FAILED: expected error not found")
		}
	})
}

func getNewS3User() *s3user {
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IkxnLXg4TUxkTVNv"
	userName := "bucket1_user"
	systemId := "DUMMY_SETUP"
	proxy := "http://dummy_proxy:8080"
	apiClient, _ := NewAPIClient(host, token, proxy).GetAPIClient()
	return NewS3User(userName, systemId, apiClient, context.Background())
}
