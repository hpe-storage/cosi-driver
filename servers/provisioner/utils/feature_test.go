// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Package utils
package utils

import (
	"encoding/json"
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

// TestBucketRequest_UnmarshalJSON ensures the enum constraint flows through
// to fields of type Feature on real request structs.
func TestBucketRequest_UnmarshalJSON(t *testing.T) {
	t.Run("valid payload unmarshals", func(t *testing.T) {
		var br BucketRequest
		if err := json.Unmarshal([]byte(`{"Compression":"Enabled","Versioning":"Enabled","Locking":"Disabled"}`), &br); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if br.Compression != FeatureEnabled || br.Versioning != FeatureEnabled || br.Locking != FeatureDisabled {
			t.Errorf("unexpected values: %+v", br)
		}
	})
	t.Run("invalid Compression rejected", func(t *testing.T) {
		var br BucketRequest
		if err := json.Unmarshal([]byte(`{"Compression":"On"}`), &br); err == nil {
			t.Fatal("expected error, got nil")
		}
	})
	t.Run("invalid Versioning rejected", func(t *testing.T) {
		var br BucketRequest
		if err := json.Unmarshal([]byte(`{"Versioning":"yes"}`), &br); err == nil {
			t.Fatal("expected error, got nil")
		}
	})
	t.Run("invalid Locking rejected", func(t *testing.T) {
		var br BucketRequest
		if err := json.Unmarshal([]byte(`{"Locking":"off"}`), &br); err == nil {
			t.Fatal("expected error, got nil")
		}
	})
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
