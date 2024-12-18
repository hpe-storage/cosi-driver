// © Copyright 2024 Hewlett Packard Enterprise Development LP
package iam

import (
	"context"
	"errors"
	sdk "hpe-cosi-osp/glcp_tasks_sdk"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"gotest.tools/v3/assert"
)

func Test_TaskStatus(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IkxnLXg4TUxkTVNv"
	systemId := "DUMMY_SETUP"
	proxy := "http://dummy_proxy:8080"
	apiClient, _ := NewAPIClient(host, token, proxy).GetTaskAPIClient()
	taskUi := GetMockTaskResponseUi()
	task := NewTask(taskUi.TaskUri, systemId, apiClient, context.Background())
	t.Run("TaskStatus successful", func(t *testing.T) {
		expState := string(SUCCEEDED)
		status := sdk.Task{
			ResourceUri: &task.taskId,
			State:       &expState,
		}
		patch := gomonkey.ApplyFuncReturn(NewTask, task)
		patch = patch.ApplyMethodReturn(task, "GetTask", &status, nil)
		defer patch.Reset()

		taskResp, err := TaskStatus(taskUi.TaskUri, systemId, apiClient, context.Background())
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, taskResp, true)
	})
	t.Run("TaskStatus failed", func(t *testing.T) {
		expState := string(FAILED)
		status := sdk.Task{
			ResourceUri: &task.taskId,
			State:       &expState,
		}
		patch := gomonkey.ApplyFuncReturn(NewTask, task)
		patch = patch.ApplyMethodReturn(task, "GetTask", &status, nil)
		defer patch.Reset()

		taskResp, err := TaskStatus(taskUi.TaskUri, systemId, apiClient, context.Background())
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, taskResp, false)
	})

	t.Run("TaskStatus timedout", func(t *testing.T) {
		expState := string(TIMEDOUT)
		status := sdk.Task{
			ResourceUri: &task.taskId,
			State:       &expState,
		}
		patch := gomonkey.ApplyFuncReturn(NewTask, task)
		patch = patch.ApplyMethodReturn(task, "GetTask", &status, nil)
		defer patch.Reset()

		taskResp, err := TaskStatus(taskUi.TaskUri, systemId, apiClient, context.Background())
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, taskResp, false)
	})

	t.Run("TaskStatus unspecified", func(t *testing.T) {
		expState := string(UNSPECIFIED)
		status := sdk.Task{
			ResourceUri: &task.taskId,
			State:       &expState,
		}
		patch := gomonkey.ApplyFuncReturn(NewTask, task)
		patch = patch.ApplyMethodReturn(task, "GetTask", &status, nil)
		defer patch.Reset()

		taskResp, err := TaskStatus(taskUi.TaskUri, systemId, apiClient, context.Background())
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, taskResp, false)
	})

	t.Run("TaskStatus failed with error", func(t *testing.T) {
		patch := gomonkey.ApplyFuncReturn(NewTask, task)
		patch = patch.ApplyMethodReturn(task, "GetTask", nil, errors.New("error fetching task"))
		defer patch.Reset()

		taskResp, err := TaskStatus(taskUi.TaskUri, systemId, apiClient, context.Background())
		if err == nil {
			t.Errorf("FAILED: expected error not found")
		}
		assert.Equal(t, taskResp, false)
	})
}
