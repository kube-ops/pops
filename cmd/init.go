package cmd

import (
	"github.com/kube-ops/pops/config"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// buildCmd represents the build command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Pops for current project",
	Long: `Initialize Pops for the current project (cwd):
	- creates a default configuration file (if not exist)
	- Initialize Helm
	- Initialize the Helm repository`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.SafeWriteConfig(); err != nil {
			log.Fatal("Failed to write configuration file: ", err)
		}

		// TODO Initialize Helm
		// TODO Initialize Helm repository
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
