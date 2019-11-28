package storage

import (
	"fmt"
	"time"

	"github.com/clintjedwards/comet/utils"

	"github.com/golang/protobuf/proto"

	"github.com/boltdb/bolt"
	"github.com/clintjedwards/comet/api"
	"github.com/clintjedwards/comet/config"
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

	err = boltDB.createBuckets(PipelinesBucket)
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

func (boltDB *boltDB) GetAllPipelines() (map[string]*api.Pipeline, error) {
	results := map[string]*api.Pipeline{}

	boltDB.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(PipelinesBucket))

		err := bucket.ForEach(func(key, value []byte) error {
			var pipeline api.Pipeline
			err := proto.Unmarshal(value, &pipeline)
			if err != nil {
				utils.StructuredLog(utils.LogLevelError,
					"could not unmarshal pipeline while trying to retrieve all",
					map[string]string{"pipeline_id": string(key), "error": err.Error()})
				return nil
			}
			results[string(key)] = &pipeline
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})

	return results, nil
}

func (boltDB *boltDB) GetPipeline(id string) (*api.Pipeline, error) {

	var storedPipeline api.Pipeline

	err := boltDB.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(PipelinesBucket))

		pipelineRaw := bucket.Get([]byte(id))
		if pipelineRaw == nil {
			return utils.ErrEntityNotFound
		}

		err := proto.Unmarshal(pipelineRaw, &storedPipeline)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &storedPipeline, nil
}

func (boltDB *boltDB) AddPipeline(id string, pipeline *api.Pipeline) error {
	err := boltDB.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(PipelinesBucket))

		// First check if key exists
		currentPipeline := bucket.Get([]byte(id))
		if currentPipeline != nil {
			return utils.ErrEntityExists
		}

		pipelineRaw, err := proto.Marshal(pipeline)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), pipelineRaw)
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

func (boltDB *boltDB) UpdatePipeline(id string, pipeline *api.Pipeline) error {
	err := boltDB.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(PipelinesBucket))

		// First check if key exists
		currentPipeline := bucket.Get([]byte(id))
		if currentPipeline == nil {
			return utils.ErrEntityNotFound
		}

		pipelineRaw, err := proto.Marshal(pipeline)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(id), pipelineRaw)
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

func (boltDB *boltDB) DeletePipeline(id string) error {
	err := boltDB.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(PipelinesBucket))

		// First check if key exists
		currentPipeline := bucket.Get([]byte(id))
		if currentPipeline == nil {
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
