// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Package utils
package utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	ACCESS_POLICY_PREFIX = "acp_"
	USER_PREFIX          = "user_"
	//KEYS in glcp secret
	GLCP_USER_CLIENTID   = "glcpUserClientId"
	GLCP_USER_SECRET_KEY = "glcpUserSecretKey"
	GLCP_WORKSPACE_ID    = "glcpWorkspaceId"
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
	GLCPWorkspaceId   string
	GLCPCommonCloud   string
	DSCCZone          string
	SystemId          string
	Endpoint          string
	Proxy             string
	OnPremCloudCA     string // Base64 encoded CA certificate for on-premise DSCC
}

// BucketRequest carries every BucketClass parameter consumed by bucket
// creation.
//
// It plays two distinct roles:
//
//  1. Internal: a single validated value passed from parseBucketParams into
//     DriverCreateBucket. Every field is read by Go field access (the
//     retention trio is forwarded to SetObjectLockConfiguration, which is
//     a separate REST call from CreateBucket).
//
//  2. Wire: the CreateBucket REST endpoint accepts JSON of the form
//     {"Compression": <bool>, "Versioning": "<Enabled|Disabled>", "Locking":
//     "<Enabled|Disabled>"} and rejects any other field or type with HTTP
//     400 InvalidRequest. In particular Compression is a JSON boolean, not
//     the string "Enabled"/"Disabled" used everywhere else.
//
// To keep field types convenient for internal use (all three feature flags
// share the Feature enum) while honouring the wire contract, the legacy
// wire format is produced by the MarshalJSON method below rather than by
// struct tags. The retention trio is driver-internal only and is omitted
// from the wire entirely.
type BucketRequest struct {
	Compression     Feature
	Versioning      Feature
	Locking         Feature
	RetentionMode   RetentionMode
	ObjectLockDays  int
	ObjectLockYears int
}

// MarshalJSON emits the exact wire payload accepted by the HomeFleet
// CreateBucket REST endpoint:
//
//	{"Compression": <bool>, "Versioning": "<value>", "Locking": "<value>"}
//
// Compression is encoded as a JSON boolean (true when FeatureEnabled).
// Versioning and Locking are encoded as strings and omitted when empty.
// The retention-related fields on BucketRequest are intentionally absent
// from the wire payload — they drive a separate SetObjectLockConfiguration
// call in DriverCreateBucket.
func (b BucketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Compression bool   `json:"Compression"`
		Versioning  string `json:"Versioning,omitempty"`
		Locking     string `json:"Locking,omitempty"`
	}{
		Compression: b.Compression == FeatureEnabled,
		Versioning:  string(b.Versioning),
		Locking:     string(b.Locking),
	})
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

// --- Generic validated string-enum support ---------------------------------
//
// stringEnum constrains a type to any concrete string-based named type, so we
// can build one reusable registry that performs parsing, validation, JSON
// unmarshalling and parameter lookup for every enum in this package.
//
// Adding a new enum requires only:
//  1. Declaring `type Foo string` and its constants.
//  2. Declaring `var fooEnum = enumSet[Foo]{name: "Foo", values: []Foo{...}}`.
//  3. Wiring the four one-line methods/helpers (IsValid / Parse / UnmarshalJSON
//     / FromParams) — all delegate to the generic helper, no new logic.
type stringEnum interface{ ~string }

// unmarshalStringJSON is the shared JSON-string decode + validate adapter for
// any named string type. It decodes the JSON string literal first (handling
// quotes and escapes, rejecting non-string tokens), then defers to the
// supplied parse function for type-specific validation. typeName is used
// only for error messages.
//
// Used by every UnmarshalJSON method in this file so the decode step is
// implemented exactly once.
func unmarshalStringJSON[T stringEnum](data []byte, target *T, parse func(string) (T, error), typeName string) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("%s: expected JSON string: %w", typeName, err)
	}
	parsed, err := parse(s)
	if err != nil {
		return err
	}
	*target = parsed
	return nil
}

// lookupStringParam is the shared "find a key in a BucketClass parameters
// map (case-insensitive) and validate its value" adapter. It returns the
// zero value (with no error) when the key is absent.
//
// When parse rejects the value, the error is formatted in the uniform
// user-facing shape:
//
//	Invalid bucket parameters: <key>="<value>" is invalid; use <hint> in BucketClass and recreate BucketClaim
//
// allowedHint is a short human-readable description of the accepted values
// (e.g. "Enabled or Disabled", "values like 30d, 2m or 1y"). The technical
// detail from parse is intentionally discarded so the message surfaced as a
// Kubernetes event stays short and actionable.
//
// Used by every *FromParams helper in this file so the lookup step and the
// error shape are implemented exactly once.
func lookupStringParam[T stringEnum](param map[string]string, key string,
	parse func(string) (T, error), allowedHint string) (T, error) {
	var zero T
	for k, v := range param {
		if strings.EqualFold(k, key) {
			parsed, err := parse(v)
			if err != nil {
				return zero, fmt.Errorf(
					"Invalid bucket parameters: %s=%q is invalid; use %s in BucketClass and recreate BucketClaim",
					key, v, allowedHint)
			}
			return parsed, nil
		}
	}
	return zero, nil
}

