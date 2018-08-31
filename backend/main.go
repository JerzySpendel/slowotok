package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"bufio"
	"os"
	"math"
	"strconv"
	"fmt"
	"math/rand"
	"time"
)


type User struct {
	Id string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Score int16 `json:"score,omitempty"`
}


var charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var users []User


func main() {
	router := mux.NewRouter()

	users = append(users, User{Id: "1", Username: "FirstUser", Score: 0})
	users = append(users, User{Id: "2", Username: "SecondUser", Score: 0})
	users = append(users, User{Id: "3", Username: "ThirdUser", Score: 0})

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/board/{level}", getBoard).Methods("GET")

	dictionary, err := readWords("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func readWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func getBoard(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	params := mux.Vars(r)
	i, err := strconv.ParseFloat(params["level"], 64)
	if err != nil {
		fmt.Println(err)
	}
	level := int(math.Pow(i, 2))

	board := make([]byte, level)
	// fmt.Println(board)

	words, err := readWords("game-words.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(words)

	for a := 0; a < len(board); a++ {
		board[a] = byte(65 + rand.Intn(25))
	}

	json.NewEncoder(w).Encode(string(board))
}
