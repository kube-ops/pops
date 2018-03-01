package cmd

import (
	"github.com/kube-ops/pops/image"
	"github.com/spf13/cobra"
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
		image.PrintList(dockersDir)
	},
}

func init() {
	listImageCmd.Flags().StringVarP(&dockersDir, "dockers-dir", "d", "/tmp/test", "Directory where docker definition folders can be found")
	listCmd.AddCommand(listImageCmd)
	rootCmd.AddCommand(listCmd)
}
