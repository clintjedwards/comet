package plugin

import (
	"github.com/clintjedwards/comet/backend/proto"
	"github.com/hashicorp/go-plugin"
)

// This file contains structures that both the plugin and the plugin host has to implement

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "COMET_PLUGIN",
	MagicCookieValue: "cKykOnGDBJ",
}

// BackendDefinition is the interface in which both the plugin and the host has to implement
type BackendDefinition interface {
	CreateMachine(request *proto.CreateMachineRequest) (*proto.CreateMachineResponse, error)
}

// BackendPlugin is just a wrapper so we implement the correct go-plugin interface
// it allows us to serve/consume the plugin
type BackendPlugin struct {
	plugin.Plugin
	Impl BackendDefinition
}
