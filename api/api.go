package api

import (
	"net"

	"github.com/clintjedwards/comet/backend"
	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/proto"
	"github.com/clintjedwards/comet/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hashicorp/go-plugin"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/rs/zerolog/log"
)

// API represents the grpc backend service
type API struct {
	config        *config.Config
	storage       storage.Engine
	backendPlugin map[string]plugin.Plugin
	// we add this so we aren't forced to immediately implement all methods
	// for a valid api server
	proto.UnimplementedCometAPIServer
}

// NewAPI inits a grpc api service
func NewAPI(config *config.Config, storage storage.Engine) *API {

	// Declare an empty backend plugin so that we can init it later
	backendPlugin := map[string]plugin.Plugin{
		"backend": &backend.Plugin{},
	}

	return &API{
		backendPlugin: backendPlugin,
		config:        config,
		storage:       storage,
	}
}

// CreateGRPCServer creates a grpc server with all the proper settings; TLS enabled
func CreateGRPCServer(api *API) *grpc.Server {

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)

	grpc_prometheus.EnableHandlingTimeHistogram()

	reflection.Register(grpcServer)
	grpc_prometheus.Register(grpcServer)
	proto.RegisterCometAPIServer(grpcServer, api)

	return grpcServer
}

// InitGRPCService starts a GPRC server
func InitGRPCService(config *config.Config, server *grpc.Server) {

	listen, err := net.Listen("tcp", config.Comet.GRPCURL)
	if err != nil {
		log.Panic().Err(err).Msg("could not init tcp listener")
	}

	log.Info().Str("url", config.Comet.GRPCURL).Msg("starting comet grpc service")

	log.Fatal().Err(server.Serve(listen)).Msg("server exited abnormally")
}
