package app

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/clintjedwards/comet/api"
	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/metrics"
	"github.com/clintjedwards/comet/storage"
	"github.com/clintjedwards/comet/storage/bolt"
)

// StartServices initializes all required services,the raw GRPC service, and the metrics endpoint
func StartServices() {
	config, err := config.FromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("could not get config in order to start services")
	}

	storage, err := InitStorage(storage.EngineType(config.Database.Engine))
	if err != nil {
		log.Panic().Err(err).Msg("could not init storage")
	}

	cometAPI := api.NewAPI(config, storage)
	grpcServer := api.CreateGRPCServer(cometAPI)

	go metrics.InitPrometheusService(config)
	api.InitGRPCService(config, grpcServer)
}

// InitStorage creates a storage object with the appropriate engine
func InitStorage(engineType storage.EngineType) (storage.Engine, error) {

	config, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	switch engineType {
	case storage.BoltEngine:

		boltStorageEngine, err := bolt.Init(config.Database.Bolt)
		if err != nil {
			return nil, err
		}

		return &boltStorageEngine, nil
	default:
		return nil, fmt.Errorf("storage backend not implemented: %s", engineType)
	}
}
