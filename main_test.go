package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ---------- Test the HTTP handler ----------
func TestAnalyzeHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   Response
	}{
		{
			name:           "valid sentence",
			requestBody:    `{"sentence":"Hello Go"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Words: 2, Vowels: 3, Consonants: 4},
		},
		{
			name:           "empty sentence",
			requestBody:    `{"sentence":""}`,
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Words: 0, Vowels: 0, Consonants: 0},
		},
		{
			name:           "invalid JSON",
			requestBody:    `{"sentence":`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	//* ---------- Check body if request  ----------
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/analyze", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			analyzeHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK {
				var body Response
				if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if body != tt.expectedBody {
					t.Errorf("expected body %v, got %v", tt.expectedBody, body)
				}
			}
		})
	}
}
