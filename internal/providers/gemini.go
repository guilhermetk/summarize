package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
		FinishReason     string  `json:"finishReason"`
		AvgLogprobs      float64 `json:"avgLogprobs"`
		CitationMetadata struct {
			CitationSources []struct {
				StartIndex int     `json:"startIndex"`
				EndIndex   int     `json:"endIndex"`
				URI        *string `json:"uri,omitempty"`
			} `json:"citationSources"`
		} `json:"citationMetadata"`
	} `json:"candidates"`
	UsageMetadata struct {
		PromptTokenCount     int `json:"promptTokenCount"`
		CandidatesTokenCount int `json:"candidatesTokenCount"`
		TotalTokenCount      int `json:"totalTokenCount"`
		PromptTokensDetails  []struct {
			Modality   string `json:"modality"`
			TokenCount int    `json:"tokenCount"`
		} `json:"promptTokensDetails"`
		CandidatesTokensDetails []struct {
			Modality   string `json:"modality"`
			TokenCount int    `json:"tokenCount"`
		} `json:"candidatesTokensDetails"`
	} `json:"usageMetadata"`
	ModelVersion string `json:"modelVersion"`
}

type GeminiProvider struct {
}

func (p *GeminiProvider) Summarize(s string) string {
	return executePost(s)
}

func executePost(s string) string {
	apiKey := os.Getenv("GOOGLE_GEMINI_KEY")
	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s",
		apiKey,
	)

	// Create the request payload
	payload := map[string]any{
		"contents": []map[string]any{
			{
				"parts": []map[string]string{
					{
						"text": "Your job is to summarize any given test, mantaining the meaning of the original one but trying to keep the response as short as possible. Here is the text you should work on: " + s,
					},
				},
			},
		},
	}

	// Encode the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return ""
	}

	// Make the HTTP POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("HTTP request error:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	geminiResponse := new(geminiResponse)
	parseError := json.Unmarshal(body, geminiResponse)

	if parseError != nil {
		log.Print("error parsing gemini response", parseError)
		return ""
	}

	return geminiResponse.Candidates[0].Content.Parts[0].Text
}
