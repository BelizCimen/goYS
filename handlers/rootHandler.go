package handlers

import "net/http"

//RootHandler is...
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("NOT FOUND\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(([]byte("RUNNING API V1\n")))

}
