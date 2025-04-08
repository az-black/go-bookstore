package main

import (
	"github.com/az-black/bookstore/pkg/config"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/az-black/bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import dengan underscore karena hanya untuk init
)

func main() {
	// Inisialisasi koneksi DB lebih awal
	config.Connect()

	// Inisialisasi router
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)

	// Jalankan server
	log.Println("Server berjalan di http://localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
