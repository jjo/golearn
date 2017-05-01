package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

var port = flag.Int("p", 8000, "Port to listen")

func _init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("http.Request: %s", r)
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
}

func main() {
	_init()
	http.HandleFunc("/", handler)
	log.Printf("Listening on port=%d\n", *port)
	var err = http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	log.Fatalf("FATAL: %s\n", err)
}