// enumSet is the single source of truth for a validated string enum: it owns
// the human-readable type name (used in error messages) and the registry of
// declared members. All behaviour lives here so individual enum types do not
// duplicate the loop / decode / lookup code.
type enumSet[T stringEnum] struct {
	name   string
	values []T
}

// parse converts a raw string into a T in a case-insensitive manner and
// returns an error if the value is not one of the declared members. An empty
// input returns the zero value with no error so that omitempty fields and
// absent map keys remain valid.
func (e enumSet[T]) parse(s string) (T, error) {
	var zero T
	if s == "" {
		return zero, nil
	}
	for _, v := range e.values {
		if strings.EqualFold(s, string(v)) {
			return v, nil
		}
	}
	return zero, fmt.Errorf("invalid %s value %q: must be one of %s",
		e.name, s, e.join())
}

// isValid reports whether v is the zero value or one of the declared members.
func (e enumSet[T]) isValid(v T) bool {
	if v == "" {
		return true
	}
	for _, x := range e.values {
		if v == x {
			return true
		}
	}
	return false
}

// unmarshalJSON delegates to the shared adapter, passing the enum's own
// parse method as the validator.
func (e enumSet[T]) unmarshalJSON(data []byte, target *T) error {
	return unmarshalStringJSON(data, target, e.parse, e.name)
}

// fromParams delegates to the shared adapter, passing the enum's own parse
// method as the validator and a human-readable hint listing the accepted
// values for user-facing error messages.
func (e enumSet[T]) fromParams(param map[string]string, key string) (T, error) {
	return lookupStringParam(param, key, e.parse, e.allowedHint())
}

// join renders the declared enum members for error messages,
// e.g. `"Enabled", "Disabled"`.
func (e enumSet[T]) join() string {
	parts := make([]string, len(e.values))
	for i, v := range e.values {
		parts[i] = fmt.Sprintf("%q", string(v))
	}
	return strings.Join(parts, ", ")
}

// allowedHint renders the declared enum members in a form suitable for
// user-facing messages: "Enabled or Disabled" for two values,
// "A, B or C" for three or more.
func (e enumSet[T]) allowedHint() string {
	parts := make([]string, len(e.values))
	for i, v := range e.values {
		parts[i] = string(v)
	}
	switch len(parts) {
	case 0:
		return ""
	case 1:
		return parts[0]
	case 2:
		return parts[0] + " or " + parts[1]
	default:
		return strings.Join(parts[:len(parts)-1], ", ") + " or " + parts[len(parts)-1]
	}
}

// --- Feature enum (Compression / Versioning / Locking) ---------------------

// Feature represents a feature toggle with enabled or disabled states.
type Feature string

const (
	FeatureEnabled  Feature = "Enabled"
	FeatureDisabled Feature = "Disabled"
)

var featureEnum = enumSet[Feature]{
	name:   "Feature",
	values: []Feature{FeatureEnabled, FeatureDisabled},
}

// IsValid reports whether f is one of the declared Feature values (or empty).
func (f Feature) IsValid() bool { return featureEnum.isValid(f) }

// ParseFeature converts a raw string into a Feature, rejecting any value
// outside the declared enum members.
func ParseFeature(s string) (Feature, error) { return featureEnum.parse(s) }

// UnmarshalJSON enforces the enum constraint when decoding JSON.
func (f *Feature) UnmarshalJSON(data []byte) error {
	return featureEnum.unmarshalJSON(data, f)
}

// FeatureFromParams looks up key in param (case-insensitive) and returns a
// validated Feature. Out-of-enum values are rejected with an error naming
// the offending key.
func FeatureFromParams(param map[string]string, key string) (Feature, error) {
	return featureEnum.fromParams(param, key)
}

// --- RetentionMode enum ----------------------------------------------------

// RetentionMode selects the S3 object-lock mode applied to a bucket.
//   - GOVERNANCE: users without special permissions cannot overwrite or
//     delete locked objects.
//   - COMPLIANCE: no user (including root) can overwrite or delete locked
//     objects until the retention period expires.
type RetentionMode string

const (
	RetentionModeGovernance RetentionMode = "GOVERNANCE"
	RetentionModeCompliance RetentionMode = "COMPLIANCE"
)

var retentionModeEnum = enumSet[RetentionMode]{
	name:   "RetentionMode",
	values: []RetentionMode{RetentionModeGovernance, RetentionModeCompliance},
}

// IsValid reports whether r is one of the declared RetentionMode values.
func (r RetentionMode) IsValid() bool { return retentionModeEnum.isValid(r) }

