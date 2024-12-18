// © Copyright 2024 Hewlett Packard Enterprise Development LP
package iam

import (
	"context"
	"encoding/json"
	"fmt"
	sdk "hpe-cosi-osp/alletraMPX10000_sdk"
	"hpe-cosi-osp/servers/provisioner/utils"
	"io"
	"net/http"
)

// Defines attributes to create an S3 user instance( used to manage S3 user in DSCC )
// These attributes are used to identify user & system-id in DSCC
type s3user struct {
	name     string
	systemId string
	client   *sdk.APIClient
	context  context.Context
}

// Returns an instance of the s3user struct
func NewS3User(userName, systemId string, client *sdk.APIClient, context context.Context) *s3user {
	return &s3user{
		name:     userName,
		systemId: systemId,
		client:   client,
		context:  context,
	}
}

// Returns the access policy by name(passed with s3policy instance)
func (u *s3user) GetS3User() (*sdk.UserListDetails, error) {
	user, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7GetSingleUser(u.context, u.systemId, u.name).Execute()
	if err == nil && r.StatusCode != http.StatusOK {
		err = fmt.Errorf("request failed while fetching s3 user %s, err: %v", u.name, r)
	}
	return user, err
}

// Returns true/false, if user of name(passed with s3policy instance) is present under DSCC
func (u *s3user) UserExists() (bool, error) {
	log := utils.GetLoggerFromContext(u.context)
	ap, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7GetSingleUser(u.context, u.systemId, u.name).Execute()
	if err != nil && r != nil && r.StatusCode == http.StatusNotFound {
		e, body := getResponseErrorCode(u.context, r)
		if e.GetErrorCode() == string(ERR_RESOURCE_NOT_FOUND) {
			return false, nil
		}
		log.Error(err, "Received error while fetching user details from DSCC.", "response body", body)
		return false, err
	} else if ap != nil && err == nil {
		return true, nil
	}
	return false, err
}

// Creates a user with the passed name in DSCC
// returns the new secret key of the user
func (u *s3user) CreateS3User() (string, error) {
	userDetails := *sdk.NewUserDetails(u.name)
	resp, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7CreateUser(u.context, u.systemId).
		UserDetails(userDetails).Execute()
	if err != nil {
		return "", err
	} else if r.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("request failed while creating s3 user %s, err: %v", u.name, r)
	}
	return resp.GetSecretKey(), nil
}

// Reset password a user,
// returns the new secret key of the user
// returns an error, if the user is not existing under DSCC
func (u *s3user) ResetPassword() (string, error) {
	userDetails := *sdk.NewUserDetailsEdit(true)
	resp, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7EditUser(u.context, u.systemId, u.name).UserDetailsEdit(userDetails).Execute()
	if err != nil {
		return "", err
	} else if r.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("request failed while resetting s3 user %s password, err: %v", u.name, r)
	}
	return resp.GetSecretKey(), nil
}

// Applies a created access policy to the existing user
// returns the task details created for this process
// returns an error, if the user or policy is not existing under DSCC
func (u *s3user) ApplyPolicy(policies []string) (*sdk.TaskResponseUi, error) {
	policy := *sdk.NewApplyPolicy(u.name, policies, string(USER))

	taskUI, r, err := u.client.ObjectstoreIdentitiesAPI.ApplyPolicy(u.context, u.systemId).ApplyPolicy(policy).Execute()
	if err == nil && (r.StatusCode != http.StatusAccepted || taskUI.GetStatus() != string(SUBMITTED)) {
		err = fmt.Errorf("request failed while applying policies %v, to the user %s, err: %v", policies, u.name, r)
	}
	return taskUI, err
}

// Deletes an existing user
// returns the task details created for this process
// returns an error, if the user is not existing under DSCC
func (u *s3user) DeleteS3User() (*sdk.TaskResponseUi, error) {
	taskUI, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7DeleteUserById(u.context, u.systemId, u.name).Execute()
	if err == nil && (r.StatusCode != http.StatusAccepted || taskUI.GetStatus() != string(SUBMITTED)) {
		err = fmt.Errorf("request failed while delete s3 user %s, err: %v", u.name, r)
	}
	return taskUI, err
}

func getResponseErrorCode(ctx context.Context, r *http.Response) (sdk.InternalErrorResponseUi, string) {
	log := utils.GetLoggerFromContext(ctx)
	var b sdk.InternalErrorResponseUi
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err, "Unable to retrieve body from http response")
		return b, ""
	}
	if err = json.Unmarshal(body, &b); err != nil {
		log.Error(err, "Failed to unmarshal body from, %s", string(body)[:])
		return b, string(body)[:]
	}
	return b, string(body)[:]
}
