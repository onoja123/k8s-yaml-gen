package applier

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Apply runs 'kubectl apply -f -' with the given manifest content.
func Apply(manifest string) error {
	cmd := exec.Command("kubectl", "apply", "-f", "-")
	cmd.Stdin = strings.NewReader(manifest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("kubectl apply failed: %w", err)
	}
	return nil
}
