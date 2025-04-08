package controllers

import (
	"encoding/json"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// contoh handler
	w.Write([]byte("CreateBook"))
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetBooks"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetBook"))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateBook"))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteBook"))
}
