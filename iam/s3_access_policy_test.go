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

func Test_GetS3AccessPolicy(t *testing.T) {
	policy := getNewS3Policy()
	policyName := policy.name
	systemId := policy.systemId
	request := sdk.ApiDeviceType7GetAccessPolicyByIdRequest{}
	t.Run("GetS3AccessPolicy successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7GetAccessPolicyById", request)
		gId := int32(1725948764)
		ap := sdk.AccessPolicy{
			Generation: &gId,
			Id:         &policyName,
			Name:       &policyName,
			SystemUid:  &systemId,
		}
		r := http.Response{StatusCode: http.StatusOK}
		patch = patch.ApplyMethodReturn(request, "Execute", &ap, &r, nil)
		defer patch.Reset()
		policy, err := policy.GetS3AccessPolicy()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.DeepEqual(t, policy, &ap)
	})
	t.Run("GetS3AccessPolicy failure", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7GetAccessPolicyById", request)
		r := http.Response{StatusCode: http.StatusBadRequest}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, nil)
		defer patch.Reset()
		if _, err := policy.GetS3AccessPolicy(); err == nil {
			t.Errorf("FAILED: expected error not found")
		}
	})
}

func Test_PolicyExists(t *testing.T) {
	policy := getNewS3Policy()
	policyName := policy.name
	systemId := policy.systemId
	request := sdk.ApiDeviceType7GetAccessPolicyByIdRequest{}
	t.Run("PolicyExists successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7GetAccessPolicyById", request)
		gId := int32(1725948764)
		ap := sdk.AccessPolicy{
			Generation: &gId,
			Id:         &policyName,
			Name:       &policyName,
			SystemUid:  &systemId,
		}
		r := http.Response{StatusCode: http.StatusOK}
		patch = patch.ApplyMethodReturn(request, "Execute", &ap, &r, nil)
		defer patch.Reset()
		isExist, err := policy.PolicyExists()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, isExist, true)
	})
	t.Run("PolicyExists false", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7GetAccessPolicyById", request)
		body, _ := json.Marshal(resourceNotFoundResponse)
		r := http.Response{StatusCode: http.StatusNotFound, Body: io.NopCloser(bytes.NewBuffer(body))}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, errors.New("404 Not Found"))
		defer patch.Reset()
		isExist, err := policy.PolicyExists()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, isExist, false)
	})

	t.Run("InCorrect DSCC URL", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7GetAccessPolicyById", request)
		body, _ := json.Marshal(badRoutingResponse)
		r := http.Response{StatusCode: http.StatusNotFound, Body: io.NopCloser(bytes.NewBuffer(body))}
		e := errors.New("404 Not Found")
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, e)
		defer patch.Reset()
		isExist, err := policy.PolicyExists()
		if err == nil {
			t.Errorf("FAILED: expected error not found")
		}
		assert.Equal(t, isExist, false)
		assert.Equal(t, err, e)
	})

	t.Run("Invalid URL with inline character", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7GetAccessPolicyById", request)
		err := fmt.Errorf("parse \"https://%s\\n/api/v1/storage-systems/device-type7/HFDT1Y87BC/access-policies/acp_ba-4ea8396e-53eb-43aa-ac52-10224f0c4273\": net/url: invalid control character in URL", host)
		patch = patch.ApplyMethodReturn(request, "Execute", nil, nil, err)
		defer patch.Reset()
		isExist, e := policy.PolicyExists()
		if e == nil {
			t.Errorf("FAILED: expected error not found")
		}
		assert.Equal(t, isExist, false)
		assert.Equal(t, err, e)
	})
}

