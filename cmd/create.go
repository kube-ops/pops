package cmd

import (
	"path"

	"github.com/kube-ops/pops/image"
	"github.com/kube-ops/pops/stack"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// buildCmd represents the build command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a skeleton for an artefact.",
	Aliases: []string{"crt"},
	Long: `Create a skeleton for a container image or a stack description.
	These artifacts can be one of:
	- container image
	- stack description`,
}

var createStackCmd = &cobra.Command{
	Use:   "stack STACK",
	Short: "Create a stack description",
	Long: `Create a stack in a directory.
  Only helm charts are supported for now`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		currStack := stack.HelmChart{Name: args[0], Version: ""}
		destDir := path.Join(viper.GetString("ProjectRootDir"), viper.GetString("chart-dir"), args[0])
		if err := currStack.Create(destDir); err != nil {
			log.Fatalf("Couldn't create chart %s.", args[0])
		}
	},
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
		sourceDir := path.Join(viper.GetString("ProjectRootDir"), viper.GetString("image-dir"))
		docker := image.NewDockerImage(args[0], dockerRegistry, dockerTag, "")
		docker.Create(sourceDir)
	},
}

func init() {
	createCmd.AddCommand(createStackCmd)
	addStackPersistentFlags(createStackCmd)
	addImagePersistentFlags(createImageCmd)

	createImageCmd.Flags().StringVarP(&dockerTag, "tag", "t", "", "Docker image tag (Required)")
	err := createImageCmd.MarkFlagRequired("tag")
	if err != nil {
		log.Fatal(err)
	}
	createImageCmd.Flags().StringVarP(&dockerRegistry, "registry", "r", "", "Registry where the docker image will be published (Required)")
	err = createImageCmd.MarkFlagRequired("registry")
	if err != nil {
		log.Fatal(err)
	}
	createCmd.AddCommand(createImageCmd)
	rootCmd.AddCommand(createCmd)
}
