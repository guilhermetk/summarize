package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/guilhermetk/summarize/internal/handlers/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetSummarize(t *testing.T) {
	tests := []struct {
		name           string
		queryText      string
		mockSummary    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "successful summarization",
			queryText:      "test text",
			mockSummary:    "summary of test text",
			expectedStatus: http.StatusOK,
			expectedBody:   "summary of test text",
		},
		{
			name:           "empty text",
			queryText:      "",
			mockSummary:    "",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()

			// Create URL with properly encoded query parameters
			baseURL := "/summarize"
			params := url.Values{}
			params.Add("text", tt.queryText)
			requestURL := baseURL + "?" + params.Encode()

			req := httptest.NewRequest(http.MethodGet, requestURL, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Create mock provider
			mockProvider := &test.MockProvider{
				SummarizeFunc: func(text string) string {
					return tt.mockSummary
				},
			}

			// Create handler
			h := &SummarizeHandler{
				Provider: mockProvider,
			}

			// Test
			err := h.HandleGetSummarize(c)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.Equal(t, tt.expectedBody, rec.Body.String())
		})
	}
}
