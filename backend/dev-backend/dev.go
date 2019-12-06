// Package dev defines a development backend for comet
package dev

import (
	"github.com/hashicorp/go-plugin"

	backendPlugin "github.com/clintjedwards/comet/backend"
	proto "github.com/clintjedwards/comet/backend/proto"
	"time"
)

// Backend is a single implementation of the backend plugin for comet
type backend struct{}

func (*backend) GetPluginInfo(request *proto.GetPluginInfoRequest) (*proto.GetPluginInfoResponse, error) {
	return &proto.GetPluginInfoResponse{
		Id: "dev",
	}, nil
}

// CreateMachine creates a fake machine
func (*backend) CreateMachine(request *proto.CreateMachineRequest) (*proto.CreateMachineResponse, error) {
	time.Sleep(time.Second * 5)
	return &proto.CreateMachineResponse{}, nil
}

func main() {

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: backendPlugin.Handshake,
		Plugins: map[string]plugin.Plugin{
			// the key here doesn't matter
			"dev": &backendPlugin.Plugin{Impl: &backend{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})

}
