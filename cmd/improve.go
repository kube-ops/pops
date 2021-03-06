package cmd

import (
	"path"

	"github.com/kube-ops/pops/git"
	"github.com/kube-ops/pops/image"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// improveCmd represents the improve command
var improveCmd = &cobra.Command{
	Use:   "improve",
	Short: "Create a development branch and bump version",
	Long: `Create a development branch named according to the artifact type, name and version.
	It bumps the version too
	These artifacts can be one of:
	- container image
	- stack description`,
}

var bumpMajor, bumpMinor, bumpPatch bool
var improveImageCmd = &cobra.Command{
	Use:   "image IMAGE",
	Short: "Create a development branch and bump version",
	Long: `Create a development branch named according to the artifact type, name and version.
	It bumps the version too`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir := path.Join(viper.GetString("ProjectRootDir"), viper.GetString("image-dir"))
		img, err := image.NewDockerImageFromPath(sourceDir, args[0])
		if err != nil {
			log.Fatal(err)
		}
		if bumpMajor {
			img.BumpVersion("major")
		} else if bumpMinor {
			img.BumpVersion("minor")
		} else {
			img.BumpVersion("patch")
		}
		git.CreateBranch(viper.GetString("ProjectRootDir"), "image-"+img.Name+"-"+img.Version)
		img.SaveToFile(sourceDir)
	},
}

func init() {
	improveCmd.AddCommand(improveImageCmd)
	addImagePersistentFlags(improveImageCmd)

	improveImageCmd.Flags().BoolVarP(&bumpMinor, "major", "M", false, "Make a breaking change")
	improveImageCmd.Flags().BoolVarP(&bumpMinor, "minor", "m", false, "Add functionality")
	improveImageCmd.Flags().BoolVarP(&bumpPatch, "patch", "p", true, "Make a bug fix")

	rootCmd.AddCommand(improveCmd)
}
