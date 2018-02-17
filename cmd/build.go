package cmd

import (
	"fmt"

	"github.com/kube-ops/pops/image"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build an artifact to be deployed on Kubernetes",
	Long: `Build an artifact to be deployed on Kubernetes
	These artifacts can be one of:
	- container image
	- stack description`,
}

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Build a container image",
	Long: `Build a container image.
  Only docker images are supported for now`,
	Run: func(cmd *cobra.Command, args []string) {
		img := image.NewDocker("nginx", "kube-ops", "1.0")
		img.Print()
	},
}

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Build a stack artifact",
	Long: `Build a stack artifact.
	Creates a tgz of the stack description.
  Only helm charts are supported for now`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building helm chart")
	},
}

func init() {
	buildCmd.AddCommand(imageCmd)
	buildCmd.AddCommand(stackCmd)

	rootCmd.AddCommand(buildCmd)
}
