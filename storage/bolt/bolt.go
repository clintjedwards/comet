package bolt

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/proto"
	"github.com/clintjedwards/comet/storage"
	"github.com/clintjedwards/comet/utils"
	go_proto "github.com/golang/protobuf/proto"
	"github.com/rs/zerolog/log"
)

// Bolt is a representation of the bolt datastore
type Bolt struct {
	store *bolt.DB
}

const backendKey = "backend"

// Init creates a new boltdb with given settings
func Init(configuration interface{}) (Bolt, error) {
	db := Bolt{}
	conf, ok := configuration.(*config.BoltConfig)
	if !ok {
		return Bolt{}, fmt.Errorf("incorrect config type expected 'config.BoltConfig'; got '%T'", configuration)
	}

	store, err := bolt.Open(conf.Path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return Bolt{}, err
	}

	// Create root bucket if not exists
	err = store.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(storage.CometsBucket))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(storage.BackendBucket))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return Bolt{}, err
	}

	db.store = store

	return db, nil
}

// GetAllComets returns an unpaginated list of current links
func (db *Bolt) GetAllComets() (map[string]*proto.Comet, error) {
	results := map[string]*proto.Comet{}

	db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.CometsBucket))

		err := bucket.ForEach(func(key, value []byte) error {
			var comet proto.Comet

			err := go_proto.Unmarshal(value, &comet)
			if err != nil {
				log.Error().Err(err).Str("id", string(key)).Msg("could not unmarshal database object")
				return nil
			}

			results[string(key)] = &comet
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})

	return results, nil
}

// GetComet returns a single comet by id
func (db *Bolt) GetComet(id string) (*proto.Comet, error) {

	var storedComet proto.Comet

	err := db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.CometsBucket))

		cometRaw := bucket.Get([]byte(id))
		if cometRaw == nil {
			return utils.ErrEntityNotFound
		}

		err := go_proto.Unmarshal(cometRaw, &storedComet)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &storedComet, nil
}

// AddComet stores a new comet
func (db *Bolt) AddComet(id string, comet *proto.Comet) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.CometsBucket))

		// First check if key exists
		exists := bucket.Get([]byte(id))
		if exists != nil {
			return utils.ErrEntityExists
		}

		cometRaw, err := go_proto.Marshal(comet)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), cometRaw)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateComet alters comet infromation
func (db *Bolt) UpdateComet(id string, comet *proto.Comet) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.CometsBucket))

		// First check if key exists
		currentComet := bucket.Get([]byte(id))
		if currentComet == nil {
			return utils.ErrEntityNotFound
		}

		cometRaw, err := go_proto.Marshal(comet)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), cometRaw)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteComet removes a comet from the database
func (db *Bolt) DeleteComet(id string) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.CometsBucket))

		// First check if key exists
		exists := bucket.Get([]byte(id))
		if exists == nil {
			return utils.ErrEntityNotFound
		}

		err := bucket.Delete([]byte(id))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// AddBackend adds a configured backend to the database
func (db *Bolt) AddBackend(backend *proto.Backend) error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.BackendBucket))

		// First check if key exists
		exists := bucket.Get([]byte(backendKey))
		if exists != nil {
			return utils.ErrEntityExists
		}

		backendRaw, err := go_proto.Marshal(backend)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(backendKey), backendRaw)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// GetBackend returns the configured backend
func (db *Bolt) GetBackend() (*proto.Backend, error) {

	var storedBackend proto.Backend

	err := db.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.BackendBucket))

		backendRaw := bucket.Get([]byte(backendKey))
		if backendRaw == nil {
			return utils.ErrEntityNotFound
		}

		err := go_proto.Unmarshal(backendRaw, &storedBackend)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &storedBackend, nil
}

// DeleteBackend removes a configured backend
func (db *Bolt) DeleteBackend() error {
	err := db.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(storage.BackendBucket))

		// First check if key exists
		exists := bucket.Get([]byte(backendKey))
		if exists == nil {
			return utils.ErrEntityNotFound
		}

		err := bucket.Delete([]byte(backendKey))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
