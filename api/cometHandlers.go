package api

import (
	"context"
	"time"

	backendProto "github.com/clintjedwards/comet/backend/proto"
	"github.com/clintjedwards/comet/namegen"
	"github.com/clintjedwards/comet/proto"
	"github.com/clintjedwards/comet/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateComet spawns a comet creation request based on user given parameters
func (api *API) CreateComet(ctx context.Context, request *proto.CreateCometRequest) (*proto.CreateCometResponse, error) {

	// Validate user input
	if request.TimeRequested == "" {
		return &proto.CreateCometResponse{},
			status.Error(codes.FailedPrecondition, "comet duration requried")
	}

	if request.Size == 0 {
		return &proto.CreateCometResponse{},
			status.Error(codes.FailedPrecondition, "size required")
	}

	name := request.Name
	if name == "" {
		name = namegen.GenerateName()
	}

	newComet := proto.Comet{
		Id:       string(utils.GenerateRandString(api.config.Comet.IDLength)),
		Name:     name,
		Notes:    request.Notes,
		Size:     proto.Comet_Size(request.Size),
		Status:   proto.Comet_PENDING,
		Created:  time.Now().Unix(),
		Modified: time.Now().Unix(),
		// (TODO) Calculate Deletion time here
		Metadata: request.Metadata,
	}

	machineRequest := &backendProto.CreateMachineRequest{
		Id:       newComet.Id,
		Name:     newComet.Name,
		Size:     backendProto.CreateMachineRequest_Size(newComet.Size),
		Metadata: newComet.Metadata,
	}

	err := api.storage.AddComet(newComet.Id, &newComet)
	if err != nil {
		return &proto.CreateCometResponse{},
			status.Errorf(codes.Internal, "could not create comet:%v", err)
	}

	go api.spawnComet(machineRequest)

	return &proto.CreateCometResponse{Comet: &newComet}, nil
}
