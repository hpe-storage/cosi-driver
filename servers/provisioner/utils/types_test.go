// © Copyright Hewlett Packard Enterprise Development LP

// Package utils
package utils

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestFeature_IsValid(t *testing.T) {
	tests := []struct {
		name  string
		value Feature
		want  bool
	}{
		{"empty is valid (zero value)", Feature(""), true},
		{"enabled is valid", FeatureEnabled, true},
		{"disabled is valid", FeatureDisabled, true},
		{"lowercase enabled is invalid", Feature("enabled"), false},
		{"arbitrary string is invalid", Feature("On"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.IsValid(); got != tt.want {
				t.Errorf("Feature(%q).IsValid() = %v, want %v", string(tt.value), got, tt.want)
			}
		})
	}
}

func TestParseFeature(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Feature
		wantErr bool
	}{
		{"empty input returns zero value", "", "", false},
		{"exact Enabled", "Enabled", FeatureEnabled, false},
		{"exact Disabled", "Disabled", FeatureDisabled, false},
		{"case-insensitive enabled", "enabled", FeatureEnabled, false},
		{"case-insensitive DISABLED", "DISABLED", FeatureDisabled, false},
		{"unknown value rejected", "On", "", true},
		{"boolean string rejected", "true", "", true},
		{"whitespace rejected", " Enabled", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFeature(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseFeature(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseFeature(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestFeature_UnmarshalJSON(t *testing.T) {
	type payload struct {
		Compression Feature `json:"Compression,omitempty"`
	}
	tests := []struct {
		name    string
		body    string
		want    Feature
		wantErr bool
	}{
		{"valid Enabled", `{"Compression":"Enabled"}`, FeatureEnabled, false},
		{"valid Disabled", `{"Compression":"Disabled"}`, FeatureDisabled, false},
		{"case-insensitive accepted", `{"Compression":"enabled"}`, FeatureEnabled, false},
		{"empty string accepted", `{"Compression":""}`, "", false},
		{"missing field accepted", `{}`, "", false},
		{"invalid string rejected", `{"Compression":"On"}`, "", true},
		{"boolean rejected", `{"Compression":true}`, "", true},
		{"number rejected", `{"Compression":1}`, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p payload
			err := json.Unmarshal([]byte(tt.body), &p)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Unmarshal(%s) error = %v, wantErr %v", tt.body, err, tt.wantErr)
			}
			if !tt.wantErr && p.Compression != tt.want {
				t.Errorf("Unmarshal(%s) Compression = %q, want %q", tt.body, p.Compression, tt.want)
			}
		})
	}
}

// TestFeatureFromParams ensures the unified parameter helper looks up keys
// case-insensitively and rejects out-of-enum values.
func TestFeatureFromParams(t *testing.T) {
	tests := []struct {
		name    string
		param   map[string]string
		key     string
		want    Feature
		wantErr bool
	}{
		{"absent key returns zero value", map[string]string{}, "versioning", "", false},
		{"empty value returns zero value", map[string]string{"versioning": ""}, "versioning", "", false},
		{"valid Enabled", map[string]string{"versioning": "Enabled"}, "versioning", FeatureEnabled, false},
		{"valid Disabled", map[string]string{"compression": "Disabled"}, "compression", FeatureDisabled, false},
		{"case-insensitive key lookup", map[string]string{"Versioning": "Enabled"}, "versioning", FeatureEnabled, false},
		{"case-insensitive value", map[string]string{"locking": "enabled"}, "locking", FeatureEnabled, false},
		{"invalid value rejected", map[string]string{"versioning": "On"}, "versioning", "", true},
		{"boolean-like value rejected", map[string]string{"compression": "true"}, "compression", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FeatureFromParams(tt.param, tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FeatureFromParams() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("FeatureFromParams() = %q, want %q", got, tt.want)
			}
		})
	}
}

// --- RetentionMode -------------------------------------------------------

func TestParseRetentionMode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    RetentionMode
		wantErr bool
	}{
		{"empty returns zero value", "", "", false},
		{"exact GOVERNANCE", "GOVERNANCE", RetentionModeGovernance, false},
		{"exact COMPLIANCE", "COMPLIANCE", RetentionModeCompliance, false},
		{"case-insensitive governance", "governance", RetentionModeGovernance, false},
		{"case-insensitive Compliance", "Compliance", RetentionModeCompliance, false},
		{"unknown rejected", "Strict", "", true},
		{"whitespace rejected", " GOVERNANCE", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRetentionMode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseRetentionMode(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseRetentionMode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestRetentionMode_UnmarshalJSON(t *testing.T) {
	type payload struct {
		RetentionMode RetentionMode `json:"RetentionMode,omitempty"`
	}
	tests := []struct {
		name    string
		body    string
		want    RetentionMode
		wantErr bool
	}{
		{"valid GOVERNANCE", `{"RetentionMode":"GOVERNANCE"}`, RetentionModeGovernance, false},
		{"valid lowercase compliance", `{"RetentionMode":"compliance"}`, RetentionModeCompliance, false},
		{"missing field accepted", `{}`, "", false},
		{"invalid string rejected", `{"RetentionMode":"Strict"}`, "", true},
		{"number rejected", `{"RetentionMode":3}`, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p payload
			err := json.Unmarshal([]byte(tt.body), &p)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Unmarshal(%s) error = %v, wantErr %v", tt.body, err, tt.wantErr)
			}
			if !tt.wantErr && p.RetentionMode != tt.want {
				t.Errorf("Unmarshal(%s) RetentionMode = %q, want %q", tt.body, p.RetentionMode, tt.want)
			}
		})
	}
}

func TestRetentionModeFromParams(t *testing.T) {
	tests := []struct {
		name    string
		param   map[string]string
		key     string
		want    RetentionMode
		wantErr bool
	}{
		{"absent key", map[string]string{}, "retentionMode", "", false},
		{"empty value", map[string]string{"retentionMode": ""}, "retentionMode", "", false},
		{"valid governance", map[string]string{"retentionMode": "GOVERNANCE"}, "retentionMode", RetentionModeGovernance, false},
		{"case-insensitive key", map[string]string{"RetentionMode": "COMPLIANCE"}, "retentionMode", RetentionModeCompliance, false},
		{"invalid value", map[string]string{"retentionMode": "Strict"}, "retentionMode", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RetentionModeFromParams(tt.param, tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("RetentionModeFromParams() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("RetentionModeFromParams() = %q, want %q", got, tt.want)
			}
		})
	}
}

// --- RetentionInterval ---------------------------------------------------

func TestParseRetentionInterval(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    RetentionInterval
		wantErr bool
	}{
		{"empty returns zero", "", "", false},
		{"30d valid", "30d", "30d", false},
		{"2m valid", "2m", "2m", false},
		{"1y valid", "1y", "1y", false},
		{"large day count valid", "365d", "365d", false},
		{"zero rejected", "0d", "", true},
		{"leading zero rejected", "01d", "", true},
		{"negative rejected", "-1d", "", true},
		{"missing suffix rejected", "30", "", true},
		{"wrong suffix rejected", "30h", "", true},
		{"whitespace rejected", " 30d", "", true},
		{"trailing whitespace rejected", "30d ", "", true},
		{"uppercase suffix rejected", "30D", "", true},
		{"empty number rejected", "d", "", true},
		{"double suffix rejected", "30dd", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRetentionInterval(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseRetentionInterval(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseRetentionInterval(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestRetentionInterval_UnmarshalJSON(t *testing.T) {
	type payload struct {
		DefaultRetentionInterval RetentionInterval `json:"defaultRetentionInterval,omitempty"`
	}
	tests := []struct {
		name    string
		body    string
		want    RetentionInterval
		wantErr bool
	}{
		{"valid 30d", `{"defaultRetentionInterval":"30d"}`, "30d", false},
		{"valid 1y", `{"defaultRetentionInterval":"1y"}`, "1y", false},
		{"missing field accepted", `{}`, "", false},
		{"invalid string rejected", `{"defaultRetentionInterval":"forever"}`, "", true},
		{"number rejected", `{"defaultRetentionInterval":30}`, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p payload
			err := json.Unmarshal([]byte(tt.body), &p)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Unmarshal(%s) error = %v, wantErr %v", tt.body, err, tt.wantErr)
			}
			if !tt.wantErr && p.DefaultRetentionInterval != tt.want {
				t.Errorf("Unmarshal(%s) = %q, want %q", tt.body, p.DefaultRetentionInterval, tt.want)
			}
		})
	}
}

func TestRetentionInterval_ToDaysYears(t *testing.T) {
	tests := []struct {
		name      string
		input     RetentionInterval
		wantDays  int
		wantYears int
		wantErr   bool
	}{
		{"empty is valid zero", "", 0, 0, false},
		{"30d", "30d", 30, 0, false},
		{"2m -> 60 days", "2m", 60, 0, false},
		{"1m -> 30 days", "1m", 30, 0, false},
		{"1y -> 1 year", "1y", 0, 1, false},
		{"5y -> 5 years", "5y", 0, 5, false},
		// Values that should never appear in production (they can only be
		// constructed by skipping ParseRetentionInterval). The contract is
		// that ToDaysYears surfaces them rather than returning (0, 0).
		{"malformed garbage rejected", "garbage", 0, 0, true},
		{"unit-only rejected", "d", 0, 0, true},
		{"unknown unit rejected", "5x", 0, 0, true},
		{"leading zero rejected", "05d", 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, y, err := tt.input.ToDaysYears()
			if (err != nil) != tt.wantErr {
				t.Fatalf("(%q).ToDaysYears() err = %v, wantErr=%v", tt.input, err, tt.wantErr)
			}
			if d != tt.wantDays || y != tt.wantYears {
				t.Errorf("(%q).ToDaysYears() = (%d, %d), want (%d, %d)",
					tt.input, d, y, tt.wantDays, tt.wantYears)
			}
		})
	}
}

func TestRetentionIntervalFromParams(t *testing.T) {
	tests := []struct {
		name    string
		param   map[string]string
		key     string
		want    RetentionInterval
		wantErr bool
	}{
		{"absent key", map[string]string{}, "defaultRetentionInterval", "", false},
		{"valid 30d", map[string]string{"defaultRetentionInterval": "30d"}, "defaultRetentionInterval", "30d", false},
		{"case-insensitive key", map[string]string{"DefaultRetentionInterval": "1y"}, "defaultRetentionInterval", "1y", false},
		{"invalid value", map[string]string{"defaultRetentionInterval": "forever"}, "defaultRetentionInterval", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RetentionIntervalFromParams(tt.param, tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("RetentionIntervalFromParams() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("RetentionIntervalFromParams() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestBucketRequest_WirePayload guards the contract that CreateBucket's
// JSON wire payload matches what the HomeFleet REST endpoint accepts:
//
//	{"Compression": <bool>, "Versioning": "<value>", "Locking": "<value>"}
//
// In particular Compression is a JSON *boolean* (not the string "Enabled"
// used internally) and the retention-related fields never appear on the
// wire — they drive a separate SetObjectLockConfiguration call. The
// backend rejects any deviation with HTTP 400 InvalidRequest, so this is
// a hard contract worth pinning down.
func TestBucketRequest_WirePayload(t *testing.T) {
	t.Run("all enabled with retention", func(t *testing.T) {
		br := BucketRequest{
			Compression:     FeatureEnabled,
			Versioning:      FeatureEnabled,
			Locking:         FeatureEnabled,
			RetentionMode:   RetentionModeGovernance,
			ObjectLockDays:  30,
			ObjectLockYears: 0,
		}
		out, err := json.Marshal(br)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		got := string(out)
		want := `{"Compression":true,"Versioning":"Enabled","Locking":"Enabled"}`
		if got != want {
			t.Errorf("wire payload mismatch\n got: %s\nwant: %s", got, want)
		}
	})
	t.Run("compression disabled emits boolean false", func(t *testing.T) {
		br := BucketRequest{
			Compression: FeatureDisabled,
			Versioning:  FeatureDisabled,
		}
		out, err := json.Marshal(br)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		got := string(out)
		want := `{"Compression":false,"Versioning":"Disabled"}`
		if got != want {
			t.Errorf("wire payload mismatch\n got: %s\nwant: %s", got, want)
		}
	})
	t.Run("driver-internal fields never appear on the wire", func(t *testing.T) {
		br := BucketRequest{
			Compression:     FeatureEnabled,
			Versioning:      FeatureEnabled,
			Locking:         FeatureEnabled,
			RetentionMode:   RetentionModeCompliance,
			ObjectLockDays:  0,
			ObjectLockYears: 7,
		}
		out, err := json.Marshal(br)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		got := string(out)
		for _, forbidden := range []string{"RetentionMode", "ObjectLockDays", "ObjectLockYears", "COMPLIANCE", "GOVERNANCE"} {
			if strings.Contains(got, forbidden) {
				t.Errorf("wire payload leaked driver-internal token %q: %s", forbidden, got)
			}
		}
	})
}

func TestMaskName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "typical user name with UUID",
			input: "user_ba-edummy-forthe-sake-ofex-ampled918fa73",
			want:  "usxx_xx-xxxxxx-xxxxxx-xxxx-xxxx-xxxxxxxxxfa73",
		},
		{
			name:  "short string exactly 4 chars",
			input: "abcd",
			want:  "xxxd",
		},
		{
			name:  "5 chars masks all but last",
			input: "abcde",
			want:  "xxxxe",
		},
		{
			name:  "6 chars masks all but last",
			input: "abcdef",
			want:  "xxxxxf",
		},
		{
			name:  "7 chars keeps first 2 and last 4",
			input: "abcdefg",
			want:  "abxdefg",
		},
		{
			name:  "short string 3 chars",
			input: "abc",
			want:  "xxc",
		},
		{
			name:  "short string 2 chars",
			input: "ab",
			want:  "xb",
		},
		{
			name:  "single char",
			input: "a",
			want:  "a",
		},
		{
			name:  "preserves hyphens in long string",
			input: "a-b-c-d-e",
			want:  "a-x-x-d-e",
		},
		{
			name:  "preserves underscores in long string",
			input: "a_b_c_d_e",
			want:  "a_x_x_d_e",
		},
		{
			name:  "prefix with underscore",
			input: "acp_12345678",
			want:  "acx_xxxx5678",
		},
		{
			name:  "all hyphens preserved",
			input: "aa----bb",
			want:  "aa----bb",
		},
		{
			name:  "numeric string",
			input: "1234567890",
			want:  "12xxxx7890",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaskName(tt.input)
			if got != tt.want {
				t.Errorf("MaskName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
