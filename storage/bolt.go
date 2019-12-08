package storage

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/clintjedwards/comet/config"
	"github.com/clintjedwards/comet/proto"
	"github.com/clintjedwards/comet/utils"
	go_proto "github.com/golang/protobuf/proto"
)

type boltDB struct {
	filePath string
	store    *bolt.DB
}

func (boltDB *boltDB) Init(config *config.Config) error {

	db, err := bolt.Open(config.Database.BoltDB.Path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	boltDB.store = db

	err = boltDB.createBuckets(CometsBucket)
	if err != nil {
		return err
	}

	return nil
}

func (boltDB *boltDB) createBuckets(names ...Bucket) error {

	for _, name := range names {
		err := boltDB.store.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return fmt.Errorf("could not create bucket: %s; %v", name, err)
			}

			return nil
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func (boltDB *boltDB) GetAllComets() (map[string]*proto.Comet, error) {
	results := map[string]*proto.Comet{}

	boltDB.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(CometsBucket))

		err := bucket.ForEach(func(key, value []byte) error {
			var comet proto.Comet
			err := go_proto.Unmarshal(value, &comet)
			if err != nil {
				utils.Log().Errorf("could not unmarshal database object",
					"id", key,
					"error", err)
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

func (boltDB *boltDB) GetComet(id string) (*proto.Comet, error) {

	var storedComet proto.Comet

	err := boltDB.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(CometsBucket))

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

func (boltDB *boltDB) AddComet(id string, comet *proto.Comet) error {
	err := boltDB.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(CometsBucket))

		// First check if key exists
		currentComet := bucket.Get([]byte(id))
		if currentComet != nil {
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

func (boltDB *boltDB) UpdateComet(id string, comet *proto.Comet) error {
	err := boltDB.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(CometsBucket))

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

func (boltDB *boltDB) DeleteComet(id string) error {
	err := boltDB.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(CometsBucket))

		// First check if key exists
		currentComet := bucket.Get([]byte(id))
		if currentComet == nil {
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

func (boltDB *boltDB) AddBackend(backend *proto.Backend) error {

	return nil
}

func (boltDB *boltDB) GetBackend() (*proto.Backend, error) {

	return nil, nil
}

func (boltDB *boltDB) DeleteBackend() error {

	return nil
}
