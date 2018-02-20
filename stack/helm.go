package stack

import (
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
func (helm *HelmChart) Create() error {
	// nolint: gas
	// Disabling warning from gas caused by a variable in parameter.
	cmd := exec.Command("helm", "create", helm.Name)
	_, err := cmd.CombinedOutput()
	return err
}

// Build builds the Helm chart and validates it.
func (helm *HelmChart) Build() error {
	return nil
}

// Publish sends the Helm chart to a repository.
func (helm *HelmChart) Publish() error {
	return nil
}
