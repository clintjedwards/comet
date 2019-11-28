package api

import (
	"log"
	"net"

	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/search"
	"github.com/clintjedwards/comet/storage"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// API represents the grpc backend service
type API struct {
	storage storage.Engine
	config  *config.Config
	search  *search.Search
}

// NewAPI inits a grpc api service
func NewAPI(config *config.Config) *API {
	newAPI := API{}

	// storage, err := storage.InitStorage()
	// if err != nil {
	// 	utils.StructuredLog(utils.LogLevelFatal, "failed to initialize storage", err)
	// }

	// searchIndex, err := search.InitSearch()
	// if err != nil {
	// 	utils.StructuredLog(utils.LogLevelFatal, "failed to initialize search functions", err)
	// }

	// go searchIndex.BuildIndex()

	// basecoatAPI.config = config
	// basecoatAPI.storage = storage
	// basecoatAPI.search = searchIndex

	return &newAPI
}

// CreateGRPCServer creates a grpc server with all the proper settings; TLS enabled
func CreateGRPCServer(cometAPI *API) *grpc.Server {

	creds, err := credentials.NewServerTLSFromFile(cometAPI.config.TLSCertPath, cometAPI.config.TLSKeyPath)
	if err != nil {
		utils.StructuredLog(utils.LogLevelFatal, "failed to get certificates", err)
	}

	serverOption := grpc.Creds(creds)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(cometAPI.authenticate),
			grpc_prometheus.UnaryServerInterceptor,
		)),
		serverOption,
	)

	grpc_prometheus.EnableHandlingTimeHistogram()

	reflection.Register(grpcServer)
	grpc_prometheus.Register(grpcServer)
	api.RegisterCometServer(grpcServer, cometAPI)

	return grpcServer
}

// InitGRPCService starts a GPRC server
func InitGRPCService(config *config.Config, server *grpc.Server) {

	listen, err := net.Listen("tcp", config.Backend.GRPCURL)
	if err != nil {
		utils.StructuredLog(utils.LogLevelFatal, "could not initialize tcp listener", err)
	}

	utils.StructuredLog(utils.LogLevelInfo, "starting basecoat grpc service",
		map[string]string{"url": config.Backend.GRPCURL})

	log.Fatal(server.Serve(listen))
}
