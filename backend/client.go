package backend

import (
	"context"

	"github.com/clintjedwards/comet/backend/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// GRPCClient represents the implementation for a client that can talk to plugins
type GRPCClient struct{ client proto.BackendPluginClient }

// GRPCClient is the client implementation that allows our host to send RPCs to plugins
func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewBackendPluginClient(c)}, nil
}

// Below are wrappers for how plugins should respond to the RPC in question
// They are all pretty simple since the general flow is to just call the implementation
// of the rpc method for that specific plugin and return the result

// CreateMachine calls CreateMachine on the plugin through the GRPC client
func (m *GRPCClient) CreateMachine(request *proto.CreateMachineRequest) (*proto.CreateMachineResponse, error) {
	response, err := m.client.CreateMachine(context.Background(), request)
	if err != nil {
		return &proto.CreateMachineResponse{}, err
	}
	return response, nil
}
