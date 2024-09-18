package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/index", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	defer mu.Unlock()
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count %d\n", count)
	defer mu.Unlock()
}
