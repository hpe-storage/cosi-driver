// © Copyright 2024 Hewlett Packard Enterprise Development LP

package iam

import (
	"time"
)

var (
	ACCESS_GRANT_TYPE       = "client_credentials"
	GLCP_ACCESS_TOKEN_URL   = "as/token.oauth2"
	HTTP_REQUEST_TIMEOUT    = 30 * time.Second
	APPLICATION_URL_ENCODED = "application/x-www-form-urlencoded"
	DSCC_POLICY_VERSION     = "2012-10-17"
	WAIT_TIME               = 2 * time.Second
)

// Defines types of entities in DSCC , ex: USER|POLICY
type EntityType string

const (
	USER EntityType = "USER"
)

type Status string

const (
	SUBMITTED = "SUBMITTED"
)

// S3 PROTOCOL constants
// TODO : revisit hardcoding of these values, after R1

var S3_RESOURCE_PREFIX = "arn:aws:s3"

// Defines rules of restriction for a policy to a bucket
type Effect string

const (
	ALLOW Effect = "ALLOW"
)

// Defines access privileges for a policy to a bucket
type Action string

const (
	FULL_ACCESS Action = "s3:*"
)

// Defines different States of a Task run/finished/ongoing in DSCC
type State string

const (
	SUCCEEDED   State = "SUCCEEDED"
	INITIALIZED State = "INITIALIZED"
	RUNNING     State = "RUNNING"
	FAILED      State = "FAILED"
	TIMEDOUT    State = "TIMEDOUT"
	PAUSED      State = "PAUSED"
	UNSPECIFIED State = "UNSPECIFIED"
)

// Error codes returned by DSCC
type HttpErrorCode string

const (
	ERR_RESOURCE_NOT_FOUND HttpErrorCode = "ERR_RESOURCE_NOT_FOUND"
)
