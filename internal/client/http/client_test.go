package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClassify(t *testing.T) {
	tests := []struct {
		name          string
		systemPromt   string
		userInput     string
		apiKey        string
		mockResponse  any
		statusCode    int
		expectedError bool
	}{
		{
			name:        "Ok",
			systemPromt: "test",
			userInput:   "test",
			apiKey:      "test-key",
			mockResponse: map[string]any{
				"choices": []map[string]any{
					{"message": map[string]string{"content": "Привет! Чем помочь?"}},
				},
			},
			statusCode: http.StatusOK,
		},
		{
			name:          "Err-ChatCompletions",
			systemPromt:   "test",
			userInput:     "test",
			apiKey:        "test-key",
			statusCode:    http.StatusOK,
			expectedError: true,
		},
		{
			name:        "Err-EmptyChoices",
			systemPromt: "test",
			userInput:   "test",
			apiKey:      "test-key",
			mockResponse: map[string]any{
				"choices": []map[string]any{},
			},
			statusCode:    http.StatusOK,
			expectedError: true,
		},
		{
			name:        "Err-HTTP-400",
			systemPromt: "test",
			userInput:   "test",
			apiKey:      "test-key",
			statusCode:  http.StatusBadRequest,
			mockResponse: map[string]any{
				"error": map[string]string{"message": "Bad request"},
			},
			expectedError: true,
		},
		{
			name:        "Err-HTTP-401",
			systemPromt: "test",
			userInput:   "test",
			apiKey:      "key",
			statusCode:  http.StatusUnauthorized,
			mockResponse: map[string]any{
				"error": map[string]string{"message": "Unauthorized"},
			},
			expectedError: true,
		},
		{
			name:          "Err-Invalid-JSON",
			systemPromt:   "test",
			userInput:     "test",
			apiKey:        "test-key",
			mockResponse:  "invalid json",
			statusCode:    http.StatusOK,
			expectedError: true,
		},
		{
			name:          "Err-HTTP-500",
			systemPromt:   "test",
			userInput:     "test",
			apiKey:        "test-key",
			statusCode:    http.StatusInternalServerError,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, "application/json", r.Header.Get("Content-Type"))

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				if tt.mockResponse != nil {
					json.NewEncoder(w).Encode(tt.mockResponse)
				}
			}))
			defer server.Close()

			client := New("model", server.URL, tt.apiKey)
			resp, err := client.DoRequest(context.Background(), tt.systemPromt, tt.userInput)
			if !tt.expectedError {
				require.NoError(t, err)
				require.NotNil(t, resp)
			} else {
				require.Error(t, err)
			}
		})
	}
}
