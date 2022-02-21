package db

import (
	"github.com/boltdb/bolt"
	"github.com/eslerkang/tjcoin/utils"
)

var db *bolt.DB

const (
	dbName       = "blockchain.db"
	databucket   = "data"
	blocksBucket = "blocks"
)

func DB() *bolt.DB {
	if db == nil {
		// initialize db
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		utils.HandleError(err)
		db = dbPointer
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(databucket))
			utils.HandleError(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			utils.HandleError(err)
			return err
		})
		utils.HandleError(err)
	}
	return db
}
