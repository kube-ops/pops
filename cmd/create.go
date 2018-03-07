package cmd

import (
	"path"

	"github.com/kube-ops/pops/stack"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

// buildCmd represents the build command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a skeleton for an artefact.",
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

func init() {
	createCmd.AddCommand(createStackCmd)
	addStackPersistentFlags(createStackCmd)

	rootCmd.AddCommand(createCmd)
}
