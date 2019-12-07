package cmd

import "github.com/spf13/cobra"

var cmdBackend = &cobra.Command{
	Use:   "backend",
	Short: "Controls operations around comet's backend",
	Long: `Comet allows users to use different backends for machine procurement.

A backend might be a cloud provider, docker container, or anything that
might serve as a conduit for creating machines.

To start using comet you must first install a backend using the install command.
`,
}

func init() {
	RootCmd.AddCommand(cmdBackend)
}
