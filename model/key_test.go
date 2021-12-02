package model

import (
	"reflect"
	"testing"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

func TestCRUD(t *testing.T) {

	t.Log("Create")
	k := &Key{
		ID:    bson.NewObjectId(),
		Key:   "key1",
		Value: "value1",
	}

	err := k.Save()
	if err != nil {
		t.Fatalf("Error saving key")
	}

	t.Log("Read")
	k2, err := One(k.ID)
	if err != nil {
		t.Fatalf("Error retrieving key")
	}
	if !reflect.DeepEqual(k2, k) {
		t.Error("Records do not match")
	}

	t.Log("Delete")
	err = Delete(k.ID)
	if err != nil {
		t.Fatalf("Error deleting key")
	}

	_, err = One(k.ID)
	if err == nil {
		t.Fatal("Record should not exist anymore")
	}
	if err != storm.ErrNotFound {
		t.Fatalf("Error retriving non-exist key")
	}
}
