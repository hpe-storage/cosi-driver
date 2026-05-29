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
	maskedName := utils.MaskName(u.name)
	user, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7GetSingleUser(u.context, u.systemId, u.name).Execute()
	if err != nil {
		return user, utils.FormatIAMError(fmt.Sprintf("failed to fetch s3 user %s", maskedName), err)
	}
	if r.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to fetch s3 user %s (HTTP %d): %s", maskedName, r.StatusCode, readResponseBody(r))
	}
	return user, err
}

// Returns true/false, if user of name(passed with s3policy instance) is present under DSCC
func (u *s3user) UserExists() (bool, error) {
	log := utils.GetLoggerFromContext(u.context)
	ap, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7GetSingleUser(u.context, u.systemId, u.name).Execute()
	if err != nil && r != nil && r.StatusCode == http.StatusNotFound {
		e, body := getResponseErrorCode(r)
		if e.GetErrorCode() == string(ERR_RESOURCE_NOT_FOUND) {
			return false, nil
		}
		log.Error(err, "Received error while fetching user details from DSCC.", "response body", body)
		return false, err
	} else if ap != nil && err == nil {
		return true, nil
	} else if r != nil {
		e, body := getResponseErrorCode(r)
		if e.GetErrorCode() == string(DOWNSTREAM_SERVICE_ERROR) {
			log.Error(err, "Downstream service error from DSCC;"+
				" verify GLCP credentials, workspace in mounted"+
				" secret, and GLCP_COMMON_CLOUD in COSI pod",
				"response body", body)
			return false, err
		} else {
			log.Error(err, "Received error while fetching user details from DSCC.", "statusCode", r.StatusCode, "status", r.Status, "response body", body)
		}
	}
	return false, err
}

// Creates a user with the passed name in DSCC
// returns the new secret key of the user
func (u *s3user) CreateS3User() (string, error) {
	maskedName := utils.MaskName(u.name)
	userDetails := *sdk.NewUserDetails(u.name)
	resp, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7CreateUser(u.context, u.systemId).
		UserDetails(userDetails).Execute()
	if err != nil {
		return "", utils.FormatIAMError(fmt.Sprintf("failed to create s3 user %s", maskedName), err)
	} else if r.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to create s3 user %s (HTTP %d): %s", maskedName, r.StatusCode, readResponseBody(r))
	}
	return resp.GetSecretKey(), nil
}

// Reset password a user,
// returns the new secret key of the user
// returns an error, if the user is not existing under DSCC
func (u *s3user) ResetPassword() (string, error) {
	maskedName := utils.MaskName(u.name)
	userDetails := *sdk.NewUserDetailsEdit(true)
	resp, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7EditUser(u.context, u.systemId, u.name).UserDetailsEdit(userDetails).Execute()
	if err != nil {
		return "", utils.FormatIAMError(fmt.Sprintf("failed to reset s3 user %s password", maskedName), err)
	} else if r.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to reset s3 user %s password (HTTP %d): %s", maskedName, r.StatusCode, readResponseBody(r))
	}
	return resp.GetSecretKey(), nil
}

// Applies a created access policy to the existing user
// returns the task details created for this process
// returns an error, if the user or policy is not existing under DSCC
func (u *s3user) ApplyPolicy(policies []string) (*sdk.TaskResponseUi, error) {
	maskedName := utils.MaskName(u.name)
	policy := *sdk.NewApplyPolicy(u.name, policies, string(USER))

	taskUI, r, err := u.client.ObjectstoreIdentitiesAPI.ApplyPolicy(u.context, u.systemId).ApplyPolicy(policy).Execute()
	if err != nil {
		return taskUI, utils.FormatIAMError(fmt.Sprintf("failed to apply policies to user %s", maskedName), err)
	}
	if r.StatusCode != http.StatusAccepted || taskUI.GetStatus() != string(SUBMITTED) {
		err = fmt.Errorf("failed to apply policies to user %s (HTTP %d): %s", maskedName, r.StatusCode, readResponseBody(r))
	}
	return taskUI, err
}

// Deletes an existing user
// returns the task details created for this process
// returns an error, if the user is not existing under DSCC
func (u *s3user) DeleteS3User() (*sdk.TaskResponseUi, error) {
	maskedName := utils.MaskName(u.name)
	taskUI, r, err := u.client.ObjectstoreIdentitiesAPI.DeviceType7DeleteUserById(u.context, u.systemId, u.name).Execute()
	if err != nil {
		return taskUI, utils.FormatIAMError(fmt.Sprintf("failed to delete s3 user %s", maskedName), err)
	}
	if r.StatusCode != http.StatusAccepted || taskUI.GetStatus() != string(SUBMITTED) {
		err = fmt.Errorf("failed to delete s3 user %s (HTTP %d): %s", maskedName, r.StatusCode, readResponseBody(r))
	}
	return taskUI, err
}

// readResponseBody reads the HTTP response body and returns a parsed,
// human-readable error string using ParseIAMErrorResponse.
// Returns an empty string if the body cannot be read.
func readResponseBody(r *http.Response) string {
	if r == nil || r.Body == nil {
		return ""
	}
	_, body := getResponseErrorCode(r)
	return body
}

func getResponseErrorCode(r *http.Response) (sdk.InternalErrorResponseUi, string) {
	var b sdk.InternalErrorResponseUi
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		return b, ""
	}
	// Best-effort unmarshal into structured error; ignore failures
	_ = json.Unmarshal(body, &b)
	return b, utils.ParseIAMErrorResponse(body)
}
