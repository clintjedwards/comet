package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appVersion = "v0.0.dev"

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Comet %s\n", appVersion)
	},
}

func init() {
	RootCmd.AddCommand(cmdVersion)
}
