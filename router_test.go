package stremigo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetHeaders(t *testing.T) {
	tests := []struct {
		name        string
		expected    map[string]string
		notExpected []string // optional: headers not expected to be included
	}{
		{
			name: "default headers set",
			expected: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type",
				"Content-Type":                 "application/json",
			},
			notExpected: nil,
		},
		{
			name:        "no unexpected headers",
			expected:    nil,
			notExpected: []string{"Custom-Header", "X-Some-Header"},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Use a ResponseRecorder to capture the headers
				rr := &httptest.ResponseRecorder{HeaderMap: http.Header{}}

				// Call the setHeaders function
				setHeaders(rr)

				// Check expected headers
				for key, value := range tt.expected {
					if got := rr.Header().Get(key); got != value {
						t.Errorf("Header %q = %q, want %q", key, got, value)
					}
				}

				// Check that not expected headers are not present
				for _, header := range tt.notExpected {
					if got := rr.Header().Get(header); got != "" {
						t.Errorf("Header %q was not expected but was present with value %q", header, got)
					}
				}
			},
		)
	}
}

func TestIsEnabledEndpoint(t *testing.T) {
	// Table driven tests
	tests := []struct {
		name     string
		endpoint string
		want     bool
	}{
		{
			name:     "Valid endpoint - " + PathManifest,
			endpoint: PathManifest,
			want:     true,
		},
		{
			name:     "Valid endpoint - " + PathCatalog,
			endpoint: PathCatalog,
			want:     true,
		},
		{
			name:     "Valid endpoint - " + PathMeta,
			endpoint: PathMeta,
			want:     true,
		},
		{
			name:     "Valid endpoint - " + PathStream,
			endpoint: PathStream,
			want:     true,
		},
		{
			name:     "Valid endpoint - " + PathSubtitles,
			endpoint: PathSubtitles,
			want:     true,
		},
		{
			name:     "Valid endpoint - " + PathConfigure,
			endpoint: PathConfigure,
			want:     true,
		},
		{
			name:     "invalid endpoint",
			endpoint: "invalid",
			want:     false,
		},
		{
			name:     "empty endpoint",
			endpoint: "",
			want:     false,
		},
		{
			name:     "partial match",
			endpoint: "strea",
			want:     false, // assuming "/api" is not in EnabledEndpoints
		},
		{
			name:     "additional slashes",
			endpoint: "/" + PathStream,
			want:     false, // assuming exact match is required
		},
		{
			name:     "additional slashes 2",
			endpoint: PathStream + "/",
			want:     false, // assuming exact match is required
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := isEnabledEnpoint(tt.endpoint)
				if got != tt.want {
					t.Errorf("isEnabledEnpoint(%q) = %v; want %v", tt.endpoint, got, tt.want)
				}
			},
		)
	}
}
