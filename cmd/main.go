package main

import (
	"log"
	"net/http"

	"github.com/coderj001/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

const port string = "localhost"
const host string = ":3000"

func main() {
	var r = mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(host+port, r))
}
