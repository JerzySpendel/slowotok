package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	users = append(users, User{Id: "1", Username: "FirstUser", Score: 0})
	users = append(users, User{Id: "2", Username: "SecondUser", Score: 0})
	users = append(users, User{Id: "3", Username: "ThirdUser", Score: 0})

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/board/{level}", GetBoard).Methods("GET")
	router.HandleFunc("/check_word/", CheckWord).Methods("POST")

	// dictionary, err := ReadWords("dictionaries/words.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(dictionary)

	log.Fatal(http.ListenAndServe(":3000", router))
}
