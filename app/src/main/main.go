package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tvs", getTvs).Methods("GET")
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/tvs/{id}", getTv).Methods("GET")
	r.HandleFunc("/tvs", createTv).Methods("POST")

	port := os.Getenv("go_app_port")
	if port == "" {
		port = "80"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
