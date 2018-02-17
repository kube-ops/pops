package image

import (
	"fmt"
)

// Docker docker image type
type Docker struct {
	Name     string
	Registry string
	Tag      string
}

// NewDocker create a new docker image
func NewDocker(name string, registry string, tag string) *Docker {
	return &Docker{name, registry, tag}
}

// Print print docker attributes
func (docker *Docker) Print() {
	fmt.Printf("name: %s, registry: %s, tag: %s", docker.Name, docker.Registry, docker.Tag)

}
