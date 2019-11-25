package kubernetes

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/pkg/errors"
)

// Apply config using kubectl
func Apply(config string) error {
	cmd := exec.Command("kubectl", "apply", "-f", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "could not create stdin pipe")
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, config)
	}()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, config)
	}
	if !cmd.ProcessState.Success() {
		return errors.Wrap(err, string(output))
	}
	fmt.Print(string(output))
	return nil
}