func Test_CreateS3AccessPolicy(t *testing.T) {
	policy := getNewS3Policy()
	request := sdk.ApiDeviceType7CreateAccessPolicyRequest{}
	request1 := sdk.ApiDeviceType7CreateAccessPolicyRequest{}
	t.Run("CreateS3AccessPolicy successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7CreateAccessPolicy", request)
		patch = patch.ApplyMethodReturn(request, "CreateAccessPolicyBody", request1)
		taskUi := GetMockTaskResponseUi()
		r := http.Response{StatusCode: http.StatusAccepted}
		patch = patch.ApplyMethodReturn(request1, "Execute", taskUi, &r, nil)
		defer patch.Reset()
		taskResp, err := policy.CreateS3AccessPolicy()
		if err != nil {
			fmt.Println(err)
			t.Errorf("FAILED: unexpected error")
		}
		assert.DeepEqual(t, taskResp, taskUi)
	})
	t.Run("CreateS3AccessPolicy failure", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7CreateAccessPolicy", request)
		patch = patch.ApplyMethodReturn(request, "CreateAccessPolicyBody", request1)
		r := http.Response{StatusCode: http.StatusBadRequest}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, nil)
		defer patch.Reset()
		if _, err := policy.CreateS3AccessPolicy(); err == nil {
			t.Errorf("FAILED: expected error not found")
		}
	})
}

func Test_UpdateS3AccessPolicy(t *testing.T) {
	policy := getNewS3Policy()
	request := sdk.ApiDeviceType7UpdateAccessPolicyByIdRequest{}
	request1 := sdk.ApiDeviceType7UpdateAccessPolicyByIdRequest{}
	t.Run("UpdateS3AccessPolicy successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7UpdateAccessPolicyById", request)
		patch = patch.ApplyMethodReturn(request, "UpdateAccessPolicyBody", request1)
		taskUi := GetMockTaskResponseUi()
		r := http.Response{StatusCode: http.StatusAccepted}
		patch = patch.ApplyMethodReturn(request1, "Execute", taskUi, &r, nil)
		defer patch.Reset()
		taskResp, err := policy.UpdateS3AccessPolicy()
		if err != nil {
			fmt.Println(err)
			t.Errorf("FAILED: unexpected error")
		}
		assert.DeepEqual(t, taskResp, taskUi)
	})
	t.Run("UpdateS3AccessPolicy failure", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7UpdateAccessPolicyById", request)
		patch = patch.ApplyMethodReturn(request, "UpdateAccessPolicyBody", request1)
		r := http.Response{StatusCode: http.StatusBadRequest}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, nil)
		defer patch.Reset()
		if _, err := policy.UpdateS3AccessPolicy(); err == nil {
			t.Errorf("FAILED: expected error not found")
		}
	})
}

func Test_DeleteS3AccessPolicy(t *testing.T) {
	policy := getNewS3Policy()
	request := sdk.ApiDeviceType7DeleteAccessPolicyByIdRequest{}
	t.Run("DeleteS3AccessPolicy successful", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7DeleteAccessPolicyById", request)
		taskUi := GetMockTaskResponseUi()
		r := http.Response{StatusCode: http.StatusAccepted}
		patch = patch.ApplyMethodReturn(request, "Execute", taskUi, &r, nil)
		defer patch.Reset()
		taskResp, err := policy.DeleteS3AccessPolicy()
		if err != nil {
			fmt.Println(err)
			t.Errorf("FAILED: unexpected error")
		}
		assert.DeepEqual(t, taskResp, taskUi)
	})
	t.Run("DeleteS3AccessPolicy failure", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(policy.client.ObjectstoreIdentitiesAPI, "DeviceType7DeleteAccessPolicyById", request)
		r := http.Response{StatusCode: http.StatusBadRequest}
		patch = patch.ApplyMethodReturn(request, "Execute", nil, &r, nil)
		defer patch.Reset()
		if _, err := policy.DeleteS3AccessPolicy(); err == nil {
			t.Errorf("FAILED: expected error not found")
		}
	})
}

func getNewS3Policy() *s3policy {
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IkxnLXg4TUxkTVNv"
	policyName := "bucket1_policy"
	bucket := "bucket1"
	systemId := "DUMMY_SETUP"
	proxy := "http://dummy_proxy:8080"
	apiClient, _ := NewAPIClient(host, token, proxy).GetAPIClient()
	return NewS3Policy(policyName, bucket, systemId, apiClient, context.Background())
}
