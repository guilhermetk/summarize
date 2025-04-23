package providers

import (
	"os"
	"testing"

	"github.com/guilhermetk/summarize/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestGeminiProvider_Summarize(t *testing.T) {
	config.LoadEnv()

	if os.Getenv("GOOGLE_GEMINI_KEY") == "" {
		t.Skip("Skipping test: GOOGLE_GEMINI_KEY not set")
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic summarization",
			input:    "The quick brown fox jumps over the lazy dog. This is a test sentence that should be summarized.",
			expected: "", // We can't predict the exact output, but it shouldn't be empty
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
	}

	provider := &GeminiProvider{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := provider.Summarize(tt.input)

			if tt.input == "" {
				assert.Empty(t, result, "Empty input should result in empty output")
			} else {
				assert.NotEmpty(t, result, "Non-empty input should result in non-empty output")
				assert.NotEqual(t, tt.input, result, "Summary should be different from input")
				assert.Less(t, len(result), len(tt.input))
			}
		})
	}
}
