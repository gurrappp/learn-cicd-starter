package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
		errorContains string
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key-123"},
			},
			expectedKey:   "test-api-key-123",
			expectedError: nil,
		},
		{
			name: "malformed Authorization header - No Space",
			headers: http.Header{
				"Authorization": []string{"ApiKeytest-api-key-123"},
			},
			expectedKey:   "",
			errorContains: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			// Check the returned key
			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() got key = %v, want %v", key, tt.expectedKey)
			}

			// Check the error
			if tt.expectedError != nil {
				if err != tt.expectedError {
					t.Errorf("GetAPIKey() got error = %v, want %v", err, tt.expectedError)
				}
			} else if tt.errorContains != "" {
				if err == nil {
					t.Errorf("GetAPIKey() expected error containing %q but got nil", tt.errorContains)
				} else if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("GetAPIKey() error = %v, should contain %q", err, tt.errorContains)
				}
			} else if err != nil {
				t.Errorf("GetAPIKey() unexpected error: %v", err)
			}
		})
	}
}