// ParseRetentionMode converts a raw string into a RetentionMode, rejecting
// any value outside the declared enum members.
func ParseRetentionMode(s string) (RetentionMode, error) {
	return retentionModeEnum.parse(s)
}

// UnmarshalJSON enforces the enum constraint when decoding JSON.
func (r *RetentionMode) UnmarshalJSON(data []byte) error {
	return retentionModeEnum.unmarshalJSON(data, r)
}

// RetentionModeFromParams looks up key in param (case-insensitive) and
// returns a validated RetentionMode.
func RetentionModeFromParams(param map[string]string, key string) (RetentionMode, error) {
	return retentionModeEnum.fromParams(param, key)
}

// --- RetentionInterval (regex-validated, not a finite enum) ----------------

// RetentionInterval expresses the default retention period applied to
// locked objects, written as a positive integer followed by a unit suffix:
//
//	"30d" — 30 days
//	"2m"  — 2 "months" (interpreted as 2 * 30 = 60 days)
//	"1y"  — 1 year (mapped to the ObjectLockYears field on the wire)
//
// Months are intentionally normalised to 30-day blocks so the backend only
// has to reason about days/years.
type RetentionInterval string

// retentionIntervalRe matches a positive natural number (no leading zero)
// followed by exactly one of d/m/y. Whitespace and signs are rejected.
// Submatch 1 captures the integer; submatch 2 captures the unit.
var retentionIntervalRe = regexp.MustCompile(`^([1-9][0-9]*)([dmy])$`)

// ParseRetentionInterval validates the textual form of a retention interval.
// An empty input yields the zero value with no error.
func ParseRetentionInterval(s string) (RetentionInterval, error) {
	if s == "" {
		return "", nil
	}
	if !retentionIntervalRe.MatchString(s) {
		return "", fmt.Errorf(
			"invalid RetentionInterval %q: expected positive integer followed by 'd', 'm' or 'y' (e.g., \"30d\", \"2m\", \"1y\")",
			s)
	}
	return RetentionInterval(s), nil
}

// UnmarshalJSON enforces the format constraint when decoding JSON.
func (r *RetentionInterval) UnmarshalJSON(data []byte) error {
	return unmarshalStringJSON(data, r, ParseRetentionInterval, "RetentionInterval")
}

// RetentionIntervalFromParams looks up key in param (case-insensitive) and
// returns a validated RetentionInterval.
func RetentionIntervalFromParams(param map[string]string, key string) (RetentionInterval, error) {
	return lookupStringParam(param, key, ParseRetentionInterval, "values like 30d, 2m or 1y")
}

// ToDaysYears decomposes the interval into (days, years) using the rules in
// the type's doc comment. The zero value (empty string) returns (0, 0, nil)
// — it represents "interval not set" and is a valid input.
//
// A non-empty value that does not match the validation regex returns an
// error. This can only happen if a caller constructed a RetentionInterval
// outside ParseRetentionInterval / UnmarshalJSON / RetentionIntervalFromParams
// (e.g. by direct string conversion). Surfacing it rather than silently
// returning (0, 0) prevents an unvalidated interval from being indistinguishable
// from an absent one.
func (r RetentionInterval) ToDaysYears() (days, years int, err error) {
	if r == "" {
		return 0, 0, nil
	}
	m := retentionIntervalRe.FindStringSubmatch(string(r))
	if m == nil {
		return 0, 0, fmt.Errorf(
			"invalid RetentionInterval %q: value was not produced by ParseRetentionInterval",
			string(r))
	}
	n, convErr := strconv.Atoi(m[1])
	if convErr != nil || n <= 0 {
		return 0, 0, fmt.Errorf(
			"invalid RetentionInterval %q: integer component is not a positive number",
			string(r))
	}
	switch m[2] {
	case "d":
		return n, 0, nil
	case "m":
		return n * 30, 0, nil
	case "y":
		return 0, n, nil
	}
	return 0, 0, fmt.Errorf(
		"invalid RetentionInterval %q: unrecognised unit %q",
		string(r), m[2])
}

// MaskName masks the UUID portion of a name (e.g. user_ba-edummy-forthe-sake-ofex-ampled918fa73)
// keeping the prefix and last 4 characters visible, replacing the rest with 'x'.
// Example: "user_ba-adummy-forthe-sake-ofex-mplein918fa73" -> "user_ba-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxfa73"
func MaskName(name string) string {
	l := len(name)
	if l <= 1 {
		return name
	}
	if l > 6 {
		return replaceWithXx(name, 2, l-4)
	}
	return replaceWithXx(name, 0, l-1)
}

func replaceWithXx(str string, begin, end int) string {
	chars := strings.Split(str, "")
	for i, v := range chars[begin:end] {
		if v != "-" && v != "_" {
			chars[i+begin] = "x"
		}
	}
	return strings.Join(chars, "")
}
