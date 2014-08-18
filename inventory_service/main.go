package main

import (
	"github.com/deivinsontejeda/inventory_service/api"
	"log"
	"net/http"
)

var apiRouter = inventory_service.NewAPIRouter()

func main() {
	bind := ":7777"
	log.Printf("Listening on %s", bind)
	log.Println(http.ListenAndServe(bind, nil))
}
