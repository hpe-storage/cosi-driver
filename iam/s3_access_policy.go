// © Copyright 2024 Hewlett Packard Enterprise Development LP
package iam

import (
	"context"
	"fmt"
	sdk "hpe-cosi-osp/alletraMPX10000_sdk"
	"hpe-cosi-osp/servers/provisioner/utils"
	"net/http"

	"golang.org/x/exp/rand"
)

// Defines attributes to create an S3 policy instance( used to manage Access Policy in DSCC )
// These attributes are used to identify access policy, system-id & bucket in DSCC
type s3policy struct {
	name       string
	systemId   string
	client     *sdk.APIClient
	context    context.Context
	bucketName string
	//TODO: Pass access permissions as well, i.e Full access ,  Read/Write etc
}

// Returns an instance of the s3policy struct
func NewS3Policy(policyName, bucketName, systemId string, client *sdk.APIClient, context context.Context) *s3policy {
	return &s3policy{
		name:       policyName,
		bucketName: bucketName,
		systemId:   systemId,
		client:     client,
		context:    context,
	}
}

// Returns the access policy by name(passed with s3policy instance)
func (p *s3policy) GetS3AccessPolicy() (*sdk.AccessPolicy, error) {
	ap, r, err := p.client.ObjectstoreIdentitiesAPI.DeviceType7GetAccessPolicyById(p.context, p.systemId, p.name).Execute()
	if err == nil && r.StatusCode != http.StatusOK {
		err = fmt.Errorf("request failed while fetching s3 access policy %s, err : %v", p.name, r)
	}
	return ap, err
}

// Returns true/false, if access policy of name(passed with s3policy instance) is present under DSCC
func (p *s3policy) PolicyExists() (bool, error) {
	log := utils.GetLoggerFromContext(p.context)
	ap, r, err := p.client.ObjectstoreIdentitiesAPI.DeviceType7GetAccessPolicyById(p.context, p.systemId, p.name).Execute()
	if err != nil && r != nil && r.StatusCode == http.StatusNotFound {
		e, body := getResponseErrorCode(p.context, r)
		if e.GetErrorCode() == string(ERR_RESOURCE_NOT_FOUND) {
			return false, nil
		}
		log.Error(err, "Received error while fetching policy details from DSCC.", "response body", body)
		return false, err
	} else if ap != nil && err == nil {
		return true, nil
	}
	return false, err
}

// Returns true/false, if access policy of name has full access to a bucket(passed with s3policy instance)
func (p *s3policy) PolicyHasFullBucketAccess() (bool, error) {
	policy, err := p.GetS3AccessPolicy()
	if err != nil {
		return false, err
	}
	return hasFullAccess(p.bucketName, policy), nil
}

// Creates an access policy with the passed name in DSCC
// returns the task details created for this process
func (p *s3policy) CreateS3AccessPolicy() (*sdk.TaskResponseUi, error) {

	policyStatement := p.formFullAccessRuleForBucket()

	createAccessPolicyBody := *sdk.NewCreateAccessPolicyBody(p.name, []sdk.PolicyStatementInput{policyStatement}, DSCC_POLICY_VERSION)
	taskUI, r, err := p.client.ObjectstoreIdentitiesAPI.DeviceType7CreateAccessPolicy(p.context, p.systemId).
		CreateAccessPolicyBody(createAccessPolicyBody).Execute()
	if err == nil && (r.StatusCode != http.StatusAccepted || taskUI.GetStatus() != string(SUBMITTED)) {
		err = fmt.Errorf("request failed while creating s3 access policy %s, err: %v", p.name, r)
	}
	return taskUI, err
}

// Updates an existing access policy with the passed attributes
// returns the task details created for this process
// returns an error, if the policy is not existing under DSCC
func (p *s3policy) UpdateS3AccessPolicy() (*sdk.TaskResponseUi, error) {
	policyStatement := p.formFullAccessRuleForBucket()
	updateAccessPolicyBody := *sdk.NewUpdateAccessPolicyBody([]sdk.PolicyStatementInput{policyStatement}, DSCC_POLICY_VERSION)
	taskUI, r, err := p.client.ObjectstoreIdentitiesAPI.DeviceType7UpdateAccessPolicyById(p.context, p.systemId, p.name).
		UpdateAccessPolicyBody(updateAccessPolicyBody).Execute()
	if err == nil && (r.StatusCode != http.StatusAccepted || taskUI.GetStatus() != string(SUBMITTED)) {
		err = fmt.Errorf("request failed while updating s3 access policy %s, err: %v", p.name, r)
	}
	return taskUI, err
}

// Deletes an existing access policy
// returns the task details created for this process
// returns an error, if the policy is not existing under DSCC
func (p *s3policy) DeleteS3AccessPolicy() (*sdk.TaskResponseUi, error) {
	taskUI, r, err := p.client.ObjectstoreIdentitiesAPI.DeviceType7DeleteAccessPolicyById(p.context, p.systemId, p.name).Execute()

	if err == nil && (r.StatusCode != http.StatusAccepted || taskUI.GetStatus() != string(SUBMITTED)) {
		err = fmt.Errorf("request failed while deleting s3 access policy %s, err: %v", p.name, r)
	}
	return taskUI, err
}

func getBucketActions() []string {
	return []string{string(FULL_ACCESS)}
}

func (p *s3policy) getResources() []string {
	resource := fmt.Sprintf("%s:::%s", S3_RESOURCE_PREFIX, p.bucketName)
	return []string{resource, resource + "/*"}
}

func (p *s3policy) formFullAccessRuleForBucket() sdk.PolicyStatementInput {
	action := getBucketActions()
	resources := p.getResources()
	return *sdk.NewPolicyStatementInput(action, string(ALLOW), sdk.PolicyStatementInputResource{ArrayOfString: &resources})
}

func hasFullAccess(bucketName string, policy *sdk.AccessPolicy) bool {
	return false
}

// Visible & to be used for UT alone
func GetMockTaskResponseUi() *sdk.TaskResponseUi {
	m := "Mock Task response raised for Unit Testing"
	s := "SUBMITTED"
	pattern := "f3073f7f-97a9-49c3-8b4d-45038b1aaeee"
	return &sdk.TaskResponseUi{
		Message: &m,
		Status:  &s,
		TaskUri: string(pattern[rand.Intn(len(pattern))]),
	}
}
