package backend

import (
	"context"

	"github.com/clintjedwards/comet/backend/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// GRPCServer is the implementation that allows the plugin to respond to requests from the host
type GRPCServer struct {
	Impl PluginDefinition
}

// GRPCServer is the server implementation that allows our plugins to recieve RPCs
func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterBackendPluginServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

// Below are wrappers for how plugins should respond to the RPC in question
// They are all pretty simple since the general flow is to just call the implementation
// of the rpc method for that specific plugin and return the result

// CreateMachine creates a machine and passes relevant details back to comet
func (m *GRPCServer) CreateMachine(ctx context.Context, request *proto.CreateMachineRequest) (*proto.CreateMachineResponse, error) {
	response, err := m.Impl.CreateMachine(request)
	return response, err
}
