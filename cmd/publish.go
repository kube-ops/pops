package cmd

import (
	"github.com/kube-ops/pops/image"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish an artifact to its repository",
	Long: `Publish an artifact to its repository
	These artifacts can be one of:
	- container image
	- stack description`,
}

var publishImageCmd = &cobra.Command{
	Use:   "image IMAGE",
	Short: "Publish a container image to its repository",
	Long: `Publish a container image to its repository
  Only docker images are supported for now`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		img, err := image.NewDockerImageFromPath(dockersDir, args[0])
		if err != nil {
			logrus.Fatal(err)
		}
		img.Publish()
	},
}

func init() {
	publishImageCmd.Flags().StringVarP(&dockersDir, "dockers-dir", "d", "/tmp/test", "Directory where docker definition folders can be found")
	publishCmd.AddCommand(publishImageCmd)

	rootCmd.AddCommand(publishCmd)
}
