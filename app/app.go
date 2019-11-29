package app

import (
	"log"

	"github.com/clintjedwards/comet/api"
	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/metrics"
)

// StartServices initializes a GRPC-web compatible webserver and a GPRC service
func StartServices() {
	config, err := config.FromEnv()
	if err != nil {
		log.Fatal(err)
	}

	cometAPI := api.NewAPI(config)
	grpcServer := api.CreateGRPCServer(cometAPI)

	go metrics.InitPrometheusService(config)
	api.InitGRPCService(config, grpcServer)
}