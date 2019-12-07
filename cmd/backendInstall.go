package cmd

import (
	"fmt"
	"log"

	"github.com/clintjedwards/comet/backend"
	"github.com/spf13/cobra"
)

var cmdBackendInstall = &cobra.Command{
	Use:   "install <location>",
	Short: "Install a comet backend",
	Long: `Backends can be installed from a variety of sources.

Location is where your plugin code is hosted. Valid locations include common
places to host code including github url, fileserver path, or even just
the path to a local directory.

For more information on what formats location accepts see:
https://github.com/hashicorp/go-getter#supported-protocols-and-detectors

Examples:
comet backend install ~/comet/backend/dev-backend
comet backend install github.com/hashicorp/go-getter
`,
	Args: cobra.MaximumNArgs(1),
	Run:  runBackendInstallCmd,
}

func runBackendInstallCmd(cmd *cobra.Command, args []string) {
	location := args[0]
	update, _ := cmd.Flags().GetBool("update")

	if backend.PluginExists() && !update {
		fmt.Println("Backend already exists")
		return
	}

	err := backend.GetPluginRaw(location)
	if err != nil {
		log.Fatalf("Could not get plugin: %v", err)
	}

	err = backend.BuildPlugin()
	if err != nil {
		log.Fatalf("Could not build plugin: %v", err)
	}
}

func init() {
	var update bool
	cmdBackendInstall.Flags().BoolVarP(&update, "update", "u", false,
		"force backend redownload and update")

	cmdBackend.AddCommand(cmdBackendInstall)
}
