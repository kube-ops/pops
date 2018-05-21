package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docCmd generate the command help in markdown
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate the command help in markdown",
	Long:  "Generate the command help in markdown",
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
}
