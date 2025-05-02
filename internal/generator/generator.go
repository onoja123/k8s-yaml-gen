package generator

import (
	"context"
	"fmt"
	"time"

	"github.com/onoja123/k8s-yaml-gen/internal/ai"
)

// GenerateManifest invokes the AI client to produce the Kubernetes YAML manifest.
func GenerateManifest(opts Options) (string, error) {
	prompt := BuildPrompt(opts)
	// Use a context with timeout to avoid hanging if the API is slow.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	yaml, err := ai.DefaultClient.GenerateYAML(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("AI generation error: %w", err)
	}
	return yaml, nil
}

func (o Options) Validate() error {
	if o.Replicas <= 0 {
		return fmt.Errorf("replicas must be greater than 0")
	}
	if o.Memory == "" {
		return fmt.Errorf("memory must not be empty")
	}
	if o.Port <= 0 || o.Port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}
	return nil
}
