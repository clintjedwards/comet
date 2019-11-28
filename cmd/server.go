package cmd

import (
	"github.com/clintjedwards/comet/app"
	"github.com/spf13/cobra"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Starts the service and runs until an interrupt is received",
	Long:  `Starts the service and runs until an interrupt is received`,
	Run:   runServerCmd,
}

func runServerCmd(cmd *cobra.Command, args []string) {
	app.StartServices()
}

func init() {
	RootCmd.AddCommand(cmdServer)
}
