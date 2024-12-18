// © Copyright 2024 Hewlett Packard Enterprise Development LP

package iam

import (
	"testing"
)

var host = "us1.xxxx.xxxxx.hpe.com"

func Test_GetAPIClient(t *testing.T) {
	token := "bearerdummyoxyzxxzzz12xxxx341111zzzzyyyyyyQQQQQHHHHH"
	t.Run("GetAPIClient successful", func(t *testing.T) {
		proxy := "http://dummy_proxy:8080"
		apiClient := NewAPIClient(host, token, proxy)
		client, err := apiClient.GetAPIClient()
		if err != nil || client == nil {
			t.Errorf("FAILED: unexpected error")
		}
	})
	t.Run("GetAPIClient with proxy & password", func(t *testing.T) {
		proxy := "http://dummy_user:dummy_password@dummy_proxy:8080"
		apiClient := NewAPIClient(host, token, proxy)
		client, err := apiClient.GetAPIClient()
		if err != nil || client == nil {
			t.Errorf("FAILED: unexpected error")
		}
	})
}

func Test_GetTaskAPIClient(t *testing.T) {
	token := "bearerdummyoxyzxxzzz12xxxx341111zzzzyyyyyyQQQQQHHHHH"
	t.Run("GetTaskAPIClient successful", func(t *testing.T) {
		proxy := "http://dummy_proxy:8080"
		apiClient := NewAPIClient(host, token, proxy)
		client, err := apiClient.GetTaskAPIClient()
		if err != nil || client == nil {
			t.Errorf("FAILED: unexpected error")
		}
	})
	t.Run("GetTaskAPIClient with proxy & password", func(t *testing.T) {
		proxy := "http://dummy_user:dummy_password@dummy_proxy:8080"
		apiClient := NewAPIClient(host, token, proxy)
		client, err := apiClient.GetTaskAPIClient()
		if err != nil || client == nil {
			t.Errorf("FAILED: unexpected error")
		}
	})
}
