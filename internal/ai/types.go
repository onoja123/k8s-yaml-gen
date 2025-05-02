package ai

import openai "github.com/sashabaranov/go-openai"

// ChatModel wraps the OpenAI model identifier.
type ChatModel string

const (
	// GPT35Turbo specifies the gpt-3.5-turbo model.
	GPT35Turbo ChatModel = openai.GPT3Dot5Turbo
	
	// GPT4 specifies the GPT-4 model (if available in your API plan).
	GPT4       ChatModel = openai.GPT4
)

// DefaultSystemPrompt guides the AI to output only valid Kubernetes YAML manifests.
const DefaultSystemPrompt = "You are a Kubernetes YAML generator. Only respond with valid YAML manifests."
