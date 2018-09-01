package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"strings"
	"github.com/gorilla/mux"
)

func Shuffle(src []string) []string {
 	final := make([]string, len(src))
 	rand.Seed(time.Now().UTC().UnixNano())
 	perm := rand.Perm(len(src))

 	for i, v := range perm {
 		final[v] = src[i]
 	}
 	return final
 }

func GetBoard(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	params := mux.Vars(r)
	i, err := strconv.ParseFloat(params["level"], 64)
	if err != nil {
		fmt.Println(err)
	}
	level := int(math.Pow(i, 2))

	words, err := ReadWords("dictionaries/game-words.txt")
	if err != nil {
		fmt.Println(err)
	}

	var board = ""
	x := 0
	for x < level {
		board += words[rand.Intn(len(words))]
		x++
	}

	asdfboard := strings.Split(board, "")

	json.NewEncoder(w).Encode(Shuffle(asdfboard[:level]))
}
