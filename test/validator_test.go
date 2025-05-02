package generator_test

import (
	"testing"

	"github.com/onoja123/k8s-yaml-gen/internal/generator"
)

func TestBuildPromptValidation(t *testing.T) {
	tests := []struct {
		opts     generator.Options
		expected string
	}{
		{
			opts:     generator.Options{Replicas: 3, Memory: "512Mi", Port: 30080},
			expected: "Generate a Kubernetes Deployment and Service YAML for a web app with 3 replicas, 512Mi memory per pod, and NodePort 30080. Only output valid YAML.",
		},
		{
			opts:     generator.Options{Replicas: 1, Memory: "256Mi", Port: 31000},
			expected: "Generate a Kubernetes Deployment and Service YAML for a web app with 1 replicas, 256Mi memory per pod, and NodePort 31000. Only output valid YAML.",
		},
	}

	for _, tt := range tests {
		prompt := generator.BuildPrompt(tt.opts)
		if prompt != tt.expected {
			t.Errorf("BuildPrompt() = %q, want %q", prompt, tt.expected)
		}
	}
}

func TestOptionsValidation(t *testing.T) {
	tests := []struct {
		opts    generator.Options
		wantErr bool
	}{
		{opts: generator.Options{Replicas: 2, Memory: "128Mi", Port: 30080}, wantErr: false},
		{opts: generator.Options{Replicas: 0, Memory: "128Mi", Port: 30080}, wantErr: true},
		{opts: generator.Options{Replicas: 1, Memory: "", Port: 30080}, wantErr: true},
		{opts: generator.Options{Replicas: 1, Memory: "64Mi", Port: 0}, wantErr: true},
		{opts: generator.Options{Replicas: 1, Memory: "64Mi", Port: 70000}, wantErr: true},
	}

	for _, tt := range tests {
		err := tt.opts.Validate()
		if (err != nil) != tt.wantErr {
			t.Errorf("Options.Validate() with opts=%+v, wantErr=%v, gotErr=%v", tt.opts, tt.wantErr, err != nil)
		}
	}
}
