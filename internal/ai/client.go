package ai

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// Client wraps the OpenAI API client.
type Client struct {
	cli *openai.Client
}

// NewClient initializes an OpenAI client using the OPENAI_API_KEY env var.
// Panics if the API key is not set.
func NewClient() *Client {
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		panic("OPENAI_API_KEY environment variable not set")
	}

	client := openai.NewClient(apiKey)
	return &Client{cli: client}
}

// DefaultClient is the shared instance used by the application.
var DefaultClient = NewClient()

// GenerateYAML sends the prompt to OpenAI and returns the generated YAML string.
func (c *Client) GenerateYAML(ctx context.Context, prompt string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: "You are a Kubernetes YAML generator. Only respond with valid YAML manifests."},
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		Temperature: 0.2,
	}

	resp, err := c.cli.CreateChatCompletion(ctx, req)

	if err != nil {
		return "", fmt.Errorf("OpenAI request failed: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}
