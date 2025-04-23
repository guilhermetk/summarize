package test

// MockProvider implements the Provider interface for testing
type MockProvider struct {
	SummarizeFunc func(string) string
}

// Summarize is the mock implementation of the Provider interface
func (m *MockProvider) Summarize(text string) string {
	if m.SummarizeFunc != nil {
		return m.SummarizeFunc(text)
	}
	return "mock summary: " + text
}
