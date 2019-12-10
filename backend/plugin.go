package backend

import (
	"github.com/clintjedwards/comet/backend/proto"
	"github.com/hashicorp/go-plugin"
)

const (
	// GolangBinaryName is typically used in searching for the binary on file systems
	GolangBinaryName string = "go"
	// PluginBinaryName is the name of the backend plugin once compiled into a binary
	PluginBinaryName string = "backend"
	// TmpDir is the directory where we download plugin src files to before compiling
	TmpDir string = "/tmp"
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
