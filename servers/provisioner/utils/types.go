// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Package utils
package utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
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
	Compression     Feature `json:"Compression,omitempty"`
	Versioning      Feature `json:"Versioning,omitempty"`
	Locking         Feature `json:"Locking,omitempty"`
	RetentionMode   string  `json:"RetentionMode,omitempty"`
	ObjectLockDays  int     `json:"ObjectLockDays,omitempty"`
	ObjectLockYears int     `json:"ObjectLockYears,omitempty"`
}

// SpaceQuota defines the quota type and limit for a bucket.
type SpaceQuota struct {
	QuotaType     string `json:"QuotaType"`
	QuotaLimitMiB int    `json:"QuotaLimitMiB"`
}

// CreateBucketRequest defines the structure for creating a bucket with various configurations.
type CreateBucketRequest struct {
	LocationConstraint string  `json:"LocationConstraint"`
	Compression        Feature `json:"Compression,omitempty"`
	BucketPolicy       string  `json:"BucketPolicy"`
	Versioning         Feature `json:"Versioning,omitempty"`
	ObjectLockEnabled  string  `json:"ObjectLockEnabled,omitempty"`
	ObjectLockMode     string  `json:"ObjectLockMode,omitempty"`
	ObjectLockDays     int     `json:"ObjectLockDays,omitempty"`
	ObjectLockYears    int     `json:"ObjectLockYears,omitempty"`
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
// Acts as a validated enum: only FeatureEnabled, FeatureDisabled, or the
// zero value ("") are accepted. Any other value is rejected during JSON
// unmarshalling or when constructed via ParseFeature.
type Feature string

const (
	FeatureEnabled  Feature = "Enabled"
	FeatureDisabled Feature = "Disabled"
)

// IsValid reports whether f is one of the declared Feature values.
// The zero value ("") is considered valid so that omitempty fields can be
// left unset.
func (f Feature) IsValid() bool {
	switch f {
	case "", FeatureEnabled, FeatureDisabled:
		return true
	default:
		return false
	}
}

// ParseFeature converts a raw string into a Feature in a case-insensitive
// manner, returning an error if the value is not one of the declared enum
// members. An empty input returns the zero value with no error.
func ParseFeature(s string) (Feature, error) {
	switch {
	case s == "":
		return "", nil
	case strings.EqualFold(s, string(FeatureEnabled)):
		return FeatureEnabled, nil
	case strings.EqualFold(s, string(FeatureDisabled)):
		return FeatureDisabled, nil
	default:
		return "", fmt.Errorf("invalid Feature value %q: must be %q or %q",
			s, string(FeatureEnabled), string(FeatureDisabled))
	}
}

// UnmarshalJSON enforces the enum constraint when decoding JSON into a
// Feature value. Any string outside the declared set causes the unmarshal
// to fail, so consumers cannot silently accept unknown values.
func (f *Feature) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Feature: expected JSON string: %w", err)
	}
	parsed, err := ParseFeature(s)
	if err != nil {
		return err
	}
	*f = parsed
	return nil
}

// FeatureFromParams looks up key in param (case-insensitive) and converts
// the associated value into a validated Feature. An absent key or empty
// value yields the zero Feature with no error. Any non-empty value that is
// not "Enabled" or "Disabled" (case-insensitive) is rejected with an error
// that names the offending key, so callers do not need to perform any
// additional string-level validation.
func FeatureFromParams(param map[string]string, key string) (Feature, error) {
	for k, v := range param {
		if strings.EqualFold(k, key) {
			f, err := ParseFeature(v)
			if err != nil {
				return "", fmt.Errorf("invalid value for %s: %w", key, err)
			}
			return f, nil
		}
	}
	return "", nil
}
