package openaiclient

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type OpenAiClient struct {
	client *openai.Client
	model  string
}

func New(model string, url string, key string) *OpenAiClient {
	client := openai.NewClient(
		option.WithBaseURL(url),
		option.WithAPIKey(key),
	)

	return &OpenAiClient{
		client: &client,
		model:  model,
	}
}

func (c *OpenAiClient) DoRequest(ctx context.Context, systemPrompt string, userInput string) (string, error) {
	const op = "client.openai.OpenAiClient.DoRequest"

	chatCompletion, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model:       c.model,
		Temperature: openai.Float(0.0),
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(userInput),
		},
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if len(chatCompletion.Choices) == 0 {
		return "", fmt.Errorf("%s: empty response from the model", op)
	}

	response := chatCompletion.Choices[0].Message.Content

	return response, nil
}
