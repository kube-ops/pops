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
	Short: "Create a develoment branch and bump version",
	Long: `Create a develoment branch named according to the artifact type, name and version.
	It bumps the version too
	These artifacts can be one of:
	- container image
	- stack description`,
}

var major, minor, patch bool
var improveImageCmd = &cobra.Command{
	Use:   "image IMAGE",
	Short: "Create a develoment branch and bump version",
	Long: `Create a develoment branch named according to the artifact type, name and version.
	It bumps the version too`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir := path.Join(viper.GetString("ProjectRootDir"), viper.GetString("image-dir"))
		img, err := image.NewDockerImageFromPath(sourceDir, args[0])
		if err != nil {
			log.Fatal(err)
		}
		if major {
			img.BumpVersion("major")
		} else if minor {
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

	improveImageCmd.Flags().BoolVarP(&major, "major", "M", false, "Make a braking change")
	improveImageCmd.Flags().BoolVarP(&minor, "minor", "m", false, "Add functionality")
	improveImageCmd.Flags().BoolVarP(&patch, "patch", "p", true, "Make a bug fix")

	rootCmd.AddCommand(improveCmd)
}
