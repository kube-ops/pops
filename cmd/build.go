package cmd

import (
	"fmt"

	"github.com/kube-ops/pops/image"
	"github.com/sirupsen/logrus"
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

var dockersDir string
var buildImageCmd = &cobra.Command{
	Use:   "image IMAGE",
	Short: "Build a container image",
	Long: `Build a container image.
  Only docker images are supported for now`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		img, err := image.NewDockerImageFromPath(dockersDir, args[0])
		if err != nil {
			logrus.Fatal(err)
		}
		img.Build()
	},
}

var buildStackCmd = &cobra.Command{
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
	buildImageCmd.Flags().StringVarP(&dockersDir, "dockers-dir", "d", "/tmp/test", "Directory where docker definition folders can be found")
	buildCmd.AddCommand(buildImageCmd)
	buildCmd.AddCommand(buildStackCmd)

	rootCmd.AddCommand(buildCmd)
}
