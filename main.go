package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", DeleteUser).Methods("DELETE")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
