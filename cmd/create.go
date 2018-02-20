package cmd

import (
	"github.com/kube-ops/pops/stack"

	"github.com/spf13/cobra"
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
		if err := currStack.Create(); err != nil {
			panic(err)
		}
	},
}

func init() {
	createCmd.AddCommand(createStackCmd)

	rootCmd.AddCommand(createCmd)
}
