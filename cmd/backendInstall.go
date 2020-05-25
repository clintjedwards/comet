package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/clintjedwards/comet/app"
	"github.com/clintjedwards/comet/backend"
	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/proto"
	"github.com/clintjedwards/comet/storage"
	"github.com/clintjedwards/comet/utils"
	"github.com/theckman/yacspin"

	"github.com/hashicorp/go-getter"
	"github.com/spf13/cobra"

	"github.com/rs/zerolog/log"
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
comet backend install ~/comet/dev-backend
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
	info, err := os.Stat(fmt.Sprintf("%s/%s", path, backend.PluginBinaryName))
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
	err := getter.GetAny(fmt.Sprintf("%s/%s", backend.TmpDir, backend.PluginBinaryName), location)
	return err
}

// buildPlugin builds the plugin from srcPath and stores it in dstPath
// with the provided name
func buildPlugin(path string) ([]byte, error) {
	fullPluginPath := fmt.Sprintf("%s/%s", path, backend.PluginBinaryName)
	tmpPath := fmt.Sprintf("%s/%s", backend.TmpDir, backend.PluginBinaryName)

	buildArgs := []string{"build", "-o", fullPluginPath}

	golangBinaryPath, err := exec.LookPath(backend.GolangBinaryName)
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

func initSpinner(suffix string) (*yacspin.Spinner, error) {
	cfg := yacspin.Config{
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[14],
		Suffix:            " " + suffix,
		SuffixAutoColon:   true,
		StopCharacter:     "✓",
		StopColors:        []string{"fgGreen"},
		StopFailCharacter: "x",
		StopFailColors:    []string{"fgRed"},
	}

	spinner, err := yacspin.New(cfg)
	if err != nil {
		return nil, err
	}

	return spinner, nil
}

func runBackendInstallCmd(cmd *cobra.Command, args []string) {
	spinner, err := initSpinner("installing comet backend")
	if err != nil {
		log.Error().Err(err).Msg("could not init spinner")
		return
	}
	spinner.Message("retrieving configuration")
	spinner.Start()

	config, err := config.FromEnv()
	if err != nil {
		spinner.StopFailMessage(fmt.Sprintf("could not get config: %v", err))
		spinner.StopFail()
		return
	}

	spinner.Message("connecting to storage")

	database, err := app.InitStorage(storage.EngineType(config.Database.Engine))
	if err != nil {
		spinner.StopFailMessage(fmt.Sprintf("could not init storage: %v", err))
		spinner.StopFail()
		return
	}

	location := args[0]
	update, _ := cmd.Flags().GetBool("update")

	if pluginExists(config.Backend.PluginDirectoryPath) && !update {
		spinner.StopFailMessage("backend already exists")
		spinner.StopFail()
		return
	}

	err = createDirectories(backend.TmpDir, config.Backend.PluginDirectoryPath)
	if err != nil {
		spinner.StopFailMessage(fmt.Sprintf("could not create required directories: %v", err))
		spinner.StopFail()
		return
	}

	spinner.Message("downloading backend plugin")

	err = getPluginRaw(location)
	if err != nil {
		spinner.StopFailMessage(fmt.Sprintf("could not get plugin: %v", err))
		spinner.StopFail()
		return
	}

	spinner.Message("building backend plugin")

	output, err := buildPlugin(config.Backend.PluginDirectoryPath)
	if err != nil {
		spinner.StopFailMessage(fmt.Sprintf("could not build plugin: %v\n%s", err, output))
		spinner.StopFail()
		return
	}

	spinner.Message("adding backend to database")

	database.AddBackend(&proto.Backend{
		Location: location,
		Created:  time.Now().Unix(),
		Modified: time.Now().Unix(),
	})

	spinner.Suffix(" installed comet backend")
	spinner.Stop()
}

func init() {
	var update bool
	cmdBackendInstall.Flags().BoolVarP(&update, "update", "u", false,
		"force backend redownload and update")

	cmdBackend.AddCommand(cmdBackendInstall)
}
