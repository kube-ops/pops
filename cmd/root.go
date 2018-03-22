package cmd

import (
	"os"

	"github.com/kube-ops/pops/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var rootCmd = &cobra.Command{
	Use:   "Pops",
	Short: "CLI tool to help build, version and publish Ops files to be deployed in Kubernetes",
	Long: `Pops helps managing the lifecycle of ops files destined to be deployed in Kubernetes.
For now, Pops handles Docker images, and Helm charts only.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(1)
		}
	},
}

func processPersistentFlags() {
	if viper.GetBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}
}

// AddStackPersistentFlags add the common persistent flags for stacks.
func addStackPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("chart-dir", "s", "charts", "Directory containing the Helm charts")
	// nolint: gas
	_ = viper.BindPFlag("chart-dir", cmd.PersistentFlags().Lookup("chart-dir"))
}

// Execute execute the root command.
func Execute() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Activates verbose mode")
	// nolint: gas
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().StringP("out-dir", "o", ".out", "Storage directory for artefacts")
	// nolint: gas
	_ = viper.BindPFlag("out-dir", rootCmd.PersistentFlags().Lookup("out-dir"))

	cobra.OnInitialize(config.InitializeConfig)
	cobra.OnInitialize(processPersistentFlags)
	log.SetLevel(log.WarnLevel)

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
