package api

import (
	"context"

	"github.com/clintjedwards/comet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (api *API) CreateComet(context context.Context, request *proto.CreateCometRequest) (*proto.CreateCometResponse, error) {

	// Validate user input
	if request.Name == "" {
		return &proto.CreateCometResponse{}, status.Error(codes.FailedPrecondition, "name requried")
	}

	return &proto.CreateCometResponse{}, nil
}
