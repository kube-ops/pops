package cmd

import (
	"path"

	"github.com/kube-ops/pops/image"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// buildCmd represents the build command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List artifacts",
	Aliases: []string{"ls"},
	Long: `List artifact
	These artifacts can be one of:
	- container image
	- stack description`,
}

var listImageCmd = &cobra.Command{
	Use:     "image",
	Short:   "List container images",
	Aliases: []string{"images", "img", "im", "i"},
	Long: `List container images.
  Only docker images are supported for now`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir := path.Join(viper.GetString("ProjectRootDir"), viper.GetString("image-dir"))
		image.PrintList(sourceDir)
	},
}

func init() {
	listCmd.AddCommand(listImageCmd)
	addImagePersistentFlags(listImageCmd)
	rootCmd.AddCommand(listCmd)
}
