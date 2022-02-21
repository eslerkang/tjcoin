package db

import (
	"github.com/boltdb/bolt"
	"github.com/eslerkang/tjcoin/utils"
)

var db *bolt.DB

type Bucket string

const (
	dbName       = "blockchain.db"
	Databucket   = Bucket("data")
	BlocksBucket = Bucket("blocks")
)

func DB() *bolt.DB {
	if db == nil {
		// initialize db
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		utils.HandleError(err)
		db = dbPointer
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(Databucket))
			utils.HandleError(err)
			_, err = t.CreateBucketIfNotExists([]byte(BlocksBucket))
			utils.HandleError(err)
			return nil
		})
		utils.HandleError(err)
	}
	return db
}

func SaveInBucket(bucketName Bucket, key string, data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(bucketName))
		err := bucket.Put([]byte(key), data)
		return err
	})
	utils.HandleError(err)
}