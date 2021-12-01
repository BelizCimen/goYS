package handlers

import (
	"encoding/json"
	"errors"
	"goYS/model"
	"io/ioutil"
	"net/http"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

func bodyToKey(r *http.Request, k *model.Key) error {
	if r.Body == nil {
		return errors.New("REQUEST BODY EMPTY")
	}
	if r.Body == nil {
		return errors.New("KEY EMPTY")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, k)
}
func keysGetAll(w http.ResponseWriter, r *http.Request) {
	keys, err := model.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"keys ": keys})
	return
}

func keysPostOne(w http.ResponseWriter, r *http.Request) {
	k := new(model.Key)
	err := bodyToKey(r, k)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	k.ID = bson.NewObjectId()
	err = k.Save()
	if err != nil {
		if err == model.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/keys/"+k.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}

func keysGetOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	k, err := model.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"key": k})
}

func keysPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	k := new(model.Key)
	err := bodyToKey(r, k)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	k.ID = id
	err = k.Save()
	if err != nil {
		if err == model.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"key": k})
}
func keysDeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := model.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
