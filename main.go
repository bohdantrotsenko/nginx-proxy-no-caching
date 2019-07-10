package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var counter int64

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fl := w.(http.Flusher)
		for {
			fmt.Fprintf(w, "%d\n", atomic.AddInt64(&counter, 1))
			fl.Flush()
			time.Sleep(200 * time.Millisecond)
		}
	})

	go http.ListenAndServe(":8080", nil)

	for {
		var placeholder string
		fmt.Scanln(&placeholder)
		log.Printf("counter reset\n")
		atomic.StoreInt64(&counter, 0)
	}
}
