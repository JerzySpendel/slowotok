package main


import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"math"
	"strconv"
	"fmt"
	"math/rand"
	"time"
)


func GetBoard(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	params := mux.Vars(r)
	i, err := strconv.ParseFloat(params["level"], 64)
	if err != nil {
		fmt.Println(err)
	}
	level := int(math.Pow(i, 2))

	board := make([]byte, level)
	// fmt.Println(board)

	words, err := ReadWords("game-words.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(words)

	for a := 0; a < len(board); a++ {
		board[a] = byte(65 + rand.Intn(25))
	}

	json.NewEncoder(w).Encode(string(board))
}
