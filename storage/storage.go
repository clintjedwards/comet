package storage

import (
	"github.com/clintjedwards/comet/proto"
)

// Bucket represents the name of a section of key/value pairs
// usually a grouping of some sort
// ex. A key/value pair of userid-userdata would belong in the users bucket
type Bucket string

const (
	// CometsBucket represents the container in which comets are managed
	CometsBucket Bucket = "comets"
	// BackendBucket represents the container in which the backend is managed
	BackendBucket Bucket = "backend"
)

// EngineType represents the different possible storage engines available
type EngineType string

const (
	// BoltEngine represents a bolt storage engine.
	// A file based key-value store.(https://github.com/boltdb/bolt)
	BoltEngine EngineType = "bolt"
)

// Engine represents backend storage implementations where items can be persisted
type Engine interface {
	GetAllComets() (map[string]*proto.Comet, error)
	GetComet(id string) (*proto.Comet, error)
	AddComet(id string, comet *proto.Comet) error
	UpdateComet(id string, comet *proto.Comet) error
	DeleteComet(id string) error
	AddBackend(backend *proto.Backend) error
	GetBackend() (*proto.Backend, error)
	DeleteBackend() error
}
