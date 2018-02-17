package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "Pops",
	Short: "CLI tool to help build, version and publish Ops files to be deployed in Kubernetes",
	Long: `Pops helps managing the lifecycle of ops files destined to be deployed in Kubernetes.
For now, Pops handles Docker images, and Helm charts only.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute execute the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
