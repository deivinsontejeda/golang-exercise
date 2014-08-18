package inventory_service

import (
	"encoding/json"
	d "github.com/deivinsontejeda/inventory_service/database"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	ProductGetRoute = "product"
)

func NewAPIRouter() *mux.Router {
	m := mux.NewRouter()
	m.Path("/api/v1/products/{uuid}").Name(ProductGetRoute)

	return m
}

func init() {
	m := NewAPIRouter()
	m.Get(ProductGetRoute).HandlerFunc(handleProductGet).Methods("GET")
	http.Handle("/api/", m)
}

var dsn string = "root@tcp(127.0.0.1:3306)/glive_inventory_development"
var productDataStore d.ProductsService = d.NewConn(dsn)

func handleProductGet(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	product, _ := productDataStore.Get(uuid)
	b, err := json.Marshal(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
