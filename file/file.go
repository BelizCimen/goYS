package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

//Readfile is...
func ReadFile() {
	fmt.Println("Scheduled save file job is starting")
	timer1 := time.NewTimer(5 * time.Minute)
	done := make(chan bool)
	go func() {
		<-timer1.C
		r, e := http.Get("localhost:1111/keys")
		if e != nil {
			panic(e)
		}
		defer r.Body.Close()
		f, e := os.Create("tmp/TIMESTAMP-data.json")
		if e != nil {
			panic(e)
		}
		defer f.Close()
		io.Copy(f, r.Body)
		fmt.Println("data saved into TIMESTAMP-data.json")
		done <- true
	}()
	<-done
}
