// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Package utils
package utils

import (
	"encoding/xml"
)

const (
	ACCESS_POLICY_PREFIX = "acp_"
	USER_PREFIX          = "user_"
	//KEYS in glcp secret
	GLCP_USER_CLIENTID   = "glcpUserClientId"
	GLCP_USER_SECRET_KEY = "glcpUserSecretKey"
	GLCP_COMMON_CLOUD    = "GLCP_COMMON_CLOUD"
	DSCC_ZONE            = "dsccZone"
	ALLETRA_MP_X10K_SNO  = "clusterSerialNumber"
	ENDPOINT             = "endpoint"
	ON_PREM_CLOUD_CA     = "onPremCloudCA"
	RETRY_ATTEMPT        = 3
	PROXY                = "PROXY"
)

// IAMCredentials defines credentials to access DSCC through GLCP API user.
type IAMCredentials struct {
	GLCPUser          string
	GLCPUserSecretKey string
	GLCPCommonCloud   string
	DSCCZone          string
	SystemId          string
	Endpoint          string
	Proxy             string
	OnPremCloudCA     string // Base64 encoded CA certificate for on-premise DSCC
}

// BucketRequest defines the structure for bucket creation options like versioning, locking, and compression.
type BucketRequest struct {
	Compression     string `json:"Compression,omitempty"`
	Versioning      string `json:"Versioning,omitempty"`
	Locking         string `json:"Locking,omitempty"`
	RetentionMode   string `json:"RetentionMode,omitempty"`
	ObjectLockDays  int    `json:"ObjectLockDays,omitempty"`
	ObjectLockYears int    `json:"ObjectLockYears,omitempty"`
}

// SpaceQuota defines the quota type and limit for a bucket.
type SpaceQuota struct {
	QuotaType     string `json:"QuotaType"`
	QuotaLimitMiB int    `json:"QuotaLimitMiB"`
}

// CreateBucketRequest defines the structure for creating a bucket with various configurations.
type CreateBucketRequest struct {
	LocationConstraint string `json:"LocationConstraint"`
	Compression        string `json:"Compression,omitempty"`
	BucketPolicy       string `json:"BucketPolicy"`
	Versioning         string `json:"Versioning,omitempty"`
	ObjectLockEnabled  string `json:"ObjectLockEnabled,omitempty"`
	ObjectLockMode     string `json:"ObjectLockMode,omitempty"`
	ObjectLockDays     int    `json:"ObjectLockDays,omitempty"`
	ObjectLockYears    int    `json:"ObjectLockYears,omitempty"`
}

// ObjectLockConfiguration represents the XML structure for object lock configuration.
type ObjectLockConfiguration struct {
	XMLName           xml.Name        `xml:"ObjectLockConfiguration"`
	Xmlns             string          `xml:"xmlns,attr"`
	ObjectLockEnabled string          `xml:"ObjectLockEnabled"`
	Rule              *ObjectLockRule `xml:"Rule,omitempty"`
}

// ObjectLockRule defines the default retention rule for object locking.
type ObjectLockRule struct {
	DefaultRetention ObjectLockDefaultRetention `xml:"DefaultRetention"`
}

// ObjectLockDefaultRetention defines the retention mode and duration for object locking.
type ObjectLockDefaultRetention struct {
	Mode  string `xml:"Mode,omitempty"`
	Days  int    `xml:"Days,omitempty"`
	Years int    `xml:"Years,omitempty"`
}

// Feature represents a feature toggle with enabled or disabled states.
type Feature string

const (
	FeatureEnabled  Feature = "Enabled"
	FeatureDisabled Feature = "Disabled"
)
