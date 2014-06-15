package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to lister HTTP requests on")
)

func main() {
	flag.Parse()

	fmt.Println("Starting Dispatcher")
	StartDispatcher(*NWorkers)

	fmt.Println("Registering the collectior")
	http.HandleFunc("/work", Collector)

	fmt.Println("HTTP server lintening on: ", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
