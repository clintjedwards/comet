package backend

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"

	"github.com/clintjedwards/comet/backend/proto"
	"github.com/clintjedwards/comet/utils"
	"github.com/hashicorp/go-getter"
	"github.com/hashicorp/go-plugin"
)

const (
	golangBinaryName = "go"
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

// pluginExists checks the plugin directory to see if we already have a built version
// of the plugin we want
func pluginExists(pluginPath string) bool {
	info, err := os.Stat(pluginPath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// getPluginRaw is used to retrieve a plugin from either a repo or local path.
// Should be able to download from most common sources. (eg: git, http, mercurial)
// See (https://github.com/hashicorp/go-getter#url-format) for more information
// on how to form input
func getPluginRaw(id, dstPath, location string) error {
	err := getter.GetAny(dstPath, location)
	return err
}

// buildPlugin builds the plugin from srcPath and stores it in dstPath with the provided name
// id refers to the unique hash of the plugin
func buildPlugin(id, srcPath, dstPath string) error {
	buildArgs := []string{"build", "-o", fmt.Sprintf("%s/%s", dstPath, id)}

	golangBinaryPath, err := exec.LookPath(golangBinaryName)
	if err != nil {
		return err
	}

	// go build <args> <path_to_plugin_src_files>
	_, err = utils.ExecuteCmd(golangBinaryPath, buildArgs, nil, srcPath)
	if err != nil {
		return err
	}

	return nil
}

func importBackendPlugin(repoDir, binaryDir, location string) error {
	id := getMD5Hash(location)
	err := getPluginRaw(id, repoDir, location)
	if err != nil {
		return err
	}
	err = buildPlugin(id, repoDir, binaryDir)
	if err != nil {
		return err
	}

	return nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
