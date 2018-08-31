package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


type User struct {
	Id string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}


var users []User


func main() {
	router := mux.NewRouter()

	users = append(users, User{Id: "1", Username: "FirstUser"})
	users = append(users, User{Id: "2", Username: "SecondUser"})
	users = append(users, User{Id: "3", Username: "ThirdUser"})

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}