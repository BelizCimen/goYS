package model

import (
	"errors"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

//Key is
type Key struct {
	ID    bson.ObjectId `json:"id" storm:"id"`
	Key   string        `json:"key"`
	Value string        `json:"value"`
}

const (
	dbPath = "model.db"
)

// ErrRecordInvalid for invalid record alert
var (
	ErrRecordInvalid = errors.New("INVALID RECORD")
)

//All is
func All() ([]Key, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}

	defer db.Close()
	keys := []Key{}
	err = db.All(&keys)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

//One is
func One(id bson.ObjectId) (*Key, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	k := new(Key)
	err = db.One("ID", id, k)
	if err != nil {
		return nil, err
	}
	return k, nil
}

//Delete is
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	k := new(Key)
	err = db.One("ID", id, k)
	if err != nil {
		return err
	}
	return db.DeleteStruct(k)
}

//Save is
func (k *Key) Save() error {
	if err := k.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(k)
}

//Validate is
func (k *Key) validate() error {
	if k.Key == "" {
		return ErrRecordInvalid
	}
	return nil
}
