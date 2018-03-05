package cmd

import (
	"github.com/kube-ops/pops/image"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create artifacts",
	Aliases: []string{"crt"},
	Long: `List artifact
	These artifacts can be one of:
	- container image
	- stack description`,
}

var dockerRegistry string
var dockerTag string
var createImageCmd = &cobra.Command{
	Use:     "image IMAGE",
	Short:   "Create container images",
	Aliases: []string{"images", "img", "im", "i"},
	Args:    cobra.ExactArgs(1),
	Long: `Create container images directory layout in dockers-dir.
  Only docker images are supported for now`,
	Run: func(cmd *cobra.Command, args []string) {
		docker := image.NewDockerImage(args[0], dockerRegistry, dockerTag, "")
		docker.Create(dockersDir)
	},
}

func init() {
	createImageCmd.Flags().StringVarP(&dockersDir, "dockers-dir", "d", "/tmp/test", "Directory where docker definition folders can be found")
	createImageCmd.Flags().StringVarP(&dockerTag, "tag", "t", "", "Tag og the docker image (Required)")
	err := createImageCmd.MarkFlagRequired("tag")
	if err != nil {
		logrus.Fatal(err)
	}
	createImageCmd.Flags().StringVarP(&dockerRegistry, "registry", "r", "", "Registry where the docker image will be published (Required)")
	err = createImageCmd.MarkFlagRequired("registry")
	if err != nil {
		logrus.Fatal(err)
	}
	createCmd.AddCommand(createImageCmd)
	rootCmd.AddCommand(createCmd)
}
