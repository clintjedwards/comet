package backend

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/clintjedwards/comet/backend/proto"
	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/utils"
	"github.com/hashicorp/go-getter"
	"github.com/hashicorp/go-plugin"
)

const (
	golangBinaryName = "go"
	pluginBinaryName = "backend"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "COMET_PLUGIN",
	MagicCookieValue: "cKykOnGDBJ",
}

// PluginDefinition is the interface in which both the plugin and the host has to implement
type PluginDefinition interface {
	CreateMachine(request *proto.CreateMachineRequest) (*proto.CreateMachineResponse, error)
	GetPluginInfo(request *proto.GetPluginInfoRequest) (*proto.GetPluginInfoResponse, error)
}

// Plugin is just a wrapper so we implement the correct go-plugin interface
// it allows us to serve/consume the plugin
type Plugin struct {
	plugin.Plugin
	Impl PluginDefinition
}

// PluginExists checks the plugin directory to see if we already have a built version
// of the plugin we want
func PluginExists() bool {
	config, _ := config.FromEnv()

	info, err := os.Stat(fmt.Sprintf("%s/%s", config.Backend.BinaryPath, pluginBinaryName))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// GetPluginRaw is used to retrieve a plugin from either a repo or local path.
// Should be able to download from most common sources. (eg: git, http, mercurial)
// See (https://github.com/hashicorp/go-getter#url-format) for more information
// on how to form input
func GetPluginRaw(location string) error {
	config, err := config.FromEnv()
	if err != nil {
		return err
	}

	err = getter.GetAny(fmt.Sprintf("%s/%s", config.Backend.RepoPath, pluginBinaryName), location)
	return err
}

// BuildPlugin builds the plugin from srcPath and stores it in dstPath
// with the provided name
// id refers to the unique hash of the plugin
func BuildPlugin() error {
	config, err := config.FromEnv()
	if err != nil {
		return err
	}

	fullBinaryPath := fmt.Sprintf("%s/%s", config.Backend.BinaryPath, pluginBinaryName)

	buildArgs := []string{"build", "-o", fullBinaryPath}

	golangBinaryPath, err := exec.LookPath(golangBinaryName)
	if err != nil {
		return err
	}

	// go build <args> <path_to_plugin_src_files>
	_, err = utils.ExecuteCmd(golangBinaryPath, buildArgs, nil, config.Backend.RepoPath)
	if err != nil {
		return err
	}

	return nil
}
