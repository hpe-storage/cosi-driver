// © Copyright 2024 Hewlett Packard Enterprise Development LP
package iam

import (
	"context"
	"fmt"
	sdk "hpe-cosi-osp/glcp_tasks_sdk"
	"hpe-cosi-osp/servers/provisioner/utils"
	"net/http"
	"time"
)

// Defines attributes to create an task instance( used to manage Tasks in DSCC )
// These attributes are used to identify tasks, system-id & bucket in DSCC
type task struct {
	taskId   string
	systemId string
	client   *sdk.APIClient
	context  context.Context
}

// Returns an instance of the task struct
func NewTask(taskId, systemId string, client *sdk.APIClient, ctx context.Context) *task {
	return &task{
		taskId:   taskId,
		systemId: systemId,
		client:   client,
		context:  ctx,
	}
}

// Returns the running task by id (passed with s3policy instance)
func (t *task) GetTask() (*sdk.Task, error) {
	task, r, err := t.client.TasksAPI.GetTask(t.context, t.taskId).Execute()
	if err == nil && r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed while fetching taskId %s, err: %v", t.taskId, r)
	}
	return task, err
}

func isTaskFailed(state string) bool {
	if state == string(FAILED) || state == string(TIMEDOUT) || state == string(UNSPECIFIED) {
		return true
	}
	return false
}

func isTaskOngoing(state string) (b bool) {
	if state == string(INITIALIZED) || state == string(PAUSED) || state == string(RUNNING) {
		b = true
	}
	return
}

// Returns true/false based on the DSCC task status (succeeeded/bailed out or failed)
// returns error, if the task is not existing
func TaskStatus(taskUri, systemId string, client *sdk.APIClient, ctx context.Context) (bool, error) {
	log := utils.GetLoggerFromContext(ctx)
	t := NewTask(taskUri, systemId, client, ctx)
	for {
		time.Sleep(WAIT_TIME)
		status, err := t.GetTask()
		if status != nil {
			log.Info("Task status : " + *status.State)
		}
		if err != nil && status == nil {
			log.Error(err, "error while fetching task status")
			return false, err
		}
		if isTaskOngoing(*status.State) {
			continue
		} else if isTaskFailed(*status.State) {
			return false, nil
		} else {
			break
		}
	}
	return true, nil
}
