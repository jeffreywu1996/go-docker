package main

import (
	"encoding/json"
	"fmt"
	"net/http"
  "log"

	"github.com/gorilla/mux"
)

// Struct (our models)
type User struct {
	ID  string `json:"user_id"`
	AGE int32  `json:"age"`
}

var users []User

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, user := range users {
		if user.ID == params["user_id"] {
			fmt.Println(params["user_id"])
			json.NewEncoder(w).Encode(user)
      log.Println("User get", user.ID)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)

	w.WriteHeader(http.StatusNoContent)
  log.Println("User added", user.ID)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, user := range users {
		if user.ID == params["user_id"] {
			users = append(users[:index], users[index+1:]...)
      log.Println("User deleted", user.ID)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
