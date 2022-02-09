package routes

import (
	"github.com/coderj001/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Route) {
	router.HandlerFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandlerFunc("/book/{bookid}", controllers.GetBookId).Methods("GET")
	router.HandlerFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandlerFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
}
