// © Copyright Hewlett Packard Enterprise Development LP

// Package utils
package utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// S3ErrorResponse represents the standard S3 XML error response.
type S3ErrorResponse struct {
	XMLName   xml.Name `xml:"Error"`
	Code      string   `xml:"Code"`
	Message   string   `xml:"Message"`
	Resource  string   `xml:"Resource"`
	RequestID string   `xml:"RequestId"`
}

// ParseS3ErrorResponse attempts to parse a raw S3 XML error body into a
// concise, human-readable string. If parsing fails it falls back to returning
// the raw body (truncated to 256 chars) so the error is still actionable.
func ParseS3ErrorResponse(body []byte) string {
	var s3Err S3ErrorResponse
	if err := xml.Unmarshal(body, &s3Err); err == nil && s3Err.Code != "" {
		if s3Err.Message != "" {
			return fmt.Sprintf("S3 error %s: %s", s3Err.Code, s3Err.Message)
		}
		return fmt.Sprintf("S3 error %s", s3Err.Code)
	}
	// Fallback: return truncated raw body
	raw := strings.TrimSpace(string(body))
	if len(raw) > 256 {
		raw = raw[:256] + "..."
	}
	return raw
}

// IAMErrorResponse represents the JSON error response from DSCC/IAM APIs.
type IAMErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode string `json:"errorCode"`
}

// ParseIAMErrorResponse attempts to parse a raw DSCC/IAM JSON error body into a
// concise, human-readable string. If parsing fails it falls back to returning
// the raw body (truncated to 256 chars) so the error is still actionable.
func ParseIAMErrorResponse(body []byte) string {
	var iamErr IAMErrorResponse
	if err := json.Unmarshal(body, &iamErr); err == nil && (iamErr.ErrorCode != "" || iamErr.Error != "") {
		if iamErr.ErrorCode != "" && iamErr.Error != "" {
			return fmt.Sprintf("DSCC error %s: %s", iamErr.ErrorCode, iamErr.Error)
		}
		if iamErr.ErrorCode != "" {
			return fmt.Sprintf("DSCC error %s", iamErr.ErrorCode)
		}
		return fmt.Sprintf("DSCC error: %s", iamErr.Error)
	}
	// Fallback: return truncated raw body
	raw := strings.TrimSpace(string(body))
	if len(raw) > 256 {
		raw = raw[:256] + "..."
	}
	return raw
}

// BodyExtractor is implemented by errors that carry a raw response body
// (e.g. the SDK's GenericOpenAPIError). This avoids a direct dependency
// on the SDK package.
type BodyExtractor interface {
	Body() []byte
}

// FormatIAMError enriches a DSCC/IAM API error with the parsed response body.
// If the error implements BodyExtractor (e.g. GenericOpenAPIError), the body is
// parsed via ParseIAMErrorResponse for a human-readable message. Otherwise the
// original error string is returned as-is.
func FormatIAMError(action string, err error) error {
	if err == nil {
		return nil
	}
	if be, ok := err.(BodyExtractor); ok {
		if body := be.Body(); len(body) > 0 {
			parsed := ParseIAMErrorResponse(body)
			return fmt.Errorf("%s: %s", action, parsed)
		}
	}
	return fmt.Errorf("%s: %s", action, err.Error())
}

