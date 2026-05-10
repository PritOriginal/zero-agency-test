package classifier

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

type СlassifyResponse struct {
	Category string `json:"category"`
}

var classifyResponseExampleStr string

func init() {
	jsonExample := СlassifyResponse{
		Category: "название_тега",
	}
	jsonExampleBytes, err := json.Marshal(jsonExample)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal example classify response: %v", err))
	}
	classifyResponseExampleStr = string(jsonExampleBytes)
}

type Client interface {
	DoRequest(ctx context.Context, systemPrompt string, userInput string) (string, error)
}

type Classifier struct {
	log                  *slog.Logger
	client               Client
	allowedTags          []string
	classifySystemPrompt string
}

func New(log *slog.Logger, client Client, allowedTags []string) *Classifier {
	classifier := &Classifier{
		log:         log,
		client:      client,
		allowedTags: allowedTags,
	}
	classifier.initSystemPrompt()
	return classifier
}

func (c *Classifier) initSystemPrompt() {
	c.classifySystemPrompt = fmt.Sprintf(
		"Ты диспетчер. Прочитай сообщение и верни ТОЛЬКО JSON формата %s. Допустимые теги: %s. Никакого другого текста писать нельзя.",
		classifyResponseExampleStr, strings.Join(c.allowedTags, ","),
	)
}

func (c *Classifier) Classify(ctx context.Context, userInput string) (string, error) {
	const op = "classifier.Classifier.Classify"

	if userInput == "" {
		return "", fmt.Errorf("%s: user input is empty", op)
	}

	aiResponse, err := c.client.DoRequest(ctx, c.classifySystemPrompt, userInput)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var response СlassifyResponse
	if err := json.Unmarshal([]byte(aiResponse), &response); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	tag := response.Category

	if !slices.Contains(c.allowedTags, tag) {
		return "", fmt.Errorf("%s: tag %s not supported", op, tag)
	}

	return tag, nil
}
