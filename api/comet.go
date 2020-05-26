package api

import (
	"fmt"
	"os/exec"

	"github.com/clintjedwards/comet/backend"
	backendProto "github.com/clintjedwards/comet/backend/proto"
	"github.com/clintjedwards/comet/proto"
	"github.com/hashicorp/go-plugin"
	"github.com/rs/zerolog/log"
)

// spawnComet starts and tracks the creation of a comet
func (api *API) spawnComet(request *backendProto.CreateMachineRequest) {

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
		log.Error().Err(err).Msg("could not connect to backend plugin")
		return
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(backend.PluginBinaryName)
	if err != nil {
		log.Error().Err(err).Msg("could not connect to backend plugin")
		return
	}

	// It should never be possible to get into this state unless there is database corruption
	comet, err := api.storage.GetComet(request.Id)
	if err != nil {
		log.Error().Err(err).Msg("could not get newly created comet")
		return
	}

	updatedComet := &proto.Comet{
		Id:         comet.Id,
		InstanceId: comet.InstanceId,
		Name:       comet.Name,
		Notes:      comet.Notes,
		Size:       comet.Size,
		Address:    comet.Address,
		Created:    comet.Created,
		Modified:   comet.Modified,
		Deletion:   comet.Deletion,
	}

	backend := raw.(backend.PluginDefinition)
	_, err = backend.CreateMachine(request)
	if err != nil {
		log.Error().Err(err).Interface("comet", updatedComet).Msg("could not create comet")
		updatedComet.Status = proto.Comet_STOPPED
	} else {
		log.Info().Interface("comet", updatedComet).Msg("created new comet")
		updatedComet.Status = proto.Comet_RUNNING
	}

	err = api.storage.UpdateComet(request.Id, updatedComet)
	if err != nil {
		log.Error().Err(err).Msg("could not update comet info")
		return
	}

	return
}
