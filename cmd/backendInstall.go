package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/proto"
	"github.com/clintjedwards/comet/storage"
	"github.com/clintjedwards/comet/utils"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/hashicorp/go-getter"
	"github.com/spf13/cobra"
)

const (
	golangBinaryName = "go"
	pluginBinaryName = "backend"
	tmpDir           = "/tmp"
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

// createDirectroies attempts to create the needed directories to store plugins and repositories
func createDirectories(directories ...string) error {

	for _, path := range directories {

		_, err := os.Stat(path)

		if os.IsNotExist(err) {
			err := os.MkdirAll(path, 0755)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}
	return nil
}

// pluginExists checks the plugin directory to see if we already have a built version
// of the plugin we want
// path points to the directory where plugins are stored
func pluginExists(path string) bool {
	info, err := os.Stat(fmt.Sprintf("%s/%s", path, pluginBinaryName))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// getPluginRaw is used to retrieve a plugin from either a repo or local path.
// Should be able to download from most common sources. (eg: git, http, mercurial)
// See (https://github.com/hashicorp/go-getter#url-format) for more information
// on how to form input
func getPluginRaw(location string) error {
	err := getter.GetAny(fmt.Sprintf("%s/%s", tmpDir, pluginBinaryName), location)
	return err
}

// buildPlugin builds the plugin from srcPath and stores it in dstPath
// with the provided name
// id refers to the unique hash of the plugin
func buildPlugin(path string) ([]byte, error) {
	fullPluginPath := fmt.Sprintf("%s/%s", path, pluginBinaryName)
	tmpPath := fmt.Sprintf("%s/%s", tmpDir, pluginBinaryName)

	buildArgs := []string{"build", "-o", fullPluginPath}

	golangBinaryPath, err := exec.LookPath(golangBinaryName)
	if err != nil {
		return nil, err
	}

	// go build <args> <path_to_plugin_src_files>
	output, err := utils.ExecuteCmd(golangBinaryPath, buildArgs, nil, tmpPath)
	if err != nil {
		return output, err
	}

	return output, nil
}

func runBackendInstallCmd(cmd *cobra.Command, args []string) {
	loading := wow.New(os.Stdout, spin.Get(spin.Dots), " Initializing assets")
	loading.Start()

	config, err := config.FromEnv()
	if err != nil {
		log.Fatalf("Could not get config: %v", err)
	}

	database, err := storage.InitStorage(storage.EngineType(config.Database.Engine))
	if err != nil {
		log.Fatalf("could not init storage: %v", err)
	}

	location := args[0]
	update, _ := cmd.Flags().GetBool("update")

	if pluginExists(config.Backend.Path) && !update {
		fmt.Println("Backend already exists")
		return
	}

	err = createDirectories(tmpDir, config.Backend.Path)
	if err != nil {
		log.Fatalf("Could not create required directories: %v", err)
	}

	loading.Text(" Downloading backend plugin")

	err = getPluginRaw(location)
	if err != nil {
		log.Fatalf("Could not get plugin: %v", err)
	}

	loading.Text(" Building backend plugin")

	output, err := buildPlugin(config.Backend.Path)
	if err != nil {
		log.Fatalf("\nCould not build plugin: %v\n%s", err, output)
	}

	database.AddBackend(&proto.Backend{
		Location: location,
		Created:  time.Now().Unix(),
		Modified: time.Now().Unix(),
	})

	loading.PersistWith(spin.Spinner{}, "Backend Installed!")
}

func init() {
	var update bool
	cmdBackendInstall.Flags().BoolVarP(&update, "update", "u", false,
		"force backend redownload and update")

	cmdBackend.AddCommand(cmdBackendInstall)
}
