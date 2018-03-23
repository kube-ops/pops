package cluster

import "fmt"

// Cluster cluster definition
type Cluster struct {
	Authorization     string `yaml:"authorization"`
	CloudProvider     string `yaml:"cloudProvider"`
	InstanceImage     string `yaml:"instanceImage"`
	KubernetesVersion string `yaml:"kubernetesVersion"`
	MasterSize        string `yaml:"masterSize"`
	NodeSize          string `yaml:"nodeSize"`
}

func main() {
	fmt.Println("vim-go")
}
