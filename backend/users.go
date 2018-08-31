package main


import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)


type User struct {
	Id string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Score int16 `json:"score,omitempty"`
}

var users []User


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
