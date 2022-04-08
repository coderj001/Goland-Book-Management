package main

import (
	"log"
	"net/http"

	"github.com/coderj001/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

const port string = ":3000"
const host string = "localhost"

func main() {
	var r = mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(host+port, r))
}
