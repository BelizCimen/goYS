package main

import (
	"fmt"
	"goYS/file"
	"goYS/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/keys", handlers.KeysRouter)
	http.HandleFunc("/keys/", handlers.KeysRouter)
	file.ReadFile()
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
