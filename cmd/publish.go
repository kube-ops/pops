package cmd

import (
	"path"

	"github.com/kube-ops/pops/git"
	"github.com/kube-ops/pops/image"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		sourceDir := path.Join(viper.GetString("ProjectRootDir"), viper.GetString("image-dir"))
		img, err := image.NewDockerImageFromPath(sourceDir, args[0])
		if err != nil {
			log.Fatal(err)
		}
		img.Publish()
		git.CreateTag(viper.GetString("ProjectRootDir"), "image-"+img.Name+"-"+img.Version)
	},
}

func init() {
	publishCmd.AddCommand(publishImageCmd)
	addImagePersistentFlags(publishImageCmd)

	rootCmd.AddCommand(publishCmd)
}
