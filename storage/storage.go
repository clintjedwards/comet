package storage

import (
	"fmt"

	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/proto"
)

// Bucket represents the name of a section of key/value pairs
// usually a grouping of some sort
// ex. A key/value pair of userid-userdata would belong in the users bucket
type Bucket string

const (
	// CometsBucket represents the container in which comets are managed
	CometsBucket Bucket = "comets"
)

// EngineType represents the different possible storage engines available
type EngineType string

const (
	// StorageEngineBoltDB represents a boltDB storage engine.
	// A file based key-value store.(https://github.com/boltdb/bolt)
	StorageEngineBoltDB EngineType = "boltdb"
)

// Engine represents backend storage implementations where items can be persisted
type Engine interface {
	Init(config *config.Config) error
	GetAllComets() (map[string]*proto.Comet, error)
	GetComet(id string) (*proto.Comet, error)
	AddComet(id string, comet *proto.Comet) error
	UpdateComet(id string, comet *proto.Comet) error
	DeleteComet(id string) error
	AddBackend(backend *proto.Backend) error
	GetBackend() (*proto.Backend, error)
	DeleteBackend() error
}

// InitStorage creates a storage object with the appropriate engine
func InitStorage(engineType EngineType) (Engine, error) {

	switch engineType {
	case StorageEngineBoltDB:
		config, err := config.FromEnv()
		if err != nil {
			return nil, err
		}

		boltDBStorageEngine := boltDB{}
		err = boltDBStorageEngine.Init(config)
		if err != nil {
			return nil, err
		}

		return &boltDBStorageEngine, nil
	default:
		return nil, fmt.Errorf("storage backend not implemented: %s", engineType)
	}
}
