package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//KeysRouter is
func KeysRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/keys" {
		switch r.Method {
		case http.MethodGet:
			keysGetAll(w, r)
			return
		case http.MethodPost:
			keysPostOne(w, r)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/keys/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		keysGetOne(w, r, id)
		return
	case http.MethodPost:
		keysPostOne(w, r)
		return
	case http.MethodDelete:
		keysDeleteOne(w, r, id)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
