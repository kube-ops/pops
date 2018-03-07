package stack

import (
	log "github.com/sirupsen/logrus"
	"net/url"
	"os/exec"
)

// HelmChart is an struct representing a Helm chart data.
type HelmChart struct {
	Name    string
	Version string
}

// HelmRepository is an struct representing an Helm Repository data.
type HelmRepository struct {
	Name string
	URI  url.URL
}

// Create creates an Helm chart skeleton.
func (helm *HelmChart) Create(directory string) error {
	// nolint: gas
	// Disabling warning from gas caused by a variable in parameter.
	cmd := exec.Command("helm", "create", directory)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Errorf("\"helm create\" failed. Command output:\n%s", output)
		return err
	}

	return nil
}

// Build builds the Helm chart and validates it.
func (helm *HelmChart) Build() error {
	return nil
}

// Publish sends the Helm chart to a repository.
func (helm *HelmChart) Publish(sourceFile string) error {
	return nil
}
