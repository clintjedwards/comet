// Package dev defines a development backend for comet
package main

import (
	"github.com/hashicorp/go-plugin"

	"time"

	backendPlugin "github.com/clintjedwards/comet/backend"
	proto "github.com/clintjedwards/comet/backend/proto"
)

// Backend is a single implementation of the backend plugin for comet
type backend struct{}

func (*backend) GetPluginInfo(request *proto.GetPluginInfoRequest) (*proto.GetPluginInfoResponse, error) {
	return &proto.GetPluginInfoResponse{
		Version:  "0.0.1",
		Name:     "development backend",
		Provider: "in-memory",
	}, nil
}

// CreateMachine creates a fake machine
func (*backend) CreateMachine(request *proto.CreateMachineRequest) (*proto.CreateMachineResponse, error) {
	time.Sleep(time.Second * 5)
	return &proto.CreateMachineResponse{
		Machine: &proto.Machine{
			InstanceId: "fakeinstanceid",
			Address:    "127.0.0.1",
			Metadata: map[string]string{
				"dev": "backend",
			},
		},
		StatusMessage: "some special message about how this deployment went",
	}, nil
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