// ToGRPCStatusError inspects the error message and returns a gRPC status error
// with the appropriate error code and a concise, actionable description.
//
// Mapping rules (evaluated in order – S3 error codes first, then general):
//
//	S3: AccessDenied, AccountProblem,
//	    InvalidAccessKeyId, SignatureDoesNotMatch  → codes.PermissionDenied
//	S3: BucketAlreadyExists,
//	    BucketAlreadyOwnedByYou                    → codes.AlreadyExists
//	S3: NoSuchBucket, NoSuchKey                    → codes.NotFound
//	S3: BucketNotEmpty                             → codes.FailedPrecondition
//	IAM: DOWNSTREAM_SERVICE_ERROR, empty token,
//	     401 Unauthorized                          → codes.Unauthenticated
//	General: "already exists"                      → codes.AlreadyExists
//	General: "permission denied", "forbidden",
//	         "403", "accessdenied"                 → codes.PermissionDenied
//	General: "invalid", "required",
//	         "unsupported"                         → codes.InvalidArgument
//	General: "not found", "404",
//	         "nosuchbucket", "nosuchkey"           → codes.NotFound
//	General: "not empty", "bucketnotempty"         → codes.FailedPrecondition
//	General: connection / timeout errors           → codes.Unavailable
//	everything else                                → codes.Internal
func ToGRPCStatusError(msg string, err error) error {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	combined := strings.ToLower(msg + " " + errMsg)

	var code codes.Code
	switch {
	// --- S3 & general: already exists ---
	case strings.Contains(combined, "bucketalreadyexists") ||
		strings.Contains(combined, "bucketalreadyownedbyyou") ||
		strings.Contains(combined, "already exists") ||
		strings.Contains(combined, "already present"):
		code = codes.AlreadyExists

	// --- IAM/GLCP: authentication (token / credentials) ---
	case strings.Contains(combined, "empty access token") ||
		strings.Contains(combined, "empty token received") ||
		strings.Contains(combined, "failed to generate token") ||
		strings.Contains(combined, "401 unauthorized") ||
		strings.Contains(combined, "token rest api returned status 401") ||
		strings.Contains(combined, "downstream_service_error") ||
		strings.Contains(combined, "unauthenticated"):
		code = codes.Unauthenticated

	// --- S3 & general: permission / auth ---
	case strings.Contains(combined, "accessdenied") ||
		strings.Contains(combined, "invalidaccesskeyid") ||
		strings.Contains(combined, "signaturedoesnotmatch") ||
		strings.Contains(combined, "accountproblem") ||
		strings.Contains(combined, "permission denied") ||
		strings.Contains(combined, "forbidden") ||
		strings.Contains(combined, "403") ||
		strings.Contains(combined, "token rest api returned status 403"):
		code = codes.PermissionDenied

	// --- S3 & general: not found ---
	case strings.Contains(combined, "nosuchbucket") ||
		strings.Contains(combined, "nosuchkey") ||
		strings.Contains(combined, "not found") ||
		strings.Contains(combined, "err_resource_not_found") ||
		strings.Contains(combined, "404"):
		code = codes.NotFound

	// --- S3 & general: precondition (e.g. bucket not empty) ---
	case strings.Contains(combined, "bucketnotempty") ||
		strings.Contains(combined, "not empty"):
		code = codes.FailedPrecondition

	// --- connectivity / timeout ---
	case strings.Contains(combined, "connection refused") ||
		strings.Contains(combined, "no such host") ||
		strings.Contains(combined, "timeout") ||
		strings.Contains(combined, "dial tcp") ||
		strings.Contains(combined, "503 service unavailable") ||
		strings.Contains(combined, "token rest api returned status 503") ||
		strings.Contains(combined, "verify the connectivity"):
		code = codes.Unavailable

	// --- general: invalid input ---
	case strings.Contains(combined, "invalid") ||
		strings.Contains(combined, "required") ||
		strings.Contains(combined, "unsupported"):
		code = codes.InvalidArgument

	default:
		code = codes.Internal
	}

	// Build a concise, actionable message for the Kubernetes event.
	detail := msg
	if err != nil {
		detail = msg + ": " + err.Error()
	}
	if hint := actionHint(code); hint != "" {
		detail = detail + " — " + hint
	}
	return status.Error(code, detail)
}

// actionHint returns a short, user-facing remediation hint for common error
// codes so that Kubernetes events are immediately actionable.
func actionHint(code codes.Code) string {
	switch code {
	case codes.Unauthenticated:
		return "verify GLCP API client credentials and token endpoint in the BucketAccessClass secret"
	case codes.PermissionDenied:
		return "verify credentials and access policies in Alletra Storage MP X10000"
	case codes.Unavailable:
		return "re-check connectivity to Alletra Storage MP X10000"
	case codes.AlreadyExists:
		return "resource already exists, no action needed if this is expected"
	case codes.NotFound:
		return "resource not found, verify the resource name and backend state"
	case codes.FailedPrecondition:
		return "precondition not met (e.g. bucket is not empty), resolve before retrying"
	default:
		return ""
	}
}
