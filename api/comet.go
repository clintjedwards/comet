package api

import (
	"fmt"
	"os/exec"

	"github.com/clintjedwards/comet/backend"
	backendProto "github.com/clintjedwards/comet/backend/proto"
	"github.com/clintjedwards/comet/proto"
	"github.com/hashicorp/go-plugin"
)

func (api *API) spawnComet(request *proto.CreateCometRequest) (*proto.CreateCometResponse, error) {

	_ = proto.Comet{}

	pluginPath := fmt.Sprintf("%s/%s", api.config.Backend.PluginDirectoryPath, backend.PluginBinaryName)

	// We create a client so that we can communicate with backend plugin
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  backend.Handshake,
		Plugins:          api.backendPlugin,
		Cmd:              exec.Command(pluginPath),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return &proto.CreateCometResponse{}, nil
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(backend.PluginBinaryName)
	if err != nil {
		return &proto.CreateCometResponse{}, nil
	}

	backend := raw.(backend.PluginDefinition)
	_, err = backend.CreateMachine(&backendProto.CreateMachineRequest{})
	if err != nil {
		return &proto.CreateCometResponse{}, nil
	}

	return &proto.CreateCometResponse{}, nil
}
