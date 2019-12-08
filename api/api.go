package api

import (
	"log"
	"net"

	"github.com/clintjedwards/comet/backend"
	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/proto"
	"github.com/clintjedwards/comet/storage"
	"github.com/clintjedwards/comet/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/hashicorp/go-plugin"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

// API represents the grpc backend service
type API struct {
	config *config.Config
	//search  *search.Search
	storage       storage.Engine
	backendPlugin map[string]plugin.Plugin
	proto.UnimplementedCometAPIServer
}

// NewAPI inits a grpc api service
func NewAPI(config *config.Config) *API {

	backendPlugin := map[string]plugin.Plugin{
		"backend": &backend.Plugin{},
	}

	storage, err := storage.InitStorage(storage.EngineType(config.Database.Engine))
	if err != nil {
		utils.Log().Panicf("could not init storage: %v", err)
	}

	// searchIndex, err := search.InitSearch()

	// go searchIndex.BuildIndex()

	return &API{
		backendPlugin: backendPlugin,
		config:        config,
		storage:       storage,
	}
}

// CreateGRPCServer creates a grpc server with all the proper settings; TLS enabled
func CreateGRPCServer(api *API) *grpc.Server {

	creds, err := credentials.NewServerTLSFromFile(api.config.TLSCertPath, api.config.TLSKeyPath)
	if err != nil {
		utils.Log().Panicf("failed to get certificates: %v", err)
	}

	serverOption := grpc.Creds(creds)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
		)),
		serverOption,
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
		utils.Log().Panicf("could not init tcp listener", err)
	}

	utils.Log().Infow("starting comet grpc service",
		"url", config.Comet.GRPCURL)

	log.Fatal(server.Serve(listen))
}
