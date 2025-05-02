package generator

import (
	"context"
	"fmt"
	"time"

	"github.com/onoja123/k8s-yaml-gen/internal/ai"
)

// Options holds the user-specified parameters for manifest generation.
type Options struct {
	Replicas int
	Memory   string
	Port     int
}

// BuildPrompt constructs the natural-language prompt for the AI model.
func BuildPrompt(opts Options) string {
	return fmt.Sprintf(
		"Generate a Kubernetes Deployment and Service YAML for a web app with %d replicas, %s memory per pod, and NodePort %d. Only output valid YAML.",
		opts.Replicas, opts.Memory, opts.Port,
	)
}

// GenerateManifest invokes the AI client to produce the Kubernetes YAML manifest.
func GenerateManifestFromOptions(opts Options) (string, error) {
	prompt := BuildPrompt(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	yaml, err := ai.DefaultClient.GenerateYAML(ctx, prompt)

	if err != nil {
		return "", fmt.Errorf("AI generation error: %w", err)
	}

	return yaml, nil
}
