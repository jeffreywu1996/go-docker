package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// init mux router
	// type inference :=
	log.Println("Server started")
	r := mux.NewRouter()
	// Route handler / Endpoints
	r.HandleFunc("/user/{user_id}", getUser).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{user_id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
