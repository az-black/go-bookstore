package routes

import (
	"github.com/gorilla/mux"
	"go-bookstore/pkg/controllers"

)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.Handle("/books", controllers.CreateBook).Methods("POST")
	router.Handle("/books", controllers.GetBooks).Methods("GET")
	router.Handle("/books/{id}", controllers.GetBook).Methods("GET")
	router.Handle("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.Handle("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}