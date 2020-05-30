package api

import (
	backendProto "github.com/clintjedwards/comet/backend/proto"
	"github.com/clintjedwards/comet/proto"
	"github.com/rs/zerolog/log"
)

// spawnComet starts and tracks the creation of a comet
func (api *API) spawnComet(request *backendProto.CreateMachineRequest) {

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

	response, err := api.backendPlugin.CreateMachine(request)
	if err != nil {
		log.Error().Err(err).Interface("comet", updatedComet).Msg("could not create comet")
		updatedComet.Status = proto.Comet_STOPPED
	} else {
		log.Info().Interface("comet", updatedComet).Msg("created new comet")
		updatedComet.Status = proto.Comet_RUNNING
		updatedComet.InstanceId = response.Machine.InstanceId
	}

	err = api.storage.UpdateComet(request.Id, updatedComet)
	if err != nil {
		log.Error().Err(err).Msg("could not update comet info")
		return
	}

	return
}
