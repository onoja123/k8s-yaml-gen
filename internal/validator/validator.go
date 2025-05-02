package validator

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/serializer"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
)

// Validate attempts to decode each document in the manifest and ensures
// they are valid Kubernetes resources of allowed kinds (Deployment or Service).
func Validate(manifest string) error {
	// UniversalDeserializer handles all registered k8s types.
	codecs := serializer.NewCodecFactory(clientgoscheme.Scheme)
	decoder := codecs.UniversalDeserializer()


	docs := strings.Split(manifest, "---")
	for _, doc := range docs {
		raw := strings.TrimSpace(doc)
		if raw == "" {
			continue
		}

		_, gvk, err := decoder.Decode([]byte(raw), nil, nil)

		if err != nil {
			return fmt.Errorf("failed to decode YAML: %w", err)
		}

		// Only allow Deployment and Service.
		switch gvk.Kind {
		case "Deployment", "Service":

		default:
			return fmt.Errorf("unexpected resource kind: %s", gvk.Kind)
		}
	}

	return nil
}
