package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/onoja123/k8s-yaml-gen/internal/applier"
	"github.com/onoja123/k8s-yaml-gen/internal/generator"
	"github.com/onoja123/k8s-yaml-gen/internal/validator"
)

var (
	replicas int
	memory   string
	port     int
	apply    bool
)

var rootCmd = &cobra.Command{
	Use:   "k8s-yaml-gen",
	Short: "Generate Kubernetes YAML manifests via AI",
	Long: `k8s-yaml-gen is a CLI tool that generates Kubernetes YAML manifests
based on plain-English specifications using the OpenAI API.`,
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Kubernetes manifest",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := generator.Options{
			Replicas: replicas,
			Memory:   memory,
			Port:     port,
		}

		manifest, err := generator.GenerateManifest(opts)
		
		if err != nil {
			return fmt.Errorf("failed to generate manifest: %w", err)
		}

		if err := validator.Validate(manifest); err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}

		fmt.Println(manifest)

		if apply {
			return applier.Apply(manifest)
		}
		return nil
	},
}

func init() {
	generateCmd.Flags().IntVarP(&replicas, "replicas", "r", 1, "Number of replicas")
	generateCmd.Flags().StringVarP(&memory, "memory", "m", "512Mi", "Memory per pod (e.g. 512Mi)")
	generateCmd.Flags().IntVarP(&port, "port", "p", 30080, "NodePort (e.g. 30080)")
	generateCmd.Flags().BoolVar(&apply, "apply", false, "If set, run kubectl apply on the generated manifest")

	rootCmd.AddCommand(generateCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
