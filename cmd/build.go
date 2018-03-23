package cmd

import (
	"path"

	"github.com/kube-ops/pops/image"
	"github.com/kube-ops/pops/stack"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
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

var buildImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Build a container image",
	Long: `Build a container image.
  Only docker images are supported for now`,
	Run: func(cmd *cobra.Command, args []string) {
		img := image.NewDocker("nginx", "kube-ops", "1.0")
		img.Print()
	},
}

var buildStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Build a stack artifact",
	Long: `Build a stack artifact.
	Creates a tgz of the stack description.
  Only helm charts are supported for now`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		currStack := stack.HelmChart{Name: args[0], Version: ""}
		sourceDir := path.Join(viper.GetString("ProjectRootDir"), viper.GetString("chart-dir"), args[0])
		destDir := path.Join(viper.GetString("ProjectRootDir"), viper.GetString("out-dir"))
		if err := currStack.Build(sourceDir, destDir); err != nil {
			log.Fatalf("Couldn't build chart %s.", args[0])
		}
	},
}

func init() {
	buildCmd.AddCommand(buildImageCmd)

	buildCmd.AddCommand(buildStackCmd)
	addStackPersistentFlags(buildStackCmd)

	rootCmd.AddCommand(buildCmd)
}
