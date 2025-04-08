package routes

import (
	"github.com/gorilla/mux"
	"github.com/az-black/go-bookstore/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/v1/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/v1/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/v1/books/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/v1/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/v1/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}