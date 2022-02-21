package db

import (
	"github.com/boltdb/bolt"
	"github.com/eslerkang/tjcoin/utils"
)

var db *bolt.DB

type Bucket string

const (
	DB_NAME      = "blockchain.db"
	DATA_BUCKET  = Bucket("data")
	BLOCK_BUCKET = Bucket("blocks")
	CHECKPOINT   = "checkpoint"
)

func DB() *bolt.DB {
	if db == nil {
		// initialize db
		dbPointer, err := bolt.Open(DB_NAME, 0600, nil)
		utils.HandleError(err)
		db = dbPointer
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(DATA_BUCKET))
			utils.HandleError(err)
			_, err = t.CreateBucketIfNotExists([]byte(BLOCK_BUCKET))
			utils.HandleError(err)
			return nil
		})
		utils.HandleError(err)
	}
	return db
}

func Close() {
	DB().Close()
}

func SaveInBucket(bucketName Bucket, key string, data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(bucketName))
		err := bucket.Put([]byte(key), data)
		return err
	})
	utils.HandleError(err)
}

func CheckPoint() []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(DATA_BUCKET))
		data = bucket.Get([]byte(CHECKPOINT))
		return nil
	})
	return data
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(BLOCK_BUCKET))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}
