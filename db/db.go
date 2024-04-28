package db

import (
	bolt "go.etcd.io/bbolt"
)

type Database struct {
	db *bolt.DB
}

func NewDatabase(path string) (db *Database, closeFunc func() error, err error) {
	boltDb, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, nil, err
	}
	closeFunc = boltDb.Close
	return &Database{db: boltDb}, closeFunc, nil
}
