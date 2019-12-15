package api

import (
	"context"
	"strings"

	"github.com/clintjedwards/comet/proto"
)

var appVersion = "v0.0.dev <commit>"

// GetSystemInfo returns system information and health
func (api *API) GetSystemInfo(context context.Context, request *proto.GetSystemInfoRequest) (*proto.GetSystemInfoResponse, error) {

	versionTuple := strings.Split(appVersion, " ")

	return &proto.GetSystemInfoResponse{
		DebugEnabled:   api.config.Debug,
		Version:        versionTuple[0],
		Commit:         versionTuple[1],
		DatabaseEngine: api.config.Database.Engine,
	}, nil
}
