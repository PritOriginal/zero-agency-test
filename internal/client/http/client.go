package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct {
	client *http.Client
	model  string
	url    string
	key    string
}

type chatCompletionRequest struct {
	Model       string                  `json:"model"`
	Temperature float64                 `json:"temperature"`
	Messages    []chatCompletionMessage `json:"messages"`
}

type chatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatCompletionResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func New(model string, url string, key string) *HttpClient {
	return &HttpClient{
		client: &http.Client{},
		model:  model,
		url:    url,
		key:    key,
	}
}

func (c *HttpClient) DoRequest(ctx context.Context, systemPrompt string, userInput string) (string, error) {
	const op = "client.openai.HttpClient.DoRequest"

	request := chatCompletionRequest{
		Model:       c.model,
		Temperature: 0.0,
		Messages: []chatCompletionMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userInput},
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	url := fmt.Sprintf("%s/chat/completions", c.url)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.key))

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp chatCompletionResponse
		if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.Error != nil {
			return "", fmt.Errorf("%s: API error (status %d): %s", op, resp.StatusCode, errorResp.Error.Message)
		}
		return "", fmt.Errorf("%s: unexpected response status: %d, body: %s", op, resp.StatusCode, string(body))
	}

	var chatResp chatCompletionResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("%s: empty response from the model", op)
	}

	response := chatResp.Choices[0].Message.Content

	return response, nil
}
